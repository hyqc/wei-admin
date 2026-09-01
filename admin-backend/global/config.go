package global

import (
	"admin/pkg/captcha"
	"admin/pkg/config"
	"admin/pkg/utils"
	"admin/pkg/utils/jwt"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go-micro.dev/v5/logger"
	"go.uber.org/zap"
)

type Config struct {
	// ConfigSource 配置来源：file（本地配置文件，默认）| nacos（Nacos 配置中心）
	ConfigSource string         `json:"configSource"`
	Server       config.Server  `json:"server"`
	Logger       config.Logger  `json:"logger"`
	JWT          config.Jwt     `json:"JWT"`
	Redis        config.Redis   `json:"redis"`
	Store        config.Store   `json:"store"`
	Captcha      config.Captcha `json:"captcha"`
}

func (c *Config) Handle() error {
	return nil
}

func (c *Config) Original() any {
	return c
}

// GetConfigSource 实现 config.IConfigSource，供配置加载决定是否需要从 Nacos 拉取
func (c *Config) GetConfigSource() string {
	return c.ConfigSource
}

type initConfigFunc func() error

var (
	AppConfig  = &Config{}
	Log        *zap.Logger
	LogSugar   *zap.SugaredLogger
	AppDB      *DBClient
	AppAuth    *jwt.Auth
	AppRedis   *redis.Client
	AppCaptcha *captcha.Manager
)

var (
	//注册初始化方法
	initConfigCall = []initConfigFunc{
		initConfig,
		initLogger,
		initDatabase,
		initJwt,
		initServer,
		initRedis,
		initCaptcha,
	}
)

func init() {
	initValidator()
}

func initConfig() error {
	return config.Init(AppConfig)
}

func initServer() error {
	ip, err := utils.GetOutBoundIP()
	if err != nil {
		logger.Errorf("initServer utils.GetOutBoundIP error: %v", err)
		return err
	}
	AppConfig.Server.Id = utils.Md5(fmt.Sprintf("%s=%s:%d", AppConfig.Server.Name, ip, AppConfig.Server.Port))
	return nil
}
