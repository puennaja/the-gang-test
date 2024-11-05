package middleware

import (
	"encoding/json"

	gologger "daveslist/pkg/go-logger/iface"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Logger(logger gologger.Logger) echo.MiddlewareFunc {
	return middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		var body, resp interface{}

		_ = json.Unmarshal(reqBody, &body)
		_ = json.Unmarshal(resBody, &resp)

		if c.Response().Status >= 200 && c.Response().Status < 300 {
			logger.InfoW(
				"success",
				"method", c.Request().Method,
				"path", c.Request().URL,
				"body", body,
				"resp", resp,
			)
		} else if c.Response().Status >= 400 {
			logger.WarnW(
				"warn",
				"method", c.Request().Method,
				"path", c.Request().URL,
				"body", body,
				"resp", resp,
			)
		} else if c.Response().Status > 499 {
			logger.ErrorW(
				"error",
				"method", c.Request().Method,
				"path", c.Request().URL,
				"body", body,
				"resp", resp,
			)
		} else {
			logger.ErrorW(
				"error",
				"method", c.Request().Method,
				"path", c.Request().URL,
				"body", body,
				"resp", resp,
			)
		}
	})
}
