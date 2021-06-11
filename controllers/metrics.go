package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/prometheus/client_golang/prometheus"
	"kong-exporter/models"
	"net/http"
	"strconv"
)

var (
	DEFAULT_BUCKETS = []float64{1, 50, 300, 500, 1000,
		2000, 5000, 10000, 30000, 60000}
	consumerRequest = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "kong_http_con_status",
			Help: "HELP kong_http_status HTTP status codes per service/route in Kong",
		}, []string{"service", "route", "status", "consumer", "clientIp", "recTime"})

	kongConBandwidth = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "kong_con_bandwidth",
			Help: "kong con bandwidth",
		}, []string{"service", "consumer", "route", "status"})
)

func init() {
	prometheus.MustRegister(consumerRequest)
	prometheus.MustRegister(kongConBandwidth)
}

// Operations about object
type LogController struct {
	beego.Controller
}

// @Title Create
// @Description create object
// @Param	body		body 	models.Object	true		"The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [post]
func (l *LogController) Post() {
	var kongLog models.KongLogs
	err := json.Unmarshal(l.Ctx.Input.RequestBody, &kongLog)
	if err != nil {
		logs.Error("")
		l.ServeJSON()
	}
	//logs.Info(kongLog)
	logs.Info(kongLog.Request)
	logs.Info(kongLog.ClientIP)
	//1.kong_http_status
	assigment := models.AssignStatus(kongLog)
	consumerRequest.With(
		prometheus.Labels{
			"service": assigment.Service,
			"route": assigment.Route,
			"status": assigment.Code,
			"consumer": assigment.Consumer,
			"clientIp": assigment.ClientIP,
			"recTime": assigment.StartedAt}).Inc()
	//2.kong_con_bandwidth
	resSize, err := strconv.Atoi(kongLog.Response.Size)
	if err != nil {
		logs.Error(" ")
		l.ServeJSON()
	}
	reqSize, err := strconv.Atoi(kongLog.Request.Size)
	if err != nil {
		logs.Error("")
		l.ServeJSON()
	}
	totalSize := resSize + reqSize
	kongConBandwidth.With(
		prometheus.Labels{"service": assigment.Service,
			"route": assigment.Route,
			"consumer": assigment.Consumer,
			"status": assigment.Code}).Add(float64(totalSize))
	l.Ctx.Output.SetStatus(http.StatusOK)
	l.Data["json"] = "success"
	l.ServeJSON()
}
