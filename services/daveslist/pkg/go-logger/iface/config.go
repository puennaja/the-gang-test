package iface

type Config interface {
	SetLogLevel(set string)
	SetTimeKey(set string)
	SetStacktraceKey(set string)
	SetCallerKey(set string)
	SetMessageKey(set string)
	SetTimeFormat(set string)
	SetDisableStacktrace(set bool)
	SetDisableCaller(set bool)
	SetInitialFields(set map[string]interface{})
	LogLevel() string
	TimeKey() string
	StacktraceKey() string
	CallerKey() string
	MessageKey() string
	TimeFormat() string
	DisableStacktrace() bool
	DisableCaller() bool
	InitialFields() map[string]interface{}
}
