package detail_course

import (
	"errors"
	"net/http"

	detailCourse "github.com/Kelompok14-LMS/backend-go/businesses/detailCourse"
	"github.com/Kelompok14-LMS/backend-go/controllers/detailCourse/response"
	"github.com/Kelompok14-LMS/backend-go/helper"
	"github.com/Kelompok14-LMS/backend-go/pkg"
	"github.com/labstack/echo/v4"
)

type DetailCourseController struct {
	detailCourseUsecase detailCourse.Usecase
}

func NewDetailCourseController(detailCourseUsecase detailCourse.Usecase) *DetailCourseController {
	return &DetailCourseController{
		detailCourseUsecase: detailCourseUsecase,
	}
}

func (ctrl *DetailCourseController) HandlerDetailCourse(c echo.Context) error {
	courseId := c.Param("courseId")

	course, err := ctrl.detailCourseUsecase.DetailCourse(courseId)

	if err != nil {
		switch {
		case errors.Is(err, pkg.ErrCourseNotFound):
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(pkg.ErrInternalServerError.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Sukses get detail kursus", response.FullDetailCourse(course)))
}

func (ctrl *DetailCourseController) HandlerDetailCourseEnrolled(c echo.Context) error {
	menteeId := c.Param("menteeId")
	courseId := c.Param("courseId")

	course, err := ctrl.detailCourseUsecase.DetailCourseEnrolled(menteeId, courseId)

	if err != nil {
		switch {
		case errors.Is(err, pkg.ErrMenteeNotFound):
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrMenteeNotFound.Error()))
		case errors.Is(err, pkg.ErrCourseNotFound):
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrCourseNotFound.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(pkg.ErrInternalServerError.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Sukses get detail kursus yang ter-enroll", response.FullDetailCourseEnrolled(course)))
}
