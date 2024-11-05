package iface

// go:generate mockgen -source=baselog.go -destination=mocks/baselog.go -package=mocks
type BaseLog interface {
	Info(args ...interface{})
	Debug(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Panic(args ...interface{})
	Fatal(args ...interface{})
	Infow(message string, kv ...interface{})
	Debugw(message string, kv ...interface{})
	Warnw(message string, kv ...interface{})
	Errorw(message string, kv ...interface{})
	Panicw(message string, kv ...interface{})
	Fatalw(message string, kv ...interface{})
}
