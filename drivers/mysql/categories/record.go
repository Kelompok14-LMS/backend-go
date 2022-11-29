package categories

import (
	"time"

	"github.com/Kelompok14-LMS/backend-go/businesses/categories"
)

type Category struct {
	ID        string    `gorm:"primaryKey;size:200" json:"id"`
	Name      string    `gorm:"size:255" json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(categoryDomain *categories.Domain) *Category {
	return &Category{
		ID:        categoryDomain.ID,
		Name:      categoryDomain.Name,
		CreatedAt: categoryDomain.CreatedAt,
		UpdatedAt: categoryDomain.UpdatedAt,
	}
}

func (rec *Category) ToDomain() *categories.Domain {
	return &categories.Domain{
		ID:        rec.ID,
		Name:      rec.Name,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}
