package response

import (
	"time"

	"github.com/Kelompok14-LMS/backend-go/businesses/modules"
)

type FindByIdModule struct {
	ID        string    `json:"id"`
	CourseId  string    `json:"course_id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func DetailModule(moduleDomain *modules.Domain) *FindByIdModule {
	return &FindByIdModule{
		ID:        moduleDomain.ID,
		CourseId:  moduleDomain.CourseId,
		Title:     moduleDomain.Title,
		CreatedAt: moduleDomain.CreatedAt,
		UpdatedAt: moduleDomain.UpdatedAt,
	}
}
