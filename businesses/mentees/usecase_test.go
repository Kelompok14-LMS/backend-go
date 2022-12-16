package mentees_test

import (
	"errors"
	"testing"
	"time"

	"github.com/Kelompok14-LMS/backend-go/businesses/mentees"
	_menteeMock "github.com/Kelompok14-LMS/backend-go/businesses/mentees/mocks"
	"github.com/Kelompok14-LMS/backend-go/constants"
	"github.com/Kelompok14-LMS/backend-go/helper"
	"github.com/Kelompok14-LMS/backend-go/pkg"
	"github.com/Kelompok14-LMS/backend-go/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	_otpMock "github.com/Kelompok14-LMS/backend-go/businesses/otp/mocks"

	"github.com/Kelompok14-LMS/backend-go/businesses/users"
	_userMock "github.com/Kelompok14-LMS/backend-go/businesses/users/mocks"
)

var (
	menteeRepository _menteeMock.Repository
	menteeService    mentees.Usecase

	otpRepository  _otpMock.Repository
	userRepository _userMock.UserRepositoryMock
	jwtConfig      utils.JWTConfig
	mailerConfig   pkg.MailerConfig
	storage        helper.StorageConfig

	menteeDomain         mentees.Domain
	menteeAuth           mentees.MenteeAuth
	menteeRegister       mentees.MenteeRegister
	menteeForgotPassword mentees.MenteeForgotPassword

	userDomain users.Domain
)

func TestMain(m *testing.M) {
	menteeService = mentees.NewMenteeUsecase(&menteeRepository, &userRepository, &otpRepository, &jwtConfig, &mailerConfig, &storage)

	userDomain = users.Domain{
		ID:        uuid.NewString(),
		Email:     "test@gmail.com",
		Password:  "testtest",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	menteeDomain = mentees.Domain{
		ID:             uuid.NewString(),
		UserId:         userDomain.ID,
		Fullname:       "test",
		Phone:          "test",
		Role:           "mentee",
		BirthDate:      time.Now(),
		Address:        "test",
		ProfilePicture: "test.com",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	menteeAuth = mentees.MenteeAuth{
		Email:    userDomain.Email,
		Password: userDomain.Password,
	}

	menteeRegister = mentees.MenteeRegister{
		Fullname: menteeDomain.Fullname,
		Phone:    menteeDomain.Phone,
		Email:    userDomain.Email,
		Password: userDomain.Password,
		OTP:      "0000",
	}

	menteeForgotPassword = mentees.MenteeForgotPassword{
		Email:            userDomain.Email,
		Password:         userDomain.Password,
		RepeatedPassword: userDomain.Password,
		OTP:              "0000",
	}

	m.Run()
}

func TestRegister(t *testing.T) {
	t.Run("Test Register | Success register", func(t *testing.T) {
		userRepository.Mock.On("FindByEmail", menteeAuth.Email).Return(nil, nil).Once()

		otpRepository.Mock.On("Save", mock.Anything, menteeAuth.Email, mock.Anything, constants.TIME_TO_LIVE).Return(nil).Once()

		err := menteeService.Register(&menteeAuth)

		assert.NoError(t, err)
	})

	t.Run("Test Register | Failed register | Invalid password length", func(t *testing.T) {
		userRepository.Mock.On("FindByEmail", menteeAuth.Email).Return(nil, nil).Once()

		otpRepository.Mock.On("Save", mock.Anything, menteeAuth.Email, mock.Anything, constants.TIME_TO_LIVE).Return(nil).Once()

		menteeAuth.Password = "test"

		err := menteeService.Register(&menteeAuth)

		assert.Error(t, err)
	})

	t.Run("Test Register | Failed register | Error set otp", func(t *testing.T) {
		userRepository.Mock.On("FindByEmail", menteeAuth.Email).Return(nil, nil).Once()

		otpRepository.Mock.On("Save", mock.Anything, menteeAuth.Email, mock.Anything, constants.TIME_TO_LIVE).Return(errors.New("error occured")).Once()

		err := menteeService.Register(&menteeAuth)

		assert.Error(t, err)
	})
}

func TestVerifyRegister(t *testing.T) {
	t.Run("Test VerifyRegister | Success register verified", func(t *testing.T) {
		otpRepository.Mock.On("Get", mock.Anything, menteeRegister.Email).Return(menteeRegister.OTP, nil).Once()

		userRepository.Mock.On("Create", mock.Anything).Return(nil).Once()

		menteeRepository.Mock.On("Create", mock.Anything).Return(nil).Once()

		err := menteeService.VerifyRegister(&menteeRegister)

		assert.NoError(t, err)
	})

	t.Run("Test VerifyRegister | Failed register | OTP not match", func(t *testing.T) {
		otpRepository.Mock.On("Get", mock.Anything, menteeRegister.Email).Return("9999", pkg.ErrOTPNotMatch).Once()

		err := menteeService.VerifyRegister(&menteeRegister)

		assert.Error(t, err)
	})

	t.Run("Test VerifyRegister | Failed register verified | Error on create user", func(t *testing.T) {
		otpRepository.Mock.On("Get", mock.Anything, menteeRegister.Email).Return(menteeRegister.OTP, nil).Once()

		userRepository.Mock.On("Create", mock.Anything).Return(errors.New("error occurred")).Once()

		err := menteeService.VerifyRegister(&menteeRegister)

		assert.Error(t, err)
	})

	t.Run("Test VerifyRegister | Failed register verified | Error on create mentee", func(t *testing.T) {
		otpRepository.Mock.On("Get", mock.Anything, menteeRegister.Email).Return(menteeRegister.OTP, nil).Once()

		userRepository.Mock.On("Create", mock.Anything).Return(nil).Once()

		menteeRepository.Mock.On("Create", mock.Anything).Return(errors.New("error occurred"))

		err := menteeService.VerifyRegister(&menteeRegister)

		assert.Error(t, err)
	})
}

// func TestForgotPassword(t *testing.T) {}

// func TestLogin(t *testing.T) {}
