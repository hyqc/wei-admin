package global

import (
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Server   Server   `yaml:"server"`
	Database Database `yaml:"database"`
	Logger   Logger   `yaml:"logger"`
	DBClient *DBClient
}

var (
	AppConfig        = &Config{}
	AppLogger        *zap.Logger
	AppLoggerSugared *zap.SugaredLogger
)

func ParseConfig(name string) error {
	body, err := os.ReadFile(name)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(body, AppConfig); err != nil {
		return err
	}
	if AppConfig.Server.JWT.UsefulLife == 0 {
		AppConfig.Server.JWT.UsefulLife = 3600 * 24
	}
	return nil
}
