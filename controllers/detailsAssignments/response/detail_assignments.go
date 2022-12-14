package response

import (
	"time"

	detailAssignment "github.com/Kelompok14-LMS/backend-go/businesses/detailsAssignments"
)

type Assignment struct {
	AssignmentID     string             `json:"assignment_id"`
	CourseId         string             `json:"course_id"`
	NameCourse       string             `json:"course_title"`
	Title            string             `json:"title"`
	Description      string             `json:"description"`
	AssignmentMentee []AssignmentMentee `json:"assignment_mentee"`
	CreatedAt        time.Time          `json:"created_at"`
	UpdatedAt        time.Time          `json:"updated_at"`
}

type AssignmentMentee struct {
	AssignmentMenteeID string    `json:"assignment_mentee_id"`
	MenteeId           string    `json:"mentee_id"`
	AssignmentId       string    `json:"assignment_id"`
	Name               string    `json:"name"`
	AssignmentURL      string    `json:"assignment_url"`
	Grade              int       `json:"grade"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

func FullDetailAssignement(domain *detailAssignment.Assignment) *Assignment {

	var assignmentMentee []AssignmentMentee

	for _, assignment := range domain.AssignmentMentee {
		assignmentMentee = append(assignmentMentee, AssignmentMentee{
			AssignmentMenteeID: assignment.ID,
			MenteeId:           assignment.MenteeId,
			AssignmentId:       assignment.AssignmentId,
			Name:               assignment.Name,
			AssignmentURL:      assignment.AssignmentURL,
			Grade:              assignment.Grade,
			CreatedAt:          assignment.CreatedAt,
			UpdatedAt:          assignment.UpdatedAt,
		})
	}

	return &Assignment{
		AssignmentID:     domain.AssignmentID,
		CourseId:         domain.CourseId,
		NameCourse:       domain.NameCourse,
		Title:            domain.Title,
		Description:      domain.Description,
		AssignmentMentee: assignmentMentee,
		CreatedAt:        domain.CreatedAt,
		UpdatedAt:        domain.UpdatedAt,
	}
}
