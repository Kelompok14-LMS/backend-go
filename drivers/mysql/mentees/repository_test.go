package mentees

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Kelompok14-LMS/backend-go/businesses/mentees"

	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type suiteMentee struct {
	suite.Suite
	mock             sqlmock.Sqlmock
	menteeRepository mentees.Repository
}

func (s *suiteMentee) SetupSuite() {
	db, mock, err := sqlmock.New()

	s.NoError(err)

	s.mock = mock

	dbGorm, _ := gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      db,
	}))

	s.menteeRepository = NewSQLRepository(dbGorm)
}

func (s *suiteMentee) TestCreate_Success() {
	birthDate := time.Date(2022, 8, 12, 0, 0, 0, 0, time.Local)

	menteeRow := sqlmock.NewRows([]string{"id", "user_id", "fullname", "phone", "role", "birth_date", "address", "profile_picture", "created_at", "updated_at"}).
		AddRow("MID1", "UID1", "Mentee 1", "087654321", "mentee", birthDate, "Jl Rinjani", "https://examples.com/to/bucket", time.Now(), time.Now())

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `mentees` WHERE user_id = ? ORDER BY `mentees`.`id` LIMIT 1")).
		WithArgs("UID1").
		WillReturnRows(menteeRow)

	mentee := mentees.Domain{
		ID:        "MID1",
		UserId:    "UID1",
		Fullname:  "Mentee 1",
		Phone:     "087654321",
		Role:      "mentee",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `mentees` (`id`,`user_id`,`fullname`,`phone`,`role`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?,?)")).
		WithArgs("MID1", "UID1", "Mentee 1", "087654321", "mentee", time.Now(), time.Now()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	if err := s.menteeRepository.Create(&mentee); err != nil {
		s.Error(err)
	}

	if err := s.mock.ExpectationsWereMet(); err != nil {
		s.Error(err)
	}
}

//func (s *suiteMentee) TestCreate_Failure() {
//	mentee := mentees.Domain{
//		ID:        "MID1",
//		UserId:    "UID1",
//		Fullname:  "Mentee 1",
//		Role:      "mentee",
//		Phone:     "087654321",
//		CreatedAt: time.Now(),
//		UpdatedAt: time.Now(),
//	}
//
//	s.mock.ExpectBegin()
//	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `mentees` (`id`,`user_id`,`fullname`,`phone`,`role`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?,?)")).
//		WithArgs("MID1", "UID1", "Mentee 1", "087654321", "mentee", time.Now(), time.Now()).
//		WillReturnResult(nil).
//		WillReturnError(fmt.Errorf("error occured"))
//	s.mock.ExpectRollback()
//
//	if err := s.menteeRepository.Create(&mentee); err != nil {
//		s.Error(err)
//	}
//}

func (s *suiteMentee) TestFindByIdUser() {
	birthDate := time.Date(2022, 8, 12, 0, 0, 0, 0, time.Local)

	mentee := sqlmock.NewRows([]string{"id", "user_id", "fullname", "phone", "role", "birth_date", "address", "profile_picture", "created_at", "updated_at"}).
		AddRow("MID1", "UID1", "Mentee 1", "087654321", "mentee", birthDate, "Jl Rinjani", "https://examples.com/to/bucket", time.Now(), time.Now())

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM mentees WHERE user_id = ?")).
		WithArgs("UID1").
		WillReturnRows(mentee)

	if _, err := s.menteeRepository.FindByIdUser("UID1"); err != nil {
		s.Error(err)
	}

	if err := s.mock.ExpectationsWereMet(); err != nil {
		s.Error(err)
	}
}

func TestSuiteMentee(t *testing.T) {
	suite.Run(t, new(suiteMentee))
}
