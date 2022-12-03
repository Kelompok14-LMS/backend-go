package response

import (
	"time"

	"github.com/Kelompok14-LMS/backend-go/businesses/courses"
)

type FindByIdCourse struct {
	CourseId    string    `json:"course_id"`
	Mentor      string    `json:"mentor"`
	Category    string    `json:"category"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Thumbnail   string    `json:"thumbnail"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func DetailCourse(res *courses.Domain) FindByIdCourse {
	return FindByIdCourse{
		CourseId:    res.ID,
		Mentor:      res.Mentor.FullName,
		Category:    res.Category.Name,
		Title:       res.Title,
		Description: res.Description,
		Thumbnail:   res.Thumbnail,
		CreatedAt:   res.CreatedAt,
		UpdatedAt:   res.UpdatedAt,
	}
}
