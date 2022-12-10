package mentee_progresses_test

import (
	"errors"
	"testing"
	"time"

	"github.com/Kelompok14-LMS/backend-go/businesses/courses"
	_courseMock "github.com/Kelompok14-LMS/backend-go/businesses/courses/mocks"
	"github.com/Kelompok14-LMS/backend-go/businesses/materials"
	_materialMock "github.com/Kelompok14-LMS/backend-go/businesses/materials/mocks"
	menteeProgresses "github.com/Kelompok14-LMS/backend-go/businesses/menteeProgresses"
	_menteeProgressMock "github.com/Kelompok14-LMS/backend-go/businesses/menteeProgresses/mocks"
	"github.com/Kelompok14-LMS/backend-go/businesses/mentees"
	_menteeMock "github.com/Kelompok14-LMS/backend-go/businesses/mentees/mocks"
	"github.com/Kelompok14-LMS/backend-go/pkg"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	menteeProgressRepository _menteeProgressMock.Repository
	menteeRepository         _menteeMock.MenteeRepositoryMock
	courseRepository         _courseMock.Repository
	materialRepository       _materialMock.Repository

	menteeProgressService menteeProgresses.Usecase

	menteeProgressDomain menteeProgresses.Domain
	menteeDomain         mentees.Domain
	courseDomain         courses.Domain
	materialDomain       materials.Domain
)

func TestMain(m *testing.M) {
	menteeProgressService = menteeProgresses.NewMenteeProgressUsecase(
		&menteeProgressRepository,
		&menteeRepository,
		&courseRepository,
		&materialRepository,
	)

	courseDomain = courses.Domain{
		ID:          uuid.NewString(),
		MentorId:    uuid.NewString(),
		CategoryId:  uuid.NewString(),
		Title:       "test",
		Description: "test",
		Thumbnail:   "test.com",
	}

	menteeDomain = mentees.Domain{
		ID:             uuid.NewString(),
		UserId:         uuid.NewString(),
		Fullname:       "test",
		Phone:          "test",
		Role:           "mentee",
		BirthDate:      time.Now(),
		Address:        "test",
		ProfilePicture: "test.com",
	}

	materialDomain = materials.Domain{
		ID:          uuid.NewString(),
		ModuleId:    uuid.NewString(),
		Title:       "test",
		URL:         "test.com",
		Description: "test",
	}

	menteeProgressDomain = menteeProgresses.Domain{
		ID:         uuid.NewString(),
		MenteeId:   menteeDomain.ID,
		CourseId:   courseDomain.ID,
		MaterialId: materialDomain.ID,
	}

	m.Run()
}

func TestAdd(t *testing.T) {
	t.Run("Test Add | Success add progress", func(t *testing.T) {
		menteeRepository.Mock.On("FindById", menteeDomain.ID).Return(&menteeDomain, nil).Once()

		courseRepository.Mock.On("FindById", courseDomain.ID).Return(&courseDomain, nil).Once()

		materialRepository.Mock.On("FindById", materialDomain.ID).Return(&materialDomain, nil).Once()

		menteeProgressRepository.Mock.On("Add", mock.Anything).Return(nil).Once()

		err := menteeProgressService.Add(&menteeProgressDomain)

		assert.NoError(t, err)
	})

	t.Run("Test Add | Failed add progress | Mentee not found", func(t *testing.T) {
		menteeRepository.Mock.On("FindById", menteeDomain.ID).Return(&mentees.Domain{}, pkg.ErrMenteeNotFound).Once()

		err := menteeProgressService.Add(&menteeProgressDomain)

		assert.Error(t, err)
	})

	t.Run("Test Add | Failed add progress | Course not found", func(t *testing.T) {
		menteeRepository.Mock.On("FindById", menteeDomain.ID).Return(&menteeDomain, nil).Once()

		courseRepository.Mock.On("FindById", courseDomain.ID).Return(&courses.Domain{}, pkg.ErrCourseNotFound).Once()

		err := menteeProgressService.Add(&menteeProgressDomain)

		assert.Error(t, err)
	})

	t.Run("Test Add | Failed add progress | Material not found", func(t *testing.T) {
		menteeRepository.Mock.On("FindById", menteeDomain.ID).Return(&menteeDomain, nil).Once()

		courseRepository.Mock.On("FindById", courseDomain.ID).Return(&courseDomain, nil).Once()

		materialRepository.Mock.On("FindById", materialDomain.ID).Return(&materials.Domain{}, pkg.ErrMaterialNotFound).Once()

		err := menteeProgressService.Add(&menteeProgressDomain)

		assert.Error(t, err)
	})

	t.Run("Test Add | Failed add progress | Error occurred", func(t *testing.T) {
		menteeRepository.Mock.On("FindById", menteeDomain.ID).Return(&menteeDomain, nil).Once()

		courseRepository.Mock.On("FindById", courseDomain.ID).Return(&courseDomain, nil).Once()

		materialRepository.Mock.On("FindById", materialDomain.ID).Return(&materials.Domain{}, pkg.ErrMaterialNotFound).Once()

		menteeProgressRepository.Mock.On("Add", mock.Anything).Return(errors.New("failed add progress")).Once()

		err := menteeProgressService.Add(&menteeProgressDomain)

		assert.Error(t, err)
	})
}
