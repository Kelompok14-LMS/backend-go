package materials

import (
	"errors"

	"github.com/Kelompok14-LMS/backend-go/businesses/materials"
	"github.com/Kelompok14-LMS/backend-go/pkg"
	"gorm.io/gorm"
)

type materialRepository struct {
	conn *gorm.DB
}

func NewSQLRepository(conn *gorm.DB) materials.Repository {
	return materialRepository{
		conn: conn,
	}
}

func (mr materialRepository) Create(materialDomain *materials.Domain) error {
	rec := FromDomain(materialDomain)

	err := mr.conn.Model(&Material{}).Create(rec).Error

	if err != nil {
		return err
	}

	return nil
}

func (mr materialRepository) FindById(materialId string) (*materials.Domain, error) {
	rec := Material{}

	err := mr.conn.Model(&Material{}).Where("id = ?", materialId).First(&rec).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, pkg.ErrMaterialAssetNotFound
		}

		return nil, err
	}

	return rec.ToDomain(), nil
}

func (mr materialRepository) Update(materialId string, materialDomain *materials.Domain) error {
	rec := FromDomain(materialDomain)

	err := mr.conn.Model(&Material{}).Where("id = ?", materialId).Updates(&rec).Error

	if err != nil {
		return err
	}

	return nil
}

func (mr materialRepository) Delete(materialId string) error {
	err := mr.conn.Model(&Material{}).Where("id = ?", materialId).Delete(&Material{}).Error

	if err != nil {
		return err
	}

	return nil
}

func (mr materialRepository) Deletes(moduleId string) error {
	err := mr.conn.Model(&Material{}).Where("module_id = ?", moduleId).Delete(&Material{}).Error

	if err != nil {
		return err
	}

	return nil
}
