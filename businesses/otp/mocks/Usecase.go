package mocks

import (
	"github.com/Kelompok14-LMS/backend-go/businesses/otp"
	"github.com/stretchr/testify/mock"
)

type OTPUsecaseMock struct {
	Mock mock.Mock
}

func (o *OTPUsecaseMock) SendOTP(otpDomain *otp.Domain) error {
	ret := o.Mock.Called(otpDomain)

	return ret.Error(0)
}

func (o *OTPUsecaseMock) CheckOTP(otpDomain *otp.Domain) error {
	ret := o.Mock.Called(otpDomain)

	return ret.Error(0)
}
