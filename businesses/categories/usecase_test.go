package categories_test

import (
	"testing"

	"github.com/Kelompok14-LMS/backend-go/businesses/categories"
	_categoryMock "github.com/Kelompok14-LMS/backend-go/businesses/categories/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	categoryRepository _categoryMock.CategoryRepositoryMock
	categoryService    categories.Usecase

	categoryDomain categories.Domain
)

func TestMain(m *testing.M) {
	categoryRepository = _categoryMock.CategoryRepositoryMock{Mock: mock.Mock{}}

	categoryService = categories.NewCategoryUsecase(&categoryRepository)

	categoryDomain = categories.Domain{
		ID:   "CID1",
		Name: "Programming",
	}

	m.Run()
}

func TestCreate(t *testing.T) {
	t.Run("Create | Success", func(t *testing.T) {
		categoryRepository.Mock.On("Create", mock.Anything).Return(nil)

		err := categoryService.Create(&categoryDomain)

		assert.NoError(t, err)
	})
}

func TestFindAll(t *testing.T) {
	t.Run("FindAll | Success", func(t *testing.T) {
		var categories []categories.Domain

		categories = append(categories, categoryDomain)

		categoryRepository.Mock.On("FindAll").Return(&categories, nil)

		results, err := categoryService.FindAll()

		assert.NoError(t, err)
		assert.NotNil(t, results)
	})
}

func TestFindById(t *testing.T) {
	t.Run("FindById", func(t *testing.T) {
		categoryRepository.Mock.On("FindById", "CID1").Return(&categoryDomain, nil)

		result, err := categoryService.FindById("CID1")

		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update | Success", func(t *testing.T) {
		categoryRepository.Mock.On("FindById", "CID1").Return(&categoryDomain, nil)

		categoryRepository.Mock.On("Update", "CID1", &categoryDomain).Return(nil)

		err := categoryService.Update("CID1", &categoryDomain)

		assert.NoError(t, err)
	})
}
