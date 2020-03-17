package main

import (
	"boot_beego/bootconfig"
	_ "boot_beego/routers"
	"github.com/astaxie/beego"
)

func main() {

	bootconfig.InitEnv()

	initBeegoConfig()

	beego.Run()
}

func initBeegoConfig() {
	//配置全局
	//beego.BConfig.WebConfig.Session.SessionOn=true
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 86400 //设置Session有效期,单位秒

	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"

}
