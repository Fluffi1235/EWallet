package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var Logger *zap.SugaredLogger

func NewLogger() error {
	var err error
	cfgLogger := zap.NewProductionConfig()
	cfgLogger.DisableStacktrace = true
	cfgLogger.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)

	prodLogger, err := cfgLogger.Build()
	if err != nil {
		return err
	}

	Logger = prodLogger.Sugar()

	return nil
}
