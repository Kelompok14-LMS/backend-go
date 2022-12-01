package response

import (
	"time"

	"github.com/Kelompok14-LMS/backend-go/businesses/mentors"
)

type FindMentorAll struct {
	ID             string    `json:"id,omitempty"`
	UserID         string    `json:"user_id"`
	FullName       string    `json:"fullname"`
	Email          string    `json:"email"`
	Phone          string    `json:"phone"`
	Role           string    `json:"role"`
	Jobs           string    `json:"jobs"`
	Gender         string    `json:"gender"`
	BirthPlace     string    `json:"birth_place"`
	BirthDate      string    `json:"birth_date"`
	ProfilePicture string    `json:"profile_picture"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func FromDomainAll(user *mentors.Domain) *FindMentorAll {
	return &FindMentorAll{
		ID:             user.ID,
		UserID:         user.UserId,
		FullName:       user.FullName,
		Email:          user.Email,
		Phone:          user.Phone,
		Role:           user.Role,
		Jobs:           user.Jobs,
		Gender:         user.Gender,
		BirthPlace:     user.BirthPlace,
		BirthDate:      user.BirthDate.String(),
		ProfilePicture: user.ProfilePicture,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
	}
}
