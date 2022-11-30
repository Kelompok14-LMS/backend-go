package mentors_test

import (
	"testing"
	"time"

	"github.com/Kelompok14-LMS/backend-go/businesses/mentors"
	_mentorMock "github.com/Kelompok14-LMS/backend-go/businesses/mentors/mocks"
	"github.com/Kelompok14-LMS/backend-go/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/Kelompok14-LMS/backend-go/businesses/users"
	_userMock "github.com/Kelompok14-LMS/backend-go/businesses/users/mocks"
)

var (
	mentorRepository _mentorMock.Repository
	mentorService    mentors.Usecase

	userRepository _userMock.UserRepositoryMock
	jwtConfig      utils.JWTConfig

	mentorDomain   mentors.Domain
	mentorAuth     mentors.MentorAuth
	mentorRegister mentors.MentorRegister

	userDomain users.Domain
)

func TestMain(m *testing.M) {
	mentorRepository = _mentorMock.Repository{Mock: mock.Mock{}}
	userRepository = _userMock.UserRepositoryMock{Mock: mock.Mock{}}
	jwtConfig = utils.JWTConfig{JWTSecret: "secret"}

	mentorService = mentors.NewMentorUsecase(&mentorRepository, &userRepository, &jwtConfig)

	// birth date
	birthDate := time.Date(2021, 8, 11, 0, 0, 0, 0, time.Local)

	mentorDomain = mentors.Domain{
		ID:             "MID1",
		UserId:         "UID1",
		FullName:       "Mentors Test",
		Phone:          "0857654378",
		Role:           "mentor",
		BirthDate:      birthDate,
		Address:        "Jl Ahmad Yani",
		ProfilePicture: "https://example.com/to/bucket",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	mentorAuth = mentors.MentorAuth{
		Email:    "mentor@gmail.com",
		Password: "hashedpassword",
	}

	mentorRegister = mentors.MentorRegister{
		FullName: "Mentor Test",
		Email:    "mentor@gmail.com",
		Password: "hashedpassword",
	}

	userDomain = users.Domain{
		ID:        "UID1",
		Email:     "mentor@gmail.com",
		Password:  "hashedpassword",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	m.Run()
}

func TestRegister(t *testing.T) {
	t.Run("Register | Success", func(t *testing.T) {

		userRepository.Mock.On("FindByEmail", mentorRegister.Email).Return(nil, nil)

		userRepository.Mock.On("Create", mock.Anything).Return(nil)

		mentorRepository.Mock.On("Create", mock.Anything).Return(nil)

		err := mentorService.Register(&mentorRegister)

		assert.NoError(t, err)
	})
}

// TODO: Create test Login
