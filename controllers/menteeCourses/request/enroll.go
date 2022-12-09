package request

import (
	menteeCourses "github.com/Kelompok14-LMS/backend-go/businesses/menteeCourses"
	"github.com/go-playground/validator/v10"
)

type EnrollCourse struct {
	MenteeId string `json:"mentee_id" form:"mentee_id" validate:"required"`
	CourseId string `json:"course_id" form:"course_id" validate:"required"`
}

func (req *EnrollCourse) ToDomain() *menteeCourses.Domain {
	return &menteeCourses.Domain{
		MenteeId: req.MenteeId,
		CourseId: req.CourseId,
	}
}

func (req *EnrollCourse) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	if err != nil {
		return err
	}

	return nil
}
