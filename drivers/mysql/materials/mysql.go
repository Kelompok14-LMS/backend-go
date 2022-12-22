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

	err := mr.conn.Model(&Material{}).Preload("Module").
		Joins("INNER JOIN modules ON modules.id = materials.module_id").
		Where("materials.id = ?", materialId).First(&rec).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, pkg.ErrMaterialAssetNotFound
		}

		return nil, err
	}

	return rec.ToDomain(), nil
}

func (mr materialRepository) FindByModule(moduleIds []string) ([]materials.Domain, error) {
	rec := []Material{}

	err := mr.conn.Model(&Material{}).Preload("Module").
		Joins("INNER JOIN modules ON modules.id = materials.module_id").
		Where("materials.module_id IN ?", moduleIds).
		Order("materials.created_at ASC").
		Find(&rec).Error

	if err != nil {
		return nil, err
	}

	materialsDomain := []materials.Domain{}

	for _, material := range rec {
		materialsDomain = append(materialsDomain, *material.ToDomain())
	}

	return materialsDomain, nil
}

func (mr materialRepository) FindByCourse(courseIds []string, title string, status string) ([]materials.Domain, []int64, error) {
	totalRows := []int64{}

	_ = mr.conn.Model(&Material{}).Select("COUNT(DISTINCT materials.id)").
		Joins("LEFT JOIN modules ON modules.id = materials.module_id").
		Joins("LEFT JOIN courses ON courses.id = modules.course_id").
		Joins("LEFT JOIN mentee_courses ON courses.id = mentee_courses.course_id").
		Where("modules.course_id IN ? AND courses.title LIKE ? AND mentee_courses.status LIKE ? AND courses.deleted_at IS NULL AND modules.deleted_at IS NULL AND materials.deleted_at IS NULL", courseIds, "%"+title+"%", "%"+status+"%").
		Group("modules.course_id").
		Find(&totalRows).Error

	rec := []Material{}

	err := mr.conn.Model(&Material{}).Preload("Module").
		Joins("LEFT JOIN modules ON modules.id = materials.module_id").
		Joins("LEFT JOIN courses ON courses.id = modules.course_id").
		Joins("LEFT JOIN mentee_courses ON courses.id = mentee_courses.course_id").
		Where("modules.course_id IN ? AND courses.title LIKE ? AND mentee_courses.status LIKE ? AND courses.deleted_at IS NULL AND modules.deleted_at IS NULL AND materials.deleted_at IS NULL", courseIds, "%"+title+"%", "%"+status+"%").
		Group("modules.course_id").
		Find(&rec).Error

	if err != nil {
		return nil, nil, err
	}

	var materialDomain []materials.Domain

	for _, material := range rec {
		materialDomain = append(materialDomain, *material.ToDomain())
	}

	return materialDomain, totalRows, nil
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
