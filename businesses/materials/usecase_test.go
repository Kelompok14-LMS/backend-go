package materials_test

import (
	"testing"
	"time"

	"github.com/Kelompok14-LMS/backend-go/businesses/materials"
	_materialMock "github.com/Kelompok14-LMS/backend-go/businesses/materials/mocks"
	"github.com/Kelompok14-LMS/backend-go/businesses/modules"
	_moduleMock "github.com/Kelompok14-LMS/backend-go/businesses/modules/mocks"
	"github.com/Kelompok14-LMS/backend-go/helper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var (
	materialRepository _materialMock.Repository
	materialService    materials.Usecase

	moduleRepository _moduleMock.Repository
	storageClient    helper.StorageConfig
)

func TestMain(m *testing.M) {
	materialRepository = _materialMock.Repository{Mock: mock.Mock{}}
	moduleRepository = _moduleMock.Repository{Mock: mock.Mock{}}
	storageClient = helper.StorageConfig{}

	materialService = materials.NewMaterialUsecase(&materialRepository, &moduleRepository, &storageClient)

	m.Run()
}

func TestFindById(t *testing.T) {
	mockMaterial := materials.Domain{
		ID:          "MATERIAL_1",
		ModuleId:    "MODULE_1",
		Title:       "Title test",
		URL:         "https://storage.com/to/bucket/object.mp4",
		Description: "Description test",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   gorm.DeletedAt{},
	}

	materialRepository.Mock.On("FindById", "MATERIAL_1").Return(&mockMaterial, nil)

	result, err := materialService.FindById("MATERIAL_1")

	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestUpdate(t *testing.T) {
	moduleDomain := modules.Domain{
		ID:        "MODULE_1",
		CourseId:  "COURSE_1",
		Title:     "Course Title",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: gorm.DeletedAt{},
	}

	moduleRepository.Mock.On("FindById", "MODULE_1").Return(&moduleDomain, nil)

	mockMaterial := materials.Domain{
		ID:          "MATERIAL_1",
		ModuleId:    "MODULE_1",
		Title:       "Title test",
		URL:         "https://storage.com/to/bucket/object.mp4",
		Description: "Description test",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   gorm.DeletedAt{},
	}

	materialRepository.Mock.On("FindById", "MATERIAL_1").Return(&mockMaterial, nil)

	materialDomain := materials.Domain{
		ModuleId:    "MODULE_1",
		Title:       "Title test",
		URL:         "",
		Description: "Description test",
	}

	materialRepository.Mock.On("Update", "MATERIAL_1", &materialDomain).Return(nil)

	updatedMaterial := materials.Domain{
		ID:          "MATERIAL_1",
		ModuleId:    "MODULE_1",
		Title:       "Title test",
		File:        nil,
		Description: "Description test",
	}

	err := materialService.Update("MATERIAL_1", &updatedMaterial)

	assert.Nil(t, err)
}

func TestDelete(t *testing.T) {
	mockMaterial := materials.Domain{
		ID:          "MATERIAL_1",
		ModuleId:    "MODULE_1",
		Title:       "Title test",
		URL:         "https://storage.com/to/bucket/object.mp4",
		Description: "Description test",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   gorm.DeletedAt{},
	}

	materialRepository.Mock.On("FindById", "MATERIAL_1").Return(&mockMaterial, nil)

	materialRepository.Mock.On("Delete", "MATERIAL_1").Return(nil)

	err := materialService.Delete("MATERIAL_1")

	assert.Nil(t, err)
}

func TestDeletes(t *testing.T) {
	moduleDomain := modules.Domain{
		ID:        "MODULE_1",
		CourseId:  "COURSE_1",
		Title:     "Course Title",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: gorm.DeletedAt{},
	}

	moduleRepository.Mock.On("FindById", "MODULE_1").Return(&moduleDomain, nil)

	materialRepository.Mock.On("Deletes", "MODULE_1").Return(nil)

	err := materialService.Deletes("MODULE_1")

	assert.Nil(t, err)
}
