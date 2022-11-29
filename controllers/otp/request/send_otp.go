package request

import (
	otpDomain "github.com/Kelompok14-LMS/backend-go/businesses/otp"
	"github.com/go-playground/validator/v10"
)

type OTP struct {
	Key string `json:"email" form:"email" validate:"required,email"`
}

func (req *OTP) ToDomain() *otpDomain.Domain {
	return &otpDomain.Domain{
		Key: req.Key,
	}
}

func (req *OTP) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	if err != nil {
		return err
	}

	return nil
}
