package httphdl

import (
	"daveslist/internal/core/domain/dto"
	"net/http"

	errors "daveslist/pkg/go-errors"
	"daveslist/pkg/response"

	"github.com/labstack/echo/v4"
)

func (hdl *HTTPHandler) CreateCategory(c echo.Context) error {
	ctx := c.Request().Context()

	var request dto.CreateCategoryRequest
	if err := c.Bind(&request); err != nil {
		return errors.NewResponseErr(err)
	}

	if err := hdl.validate(&request); err != nil {
		return err
	}

	resp, err := hdl.categoryService.CreateCategory(ctx, &request)
	if err != nil {
		return errors.NewResponseErr(err)
	}
	return c.JSON(http.StatusCreated, response.SuccessResponse.SetData(resp))
}

func (hdl *HTTPHandler) GetCategoryList(c echo.Context) error {
	var request dto.CategoryQuery
	if err := c.Bind(&request); err != nil {
		return errors.NewResponseErr(err)
	}

	if err := hdl.validate(&request); err != nil {
		return err
	}

	ctx := c.Request().Context()
	resp, err := hdl.categoryService.GetCategoryList(ctx, &request)
	if err != nil {
		return errors.NewResponseErr(err)
	}
	return c.JSON(http.StatusOK, response.SuccessResponse.SetData(resp))
}

func (hdl *HTTPHandler) DeleteCategory(c echo.Context) error {
	ctx := c.Request().Context()

	var request dto.DeleteCategoryRequest
	if err := c.Bind(&request); err != nil {
		return errors.NewResponseErr(err)
	}

	if err := hdl.validate(&request); err != nil {
		return err
	}

	resp, err := hdl.categoryService.DeleteCategory(ctx, request.ID)
	if err != nil {
		return errors.NewResponseErr(err)
	}
	return c.JSON(http.StatusOK, response.SuccessResponse.SetData(resp))
}
