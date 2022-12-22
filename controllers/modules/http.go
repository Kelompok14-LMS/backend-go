package modules

import (
	"net/http"

	"github.com/Kelompok14-LMS/backend-go/businesses/modules"
	"github.com/Kelompok14-LMS/backend-go/controllers/modules/request"
	"github.com/Kelompok14-LMS/backend-go/controllers/modules/response"
	"github.com/Kelompok14-LMS/backend-go/helper"
	"github.com/Kelompok14-LMS/backend-go/pkg"
	"github.com/labstack/echo/v4"
)

type ModuleController struct {
	moduleUsecase modules.Usecase
}

func NewModuleController(moduleUsecase modules.Usecase) *ModuleController {
	return &ModuleController{
		moduleUsecase: moduleUsecase,
	}
}

func (ctrl *ModuleController) HandlerCreateModule(c echo.Context) error {
	moduleInput := request.CreateModuleInput{}

	if err := c.Bind(&moduleInput); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrInvalidRequest.Error()))
	}

	if err := moduleInput.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrInvalidRequest.Error()))
	}

	err := ctrl.moduleUsecase.Create(moduleInput.ToDomain())

	if err != nil {
		switch err {
		case pkg.ErrCourseNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrCourseNotFound.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(pkg.ErrInternalServerError.Error()))
		}
	}

	return c.JSON(http.StatusCreated, helper.SuccessCreatedResponse("Sukses menambahkan modul", nil))
}

func (ctrl *ModuleController) HandlerFindByIdModule(c echo.Context) error {
	moduleId := c.Param("moduleId")

	module, err := ctrl.moduleUsecase.FindById(moduleId)

	if err != nil {
		switch err {
		case pkg.ErrModuleNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrModuleNotFound.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(pkg.ErrInternalServerError.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Sukses get modul berdasarkan id", response.DetailModule(module)))
}

func (ctrl *ModuleController) HandlerUpdateModule(c echo.Context) error {
	moduleId := c.Param("moduleId")
	moduleInput := request.UpdateModuleInput{}

	if err := c.Bind(&moduleInput); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrInvalidRequest.Error()))
	}

	if err := moduleInput.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(err.Error()))
	}

	err := ctrl.moduleUsecase.Update(moduleId, moduleInput.ToDomain())

	if err != nil {
		switch err {
		case pkg.ErrCourseNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrCourseNotFound.Error()))
		case pkg.ErrModuleNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrModuleNotFound.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(pkg.ErrInternalServerError.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Sukses update modul", nil))
}

func (ctrl *ModuleController) HandlerDeleteModule(c echo.Context) error {
	moduleId := c.Param("moduleId")

	err := ctrl.moduleUsecase.Delete(moduleId)

	if err != nil {
		switch err {
		case pkg.ErrModuleNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrModuleNotFound.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(pkg.ErrInternalServerError.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Modul dihapus", nil))
}
