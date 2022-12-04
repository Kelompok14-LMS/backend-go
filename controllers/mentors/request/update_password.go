package request

import (
	"github.com/Kelompok14-LMS/backend-go/businesses/mentors"
	"github.com/go-playground/validator/v10"
)

type MentorUpdatePassword struct {
	UserID      string `json:"user_id,omitempty" form:"user_id,omitempty" validate:"required"`
	OldPassword string `json:"old_password" form:"old_password" validate:"required"`
	NewPassword string `json:"new_password" form:"new_password" validate:"required"`
}

func (req *MentorUpdatePassword) ToDomain() *mentors.MentorUpdatePassword {
	return &mentors.MentorUpdatePassword{
		UserID:      req.UserID,
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	}
}

func (req *MentorUpdatePassword) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	if err != nil {
		return err
	}

	return nil
}
