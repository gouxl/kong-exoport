package models

import (
	"strconv"
)

//解析kong-log对象
type KongLogs struct {
	AuthenticatedEntity struct {
		ConsumerID string `json:"consumer_id"`
		ID         string `json:"id"`
	} `json:"authenticated_entity"`
	ClientIP string `json:"client_ip"`
	Consumer struct {
		CreatedAt int64  `json:"created_at"`
		ID        string `json:"id"`
		Username  string `json:"username"`
	} `json:"consumer"`
	Latencies struct {
		Kong    int64 `json:"kong"`
		Proxy   int64 `json:"proxy"`
		Request int64 `json:"request"`
	} `json:"latencies"`
	Request struct {
		Headers struct {
			Accept     string `json:"accept"`
			Host       string `json:"host"`
			User_agent string `json:"user-agent"`
		} `json:"headers"`
		Method      string   `json:"method"`
		Querystring struct{} `json:"querystring"`
		Size        string   `json:"size"`
		TLS         struct {
			Cipher                 string `json:"cipher"`
			ClientVerify           string `json:"client_verify"`
			SupportedClientCiphers string `json:"supported_client_ciphers"`
			Version                string `json:"version"`
		} `json:"tls"`
		URI string `json:"uri"`
		URL string `json:"url"`
	} `json:"request"`
	Response struct {
		Headers struct {
			Connection                       string `json:"Connection"`
			Content_Length                   string `json:"Content-Length"`
			Content_Type                     string `json:"Content-Type"`
			Access_control_allow_credentials string `json:"access-control-allow-credentials"`
			Access_control_allow_origin      string `json:"access-control-allow-origin"`
			Server                           string `json:"server"`
			Via                              string `json:"via"`
		} `json:"headers"`
		Size   string `json:"size"`
		Status int64  `json:"status"`
	} `json:"response"`
	Route struct {
		CreatedAt     int64       `json:"created_at"`
		Hosts         interface{} `json:"hosts"`
		ID            string      `json:"id"`
		Methods       interface{} `json:"methods"`
		Paths         []string    `json:"paths"`
		PreserveHost  bool        `json:"preserve_host"`
		Protocols     []string    `json:"protocols"`
		RegexPriority int64       `json:"regex_priority"`
		Service       struct {
			ID string `json:"id"`
		} `json:"service"`
		StripPath bool  `json:"strip_path"`
		UpdatedAt int64 `json:"updated_at"`
	} `json:"route"`
	Service struct {
		ConnectTimeout int64  `json:"connect_timeout"`
		CreatedAt      int64  `json:"created_at"`
		Host           string `json:"host"`
		ID             string `json:"id"`
		Name           string `json:"name"`
		Path           string `json:"path"`
		Port           int64  `json:"port"`
		Protocol       string `json:"protocol"`
		ReadTimeout    int64  `json:"read_timeout"`
		Retries        int64  `json:"retries"`
		UpdatedAt      int64  `json:"updated_at"`
		WriteTimeout   int64  `json:"write_timeout"`
	} `json:"service"`
	StartedAt int64 `json:"started_at"`
	Tries     []struct {
		Code  int64  `json:"code"`
		IP    string `json:"ip"`
		Port  int64  `json:"port"`
		State string `json:"state"`
	} `json:"tries"`
	UpstreamURI string `json:"upstream_uri"`
	Workspaces  []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"workspaces"`
}

type ConsumerReqStatus struct {
	Service  string
	Route    string
	Code     string
	Consumer string
	Url      string
	ClientIP string
	StartedAt string
}


func AssignStatus(l KongLogs) (c ConsumerReqStatus) {
	c.Service = l.Service.Name
	if len(l.Service.Name) == 0 {
		c.Service = l.Service.ID
	}
	c.Route = l.Route.ID
	//c.Url = l.Request.URL
	c.Code = strconv.FormatInt(l.Response.Status, 10)
	c.Consumer = l.Consumer.Username
	if len(l.Consumer.Username) == 0 {
		c.Consumer = l.Consumer.ID
	}
	c.ClientIP = l.ClientIP
	c.StartedAt = strconv.FormatInt(l.StartedAt, 10)
	return
}

