package controllers

import (
	"context"
	"docker-management-platform/models"
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
)

type DockerVolumeController struct {
	beego.Controller
}

func (this *DockerVolumeController) GetVolumes() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	var filters filters.Args
	volumeList, err := cli.VolumeList(ctx, filters)
	if err != nil {
		panic(err)
		this.Ctx.WriteString("FAIL")
	}
	jsonVolumeList, err := json.Marshal(volumeList.Volumes)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	fmt.Println(string(jsonVolumeList))
	this.Ctx.WriteString(string(jsonVolumeList))
}

func (this *DockerVolumeController) CreateVolume() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	var options volume.VolumeCreateBody
	data := this.Ctx.Input.RequestBody
	err = json.Unmarshal(data, &options)
	if err != nil {
		panic(err)
	}
	options.DriverOpts["device"] = "/storage/newVolume/" + options.Name
	volume, err := cli.VolumeCreate(ctx, options)
	if err != nil {
		panic(err)
		this.Ctx.WriteString("FAIL")
	}
	jsonVolume, err := json.Marshal(volume)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	fmt.Println(string(jsonVolume))
	this.Ctx.WriteString(string(jsonVolume))
}

func (this *DockerApiController) GetNewVolume() {
	user := this.GetString(":user")
	repo := this.GetString(":repo")
	address := "/images/" + user + "/" + repo + "/json"
	result := RequestUnixSocket(address, "GET")
	this.Ctx.WriteString(result)
}

func (this *DockerVolumeController) RemoveVolume() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	volumeId := this.GetString(":name")

	err = cli.VolumeRemove(ctx, volumeId, false)
	if err != nil {
		panic(err)
		this.Ctx.WriteString("FAIL")
	}
	var resp models.Response
	resp.Code = 200
	resp.Msg = "ID:" + volumeId + " is removed successful"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err)
		return
	}
	this.Ctx.WriteString(string(jsonResp))
}
