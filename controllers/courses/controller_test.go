package courses

import (
	"database/sql/driver"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Kelompok14-LMS/backend-go/businesses/courses"
	"github.com/Kelompok14-LMS/backend-go/businesses/courses/mocks"
	"github.com/Kelompok14-LMS/backend-go/controllers/courses/response"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type Anytime struct{}

func (a Anytime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)

	return ok
}

type suiteCourse struct {
	suite.Suite
	mock    *mocks.Usecase
	handler *CourseController
}

func (s *suiteCourse) SetupSuite() {
	mock := mocks.Usecase{}
	s.mock = &mock

	s.handler = NewCourseController(s.mock)
}

func (s *suiteCourse) TestHandlerFindAllCourses() {
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

	s.mock.On("FindAll", "").Return(&mockCourse, nil)

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
				"message": "Success get all courses",
				"data": []response.FindCourses{
					{
						CourseId:  "COURSE_1",
						Mentor:    "MENTOR_1",
						Category:  "CAT_1",
						Title:     "Course Title",
						Thumbnail: "https://imgurl.com/bucket/object",
						CreatedAt: time.Now(),
						UpdatedAt: time.Now(),
					},
				},
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.Name, func(t *testing.T) {
			r := httptest.NewRequest(v.Method, "/courses", nil)
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.QueryParam("")

			err := s.handler.HandlerFindAllCourses(ctx)
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

func (s *suiteCourse) TestHadlerFindById() {
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

	s.mock.On("FindById", "COURSE_1").Return(&mockCourse, nil)

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
				"message": "Success get course by id",
				"data": map[string]interface{}{
					"course_id":   "COURSE_1",
					"mentor":      "Mentor 1",
					"category":    "Programming",
					"title":       "Course Title",
					"description": "Course description",
					"thumbnail":   "https://imgurl.com/bucket/object",
					"created_at":  time.Now(),
					"updated_at":  time.Now(),
				},
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.Name, func(t *testing.T) {
			r := httptest.NewRequest(v.Method, "/courses", nil)
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/:courseId")
			ctx.SetParamNames("courseId")
			ctx.SetParamValues("COURSE_1")

			err := s.handler.HandlerFindByIdCourse(ctx)
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

func (s *suiteCourse) TestHandlerFindByCategory() {
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

	s.mock.On("FindByCategory", "CAT_1").Return(&mockCourse, nil)

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
				"message": "Success get courses by category",
				"data": map[string]interface{}{
					"course_id":   "COURSE_1",
					"mentor":      "Mentor 1",
					"category":    "Programming",
					"title":       "Course Title",
					"description": "Course description",
					"thumbnail":   "https://imgurl.com/bucket/object",
					"created_at":  time.Now(),
					"updated_at":  time.Now(),
				},
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.Name, func(t *testing.T) {
			r := httptest.NewRequest(v.Method, "/courses", nil)
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/:categoryId")
			ctx.SetParamNames("categoryId")
			ctx.SetParamValues("CAT_1")

			err := s.handler.HandlerFindByCategory(ctx)
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

func (s *suiteCourse) TestHandlerSoftDeleteCourse() {
	s.mock.On("Delete", "COURSE_1").Return(nil)

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
				"message": "Course deleted",
				"data":    nil,
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.Name, func(t *testing.T) {
			r := httptest.NewRequest(v.Method, "/courses", nil)
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/:courseId")
			ctx.SetParamNames("courseId")
			ctx.SetParamValues("COURSE_1")

			err := s.handler.HandlerSoftDeleteCourse(ctx)
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

func TestSuiteCourse(t *testing.T) {
	suite.Run(t, new(suiteCourse))
}
