package otp

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Kelompok14-LMS/backend-go/businesses/otp/mocks"
	"github.com/Kelompok14-LMS/backend-go/controllers/otp/request"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

type suiteOtp struct {
	suite.Suite
	mock    *mocks.OTPUsecaseMock
	handler *OTPController
}

func (s *suiteOtp) SetupSuite() {
	mock := mocks.OTPUsecaseMock{}
	s.mock = &mock

	s.handler = &OTPController{
		otpUsecase: s.mock,
	}
}

func (s *suiteOtp) TestSendOTP() {
	otpInput := request.OTP{
		Key: "mentee@gmail.com",
	}

	s.mock.Mock.On("SendOTP", otpInput.ToDomain()).Return(nil)

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
				"email": "mentee@gmail.com",
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
			r := httptest.NewRequest(v.Method, "/auth/send-otp", bytes.NewReader(res))
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.Request().Header.Set("Content-Type", v.Header["Content-Type"])

			err := s.handler.HandlerSendOTP(ctx)
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

func (s *suiteOtp) TestCheckOTP() {
	otpInput := request.CheckOTP{
		Key:   "mentee@gmail.com",
		Value: "7339",
	}

	s.mock.Mock.On("CheckOTP", otpInput.ToDomain()).Return(nil)

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
				"email": "mentee@gmail.com",
				"otp":   "7339",
			},
			HasReturnBody: true,
			ExpectedResult: map[string]interface{}{
				"code":    200,
				"status":  "success",
				"message": "OTP matched",
				"data":    nil,
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(v.Method, "/auth/check-otp", bytes.NewReader(res))
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.Request().Header.Set("Content-Type", v.Header["Content-Type"])

			err := s.handler.HandlerCheckOTP(ctx)
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

func TestOTPController(t *testing.T) {
	suite.Run(t, new(suiteOtp))
}
