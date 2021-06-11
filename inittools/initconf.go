package inittools

import (
	"flag"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func InitLoadConfig() {
	configpath := flag.String("config", "./conf/app.conf", "config:default is app.conf")
	flag.Parse()
	logs.Info("config path:", *configpath)
	beego.LoadAppConfig("ini", *configpath)
}
