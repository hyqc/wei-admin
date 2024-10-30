package config

import (
	"hope_city/lib/sms"
)

type Jaeger struct {
	Addr string `json:"addr"`
}

type Sms struct {
	//当前使用短信的类型
	Type string `json:"type"`
	//短信宝
	Smsbao *sms.SmsBaoConfig `json:"smsbao"`
}

// Registry 服务发现配置
type Registry struct {
	Endpoints []string `json:"endpoints"` //地址
}

type Server struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Addr    string `json:"addr"`
	Id      string `json:"id"`
}

type HttpServer struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Addr    string `json:"addr"`
}

type GrpcServer struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Addr    string `json:"addr"`
}

type Broker struct {
	Address []string `json:"address"`
}

type Jwt struct {
	Private string `json:"private"`
	Public  string `json:"public"`
	Expire  string `json:"expire"`
}
