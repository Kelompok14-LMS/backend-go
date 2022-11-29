package categories

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Kelompok14-LMS/backend-go/businesses/categories"
	"github.com/Kelompok14-LMS/backend-go/businesses/categories/mocks"
	"github.com/Kelompok14-LMS/backend-go/controllers/categories/request"
	"github.com/Kelompok14-LMS/backend-go/controllers/categories/response"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

type suiteCategory struct {
	suite.Suite
	mock    mocks.CategoryUsecaseMock
	handler *CategoryController
}

func (s *suiteCategory) SetupSuite() {
	mock := mocks.CategoryUsecaseMock{}
	s.mock = mock

	s.handler = NewCategoryController(&s.mock)
}

func (s *suiteCategory) TestHandlerCreateCategory() {
	categoryInput := request.Category{
		Name: "Programming",
	}

	s.mock.Mock.On("Create", categoryInput.ToDomain()).Return(nil)

	testCases := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		Header             map[string]string
		Body               map[string]interface{}
		HasReturnBody      bool
		ExpectedResult     map[string]interface{}
	}{
		{
			Name:               "success",
			ExpectedStatusCode: http.StatusCreated,
			Method:             "POST",
			Header: map[string]string{
				"Content-Type": "application/json",
			},
			Body: map[string]interface{}{
				"name": "Programming",
			},
			HasReturnBody: true,
			ExpectedResult: map[string]interface{}{
				"code":    201,
				"status":  "success",
				"message": "Success add category",
				"data":    nil,
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(v.Method, "/categories", bytes.NewReader(res))
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.Request().Header.Set("Content-Type", v.Header["Content-Type"])

			err := s.handler.HandlerCreateCategory(ctx)
			s.NoError(err)

			s.Equal(v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp map[string]interface{}
				err := json.NewDecoder(w.Result().Body).Decode(&resp)
				s.NoError(err)

				s.Equal(v.ExpectedResult["status"], resp["status"])
				s.Equal(v.ExpectedResult["message"], resp["message"])
			}
		})
	}
}

func (s *suiteCategory) TestHandlerFindAllCategories() {
	categories := []categories.Domain{
		{
			ID:        "CID1",
			Name:      "Programming",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	s.mock.Mock.On("FindAll").Return(&categories, nil)

	testCases := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		HasReturnBody      bool
		ExpectedResult     map[string]interface{}
	}{
		{
			Name:               "success",
			ExpectedStatusCode: http.StatusOK,
			Method:             "GET",
			HasReturnBody:      true,
			ExpectedResult: map[string]interface{}{
				"code":    200,
				"status":  "success",
				"message": "Success get all categories",
				"data": []response.Category{
					{
						ID:        "CID1",
						Name:      "Programming",
						CreatedAt: time.Now(),
						UpdatedAt: time.Now(),
					},
				},
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.Name, func(t *testing.T) {
			r := httptest.NewRequest(v.Method, "/categories", nil)
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)

			err := s.handler.HandlerFindAllCategories(ctx)
			s.NoError(err)

			s.Equal(v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp map[string]interface{}
				err := json.NewDecoder(w.Result().Body).Decode(&resp)
				s.NoError(err)

				s.Equal(v.ExpectedResult["status"], resp["status"])
				s.Equal(v.ExpectedResult["message"], resp["message"])

				s.NotNil(resp["data"])
			}
		})
	}
}

func (s *suiteCategory) TestHandlerFindByIdCategory() {
	category := categories.Domain{
		ID:        "CID1",
		Name:      "Programming",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	s.mock.Mock.On("FindById", "CID1").Return(&category, nil)

	testCases := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		HasReturnBody      bool
		ExpectedResult     map[string]interface{}
	}{
		{
			Name:               "success",
			ExpectedStatusCode: http.StatusOK,
			Method:             "GET",
			HasReturnBody:      true,
			ExpectedResult: map[string]interface{}{
				"code":    200,
				"status":  "success",
				"message": "Success get category by id",
				"data": response.Category{
					ID:        "CID1",
					Name:      "Programming",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.Name, func(t *testing.T) {
			r := httptest.NewRequest(v.Method, "/categories", nil)
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/:categoryId")
			ctx.SetParamNames("categoryId")
			ctx.SetParamValues("CID1")

			err := s.handler.HandlerFindByIdCategory(ctx)
			s.NoError(err)

			s.Equal(v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp map[string]interface{}
				err := json.NewDecoder(w.Result().Body).Decode(&resp)
				s.NoError(err)

				s.Equal(v.ExpectedResult["status"], resp["status"])
				s.Equal(v.ExpectedResult["message"], resp["message"])

				s.NotNil(resp["data"])
			}
		})
	}
}

func (s *suiteCategory) TestHandlerUpdateCategory() {
	category := categories.Domain{
		ID:        "CID1",
		Name:      "Programming",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	s.mock.Mock.On("FindById", "CID1").Return(&category, nil)

	updatedCategory := request.Category{
		Name: "UI/UX",
	}

	s.mock.Mock.On("Update", "CID1", updatedCategory.ToDomain()).Return(nil)

	testCases := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		Header             map[string]string
		Body               map[string]interface{}
		HasReturnBody      bool
		ExpectedResult     map[string]interface{}
	}{
		{
			Name:               "success",
			ExpectedStatusCode: http.StatusOK,
			Method:             "PUT",
			Header: map[string]string{
				"Content-Type": "application/json",
			},
			Body: map[string]interface{}{
				"name": "UI/UX",
			},
			HasReturnBody: true,
			ExpectedResult: map[string]interface{}{
				"code":    200,
				"status":  "success",
				"message": "Success update category",
				"data":    nil,
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(v.Method, "/categories", bytes.NewReader(res))
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.Request().Header.Set("Content-Type", v.Header["Content-Type"])
			ctx.SetPath("/:categoryId")
			ctx.SetParamNames("categoryId")
			ctx.SetParamValues("CID1")

			err := s.handler.HandlerUpdateCategory(ctx)
			s.NoError(err)

			s.Equal(v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp map[string]interface{}
				err := json.NewDecoder(w.Result().Body).Decode(&resp)
				s.NoError(err)

				s.Equal(v.ExpectedResult["status"], resp["status"])
				s.Equal(v.ExpectedResult["message"], resp["message"])
			}
		})
	}
}

func TestSuiteCategory(t *testing.T) {
	suite.Run(t, new(suiteCategory))
}
