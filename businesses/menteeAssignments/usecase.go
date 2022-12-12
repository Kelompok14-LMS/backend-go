package mentee_assignments

import (
	"context"

	"github.com/Kelompok14-LMS/backend-go/businesses/assignments"
	"github.com/Kelompok14-LMS/backend-go/helper"
	"github.com/Kelompok14-LMS/backend-go/pkg"
	"github.com/Kelompok14-LMS/backend-go/utils"
	"github.com/google/uuid"
)

type assignmentMenteeUsecase struct {
	assignmentMenteeRepository Repository
	assignmentRepository       assignments.Repository
	storage                    *helper.StorageConfig
}

func NewMenteeAssignmentUsecase(assignmentMenteeRepository Repository,
	assignmentRepository assignments.Repository,
	storage *helper.StorageConfig) Usecase {
	return assignmentMenteeUsecase{
		assignmentMenteeRepository: assignmentMenteeRepository,
		assignmentRepository:       assignmentRepository,
		storage:                    storage,
	}
}

func (mu assignmentMenteeUsecase) Create(assignmentMenteeDomain *Domain) error {
	if _, err := mu.assignmentRepository.FindById(assignmentMenteeDomain.AssignmentId); err != nil {
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
	if _, err := mu.assignmentMenteeRepository.FindById(assignmentMenteeId); err != nil {
		return err
	}

	err := mu.assignmentMenteeRepository.Delete(assignmentMenteeId)

	if err != nil {
		return err
	}

	return nil
}
