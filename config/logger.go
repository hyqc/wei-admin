package config

import (
	"ginweb/pkg/logger"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger struct {
	Filename   string `yaml:"filename"`
	MaxSize    int    `yaml:"max_size"`
	MaxBackups int    `yaml:"max_backups"`
	MaxAge     int    `yaml:"max_age"`
	Compress   bool   `yaml:"compress"`
	Level      int    `yaml:"level"`
	Json       bool   `yaml:"json"`
}

func InitLogger() error {
	cf := AppConfig.Logger
	lg, err := logger.NewLogger(&logger.Logger{
		Writer: &lumberjack.Logger{
			Filename:   cf.Filename,
			MaxSize:    cf.MaxSize,
			MaxBackups: cf.MaxBackups,
			MaxAge:     cf.MaxAge,
			Compress:   cf.Compress,
		},
		Level: zapcore.Level(cf.Level),
		Json:  cf.Json,
	})
	if err != nil {
		return err
	}
	AppLogger = lg
	return nil
}

func LoggerClose() {
	if AppLogger != nil {
		err := AppLogger.Sync()
		if err != nil {
			return
		}
	}
}
