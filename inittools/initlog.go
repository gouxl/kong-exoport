package inittools

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func InitLog() {
	logFile := beego.AppConfig.String("logFile")
	if logFile != "" {
		err := logs.SetLogger(logs.AdapterFile, fmt.Sprintf(`{"filename":"%s","maxlines":10000,"daily":true,"maxdays":10,"color":true}`, beego.AppConfig.String("logfile")))
		if err != nil {
			logs.Error("log file err:", err)
			return
		}
	} else {
		logs.SetLogger(logs.AdapterConsole, "")
	}
	logLevel, err := beego.AppConfig.Int("loglevel")
	if err != nil {
		logs.Error("log file err:", err)
		return
	}
	logs.SetLevel(logLevel)
	logs.SetLogFuncCall(true)
}
