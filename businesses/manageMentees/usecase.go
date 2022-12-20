package manage_mentees

import (
	"context"

	menteeAssignments "github.com/Kelompok14-LMS/backend-go/businesses/menteeAssignments"
	menteeCourses "github.com/Kelompok14-LMS/backend-go/businesses/menteeCourses"
	menteeProgresses "github.com/Kelompok14-LMS/backend-go/businesses/menteeProgresses"
	"github.com/Kelompok14-LMS/backend-go/helper"
)

type manageMenteeUsecase struct {
	menteeCourse     menteeCourses.Repository
	menteeProgress   menteeProgresses.Repository
	menteeAssignment menteeAssignments.Repository
	storage          *helper.StorageConfig
}

func NewManageMenteeUsecase(
	menteeCourse menteeCourses.Repository,
	menteeProgress menteeProgresses.Repository,
	menteeAssignment menteeAssignments.Repository,
	storage *helper.StorageConfig,
) Usecase {
	return manageMenteeUsecase{
		menteeCourse:     menteeCourse,
		menteeProgress:   menteeProgress,
		menteeAssignment: menteeAssignment,
		storage:          storage,
	}
}

func (mm manageMenteeUsecase) DeleteAccess(menteeId string, courseId string) error {
	if _, err := mm.menteeCourse.CheckEnrollment(menteeId, courseId); err != nil {
		return err
	}

	assignment, _ := mm.menteeAssignment.FindByCourse(menteeId, courseId)

	if err := mm.menteeProgress.DeleteMenteeProgressesByCourse(menteeId, courseId); err != nil {
		return err
	}

	if err := mm.menteeCourse.DeleteEnrolledCourse(menteeId, courseId); err != nil {
		return err
	}

	if assignment != nil {
		if err := mm.menteeAssignment.Delete(assignment.ID); err != nil {
			return err
		}

		ctx := context.Background()

		if err := mm.storage.DeleteObject(ctx, assignment.AssignmentURL); err != nil {
			return err
		}
	}

	return nil
}
