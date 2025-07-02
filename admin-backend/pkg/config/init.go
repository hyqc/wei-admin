package config

import (
	"admin/pkg/config/store/mysql"
	"admin/pkg/config/store/redis"
)

type Config struct {
	Nacos      *NacosConfig             `json:"nacos"`
	Server     *Server                  `json:"server"`
	HttpServer *Server                  `json:"http_server"`
	Broker     *Broker                  `json:"broker"`
	Jwt        *Jwt                     `json:"jwt"`
	Redis      map[string]*redis.Config `json:"redis"`
	Mysqldb    map[string]*mysql.Config `json:"mysqldb"`
}

func (c *Config) Handle() error {
	return nil
}

func (c *Config) Original() interface{} {
	return c
}
