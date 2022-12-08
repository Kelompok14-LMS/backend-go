package assignments

import (
	"time"

	"github.com/Kelompok14-LMS/backend-go/businesses/assignments"
	"github.com/Kelompok14-LMS/backend-go/drivers/mysql/modules"
	"gorm.io/gorm"
)

type Assignment struct {
	ID          string `gorm:"primaryKey;size:200" json:"id"`
	ModuleID    string `gorm:"size:200" json:"module_id"`
	Module      modules.Module
	Title       string         `gorm:"size:225" json:"title"`
	Description string         `json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (rec *Assignment) ToDomain() *assignments.Domain {
	return &assignments.Domain{
		ID:          rec.ID,
		ModuleID:    rec.ModuleID,
		Title:       rec.Title,
		Description: rec.Description,
		CreatedAt:   rec.CreatedAt,
		UpdatedAt:   rec.UpdatedAt,
		DeletedAt:   rec.DeletedAt,
	}
}

func FromDomain(assignmentDomain *assignments.Domain) *Assignment {
	return &Assignment{
		ID:          assignmentDomain.ID,
		ModuleID:    assignmentDomain.ModuleID,
		Title:       assignmentDomain.Title,
		Description: assignmentDomain.Description,
		CreatedAt:   assignmentDomain.CreatedAt,
		UpdatedAt:   assignmentDomain.UpdatedAt,
		DeletedAt:   assignmentDomain.DeletedAt,
	}
}
