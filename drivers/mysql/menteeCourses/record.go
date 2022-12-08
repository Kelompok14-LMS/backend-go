package mentee_courses

import (
	"time"

	menteeCourses "github.com/Kelompok14-LMS/backend-go/businesses/menteeCourses"
)

type MenteeCourse struct {
	ID        string    `json:"id" gorm:"primaryKey;size:200"`
	MenteeId  string    `json:"mentee_id" gorm:"size:200"`
	CourseId  string    `json:"course_id" gorm:"size:200"`
	Status    string    `json:"status" gorm:"size:50"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (rec *MenteeCourse) ToDomain() *menteeCourses.Domain {
	return &menteeCourses.Domain{
		ID:        rec.ID,
		MenteeId:  rec.MenteeId,
		CourseId:  rec.CourseId,
		Status:    rec.Status,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func FromDomain(menteeCourseDomain *menteeCourses.Domain) *MenteeCourse {
	return &MenteeCourse{
		ID:        menteeCourseDomain.ID,
		MenteeId:  menteeCourseDomain.MenteeId,
		CourseId:  menteeCourseDomain.CourseId,
		Status:    menteeCourseDomain.Status,
		CreatedAt: menteeCourseDomain.CreatedAt,
		UpdatedAt: menteeCourseDomain.UpdatedAt,
	}
}