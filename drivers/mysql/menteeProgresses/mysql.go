package mentee_progresses

import (
	menteeProgresses "github.com/Kelompok14-LMS/backend-go/businesses/menteeProgresses"
	"gorm.io/gorm"
)

type menteeProgressRepository struct {
	conn *gorm.DB
}

func NewSQLRepository(conn *gorm.DB) menteeProgresses.Repository {
	return menteeProgressRepository{
		conn: conn,
	}
}

func (m menteeProgressRepository) Add(menteeProgressDomain *menteeProgresses.Domain) error {
	rec := FromDomain(menteeProgressDomain)

	err := m.conn.Model(&MenteeProgress{}).Create(&rec).Error

	if err != nil {
		return err
	}

	return nil
}

func (m menteeProgressRepository) Count(menteeId string) ([]int64, error) {
	rec := []int64{}

	err := m.conn.Model(&MenteeProgress{}).Select("COUNT(material_id)").
		Where("mentee_id = ?", menteeId).Group("mentee_id").Group("course_id").Order("course_id ASC").
		Find(&rec).Error

	if err != nil {
		return nil, err
	}

	return rec, nil
}
