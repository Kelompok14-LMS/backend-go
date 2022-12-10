package response

import (
	"time"

	"github.com/Kelompok14-LMS/backend-go/businesses/materials"
)

type FindByIdMaterial struct {
	ID          string    `json:"id"`
	CourseId    string    `json:"course_id"`
	ModuleId    string    `json:"module_id"`
	Title       string    `json:"title"`
	URL         string    `json:"url"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func MaterialDetail(res *materials.Domain) *FindByIdMaterial {
	return &FindByIdMaterial{
		ID:          res.ID,
		CourseId:    res.Module.CourseId,
		ModuleId:    res.ModuleId,
		Title:       res.Title,
		URL:         res.URL,
		Description: res.Description,
		CreatedAt:   res.CreatedAt,
		UpdatedAt:   res.UpdatedAt,
	}
}
