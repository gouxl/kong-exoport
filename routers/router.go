// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"kong-exporter/controllers"
	"path"

	"github.com/astaxie/beego"
)

func init() {

	// Beego注解路由代码生成规则和程序运行路径相关，需要改写一下避免产生不一致的文件名
	AppPath := path.Base(beego.AppPath)
	if beego.BConfig.RunMode == "dev" && AppPath == "_build" {
		beego.AppPath = path.Join(path.Dir(beego.AppPath), "kong-exporter")
	}
	logs.Info("AppPath:", AppPath, "beego.AppPath:", beego.AppPath)

	//swagger
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	//cors 解决跨域问题
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))

	ns := beego.NewNamespace("/kong",
		beego.NSNamespace("/",
			beego.NSInclude(
				&controllers.LogController{},
			),
		),
	)
	beego.AddNamespace(ns)
	beego.Handler("/metrics", promhttp.Handler())
}
