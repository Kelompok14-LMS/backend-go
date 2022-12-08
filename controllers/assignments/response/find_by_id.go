package response

import (
	"time"

	"github.com/Kelompok14-LMS/backend-go/businesses/assignments"
)

type FindByIdAssignments struct {
	ID          string    `json:"id"`
	ModuleID    string    `json:"module_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func DetailAssignment(assignmentDomain *assignments.Domain) *FindByIdAssignments {
	return &FindByIdAssignments{
		ID:          assignmentDomain.ID,
		ModuleID:    assignmentDomain.ModuleID,
		Title:       assignmentDomain.Title,
		Description: assignmentDomain.Description,
		CreatedAt:   assignmentDomain.CreatedAt,
		UpdatedAt:   assignmentDomain.UpdatedAt,
	}
}
