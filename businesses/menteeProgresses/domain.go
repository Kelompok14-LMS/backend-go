package mentee_progresses

import (
	"time"

	"github.com/Kelompok14-LMS/backend-go/businesses/courses"
	"github.com/Kelompok14-LMS/backend-go/businesses/materials"
	"github.com/Kelompok14-LMS/backend-go/businesses/mentees"
)

type Domain struct {
	ID            string
	MenteeId      string
	CourseId      string
	MaterialId    string
	ProgressCount int64
	Mentee        mentees.Domain
	Course        courses.Domain
	Material      materials.Domain
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Repository interface {
	// Add repository add new progress
	Add(menteeProgressDomain *Domain) error

	// repository get mentee progress
	Count(menteeId string) ([]int64, error)
}

type Usecase interface {
	// Add usecase add new progress
	Add(menteeProgressDomain *Domain) error
}
