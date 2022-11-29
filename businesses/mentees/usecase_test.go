package mentees_test

import (
	"testing"
	"time"

	"github.com/Kelompok14-LMS/backend-go/businesses/mentees"
	_menteeMock "github.com/Kelompok14-LMS/backend-go/businesses/mentees/mocks"
	"github.com/Kelompok14-LMS/backend-go/constants"
	"github.com/Kelompok14-LMS/backend-go/pkg"
	"github.com/Kelompok14-LMS/backend-go/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	_otpMock "github.com/Kelompok14-LMS/backend-go/businesses/otp/mocks"

	"github.com/Kelompok14-LMS/backend-go/businesses/users"
	_userMock "github.com/Kelompok14-LMS/backend-go/businesses/users/mocks"
)

var (
	menteeRepository _menteeMock.MenteeRepositoryMock
	menteeService    mentees.Usecase

	otpRepository  _otpMock.OTPRepositoryMock
	userRepository _userMock.UserRepositoryMock
	jwtConfig      utils.JWTConfig
	mailerConfig   pkg.MailerConfig

	menteeDomain         mentees.Domain
	menteeAuth           mentees.MenteeAuth
	menteeRegister       mentees.MenteeRegister
	menteeForgotPassword mentees.MenteeForgotPassword

	userDomain users.Domain
)

func TestMain(m *testing.M) {
	menteeRepository = _menteeMock.MenteeRepositoryMock{Mock: mock.Mock{}}
	userRepository = _userMock.UserRepositoryMock{Mock: mock.Mock{}}
	otpRepository = _otpMock.OTPRepositoryMock{Mock: mock.Mock{}}
	jwtConfig = utils.JWTConfig{JWTSecret: "secret"}
	mailerConfig = pkg.MailerConfig{}

	menteeService = mentees.NewMenteeUsecase(&menteeRepository, &userRepository, &otpRepository, &jwtConfig, &mailerConfig)

	// birth date
	birthDate := time.Date(2022, 8, 12, 0, 0, 0, 0, time.Local)

	menteeDomain = mentees.Domain{
		ID:             "MID1",
		UserId:         "UID1",
		Fullname:       "Mentee Test",
		Phone:          "0987654321",
		Role:           "mentee",
		BirthDate:      birthDate,
		Address:        "Jl Sengon",
		ProfilePicture: "https://example.com/to/bucket",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	menteeAuth = mentees.MenteeAuth{
		Email:    "mentee@gmail.com",
		Password: "hashedpassword",
	}

	menteeRegister = mentees.MenteeRegister{
		Fullname: "Mentee Test",
		Phone:    "0987654321",
		Email:    "mentee@gmail.com",
		Password: "hashedpassword",
		OTP:      "7339",
	}

	menteeForgotPassword = mentees.MenteeForgotPassword{
		Email:            "mentee@gmail.com",
		Password:         "updatedhashedpassword",
		RepeatedPassword: "updatedhashedpassword",
		OTP:              "7339",
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

// FIXME: error init get env
func TestRegister(t *testing.T) {
	t.Run("Register | Success", func(t *testing.T) {
		userRepository.Mock.On("FindByEmail", menteeAuth.Email).Return(nil, nil)

		otpRepository.Mock.On("Save", mock.Anything, menteeAuth.Email, mock.Anything, constants.TIME_TO_LIVE).Return(nil)

		err := menteeService.Register(&menteeAuth)

		assert.NoError(t, err)
	})
}

func TestVerifyRegister(t *testing.T) {
	t.Run("VerifyRegister | Success", func(t *testing.T) {
		otpRepository.Mock.On("Get", mock.Anything, menteeRegister.Email).Return("7339", nil)

		userRepository.Mock.On("Create", mock.Anything).Return(nil)

		menteeRepository.Mock.On("Create", mock.Anything).Return(nil)

		err := menteeService.VerifyRegister(&menteeRegister)

		assert.NoError(t, err)
	})
}

// TODO: Create test ForgotPassword
// func TestForgotPassword(t *testing.T) {}

// TODO: Create test Login
// func TestLogin(t *testing.T) {}
