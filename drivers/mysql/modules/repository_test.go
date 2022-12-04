package modules

import (
	"database/sql/driver"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Kelompok14-LMS/backend-go/businesses/modules"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Anytime struct{}

func (a Anytime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)

	return ok
}

type suiteModule struct {
	suite.Suite
	mock             sqlmock.Sqlmock
	moduleRepository modules.Repository
}

func (s *suiteModule) SetupSuite() {
	db, mock, err := sqlmock.New()

	s.NoError(err)

	s.mock = mock

	dbGorm, _ := gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      db,
	}))

	s.moduleRepository = NewSQLRepository(dbGorm)
}

func (s *suiteModule) TestCreate() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `modules` (`id`,`course_id`,`title`,`created_at`,`updated_at`,`deleted_at`) VALUES (?,?,?,?,?,?)")).
		WithArgs("MOD_1", "COURSE_1", "Module Test", time.Now(), time.Now(), nil).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	courseDomain := modules.Domain{
		ID:        "MOD_1",
		CourseId:  "COURSE_1",
		Title:     "Module Test",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: gorm.DeletedAt{},
	}

	err := s.moduleRepository.Create(&courseDomain)

	s.NoError(err)
}

func (s *suiteModule) TestFindById() {
	row := sqlmock.NewRows([]string{"id", "course_id", "title", "created_at", "updated_at", "deleted_at"}).
		AddRow("MOD_1", "COURSE_1", "Module Test", time.Now(), time.Now(), nil)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `modules` WHERE id = ? AND `modules`.`deleted_at` IS NULL ORDER BY `modules`.`id` LIMIT 1")).
		WithArgs("MOD_1").
		WillReturnRows(row)

	result, err := s.moduleRepository.FindById("MOD_1")

	s.Nil(err)
	s.NotNil(result)
}

func (s *suiteModule) TestUpdate() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `modules` SET `course_id`=?,`title`=?,`updated_at`=? WHERE id = ? AND `modules`.`deleted_at` IS NULL")).
		WithArgs("COURSE_1", "Module Test Updated", Anytime{}, "MOD_1").
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	courseDomain := modules.Domain{
		CourseId:  "COURSE_1",
		Title:     "Module Test Updated",
		UpdatedAt: time.Now(),
	}

	err := s.moduleRepository.Update("MOD_1", &courseDomain)

	s.NoError(err)
}

func (s *suiteModule) TestDelete() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `modules` SET `deleted_at`=? WHERE id = ? AND `modules`.`deleted_at` IS NULL")).
		WithArgs(Anytime{}, "MOD_1").
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	err := s.moduleRepository.Delete("MOD_1")

	s.NoError(err)
}

func TestSuiteModule(t *testing.T) {
	suite.Run(t, new(suiteModule))
}
