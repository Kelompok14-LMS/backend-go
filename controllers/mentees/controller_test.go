package mentees

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Kelompok14-LMS/backend-go/businesses/mentees/mocks"
	"github.com/Kelompok14-LMS/backend-go/controllers/mentees/request"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

type suiteMentee struct {
	suite.Suite
	mock    *mocks.MenteeUsecaseMock
	handler *MenteeController
}

func (s *suiteMentee) SetupSuite() {
	mock := mocks.MenteeUsecaseMock{}
	s.mock = &mock

	s.handler = &MenteeController{
		menteeUsecase: s.mock,
	}
}

func (s *suiteMentee) TestHandlerRegisterMentee() {
	menteeInput := request.AuthMenteeInput{
		Email:    "mentee@gmail.com",
		Password: "123456",
	}

	s.mock.Mock.On("Register", menteeInput.ToDomain()).Return(nil)

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
				"email":    "mentee@gmail.com",
				"password": "123456",
			},
			HasReturnBody: true,
			ExpectedResult: map[string]interface{}{
				"code":    200,
				"status":  "success",
				"message": "Success send OTP to email",
				"data":    nil,
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(v.Method, "/auth/mentee/register", bytes.NewReader(res))
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.Request().Header.Set("Content-Type", v.Header["Content-Type"])

			err := s.handler.HandlerRegisterMentee(ctx)
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

func (s *suiteMentee) TestHandlerVerifyRegister() {
	menteeInput := request.MenteeRegisterInput{
		Fullname: "Mentee Test",
		Phone:    "087654321",
		Email:    "mentee@gmail.com",
		Password: "12345678",
		OTP:      "7339",
	}

	s.mock.Mock.On("VerifyRegister", menteeInput.ToDomain()).Return(nil)

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
				"fullname": "Mentee Test",
				"phone":    "087654321",
				"email":    "mentee@gmail.com",
				"password": "12345678",
				"otp":      "7339",
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
			r := httptest.NewRequest(v.Method, "/auth/mentee/register/verify", bytes.NewReader(res))
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.Request().Header.Set("Content-Type", v.Header["Content-Type"])

			err := s.handler.HandlerVerifyRegisterMentee(ctx)
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

func (s *suiteMentee) TestHandlerLogin() {
	menteeInput := request.AuthMenteeInput{
		Email:    "mentee@gmail.com",
		Password: "123456",
	}

	token := "token"

	s.mock.Mock.On("Login", menteeInput.ToDomain()).Return(&token, nil)

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
				"email":    "mentee@gmail.com",
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
			r := httptest.NewRequest(v.Method, "/auth/mentee/login", bytes.NewReader(res))
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.Request().Header.Set("Content-Type", v.Header["Content-Type"])

			err := s.handler.HandlerLoginMentee(ctx)
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

func (s *suiteMentee) TestHandlerForgotPassword() {
	menteeInput := request.ForgotPasswordInput{
		Email:            "mentee@gmail.com",
		Password:         "123456",
		RepeatedPassword: "123456",
		OTP:              "7339",
	}

	s.mock.Mock.On("ForgotPassword", menteeInput.ToDomain()).Return(nil)

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
				"email":             "mentee@gmail.com",
				"password":          "123456",
				"repeated_password": "123456",
				"otp":               "7339",
			},
			HasReturnBody: true,
			ExpectedResult: map[string]interface{}{
				"code":    200,
				"status":  "success",
				"message": "Success reset password",
				"data":    nil,
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(v.Method, "/auth/forgot-password", bytes.NewReader(res))
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.Request().Header.Set("Content-Type", v.Header["Content-Type"])

			err := s.handler.HandlerForgotPassword(ctx)
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

func TestMenteeController(t *testing.T) {
	suite.Run(t, new(suiteMentee))
}
