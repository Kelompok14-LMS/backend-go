package response

import (
	"time"

	"github.com/Kelompok14-LMS/backend-go/businesses/mentees"
)

type FindAllMentees struct {
	ID             string    `json:"id"`
	UserId         string    `json:"user_id"`
	Fullname       string    `json:"fullname"`
	Phone          string    `json:"phone"`
	Role           string    `json:"role"`
	BirthDate      time.Time `json:"birth_date"`
	Address        string    `json:"address"`
	ProfilePicture string    `json:"profile_picture"`
	Email          string    `json:"email"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func AllMentees(menteeDomain *mentees.Domain) FindAllMentees {
	return FindAllMentees{
		ID:             menteeDomain.ID,
		UserId:         menteeDomain.UserId,
		Fullname:       menteeDomain.Fullname,
		Phone:          menteeDomain.Phone,
		Role:           menteeDomain.Role,
		BirthDate:      menteeDomain.BirthDate,
		Address:        menteeDomain.Address,
		ProfilePicture: menteeDomain.ProfilePicture,
		Email:          menteeDomain.User.Email,
		CreatedAt:      menteeDomain.CreatedAt,
		UpdatedAt:      menteeDomain.UpdatedAt,
	}
}
