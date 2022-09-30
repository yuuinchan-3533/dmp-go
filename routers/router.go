package routers

import (
	"docker-management-platform/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	/* Pass to Angular router */
	beego.Router("/", &controllers.MainController{})
	beego.Router("/containers", &controllers.MainController{})
	beego.Router("/containers/:id", &controllers.MainController{})
	beego.Router("/images", &controllers.MainController{})
	beego.Router("/images/:id", &controllers.MainController{})
	beego.Router("/images/:user/:repo", &controllers.MainController{})
	beego.Router("/configuration", &controllers.MainController{})
	beego.Router("/dockerhub", &controllers.MainController{})

	/* HTTP API for docker remote API */
	beego.Router("/dockerapi/containers/json", &controllers.DockerApiController{}, "get:GetContainers")
	//beego.Router("/dockerapi/containers/all", &controllers.DockerApiController{}, "get:GetContainerList")

	/* HTTP API for docker container API */
	beego.Router("/dockerapi/containers/:id/json", &controllers.DockerApiController{}, "get:GetContainer")
	beego.Router("/dockerapi/containers/:id/top", &controllers.DockerApiController{}, "get:TopContainer")
	beego.Router("/dockerapi/containers/:id/start", &controllers.DockerApiController{}, "post:StartContainer")
	beego.Router("/dockerapi/containers/:id/stop", &controllers.DockerApiController{}, "post:StopContainer")
	beego.Router("/dockerapi/containers/:id/pause", &controllers.DockerApiController{}, "post:PauseContainer")
	beego.Router("/dockerapi/containers/:id/kill", &controllers.DockerApiController{}, "post:KillContainer")

	beego.Router("/dockerapi/containers/:id/resume", &controllers.DockerContainerController{}, "post:UnpauseContainer")
	beego.Router("/dockerapi/containers/:id/remove", &controllers.DockerContainerController{}, "post:RemoveContainer")

	beego.Router("/dockerapi/containers/:id", &controllers.DockerApiController{}, "delete:DeleteContainer")
	beego.Router("/dockerapi/containers/:id/stats", &controllers.DockerApiController{}, "get:GetContainerStats")

	/* HTTP API for docker images remote API */
	beego.Router("/dockerapi/images/json", &controllers.DockerApiController{}, "get:GetImages")
	beego.Router("/dockerapi/images/:id/json", &controllers.DockerApiController{}, "get:GetImage")
	beego.Router("/dockerapi/images/:user/:repo/json", &controllers.DockerApiController{}, "get:GetUserImage")
	beego.Router("/dockerapi/images/:id", &controllers.DockerApiController{}, "delete:DeleteImage")
	beego.Router("/dockerapi/version", &controllers.DockerApiController{}, "get:GetVersion")
	beego.Router("/dockerapi/info", &controllers.DockerApiController{}, "get:GetInfo")
	beego.Router("/dockerapi/images/search", &controllers.DockerApiController{}, "get:GetSearchImages")

	beego.Router("/dockerapi/images/pull", &controllers.DockerImageController{}, "post:PullImage")

	// beego.Router("/dockerapi/events", &controllers.DockerapiController{}, "get:GetEvents") // Not support yet
	beego.Router("/dockerapi/_ping", &controllers.DockerApiController{}, "get:Ping")

	/* HTTP API for docker security API */
	beego.Router("/dockerapi/security/image", &controllers.DockerSecurityController{}, "get:ScanImage")

	/* HTTP API for docker volume API */
	beego.Router("/dockerapi/volumes/json", &controllers.DockerVolumeController{}, "get:GetVolumes")
	beego.Router("/dockerapi/volumes/create", &controllers.DockerVolumeController{}, "post:CreateVolume")
	beego.Router("/dockerapi/volumes/:name", &controllers.DockerVolumeController{}, "delete:RemoveVolume")

}
