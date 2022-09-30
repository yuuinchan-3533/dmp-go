package controllers

import (
	"context"
	"docker-management-platform/models"
	"docker-management-platform/models/common"
	"docker-management-platform/models/req"
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"io"
)

type DockerImageController struct {
	beego.Controller
}

func (this *DockerImageController) PullImage() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	var req req.PullImageReq
	var options types.ImagePullOptions
	var resp models.Response

	data := this.Ctx.Input.RequestBody
	err = json.Unmarshal(data, &req)

	events, err := cli.ImagePull(ctx, req.ImageName, options)
	if err != nil {
		panic(err)
		this.Ctx.WriteString("FAIL")
	}
	d := json.NewDecoder(events)
	var event *common.Event
	for {
		if err := d.Decode(&event); err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
	}

	if event != nil {
		resp.Code = 200
		resp.Msg = event.Status
	}
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err)
		return
	}
	this.Ctx.WriteString(string(jsonResp))
}
