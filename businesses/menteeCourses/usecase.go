package mentee_courses

import (
	"github.com/Kelompok14-LMS/backend-go/businesses/courses"
	"github.com/Kelompok14-LMS/backend-go/businesses/materials"
	menteeProgresses "github.com/Kelompok14-LMS/backend-go/businesses/menteeProgresses"
	"github.com/Kelompok14-LMS/backend-go/businesses/mentees"
	"github.com/Kelompok14-LMS/backend-go/pkg"
	"github.com/google/uuid"
)

type menteeCourseUsecase struct {
	menteeCourseRepository   Repository
	menteeRepository         mentees.Repository
	courseRepository         courses.Repository
	materialRepository       materials.Repository
	menteeProgressRepository menteeProgresses.Repository
}

func NewMenteeCourseUsecase(
	menteeCourseRepository Repository,
	menteeRepository mentees.Repository,
	courseRepository courses.Repository,
	materialRepository materials.Repository,
	menteeProgressRepository menteeProgresses.Repository,
) Usecase {
	return menteeCourseUsecase{
		menteeCourseRepository:   menteeCourseRepository,
		menteeRepository:         menteeRepository,
		courseRepository:         courseRepository,
		materialRepository:       materialRepository,
		menteeProgressRepository: menteeProgressRepository,
	}
}

func (m menteeCourseUsecase) Enroll(menteeCourseDomain *Domain) error {
	if _, err := m.menteeRepository.FindById(menteeCourseDomain.MenteeId); err != nil {
		return err
	}

	if _, err := m.courseRepository.FindById(menteeCourseDomain.CourseId); err != nil {
		return err
	}

	isEnrolled, _ := m.menteeCourseRepository.CheckEnrollment(menteeCourseDomain.MenteeId, menteeCourseDomain.CourseId)

	if isEnrolled != nil {
		return pkg.ErrAlreadyEnrolled
	}

	menteeCourseId := uuid.NewString()

	menteeCourse := Domain{
		ID:       menteeCourseId,
		MenteeId: menteeCourseDomain.MenteeId,
		CourseId: menteeCourseDomain.CourseId,
		Status:   "ongoing",
	}

	if err := m.menteeCourseRepository.Enroll(&menteeCourse); err != nil {
		return err
	}

	return nil
}

func (m menteeCourseUsecase) FindMenteeCourses(menteeId string, title string, status string) (*[]Domain, error) {
	menteeCourses, err := m.menteeCourseRepository.FindCoursesByMentee(menteeId, title, status)

	if err != nil {
		return nil, err
	}

	progresses, err := m.menteeProgressRepository.Count(menteeId, title, status)

	if err != nil {
		return nil, err
	}

	courseIds := []string{}

	for _, course := range *menteeCourses {
		courseIds = append(courseIds, course.CourseId)
	}

	totalMaterials, err := m.materialRepository.CountByCourse(courseIds)

	if err != nil {
		return nil, err
	}

	for i, progress := range progresses {
		(*menteeCourses)[i].ProgressCount = progress
	}

	for i, material := range totalMaterials {
		(*menteeCourses)[i].TotalMaterials = material
	}

	return menteeCourses, nil
}

func (m menteeCourseUsecase) CheckEnrollment(menteeId string, courseId string) (bool, error) {
	menteeCourseDomain, _ := m.menteeCourseRepository.CheckEnrollment(menteeId, courseId)

	isEnrolled := menteeCourseDomain != nil

	return isEnrolled, nil
}
