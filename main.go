package main

import (
	"kong-exporter/inittools"
	_ "kong-exporter/routers"

	"github.com/astaxie/beego"
)

func main() {
	//1.指定配置文件路径
	inittools.InitLoadConfig()
	//2.初始化日志配置
	inittools.InitLog()
	beego.Run()
}
