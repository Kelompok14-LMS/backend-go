package details_assignments

import (
	"errors"
	"net/http"

	detailAssignment "github.com/Kelompok14-LMS/backend-go/businesses/detailsAssignments"
	"github.com/Kelompok14-LMS/backend-go/controllers/detailsAssignments/response"
	"github.com/Kelompok14-LMS/backend-go/helper"
	"github.com/Kelompok14-LMS/backend-go/pkg"
	"github.com/labstack/echo/v4"
)

type DetailAssignmentController struct {
	detailAssignmentUsecase detailAssignment.Usecase
}

func NewDetailAssignmentController(detailAssignmentUsecase detailAssignment.Usecase) *DetailAssignmentController {
	return &DetailAssignmentController{
		detailAssignmentUsecase: detailAssignmentUsecase,
	}
}

func (ctrl *DetailAssignmentController) HandlerDetailAssignment(c echo.Context) error {
	assignmentid := c.Param("assignmentid")

	assignment, err := ctrl.detailAssignmentUsecase.DetailAssignment(assignmentid)

	if err != nil {
		if errors.Is(err, pkg.ErrCourseNotFound) {
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		} else if errors.Is(err, pkg.ErrModuleNotFound) {
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		} else if errors.Is(err, pkg.ErrMaterialNotFound) {
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Success get full detail assignment", response.FullDetailAssignement(assignment)))
}
