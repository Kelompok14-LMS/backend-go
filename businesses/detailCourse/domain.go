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
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Usecase interface {
	// detail course with modules and materials
	DetailCourse(courseId string) (*Domain, error)
}
