package mentee_assignments

import (
	"errors"

	menteeAssignments "github.com/Kelompok14-LMS/backend-go/businesses/menteeAssignments"
	"gorm.io/gorm"
)

type assignmentMenteeRepository struct {
	conn *gorm.DB
}

func NewSQLRepository(conn *gorm.DB) menteeAssignments.Repository {
	return assignmentMenteeRepository{
		conn: conn,
	}
}

func (am assignmentMenteeRepository) Create(assignmentmenteeDomain *menteeAssignments.Domain) error {
	rec := FromDomain(assignmentmenteeDomain)

	err := am.conn.Model(&MenteeAssignment{}).Create(&rec).Error

	if err != nil {
		return err
	}

	return nil
}

func (am assignmentMenteeRepository) FindById(assignmentMenteeId string) (*menteeAssignments.Domain, error) {
	rec := MenteeAssignment{}

	err := am.conn.Model(&MenteeAssignment{}).Where("id = ?", assignmentMenteeId).Preload("Mentee").First(&rec).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

		return nil, err
	}

	return rec.ToDomain(), nil
}

func (am assignmentMenteeRepository) Update(assignmentMenteeId string, assignmentmenteeDomain *menteeAssignments.Domain) error {
	rec := FromDomain(assignmentmenteeDomain)

	err := am.conn.Model(&MenteeAssignment{}).Where("id = ?", assignmentMenteeId).Updates(&rec).Error

	if err != nil {
		return err
	}

	return nil
}

func (am assignmentMenteeRepository) Delete(assignmentMenteeId string) error {
	err := am.conn.Model(&MenteeAssignment{}).Where("id = ?", assignmentMenteeId).Delete(&MenteeAssignment{}).Error

	if err != nil {
		return err
	}

	return nil
}
