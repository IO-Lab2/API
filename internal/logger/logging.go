package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var LoggerRaw *zap.Logger
var Logger *zap.SugaredLogger

func init() {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	LoggerRaw, _ = config.Build()

	Logger = LoggerRaw.Sugar() // Prettier logging
	Logger.Debug("Logger initialized")
}

func Sync() {
	if err := LoggerRaw.Sync(); err != nil {
		// handle the error, e.g., log it
		Logger.Error("Failed to sync logger", zap.Error(err))
	}
}
