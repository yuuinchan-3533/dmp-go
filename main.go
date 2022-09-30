package main

import (
	_ "docker-management-platform/routers"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
)

func init() {
	//跨域配置
	beego.InsertFilter("/*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "DELETE"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

}

func main() {

	beego.Run()
}
