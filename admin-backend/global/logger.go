package global

import (
	"admin/pkg/logger"
	"gopkg.in/natefinch/lumberjack.v2"
)

func initLogger() error {
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
	Log = lg
	LogSugar = lg.Sugar()
	return nil
}

func CloseLogger() {
	if Log != nil {
		err := Log.Sync()
		if err != nil {
			return
		}
	}
	if LogSugar != nil {
		err := LogSugar.Sync()
		if err != nil {
			return
		}
	}
}
