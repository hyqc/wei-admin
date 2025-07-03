package global

import (
	"admin/pkg/config"
	"admin/pkg/utils"
	"admin/pkg/utils/jwt"
	"fmt"
	"go-micro.dev/v5/logger"
	"go.uber.org/zap"
)

type Config struct {
	Server config.Server `json:"server"`
	Logger config.Logger `json:"logger"`
	JWT    config.Jwt    `json:"JWT"`
	Store  Store         `json:"store"`
}

func (c *Config) Handle() error {
	return nil
}

func (c *Config) Original() any {
	return c
}

type initConfigFunc func() error

var (
	AppConfig = &Config{}
	Log       *zap.Logger
	LogSugar  *zap.SugaredLogger
	AppDB     *DBClient
	AppAuth   *jwt.Auth
)

var (
	//注册初始化方法
	initConfigCall = []initConfigFunc{
		initConfig,
		initLogger,
		initMySQLDB,
		initJwt,
		initServer,
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
