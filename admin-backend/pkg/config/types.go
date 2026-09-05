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

// Upload 文件上传配置
type Upload struct {
	// Driver 存储驱动：local（默认，单机部署）| aliyun（阿里云 OSS）| qcloud（腾讯云 COS）| s3（亚马逊 S3 及兼容实现）
	Driver string `json:"driver" yaml:"driver"`
	// Domain 访问域名前缀（结尾不带斜杠），用于拼接文件访问链接
	// 留空时：本地存储返回站点相对路径，云存储返回厂商默认公网地址
	Domain string `json:"domain" yaml:"domain"`
	// MaxSize 单文件大小上限（MB），0 表示不限制
	MaxSize int64 `json:"max_size" yaml:"max_size"`
	// AllowedExts 允许上传的扩展名（小写、不含点），为空表示不限制
	AllowedExts []string `json:"allowed_exts" yaml:"allowed_exts"`
	// Local 本地存储配置
	Local UploadLocal `json:"local" yaml:"local"`
	// Aliyun 阿里云 OSS 配置
	Aliyun UploadAliyun `json:"aliyun" yaml:"aliyun"`
	// Qcloud 腾讯云 COS 配置
	Qcloud UploadQcloud `json:"qcloud" yaml:"qcloud"`
	// S3 亚马逊 S3（及 MinIO 等 S3 兼容实现）配置
	S3 UploadS3 `json:"s3" yaml:"s3"`
}

// UploadLocal 本地存储配置（单机部署默认方案）
type UploadLocal struct {
	// Root 存储根目录，相对后端工作目录，如 ./upload
	Root string `json:"root" yaml:"root"`
	// Prefix 静态访问前缀，需与注册到路由的静态目录一致，如 /upload
	Prefix string `json:"prefix" yaml:"prefix"`
}

// UploadAliyun 阿里云 OSS 配置
type UploadAliyun struct {
	Endpoint        string `json:"endpoint" yaml:"endpoint"`
	AccessKeyID     string `json:"access_key_id" yaml:"access_key_id"`
	AccessKeySecret string `json:"access_key_secret" yaml:"access_key_secret"`
	Bucket          string `json:"bucket" yaml:"bucket"`
	// BasePath 对象键统一前缀（可留空），如 admin
	BasePath string `json:"base_path" yaml:"base_path"`
}

// UploadQcloud 腾讯云 COS 配置
type UploadQcloud struct {
	Region    string `json:"region" yaml:"region"`
	SecretID  string `json:"secret_id" yaml:"secret_id"`
	SecretKey string `json:"secret_key" yaml:"secret_key"`
	Bucket    string `json:"bucket" yaml:"bucket"`
	// BasePath 对象键统一前缀（可留空）
	BasePath string `json:"base_path" yaml:"base_path"`
}

// UploadS3 亚马逊 S3 / S3 兼容存储配置
type UploadS3 struct {
	Region          string `json:"region" yaml:"region"`
	AccessKeyID     string `json:"access_key_id" yaml:"access_key_id"`
	SecretAccessKey string `json:"secret_access_key" yaml:"secret_access_key"`
	Bucket          string `json:"bucket" yaml:"bucket"`
	// Endpoint 自定义端点（MinIO 等 S3 兼容实现填写），留空使用厂商默认端点
	Endpoint string `json:"endpoint" yaml:"endpoint"`
	// BasePath 对象键统一前缀（可留空）
	BasePath string `json:"base_path" yaml:"base_path"`
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
