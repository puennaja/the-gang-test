package config

import (
	"daveslist/pkg/go-logger/iface"
	"time"
)

type Config struct {
	logLevel          string `default:"info"`
	timeKey           string `default:"timestamp"`
	stacktraceKey     string `default:"stacktrace"`
	callerKey         string `default:"caller"`
	messageKey        string `default:"message"`
	timeFormat        string `default:"2006-01-02T15:04:05.000Z07:00"`
	disableStacktrace bool   `default:"true"`
	disableCaller     bool   `default:"true"`
	initialFields     map[string]interface{}
}

func NewDefaultConfig() iface.Config {
	return &Config{
		logLevel:          "info",
		timeKey:           "timestamp",
		stacktraceKey:     "stacktrace",
		callerKey:         "caller",
		messageKey:        "message",
		timeFormat:        time.RFC3339,
		disableStacktrace: true,
		disableCaller:     true,
	}
}

func (c *Config) SetLogLevel(set string) {
	c.logLevel = set
}

func (c *Config) SetTimeKey(set string) {
	c.timeKey = set
}

func (c *Config) SetStacktraceKey(set string) {
	c.stacktraceKey = set
}

func (c *Config) SetCallerKey(set string) {
	c.callerKey = set
}

func (c *Config) SetMessageKey(set string) {
	c.messageKey = set
}

func (c *Config) SetTimeFormat(set string) {
	c.timeFormat = set
}

func (c *Config) SetDisableStacktrace(set bool) {
	c.disableStacktrace = set
}

func (c *Config) SetDisableCaller(set bool) {
	c.disableCaller = set
}

func (c *Config) SetInitialFields(set map[string]interface{}) {
	c.initialFields = set
}

func (c *Config) LogLevel() string {
	return c.logLevel
}

func (c *Config) TimeKey() string {
	return c.timeKey
}

func (c *Config) StacktraceKey() string {
	return c.stacktraceKey
}

func (c *Config) CallerKey() string {
	return c.callerKey
}

func (c *Config) MessageKey() string {
	return c.messageKey
}

func (c *Config) TimeFormat() string {
	return c.timeFormat
}

func (c *Config) DisableStacktrace() bool {
	return c.disableStacktrace
}

func (c *Config) DisableCaller() bool {
	return c.disableCaller
}

func (c *Config) InitialFields() map[string]interface{} {
	return c.initialFields
}
