package mentee_courses

import (
	menteeCourses "github.com/Kelompok14-LMS/backend-go/businesses/menteeCourses"
	"gorm.io/gorm"
)

type menteeCourseRepository struct {
	conn *gorm.DB
}

func NewSQLRepository(conn *gorm.DB) menteeCourses.Repository {
	return menteeCourseRepository{
		conn: conn,
	}
}

func (m menteeCourseRepository) Enroll(menteeCourseDomain *menteeCourses.Domain) error {
	rec := FromDomain(menteeCourseDomain)

	err := m.conn.Model(&MenteeCourse{}).Create(&rec).Error

	if err != nil {
		return err
	}

	return nil
}

func (m menteeCourseRepository) FindCoursesByMentee(menteeId string, title string, status string) (*[]menteeCourses.Domain, error) {
	var rec []MenteeCourse

	err := m.conn.Model(&MenteeCourse{}).Preload("Course.Mentor").
		Joins("INNER JOIN courses ON courses.id = mentee_courses.course_id").
		Where("mentee_courses.mentee_id = ? AND courses.title LIKE ? AND status = ? AND courses.deleted_at IS NULL", menteeId, "%"+title+"%", status).
		Order("mentee_courses.course_id ASC").
		Find(&rec).Error

	if err != nil {
		return nil, err
	}

	var courses []menteeCourses.Domain

	for _, course := range rec {
		courses = append(courses, *course.ToDomain())
	}

	return &courses, nil
}

func (m menteeCourseRepository) CheckEnrollment(menteeId string, courseId string) (*menteeCourses.Domain, error) {
	rec := MenteeCourse{}

	err := m.conn.Model(&MenteeCourse{}).
		Where("mentee_courses.mentee_id = ? AND mentee_courses.course_id = ?", menteeId, courseId).
		First(&rec).Error

	if err != nil {
		return nil, err
	}

	return rec.ToDomain(), nil
}
