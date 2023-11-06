package config

import (
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
	"gorm.io/gorm"
	"os"
)

type Config struct {
	Server   Server   `yaml:"server"`
	Database Database `yaml:"database"`
	Logger   Logger   `yaml:"logger"`
}

type DB struct {
	Demo *gorm.DB
}

var (
	AppConfig = &Config{}
	AppLogger *zap.Logger
	AppDB     = &DB{}
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
