package mentors

import (
	"time"

	"github.com/Kelompok14-LMS/backend-go/businesses/mentors"
	"github.com/Kelompok14-LMS/backend-go/drivers/mysql/users"
)

type Mentor struct {
	ID             string     `gorm:"primaryKey;" json:"id"`
	UserId         string     `gorm:"size:200" json:"user_id"`
	FullName       string     `gorm:"size:255" json:"name"`
	Phone          string     `gorm:"size:15" json:"phone"`
	Role           string     `gorm:"size:50" json:"role"`
	Jobs           string     `json:"jobs"`
	Gender         string     `json:"gender"`
	BirthPlace     string     `json:"birth_place"`
	BirthDate      time.Time  `json:"birth_date"`
	Address        string     `json:"address"`
	ProfilePicture string     `json:"profile_picture"`
	User           users.User `json:"user"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}

func (rec *Mentor) ToDomain() *mentors.Domain {
	return &mentors.Domain{
		ID:             rec.ID,
		UserId:         rec.UserId,
		FullName:       rec.FullName,
		Email:          rec.User.Email,
		Phone:          rec.Phone,
		Role:           rec.Role,
		Jobs:           rec.Jobs,
		Gender:         rec.Gender,
		BirthPlace:     rec.BirthPlace,
		BirthDate:      rec.BirthDate,
		Address:        rec.Address,
		ProfilePicture: rec.ProfilePicture,
		CreatedAt:      rec.CreatedAt,
		UpdatedAt:      rec.UpdatedAt,
	}
}

func FromDomain(mentor *mentors.Domain) *Mentor {
	return &Mentor{
		ID:             mentor.ID,
		UserId:         mentor.UserId,
		FullName:       mentor.FullName,
		Phone:          mentor.Phone,
		Role:           mentor.Role,
		Jobs:           mentor.Jobs,
		Gender:         mentor.Gender,
		BirthPlace:     mentor.BirthPlace,
		BirthDate:      mentor.BirthDate,
		Address:        mentor.Address,
		ProfilePicture: mentor.ProfilePicture,
		CreatedAt:      mentor.CreatedAt,
		UpdatedAt:      mentor.UpdatedAt,
	}
}
