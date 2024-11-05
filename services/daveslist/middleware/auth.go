package middleware

import (
	"daveslist/internal/core/domain/constant"
	"daveslist/internal/core/domain/dto"
	"daveslist/internal/core/port"

	errors "daveslist/pkg/go-errors"

	"github.com/labstack/echo/v4"
)

func Auth(authSvc port.AuthService) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.Request().Context()
			role := c.Request().Header.Get(constant.AuthRoleKey)
			path := c.Request().URL.Path
			method := c.Request().Method

			// api access
			access, err := authSvc.Authorize(ctx, &dto.AuthorizeRequest{
				Role:   role,
				Object: path,
				Action: method,
			})
			if err != nil {
				return err
			}
			if !access {
				return errors.NewResponseErr(errors.ErrPermissionDenied)
			}

			// internal service access level
			level1, _ := authSvc.Authorize(ctx, &dto.AuthorizeRequest{
				Role:   role,
				Object: path,
				Action: method + "-" + constant.AuthLevel1,
			})
			if level1 {
				c.Request().Header.Set(constant.AuthLevelKey, constant.AuthLevel1)
				return next(c)
			}

			level2, _ := authSvc.Authorize(ctx, &dto.AuthorizeRequest{
				Role:   role,
				Object: path,
				Action: method + "-" + constant.AuthLevel2,
			})
			if level2 {
				c.Request().Header.Set(constant.AuthLevelKey, constant.AuthLevel2)
				return next(c)
			}

			level3, _ := authSvc.Authorize(ctx, &dto.AuthorizeRequest{
				Role:   role,
				Object: path,
				Action: method + "-" + constant.AuthLevel3,
			})
			if level3 {
				c.Request().Header.Set(constant.AuthLevelKey, constant.AuthLevel3)
				return next(c)
			}

			c.Request().Header.Set(constant.AuthLevelKey, constant.AuthLevel0)
			return next(c)
		}
	}
}
