package manage_mentees

import (
	menteeCourses "github.com/Kelompok14-LMS/backend-go/businesses/menteeCourses"
	menteeProgresses "github.com/Kelompok14-LMS/backend-go/businesses/menteeProgresses"
)

type manageMenteeUsecase struct {
	menteeCourse   menteeCourses.Repository
	menteeProgress menteeProgresses.Repository
}

func NewManageMenteeUsecase(
	menteeCourse menteeCourses.Repository,
	menteeProgress menteeProgresses.Repository,
) Usecase {
	return manageMenteeUsecase{
		menteeCourse:   menteeCourse,
		menteeProgress: menteeProgress,
	}
}

func (mm manageMenteeUsecase) DeleteAccess(menteeId string, courseId string) error {
	if _, err := mm.menteeCourse.CheckEnrollment(menteeId, courseId); err != nil {
		return err
	}

	if err := mm.menteeProgress.DeleteMenteeProgressesByCourse(menteeId, courseId); err != nil {
		return err
	}

	if err := mm.menteeCourse.DeleteEnrolledCourse(menteeId, courseId); err != nil {
		return err
	}

	return nil
}
