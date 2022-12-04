package modules_test

import (
	"testing"
	"time"

	"github.com/Kelompok14-LMS/backend-go/businesses/courses"
	_courseMock "github.com/Kelompok14-LMS/backend-go/businesses/courses/mocks"
	"github.com/Kelompok14-LMS/backend-go/businesses/modules"
	_moduleMock "github.com/Kelompok14-LMS/backend-go/businesses/modules/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var (
	moduleRepository _moduleMock.Repository
	moduleService    modules.Usecase

	courseRepository _courseMock.Repository
)

func TestMain(m *testing.M) {
	moduleRepository = _moduleMock.Repository{Mock: mock.Mock{}}
	courseRepository = _courseMock.Repository{Mock: mock.Mock{}}

	moduleService = modules.NewModuleUsecase(&moduleRepository, &courseRepository)

	m.Run()
}

func TestCreate(t *testing.T) {
	t.Run("Create | Success", func(t *testing.T) {
		mockCourse := courses.Domain{
			ID:          "COURSE_1",
			MentorId:    "MENTOR_1",
			CategoryId:  "CAT_1",
			Title:       "Course Title",
			Description: "Course description",
			Thumbnail:   "https://imgurl.com/bucket/object",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			DeletedAt:   gorm.DeletedAt{},
		}

		courseRepository.Mock.On("FindById", "COURSE_1").Return(&mockCourse, nil)

		moduleDomain := modules.Domain{
			ID:       "MOD_1",
			CourseId: "COURSE_1",
			Title:    "Course Title",
		}

		moduleRepository.Mock.On("Create", mock.Anything).Return(nil)

		err := moduleService.Create(&moduleDomain)

		assert.NoError(t, err)
	})
}

func TestFindById(t *testing.T) {
	t.Run("FindById | Success", func(t *testing.T) {
		moduleDomain := modules.Domain{
			ID:        "MOD_1",
			CourseId:  "COURSE_1",
			Title:     "Course Title",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: gorm.DeletedAt{},
		}

		moduleRepository.Mock.On("FindById", "MOD_1").Return(&moduleDomain, nil)

		result, err := moduleService.FindById("MOD_1")

		assert.Nil(t, err)
		assert.NotNil(t, result)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update | Success", func(t *testing.T) {
		mockCourse := courses.Domain{
			ID:          "COURSE_1",
			MentorId:    "MENTOR_1",
			CategoryId:  "CAT_1",
			Title:       "Course Title",
			Description: "Course description",
			Thumbnail:   "https://imgurl.com/bucket/object",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			DeletedAt:   gorm.DeletedAt{},
		}

		courseRepository.Mock.On("FindById", "COURSE_1").Return(&mockCourse, nil)

		moduleDomain := modules.Domain{
			CourseId: "COURSE_1",
			Title:    "Course Title Updated",
		}

		moduleRepository.Mock.On("Update", "MOD_1", &moduleDomain).Return(nil)

		err := moduleService.Update("MOD_1", &moduleDomain)

		assert.NoError(t, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete | Success", func(t *testing.T) {
		moduleDomain := modules.Domain{
			ID:        "MOD_1",
			CourseId:  "COURSE_1",
			Title:     "Course Title",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: gorm.DeletedAt{},
		}

		moduleRepository.Mock.On("FindById", "MOD_1").Return(&moduleDomain, nil)

		moduleRepository.Mock.On("Delete", "MOD_1").Return(nil)

		err := moduleService.Delete("MOD_1")

		assert.NoError(t, err)
	})
}
