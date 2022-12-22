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
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(err.Error()))
	}

	err := ctrl.assignmentUsecase.Create(assignmentInput.ToDomain())

	if err != nil {
		switch {
		case errors.Is(err, pkg.ErrAssignmentAlredyExist):
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrAssignmentAlredyExist.Error()))
		case errors.Is(err, pkg.ErrCourseNotFound):
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrCourseNotFound.Error()))
		case errors.Is(err, pkg.ErrAssignmentNotFound):
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrAssignmentNotFound.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(pkg.ErrInternalServerError.Error()))
		}
	}

	return c.JSON(http.StatusCreated, helper.SuccessCreatedResponse("Sukses menambahkan tugas", nil))
}

func (ctrl *AssignmentController) HandlerFindByIdAssignment(c echo.Context) error {
	assignmentId := c.Param("assignmentId")

	assignment, err := ctrl.assignmentUsecase.FindById(assignmentId)

	if err != nil {
		switch {
		case errors.Is(err, pkg.ErrAssignmentNotFound):
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrAssignmentNotFound.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(pkg.ErrInternalServerError.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Sukses get tugas berdasarkan id", response.DetailAssignment(assignment)))
}

func (ctrl *AssignmentController) HandlerFindByCourse(c echo.Context) error {
	courseid := c.Param("courseid")

	assignmentCourse, err := ctrl.assignmentUsecase.FindByCourseId(courseid)

	if err != nil {
		switch {
		case errors.Is(err, pkg.ErrCourseNotFound):
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrCourseNotFound.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(pkg.ErrInternalServerError.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Sukses get tugas berdasarkan id kursus", *response.DetailAssignment(assignmentCourse)))
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
		switch {
		case errors.Is(err, pkg.ErrCourseNotFound):
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrCourseNotFound.Error()))
		case errors.Is(err, pkg.ErrAssignmentNotFound):
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrAssignmentNotFound.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(pkg.ErrInternalServerError.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Sukses update tugas", nil))
}

func (ctrl *AssignmentController) HandlerDeleteAssignment(c echo.Context) error {
	assignmentId := c.Param("assignmentId")

	err := ctrl.assignmentUsecase.Delete(assignmentId)

	if err != nil {
		switch {
		case errors.Is(err, pkg.ErrAssignmentNotFound):
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrAssignmentNotFound.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(pkg.ErrInternalServerError.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Tugas dihapus", nil))
}
