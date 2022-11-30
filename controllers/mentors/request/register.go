package request

import (
	"github.com/Kelompok14-LMS/backend-go/businesses/mentors"
	"github.com/go-playground/validator/v10"
)

type MentorRegisterInput struct {
	FullName string `json:"fullname" form:"fullname" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
}

func (req *MentorRegisterInput) ToDomain() *mentors.MentorRegister {
	return &mentors.MentorRegister{
		FullName: req.FullName,
		Email:    req.Email,
		Password: req.Password,
	}
}

func (req *MentorRegisterInput) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	if err != nil {
		return err
	}

	return nil
}
