package request

import (
	"github.com/Kelompok14-LMS/backend-go/businesses/modules"
	"github.com/go-playground/validator/v10"
)

type UpdateModuleInput struct {
	CourseId string `json:"course_id" form:"course_id" validate:"required"`
	Title    string `json:"title" form:"title" validate:"required"`
}

func (req *UpdateModuleInput) ToDomain() *modules.Domain {
	return &modules.Domain{
		CourseId: req.CourseId,
		Title:    req.Title,
	}
}

func (req *UpdateModuleInput) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	if err != nil {
		return err
	}

	return nil
}
