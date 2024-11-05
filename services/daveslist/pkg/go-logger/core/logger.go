package core

import (
	"context"
	"daveslist/pkg/go-logger/iface"
)

type CoreLogger struct {
	baseLog iface.BaseLog
	ctxKey  []iface.ContextKey
}

func NewLogger(baseLog iface.BaseLog, ctxKey ...iface.ContextKey) iface.Logger {
	return &CoreLogger{
		baseLog: baseLog,
		ctxKey:  ctxKey,
	}
}
func (l *CoreLogger) AppendContextKey(key ...iface.ContextKey) {
	l.ctxKey = append(l.ctxKey, key...)
}

func (l *CoreLogger) Info(args ...interface{}) {
	l.baseLog.Info(args...)
}

func (l *CoreLogger) Debug(args ...interface{}) {
	l.baseLog.Debug(args...)
}

func (l *CoreLogger) Warn(args ...interface{}) {
	l.baseLog.Warn(args...)
}

func (l *CoreLogger) Error(args ...interface{}) {
	l.baseLog.Error(args...)
}

func (l *CoreLogger) Panic(args ...interface{}) {
	l.baseLog.Panic(args...)
}

func (l *CoreLogger) Fatal(args ...interface{}) {
	l.baseLog.Fatal(args...)
}

func (l *CoreLogger) InfoW(message string, kv ...interface{}) {
	l.baseLog.Infow(message, kv...)
}

func (l *CoreLogger) DebugW(message string, kv ...interface{}) {
	l.baseLog.Debugw(message, kv...)
}

func (l *CoreLogger) WarnW(message string, kv ...interface{}) {
	l.baseLog.Warnw(message, kv...)
}

func (l *CoreLogger) ErrorW(message string, kv ...interface{}) {
	l.baseLog.Errorw(message, kv...)
}

func (l *CoreLogger) PanicW(message string, kv ...interface{}) {
	l.baseLog.Panicw(message, kv...)
}

func (l *CoreLogger) FatalW(message string, kv ...interface{}) {
	l.baseLog.Fatalw(message, kv...)
}

func (l *CoreLogger) buildContextLog(ctx context.Context) []interface{} {
	ctxLog := make([]interface{}, 0)
	for _, key := range l.ctxKey {
		ctxLog = append(ctxLog, key.String())
		ctxLog = append(ctxLog, ctx.Value(key))
	}
	return ctxLog
}

func (l *CoreLogger) InfoC(ctx context.Context, msg string, kv ...interface{}) {
	ctxLog := l.buildContextLog(ctx)
	log := append(ctxLog, kv...)
	l.baseLog.Infow(msg, log...)
}

func (l *CoreLogger) DebugC(ctx context.Context, msg string, kv ...interface{}) {
	ctxLog := l.buildContextLog(ctx)
	log := append(ctxLog, kv...)
	l.baseLog.Debugw(msg, log...)
}

func (l *CoreLogger) WarnC(ctx context.Context, msg string, kv ...interface{}) {
	ctxLog := l.buildContextLog(ctx)
	log := append(ctxLog, kv...)
	l.baseLog.Warnw(msg, log...)
}

func (l *CoreLogger) ErrorC(ctx context.Context, msg string, kv ...interface{}) {
	ctxLog := l.buildContextLog(ctx)
	log := append(ctxLog, kv...)
	l.baseLog.Errorw(msg, log...)
}

func (l *CoreLogger) PanicC(ctx context.Context, msg string, kv ...interface{}) {
	ctxLog := l.buildContextLog(ctx)
	log := append(ctxLog, kv...)
	l.baseLog.Panicw(msg, log...)
}

func (l *CoreLogger) FatalC(ctx context.Context, msg string, kv ...interface{}) {
	ctxLog := l.buildContextLog(ctx)
	log := append(ctxLog, kv...)
	l.baseLog.Fatalw(msg, log...)
}
