package controllers

/*
 * The docker API controller to access docker unix socket and reponse JSON data
 *
 * Refer to https://docs.docker.com/reference/api/docker_remote_api_v1.14/ for docker remote API
 * Refer to https://github.com/Soulou/curl-unix-socket to know how to access unix docket
 */

import (
	"context"
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"
)

func SendRequest(address, method string) string {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	images, err := cli.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		panic(err)
	}

	for _, image := range images {
		fmt.Println(image.ID)
	}
	j, _ := json.Marshal(images)
	return string(j)

}

/* Give address and method to request docker unix socket */
func RequestUnixSocket(address, method string) string {
	DOCKER_UNIX_SOCKET := "unix:///var/run/docker.sock"
	// Example: unix:///var/run/docker.sock:/images/json?since=1374067924
	unix_socket_url := DOCKER_UNIX_SOCKET + ":" + address
	u, err := url.Parse(unix_socket_url)
	if err != nil || u.Scheme != "unix" {
		fmt.Println("Error to parse unix socket url " + unix_socket_url)
		return ""
	}

	hostPath := strings.Split(u.Path, ":")
	u.Host = hostPath[0]
	u.Path = hostPath[1]

	conn, err := net.Dial("unix", u.Host)
	if err != nil {
		fmt.Println("Error to connect to", u.Host, err)
		return ""
	}

	reader := strings.NewReader("")
	query := ""
	if len(u.RawQuery) > 0 {
		query = "?" + u.RawQuery
	}

	request, err := http.NewRequest(method, u.Path+query, reader)
	if err != nil {
		fmt.Println("Error to create http request", err)
		return ""
	}

	client := httputil.NewClientConn(conn, nil)
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error to achieve http request over unix socket", err)
		return ""
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error, get invalid body in answer")
		return ""
	}

	/* An example to get continual stream from events, but it's for stdout
		_, err = io.Copy(os.Stdout, res.Body)
		if err != nil && err != io.EOF {
			fmt.Println("Error, get invalid body in answer")
			return ""
	   }
	*/

	defer response.Body.Close()

	return string(body)
}

/* It's a beego controller */
type DockerApiController struct {
	beego.Controller
}

/* Wrap docker remote API to get contaienrs */
func (this *DockerApiController) GetContainers() {
	address := "/containers/json"
	var all int

	this.Ctx.Input.Bind(&all, "all")
	address = address + "?all=" + strconv.Itoa(1)
	result := RequestUnixSocket(address, "GET")
	this.Ctx.WriteString(result)
}

/* Wrap docker remote API to get data of contaienr */
func (this *DockerApiController) GetContainer() {
	id := this.GetString(":id")
	address := "/containers/" + id + "/json"
	result := RequestUnixSocket(address, "GET")
	this.Ctx.WriteString(result)
}

/* Wrap docker remote API to get container's status */
func (this *DockerApiController) TopContainer() {
	id := this.GetString(":id")
	address := "/containers/" + id + "/top"
	result := RequestUnixSocket(address, "GET")
	this.Ctx.WriteString(result)
}

/* Wrap docker remote API to start contaienrs */
func (this *DockerApiController) StartContainer() {
	id := this.GetString(":id")
	address := "/containers/" + id + "/start"
	result := RequestUnixSocket(address, "POST")
	this.Ctx.WriteString(result)
}

/* Wrap docker remote API to stop contaienrs */
func (this *DockerApiController) StopContainer() {
	id := this.GetString(":id")
	address := "/containers/" + id + "/stop"
	result := RequestUnixSocket(address, "POST")
	this.Ctx.WriteString(result)
}

/* Wrap docker remote API to delete contaienrs */
func (this *DockerApiController) DeleteContainer() {
	id := this.GetString(":id")
	address := "/containers/" + id
	result := RequestUnixSocket(address, "DELETE")
	this.Ctx.WriteString(result)
}

/* Wrap docker remote API to get container stats */
func (this *DockerApiController) GetContainerStats() {
	id := this.GetString(":id")
	address := "/containers/" + id + "/stats?stream=False"
	result := RequestUnixSocket(address, "GET")
	this.Ctx.WriteString(result)
}

/* Wrap docker remote API to get images */
func (this *DockerApiController) GetImages() {
	address := "/images/json"
	result := SendRequest(address, "GET")
	this.Ctx.WriteString(result)
}

/* Wrap docker remote API to get data of image */
func (this *DockerApiController) GetImage() {
	id := this.GetString(":id")
	address := "/images/" + id + "/json"
	result := RequestUnixSocket(address, "GET")
	this.Ctx.WriteString(result)
}

/* Wrap docker remote API to get data of user image */
func (this *DockerApiController) GetUserImage() {
	user := this.GetString(":user")
	repo := this.GetString(":repo")
	address := "/images/" + user + "/" + repo + "/json"
	result := RequestUnixSocket(address, "GET")
	this.Ctx.WriteString(result)
}

/* Wrap docker remote API to delete image */
func (this *DockerApiController) DeleteImage() {
	id := this.GetString(":id")
	address := "/images/" + id
	result := RequestUnixSocket(address, "DELETE")
	this.Ctx.WriteString(result)
}

/* Wrap docker remote API to get version info */
func (this *DockerApiController) GetVersion() {
	address := "/version"
	result := RequestUnixSocket(address, "GET")
	this.Ctx.WriteString(result)
}

/* Wrap docker remote API to get docker info */
func (this *DockerApiController) GetInfo() {
	address := "/info"
	result := RequestUnixSocket(address, "GET")
	this.Ctx.WriteString(result)
}

/* Wrap docker remote API to get search images */
func (this *DockerApiController) GetSearchImages() {
	address := "/images/search"
	var term string
	this.Ctx.Input.Bind(&term, "term")
	address = address + "?term=" + term
	result := RequestUnixSocket(address, "GET")
	this.Ctx.WriteString(result)
}

/* Todo: Implement events API, the response is a stream so can't use ioutil.ReadAll() which will be blocked
func (this *DockerapiController) GetEvents() {
	address := "/events"
	var since int
	this.Ctx.Input.Bind(&since, "since")
	address = address + "?since=" + strconv.Itoa(since)
	result := RequestUnixSocket(address, "GET")
	this.Ctx.WriteString(result)
}
*/

/* Warp docker remote API to ping docker daemon */
func (this *DockerApiController) Ping() {
	address := "/_ping"
	result := RequestUnixSocket(address, "GET")
	this.Ctx.WriteString(result)
}

func (this *DockerApiController) PauseContainer() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	id := this.GetString(":id")
	err = cli.ContainerPause(ctx, id)
	if err != nil {
		panic(err)
		this.Ctx.WriteString("FAIL")
	}
	this.Ctx.WriteString("SUCCESS")
}

func (this *DockerApiController) UnpauseContainer() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	id := this.GetString(":id")
	err = cli.ContainerUnpause(ctx, id)
	if err != nil {
		panic(err)
		this.Ctx.WriteString("FAIL")
	}
	this.Ctx.WriteString("SUCCESS")
}

func (this *DockerApiController) KillContainer() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	id := this.GetString(":id")
	var signal string
	err = cli.ContainerKill(ctx, id, signal)
	if err != nil {
		panic(err)
		this.Ctx.WriteString("FAIL")
	}

	this.Ctx.WriteString("SUCCESS")
}
