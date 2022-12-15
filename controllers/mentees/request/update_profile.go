package request

import (
	"mime/multipart"
	"time"

	"github.com/Kelompok14-LMS/backend-go/businesses/mentees"
	"github.com/go-playground/validator/v10"
)

type MenteeUpdateProfile struct {
	Fullname           string                `json:"fullname" form:"fullname"`
	Phone              string                `json:"phone" form:"phone"`
	BirthDate          string                `json:"birth_date" form:"birth_date"`
	ProfilePictureFile *multipart.FileHeader `json:"profile_picture" form:"profile_picture"`
}

func (req *MenteeUpdateProfile) ToDomain() *mentees.Domain {
	format := "2006-01-02"
	birth, _ := time.Parse(format, req.BirthDate)
	return &mentees.Domain{
		Fullname:           req.Fullname,
		Phone:              req.Phone,
		BirthDate:          birth,
		ProfilePictureFile: req.ProfilePictureFile,
	}
}

func (req *MenteeUpdateProfile) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	if err != nil {
		return err
	}

	return nil
}
