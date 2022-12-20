package reviews

import (
	"errors"
	"net/http"

	"github.com/Kelompok14-LMS/backend-go/businesses/reviews"
	"github.com/Kelompok14-LMS/backend-go/controllers/reviews/request"
	"github.com/Kelompok14-LMS/backend-go/controllers/reviews/response"
	"github.com/Kelompok14-LMS/backend-go/helper"
	"github.com/Kelompok14-LMS/backend-go/pkg"
	"github.com/labstack/echo/v4"
)

type ReviewController struct {
	reviewUsecase reviews.Usecase
}

func NewReviewController(reviewUscase reviews.Usecase) *ReviewController {
	return &ReviewController{
		reviewUsecase: reviewUscase,
	}
}

func (ctrl *ReviewController) HandlerCreateReview(c echo.Context) error {
	reviewInput := request.CreateReviewInput{}

	if err := c.Bind(&reviewInput); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrInvalidRequest.Error()))
	}

	if err := reviewInput.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(err.Error()))
	}

	err := ctrl.reviewUsecase.Create(reviewInput.ToDomain())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(pkg.ErrInternalServerError.Error()))
	}

	return c.JSON(http.StatusCreated, helper.SuccessCreatedResponse("Sukses menambahkan ulasan", nil))
}

func (ctrl *ReviewController) HandlerFindByMentee(c echo.Context) error {
	keyword := c.QueryParam("keyword")
	menteeId := c.Param("menteeId")

	reviews, err := ctrl.reviewUsecase.FindByMentee(menteeId, keyword)

	if err != nil {
		if errors.Is(err, pkg.ErrMenteeNotFound) {
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(pkg.ErrInternalServerError.Error()))
		}
	}

	var reviewResponse []response.FindReviewByMentee

	for _, review := range reviews {
		reviewResponse = append(reviewResponse, response.ReviewsByMentee(&review))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Sukeses get ulasan mentee", reviewResponse))
}

func (ctrl *ReviewController) HandlerFindByCourse(c echo.Context) error {
	courseId := c.Param("courseId")

	reviews, err := ctrl.reviewUsecase.FindByCourse(courseId)

	if err != nil {
		if errors.Is(err, pkg.ErrCourseNotFound) {
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(pkg.ErrInternalServerError.Error()))
		}
	}

	var reviewResponse []response.FindReviewByCourse

	for _, review := range reviews {
		reviewResponse = append(reviewResponse, response.ReviewsByCourse(&review))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Sukses get ulasan kursus", reviewResponse))
}
