package httphdl

import (
	"daveslist/internal/core/domain/constant"
	"daveslist/internal/core/domain/dto"
	errors "daveslist/pkg/go-errors"
	"daveslist/pkg/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (hdl *HTTPHandler) CreateListing(c echo.Context) error {
	ctx := c.Request().Context()

	UserID := c.Request().Header.Get(constant.AuthHeaderKey)
	UserName := c.Request().Header.Get(constant.AuthNameKey)
	var request dto.CreateListingRequest
	request.UserID = UserID
	request.UserName = UserName
	if err := c.Bind(&request); err != nil {
		return errors.NewResponseErr(err)
	}

	if err := hdl.validate(&request); err != nil {
		return err
	}

	resp, err := hdl.listingService.CreateListing(ctx, &request)
	if err != nil {
		return errors.NewResponseErr(err)
	}
	return c.JSON(http.StatusCreated, response.SuccessResponse.SetData(resp))
}

func (hdl *HTTPHandler) GetListingList(c echo.Context) error {
	ctx := c.Request().Context()

	var request dto.ListingQuery
	if err := c.Bind(&request); err != nil {
		return errors.NewResponseErr(err)
	}

	if err := hdl.validate(&request); err != nil {
		return err
	}

	resp, err := hdl.listingService.GetListingList(ctx, &request)
	if err != nil {
		return errors.NewResponseErr(err)
	}
	return c.JSON(http.StatusOK, response.SuccessResponse.SetData(resp))
}

func (hdl *HTTPHandler) UpdateListing(c echo.Context) error {
	ctx := c.Request().Context()
	UserID := c.Request().Header.Get(constant.AuthHeaderKey)
	Role := c.Request().Header.Get(constant.AuthRoleKey)
	var request dto.UpdateListingRequest
	request.UserID = UserID
	request.Role = Role
	if err := c.Bind(&request); err != nil {
		return errors.NewResponseErr(err)
	}

	if err := hdl.validate(&request); err != nil {
		return err
	}

	resp, err := hdl.listingService.UpdateListing(ctx, &request)
	if err != nil {
		return errors.NewResponseErr(err)
	}
	return c.JSON(http.StatusOK, response.SuccessResponse.SetData(resp))
}

func (hdl *HTTPHandler) DeleteListing(c echo.Context) error {
	ctx := c.Request().Context()
	UserID := c.Request().Header.Get(constant.AuthHeaderKey)
	Role := c.Request().Header.Get(constant.AuthRoleKey)
	var request dto.DeleteListingRequest
	request.UserID = UserID
	request.Role = Role
	if err := c.Bind(&request); err != nil {
		return errors.NewResponseErr(err)
	}

	if err := hdl.validate(&request); err != nil {
		return err
	}

	resp, err := hdl.listingService.DeleteListing(ctx, &request)
	if err != nil {
		return errors.NewResponseErr(err)
	}
	return c.JSON(http.StatusOK, response.SuccessResponse.SetData(resp))
}

func (hdl *HTTPHandler) HideListing(c echo.Context) error {
	ctx := c.Request().Context()

	var request dto.HideListingRequest
	if err := c.Bind(&request); err != nil {
		return errors.NewResponseErr(err)
	}

	if err := hdl.validate(&request); err != nil {
		return err
	}

	resp, err := hdl.listingService.HideListing(ctx, &request)
	if err != nil {
		return errors.NewResponseErr(err)
	}
	return c.JSON(http.StatusOK, response.SuccessResponse.SetData(resp))
}

func (hdl *HTTPHandler) CreateReplyListing(c echo.Context) error {
	ctx := c.Request().Context()
	UserID := c.Request().Header.Get(constant.AuthHeaderKey)
	UserName := c.Request().Header.Get(constant.AuthNameKey)
	var request dto.CreateReplyListingRequest
	request.UserID = UserID
	request.UserName = UserName
	if err := c.Bind(&request); err != nil {
		return errors.NewResponseErr(err)
	}

	if err := hdl.validate(&request); err != nil {
		return err
	}

	resp, err := hdl.listingService.CreateReplyListing(ctx, &request)
	if err != nil {
		return errors.NewResponseErr(err)
	}
	return c.JSON(http.StatusCreated, response.SuccessResponse.SetData(resp))
}

func (hdl *HTTPHandler) GetReplyListingList(c echo.Context) error {
	ctx := c.Request().Context()

	var request dto.ReplyListingQuery
	if err := c.Bind(&request); err != nil {
		return errors.NewResponseErr(err)
	}

	if err := hdl.validate(&request); err != nil {
		return err
	}

	resp, err := hdl.listingService.GetReplyListingList(ctx, &request)
	if err != nil {
		return errors.NewResponseErr(err)
	}
	return c.JSON(http.StatusOK, response.SuccessResponse.SetData(resp))
}
