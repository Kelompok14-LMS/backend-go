package modules

import (
	"github.com/Kelompok14-LMS/backend-go/businesses/courses"
	"github.com/google/uuid"
)

type moduleUsecase struct {
	moduleRepository Repository
	courseRepository courses.Repository
}

func NewModuleUsecase(moduleRepository Repository, courseRepository courses.Repository) Usecase {
	return moduleUsecase{
		moduleRepository: moduleRepository,
		courseRepository: courseRepository,
	}
}

func (mu moduleUsecase) Create(moduleDomain *Domain) error {
	if _, err := mu.courseRepository.FindById(moduleDomain.CourseId); err != nil {
		return err
	}

	id := uuid.NewString()

	module := Domain{
		ID:       id,
		CourseId: moduleDomain.CourseId,
		Title:    moduleDomain.Title,
	}

	err := mu.moduleRepository.Create(&module)

	if err != nil {
		return err
	}

	return nil
}

func (mu moduleUsecase) FindById(moduleId string) (*Domain, error) {
	module, err := mu.moduleRepository.FindById(moduleId)

	if err != nil {
		return nil, err
	}

	return module, nil
}

func (mu moduleUsecase) Update(moduleId string, moduleDomain *Domain) error {
	if _, err := mu.courseRepository.FindById(moduleDomain.CourseId); err != nil {
		return err
	}

	if _, err := mu.moduleRepository.FindById(moduleId); err != nil {
		return err
	}

	err := mu.moduleRepository.Update(moduleId, moduleDomain)

	if err != nil {
		return err
	}

	return nil
}

func (mu moduleUsecase) Delete(moduleId string) error {
	if _, err := mu.moduleRepository.FindById(moduleId); err != nil {
		return err
	}

	err := mu.moduleRepository.Delete(moduleId)

	if err != nil {
		return err
	}

	return nil
}
