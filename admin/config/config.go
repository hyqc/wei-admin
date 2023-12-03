package config

import (
	"github.com/sony/sonyflake"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Server   Server   `yaml:"server"`
	Database Database `yaml:"database"`
	Logger   Logger   `yaml:"logger"`
}

var (
	AppConfig        = &Config{}
	AppLogger        *zap.Logger
	AppLoggerSugared *zap.SugaredLogger
	AppSnoyflake     *sonyflake.Sonyflake
)

func ParseConfig(name string) error {
	body, err := os.ReadFile(name)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(body, AppConfig); err != nil {
		return err
	}
	return nil
}
