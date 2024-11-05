package gologger

import (
	"context"
	"daveslist/pkg/go-logger/core"
	"daveslist/pkg/go-logger/iface"
)

var log iface.Logger

func InitLogger(baseLog iface.BaseLog, ctxKey ...iface.ContextKey) {
	log = core.NewLogger(baseLog, ctxKey...)
}

func GetLogger() iface.Logger {
	return log
}

func AppendContextKey(key ...iface.ContextKey) {
	log.AppendContextKey(key...)
}

func Info(args ...interface{}) {
	log.Info(args...)
}

func Debug(args ...interface{}) {
	log.Debug(args...)
}

func Warn(args ...interface{}) {
	log.Warn(args...)
}

func Error(args ...interface{}) {
	log.Error(args...)
}

func Panic(args ...interface{}) {
	log.Panic(args...)
}

func Fatal(args ...interface{}) {
	log.Fatal(args...)
}

func InfoW(message string, kv ...interface{}) {
	log.InfoW(message, kv...)
}

func DebugW(message string, kv ...interface{}) {
	log.DebugW(message, kv...)
}

func WarnW(message string, kv ...interface{}) {
	log.WarnW(message, kv...)
}

func ErrorW(message string, kv ...interface{}) {
	log.ErrorW(message, kv...)
}

func PanicW(message string, kv ...interface{}) {
	log.PanicW(message, kv...)
}

func FatalW(message string, kv ...interface{}) {
	log.FatalW(message, kv...)
}

func InfoC(ctx context.Context, msg string, kv ...interface{}) {
	log.InfoC(ctx, msg, kv...)
}

func DebugC(ctx context.Context, msg string, kv ...interface{}) {
	log.DebugC(ctx, msg, kv...)
}

func WarnC(ctx context.Context, msg string, kv ...interface{}) {
	log.WarnC(ctx, msg, kv...)
}

func ErrorC(ctx context.Context, msg string, kv ...interface{}) {
	log.ErrorC(ctx, msg, kv...)
}

func PanicC(ctx context.Context, msg string, kv ...interface{}) {
	log.PanicC(ctx, msg, kv...)
}

func FatalC(ctx context.Context, msg string, kv ...interface{}) {
	log.FatalC(ctx, msg, kv...)
}
