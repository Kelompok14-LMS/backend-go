package courses

import (
	"errors"

	"github.com/Kelompok14-LMS/backend-go/businesses/courses"
	"github.com/Kelompok14-LMS/backend-go/pkg"
	"gorm.io/gorm"
)

type courseRepository struct {
	conn *gorm.DB
}

func NewSQLRepository(conn *gorm.DB) courses.Repository {
	return courseRepository{
		conn: conn,
	}
}

func (cr courseRepository) Create(courseDomain *courses.Domain) error {
	rec := FromDomain(courseDomain)

	err := cr.conn.Model(&Course{}).Create(&rec).Error

	if err != nil {
		return err
	}

	return nil
}

func (cr courseRepository) FindAll(keyword string) (*[]courses.Domain, error) {
	var rec []Course

	err := cr.conn.Model(&Course{}).Preload("Category").Preload("Mentor").Joins("INNER JOIN categories ON categories.id = courses.category_id").Where("courses.title LIKE ? OR categories.name LIKE ?", "%"+keyword+"%", "%"+keyword+"%").Find(&rec).Error

	if err != nil {
		return nil, err
	}

	var coursesDomain []courses.Domain

	for _, course := range rec {
		coursesDomain = append(coursesDomain, *course.ToDomain())
	}

	return &coursesDomain, nil
}

func (cr courseRepository) FindById(id string) (*courses.Domain, error) {
	rec := Course{}

	err := cr.conn.Model(&Course{}).Where("id = ?", id).Preload("Category").Preload("Mentor").First(&rec).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, pkg.ErrCourseNotFound
		}

		return nil, err
	}

	return rec.ToDomain(), nil
}

func (cr courseRepository) FindByCategory(categoryId string) (*[]courses.Domain, error) {
	var rec []Course

	err := cr.conn.Model(&Course{}).Where("category_id = ?", categoryId).Preload("Category").Preload("Mentor").Find(&rec).Error

	if err != nil {
		return nil, err
	}

	var coursesDomain []courses.Domain

	for _, course := range rec {
		coursesDomain = append(coursesDomain, *course.ToDomain())
	}

	return &coursesDomain, nil
}

func (cr courseRepository) Update(id string, courseDomain *courses.Domain) error {
	rec := FromDomain(courseDomain)

	err := cr.conn.Model(&Course{}).Where("id = ?", id).Updates(&rec).Error

	if err != nil {
		return err
	}

	return nil
}

func (cr courseRepository) Delete(id string) error {
	err := cr.conn.Model(&Course{}).Where("id = ?", id).Delete(&Course{}).Error

	if err != nil {
		return err
	}

	return nil
}
