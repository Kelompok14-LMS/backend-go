package assignments_test

import (
	"testing"
	"time"

	"github.com/Kelompok14-LMS/backend-go/businesses/assignments"
	_assignmentMock "github.com/Kelompok14-LMS/backend-go/businesses/assignments/mocks"
	"github.com/Kelompok14-LMS/backend-go/businesses/modules"
	_moduleMock "github.com/Kelompok14-LMS/backend-go/businesses/modules/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var (
	assignmentRepository _assignmentMock.Repository
	assignmentService    assignments.Usecase

	moduleRepository _moduleMock.Repository
	// storageClient    helper.StorageConfig
)

func TestMain(m *testing.M) {
	assignmentRepository = _assignmentMock.Repository{Mock: mock.Mock{}}
	moduleRepository = _moduleMock.Repository{Mock: mock.Mock{}}

	assignmentService = assignments.NewAssignmentUsecase(&assignmentRepository, &moduleRepository)

	m.Run()
}

func TestFindById(t *testing.T) {
	mockAssignment := assignments.Domain{
		ID:          "Assignment_1",
		ModuleID:    "Module_1",
		Title:       "Title test",
		Description: "Description test",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   gorm.DeletedAt{},
	}

	assignmentRepository.Mock.On("FindById", "Assignment_1").Return(&mockAssignment, nil)

	result, err := assignmentService.FindById("Assignment_1")

	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestUpdate(t *testing.T) {
	moduleDomain := modules.Domain{
		ID:        "Module_1",
		CourseId:  "Course_1",
		Title:     "Course Title",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: gorm.DeletedAt{},
	}

	moduleRepository.Mock.On("FindById", "MODULE_1").Return(&moduleDomain, nil)

	mockAssignment := assignments.Domain{
		ID:          "Assignment_1",
		ModuleID:    "Module_1",
		Title:       "Title test",
		Description: "Description test",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   gorm.DeletedAt{},
	}

	assignmentRepository.Mock.On("FindById", "Assignment_1").Return(&mockAssignment, nil)

	assignmentDomain := assignments.Domain{
		ModuleID:    "MODULE_1",
		Title:       "Title test",
		Description: "Description test",
	}

	assignmentRepository.Mock.On("Update", "Assignment_1", &assignmentDomain).Return(nil)

	updatedAssignment := assignments.Domain{
		ID:          "Assignment_1",
		ModuleID:    "MODULE_1",
		Title:       "Title test",
		Description: "Description test",
	}

	err := assignmentService.Update("Assignment_1", &updatedAssignment)

	assert.Nil(t, err)
}

func TestDelete(t *testing.T) {
	mockAssignment := assignments.Domain{
		ID:          "Assignment_1",
		ModuleID:    "MODULE_1",
		Title:       "Title test",
		Description: "Description test",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   gorm.DeletedAt{},
	}

	assignmentRepository.Mock.On("FindById", "Assignment_1").Return(&mockAssignment, nil)

	assignmentRepository.Mock.On("Delete", "Assignment_1").Return(nil)

	err := assignmentService.Delete("Assignment_1")

	assert.Nil(t, err)
}
