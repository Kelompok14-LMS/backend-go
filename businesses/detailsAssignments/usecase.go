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

	// assignmentMenteeDomain := make([]AssignmentMentee, len(assignmentMentees))

	// for _, assignmentMentee := range assignmentMentees {

	// 	assignmentMenteeDomain = append(assignmentMenteeDomain, AssignmentMentee{
	// 		AssignmentMenteeID: assignmentMentee.AssignmentId,
	// 		MenteeId:           assignmentMentee.MenteeId,
	// 		Name:               assignmentMentee.Name,
	// 		AssignmentURL:      assignmentMentee.AssignmentURL,
	// 		Grade:              assignmentMentee.Grade,
	// 		CreatedAt:          assignmentMentee.CreatedAt,
	// 		UpdatedAt:          assignmentMentee.UpdatedAt,
	// 	})
	// }

	// assignmentsDomain := make([]Assignment, len(assignments))

	// for i, assignment := range assignments {
	// 	assignmentsDomain[i].AssignmentID = assignment.ID
	// 	assignmentsDomain[i].CourseId = assignment.CourseId
	// 	assignmentsDomain[i].Title = assignment.Title
	// 	assignmentsDomain[i].Description = assignment.Description
	// 	assignmentsDomain[i].CreatedAt = assignment.CreatedAt
	// 	assignmentsDomain[i].UpdatedAt = assignment.UpdatedAt
	// }

	// for i, assignment := range assignmentsDomain {
	// 	for _, assignmentMentee := range assignmentMenteeDomain {
	// 		if assignment.AssignmentID == assignmentMentee.AssignmentId {
	// 			assignmentsDomain[i].AssignmentMentee = append(assignmentsDomain[i].AssignmentMentee, assignmentMentee)
	// 		}
	// 	}
	// }

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
