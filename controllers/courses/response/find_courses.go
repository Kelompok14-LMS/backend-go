package response

import (
	"time"

	"github.com/Kelompok14-LMS/backend-go/businesses/courses"
)

type FindCourses struct {
	CourseId    string    `json:"course_id"`
	Mentor      string    `json:"mentor"`
	Category    string    `json:"category"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Thumbnail   string    `json:"thumbnail"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func AllCourses(res *courses.Domain) FindCourses {
	return FindCourses{
		CourseId:    res.ID,
		Mentor:      res.Mentor.Fullname,
		Category:    res.Category.Name,
		Title:       res.Title,
		Description: res.Description,
		Thumbnail:   res.Thumbnail,
		CreatedAt:   res.CreatedAt,
		UpdatedAt:   res.UpdatedAt,
	}
}
