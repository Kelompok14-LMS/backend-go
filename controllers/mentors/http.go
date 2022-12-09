package mentors

import (
	"errors"
	"net/http"

	"github.com/Kelompok14-LMS/backend-go/businesses/mentors"
	"github.com/Kelompok14-LMS/backend-go/controllers/mentors/request"
	"github.com/Kelompok14-LMS/backend-go/controllers/mentors/response"

	"github.com/Kelompok14-LMS/backend-go/helper"
	"github.com/Kelompok14-LMS/backend-go/pkg"
	"github.com/labstack/echo/v4"
)

type MentorController struct {
	mentorUsecase mentors.Usecase
}

func NewMentorController(mentorUsecase mentors.Usecase) *MentorController {
	return &MentorController{
		mentorUsecase: mentorUsecase,
	}
}

func (ctrl *MentorController) HandlerRegisterMentor(c echo.Context) error {
	mentorInput := request.MentorRegisterInput{}

	if err := c.Bind(&mentorInput); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrInvalidRequest.Error()))
	}

	if err := mentorInput.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(err.Error()))
	}

	err := ctrl.mentorUsecase.Register(mentorInput.ToDomain())
	if err != nil {
		if errors.Is(err, pkg.ErrPasswordLengthInvalid) {
			return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrPasswordLengthInvalid.Error()))
		} else if errors.Is(err, pkg.ErrEmailAlreadyExist) {
			return c.JSON(http.StatusConflict, helper.ConflictResponse(pkg.ErrEmailAlreadyExist.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	return c.JSON(http.StatusCreated, helper.SuccessCreatedResponse("Register success", nil))
}

func (ctrl *MentorController) HandlerLoginMentor(c echo.Context) error {
	mentorInput := request.AuthMentorInput{}

	if err := c.Bind(&mentorInput); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrInvalidRequest.Error()))
	}

	if err := mentorInput.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(err.Error()))
	}

	token, err := ctrl.mentorUsecase.Login(mentorInput.ToDomain())

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

	data := map[string]interface{}{
		"token": token,
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Login successful", data))
}

func (ctrl *MentorController) HandlerUpdatePassword(c echo.Context) error {
	mentorId := c.Param("mentorId")
	mentorInput := request.MentorUpdatePassword{}

	if err := c.Bind(&mentorInput); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrInvalidRequest.Error()))
	}
	mentorInput.UserID = mentorId

	if err := mentorInput.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(err.Error()))
	}

	err := ctrl.mentorUsecase.UpdatePassword(mentorInput.ToDomain())

	if err != nil {
		if errors.Is(err, pkg.ErrPasswordLengthInvalid) {
			return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrPasswordLengthInvalid.Error()))
		} else if errors.Is(err, pkg.ErrPasswordNotMatch) {
			return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrPasswordNotMatch.Error()))
		} else if errors.Is(err, pkg.ErrUserNotFound) {
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrUserNotFound.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Success update password", nil))
}

func (ctrl *MentorController) HandlerFindByID(c echo.Context) error {

	var id string = c.Param("mentorId")

	mentor, err := ctrl.mentorUsecase.FindById(id)

	if err != nil {
		if errors.Is(err, pkg.ErrMentorNotFound) {
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrMentorNotFound.Error()))
		} else if errors.Is(err, pkg.ErrUserNotFound) {
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrUserNotFound.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Success get Mentor by id", response.FromDomainUser(mentor)))
}

func (ctrl *MentorController) HandlerFindAll(c echo.Context) error {

	mentors, err := ctrl.mentorUsecase.FindAll()

	allMentor := []response.FindMentorAll{}

	for _, mentor := range *mentors {
		allMentor = append(allMentor, *response.FromDomainAll(&mentor))
	}

	if err != nil {
		if errors.Is(err, pkg.ErrMentorNotFound) {
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrMentorNotFound.Error()))
		} else if errors.Is(err, pkg.ErrUserNotFound) {
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrUserNotFound.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Success get all mentor ", allMentor))
}

func (ctrl *MentorController) HandlerUpdateProfile(c echo.Context) error {
	mentorInput := request.MentorUpdateProfile{}

	mentorId := c.Param("mentorId")

	ProfilePictureFile, _ := c.FormFile("profile_picture")

	if ProfilePictureFile != nil {
		mentorInput.ProfilePictureFile = ProfilePictureFile

		if err := c.Bind(&mentorInput); err != nil {
			return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrInvalidRequest.Error()))
		}
	} else {
		mentorInput.UserID = c.FormValue("user_id")
		mentorInput.Fullname = c.FormValue("fullname")
		mentorInput.Email = c.FormValue("email")
		mentorInput.Phone = c.FormValue("phone")
		mentorInput.Jobs = c.FormValue("jobs")
		mentorInput.Gender = c.FormValue("gender")
		mentorInput.BirthDate = c.FormValue("birth_date")
		mentorInput.BirthPlace = c.FormValue("birth_place")
		mentorInput.ProfilePictureFile = nil
	}

	if err := mentorInput.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(err.Error()))
	}

	err := ctrl.mentorUsecase.Update(mentorId, mentorInput.ToDomain())

	if err != nil {
		if errors.Is(err, pkg.ErrInvalidRequest) {
			return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrInvalidRequest.Error()))
		} else if errors.Is(err, pkg.ErrMentorNotFound) {
			return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrMentorNotFound.Error()))
		} else if errors.Is(err, pkg.ErrUserNotFound) {
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrUserNotFound.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Success update profile", nil))
}

func (ctrl *MentorController) HandlerForgotPassword(c echo.Context) error {
	mentorInput := request.ForgotPasswordInput{}

	if err := c.Bind(&mentorInput); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrInvalidRequest.Error()))
	}

	if err := mentorInput.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(err.Error()))
	}

	err := ctrl.mentorUsecase.ForgotPassword(mentorInput.ToDomain())

	if err != nil {
		if errors.Is(err, pkg.ErrUserNotFound) {
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrUserNotFound.Error()))
		} else if errors.Is(err, pkg.ErrOTPExpired) {
			return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrOTPExpired.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Success reset password", nil))
}
