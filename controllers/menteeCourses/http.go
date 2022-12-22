package mentee_courses

import (
	"errors"
	"net/http"

	menteeCourses "github.com/Kelompok14-LMS/backend-go/businesses/menteeCourses"
	"github.com/Kelompok14-LMS/backend-go/controllers/menteeCourses/request"
	"github.com/Kelompok14-LMS/backend-go/controllers/menteeCourses/response"
	"github.com/Kelompok14-LMS/backend-go/helper"
	"github.com/Kelompok14-LMS/backend-go/pkg"
	"github.com/labstack/echo/v4"
)

type MenteeCourseController struct {
	menteeCourseUsecase menteeCourses.Usecase
}

func NewMenteeCourseController(menteeCourseUsecase menteeCourses.Usecase) *MenteeCourseController {
	return &MenteeCourseController{
		menteeCourseUsecase: menteeCourseUsecase,
	}
}

func (ctrl *MenteeCourseController) HandlerEnrollCourse(c echo.Context) error {
	menteeCourseInput := request.EnrollCourse{}

	if err := c.Bind(&menteeCourseInput); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrInvalidRequest.Error()))
	}

	if err := menteeCourseInput.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(err.Error()))
	}

	err := ctrl.menteeCourseUsecase.Enroll(menteeCourseInput.ToDomain())

	if err != nil {
		switch {
		case errors.Is(err, pkg.ErrCourseNotFound):
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrCourseNotFound.Error()))
		case errors.Is(err, pkg.ErrMenteeNotFound):
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrMenteeNotFound.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(pkg.ErrInternalServerError.Error()))
		}
	}

	return c.JSON(http.StatusCreated, helper.SuccessCreatedResponse("Sukses menambahkan kursus", nil))
}

func (ctrl *MenteeCourseController) HandlerFindMenteeCourses(c echo.Context) error {
	title := c.QueryParam("keyword")
	status := c.QueryParam("status")
	menteeId := c.Param("menteeId")

	courses, err := ctrl.menteeCourseUsecase.FindMenteeCourses(menteeId, title, status)

	if err != nil {
		switch {
		case errors.Is(err, pkg.ErrMenteeNotFound):
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrMenteeNotFound.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(pkg.ErrInternalServerError.Error()))
		}
	}

	var menteeCoursesDomain []response.FindMenteeCourses

	for _, course := range *courses {
		menteeCoursesDomain = append(menteeCoursesDomain, *response.MenteeCourses(&course))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Sukses get semua kursus mentee", menteeCoursesDomain))
}

func (ctrl *MenteeCourseController) HandlerCheckEnrollmentCourse(c echo.Context) error {
	courseId := c.Param("courseId")
	menteeId := c.Param("menteeId")

	isEnrolled, err := ctrl.menteeCourseUsecase.CheckEnrollment(menteeId, courseId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(pkg.ErrInternalServerError.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Sukses ceck status enrollment kursus", map[string]interface{}{
		"status_enrollment": isEnrolled,
	}))
}

func (ctrl *MenteeCourseController) HandlerCompleteCourse(c echo.Context) error {
	courseId := c.Param("courseId")
	menteeId := c.Param("menteeId")

	err := ctrl.menteeCourseUsecase.CompleteCourse(menteeId, courseId)

	if err != nil {
		switch {
		case errors.Is(err, pkg.ErrRecordNotFound):
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrRecordNotFound.Error()))
		case errors.Is(err, pkg.ErrCourseNotFound):
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrCourseNotFound.Error()))
		case errors.Is(err, pkg.ErrMenteeNotFound):
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrMenteeNotFound.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(pkg.ErrInternalServerError.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Sukses menyelesaikan kursus", nil))
}
