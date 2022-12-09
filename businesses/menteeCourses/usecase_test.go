package mentee_courses_test

import (
	"errors"
	"testing"
	"time"

	"github.com/Kelompok14-LMS/backend-go/businesses/courses"
	_courseMock "github.com/Kelompok14-LMS/backend-go/businesses/courses/mocks"
	menteeCourses "github.com/Kelompok14-LMS/backend-go/businesses/menteeCourses"
	_menteeCourseMock "github.com/Kelompok14-LMS/backend-go/businesses/menteeCourses/mocks"
	"github.com/Kelompok14-LMS/backend-go/businesses/mentees"
	_menteeMock "github.com/Kelompok14-LMS/backend-go/businesses/mentees/mocks"
	"github.com/Kelompok14-LMS/backend-go/pkg"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	menteeCourseRepository _menteeCourseMock.Repository
	courseRepository       _courseMock.Repository
	menteeRepository       _menteeMock.MenteeRepositoryMock

	menteeCourseService menteeCourses.Usecase

	menteeCourseDomain menteeCourses.Domain
	courseDomain       courses.Domain
	menteeDomain       mentees.Domain
)

func TestMain(m *testing.M) {
	menteeCourseService = menteeCourses.NewMenteeCourseUsecase(&menteeCourseRepository, &menteeRepository, &courseRepository)

	courseDomain = courses.Domain{
		ID:          uuid.NewString(),
		MentorId:    "test",
		CategoryId:  "test",
		Title:       "test",
		Description: "test",
		Thumbnail:   "test.com",
	}

	menteeDomain = mentees.Domain{
		ID:             uuid.NewString(),
		UserId:         "test",
		Fullname:       "test",
		Phone:          "test",
		Role:           "mentee",
		BirthDate:      time.Now(),
		Address:        "test",
		ProfilePicture: "test.com",
	}

	menteeCourseDomain = menteeCourses.Domain{
		ID:       uuid.NewString(),
		MenteeId: menteeDomain.ID,
		CourseId: courseDomain.ID,
		Status:   "ongoing",
	}

	m.Run()
}

func TestEnroll(t *testing.T) {
	t.Run("Test Enroll | Success enroll course", func(t *testing.T) {
		menteeRepository.Mock.On("FindById", menteeDomain.ID).Return(&menteeDomain, nil).Once()

		courseRepository.Mock.On("FindById", courseDomain.ID).Return(&courseDomain, nil).Once()

		menteeCourseRepository.Mock.On("Enroll", mock.Anything).Return(nil).Once()

		err := menteeCourseService.Enroll(&menteeCourseDomain)

		assert.Nil(t, err)
	})

	t.Run("Test Enroll | Mentee not found", func(t *testing.T) {
		menteeRepository.Mock.On("FindById", menteeDomain.ID).Return(&mentees.Domain{}, pkg.ErrMenteeNotFound).Once()

		err := menteeCourseService.Enroll(&menteeCourseDomain)

		assert.NotNil(t, err)
	})

	t.Run("Test Enroll | Course not found", func(t *testing.T) {
		menteeRepository.Mock.On("FindById", menteeDomain.ID).Return(&menteeDomain, nil).Once()

		courseRepository.Mock.On("FindById", courseDomain.ID).Return(&courses.Domain{}, pkg.ErrCourseNotFound).Once()

		err := menteeCourseService.Enroll(&menteeCourseDomain)

		assert.NotNil(t, err)
	})

	t.Run("Test Enroll | Failed enroll course", func(t *testing.T) {
		menteeRepository.Mock.On("FindById", menteeDomain.ID).Return(&menteeDomain, nil).Once()

		courseRepository.Mock.On("FindById", courseDomain.ID).Return(&courseDomain, nil).Once()

		menteeCourseRepository.Mock.On("Enroll", mock.Anything).Return(errors.New("failed enroll course")).Once()

		err := menteeCourseService.Enroll(&menteeCourseDomain)

		assert.NotNil(t, err)
	})
}

func TestFindMenteeCourses(t *testing.T) {
	t.Run("Test Find Mentee Courses | Success get mentee courses", func(t *testing.T) {
		menteeCourseRepository.Mock.On("FindCoursesByMentee", menteeDomain.ID, "test", "test").Return(&[]menteeCourses.Domain{menteeCourseDomain}, nil).Once()

		results, err := menteeCourseService.FindMenteeCourses(menteeDomain.ID, "test", "test")

		assert.Nil(t, err)
		assert.NotEmpty(t, results)
	})

	t.Run("Test Find Mentee Courses | Failed get mentee courses", func(t *testing.T) {
		menteeCourseRepository.Mock.On("FindCoursesByMentee", menteeDomain.ID, "test", "test").Return(&[]menteeCourses.Domain{}, pkg.ErrCourseNotFound).Once()

		results, err := menteeCourseService.FindMenteeCourses(menteeDomain.ID, "test", "test")

		assert.NotNil(t, err)
		assert.Empty(t, results)
	})
}

func TestCheckEnrollment(t *testing.T) {
	t.Run("Test Check Enrollment | Success check enrollment", func(t *testing.T) {
		menteeCourseRepository.Mock.On("CheckEnrollment", menteeCourseDomain.MenteeId, menteeCourseDomain.CourseId).Return(&menteeCourseDomain, nil).Once()

		result, err := menteeCourseService.CheckEnrollment(menteeCourseDomain.MenteeId, menteeCourseDomain.CourseId)

		assert.Nil(t, err)
		assert.True(t, result)
	})
}
