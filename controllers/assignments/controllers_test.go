package assignments

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Kelompok14-LMS/backend-go/businesses/assignments"
	"github.com/Kelompok14-LMS/backend-go/businesses/assignments/mocks"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

type suiteAssignment struct {
	suite.Suite
	mock    *mocks.Usecase
	handler *AssignmentController
}

func (s *suiteAssignment) SetupSuite() {
	mock := mocks.Usecase{}
	s.mock = &mock

	s.handler = NewAssignmentsController(s.mock)
}

func (s *suiteAssignment) TestHandlerFindById() {
	mockAssignment := assignments.Domain{
		ID:          "ASSIGNMENT_1",
		ModuleID:    "MODULE_1",
		Title:       "Title test",
		Description: "Description test",
		PDFurl:      "https://storage.com/to/bucket/test.pdf",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	s.mock.Mock.On("FindById", "ASSIGNMENT_1").Return(&mockAssignment, nil)

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
				"message": "Success get assignment by id",
				"data": map[string]interface{}{
					"id":          "ASSIGNMENT_1",
					"module_id":   "MODULE_1",
					"title":       "Title test",
					"description": "Description test",
					"pdf":         "https://storage.com/to/bucket/object.mp4",
					"created_at":  time.Now(),
					"updated_at":  time.Now(),
				},
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.Name, func(t *testing.T) {
			r := httptest.NewRequest(v.Method, "/assignments", nil)
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/:assignmentId")
			ctx.SetParamNames("assignmentId")
			ctx.SetParamValues("ASSIGNMENT_1")

			err := s.handler.HandlerFindByIdAssignment(ctx)
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

func (s *suiteAssignment) TestHandlerDelete() {
	s.mock.Mock.On("Delete", "ASSIGNMENT_1").Return(nil)

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
				"message": "Assignment deleted",
				"data":    nil,
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.Name, func(t *testing.T) {
			r := httptest.NewRequest(v.Method, "/assignments", nil)
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/:assignmentId")
			ctx.SetParamNames("assignmentId")
			ctx.SetParamValues("ASSIGNMENT_1")

			err := s.handler.HandlerDeleteAssignment(ctx)
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

func TestSuiteAssignments(t *testing.T) {
	suite.Run(t, new(suiteAssignment))
}
