package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func InitLogger(env string) {
	var err error
	var config zap.Config
	var encodeConfig zapcore.EncoderConfig
	if env == "prd" {
		config = zap.NewProductionConfig()
		encodeConfig = zap.NewProductionEncoderConfig()
	} else {
		config = zap.NewDevelopmentConfig()
		encodeConfig = zap.NewDevelopmentEncoderConfig()
	}
	encodeConfig.TimeKey = "timestamp"
	encodeConfig.CallerKey = "source"
	encodeConfig.StacktraceKey = ""
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig = encodeConfig
	log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

func Info(msg string, field ...zap.Field) {
	log.Info(msg, field...)
}

func Warn(msg string, field ...zap.Field) {
	log.Warn(msg, field...)
}
func Debug(msg string, field ...zap.Field) {
	log.Debug(msg, field...)
}

func Error(msg string, field ...zap.Field) {
	log.Error(msg, field...)
}
