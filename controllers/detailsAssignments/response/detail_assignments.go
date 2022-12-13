package response

import (
	"time"

	detailAssignment "github.com/Kelompok14-LMS/backend-go/businesses/detailsAssignments"
)

// type course struct {
// 	CourseId    string       `json:"course_id"`
// 	CategoryId  string       `json:"category_id"`
// 	MentorId    string       `json:"mentor_id"`
// 	Mentor      string       `json:"mentor"`
// 	Category    string       `json:"category"`
// 	Title       string       `json:"title"`
// 	Description string       `json:"description"`
// 	Thumbnail   string       `json:"thumbnail"`
// 	Assignment  []Assignment `json:"assignments"`
// 	CreatedAt   time.Time    `json:"created_at"`
// 	UpdatedAt   time.Time    `json:"updated_at"`
// }

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

	// assignments := make([]Assignment, len(domain.Assignments))

	// for _, assignment := range domain.Assignments {
	// 	assignments = append(assignments, Assignment{
	// 		AssignmentID: assignment.AssignmentID,
	// 		CourseId:     assignment.CourseId,
	// 		Title:        assignment.Title,
	// 		Description:  assignment.Description,
	// 		CreatedAt:    assignment.CreatedAt,
	// 		UpdatedAt:    assignment.UpdatedAt,
	// 	})
	// }

	// for i, assignment := range assignments {
	// 	assignment.AssignmentMentee = make([]AssignmentMentee, len(domain.Assignments[i].AssignmentMentee))

	// 	for _, assignment_mentee := range domain.Assignments[i].AssignmentMentee {
	// 		assignment.AssignmentMentee = append(assignment.AssignmentMentee, AssignmentMentee{
	// 			AssignmentMenteeID: assignment_mentee.AssignmentMenteeID,
	// 			AssignmentId:       assignment_mentee.AssignmentId,
	// 			Name:               assignment_mentee.Name,
	// 			AssignmentURL:      assignment_mentee.AssignmentURL,
	// 			Grade:              assignment_mentee.Grade,
	// 			CreatedAt:          assignment.CreatedAt,
	// 			UpdatedAt:          assignment.UpdatedAt,
	// 		})
	// 	}
	// }

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
