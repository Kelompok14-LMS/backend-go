package mentees

import (
	"errors"

	"github.com/Kelompok14-LMS/backend-go/businesses/mentees"
	"github.com/Kelompok14-LMS/backend-go/pkg"
	"gorm.io/gorm"
)

type menteeRepository struct {
	conn *gorm.DB
}

func NewSQLRepository(conn *gorm.DB) mentees.Repository {
	return menteeRepository{
		conn: conn,
	}
}

func (mr menteeRepository) Create(menteeDomain *mentees.Domain) error {
	rec := FromDomain(menteeDomain)

	err := mr.conn.Model(&Mentee{}).Omit("birth_date", "address", "profile_picture").Create(&rec).Error

	if err != nil {
		return err
	}

	return nil
}

func (mr menteeRepository) FindAll() (*[]mentees.Domain, error) {
	var rec []Mentee

	err := mr.conn.Model(&Mentee{}).Find(&rec).Error

	if err != nil {
		return nil, err
	}

	menteeDomain := []mentees.Domain{}

	for _, mentee := range rec {
		menteeDomain = append(menteeDomain, *mentee.ToDomain())
	}

	return &menteeDomain, nil
}

func (mr menteeRepository) FindById(id string) (*mentees.Domain, error) {
	rec := Mentee{}

	err := mr.conn.Model(&Mentee{}).Where("id = ?", id).First(&rec).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, pkg.ErrMenteeNotFound
		}

		return nil, err
	}

	return rec.ToDomain(), nil
}

func (mr menteeRepository) FindByIdUser(userId string) (*mentees.Domain, error) {
	rec := &Mentee{}

	err := mr.conn.Model(&Mentee{}).Where("user_id = ?", userId).First(&rec).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, pkg.ErrMenteeNotFound
		}

		return nil, err
	}

	return rec.ToDomain(), nil
}

func (mr menteeRepository) Update(id string, menteeDomain *mentees.Domain) error {
	rec := FromDomain(menteeDomain)

	err := mr.conn.Model(&Mentee{}).Where("id = ?", id).Updates(&rec).Error

	if err != nil {
		return err
	}

	return nil
}
