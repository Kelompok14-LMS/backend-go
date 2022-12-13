package detail_course

import "time"

type Domain struct {
	CourseId    string
	MentorId    string
	CategoryId  string
	Title       string
	Description string
	Thumbnail   string
	Category    string
	Mentor      string
	Modules     []Module
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Module struct {
	ModuleId    string
	CourseId    string
	Title       string
	Description string
	Materials   []Material
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Material struct {
	MaterialId  string
	ModuleId    string
	Title       string
	URL         string
	Description string
	Completed   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Usecase interface {
	// DetailCourse usecase detail course with modules and materials
	DetailCourse(courseId string) (*Domain, error)

	// DetailCourseEnrolled usecase detail course with module and material
	// for mentee who already enroll the course
	DetailCourseEnrolled(menteeId string, courseId string) (*Domain, error)
}
