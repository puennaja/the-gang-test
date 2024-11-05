package middleware

import (
	errors "daveslist/pkg/go-errors"

	"github.com/labstack/echo/v4"
)

func ErrorHandler(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}

	if errResponse := errors.IsResponseErr(err); errResponse != nil {
		_ = c.JSON(errResponse.GetStatus(), errResponse)
		return
	}

	if errEcho, ok := err.(*echo.HTTPError); ok {
		err := errors.NewResponseErr(errEcho.Unwrap())
		err.SetStatus(errEcho.Code)
		_ = c.JSON(err.GetStatus(), err)
		return
	}

	defaultErr := errors.NewResponseErr(err)
	_ = c.JSON(defaultErr.GetStatus(), defaultErr)
}
