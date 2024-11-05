package goerrors

type Errors interface {
	New(text string) error
	Newf(format string, args ...interface{}) error
	Is(err, target error) bool
	As(err error) error
	Unwrap(err error) error
	NewInternalErr(status, code int, message string, err ...error) InternalErrIFace
	NewResponseErr(err error) ResponseErrIFace
}

type ConfigIFace interface {
	SetDefaultStatus(status int)
	SetDefaultCode(code int)
	SetDefaultMessage(message string)
	DefaultStatus() int
	DefaultCode() int
	DefaultMessage() string
}

type InternalErrIFace interface {
	SetStatus(status int)
	SetCode(code int)
	SetMessage(message string)
	SetError(err ...error)
	GetStatus() int
	GetCode() int
	GetMessage() string
	GetErrors() []error

	// error
	Error() string
}

type ResponseErrIFace interface {
	SetStatus(status int)
	SetCode(code int)
	SetMessage(message string)
	SetError(err ...error)
	GetStatus() int
	GetCode() int
	GetMessage() string
	GetErrors() []error

	// error
	Error() string
}
