package request

import (
	"github.com/Kelompok14-LMS/backend-go/businesses/mentors"
	"github.com/go-playground/validator/v10"
)

type ForgotPasswordInput struct {
	Email string `json:"email" form:"email" validate:"required,email"`
}

func (req *ForgotPasswordInput) ToDomain() *mentors.MentorForgotPassword {
	return &mentors.MentorForgotPassword{
		Email: req.Email,
	}
}

func (req *ForgotPasswordInput) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	if err != nil {
		return err
	}

	return nil
}
