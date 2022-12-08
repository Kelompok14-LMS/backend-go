package mentee_courses

import (
	"github.com/Kelompok14-LMS/backend-go/businesses/courses"
	"github.com/Kelompok14-LMS/backend-go/businesses/mentees"
	"github.com/google/uuid"
)

type menteeCourseUsecase struct {
	menteeCourseRepository Repository
	menteeRepository       mentees.Repository
	courseRepository       courses.Repository
}

func NewMenteeCourseUsecase(
	menteeCourseRepository Repository,
	menteeRepository mentees.Repository,
	courseRepository courses.Repository,
) Usecase {
	return menteeCourseUsecase{
		menteeCourseRepository: menteeCourseRepository,
		menteeRepository:       menteeRepository,
		courseRepository:       courseRepository,
	}
}

func (m menteeCourseUsecase) Enroll(menteeCourseDomain *Domain) error {
	if _, err := m.menteeRepository.FindById(menteeCourseDomain.MenteeId); err != nil {
		return err
	}

	if _, err := m.courseRepository.FindById(menteeCourseDomain.CourseId); err != nil {
		return err
	}

	menteeCourseId := uuid.NewString()

	menteeCourse := Domain{
		ID:       menteeCourseId,
		MenteeId: menteeCourseDomain.MenteeId,
		CourseId: menteeCourseDomain.CourseId,
		Status:   "ongoing",
	}

	err := m.menteeCourseRepository.Enroll(&menteeCourse)

	if err != nil {
		return err
	}

	return nil
}

func (m menteeCourseUsecase) FindMenteeCourses(menteeId string, title string, status string) (*[]Domain, error) {
	coursesDomain, err := m.courseRepository.FindByMentee(menteeId, title, status)

	if err != nil {
		return nil, err
	}

	menteeCoursesDomain := make([]Domain, len(*coursesDomain))

	for i := range *coursesDomain {
		menteeCoursesDomain[i].Course = (*coursesDomain)[i]
	}

	return &menteeCoursesDomain, nil
}

func (m menteeCourseUsecase) CheckEnrollment(menteeId string, courseId string) (bool, error) {
	menteeCourseDomain, err := m.menteeCourseRepository.CheckEnrollment(menteeId, courseId)

	if err != nil {
		return false, err
	}

	isEnrolled := menteeCourseDomain == nil

	return isEnrolled, nil
}
