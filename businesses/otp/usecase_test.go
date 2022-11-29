package otp_test

import (
	"testing"
	"time"

	"github.com/Kelompok14-LMS/backend-go/businesses/otp"
	_otpMock "github.com/Kelompok14-LMS/backend-go/businesses/otp/mocks"
	"github.com/Kelompok14-LMS/backend-go/businesses/users"
	_userMock "github.com/Kelompok14-LMS/backend-go/businesses/users/mocks"
	"github.com/Kelompok14-LMS/backend-go/constants"
	"github.com/Kelompok14-LMS/backend-go/pkg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	otpRepository _otpMock.OTPRepositoryMock
	otpService    otp.Usecase

	userRepository _userMock.UserRepositoryMock
	mailerConfig   pkg.MailerConfig

	otpDomain  otp.Domain
	userDomain users.Domain
)

func TestMain(m *testing.M) {
	otpRepository = _otpMock.OTPRepositoryMock{Mock: mock.Mock{}}
	userRepository = _userMock.UserRepositoryMock{Mock: mock.Mock{}}
	mailerConfig = pkg.MailerConfig{}

	otpService = otp.NewOTPUsecase(&otpRepository, &userRepository, &mailerConfig)

	otpDomain = otp.Domain{
		Key:   "mentee@gmail.com",
		Value: "7339",
	}

	userDomain = users.Domain{
		ID:        "UID1",
		Email:     "mentee@gmail.com",
		Password:  "hashedpassword",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	m.Run()
}

func TestSendOTP(t *testing.T) {
	t.Run("SendOTP | Success", func(t *testing.T) {
		userRepository.Mock.On("FindByEmail", otpDomain.Key).Return(&userDomain, nil).Once()

		otpRepository.Mock.On("Save", mock.Anything, otpDomain.Key, mock.Anything, constants.TIME_TO_LIVE).Return(nil).Once()

		err := otpService.SendOTP(&otpDomain)

		assert.Nil(t, err)
	})
}

func TestCheckOTP(t *testing.T) {
	t.Run("CheckOTP | Success", func(t *testing.T) {
		userRepository.Mock.On("FindByEmail", otpDomain.Key).Return(&userDomain, nil).Once()

		otpRepository.Mock.On("Get", mock.Anything, otpDomain.Key).Return(otpDomain.Value, nil).Once()

		err := otpService.CheckOTP(&otpDomain)

		assert.Nil(t, err)
	})

	// TODO: failed case check otp
	t.Run("CheckOTP | Failed - User Not Found", func(t *testing.T) {
		userRepository.Mock.On("FindByEmail", "notfound@gmail.com").Return(&users.Domain{}, pkg.ErrUserNotFound).Once()

		otpRepository.Mock.On("Get", mock.Anything, "notfound@gmail.com").Return("", pkg.ErrOTPExpired).Once()

		err := otpService.CheckOTP(&otp.Domain{Key: "notfound@gmail.com"})

		assert.NotNil(t, err)
	})
}
