package global

import (
	"admin/pkg/logger"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger struct {
	Filename   string `yaml:"filename"`
	MaxSize    int    `yaml:"max_size"`
	MaxBackups int    `yaml:"max_backups"`
	MaxAge     int    `yaml:"max_age"`
	Compress   bool   `yaml:"compress"`
	Level      string `yaml:"level"`
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
		Level: logger.LevelMap[cf.Level],
		Json:  cf.Json,
	})
	if err != nil {
		return err
	}
	AppLogger = lg
	AppLoggerSugared = lg.Sugar()
	return nil
}

func LoggerClose() {
	if AppLogger != nil {
		err := AppLogger.Sync()
		if err != nil {
			return
		}
	}
	if AppLoggerSugared != nil {
		err := AppLoggerSugared.Sync()
		if err != nil {
			return
		}
	}
}
