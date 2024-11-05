package goerrors

type Config struct {
	defaultStatus  int
	defaultCode    int
	defaultMessage string
}

func NewConfig() *Config {
	return &Config{
		defaultStatus:  InternalServerErrorStatus,
		defaultCode:    DefaultCode,
		defaultMessage: messages[DefaultCode],
	}
}

// SetDefaultStatus set default status
func (c *Config) SetDefaultStatus(status int) {
	c.defaultStatus = status
}

// SetDefaultCode set default code
func (c *Config) SetDefaultCode(code int) {
	c.defaultCode = code
}

// SetDefaultMessage set default message
func (c *Config) SetDefaultMessage(message string) {
	c.defaultMessage = message
}

// DefaultStatus get default status
func (c *Config) DefaultStatus() int {
	return c.defaultStatus
}

// DefaultCode get default code
func (c *Config) DefaultCode() int {
	return c.defaultCode
}

// DefaultMessage get default message
func (c *Config) DefaultMessage() string {
	return c.defaultMessage
}
