package httphdl

import (
	"daveslist/internal/core/domain/constant"
	"daveslist/internal/core/domain/dto"
	errors "daveslist/pkg/go-errors"
	"daveslist/pkg/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (hdl *HTTPHandler) CreateMessage(c echo.Context) error {
	ctx := c.Request().Context()
	UserID := c.Request().Header.Get(constant.AuthHeaderKey)
	UserName := c.Request().Header.Get(constant.AuthNameKey)
	var request dto.CreateMessageRequest
	request.SenderID = UserID
	request.SenderName = UserName
	if err := c.Bind(&request); err != nil {
		return errors.NewResponseErr(err)
	}

	if err := hdl.validate(&request); err != nil {
		return err
	}

	resp, err := hdl.messageService.CreateMessage(ctx, &request)
	if err != nil {
		return errors.NewResponseErr(err)
	}
	return c.JSON(http.StatusCreated, response.SuccessResponse.SetData(resp))
}

func (hdl *HTTPHandler) GetMessageList(c echo.Context) error {
	ctx := c.Request().Context()

	senderID := c.Request().Header.Get("x-user-id")
	var request dto.MessageQuery
	request.SenderID = senderID
	if err := c.Bind(&request); err != nil {
		return errors.NewResponseErr(err)
	}

	if err := hdl.validate(&request); err != nil {
		return err
	}

	resp, err := hdl.messageService.GetMessageList(ctx, &request)
	if err != nil {
		return errors.NewResponseErr(err)
	}
	return c.JSON(http.StatusOK, response.SuccessResponse.SetData(resp))
}
