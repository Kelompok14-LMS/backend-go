package mentees

import (
	"net/http"
	"strconv"

	"github.com/Kelompok14-LMS/backend-go/businesses/mentees"
	"github.com/Kelompok14-LMS/backend-go/controllers/mentees/request"
	"github.com/Kelompok14-LMS/backend-go/controllers/mentees/response"
	"github.com/Kelompok14-LMS/backend-go/helper"
	"github.com/Kelompok14-LMS/backend-go/pkg"
	"github.com/Kelompok14-LMS/backend-go/utils"
	"github.com/labstack/echo/v4"
)

type MenteeController struct {
	menteeUsecase mentees.Usecase
	jwtConfig     *utils.JWTConfig
}

func NewMenteeController(menteeUsecase mentees.Usecase, jwtConfig *utils.JWTConfig) *MenteeController {
	return &MenteeController{
		menteeUsecase: menteeUsecase,
		jwtConfig:     jwtConfig,
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
		switch err {
		case pkg.ErrPasswordLengthInvalid:
			return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrPasswordLengthInvalid.Error()))
		case pkg.ErrEmailAlreadyExist:
			return c.JSON(http.StatusConflict, helper.ConflictResponse(pkg.ErrEmailAlreadyExist.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(pkg.ErrInternalServerError.Error()))
		}

	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Sukses mengirim OTP ke email", nil))
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
		switch err {
		case pkg.ErrOTPExpired:
			return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrOTPExpired.Error()))
		case pkg.ErrOTPNotMatch:
			return c.JSON(http.StatusConflict, helper.ConflictResponse(pkg.ErrOTPNotMatch.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(pkg.ErrInternalServerError.Error()))
		}

	}

	return c.JSON(http.StatusCreated, helper.SuccessCreatedResponse("Register berhasil", nil))
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
		switch err {
		case pkg.ErrUserNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrUserNotFound.Error()))
		case pkg.ErrMenteeNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrMenteeNotFound.Error()))
		case pkg.ErrPasswordLengthInvalid:
			return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrPasswordLengthInvalid.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(pkg.ErrInternalServerError.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Login berhasil", res))
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
		switch err {
		case pkg.ErrPasswordLengthInvalid:
			return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrPasswordLengthInvalid.Error()))
		case pkg.ErrUserNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrUserNotFound.Error()))
		case pkg.ErrOTPExpired:
			return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrOTPExpired.Error()))
		case pkg.ErrOTPNotMatch:
			return c.JSON(http.StatusConflict, helper.ConflictResponse(pkg.ErrOTPNotMatch.Error()))
		case pkg.ErrPasswordNotMatch:
			return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrPasswordNotMatch.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(pkg.ErrInternalServerError.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Sukses ganti kata sandi", nil))
}

func (ctrl MenteeController) HandlerFindMenteesByCourse(c echo.Context) error {
	courseId := c.Param("courseId")

	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	page, _ := strconv.Atoi(c.QueryParam("page"))

	pagination := pkg.Pagination{
		Limit: limit,
		Page:  page,
	}

	res, err := ctrl.menteeUsecase.FindByCourse(courseId, pagination)

	if err != nil {
		switch err {
		case pkg.ErrCourseNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrCourseNotFound.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(pkg.ErrInternalServerError.Error()))
		}
	}

	var menteeDomain []response.FindAllMentees

	mentees := res.Result.(*[]mentees.Domain)

	for _, mentee := range *mentees {
		menteeDomain = append(menteeDomain, response.AllMentees(&mentee))
	}

	res.Result = menteeDomain

	return c.JSON(http.StatusOK, helper.SuccessResponse("Sukses get semua mentee berdasarkan kursus", res))
}

func (ctrl *MenteeController) HandlerFindByID(c echo.Context) error {
	var id string = c.Param("menteeId")

	mentee, err := ctrl.menteeUsecase.FindById(id)

	if err != nil {
		switch err {
		case pkg.ErrMenteeNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrMenteeNotFound.Error()))
		case pkg.ErrUserNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrUserNotFound.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(pkg.ErrInternalServerError.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Sukses get mentee berdasarkan id", response.FromDomain(mentee)))
}

func (ctrl *MenteeController) HandlerProfileMentee(c echo.Context) error {
	token, _ := ctrl.jwtConfig.ExtractToken(c)

	mentee, err := ctrl.menteeUsecase.FindById(token.MenteeId)

	if err != nil {
		switch err {
		case pkg.ErrMenteeNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrMenteeNotFound.Error()))
		case pkg.ErrUserNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrUserNotFound.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(pkg.ErrInternalServerError.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Sukses get mentee berdasarkan token header", response.FromDomain(mentee)))
}

func (ctrl *MenteeController) HandlerFindAll(c echo.Context) error {

	mentees, err := ctrl.menteeUsecase.FindAll()

	allMentees := []response.FindAllMentees{}

	for _, mentee := range *mentees {
		allMentees = append(allMentees, response.AllMentees(&mentee))
	}

	if err != nil {
		switch err {
		case pkg.ErrMenteeNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrMenteeNotFound.Error()))
		case pkg.ErrUserNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrUserNotFound.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(pkg.ErrInternalServerError.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Sukses get semua mentee", allMentees))
}

func (ctrl *MenteeController) HandlerUpdateProfile(c echo.Context) error {
	menteeInput := request.MenteeUpdateProfile{}

	menteeId := c.Param("menteeId")

	ProfilePictureFile, _ := c.FormFile("profile_picture")

	if ProfilePictureFile != nil {
		menteeInput.ProfilePictureFile = ProfilePictureFile

		if err := c.Bind(&menteeInput); err != nil {
			return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrInvalidRequest.Error()))
		}
	} else {
		menteeInput.Fullname = c.FormValue("fullname")
		menteeInput.Phone = c.FormValue("phone")
		menteeInput.BirthDate = c.FormValue("birth_date")
		menteeInput.ProfilePictureFile = nil
	}

	if err := menteeInput.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(err.Error()))
	}

	err := ctrl.menteeUsecase.Update(menteeId, menteeInput.ToDomain())

	if err != nil {
		switch err {
		case pkg.ErrInvalidRequest:
			return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrInvalidRequest.Error()))
		case pkg.ErrMenteeNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrMenteeNotFound.Error()))
		case pkg.ErrUnsupportedImageFile:
			return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrUnsupportedImageFile.Error()))
		case pkg.ErrUserNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrUserNotFound.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(pkg.ErrInternalServerError.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Sukses update profil mentee", nil))
}
