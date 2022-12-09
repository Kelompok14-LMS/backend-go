package assignments

import (
	"github.com/Kelompok14-LMS/backend-go/businesses/modules"
	"github.com/Kelompok14-LMS/backend-go/pkg"
	"github.com/google/uuid"
)

type assignmentUsecase struct {
	assignmentRepository Repository
	moduleRepository     modules.Repository
}

func NewAssignmentUsecase(assignmentRepository Repository, moduleRepository modules.Repository) Usecase {
	return assignmentUsecase{
		assignmentRepository: assignmentRepository,
		moduleRepository:     moduleRepository,
	}
}

func (au assignmentUsecase) Create(assignmentDomain *Domain) error {
	if _, err := au.moduleRepository.FindById(assignmentDomain.ModuleID); err != nil {
		return err
	}

	id := uuid.NewString()

	assignment := Domain{
		ID:          id,
		ModuleID:    assignmentDomain.ModuleID,
		Title:       assignmentDomain.Title,
		Description: assignmentDomain.Description,
	}

	err := au.assignmentRepository.Create(&assignment)

	if err != nil {
		return err
	}

	return nil
}

func (au assignmentUsecase) FindById(assignmentId string) (*Domain, error) {
	assignment, err := au.assignmentRepository.FindById(assignmentId)

	if err != nil {
		return nil, pkg.ErrAssignmentNotFound
	}

	return assignment, nil
}

func (au assignmentUsecase) FindByModuleId(moduleId string) (*Domain, error) {
	assignment, err := au.assignmentRepository.FindByModuleId(moduleId)

	if err != nil {
		return nil, pkg.ErrAssignmentNotFound
	}

	return assignment, nil
}

func (au assignmentUsecase) Update(assignmentId string, assignmentDomain *Domain) error {
	if _, err := au.moduleRepository.FindById(assignmentDomain.ModuleID); err != nil {
		return err
	}

	_, err := au.assignmentRepository.FindById(assignmentId)

	if err != nil {
		return pkg.ErrAssignmentNotFound
	}

	updatedAssignment := Domain{
		ModuleID:    assignmentDomain.ModuleID,
		Title:       assignmentDomain.Title,
		Description: assignmentDomain.Description,
	}

	err = au.assignmentRepository.Update(assignmentId, &updatedAssignment)

	if err != nil {
		return pkg.ErrAssignmentNotFound
	}

	return nil
}

func (au assignmentUsecase) Delete(assignmentId string) error {
	if _, err := au.assignmentRepository.FindById(assignmentId); err != nil {
		return err
	}

	err := au.assignmentRepository.Delete(assignmentId)

	if err != nil {
		return pkg.ErrAssignmentNotFound
	}

	return nil
}

func (au assignmentUsecase) DeleteByModuleId(moduleId string) error {
	if _, err := au.assignmentRepository.FindById(moduleId); err != nil {
		return err
	}

	err := au.assignmentRepository.DeleteByModuleId(moduleId)

	if err != nil {
		return pkg.ErrAssignmentNotFound
	}

	return nil
}
