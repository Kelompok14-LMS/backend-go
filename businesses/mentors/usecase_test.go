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
	mentorUpdate   mentors.MentorUpdateProfile

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

	mentorUpdate = mentors.MentorUpdateProfile{
		ID:             "MID1",
		UserID:         "UID1",
		FullName:       "Mentors Test",
		Phone:          "0857654378",
		BirthDate:      birthDate,
		Address:        "Jl Ahmad Yani",
		ProfilePicture: "https://example.com/to/bucket",
	}

	// mentorAuth = mentors.MentorAuth{
	// 	Email:    "mentor@gmail.com",
	// 	Password: "hashedpassword",
	// }

	mentorRegister = mentors.MentorRegister{
		FullName: "Mentor Test",
		Email:    "mentor@gmail.com",
		Password: "hashedpassword",
	}

	// userDomain = users.Domain{
	// 	ID:        "UID1",
	// 	Email:     "mentor@gmail.com",
	// 	Password:  "hashedpassword",
	// 	CreatedAt: time.Now(),
	// 	UpdatedAt: time.Now(),
	// }

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

// func TestLogin(t *testing.T) {
// 	t.Run("Login | Success", func(t *testing.T) {

// 		userDomain = users.Domain{
// 			ID:        "UID1",
// 			Email:     "mentor@gmail.com",
// 			Password:  "hashedpassword",
// 			CreatedAt: time.Now(),
// 			UpdatedAt: time.Now(),
// 		}

// 		mentorAuth = mentors.MentorAuth{
// 			Email:    "mentor@gmail.com",
// 			Password: "hashedpassword",
// 		}

// 		userRepository.Mock.On("FindByEmail", mentorAuth.Email).Return(&userDomain, nil)

// 		token, err := mentorService.Login(&mentorAuth)

// 		assert.Error(t, err)
// 		assert.NotNil(t, &token)
// 	})
// }

func TestFindById(t *testing.T) {
	t.Run("Find By ID | Valid", func(t *testing.T) {
		mentorRepository.On("FindById", "MID1").Return(&mentorDomain, nil).Once()

		result, err := mentorService.FindById("MID1")
		assert.Nil(t, err)
		assert.NotNil(t, result)
	})

	t.Run("Find By ID | InValid", func(t *testing.T) {
		mentorRepository.On("FindById", "-1").Return(&mentorDomain, nil).Once()

		result, err := mentorService.FindById("-1")
		assert.Nil(t, err)
		assert.NotNil(t, result)
	})
}

func TestFindAll(t *testing.T) {
	t.Run("Find All | Valid", func(t *testing.T) {
		mentorRepository.On("FindAll").Return(&[]mentors.Domain{mentorDomain}, nil).Once()

		result, err := mentorRepository.FindAll()
		assert.Nil(t, err)
		assert.Equal(t, 1, len(*result))
	})

	t.Run("Find All | InValid", func(t *testing.T) {
		mentorRepository.On("FindAll").Return(&[]mentors.Domain{}, nil).Once()

		result, err := mentorRepository.FindAll()
		assert.Nil(t, err)
		assert.Equal(t, 0, len(*result))
	})
}
