package request

import (
	otpDomain "github.com/Kelompok14-LMS/backend-go/businesses/otp"
	"github.com/go-playground/validator/v10"
)

type CheckOTP struct {
	// the email
	Key string `json:"email" form:"email" validate:"required,email"`

	// the OTP
	Value string `json:"otp" form:"otp" validate:"required"`
}

func (req *CheckOTP) ToDomain() *otpDomain.Domain {
	return &otpDomain.Domain{
		Key:   req.Key,
		Value: req.Value,
	}
}

func (req *CheckOTP) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	if err != nil {
		return err
	}

	return err
}
