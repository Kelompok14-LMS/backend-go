package mentee_assignments

import (
	"errors"

	menteeAssignments "github.com/Kelompok14-LMS/backend-go/businesses/menteeAssignments"
	"github.com/Kelompok14-LMS/backend-go/pkg"
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

func (am assignmentMenteeRepository) FindByMenteeId(menteeId string) ([]menteeAssignments.Domain, error) {
	rec := []MenteeAssignment{}

	err := am.conn.Model(&MenteeAssignment{}).Where("mentee_id = ?", menteeId).Preload("Mentee").Order("created_at DESC").Find(&rec).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

		return nil, err
	}
	assignmentMenteeDomain := []menteeAssignments.Domain{}

	for _, assignment := range rec {
		assignmentMenteeDomain = append(assignmentMenteeDomain, *assignment.ToDomain())
	}

	return assignmentMenteeDomain, nil

}

func (am assignmentMenteeRepository) FindByAssignmentId(assignmentId string) ([]menteeAssignments.Domain, error) {
	rec := []MenteeAssignment{}

	err := am.conn.Model(&MenteeAssignment{}).Where("assignment_id = ?", assignmentId).Preload("Mentee").Order("created_at DESC").Find(&rec).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, pkg.ErrAssignmentNotFound
		}

		return nil, err
	}

	assignmentDomain := []menteeAssignments.Domain{}

	for _, assignment := range rec {
		assignmentDomain = append(assignmentDomain, *assignment.ToDomain())
	}

	return assignmentDomain, nil
}

func (am assignmentMenteeRepository) FindByCourse(menteeId string, courseId string) (*menteeAssignments.Domain, error) {
	rec := MenteeAssignment{}

	err := am.conn.Model(&MenteeAssignment{}).Preload("Assignment").
		Joins("LEFT JOIN assignments ON assignments.id = mentee_assignments.assignment_id").
		Where("mentee_assignments.mentee_id = ? AND assignments.course_id = ?", menteeId, courseId).
		First(&rec).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, pkg.ErrAssignmentNotFound
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
	err := am.conn.Model(&MenteeAssignment{}).Unscoped().Where("id = ?", assignmentMenteeId).Delete(&MenteeAssignment{}).Error

	if err != nil {
		return err
	}

	return nil
}
