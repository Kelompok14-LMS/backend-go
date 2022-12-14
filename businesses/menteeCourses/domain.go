package mentee_courses

import (
	"time"

	"github.com/Kelompok14-LMS/backend-go/businesses/courses"
	"github.com/Kelompok14-LMS/backend-go/businesses/mentees"
)

type Domain struct {
	ID             string
	MenteeId       string
	CourseId       string
	Status         string
	Mentee         mentees.Domain
	Course         courses.Domain
	ProgressCount  int64
	TotalMaterials int64
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Repository interface {
	// Enroll repository enroll a course
	Enroll(menteeCourseDomain *Domain) error

	// FindCoursesByMentee repository find by mentee
	FindCoursesByMentee(menteeId string, title string, status string) (*[]Domain, error)

	// CheckEnrollment repository check enrollment course mentee
	CheckEnrollment(menteeId string, courseId string) (*Domain, error)

	// DeleteEnrolledCourse delete enrolled course mentee
	DeleteEnrolledCourse(menteeId string, courseId string) error
}

type Usecase interface {
	// Enroll usecase Enroll usecase mentee enroll course
	Enroll(menteeCourseDomain *Domain) error

	// FindMenteeCourses usecase find all enrollment courses with title and status
	FindMenteeCourses(menteeId string, title string, status string) (*[]Domain, error)

	// CheckEnrollment usecase check enrollment course mentee
	CheckEnrollment(menteeId string, courseId string) (bool, error)
}
