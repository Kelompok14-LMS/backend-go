package assignments

import (
	"context"
	"strings"

	"github.com/Kelompok14-LMS/backend-go/businesses/modules"
	"github.com/Kelompok14-LMS/backend-go/helper"
	"github.com/Kelompok14-LMS/backend-go/pkg"
	"github.com/gabriel-vasile/mimetype"
	"github.com/google/uuid"
)

type assignmentUsecase struct {
	assignmentRepository Repository
	moduleRepository     modules.Repository
	storage              *helper.StorageConfig
}

func NewAssignmentUsecase(assignmentRepository Repository, moduleRepository modules.Repository, storage *helper.StorageConfig) Usecase {
	return assignmentUsecase{
		assignmentRepository: assignmentRepository,
		moduleRepository:     moduleRepository,
		storage:              storage,
	}
}

func (au assignmentUsecase) Create(assignmentDomain *Domain) error {
	if _, err := au.moduleRepository.FindById(assignmentDomain.ModuleID); err != nil {
		return err
	}

	PDF, err := assignmentDomain.PDFfile.Open()

	if err != nil {
		return err
	}

	defer PDF.Close()

	PDFname := strings.Replace(assignmentDomain.PDFfile.Filename, " ", "", -1)
	PDFmime, err := mimetype.DetectReader(PDF)

	if err != nil {
		return err
	}

	filetype := PDFmime.String()

	// only .mp4 and .mkv format are acceptable
	if filetype != "application/pdf" {
		return pkg.ErrUnsupportedAssignmentFile
	}

	ctx := context.Background()

	pdfUrl, err := au.storage.UploadAsset(ctx, PDFname, PDF)

	if err != nil {
		return err
	}

	id := uuid.NewString()

	assignment := Domain{
		ID:          id,
		ModuleID:    assignmentDomain.ModuleID,
		Title:       assignmentDomain.Title,
		Description: assignmentDomain.Description,
		PDFurl:      pdfUrl,
	}

	err = au.assignmentRepository.Create(&assignment)

	if err != nil {
		return err
	}

	return nil
}

func (au assignmentUsecase) FindById(assignmentId string) (*Domain, error) {
	assignment, err := au.assignmentRepository.FindById(assignmentId)

	if err != nil {
		return nil, err
	}

	return assignment, nil
}

func (au assignmentUsecase) FindByModuleId(moduleId string) (*Domain, error) {
	assignment, err := au.assignmentRepository.FindByModuleId(moduleId)

	if err != nil {
		return nil, err
	}

	return assignment, nil
}

func (au assignmentUsecase) Update(assignmentId string, assignmentDomain *Domain) error {
	if _, err := au.moduleRepository.FindById(assignmentDomain.ModuleID); err != nil {
		return err
	}

	assignment, err := au.assignmentRepository.FindById(assignmentId)

	if err != nil {
		return err
	}

	var pdfUrl string

	if assignmentDomain.PDFfile != nil {
		ctx := context.Background()

		err = au.storage.DeleteObject(ctx, assignment.PDFurl)

		if err != nil {
			return err
		}

		PDF, err := assignmentDomain.PDFfile.Open()

		if err != nil {
			return err
		}

		PDFname := strings.Replace(assignmentDomain.PDFfile.Filename, " ", "", -1)
		PDFmime, err := mimetype.DetectReader(PDF)

		if err != nil {
			return err
		}

		filetype := PDFmime.String()

		// only .mp4 and .mkv format are acceptable
		if filetype != "application/pdf" {
			return pkg.ErrUnsupportedAssignmentFile
		}

		pdfUrl, err = au.storage.UploadAsset(ctx, PDFname, PDF)

		if err != nil {
			return err
		}

	}

	updatedAssignment := Domain{
		ModuleID:    assignmentDomain.ModuleID,
		Title:       assignment.Title,
		Description: assignment.Description,
		PDFurl:      pdfUrl,
	}

	err = au.assignmentRepository.Update(assignmentId, &updatedAssignment)

	if err != nil {
		return err
	}

	return nil
}

func (au assignmentUsecase) Delete(assignmentId string) error {
	if _, err := au.assignmentRepository.FindById(assignmentId); err != nil {
		return err
	}

	err := au.assignmentRepository.Delete(assignmentId)

	if err != nil {
		return err
	}

	return nil
}
