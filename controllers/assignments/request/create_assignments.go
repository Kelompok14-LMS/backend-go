package request

import (
	"github.com/Kelompok14-LMS/backend-go/businesses/assignments"
	"github.com/go-playground/validator/v10"
)

type CreateAssignment struct {
	ModuleID    string `json:"module_id" form:"module_id" validate:"required"`
	Title       string `json:"title" form:"title" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
	PDF         string `json:"pdf" form:"pdf" validate:"required"`
}

func (req *CreateAssignment) ToDomain() *assignments.Domain {
	return &assignments.Domain{
		ModuleID:    req.ModuleID,
		Title:       req.Title,
		Description: req.Description,
		PDF:         req.PDF,
	}
}

func (req *CreateAssignment) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	if err != nil {
		return err
	}

	return nil
}
