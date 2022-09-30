package controllers

import (
	"context"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"time"
)

type DockerContainerController struct {
	beego.Controller
}

func (this *DockerContainerController) KillContainer() {
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

func (this *DockerContainerController) RestartContainer() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	id := this.GetString(":id")
	var duration time.Duration = 60 * 60
	err = cli.ContainerRestart(ctx, id, &duration)
	if err != nil {
		panic(err)
		this.Ctx.WriteString("FAIL")
	}
	this.Ctx.WriteString("SUCCESS")
}

func (this *DockerContainerController) PauseContainer() {
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

func (this *DockerContainerController) UnpauseContainer() {
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

func (this *DockerContainerController) RemoveContainer() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	id := this.GetString(":id")
	err = cli.ContainerRemove(ctx, id, types.ContainerRemoveOptions{})
	if err != nil {
		panic(err)
		this.Ctx.WriteString("FAIL")
	}
	this.Ctx.WriteString("SUCCESS")
}
