package materials

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Kelompok14-LMS/backend-go/businesses/materials"
	"github.com/Kelompok14-LMS/backend-go/businesses/materials/mocks"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

type suiteMaterial struct {
	suite.Suite
	mock    *mocks.Usecase
	handler *MaterialController
}

func (s *suiteMaterial) SetupSuite() {
	mock := mocks.Usecase{}
	s.mock = &mock

	s.handler = NewMaterialController(s.mock)
}

func (s *suiteMaterial) TestHandlerFindById() {
	mockMaterial := materials.Domain{
		ID:          "MATERIAL_1",
		ModuleId:    "MODULE_1",
		Title:       "Title test",
		URL:         "https://storage.com/to/bucket/object.mp4",
		Description: "Description test",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	s.mock.Mock.On("FindById", "MATERIAL_1").Return(&mockMaterial, nil)

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
				"message": "Success get material by id",
				"data": map[string]interface{}{
					"id":          "MATERIAL_1",
					"module_id":   "MODULE_1",
					"title":       "Title test",
					"url":         "https://storage.com/to/bucket/object.mp4",
					"description": "Description test",
					"created_at":  time.Now(),
					"updated_at":  time.Now(),
				},
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.Name, func(t *testing.T) {
			r := httptest.NewRequest(v.Method, "/materials", nil)
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/:materialId")
			ctx.SetParamNames("materialId")
			ctx.SetParamValues("MATERIAL_1")

			err := s.handler.HandlerFindByIdMaterial(ctx)
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

func (s *suiteMaterial) TestHandlerSoftDeleteMaterial() {
	s.mock.Mock.On("Delete", "MATERIAL_1").Return(nil)

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
			Method:             "DELETE",
			HasReturnBody:      true,
			ExpectedResult: map[string]interface{}{
				"code":    200,
				"status":  "success",
				"message": "Material deleted",
				"data":    nil,
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.Name, func(t *testing.T) {
			r := httptest.NewRequest(v.Method, "/materials", nil)
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/:materialId")
			ctx.SetParamNames("materialId")
			ctx.SetParamValues("MATERIAL_1")

			err := s.handler.HandlerSoftDeleteMaterial(ctx)
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

func (s *suiteMaterial) TestHandlerSoftDeleteMaterialByModule() {
	s.mock.Mock.On("Deletes", "MODULE_1").Return(nil)

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
			Method:             "DELETE",
			HasReturnBody:      true,
			ExpectedResult: map[string]interface{}{
				"code":    200,
				"status":  "success",
				"message": "Materials deleted",
				"data":    nil,
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.Name, func(t *testing.T) {
			r := httptest.NewRequest(v.Method, "/materials", nil)
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/modules/:moduleId")
			ctx.SetParamNames("moduleId")
			ctx.SetParamValues("MODULE_1")

			err := s.handler.HandlerSoftDeleteMaterialByModule(ctx)
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

func TestSuiteMaterial(t *testing.T) {
	suite.Run(t, new(suiteMaterial))
}
