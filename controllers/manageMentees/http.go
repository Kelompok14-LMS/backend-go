package manage_mentees

import (
	"errors"
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
		if errors.Is(err, pkg.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Success delete access mentee", nil))
}
