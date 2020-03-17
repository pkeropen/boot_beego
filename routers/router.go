// @APIVersion 1.0.0
// @Title bootgo Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact zhoubin296@gmail.com
// @TermsOfServiceUrl https://gitee.com/ipanocloud/bootgo
// @License MIT
// @LicenseUrl https://gitee.com/ipanocloud/bootgo/blob/master/LICENSE
package routers

import (
	"boot_beego/base"
	"boot_beego/controllers"
	"encoding/json"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/plugins/cors"
	"strings"

	"github.com/astaxie/beego"
)

func init() {
	// 最后一个参数必须设置为false 不然无法打印数据
	beego.InsertFilter("/*", beego.FinishRouter, FilterLog, false)

	//beego.InsertFilter("/*", beego.BeforeRouter, FilterAdminAuth, false)

	// 这段代码放在router.go文件的init()的开头
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		//AllowOrigins: 	  []string{"http://"+beego.AppConfig.String("front_end_domain")+":"+beego.AppConfig.String("front_end_port")},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))



	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/test",
			beego.NSInclude(
				&controllers.DemoController{},
			),
		),

	)
	beego.AddNamespace(ns)


	//beego.Router("/", &controllers.MainController{})
	beego.Router("/about", &controllers.AboutController{}, "get:Get")

	//example
	//beego.Router("api/index/index", &controllers.IndexController{}, "get:Index_Index")

}


// 添加日志拦截器
var FilterLog = func(ctx *context.Context) {
	url, _ := json.Marshal(ctx.Input.Data()["RouterPattern"])
	params, _ := json.Marshal(ctx.Request.Form)
	outputBytes, _ := json.Marshal(ctx.Input.Data()["json"])
	divider := " - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -"
	topDivider := "┌" + divider
	middleDivider := "├" + divider
	bottomDivider := "└" + divider
	outputStr := "\n" + topDivider + "\n│ 请求地址:" + string(url) + "\n" + middleDivider + "\n│ 请求参数:" + string(params) + "\n│ 返回数据:" + string(outputBytes) + "\n" + bottomDivider
	logs.Info(outputStr)
}


// 后台权限校验
var FilterAdminAuth = func(ctx *context.Context) {
	session := ctx.Input.Session(base.ADMIN_SESSION_KEY)
	if session == nil && !strings.Contains(ctx.Request.RequestURI, "login") {
		//fmt.Sprintf(url, rq.Encode())
		ctx.Redirect(302, "/home/login?r_url="+ctx.Request.RequestURI)
	}
	logs.Info("dddd")
}

