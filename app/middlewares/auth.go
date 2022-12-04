package middlewares

import (
	"net/http"

	"github.com/Kelompok14-LMS/backend-go/helper"
	"github.com/Kelompok14-LMS/backend-go/utils"
	"github.com/labstack/echo/v4"
)

func CheckTokenMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, _ := utils.ExtractToken(c)

		if token.Role == "mentee" {
			return c.JSON(http.StatusForbidden, helper.ForbiddenResponse("User tidak diizinkan"))

		}

		return next(c)
	}
}
