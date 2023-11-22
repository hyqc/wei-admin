package config

import (
	"fmt"
	"github.com/go-micro/plugins/v4/config/encoder/yaml"
	"go-micro.dev/v4/config"
	"go-micro.dev/v4/config/reader"
	"go-micro.dev/v4/config/reader/json"
	"go-micro.dev/v4/config/source/file"
	"go-micro.dev/v4/logger"
)

type IConfig interface {
	Init() error
	Original() interface{}
	Handle() error
}

type Config struct {
	Server Server `yaml:"server"`
}

type Server struct {
	Name string `yaml:"name"`
	Port string `yaml:"port"`
}

var (
	AppVersion = "1.0.0"
	AppConfig  = NewConfig()
)

const (
	configFileFullPath = "./config.yaml"
)

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Original() interface{} {
	return c
}

func (c *Config) Handle() error {
	return nil
}

// Init 配置初始化
func (c *Config) Init() error {
	enc := yaml.NewEncoder()

	conf, err := config.NewConfig(
		config.WithReader(json.NewReader(reader.WithEncoder(enc))),
	)
	if err != nil {
		return err
	}

	fp := AppFlags.ConfigDir.Value.(string)

	err = conf.Load(file.NewSource(
		file.WithPath(fp),
	))
	if err != nil {
		return err
	}
	if err := conf.Scan(c); err != nil {
		return err
	}
	watch, err := conf.Watch(fp)
	if err != nil {
		return err
	}
	defer func(watch config.Watcher) {
		err := watch.Stop()
		if err != nil {
			logger.Errorf("stop config watch failed, error:%v", err)
		}
	}(watch)

	go func() {
		for {
			fmt.Println("--------------------")
			next, err := watch.Next()
			fmt.Println("====================")
			if err != nil {
				logger.Errorf("watch config file:%s failed, error: %v", fp, err)
				continue
			}
			if err := reload(next, *AppConfig); err != nil {
				logger.Errorf("watch config file:%s scan failed, error: %v", fp, err)
			}
		}
	}()

	return nil
}

func reload(read reader.Value, conf Config) error {
	logger.Infof("reload before scan config.yaml: ", conf)
	data := conf.Original()
	if err := read.Scan(data); err != nil {
		return err
	}
	logger.Infof("reload after config.yaml: ", data)
	return conf.Handle()
}
