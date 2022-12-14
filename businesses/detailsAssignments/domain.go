package details_assignments

import (
	"time"

	menteeAssignments "github.com/Kelompok14-LMS/backend-go/businesses/menteeAssignments"
)

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

type Usecase interface {
	// detail course with modules and materials
	DetailAssignment(assignmentId string) (*Assignment, error)
}
