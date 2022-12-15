package mentee_assignments_test

import (
	"errors"
	"testing"
	"time"

	"github.com/Kelompok14-LMS/backend-go/businesses/assignments"
	_assignmentMock "github.com/Kelompok14-LMS/backend-go/businesses/assignments/mocks"
	menteeAssignments "github.com/Kelompok14-LMS/backend-go/businesses/menteeAssignments"
	_menteeAssignmentMock "github.com/Kelompok14-LMS/backend-go/businesses/menteeAssignments/mocks"
	"github.com/Kelompok14-LMS/backend-go/businesses/mentees"
	"github.com/Kelompok14-LMS/backend-go/helper"
	"github.com/Kelompok14-LMS/backend-go/pkg"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	menteeAssignmentRepository _menteeAssignmentMock.Repository
	assignmentRepository       _assignmentMock.Repository
	menteeAssignmentService    menteeAssignments.Usecase
	storageClient              helper.StorageConfig

	assignmentDomain        assignments.Domain
	menteeDomain            mentees.Domain
	menteeAssignmentDomain  menteeAssignments.Domain
	createMenteeAssignment  menteeAssignments.Domain
	updatedMenteeAssignment menteeAssignments.Domain
)

func TestMain(m *testing.M) {
	menteeAssignmentService = menteeAssignments.NewMenteeAssignmentUsecase(&menteeAssignmentRepository, &assignmentRepository, &storageClient)

	assignmentDomain = assignments.Domain{
		ID:          uuid.NewString(),
		CourseId:    uuid.NewString(),
		Title:       "test",
		Description: "unit test",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	menteeDomain = mentees.Domain{
		ID:       uuid.NewString(),
		UserId:   uuid.NewString(),
		Fullname: "test",
		Phone:    "03536654457",
	}

	menteeAssignmentDomain = menteeAssignments.Domain{
		ID:            uuid.NewString(),
		MenteeId:      menteeDomain.ID,
		AssignmentId:  assignmentDomain.ID,
		AssignmentURL: "test.com",
		Grade:         80,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	createMenteeAssignment = menteeAssignments.Domain{
		ID:            uuid.NewString(),
		MenteeId:      menteeDomain.ID,
		AssignmentId:  assignmentDomain.ID,
		AssignmentURL: "test.com",
	}

	updatedMenteeAssignment = menteeAssignments.Domain{
		ID:            uuid.NewString(),
		MenteeId:      menteeDomain.ID,
		AssignmentId:  assignmentDomain.ID,
		AssignmentURL: "test.com",
		PDFfile:       nil,
	}

	m.Run()
}

func TestCreate(t *testing.T) {
	t.Run("Test Create | Failed create mentee Assignment | Assignmentnot found", func(t *testing.T) {
		assignmentRepository.Mock.On("FindById", assignmentDomain.ID).Return(&assignments.Domain{}, pkg.ErrAssignmentNotFound).Once()

		err := menteeAssignmentService.Create(&createMenteeAssignment)

		assert.Error(t, err)
	})
}

func TestFindById(t *testing.T) {
	t.Run("Test Find By Id | Success get mentee Assignment by id", func(t *testing.T) {
		menteeAssignmentRepository.Mock.On("FindById", menteeAssignmentDomain.ID).Return(&menteeAssignmentDomain, nil).Once()

		result, err := menteeAssignmentService.FindById(menteeAssignmentDomain.ID)

		assert.NoError(t, err)
		assert.NotEmpty(t, result)
	})

	t.Run("Test Find By Id | Failed mentee Assignment not found", func(t *testing.T) {
		menteeAssignmentRepository.Mock.On("FindById", menteeAssignmentDomain.ID).Return(&menteeAssignments.Domain{}, pkg.ErrAssignmentMenteeNotFound).Once()

		result, err := menteeAssignmentService.FindById(menteeAssignmentDomain.ID)

		assert.Error(t, err)
		assert.Empty(t, result)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Test Update | Success update mentee Assignment", func(t *testing.T) {
		assignmentRepository.Mock.On("FindById", assignmentDomain.ID).Return(&assignmentDomain, nil).Once()

		menteeAssignmentRepository.Mock.On("FindById", menteeAssignmentDomain.ID).Return(&menteeAssignmentDomain, nil).Once()

		menteeAssignmentRepository.Mock.On("Update", menteeAssignmentDomain.ID, mock.Anything).Return(nil).Once()

		err := menteeAssignmentService.Update(menteeAssignmentDomain.ID, &updatedMenteeAssignment)

		assert.NoError(t, err)
	})

	t.Run("Test Update | Failed update mentee Assignment | Assignment not found", func(t *testing.T) {
		assignmentRepository.Mock.On("FindById", assignmentDomain.ID).Return(&assignments.Domain{}, pkg.ErrAssignmentNotFound).Once()

		err := menteeAssignmentService.Update(menteeAssignmentDomain.ID, &updatedMenteeAssignment)

		assert.Error(t, err)
	})

	t.Run("Test Update | Failed update mentee Assignment| mentee Assignment not found", func(t *testing.T) {
		assignmentRepository.Mock.On("FindById", assignmentDomain.ID).Return(&assignmentDomain, nil).Once()

		menteeAssignmentRepository.Mock.On("FindById", menteeAssignmentDomain.ID).Return(&menteeAssignments.Domain{}, pkg.ErrAssignmentMenteeNotFound).Once()

		err := menteeAssignmentService.Update(menteeAssignmentDomain.ID, &updatedMenteeAssignment)

		assert.Error(t, err)
	})

	t.Run("Test Update | Failed update mentee Assignment | error occurred", func(t *testing.T) {
		assignmentRepository.Mock.On("FindById", assignmentDomain.ID).Return(&assignmentDomain, nil).Once()

		menteeAssignmentRepository.Mock.On("FindById", menteeAssignmentDomain.ID).Return(&menteeAssignmentDomain, nil).Once()

		menteeAssignmentRepository.Mock.On("Update", menteeAssignmentDomain.ID, mock.Anything).Return(errors.New("error occurred"))

		err := menteeAssignmentService.Update(menteeAssignmentDomain.ID, &updatedMenteeAssignment)

		assert.Error(t, err)
	})
}

func TestDelete(t *testing.T) {
	// t.Run("Test Delete | Success delete mentee Assignment", func(t *testing.T) {
	// 	menteeAssignmentRepository.Mock.On("FindById", menteeAssignmentDomain.ID).Return(&menteeAssignmentDomain, nil).Once()

	// 	menteeAssignmentRepository.Mock.On("Delete", menteeAssignmentDomain.ID).Return(nil).Once()

	// 	err := menteeAssignmentService.Delete(menteeAssignmentDomain.ID)

	// 	assert.NoError(t, err)
	// })

	// t.Run("Test Delete | Failed delete mentee Assignment | mentee Assignment not found", func(t *testing.T) {
	// 	menteeAssignmentRepository.Mock.On("FindById", menteeAssignmentDomain.ID).Return(&menteeAssignments.Domain{}, pkg.ErrAssignmentMenteeNotFound).Once()

	// 	err := menteeAssignmentService.Delete(menteeAssignmentDomain.ID)

	// 	assert.Error(t, err)
	// })

	// t.Run("Test Delete | Failed delete mentee Assignment | error occurred", func(t *testing.T) {
	// 	menteeAssignmentRepository.Mock.On("FindById", menteeAssignmentDomain.ID).Return(&menteeAssignmentDomain, nil).Once()

	// 	menteeAssignmentRepository.Mock.On("Delete", menteeAssignmentDomain.ID).Return(errors.New("error occurred")).Once()

	// 	err := menteeAssignmentService.Delete(menteeAssignmentDomain.ID)

	// 	assert.Error(t, err)
	// })
}
