package mentors

import (
	"errors"
	"net/http"

	"github.com/Kelompok14-LMS/backend-go/businesses/mentors"
	"github.com/Kelompok14-LMS/backend-go/controllers/mentors/request"
	"github.com/Kelompok14-LMS/backend-go/controllers/mentors/response"
	"github.com/Kelompok14-LMS/backend-go/utils"

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
	mentorInput := request.MentorUpdatePassword{}

	if err := c.Bind(&mentorInput); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrInvalidRequest.Error()))
	}

	user, err := utils.GetUserID(c)
	if err != nil {
		return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrUserNotFound.Error()))
	}

	mentorInput.UserID = user.UserId

	if err := mentorInput.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(err.Error()))
	}

	err = ctrl.mentorUsecase.UpdatePassword(mentorInput.ToDomain())

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

func (ctrl *MentorController) HandlerUpdateProfile(c echo.Context) error {
	mentorInput := request.MentorUpdateProfile{}

	if err := c.Bind(&mentorInput); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrInvalidRequest.Error()))
	}

	user, err := utils.GetUserID(c)
	if err != nil {
		return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrUserNotFound.Error()))
	}

	mentorInput.ID = user.MentorId
	mentorInput.UserID = user.UserId

	if err := mentorInput.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(err.Error()))
	}

	err = ctrl.mentorUsecase.Update(mentorInput.ToDomain())

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

	return c.JSON(http.StatusOK, helper.SuccessResponse("Success update profile", nil))
}

func (ctrl *MentorController) HandlerFindByID(c echo.Context) error {
	user, err := utils.GetUserID(c)
	if err != nil {
		return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrUserNotFound.Error()))
	}

	mentor, err := ctrl.mentorUsecase.FindById(user.MentorId)

	if err != nil {
		if errors.Is(err, pkg.ErrUserNotFound) {
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrUserNotFound.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Mentor found", response.FromDomainUser(mentor)))
}

func (ctrl *MentorController) HandlerFindAll(c echo.Context) error {

	mentors, err := ctrl.mentorUsecase.FindAll()

	allMentor := []response.FindMentorAll{}

	for _, mentor := range *mentors {
		allMentor = append(allMentor, *response.FromDomainAll(&mentor))
	}

	if err != nil {
		if errors.Is(err, pkg.ErrUserNotFound) {
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrUserNotFound.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("All mentor found", allMentor))
}
