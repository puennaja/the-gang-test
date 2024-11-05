package httphdl

import (
	"daveslist/pkg/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (hdl *HTTPHandler) HealthCheck(ctx echo.Context) error {
	resp := map[string]string{"message": "success"}
	return ctx.JSON(http.StatusOK, response.SuccessResponse.SetData(resp))
}
