package request

import (
	menteeAssignments "github.com/Kelompok14-LMS/backend-go/businesses/menteeAssignments"
	"github.com/go-playground/validator/v10"
)

type CreateGrade struct {
	AssignmentID string `json:"assignment_id" form:"assignment_id" validate:"required"`
	Grade        int    `json:"grade" form:"grade" validate:"required"`
}

func (req *CreateGrade) ToDomain() *menteeAssignments.Domain {
	return &menteeAssignments.Domain{
		AssignmentId: req.AssignmentID,
		Grade:        req.Grade,
	}
}

func (req *CreateGrade) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	if err != nil {
		return err
	}

	return nil
}
