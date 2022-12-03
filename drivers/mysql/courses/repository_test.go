package courses

import (
	"database/sql/driver"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Kelompok14-LMS/backend-go/businesses/courses"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Anytime struct{}

func (a Anytime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)

	return ok
}

var birthDate = time.Date(2022, 8, 12, 0, 0, 0, 0, time.Local)

type suiteCourse struct {
	suite.Suite
	mock             sqlmock.Sqlmock
	courseRepository courses.Repository
}

func (s *suiteCourse) SetupSuite() {
	db, mock, err := sqlmock.New()

	s.NoError(err)

	s.mock = mock

	dbGorm, _ := gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      db,
	}))

	s.courseRepository = NewSQLRepository(dbGorm)
}

func (s *suiteCourse) TestCreate() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `courses` (`id`,`mentor_id`,`category_id`,`title`,`description`,`thumbnail`,`created_at`,`updated_at`,`deleted_at`) VALUES (?,?,?,?,?,?,?,?,?)")).
		WithArgs("COURSE_1", "MENTOR_1", "CAT_1", "UI/UX", "Description test.", "https://storage.googleapis.com/bucket/object", time.Now(), time.Now(), nil).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	course := courses.Domain{
		ID:          "COURSE_1",
		MentorId:    "MENTOR_1",
		CategoryId:  "CAT_1",
		Title:       "UI/UX",
		Description: "Description test.",
		Thumbnail:   "https://storage.googleapis.com/bucket/object",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   gorm.DeletedAt{},
	}

	err := s.courseRepository.Create(&course)

	s.NoError(err)
}

func (s *suiteCourse) TestFindAll() {
	courseRows := sqlmock.NewRows([]string{"id", "mentor_id", "category_id", "title", "description", "thumbnail", "created_at", "updated_at", "deleted_at"}).
		AddRow("COURSE_1", "MENTOR_1", "CAT_1", "UI/UX", "Description test.", "https://storage.googleapis.com/bucket/object", time.Now(), time.Now(), nil)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT `courses`.`id`,`courses`.`mentor_id`,`courses`.`category_id`,`courses`.`title`,`courses`.`description`,`courses`.`thumbnail`,`courses`.`created_at`,`courses`.`updated_at`,`courses`.`deleted_at` FROM `courses` INNER JOIN categories ON categories.id = courses.category_id WHERE (courses.title LIKE ? OR categories.name LIKE ?) AND `courses`.`deleted_at` IS NULL")).
		WithArgs("%%", "%%").
		WillReturnRows(courseRows)

	categoryRow := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).
		AddRow("CAT_1", "Programming", time.Now(), time.Now())

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `categories` WHERE `categories`.`id` = ?")).
		WithArgs("CAT_1").
		WillReturnRows(categoryRow)

	mentorRow := sqlmock.NewRows([]string{"id", "user_id", "fullname", "phone", "role", "jobs", "gender", "birth_place", "birth_date", "address", "profile_picture", "created_at", "updated_at"}).
		AddRow("MENTOR_1", "USER_1", "Mentor 1", "0857754321", "mentor", "frontend", "laki-laki", "bogor", birthDate, "Jl Bungur", "https://examples.com/to/bucket", time.Now(), time.Now())

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `mentors` WHERE `mentors`.`id` = ?")).
		WithArgs("MENTOR_1").
		WillReturnRows(mentorRow)

	results, err := s.courseRepository.FindAll("")

	s.Nil(err)
	s.NotNil(results)
}

