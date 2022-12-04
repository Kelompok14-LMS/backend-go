package modules

import (
	"time"

	"github.com/Kelompok14-LMS/backend-go/drivers/mysql/courses"
	"gorm.io/gorm"
)

type Domain struct {
	ID        string
	CourseId  string
	Title     string
	Course    courses.Course
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type Repository interface {
	// Create repository create new module
	Create(moduleDomain *Domain) error

	// FindById repository find module by id
	FindById(moduleId string) (*Domain, error)

	// Update repository update module
	Update(moduleId string, moduleDomain *Domain) error

	// Delete repository delete module
	Delete(moduleId string) error
}

type Usecase interface {
	// Create usecase create new module
	Create(moduleDomain *Domain) error

	// FindById usecase find module by id
	FindById(moduleId string) (*Domain, error)

	// Update usecase update module
	Update(moduleId string, moduleDomain *Domain) error

	// Delete usecase delete module
	Delete(moduleId string) error
}
