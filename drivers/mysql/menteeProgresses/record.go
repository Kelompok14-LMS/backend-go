package mentee_progresses

import (
	"time"

	menteeProgresses "github.com/Kelompok14-LMS/backend-go/businesses/menteeProgresses"
	"github.com/Kelompok14-LMS/backend-go/drivers/mysql/courses"
	"github.com/Kelompok14-LMS/backend-go/drivers/mysql/materials"
	"github.com/Kelompok14-LMS/backend-go/drivers/mysql/mentees"
)

type MenteeProgress struct {
	ID         string `json:"id" gorm:"primaryKey;size:200"`
	MenteeId   string `json:"mentee_id" gorm:"size:200"`
	CourseId   string `json:"course_id" gorm:"size:200"`
	MaterialId string `json:"material_id" gorm:"size:200"`
	Mentee     mentees.Mentee
	Course     courses.Course
	Material   materials.Material
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (req *MenteeProgress) ToDomain() *menteeProgresses.Domain {
	return &menteeProgresses.Domain{
		ID:         req.ID,
		MenteeId:   req.MenteeId,
		CourseId:   req.CourseId,
		MaterialId: req.MaterialId,
		Mentee:     *req.Mentee.ToDomain(),
		Course:     *req.Course.ToDomain(),
		Material:   *req.Material.ToDomain(),
		CreatedAt:  req.CreatedAt,
		UpdatedAt:  req.UpdatedAt,
	}
}

func FromDomain(menteeProgressDomain *menteeProgresses.Domain) *MenteeProgress {
	return &MenteeProgress{
		ID:         menteeProgressDomain.ID,
		MenteeId:   menteeProgressDomain.MenteeId,
		CourseId:   menteeProgressDomain.CourseId,
		MaterialId: menteeProgressDomain.MaterialId,
		CreatedAt:  menteeProgressDomain.CreatedAt,
		UpdatedAt:  menteeProgressDomain.UpdatedAt,
	}
}
