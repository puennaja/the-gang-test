package core

import (
	"daveslist/pkg/go-logger/iface"
	"log"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func DefaultBaseLogLevel(level string) int {
	switch strings.ToLower(level) {
	case "debug":
		return -1
	case "info":
		return 0
	case "warn":
		return 1
	case "error":
		return 2
	case "panic":
		return 4
	case "fatal":
		return 5
	default:
		return 0
	}
}

func BuildDefaultBaseLog(cfg iface.Config) iface.BaseLog {
	// set zap config
	zapConfig := zap.NewProductionConfig()
	zapConfig.Level = zap.NewAtomicLevelAt(zapcore.Level(
		DefaultBaseLogLevel(cfg.LogLevel()),
	))
	zapConfig.EncoderConfig.TimeKey = cfg.TimeKey()
	zapConfig.EncoderConfig.StacktraceKey = cfg.StacktraceKey()
	zapConfig.EncoderConfig.CallerKey = cfg.CallerKey()
	zapConfig.EncoderConfig.MessageKey = cfg.MessageKey()
	zapConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(cfg.TimeFormat())
	zapConfig.DisableStacktrace = cfg.DisableStacktrace()
	zapConfig.DisableCaller = cfg.DisableCaller()
	zapConfig.InitialFields = cfg.InitialFields()
	zapLog, err := zapConfig.Build()
	if err != nil {
		log.Fatal(err)
	}
	return zapLog.Sugar()
}
