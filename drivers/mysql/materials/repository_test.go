package materials

import (
	"database/sql/driver"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Kelompok14-LMS/backend-go/businesses/materials"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Anytime struct{}

func (a Anytime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)

	return ok
}

type suiteMaterial struct {
	suite.Suite
	mock               sqlmock.Sqlmock
	materialRepository materials.Repository
}

func (s *suiteMaterial) SetupSuite() {
	db, mock, err := sqlmock.New()

	s.NoError(err)

	s.mock = mock

	dbGorm, _ := gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      db,
	}))

	s.materialRepository = NewSQLRepository(dbGorm)
}

func (s *suiteMaterial) TestCreate() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `materials` (`id`,`module_id`,`title`,`url`,`description`,`created_at`,`updated_at`,`deleted_at`) VALUES (?,?,?,?,?,?,?,?)")).
		WithArgs("MATERIAL_1", "MODULE_1", "Title test", "https://storage.com/to/bucket/object.mp4", "Description test", time.Now(), time.Now(), nil).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	materialDomain := materials.Domain{
		ID:          "MATERIAL_1",
		ModuleId:    "MODULE_1",
		Title:       "Title test",
		URL:         "https://storage.com/to/bucket/object.mp4",
		Description: "Description test",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   gorm.DeletedAt{},
	}

	err := s.materialRepository.Create(&materialDomain)

	s.NoError(err)
}

func (s *suiteMaterial) TestFindById() {
	rows := sqlmock.NewRows([]string{"id", "module_id", "title", "url", "description", "created_at", "updated_at", "deleted_at"}).
		AddRow("MATERIAL_1", "MODULE_1", "Title test", "https://storage.com/to/bucket/object.mp4", "Description test", time.Now(), time.Now(), nil)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `materials` WHERE id = ? AND `materials`.`deleted_at` IS NULL ORDER BY `materials`.`id` LIMIT 1")).
		WithArgs("MATERIAL_1").
		WillReturnRows(rows)

	result, err := s.materialRepository.FindById("MATERIAL_1")

	s.Nil(err)
	s.NotNil(result)
}

func (s *suiteMaterial) TestUpdate() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `materials` SET `module_id`=?,`title`=?,`url`=?,`description`=?,`updated_at`=? WHERE id = ? AND `materials`.`deleted_at` IS NULL")).
		WithArgs("MODULE_1", "Title test", "https://storage.com/to/bucket/object.mp4", "Description test", Anytime{}, "MATERIAL_1").
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	materialDomain := materials.Domain{
		ModuleId:    "MODULE_1",
		Title:       "Title test",
		URL:         "https://storage.com/to/bucket/object.mp4",
		Description: "Description test",
		UpdatedAt:   time.Now(),
	}

	err := s.materialRepository.Update("MATERIAL_1", &materialDomain)

	s.NoError(err)
}

func (s *suiteMaterial) TestDelete() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `materials` SET `deleted_at`=? WHERE id = ? AND `materials`.`deleted_at` IS NULL")).
		WithArgs(Anytime{}, "MATERIAL_1").
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	err := s.materialRepository.Delete("MATERIAL_1")

	s.NoError(err)
}

func TestSuiteMaterial(t *testing.T) {
	suite.Run(t, new(suiteMaterial))
}