func (s *suiteCourse) TestFindById() {
	courseRow := sqlmock.NewRows([]string{"id", "mentor_id", "category_id", "title", "description", "thumbnail", "created_at", "updated_at", "deleted_at"}).
		AddRow("COURSE_1", "MENTOR_1", "CAT_1", "UI/UX", "Description test.", "https://storage.googleapis.com/bucket/object", time.Now(), time.Now(), nil)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `courses` WHERE id = ? AND `courses`.`deleted_at` IS NULL ORDER BY `courses`.`id` LIMIT 1")).
		WithArgs("COURSE_1").
		WillReturnRows(courseRow)

	categoryRow := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).
		AddRow("CAT_1", "Programming", time.Now(), time.Now())

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `categories` WHERE `categories`.`id` = ?")).
		WithArgs("CAT_1").
		WillReturnRows(categoryRow)

	mentorRow := sqlmock.NewRows([]string{"id", "user_id", "fullname", "phone", "role", "jobs", "gender", "birth_place", "birth_date", "address", "profile_picture", "created_at", "updated_at"}).
		AddRow("MENTOR_1", "USER_1", "Mentor 1", "0857754321", "mentor", "frontend", "laki-laki", "bogor", birthDate, "Jl Bungur", "https://examples.com/to/bucket", time.Now(), time.Now())

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `mentors` WHERE `mentors`.`id` = ?")).
		WithArgs("MENTOR_1").
		WillReturnRows(mentorRow)

	result, err := s.courseRepository.FindById("COURSE_1")

	s.Nil(err)
	s.NotNil(result)
}

func (s *suiteCourse) TestFindByCategory() {
	courseRows := sqlmock.NewRows([]string{"id", "mentor_id", "category_id", "title", "description", "thumbnail", "created_at", "updated_at", "deleted_at"}).
		AddRow("COURSE_1", "MENTOR_1", "CAT_1", "UI/UX", "Description test.", "https://storage.googleapis.com/bucket/object", time.Now(), time.Now(), nil)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `courses` WHERE category_id = ? AND `courses`.`deleted_at` IS NULL")).
		WithArgs("CAT_1").
		WillReturnRows(courseRows)

	categoryRow := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).
		AddRow("CAT_1", "Programming", time.Now(), time.Now())

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `categories` WHERE `categories`.`id` = ?")).
		WithArgs("CAT_1").
		WillReturnRows(categoryRow)

	mentorRow := sqlmock.NewRows([]string{"id", "user_id", "fullname", "phone", "role", "jobs", "gender", "birth_place", "birth_date", "address", "profile_picture", "created_at", "updated_at"}).
		AddRow("MENTOR_1", "USER_1", "Mentor 1", "0857754321", "mentor", "frontend", "laki-laki", "bogor", birthDate, "Jl Bungur", "https://examples.com/to/bucket", time.Now(), time.Now())

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `mentors` WHERE `mentors`.`id` = ?")).
		WithArgs("MENTOR_1").
		WillReturnRows(mentorRow)

	results, err := s.courseRepository.FindByCategory("CAT_1")

	s.Nil(err)
	s.NotNil(results)
}

func (s *suiteCourse) TestUpdate() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `courses` SET `category_id`=?,`title`=?,`description`=?,`thumbnail`=?,`updated_at`=? WHERE id = ? AND `courses`.`deleted_at` IS NULL")).
		WithArgs("CAT_2", "Programming", "Updated desc", "https://storage.googleapis.com/bucket/object", Anytime{}, "COURSE_1").
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	course := courses.Domain{
		CategoryId:  "CAT_2",
		Title:       "Programming",
		Description: "Updated desc",
		Thumbnail:   "https://storage.googleapis.com/bucket/object",
		UpdatedAt:   time.Now(),
	}

	err := s.courseRepository.Update("COURSE_1", &course)

	s.NoError(err)
}

func (s *suiteCourse) TestDelete() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `courses` SET `deleted_at`=? WHERE id = ? AND `courses`.`deleted_at` IS NULL")).
		WithArgs(Anytime{}, "COURSE_1").
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	err := s.courseRepository.Delete("COURSE_1")

	s.NoError(err)
}

func TestSuiteCourse(t *testing.T) {
	suite.Run(t, new(suiteCourse))
}
