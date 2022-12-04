package modules

import (
	"time"

	"github.com/Kelompok14-LMS/backend-go/businesses/modules"
	"github.com/Kelompok14-LMS/backend-go/drivers/mysql/courses"
	"gorm.io/gorm"
)

type Module struct {
	ID        string `json:"id" gorm:"primaryKey;size:200"`
	CourseId  string `json:"course_id" gorm:"size:200"`
	Title     string `json:"title" gorm:"size:255"`
	Course    courses.Course
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (rec *Module) ToDomain() *modules.Domain {
	return &modules.Domain{
		ID:        rec.ID,
		CourseId:  rec.CourseId,
		Title:     rec.Title,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func FromDomain(moduleDomain *modules.Domain) *Module {
	return &Module{
		ID:        moduleDomain.ID,
		CourseId:  moduleDomain.CourseId,
		Title:     moduleDomain.Title,
		CreatedAt: moduleDomain.CreatedAt,
		UpdatedAt: moduleDomain.UpdatedAt,
	}
}
