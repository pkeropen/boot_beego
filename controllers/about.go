package controllers

import (
	"boot_beego/utils"
	"fmt"
	"github.com/astaxie/beego"
)

type AboutController struct {
	BaseController
}

func (c *AboutController) Get() {

	var cfg = beego.AppConfig

	about := fmt.Sprintf("AppName : %s , RunModel : %s", cfg.String("appname"), cfg.String("runmode"))

	utils.ReturnHTTPSuccess(&c.Controller, about)

	c.ServeJSON()
	c.StopRun()
}
