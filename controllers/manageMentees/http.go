package manage_mentees

import (
	"net/http"

	manageMentees "github.com/Kelompok14-LMS/backend-go/businesses/manageMentees"
	"github.com/Kelompok14-LMS/backend-go/helper"
	"github.com/Kelompok14-LMS/backend-go/pkg"
	"github.com/labstack/echo/v4"
)

type ManageMenteeController struct {
	manageMenteeUsecase manageMentees.Usecase
}

func NewManageMenteeController(manageMenteeUsecase manageMentees.Usecase) *ManageMenteeController {
	return &ManageMenteeController{
		manageMenteeUsecase: manageMenteeUsecase,
	}
}

func (ctrl *ManageMenteeController) HandlerDeleteAccessMentee(c echo.Context) error {
	courseId := c.Param("courseId")
	menteeId := c.Param("menteeId")

	err := ctrl.manageMenteeUsecase.DeleteAccess(menteeId, courseId)

	if err != nil {
		switch err {
		case pkg.ErrRecordNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrRecordNotFound.Error()))
		case pkg.ErrCourseNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrCourseNotFound.Error()))
		case pkg.ErrMenteeNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrMenteeNotFound.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(pkg.ErrInternalServerError.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Sukses hapus akses kursus mentee", nil))
}
