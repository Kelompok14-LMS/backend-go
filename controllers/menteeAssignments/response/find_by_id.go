package response

import (
	"time"

	menteeAssignments "github.com/Kelompok14-LMS/backend-go/businesses/menteeAssignments"
)

type AssignmentMentee struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	PDF       string    `json:"pdf"`
	Grade     int       `json:"grade"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(domain *menteeAssignments.Domain) AssignmentMentee {
	return AssignmentMentee{
		ID:        domain.ID,
		Name:      domain.Name,
		PDF:       domain.AssignmentURL,
		Grade:     domain.Grade,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
