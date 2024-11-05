package goerrors

import (
	"errors"
	"fmt"
)

var cfg ConfigIFace

func New(text string) error {
	return errors.New(text)
}

func Newf(format string, args ...interface{}) error {
	return fmt.Errorf(format, args...)
}

func Is(err, target error) bool {
	return errors.Is(err, target)
}

func As(err error, target any) bool {
	return errors.As(err, target)
}

func Unwrap(err error) error {
	if e, ok := err.(interface{ Unwrap() error }); ok {
		return e.Unwrap()
	}
	return nil
}

func Init() {
	SetConfig(InternalServerErrorStatus, DefaultCode, messages[DefaultCode])
}

func SetConfig(staus, code int, message string) {
	newCfg := NewConfig()
	newCfg.SetDefaultStatus(staus)
	newCfg.SetDefaultCode(code)
	newCfg.SetDefaultMessage(message)
	cfg = newCfg
}

func NewInternalErr(status, code int, message string, err ...error) InternalErrIFace {
	return newInternalErr(status, code, message, err...)
}

func IsInternalErr(err error) InternalErrIFace {
	if i, ok := err.(InternalErrIFace); ok {
		return i
	}
	return nil
}

func NewResponseErr(err error) *ResponseErr {
	return newResponseErr(cfg, err)
}

func IsResponseErr(err error) *ResponseErr {
	if i, ok := err.(*ResponseErr); ok {
		return i
	}
	return nil
}
