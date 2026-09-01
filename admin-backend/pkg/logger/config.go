package logger

import (
	"errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
)

type Logger struct {
	Writer io.Writer
	Level  zapcore.Level
	Json   bool
	// Stdout 是否同时输出到标准输出（容器/K8s 部署建议开启）
	Stdout bool
}

const (
	LoggerLevelDebug = "debug"
	LoggerLevelInfo  = "info"
	LoggerLevelWarn  = "warn"
	LoggerLevelError = "error"
	LoggerLevelFatal = "fatal"
)

var config = zap.NewProductionEncoderConfig()

var (
	LevelMap = map[string]zapcore.Level{
		LoggerLevelDebug: zapcore.DebugLevel,
		LoggerLevelInfo:  zapcore.InfoLevel,
		LoggerLevelWarn:  zapcore.WarnLevel,
		LoggerLevelError: zapcore.ErrorLevel,
		LoggerLevelFatal: zapcore.FatalLevel,
	}
)

func NewDefaultLogger(filename string) *Logger {
	return &Logger{
		Writer: &lumberjack.Logger{
			Filename:   filename,
			MaxSize:    5,
			MaxBackups: 7,
			MaxAge:     14,
			Compress:   true,
		},
		Level: zapcore.DebugLevel,
	}
}

func NewLogger(cf *Logger) (*zap.Logger, error) {
	if cf == nil || cf.Writer == nil {
		return nil, errors.New("cf or cf.Writer is nil")
	}
	config.EncodeTime = zapcore.RFC3339TimeEncoder
	config.EncodeLevel = zapcore.CapitalLevelEncoder
	var encoder zapcore.Encoder
	if cf.Json {
		encoder = zapcore.NewJSONEncoder(config)
	} else {
		encoder = zapcore.NewConsoleEncoder(config)
	}
	writers := []zapcore.WriteSyncer{zapcore.AddSync(cf.Writer)}
	if cf.Stdout {
		writers = append(writers, zapcore.AddSync(os.Stdout))
	}
	var sy zapcore.WriteSyncer = writers[0]
	if len(writers) > 1 {
		sy = zapcore.NewMultiWriteSyncer(writers...)
	}
	core := zapcore.NewCore(encoder, sy, cf.Level)
	return zap.New(core, zap.AddCaller()), nil
}
