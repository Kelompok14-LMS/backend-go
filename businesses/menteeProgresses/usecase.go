package mentee_progresses

import (
	"github.com/Kelompok14-LMS/backend-go/businesses/courses"
	"github.com/Kelompok14-LMS/backend-go/businesses/materials"
	"github.com/Kelompok14-LMS/backend-go/businesses/mentees"
	"github.com/google/uuid"
)

type menteeProgressUsecase struct {
	menteeProgressRepository Repository
	menteeRepository         mentees.Repository
	courseRepository         courses.Repository
	materialRepository       materials.Repository
}

func NewMenteeProgressUsecase(
	menteeProgressRepository Repository,
	menteeRepository mentees.Repository,
	courseRepository courses.Repository,
	materialRepository materials.Repository,
) Usecase {
	return menteeProgressUsecase{
		menteeProgressRepository: menteeProgressRepository,
		menteeRepository:         menteeRepository,
		courseRepository:         courseRepository,
		materialRepository:       materialRepository,
	}
}

func (m menteeProgressUsecase) Add(menteeProgressDomain *Domain) error {
	if _, err := m.menteeRepository.FindById(menteeProgressDomain.MenteeId); err != nil {
		return err
	}

	if _, err := m.courseRepository.FindById(menteeProgressDomain.CourseId); err != nil {
		return err
	}

	if _, err := m.materialRepository.FindById(menteeProgressDomain.MaterialId); err != nil {
		return err
	}

	menteeProgress := Domain{
		ID:         uuid.NewString(),
		MenteeId:   menteeProgressDomain.MenteeId,
		CourseId:   menteeProgressDomain.CourseId,
		MaterialId: menteeProgressDomain.MaterialId,
	}

	err := m.menteeProgressRepository.Add(&menteeProgress)

	if err != nil {
		return err
	}

	return nil
}
