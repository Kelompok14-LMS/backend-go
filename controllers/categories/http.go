package categories

import (
	"errors"
	"net/http"

	"github.com/Kelompok14-LMS/backend-go/businesses/categories"
	"github.com/Kelompok14-LMS/backend-go/controllers/categories/request"
	"github.com/Kelompok14-LMS/backend-go/controllers/categories/response"
	"github.com/Kelompok14-LMS/backend-go/helper"
	"github.com/Kelompok14-LMS/backend-go/pkg"
	"github.com/labstack/echo/v4"
)

type CategoryController struct {
	categoryUsecase categories.Usecase
}

func NewCategoryController(categoryUsecase categories.Usecase) *CategoryController {
	return &CategoryController{
		categoryUsecase: categoryUsecase,
	}
}

func (ctrl *CategoryController) HandlerCreateCategory(c echo.Context) error {
	categoryInput := request.Category{}

	if err := c.Bind(&categoryInput); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrInvalidRequest.Error()))
	}

	if err := categoryInput.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(err.Error()))
	}

	err := ctrl.categoryUsecase.Create(categoryInput.ToDomain())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusCreated, helper.SuccessCreatedResponse("Success add category", nil))
}

func (ctrl *CategoryController) HandlerFindAllCategories(c echo.Context) error {
	categoriesDomain, err := ctrl.categoryUsecase.FindAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
	}

	var categoriesResponse []response.Category

	for _, category := range *categoriesDomain {
		categoriesResponse = append(categoriesResponse, response.FromDomain(&category))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Success get all categories", categoriesResponse))
}

func (ctrl *CategoryController) HandlerFindByIdCategory(c echo.Context) error {
	id := c.Param("categoryId")

	category, err := ctrl.categoryUsecase.FindById(id)

	if err != nil {
		if errors.Is(err, pkg.ErrCategoryNotFound) {
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrCategoryNotFound.Error()))
		}

		return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Success get category by id", response.FromDomain(category)))
}

func (ctrl *CategoryController) HandlerUpdateCategory(c echo.Context) error {
	id := c.Param("categoryId")

	categoryInput := request.Category{}

	if err := c.Bind(&categoryInput); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrInvalidRequest.Error()))
	}

	if err := categoryInput.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(err.Error()))
	}

	err := ctrl.categoryUsecase.Update(id, categoryInput.ToDomain())

	if err != nil {
		if errors.Is(err, pkg.ErrCategoryNotFound) {
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrCategoryNotFound.Error()))
		}

		return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Success update category", nil))
}
