package request

import (
	"mime/multipart"

	menteeAssignments "github.com/Kelompok14-LMS/backend-go/businesses/menteeAssignments"
	"github.com/go-playground/validator/v10"
)

type CreateMenteeAssignment struct {
	MenteeID     string                `json:"mentee_id" form:"mentee_id" validate:"required"`
	AssignmentID string                `json:"assignment_id" form:"assignment_id" validate:"required"`
	PDF          *multipart.FileHeader `json:"pdf" form:"pdf" validate:"required"`
}

func (req *CreateMenteeAssignment) ToDomain() *menteeAssignments.Domain {
	return &menteeAssignments.Domain{
		MenteeId:     req.MenteeID,
		AssignmentId: req.AssignmentID,
		PDFfile:      req.PDF,
	}
}

func (req *CreateMenteeAssignment) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	if err != nil {
		return err
	}

	return nil
}
