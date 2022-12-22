package mentors

import (

=======
	"context"
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
	jwtConfig     *utils.JWTConfig
}

func NewMentorController(mentorUsecase mentors.Usecase, jwtConfig *utils.JWTConfig) *MentorController {
	return &MentorController{
		mentorUsecase: mentorUsecase,
		jwtConfig:     jwtConfig,
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
		switch err {
		case pkg.ErrPasswordLengthInvalid:
			return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrPasswordLengthInvalid.Error()))
		case pkg.ErrEmailAlreadyExist:
			return c.JSON(http.StatusConflict, helper.ConflictResponse(pkg.ErrEmailAlreadyExist.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(pkg.ErrInternalServerError.Error()))
		}
	}

	return c.JSON(http.StatusCreated, helper.SuccessCreatedResponse("Register berhasil", nil))
}

func (ctrl *MentorController) HandlerLoginMentor(c echo.Context) error {
	mentorInput := request.AuthMentorInput{}

	if err := c.Bind(&mentorInput); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrInvalidRequest.Error()))
	}

	if err := mentorInput.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(err.Error()))
	}

	res, err := ctrl.mentorUsecase.Login(mentorInput.ToDomain())

	if err != nil {
		switch err {
		case pkg.ErrUserNotFound:
			return c.JSON(http.StatusConflict, helper.ConflictResponse(pkg.ErrAuthenticationFailed.Error()))
		case pkg.ErrMentorNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrMentorNotFound.Error()))
		case pkg.ErrPasswordLengthInvalid:
			return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrPasswordLengthInvalid.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(pkg.ErrInternalServerError.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Login berhasil", res))
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
		switch err {
		case pkg.ErrPasswordLengthInvalid:
			return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrPasswordLengthInvalid.Error()))
		case pkg.ErrPasswordNotMatch:
			return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrPasswordNotMatch.Error()))
		case pkg.ErrUserNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrUserNotFound.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(pkg.ErrInternalServerError.Error()))
		}

	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Sukses update kata sandi", nil))
}

func (ctrl *MentorController) HandlerFindByID(c echo.Context) error {
	var id string = c.Param("mentorId")

	mentor, err := ctrl.mentorUsecase.FindById(id)

	if err != nil {
		switch err {
		case pkg.ErrMentorNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrMentorNotFound.Error()))
		case pkg.ErrUserNotFound:
			return c.JSON(http.StatusConflict, helper.ConflictResponse(pkg.ErrAuthenticationFailed.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(pkg.ErrInternalServerError.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Sukses get mentor berdasarkan id", response.FromDomainUser(mentor)))
}

func (ctrl *MentorController) HandlerProfileMentor(c echo.Context) error {
	token, _ := ctrl.jwtConfig.ExtractToken(c)

	mentor, err := ctrl.mentorUsecase.FindById(token.MentorId)

	if err != nil {
		switch err {
		case pkg.ErrMentorNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrMentorNotFound.Error()))
		case pkg.ErrUserNotFound:
			return c.JSON(http.StatusConflict, helper.ConflictResponse(pkg.ErrAuthenticationFailed.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(pkg.ErrInternalServerError.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Sukses get mentor berdasarkan token header", response.FromDomainUser(mentor)))
}

func (ctrl *MentorController) HandlerFindAll(c echo.Context) error {

	mentors, err := ctrl.mentorUsecase.FindAll()

	allMentor := []response.FindMentorAll{}

	for _, mentor := range *mentors {
		allMentor = append(allMentor, *response.FromDomainAll(&mentor))
	}

	if err != nil {
		switch err {
		case pkg.ErrMentorNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrMentorNotFound.Error()))
		case pkg.ErrUserNotFound:
			return c.JSON(http.StatusConflict, helper.ConflictResponse(pkg.ErrAuthenticationFailed.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(pkg.ErrInternalServerError.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Sukses get semua mentor", allMentor))
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

	ctx := context.Background()

	err := ctrl.mentorUsecase.Update(ctx, mentorId, mentorInput.ToDomain())

	if err != nil {
		switch err {
		case pkg.ErrInvalidRequest:
			return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrInvalidRequest.Error()))
		case pkg.ErrUnsupportedImageFile:
			return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrUnsupportedImageFile.Error()))
		case pkg.ErrMentorNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrMentorNotFound.Error()))
		case pkg.ErrUserNotFound:
			return c.JSON(http.StatusConflict, helper.ConflictResponse(pkg.ErrAuthenticationFailed.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(pkg.ErrInternalServerError.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Sukses update profile", nil))
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
		switch err {
		case pkg.ErrMentorNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrMentorNotFound.Error()))
		case pkg.ErrUserNotFound:
			return c.JSON(http.StatusConflict, helper.ConflictResponse(pkg.ErrAuthenticationFailed.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(pkg.ErrInternalServerError.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Sukses ganti kata sandi", nil))
}
