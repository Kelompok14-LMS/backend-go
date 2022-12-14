package mentees

import (
	"errors"
	"net/http"

	"github.com/Kelompok14-LMS/backend-go/businesses/mentees"
	"github.com/Kelompok14-LMS/backend-go/controllers/mentees/request"
	"github.com/Kelompok14-LMS/backend-go/controllers/mentees/response"
	"github.com/Kelompok14-LMS/backend-go/helper"
	"github.com/Kelompok14-LMS/backend-go/pkg"
	"github.com/labstack/echo/v4"
)

type MenteeController struct {
	menteeUsecase mentees.Usecase
}

func NewMenteeController(menteeUsecase mentees.Usecase) *MenteeController {
	return &MenteeController{
		menteeUsecase: menteeUsecase,
	}
}

func (ctrl *MenteeController) HandlerRegisterMentee(c echo.Context) error {
	menteeInput := request.AuthMenteeInput{}

	if err := c.Bind(&menteeInput); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrInvalidRequest.Error()))
	}

	if err := menteeInput.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(err.Error()))
	}

	err := ctrl.menteeUsecase.Register(menteeInput.ToDomain())

	if err != nil {
		if errors.Is(err, pkg.ErrPasswordLengthInvalid) {
			return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrPasswordLengthInvalid.Error()))
		} else if errors.Is(err, pkg.ErrEmailAlreadyExist) {
			return c.JSON(http.StatusConflict, helper.ConflictResponse(pkg.ErrEmailAlreadyExist.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Success send OTP to email", nil))
}

func (ctrl *MenteeController) HandlerVerifyRegisterMentee(c echo.Context) error {
	menteeInput := request.MenteeRegisterInput{}

	if err := c.Bind(&menteeInput); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrInvalidRequest.Error()))
	}

	if err := menteeInput.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(err.Error()))
	}

	err := ctrl.menteeUsecase.VerifyRegister(menteeInput.ToDomain())

	if err != nil {
		if errors.Is(err, pkg.ErrOTPExpired) {
			return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrOTPExpired.Error()))
		} else if errors.Is(err, pkg.ErrOTPNotMatch) {
			return c.JSON(http.StatusConflict, helper.ConflictResponse(pkg.ErrOTPNotMatch.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	return c.JSON(http.StatusCreated, helper.SuccessCreatedResponse("Register success", nil))
}

func (ctrl *MenteeController) HandlerLoginMentee(c echo.Context) error {
	menteeInput := request.AuthMenteeInput{}

	if err := c.Bind(&menteeInput); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrInvalidRequest.Error()))
	}

	if err := menteeInput.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(err.Error()))
	}

	res, err := ctrl.menteeUsecase.Login(menteeInput.ToDomain())

	if err != nil {
		if errors.Is(err, pkg.ErrUserNotFound) {
			return c.JSON(http.StatusConflict, helper.ConflictResponse(pkg.ErrUserNotFound.Error()))
		} else if errors.Is(err, pkg.ErrMenteeNotFound) {
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrMenteeNotFound.Error()))
		} else if errors.Is(err, pkg.ErrPasswordLengthInvalid) {
			return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrPasswordLengthInvalid.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Login successful", res))
}

func (ctrl *MenteeController) HandlerForgotPassword(c echo.Context) error {
	menteeInput := request.ForgotPasswordInput{}

	if err := c.Bind(&menteeInput); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrInvalidRequest.Error()))
	}

	if err := menteeInput.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(err.Error()))
	}

	err := ctrl.menteeUsecase.ForgotPassword(menteeInput.ToDomain())

	if err != nil {
		if errors.Is(err, pkg.ErrPasswordLengthInvalid) {
			return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrPasswordLengthInvalid.Error()))
		} else if errors.Is(err, pkg.ErrUserNotFound) {
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrUserNotFound.Error()))
		} else if errors.Is(err, pkg.ErrOTPExpired) {
			return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrOTPExpired.Error()))
		} else if errors.Is(err, pkg.ErrOTPNotMatch) {
			return c.JSON(http.StatusConflict, helper.ConflictResponse(pkg.ErrOTPNotMatch.Error()))
		} else if errors.Is(err, pkg.ErrPasswordNotMatch) {
			return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrPasswordNotMatch.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Success reset password", nil))
}

func (ctrl MenteeController) HandlerFindMenteesByCourse(c echo.Context) error {
	courseId := c.Param("courseId")

	data, err := ctrl.menteeUsecase.FindByCourse(courseId)

	if err != nil {
		if errors.Is(err, pkg.ErrCourseNotFound) {
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	var menteeDomain []response.FindAllMentees

	mentees := data["mentees"].(*[]mentees.Domain)

	for _, mentee := range *mentees {
		menteeDomain = append(menteeDomain, response.AllMentees(&mentee))
	}

	res := map[string]interface{}{
		"total":   data["total"],
		"mentees": menteeDomain,
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Success get all mentees", res))
}
