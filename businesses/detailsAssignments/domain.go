package details_assignments

import (
	"time"

	menteeAssignments "github.com/Kelompok14-LMS/backend-go/businesses/menteeAssignments"
)

// type Domain struct {
// 	CourseId    string
// 	MentorId    string
// 	CategoryId  string
// 	Title       string
// 	Description string
// 	Thumbnail   string
// 	Category    string
// 	Mentor      string
// 	Assignments []Assignment
// 	CreatedAt   time.Time
// 	UpdatedAt   time.Time
// }

type Assignment struct {
	AssignmentID     string
	CourseId         string
	NameCourse       string
	Title            string
	Description      string
	AssignmentMentee []menteeAssignments.Domain
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

// type AssignmentMentee struct {
// 	AssignmentMenteeID string
// 	MenteeId           string
// 	AssignmentId       string
// 	Name               string
// 	AssignmentURL      string
// 	Grade              int
// 	CreatedAt          time.Time
// 	UpdatedAt          time.Time
// }

type Usecase interface {
	// detail course with modules and materials
	DetailAssignment(assignmentId string) (*Assignment, error)
}
