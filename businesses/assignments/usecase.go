package assignments

import (
	"github.com/Kelompok14-LMS/backend-go/businesses/courses"
	"github.com/Kelompok14-LMS/backend-go/pkg"
	"github.com/google/uuid"
)

type assignmentUsecase struct {
	assignmentRepository Repository
	courseRepository     courses.Repository
}

func NewAssignmentUsecase(assignmentRepository Repository, courseRepository courses.Repository) Usecase {
	return assignmentUsecase{
		assignmentRepository: assignmentRepository,
		courseRepository:     courseRepository,
	}
}

func (au assignmentUsecase) Create(assignmentDomain *Domain) error {
	if _, err := au.courseRepository.FindById(assignmentDomain.CourseId); err != nil {
		return err
	}

	id := uuid.NewString()

	assignment := Domain{
		ID:          id,
		CourseId:    assignmentDomain.CourseId,
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

func (au assignmentUsecase) FindByCourseId(courseId string) (*Domain, error) {
	assignment, err := au.assignmentRepository.FindByCourseId(courseId)

	if err != nil {
		return nil, pkg.ErrCourseNotFound
	}

	return assignment, nil
}

func (au assignmentUsecase) Update(assignmentId string, assignmentDomain *Domain) error {
	if _, err := au.courseRepository.FindById(assignmentDomain.CourseId); err != nil {
		return err
	}

	_, err := au.assignmentRepository.FindById(assignmentId)

	if err != nil {
		return pkg.ErrAssignmentNotFound
	}

	// updatedAssignment := Domain{
	// 	CourseId:    assignmentDomain.CourseId,
	// 	Title:       assignmentDomain.Title,
	// 	Description: assignmentDomain.Description,
	// }

	err = au.assignmentRepository.Update(assignmentId, assignmentDomain)

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
