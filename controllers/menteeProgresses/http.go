package mentee_progresses

import (
	"errors"
	"net/http"

	menteeProgresses "github.com/Kelompok14-LMS/backend-go/businesses/menteeProgresses"
	"github.com/Kelompok14-LMS/backend-go/controllers/menteeProgresses/request"
	"github.com/Kelompok14-LMS/backend-go/controllers/menteeProgresses/response"
	"github.com/Kelompok14-LMS/backend-go/helper"
	"github.com/Kelompok14-LMS/backend-go/pkg"
	"github.com/labstack/echo/v4"
)

type MenteeProgressController struct {
	menteeProgressUsecase menteeProgresses.Usecase
}

func NewMenteeProgressController(menteeProgressUsecase menteeProgresses.Usecase) *MenteeProgressController {
	return &MenteeProgressController{
		menteeProgressUsecase: menteeProgressUsecase,
	}
}

func (ctrl *MenteeProgressController) HandlerAddProgress(c echo.Context) error {
	menteeProgressInput := request.AddProgressInput{}

	if err := c.Bind(&menteeProgressInput); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrInvalidRequest.Error()))
	}

	if err := menteeProgressInput.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(err.Error()))
	}

	err := ctrl.menteeProgressUsecase.Add(menteeProgressInput.ToDomain())

	if err != nil {
		if errors.Is(err, pkg.ErrMenteeNotFound) {
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrMenteeNotFound.Error()))
		} else if errors.Is(err, pkg.ErrCourseNotFound) {
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrCourseNotFound.Error()))
		} else if errors.Is(err, pkg.ErrMaterialAssetNotFound) {
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrMaterialAssetNotFound.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(pkg.ErrInternalServerError.Error()))
		}
	}

	return c.JSON(http.StatusCreated, helper.SuccessCreatedResponse("Sukses menambahkan progres", nil))
}

func (ctrl *MenteeProgressController) HandlerFindMaterialEnrolled(c echo.Context) error {
	menteeId := c.Param("menteeId")
	materialId := c.Param("materialId")

	progress, err := ctrl.menteeProgressUsecase.FindMaterialEnrolled(menteeId, materialId)

	if err != nil {
		if errors.Is(err, pkg.ErrMenteeNotFound) {
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		} else if errors.Is(err, pkg.ErrMaterialNotFound) {
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(pkg.ErrInternalServerError.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Sukses get materi", response.DetailMaterialEnrolled(progress)))
}
