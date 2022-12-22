package materials

import (

	"context"
	"net/http"

	"github.com/Kelompok14-LMS/backend-go/businesses/materials"
	"github.com/Kelompok14-LMS/backend-go/controllers/materials/request"
	"github.com/Kelompok14-LMS/backend-go/controllers/materials/response"
	"github.com/Kelompok14-LMS/backend-go/helper"
	"github.com/Kelompok14-LMS/backend-go/pkg"
	"github.com/labstack/echo/v4"
)

type MaterialController struct {
	materialUsecase materials.Usecase
}

func NewMaterialController(materialUsecase materials.Usecase) *MaterialController {
	return &MaterialController{
		materialUsecase: materialUsecase,
	}
}

func (ctrl *MaterialController) HandlerCreateMaterial(c echo.Context) error {
	materialInput := request.CreateMaterialInput{}

	materialInput.File, _ = c.FormFile("video")

	if err := c.Bind(&materialInput); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrInvalidRequest.Error()))
	}

	if err := materialInput.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(err.Error()))
	}

	ctx := context.Background()

	err := ctrl.materialUsecase.Create(ctx, materialInput.ToDomain())

	if err != nil {
		switch err {
		case pkg.ErrModuleNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrModuleNotFound.Error()))
		case pkg.ErrUnsupportedVideoFile:
			return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrUnsupportedVideoFile.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(pkg.ErrInternalServerError.Error()))
		}

	}

	return c.JSON(http.StatusCreated, helper.SuccessCreatedResponse("Sukses menambahkan materi", nil))
}

func (ctrl *MaterialController) HandlerFindByIdMaterial(c echo.Context) error {
	materialId := c.Param("materialId")

	material, err := ctrl.materialUsecase.FindById(materialId)

	if err != nil {
		switch err {
		case pkg.ErrMaterialAssetNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrMaterialAssetNotFound.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(pkg.ErrInternalServerError.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Sukses get materi berdasarkan id", response.MaterialDetail(material)))
}

func (ctrl *MaterialController) HandlerUpdateMaterial(c echo.Context) error {
	materialId := c.Param("materialId")
	materialInput := request.UpdateMaterialInput{}

	file, _ := c.FormFile("video")

	if file != nil {
		materialInput.File = file

		if err := c.Bind(&materialInput); err != nil {
			return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrInvalidRequest.Error()))
		}
	} else {
		materialInput.ModuleId = c.FormValue("module_id")
		materialInput.Title = c.FormValue("title")
		materialInput.Description = c.FormValue("description")
		materialInput.File = nil
	}

	if err := materialInput.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(err.Error()))
	}

	ctx := context.Background()

	err := ctrl.materialUsecase.Update(ctx, materialId, materialInput.ToDomain())

	if err != nil {
		switch err {
		case pkg.ErrModuleNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrModuleNotFound.Error()))
		case pkg.ErrMaterialAssetNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrMaterialAssetNotFound.Error()))
		case pkg.ErrUnsupportedVideoFile:
			return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrUnsupportedVideoFile.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(pkg.ErrInternalServerError.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Sukses update materi", nil))
}

func (ctrl *MaterialController) HandlerSoftDeleteMaterial(c echo.Context) error {
	materialId := c.Param("materialId")

	err := ctrl.materialUsecase.Delete(materialId)

	if err != nil {
		switch err {
		case pkg.ErrMaterialAssetNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrMaterialAssetNotFound.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(pkg.ErrInternalServerError.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Materi dihapus", nil))
}

func (ctrl *MaterialController) HandlerSoftDeleteMaterialByModule(c echo.Context) error {
	moduleId := c.Param("moduleId")

	err := ctrl.materialUsecase.Deletes(moduleId)

	if err != nil {
		switch err {
		case pkg.ErrModuleNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrModuleNotFound.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(pkg.ErrInternalServerError.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Materi dihapus", nil))
}
