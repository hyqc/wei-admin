package config

import "time"

// NacosConfig Nacos配置
type NacosConfig struct {
	Hosts       []string `json:"hosts"`
	Namespace   string   `json:"namespace"`
	User        string   `json:"user"`
	Password    string   `json:"password"` //密码
	DataId      string   `json:"dataId"`
	Group       string   `json:"group"`        //分组
	Cluster     string   `json:"cluster"`      //tag 集群名
	LocalEnable bool     `json:"local_enable"` //是否启用本地文件配置
}

func (n *NacosConfig) Original() interface{} {
	return n
}

func (n *NacosConfig) Handle() error {
	return nil
}

// Server 服务配置
type Server struct {
	Name    string `json:"name"`    //服务名
	Version string `json:"version"` //服务版本
	Port    int    `json:"port"`    //端口号
	Id      string `json:"id"`      //服务唯一ID,md5(addr)
	Debug   bool   `yaml:"debug"`   //是否开启Debug
	Pprof   bool   `yaml:"pprof"`   //启用pprof分析
}

// Logger 日志配置
type Logger struct {
	Filename   string `yaml:"filename"`    //文件名
	MaxSize    int    `yaml:"max_size"`    //文件最大xM单位M
	MaxBackups int    `yaml:"max_backups"` //最多保存备份日志数量
	MaxAge     int    `yaml:"max_age"`     //保留日志多少天
	Level      string `yaml:"level"`       //日志等级
	Compress   bool   `yaml:"compress"`    //是否压缩日志
	Json       bool   `yaml:"json"`        //是否json日志
}

// Broker 消息中间件配置
type Broker struct {
	Address []string `json:"address"`
}

type Jwt struct {
	Private    string                         `json:"private"`
	Public     string                         `json:"public"`
	Expire     time.Duration                  `json:"expire"`  //秒
	Ignores    []IgnoreUrlRule                `json:"ignores"` //不严重jwt token的路由地址
	IgnoresMap map[string]map[string]struct{} //ignore转map
}

type IgnoreUrlRule struct {
	Method string   `json:"method"`
	Paths  []string `json:"paths"`
}
