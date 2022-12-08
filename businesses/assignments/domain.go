package assignments

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID          string
	ModuleID    string
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

type Repository interface {
	// Create repository create new assignment
	Create(assignmentDomain *Domain) error

	// FindById repository find assignment by id
	FindById(assignmentId string) (*Domain, error)

	// FindByModuleId repository find assignment by moduleid
	FindByModuleId(moduleId string) (*Domain, error)

	// Update repository update assignment
	Update(assignmentId string, assignmentDomain *Domain) error

	// Delete repository delete assignment
	Delete(assignmentId string) error

	// DeleteByModuleId repository delete assignments y id module
	DeleteByModuleId(moduleId string) error
}

type Usecase interface {
	// Create usecase create new assignment
	Create(assignmentDomain *Domain) error

	// FindById usecase findfind assignment by id
	FindById(assignmentId string) (*Domain, error)

	// FindByModuleId usecase find assignment by moduleid
	FindByModuleId(moduleId string) (*Domain, error)

	// Update usecase update assignment
	Update(assignmentId string, assignmentDomain *Domain) error

	// Delete usecase delete assignment
	Delete(assignmentId string) error

	// DeleteByModuleId usecase delete assignments y id module
	DeleteByModuleId(moduleId string) error
}
