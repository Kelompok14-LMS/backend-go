package mentee_assignments

import (
	"mime/multipart"
	"time"

	"github.com/Kelompok14-LMS/backend-go/businesses/assignments"
	"github.com/Kelompok14-LMS/backend-go/businesses/mentees"
)

type Domain struct {
	ID            string
	MenteeId      string
	AssignmentId  string
	Name          string
	AssignmentURL string
	PDFfile       *multipart.FileHeader
	Grade         int
	Mentee        mentees.Domain
	Assignment    assignments.Domain
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Repository interface {
	// Create repository create new assignment mentee
	Create(assignmentmenteeDomain *Domain) error

	// FindById repository find assignment mentee by id
	FindById(assignmentmenteeId string) (*Domain, error)

	// FindByAssignmentID repository find assignment mentee by assignment id
	FindByAssignmentId(assignmentId string) ([]Domain, error)

	// FindByMenteeId repository find assignment mentee by mentee id
	FindByMenteeId(menteeId string) ([]Domain, error)

	// Update repository update assignment  mentee
	Update(assignmentmenteeId string, assignmentmenteeDomain *Domain) error

	// Delete repository delete assignment mentee
	Delete(assignmentmenteeId string) error
}

type Usecase interface {
	// Create usecase create new assignment
	Create(assignmentDomain *Domain) error

	// FindById usecase findfind assignment by id
	FindById(assignmentId string) (*Domain, error)

	// FindByAssignmentID usecase  find assignment mentee by assignment id
	FindByAssignmentId(assignmentId string) ([]Domain, error)

	// FindByMenteeId rusecase find assignment mentee by mentee id
	FindByMenteeId(menteeId string) ([]Domain, error)

	// Update usecase update assignment
	Update(assignmentId string, assignmentDomain *Domain) error

	// Delete usecase delete assignment
	Delete(assignmentId string) error
}