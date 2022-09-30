package controllers

import (
	"docker-management-platform/models/req"
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"log"
	"os/exec"
	"time"
)

type DockerSecurityController struct {
	beego.Controller
}

func (this *DockerSecurityController) ScanImage() {
	var req req.PullImageReq
	data := this.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &req)
	if err != nil {
		panic(err)
		this.Ctx.WriteString("FAIL")
	}

	timeUnix := time.Now().Unix()
	filePath := "/Users/yuuinchan/trivyOutput/"
	fileName := string(timeUnix) + ".json"
	command := "trivy image -f json" + " -o " + filePath + fileName + " " + req.ImageName

	cmd := exec.Command("/bin/sh", "-c", command)
	err = cmd.Start()
	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	err = cmd.Wait()
	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	
	fmt.Printf("combined out:\n%s\n", string(out))
	this.Ctx.WriteString(string(out))

}
