package assignments

import (
	"errors"

	"github.com/Kelompok14-LMS/backend-go/businesses/assignments"
	"github.com/Kelompok14-LMS/backend-go/pkg"
	"gorm.io/gorm"
)

type assignmentRepository struct {
	conn *gorm.DB
}

func NewSQLRepository(conn *gorm.DB) assignments.Repository {
	return assignmentRepository{
		conn: conn,
	}
}

func (ar assignmentRepository) Create(assignmentDomain *assignments.Domain) error {
	rec := FromDomain(assignmentDomain)

	err := ar.conn.Model(&Assignment{}).Create(&rec).Error

	if err != nil {
		return err
	}

	return nil
}

func (ar assignmentRepository) FindById(assignmentId string) (*assignments.Domain, error) {
	rec := Assignment{}

	err := ar.conn.Model(&Assignment{}).Where("id = ?", assignmentId).First(&rec).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, pkg.ErrAssignmentNotFound
		}

		return nil, err
	}

	return rec.ToDomain(), nil
}

func (ar assignmentRepository) FindByModuleId(moduleId string) (*assignments.Domain, error) {
	rec := Assignment{}

	err := ar.conn.Model(&Assignment{}).Where("module_id = ?", moduleId).First(&rec).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, pkg.ErrModuleNotFound
		}

		return nil, err
	}

	return rec.ToDomain(), nil
}

func (ar assignmentRepository) Update(assignmentId string, assignmentDomain *assignments.Domain) error {
	rec := FromDomain(assignmentDomain)

	err := ar.conn.Model(&Assignment{}).Where("id = ?", assignmentId).Updates(&rec).Error

	if err != nil {
		return err
	}

	return nil
}

func (ar assignmentRepository) Delete(assignmentId string) error {
	err := ar.conn.Model(&Assignment{}).Where("id = ?", assignmentId).Delete(&Assignment{}).Error

	if err != nil {
		return err
	}

	return nil
}

func (ar assignmentRepository) DeleteByModuleId(moduleId string) error {
	err := ar.conn.Model(&Assignment{}).Where("module_id = ?", moduleId).Delete(&Assignment{}).Error

	if err != nil {
		return err
	}

	return nil
}
