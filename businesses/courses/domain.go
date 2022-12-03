package courses

import (
	"mime/multipart"
	"time"

	"github.com/Kelompok14-LMS/backend-go/businesses/categories"
	"github.com/Kelompok14-LMS/backend-go/businesses/mentors"
	"gorm.io/gorm"
)

type Domain struct {
	ID                  string
	MentorId            string
	CategoryId          string
	Title               string
	Description         string
	Thumbnail           string
	ThumbnailFileHeader *multipart.FileHeader
	Category            categories.Domain
	Mentor              mentors.Domain
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt
}

type Repository interface {
	// Create repository create course
	Create(courseDomain *Domain) error

	// FindAll repository find all courses by course title and category
	FindAll(keyword string) (*[]Domain, error)

	// FindById repository find course by id
	FindById(id string) (*Domain, error)

	// FindByCategory repository find by id category
	FindByCategory(categoryId string) (*[]Domain, error)

	// Update repository update course
	Update(id string, courseDomain *Domain) error

	// Delete repository delete course
	Delete(id string) error
}

type Usecase interface {
	// Create usecase create new course
	Create(courseDomain *Domain) error

	// FindAll usecase find all courses by course title and category
	FindAll(keyword string) (*[]Domain, error)

	// FindById usecase find by id
	FindById(id string) (*Domain, error)

	// FindByCategory usecase find by category id
	FindByCategory(categoryId string) (*[]Domain, error)

	// Update usecase update
	Update(id string, courseDomain *Domain) error

	// Delete usecase delete
	Delete(id string) error
}
