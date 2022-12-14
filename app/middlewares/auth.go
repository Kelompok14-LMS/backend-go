package middlewares

import (
	"net/http"

	"github.com/Kelompok14-LMS/backend-go/helper"
	"github.com/Kelompok14-LMS/backend-go/pkg"
	"github.com/Kelompok14-LMS/backend-go/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type AuthMiddleware struct {
	jwtConfig *utils.JWTConfig
}

func NewAuthMiddleware(jwtConfig *utils.JWTConfig) *AuthMiddleware {
	return &AuthMiddleware{
		jwtConfig: jwtConfig,
	}
}

// IsMentor custom middleware to check user role is mentor
func (mid *AuthMiddleware) IsMentor(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		payloads, err := mid.jwtConfig.ExtractToken(c)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.UnauthorizedResponse(err.Error()))
		}

		if payloads.Role != "mentor" {
			return c.JSON(http.StatusForbidden, helper.ForbiddenResponse(pkg.ErrAccessForbidden.Error()))
		}

		return next(c)
	}
}

// custom middleware to check user role is mentee
func (mid *AuthMiddleware) IsMentee(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		payloads, err := mid.jwtConfig.ExtractToken(c)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.UnauthorizedResponse(err.Error()))
		}

		if payloads.Role != "mentee" {
			return c.JSON(http.StatusForbidden, helper.ForbiddenResponse(pkg.ErrAccessForbidden.Error()))
		}

		return next(c)
	}
}

// IsAuthenticated function wrapper `echo` middleware.JWT
func (mid *AuthMiddleware) IsAuthenticated() echo.MiddlewareFunc {
	return middleware.JWT([]byte(mid.jwtConfig.JWTSecret))
}
