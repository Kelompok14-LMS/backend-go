package response

import (
	"time"

	mentee_courses "github.com/Kelompok14-LMS/backend-go/businesses/menteeCourses"
)

type FindMenteeCourses struct {
	ID             string    `json:"id"`
	CourseId       string    `json:"course_id"`
	Mentor         string    `json:"mentor"`
	Title          string    `json:"title"`
	Thumbnail      string    `json:"thumbnail"`
	Progress       int64     `json:"progress"`
	TotalMaterials int64     `json:"total_materials"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func MenteeCourses(menteeCourseDomain *mentee_courses.Domain) *FindMenteeCourses {
	return &FindMenteeCourses{
		ID:             menteeCourseDomain.ID,
		CourseId:       menteeCourseDomain.Course.ID,
		Mentor:         menteeCourseDomain.Course.Mentor.Fullname,
		Title:          menteeCourseDomain.Course.Title,
		Thumbnail:      menteeCourseDomain.Course.Thumbnail,
		Progress:       menteeCourseDomain.ProgressCount,
		TotalMaterials: menteeCourseDomain.TotalMaterials,
		CreatedAt:      menteeCourseDomain.CreatedAt,
		UpdatedAt:      menteeCourseDomain.UpdatedAt,
	}
}
