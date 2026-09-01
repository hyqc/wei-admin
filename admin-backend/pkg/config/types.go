package config

import (
	"admin/pkg/config/store/mongodb"
	"admin/pkg/config/store/mysql"
	"time"
)

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
	Stdout     bool   `yaml:"stdout"`      //是否同时输出到标准输出（容器部署建议开启）
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

// Captcha 图片验证码配置
type Captcha struct {
	// Store 存储方式：memory（单机部署，默认）| redis（集群部署）
	Store string `json:"store"`
	// Length/Expire/Width/Height 为全局默认值，未单独配置的场景使用默认值
	Length int   `json:"length"` // 验证码位数
	Expire int64 `json:"expire"` // 有效期（秒）
	Width  int   `json:"width"`  // 图片宽度
	Height int   `json:"height"` // 图片高度
	// Scenes 各场景自定义配置（key 为场景名，如 login），未配置的项继承全局默认值
	Scenes map[string]CaptchaScene `json:"scenes"`
}

// CaptchaScene 单个验证码场景配置
type CaptchaScene struct {
	Length int   `json:"length"` // 验证码位数
	Expire int64 `json:"expire"` // 有效期（秒）
	Width  int   `json:"width"`  // 图片宽度
	Height int   `json:"height"` // 图片高度
}

// Redis 配置（集群部署时用于验证码等共享存储）
type Redis struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

// Store 数据源配置：按数据库类型分组，未配置的类型不会初始化
type Store struct {
	Mysql   MysqlStore `json:"mysql"`
	Mongodb MongoStore `json:"mongodb"`
}

// MysqlStore MySQL 数据源集合
type MysqlStore struct {
	// Default 默认数据源名（对应 Sources 的 key），为空时取唯一的数据源
	Default string `json:"default"`
	// Sources 数据源配置，key 为数据源名
	Sources map[string]mysql.Config `json:"sources"`
}

// MongoStore MongoDB 数据源集合
type MongoStore struct {
	// Default 默认数据源名（对应 Sources 的 key），为空时取唯一的数据源
	Default string `json:"default"`
	// Sources 数据源配置，key 为数据源名；为空表示不启用 MongoDB
	Sources map[string]mongodb.Config `json:"sources"`
}
