package mentors

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Kelompok14-LMS/backend-go/businesses/mentors/mocks"
	"github.com/Kelompok14-LMS/backend-go/controllers/mentors/request"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

type suiteMentor struct {
	suite.Suite
	mock    *mocks.Usecase
	handler *MentorController
}

func (s *suiteMentor) SetupSuite() {
	mock := mocks.Usecase{}
	s.mock = &mock

	s.handler = &MentorController{
		mentorUsecase: s.mock,
	}
}

func (s *suiteMentor) TestHandlerRegister() {
	mentorInput := request.MentorRegisterInput{
		FullName: "Mentor Test",
		Email:    "mentor@gmail.com",
		Password: "12345678",
	}

	s.mock.Mock.On("Register", mentorInput.ToDomain()).Return(nil)

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
				"fullname": "Mentor Test",
				"email":    "mentor@gmail.com",
				"password": "12345678",
			},
			HasReturnBody: true,
			ExpectedResult: map[string]interface{}{
				"code":    201,
				"status":  "success",
				"message": "Register success",
				"data":    nil,
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(v.Method, "/auth/mentor/register", bytes.NewReader(res))
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.Request().Header.Set("Content-Type", v.Header["Content-Type"])

			err := s.handler.HandlerRegisterMentor(ctx)
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

func (s *suiteMentor) TestHandlerLogin() {
	mentorInput := request.AuthMentorInput{
		Email:    "mentor@gmail.com",
		Password: "123456",
	}

	token := "token"

	s.mock.Mock.On("Login", mentorInput.ToDomain()).Return(&token, nil)

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
			Method:             "POST",
			Header: map[string]string{
				"Content-Type": "application/json",
			},
			Body: map[string]interface{}{
				"email":    "mentor@gmail.com",
				"password": "123456",
			},
			HasReturnBody: true,
			ExpectedResult: map[string]interface{}{
				"code":    200,
				"status":  "success",
				"message": "Login successful",
				"data":    nil,
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(v.Method, "/auth/mentor/login", bytes.NewReader(res))
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.Request().Header.Set("Content-Type", v.Header["Content-Type"])

			err := s.handler.HandlerLoginMentor(ctx)
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

func TestMentorController(t *testing.T) {
	suite.Run(t, new(suiteMentor))
}
