package modules

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Kelompok14-LMS/backend-go/businesses/modules"
	"github.com/Kelompok14-LMS/backend-go/businesses/modules/mocks"
	"github.com/Kelompok14-LMS/backend-go/controllers/modules/request"
	"github.com/Kelompok14-LMS/backend-go/controllers/modules/response"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

type suiteModule struct {
	suite.Suite
	mock    *mocks.Usecase
	handler *ModuleController
}

func (s *suiteModule) SetupSuite() {
	mock := mocks.Usecase{}
	s.mock = &mock

	s.handler = NewModuleController(s.mock)
}

func (s *suiteModule) TestHandlerCreateModule() {
	moduleInput := request.CreateModuleInput{
		CourseId: "COURSE_1",
		Title:    "Course Title",
	}

	s.mock.Mock.On("Create", moduleInput.ToDomain()).Return(nil)

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
				"course_id": "COURSE_1",
				"title":     "Course Title",
			},
			HasReturnBody: true,
			ExpectedResult: map[string]interface{}{
				"code":    201,
				"status":  "success",
				"message": "Success create module",
				"data":    nil,
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(v.Method, "/modules", bytes.NewReader(res))
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.Request().Header.Set("Content-Type", v.Header["Content-Type"])

			err := s.handler.HandlerCreateModule(ctx)
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

func (s *suiteModule) TestHandlerFindByIdModule() {
	moduleDomain := modules.Domain{
		ID:        "MOD_1",
		CourseId:  "COURSE_1",
		Title:     "Course Title",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	s.mock.Mock.On("FindById", "MOD_1").Return(&moduleDomain, nil)

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
				"message": "Success get module by id",
				"data": response.FindByIdModule{
					ID:        "MOD_1",
					CourseId:  "COURSE_1",
					Title:     "Course Title",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.Name, func(t *testing.T) {
			r := httptest.NewRequest(v.Method, "/modules", nil)
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/:moduleId")
			ctx.SetParamNames("moduleId")
			ctx.SetParamValues("MOD_1")

			err := s.handler.HandlerFindByIdModule(ctx)
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

func (s *suiteModule) TestHandlerUpdateModule() {
	moduleInput := request.UpdateModuleInput{
		CourseId: "COURSE_1",
		Title:    "Course Title Updated",
	}

	s.mock.Mock.On("Update", "MOD_1", moduleInput.ToDomain()).Return(nil)

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
				"course_id": "COURSE_1",
				"title":     "Course Title Updated",
			},
			HasReturnBody: true,
			ExpectedResult: map[string]interface{}{
				"code":    200,
				"status":  "success",
				"message": "Success update module",
				"data":    nil,
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(v.Method, "/modules", bytes.NewReader(res))
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.Request().Header.Set("Content-Type", v.Header["Content-Type"])
			ctx.SetPath("/:moduleId")
			ctx.SetParamNames("moduleId")
			ctx.SetParamValues("MOD_1")

			err := s.handler.HandlerUpdateModule(ctx)
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

func (s *suiteModule) TestHandlerDeleteModule() {
	s.mock.Mock.On("Delete", "MOD_1").Return(nil)

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
				"message": "Module deleted",
				"data":    nil,
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.Name, func(t *testing.T) {
			r := httptest.NewRequest(v.Method, "/modules", nil)
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/:moduleId")
			ctx.SetParamNames("moduleId")
			ctx.SetParamValues("MOD_1")

			err := s.handler.HandlerDeleteModule(ctx)
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

func TestSuiteModule(t *testing.T) {
	suite.Run(t, new(suiteModule))
}
