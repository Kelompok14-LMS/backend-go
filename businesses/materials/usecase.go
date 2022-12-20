package materials

import (
	"context"
	"path/filepath"

	"github.com/Kelompok14-LMS/backend-go/businesses/modules"
	"github.com/Kelompok14-LMS/backend-go/helper"
	"github.com/Kelompok14-LMS/backend-go/pkg"
	"github.com/Kelompok14-LMS/backend-go/utils"
	"github.com/google/uuid"
)

type materialUsecase struct {
	materialRepository Repository
	moduleRepository   modules.Repository
	storage            *helper.StorageConfig
}

func NewMaterialUsecase(
	materialRepository Repository,
	moduleRepository modules.Repository,
	storage *helper.StorageConfig,
) Usecase {
	return materialUsecase{
		materialRepository: materialRepository,
		moduleRepository:   moduleRepository,
		storage:            storage,
	}
}

func (mu materialUsecase) Create(materialDomain *Domain) error {
	if _, err := mu.moduleRepository.FindById(materialDomain.ModuleId); err != nil {
		return err
	}

	file, err := materialDomain.File.Open()

	if err != nil {
		return err
	}

	defer file.Close()

	extension := filepath.Ext(materialDomain.File.Filename)

	if extension != ".mp4" && extension != ".mkv" {
		return pkg.ErrUnsupportedVideoFile
	}

	filename, err := utils.GetFilename(materialDomain.File.Filename)

	if err != nil {
		return pkg.ErrUnsupportedVideoFile
	}

	ctx := context.Background()

	url, err := mu.storage.UploadVideo(ctx, filename, file)

	if err != nil {
		return err
	}

	material := Domain{
		ID:          uuid.NewString(),
		ModuleId:    materialDomain.ModuleId,
		Title:       materialDomain.Title,
		URL:         url,
		Description: materialDomain.Description,
	}

	if err := mu.materialRepository.Create(&material); err != nil {
		return err
	}

	return nil
}

func (mu materialUsecase) FindById(materialId string) (*Domain, error) {
	material, err := mu.materialRepository.FindById(materialId)

	if err != nil {
		return nil, err
	}

	return material, nil
}

func (mu materialUsecase) Update(materialId string, materialDomain *Domain) error {
	if _, err := mu.moduleRepository.FindById(materialDomain.ModuleId); err != nil {
		return err
	}

	material, err := mu.materialRepository.FindById(materialId)

	if err != nil {
		return err
	}

	var url string

	if materialDomain.File != nil {
		ctx := context.Background()

		err := mu.storage.DeleteObject(ctx, material.URL)

		if err != nil {
			return err
		}

		file, err := materialDomain.File.Open()

		if err != nil {
			return err
		}

		defer file.Close()

		extension := filepath.Ext(materialDomain.File.Filename)

		if extension != ".mp4" && extension != ".mkv" {
			return pkg.ErrUnsupportedVideoFile
		}

		filename, err := utils.GetFilename(materialDomain.File.Filename)

		if err != nil {
			return pkg.ErrUnsupportedVideoFile
		}

		url, err = mu.storage.UploadVideo(ctx, filename, file)

		if err != nil {
			return err
		}
	}

	updatedMaterial := Domain{
		ModuleId:    materialDomain.ModuleId,
		Title:       materialDomain.Title,
		URL:         url,
		Description: materialDomain.Description,
	}

	if err := mu.materialRepository.Update(materialId, &updatedMaterial); err != nil {
		return err
	}

	return nil
}

func (mu materialUsecase) Delete(materialId string) error {
	if _, err := mu.materialRepository.FindById(materialId); err != nil {
		return err
	}

	err := mu.materialRepository.Delete(materialId)

	if err != nil {
		return err
	}

	return nil
}

func (mu materialUsecase) Deletes(moduleId string) error {
	if _, err := mu.moduleRepository.FindById(moduleId); err != nil {
		return err
	}

	err := mu.materialRepository.Deletes(moduleId)

	if err != nil {
		return err
	}

	return nil
}
