package otp

import (
	"errors"
	"net/http"

	otpDomain "github.com/Kelompok14-LMS/backend-go/businesses/otp"
	"github.com/Kelompok14-LMS/backend-go/controllers/otp/request"
	"github.com/Kelompok14-LMS/backend-go/helper"
	"github.com/Kelompok14-LMS/backend-go/pkg"
	"github.com/labstack/echo/v4"
)

type OTPController struct {
	otpUsecase otpDomain.Usecase
}

func NewOTPController(otpUsecase otpDomain.Usecase) *OTPController {
	return &OTPController{
		otpUsecase: otpUsecase,
	}
}

func (oc OTPController) HandlerSendOTP(c echo.Context) error {
	otpInput := request.OTP{}

	if err := c.Bind(&otpInput); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrInvalidRequest.Error()))
	}

	if err := otpInput.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(err.Error()))
	}

	err := oc.otpUsecase.SendOTP(otpInput.ToDomain())

	if err != nil {
		switch {
		case errors.Is(err, pkg.ErrUserNotFound):
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrUserNotFound.Error()))

		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(pkg.ErrInternalServerError.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Sukses kirim OTP ke email", nil))
}

func (oc OTPController) HandlerCheckOTP(c echo.Context) error {
	otpInput := request.CheckOTP{}

	if err := c.Bind(&otpInput); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrInvalidRequest.Error()))
	}

	if err := otpInput.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(err.Error()))
	}

	err := oc.otpUsecase.CheckOTP(otpInput.ToDomain())

	if err != nil {
		switch {
		case errors.Is(err, pkg.ErrUserNotFound):
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrUserNotFound.Error()))
		case errors.Is(err, pkg.ErrOTPExpired):
			return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrOTPExpired.Error()))
		case errors.Is(err, pkg.ErrOTPNotMatch):
			return c.JSON(http.StatusConflict, helper.ConflictResponse(pkg.ErrOTPNotMatch.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(pkg.ErrInternalServerError.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("OTP matched", nil))
}
