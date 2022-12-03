package courses_test

import (
	"testing"
	"time"

	"github.com/Kelompok14-LMS/backend-go/businesses/categories"
	_categoryMock "github.com/Kelompok14-LMS/backend-go/businesses/categories/mocks"
	"github.com/Kelompok14-LMS/backend-go/businesses/courses"
	_courseMock "github.com/Kelompok14-LMS/backend-go/businesses/courses/mocks"
	_mentorMock "github.com/Kelompok14-LMS/backend-go/businesses/mentors/mocks"
	"github.com/Kelompok14-LMS/backend-go/helper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var (
	courseRepository _courseMock.Repository
	courseUsecase    courses.Usecase

	mentorRepository   _mentorMock.Repository
	categoryRepository _categoryMock.CategoryRepositoryMock
	storageClient      helper.StorageConfig
)

func TestMain(m *testing.M) {
	courseRepository = _courseMock.Repository{Mock: mock.Mock{}}
	mentorRepository = _mentorMock.Repository{Mock: mock.Mock{}}
	categoryRepository = _categoryMock.CategoryRepositoryMock{Mock: mock.Mock{}}
	storageClient = helper.StorageConfig{}

	courseUsecase = courses.NewCourseUsecase(&courseRepository, &mentorRepository, &categoryRepository, &storageClient)

	m.Run()
}

func TestFindAll(t *testing.T) {
	t.Run("FindAll | Success", func(t *testing.T) {
		mockCourse := []courses.Domain{
			{
				ID:          "COURSE_1",
				MentorId:    "MENTOR_1",
				CategoryId:  "CAT_1",
				Title:       "Course Title",
				Description: "Course description",
				Thumbnail:   "https://imgurl.com/bucket/object",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
				DeletedAt:   gorm.DeletedAt{},
			},
		}

		courseRepository.Mock.On("FindAll", "").Return(&mockCourse, nil)

		results, err := courseUsecase.FindAll("")

		assert.Nil(t, err)
		assert.NotNil(t, results)
	})
}

func TestFindById(t *testing.T) {
	t.Run("FindById | Success", func(t *testing.T) {
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

		result, err := courseUsecase.FindById("COURSE_1")

		assert.Nil(t, err)
		assert.NotNil(t, result)
	})
}

func TestFindByCategory(t *testing.T) {
	t.Run("FindByCategory | Success", func(t *testing.T) {
		mockCategory := categories.Domain{
			ID:        "CAT_1",
			Name:      "Programming",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		categoryRepository.Mock.On("FindById", "CAT_1").Return(&mockCategory, nil)

		mockCourse := []courses.Domain{
			{
				ID:          "COURSE_1",
				MentorId:    "MENTOR_1",
				CategoryId:  "CAT_1",
				Title:       "Course Title",
				Description: "Course description",
				Thumbnail:   "https://imgurl.com/bucket/object",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
				DeletedAt:   gorm.DeletedAt{},
			},
		}

		courseRepository.Mock.On("FindByCategory", "CAT_1").Return(&mockCourse, nil)

		results, err := courseUsecase.FindByCategory("CAT_1")

		assert.Nil(t, err)
		assert.NotNil(t, results)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update | Success", func(t *testing.T) {
		mockCategory := categories.Domain{
			ID:        "CAT_1",
			Name:      "Programming",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		categoryRepository.Mock.On("FindById", "CAT_1").Return(&mockCategory, nil)

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

		courseDomain := courses.Domain{
			CategoryId:  "CAT_1",
			Title:       "Updated Title",
			Description: "Updated Description",
			Thumbnail:   "",
		}

		courseRepository.Mock.On("Update", "COURSE_1", &courseDomain).Return(nil)

		updatedCourse := courses.Domain{
			CategoryId:          "CAT_1",
			Title:               "Updated Title",
			Description:         "Updated Description",
			ThumbnailFileHeader: nil,
		}

		err := courseUsecase.Update("COURSE_1", &updatedCourse)

		assert.Nil(t, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete | Success", func(t *testing.T) {
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

		courseRepository.Mock.On("Delete", "COURSE_1").Return(nil)

		err := courseUsecase.Delete("COURSE_1")

		assert.Nil(t, err)
	})
}
