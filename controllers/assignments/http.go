package assignments

import (
	"errors"
	"net/http"

	"github.com/Kelompok14-LMS/backend-go/businesses/assignments"
	"github.com/Kelompok14-LMS/backend-go/controllers/assignments/request"
	"github.com/Kelompok14-LMS/backend-go/controllers/assignments/response"

	"github.com/Kelompok14-LMS/backend-go/helper"
	"github.com/Kelompok14-LMS/backend-go/pkg"
	"github.com/labstack/echo/v4"
)

type AssignmentController struct {
	assignmentUsecase assignments.Usecase
}

func NewAssignmentsController(assignmentUsecase assignments.Usecase) *AssignmentController {
	return &AssignmentController{
		assignmentUsecase: assignmentUsecase,
	}
}

func (ctrl *AssignmentController) HandlerCreateAssignment(c echo.Context) error {
	assignmentInput := request.CreateAssignment{}

	if err := c.Bind(&assignmentInput); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrInvalidRequest.Error()))
	}

	if err := assignmentInput.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrInvalidRequest.Error()))
	}

	err := ctrl.assignmentUsecase.Create(assignmentInput.ToDomain())

	if err != nil {
		if errors.Is(err, pkg.ErrAssignmentNotFound) {
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrAssignmentNotFound.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	return c.JSON(http.StatusCreated, helper.SuccessCreatedResponse("Success create assignments", nil))
}

func (ctrl *AssignmentController) HandlerFindByIdAssignment(c echo.Context) error {
	assignmentId := c.Param("assignmentId")

	assignment, err := ctrl.assignmentUsecase.FindById(assignmentId)

	if err != nil {
		if errors.Is(err, pkg.ErrAssignmentNotFound) {
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrAssignmentNotFound.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Success get assignment by id", response.DetailAssignment(assignment)))
}

func (ctrl *AssignmentController) HandlerUpdateAssignment(c echo.Context) error {
	assignmentId := c.Param("assignmentId")
	assignmentInput := request.UpdateAssignment{}

	if err := c.Bind(&assignmentInput); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrInvalidRequest.Error()))
	}

	if err := assignmentInput.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(err.Error()))
	}

	err := ctrl.assignmentUsecase.Update(assignmentId, assignmentInput.ToDomain())

	if err != nil {
		if errors.Is(err, pkg.ErrModuleNotFound) {
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrModuleNotFound.Error()))
		} else if errors.Is(err, pkg.ErrAssignmentNotFound) {
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrAssignmentNotFound.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Success update assignments", nil))
}

func (ctrl *AssignmentController) HandlerDeleteAssignment(c echo.Context) error {
	assignmentId := c.Param("assignmentId")

	err := ctrl.assignmentUsecase.Delete(assignmentId)

	if err != nil {
		if errors.Is(err, pkg.ErrAssignmentNotFound) {
			c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrAssignmentNotFound.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Assignment deleted", nil))
}
