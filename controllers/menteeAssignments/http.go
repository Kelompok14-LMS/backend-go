package mentee_assignments

import (
	"errors"
	"fmt"
	"net/http"

	menteeAssignments "github.com/Kelompok14-LMS/backend-go/businesses/menteeAssignments"
	"github.com/Kelompok14-LMS/backend-go/controllers/menteeAssignments/request"
	"github.com/Kelompok14-LMS/backend-go/controllers/menteeAssignments/response"
	"github.com/Kelompok14-LMS/backend-go/helper"
	"github.com/Kelompok14-LMS/backend-go/pkg"
	"github.com/Kelompok14-LMS/backend-go/utils"
	"github.com/labstack/echo/v4"
)

type AssignmentMenteeController struct {
	assignmentMenteeUsecase menteeAssignments.Usecase
	jwtConfig               *utils.JWTConfig
}

func NewAssignmentsMenteeController(assignmentMenteeUsecase menteeAssignments.Usecase, jwtConfig *utils.JWTConfig) *AssignmentMenteeController {
	return &AssignmentMenteeController{
		assignmentMenteeUsecase: assignmentMenteeUsecase,
		jwtConfig:               jwtConfig,
	}
}

func (ctrl *AssignmentMenteeController) HandlerCreateMenteeAssignment(c echo.Context) error {
	assignmentMenteeInput := request.CreateMenteeAssignment{}

	assignmentMenteeInput.PDF, _ = c.FormFile("pdf")

	if err := c.Bind(&assignmentMenteeInput); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrInvalidRequest.Error()))
	}

	if err := assignmentMenteeInput.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrInvalidRequest.Error()))
	}

	err := ctrl.assignmentMenteeUsecase.Create(assignmentMenteeInput.ToDomain())

	if err != nil {
		if errors.Is(err, pkg.ErrAssignmentNotFound) {
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrAssignmentNotFound.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	return c.JSON(http.StatusCreated, helper.SuccessCreatedResponse("Success create mentee assignments ", nil))
}

func (ctrl *AssignmentMenteeController) HandlerUpdateMenteeAssignment(c echo.Context) error {
	assignmentMenteeId := c.Param("menteeAssignmentId")
	menteeAssignmentInput := request.UpdateMenteeAssignment{}

	menteeAssignmentInput.PDF, _ = c.FormFile("pdf")

	menteeAssignmentInput.MenteeID = c.FormValue("mentee_id")
	menteeAssignmentInput.AssignmentID = c.FormValue("assignment_id")

	fmt.Println(menteeAssignmentInput)

	if err := c.Bind(&menteeAssignmentInput); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrInvalidRequest.Error()))
	}

	if err := menteeAssignmentInput.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(err.Error()))
	}

	err := ctrl.assignmentMenteeUsecase.Update(assignmentMenteeId, menteeAssignmentInput.ToDomain())

	if err != nil {
		if errors.Is(err, pkg.ErrAssignmentMenteeNotFound) {
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrAssignmentMenteeNotFound.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Success update mentee assignments", nil))
}

func (ctrl *AssignmentMenteeController) HandlerUpdateGradeMentee(c echo.Context) error {
	id := c.Param("menteeAssignmentId")

	menteeAssignmentInput := request.CreateGrade{}

	if err := c.Bind(&menteeAssignmentInput); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrInvalidRequest.Error()))
	}

	if err := menteeAssignmentInput.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(err.Error()))
	}

	err := ctrl.assignmentMenteeUsecase.Update(id, menteeAssignmentInput.ToDomain())

	if err != nil {
		if errors.Is(err, pkg.ErrAssignmentMenteeNotFound) {
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrAssignmentMenteeNotFound.Error()))
		}

		return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Success update grade", nil))
}

func (ctrl *AssignmentMenteeController) HandlerFindByIdMenteeAssignment(c echo.Context) error {
	id := c.Param("menteeAssignmentId")

	assignmentMentee, err := ctrl.assignmentMenteeUsecase.FindById(id)

	if err != nil {
		if errors.Is(err, pkg.ErrAssignmentMenteeNotFound) {
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrAssignmentMenteeNotFound.Error()))
		}

		return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Success get assignment mentee by id", response.FromDomain(assignmentMentee)))
}

func (ctrl *AssignmentMenteeController) HandlerFindByAssignmentId(c echo.Context) error {
	id := c.Param("assignmentId")

	assignmentMentee, err := ctrl.assignmentMenteeUsecase.FindByAssignmentId(id)

	if err != nil {
		if errors.Is(err, pkg.ErrAssignmentMenteeNotFound) {
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrAssignmentMenteeNotFound.Error()))
		}

		return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
	}

	var menteeAssignmentResponse []response.AssignmentMentee

	for _, mentee_assignments := range assignmentMentee {
		menteeAssignmentResponse = append(menteeAssignmentResponse, response.FromDomain(&mentee_assignments))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Success get assignment mentee by assignment id", menteeAssignmentResponse))
}

func (ctrl *AssignmentMenteeController) HandlerFindByMenteeId(c echo.Context) error {
	token, _ := ctrl.jwtConfig.ExtractToken(c)

	menteeAssignment, err := ctrl.assignmentMenteeUsecase.FindByMenteeId(token.MenteeId)

	if err != nil {
		if errors.Is(err, pkg.ErrAssignmentMenteeNotFound) {
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrAssignmentMenteeNotFound.Error()))
		} else if errors.Is(err, pkg.ErrUserNotFound) {
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrUserNotFound.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}
	var menteeAssignmentResponse []response.AssignmentMentee

	for _, mentee_assignments := range menteeAssignment {
		menteeAssignmentResponse = append(menteeAssignmentResponse, response.FromDomain(&mentee_assignments))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Success get assignment mentee by mentee id ", menteeAssignmentResponse))
}

func (ctrl *AssignmentMenteeController) HandlerSoftDeleteMenteeAssignment(c echo.Context) error {
	id := c.Param("menteeAssignmentId")

	err := ctrl.assignmentMenteeUsecase.Delete(id)

	if err != nil {
		if errors.Is(err, pkg.ErrAssignmentMenteeNotFound) {
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrAssignmentMenteeNotFound.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("assignment mentee deleted", nil))
}
