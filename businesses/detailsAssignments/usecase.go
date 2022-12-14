package details_assignments

import (
	"github.com/Kelompok14-LMS/backend-go/businesses/assignments"
	"github.com/Kelompok14-LMS/backend-go/businesses/courses"
	menteeAssignments "github.com/Kelompok14-LMS/backend-go/businesses/menteeAssignments"
)

type detailAssignmentUsecase struct {
	courseRepository           courses.Repository
	assignmentsRepository      assignments.Repository
	menteeAssignmentRepository menteeAssignments.Repository
}

func NewDetailAssignmentUsecase(
	courseRepository courses.Repository,
	assignmentsRepository assignments.Repository,
	menteeAssignmentRepository menteeAssignments.Repository,
) Usecase {
	return detailAssignmentUsecase{
		courseRepository:           courseRepository,
		assignmentsRepository:      assignmentsRepository,
		menteeAssignmentRepository: menteeAssignmentRepository,
	}
}

func (da detailAssignmentUsecase) DetailAssignment(assignmentId string) (*Assignment, error) {
	assignment, err := da.assignmentsRepository.FindById(assignmentId)
	if err != nil {
		return nil, err
	}

	course, err := da.courseRepository.FindById(assignment.CourseId)

	if err != nil {
		return nil, err
	}

	assignmentMentees, err := da.menteeAssignmentRepository.FindByAssignmentId(assignmentId)

	if err != nil {
		return nil, err
	}

	assignmentsDetail := Assignment{
		AssignmentID:     assignment.ID,
		CourseId:         course.ID,
		NameCourse:       course.Title,
		Title:            assignment.Title,
		Description:      assignment.Description,
		AssignmentMentee: assignmentMentees,
		CreatedAt:        assignment.CreatedAt,
		UpdatedAt:        assignment.UpdatedAt,
	}

	return &assignmentsDetail, nil
}
