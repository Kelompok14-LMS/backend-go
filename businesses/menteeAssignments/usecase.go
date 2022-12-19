package mentee_assignments

import (
	"context"
	"math"

	"github.com/Kelompok14-LMS/backend-go/businesses/assignments"
	menteeCourses "github.com/Kelompok14-LMS/backend-go/businesses/menteeCourses"
	"github.com/Kelompok14-LMS/backend-go/businesses/mentees"
	"github.com/Kelompok14-LMS/backend-go/helper"
	"github.com/Kelompok14-LMS/backend-go/pkg"
	"github.com/Kelompok14-LMS/backend-go/utils"
	"github.com/google/uuid"
)

type assignmentMenteeUsecase struct {
	assignmentMenteeRepository Repository
	assignmentRepository       assignments.Repository
	menteeCourseRepository     menteeCourses.Repository
	menteeRepository           mentees.Repository
	storage                    *helper.StorageConfig
}

func NewMenteeAssignmentUsecase(assignmentMenteeRepository Repository,
	assignmentRepository assignments.Repository, menteeCourseRepository menteeCourses.Repository, menteeRepository mentees.Repository,
	storage *helper.StorageConfig) Usecase {
	return assignmentMenteeUsecase{
		assignmentMenteeRepository: assignmentMenteeRepository,
		assignmentRepository:       assignmentRepository,
		menteeCourseRepository:     menteeCourseRepository,
		menteeRepository:           menteeRepository,
		storage:                    storage,
	}
}

func (mu assignmentMenteeUsecase) Create(assignmentMenteeDomain *Domain) error {
	assignments, err := mu.assignmentRepository.FindById(assignmentMenteeDomain.AssignmentId)
	if err != nil {
		return err
	}

	if _, err := mu.menteeCourseRepository.CheckEnrollment(assignmentMenteeDomain.MenteeId, assignments.CourseId); err != nil {
		return pkg.ErrNoEnrolled
	}
	PDF, err := assignmentMenteeDomain.PDFfile.Open()

	if err != nil {
		return err
	}

	defer PDF.Close()

	filename, err := utils.GetFilename(assignmentMenteeDomain.PDFfile.Filename)

	if err != nil {
		return pkg.ErrUnsupportedAssignmentFile
	}

	ctx := context.Background()

	pdfUrl, err := mu.storage.UploadAsset(ctx, filename, PDF)

	if err != nil {
		return err
	}

	id := uuid.NewString()

	assignmentMentee := Domain{
		ID:            id,
		MenteeId:      assignmentMenteeDomain.MenteeId,
		AssignmentId:  assignmentMenteeDomain.AssignmentId,
		AssignmentURL: pdfUrl,
		Grade:         assignmentMenteeDomain.Grade,
	}

	err = mu.assignmentMenteeRepository.Create(&assignmentMentee)

	if err != nil {
		return err
	}

	return nil
}

func (mu assignmentMenteeUsecase) FindById(assignmentMenteeId string) (*Domain, error) {
	assignmentMentee, err := mu.assignmentMenteeRepository.FindById(assignmentMenteeId)

	if err != nil {
		return nil, err
	}

	return assignmentMentee, nil
}

func (mu assignmentMenteeUsecase) FindMenteeAssignmentEnrolled(menteeId string, assignmentId string) (*Domain, error) {
	if _, err := mu.menteeRepository.FindById(menteeId); err != nil {
		return nil, err
	}

	_, err := mu.assignmentRepository.FindById(assignmentId)
	if err != nil {
		return nil, err
	}

	assignmentMentee, err := mu.assignmentMenteeRepository.FindMenteeAssignmentEnrolled(menteeId, assignmentId)
	if err != nil {
		return nil, err
	}

	var completed bool

	if assignmentMentee == nil {
		completed = false
	} else {
		completed = true
	}

	menteeAssignment := Domain{
		ID:             assignmentMentee.ID,
		MenteeId:       menteeId,
		AssignmentId:   assignmentId,
		Name:           assignmentMentee.Name,
		ProfilePicture: assignmentMentee.ProfilePicture,
		AssignmentURL:  assignmentMentee.AssignmentURL,
		Grade:          assignmentMentee.Grade,
		Completed:      completed,
		CreatedAt:      assignmentMentee.CreatedAt,
		UpdatedAt:      assignmentMentee.UpdatedAt,
	}

	return &menteeAssignment, nil
}

func (mu assignmentMenteeUsecase) FindByMenteeId(menteeId string) ([]Domain, error) {
	assignmentMentee, err := mu.assignmentMenteeRepository.FindByMenteeId(menteeId)

	if err != nil {
		return nil, err
	}

	return assignmentMentee, nil
}

func (mu assignmentMenteeUsecase) FindByAssignmentId(assignmentId string, pagination pkg.Pagination) (*pkg.Pagination, error) {
	menteeAssignments, totalRows, err := mu.assignmentMenteeRepository.FindByAssignmentId(assignmentId, pagination.GetLimit(), pagination.GetOffset())

	if err != nil {
		return nil, err
	}

	pagination.Result = menteeAssignments
	pagination.TotalRows = totalRows
	pagination.TotalPages = int(math.Ceil(float64(totalRows) / float64(pagination.GetLimit())))

	return &pagination, nil
}

func (mu assignmentMenteeUsecase) Update(assignmentMenteeId string, assignmentMenteeDomain *Domain) error {
	if _, err := mu.assignmentRepository.FindById(assignmentMenteeDomain.AssignmentId); err != nil {
		return err
	}

	assignmentMentee, err := mu.assignmentMenteeRepository.FindById(assignmentMenteeId)

	if err != nil {
		return err
	}

	var pdfUrl string

	if assignmentMenteeDomain.PDFfile != nil {
		ctx := context.Background()

		err = mu.storage.DeleteObject(ctx, assignmentMentee.AssignmentURL)

		if err != nil {
			return err
		}

		PDF, err := assignmentMenteeDomain.PDFfile.Open()

		if err != nil {
			return err
		}

		defer PDF.Close()

		filename, err := utils.GetFilename(assignmentMenteeDomain.PDFfile.Filename)

		if err != nil {
			return pkg.ErrUnsupportedAssignmentFile
		}

		pdfUrl, err = mu.storage.UploadAsset(ctx, filename, PDF)

		if err != nil {
			return err
		}

	}

	updatedMenteeAssignment := Domain{
		MenteeId:      assignmentMentee.MenteeId,
		AssignmentId:  assignmentMentee.AssignmentId,
		AssignmentURL: pdfUrl,
		Grade:         assignmentMenteeDomain.Grade,
	}

	err = mu.assignmentMenteeRepository.Update(assignmentMenteeId, &updatedMenteeAssignment)

	if err != nil {
		return err
	}

	return nil
}

func (mu assignmentMenteeUsecase) Delete(assignmentMenteeId string) error {
	assignmentMentee, err := mu.assignmentMenteeRepository.FindById(assignmentMenteeId)

	if err != nil {
		return err
	}

	ctx := context.Background()

	if err := mu.storage.DeleteObject(ctx, assignmentMentee.AssignmentURL); err != nil {
		return err
	}

	if err := mu.assignmentMenteeRepository.Delete(assignmentMenteeId); err != nil {
		return err
	}

	return nil
}
