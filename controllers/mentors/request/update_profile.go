package request

import (
	"time"

	"github.com/Kelompok14-LMS/backend-go/businesses/mentors"
	"github.com/go-playground/validator/v10"
)

type MentorUpdateProfile struct {
	ID             string `json:"id,omitempty" form:"id,omitempty" validate:"required"`
	UserID         string `json:"user_id,omitempty" form:"user_id,omitempty" validate:"required"`
	Fullname       string `json:"fullname" form:"fullname" validate:"required"`
	Email          string `json:"email" form:"fullname" validate:"required"`
	Phone          string `json:"phone" form:"phone" validate:"required"`
	Jobs           string `json:"jobs" form:"jobs" validate:"required"`
	Gender         string `json:"gender" form:"gender" validate:"required"`
	BirthPlace     string `json:"birth_place" form:"birth_place" validate:"required"`
	BirthDate      string `json:"birth_date" form:"birth_date" validate:"required"`
	ProfilePicture string `json:"profile_picture" form:"profile_picture" validate:"required"`
}

func (req *MentorUpdateProfile) ToDomain() *mentors.MentorUpdateProfile {

	birth, _ := time.Parse("2006-01-02", req.BirthDate)

	return &mentors.MentorUpdateProfile{

		ID:             req.ID,
		UserID:         req.UserID,
		Fullname:       req.Fullname,
		Email:          req.Email,
		Phone:          req.Phone,
		Jobs:           req.Jobs,
		Gender:         req.Gender,
		BirthPlace:     req.BirthPlace,
		BirthDate:      birth,
		ProfilePicture: req.ProfilePicture,
	}
}

func (req *MentorUpdateProfile) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	if err != nil {
		return err
	}

	return nil
}
