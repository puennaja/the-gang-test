package iface

import "context"

// go:generate mockgen -source=logger.go -destination=mocks/logger.go -package=mocks
type Logger interface {
	AppendContextKey(key ...ContextKey)
	Info(args ...interface{})
	Debug(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Panic(args ...interface{})
	Fatal(args ...interface{})
	InfoW(message string, kv ...interface{})
	DebugW(message string, kv ...interface{})
	WarnW(message string, kv ...interface{})
	ErrorW(message string, kv ...interface{})
	PanicW(message string, kv ...interface{})
	FatalW(message string, kv ...interface{})
	InfoC(ctx context.Context, msg string, kv ...interface{})
	DebugC(ctx context.Context, msg string, kv ...interface{})
	WarnC(ctx context.Context, msg string, kv ...interface{})
	ErrorC(ctx context.Context, msg string, kv ...interface{})
	PanicC(ctx context.Context, msg string, kv ...interface{})
	FatalC(ctx context.Context, msg string, kv ...interface{})
}
