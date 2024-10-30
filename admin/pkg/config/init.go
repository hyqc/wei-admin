package config

import (
	"hope_city/lib/log"
	"hope_city/lib/mongo"
	"hope_city/lib/mysql"
	"hope_city/lib/oauth"
	"hope_city/lib/redis"
	"hope_city/lib/tencent_im"
	"hope_city/lib/trace"
	"hope_city/lib/zinx_defined/utils"
)

type ConfInit struct {
	Nacos      *NaCos                    `json:"nacos"`
	Server     *Server                   `json:"server"`
	HttpServer *Server                   `json:"http_server"`
	TcpServer  *utils.Config             `json:"tcp_server"`
	Broker     *Broker                   `json:"broker"`
	Jwt        *Jwt                      `json:"jwt"`
	Log        *log.Conf                 `json:"log"`
	Trace      *trace.Config             `json:"trace"`
	Sls        *log.SlsConf              `json:"sls"`
	Registry   *Registry                 `json:"registry"`
	Redis      map[string]*redis.Config  `json:"redis"`
	Mysqldb    map[string]*mysql.Config  `json:"mysqldb"`
	Mongo      map[string]*mongo.Config  `json:"mongo"`
	Sms        *Sms                      `json:"sms"`
	TapSDK     *oauth.TapSDKConfig       `json:"tap_sdk"`
	Airmart    *oauth.AirMartOAuthConfig `json:"airmart"`
	IM         *tencent_im.IM            `json:"im"`
}

func (c *ConfInit) Handle() error {
	return nil
}

func (c *ConfInit) Original() interface{} {
	return c
}

type AppSecret struct {
	Url    string
	Key    string
	Secret string
}
