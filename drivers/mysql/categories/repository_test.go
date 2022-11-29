package categories

import (
	"database/sql/driver"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Kelompok14-LMS/backend-go/businesses/categories"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Anytime struct{}

func (a Anytime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)

	return ok
}

type suiteCategory struct {
	suite.Suite
	mock               sqlmock.Sqlmock
	categoryRepository categories.Repository
}

func (s *suiteCategory) SetupSuite() {
	db, mock, err := sqlmock.New()

	s.NoError(err)

	s.mock = mock

	dbGorm, _ := gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      db,
	}))

	s.categoryRepository = NewSQLRepository(dbGorm)
}

func (s *suiteCategory) TestCreate() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `categories` (`id`,`name`,`created_at`,`updated_at`) VALUES (?,?,?,?)")).
		WithArgs("CID1", "Programming", time.Now(), time.Now()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	category := categories.Domain{
		ID:        "CID1",
		Name:      "Programming",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := s.categoryRepository.Create(&category)

	s.NoError(err)
}

func (s *suiteCategory) TestFindAll() {
	rows := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).
		AddRow("CID1", "Programming", time.Now(), time.Now())

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `categories`")).
		WillReturnRows(rows)

	results, err := s.categoryRepository.FindAll()

	s.Nil(err)
	s.NotNil(results)
}

func (s *suiteCategory) TestFindById() {
	rows := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).
		AddRow("CID1", "Programming", time.Now(), time.Now())

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `categories` WHERE id = ?")).
		WithArgs("CID1").
		WillReturnRows(rows)

	result, err := s.categoryRepository.FindById("CID1")

	s.Nil(err)
	s.NotNil(result)
}

func (s *suiteCategory) TestUpdate() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `categories` SET `name`=?,`updated_at`=? WHERE id = ?")).
		WithArgs("UI/UX", Anytime{}, "CID1").
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	category := categories.Domain{
		Name: "UI/UX",
	}

	err := s.categoryRepository.Update("CID1", &category)

	s.NoError(err)
}

func TestSuiteCategory(t *testing.T) {
	suite.Run(t, new(suiteCategory))
}
