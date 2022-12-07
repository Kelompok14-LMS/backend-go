package mentors

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Kelompok14-LMS/backend-go/businesses/mentors"

	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type suiteMentor struct {
	suite.Suite
	mock             sqlmock.Sqlmock
	mentorRepository mentors.Repository
}

func (s *suiteMentor) SetupSuite() {
	db, mock, err := sqlmock.New()

	s.NoError(err)

	s.mock = mock

	dbGorm, _ := gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      db,
	}))

	s.mentorRepository = NewSQLRepository(dbGorm)
}

func (s *suiteMentor) TestCreate_Success() {
	birthDate := time.Date(2022, 8, 12, 0, 0, 0, 0, time.Local)

	mentorRow := sqlmock.NewRows([]string{"id", "user_id", "fullname", "phone", "role", "jobs", "gender", "birth_place", "birth_date", "address", "profile_picture", "created_at", "updated_at"}).
		AddRow("MID1", "UID1", "Mentor 1", "0857754321", "mentor", "frontend", "laki-laki", "bogor", birthDate, "Jl Bungur", "https://examples.com/to/bucket", time.Now(), time.Now())

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `mentors` WHERE user_id = ? ORDER BY `mentors`.`id` LIMIT 1")).
		WithArgs("UID1").
		WillReturnRows(mentorRow)

	mentor := mentors.Domain{
		ID:        "MID1",
		UserId:    "UID1",
		Fullname:  "Mentor 1",
		Role:      "mentor",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `mentors` (`id`,`user_id`,`fullname`,`role`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?)")).
		WithArgs("MID1", "UID1", "Mentor 1", "0857754321", "mentor", time.Now(), time.Now()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	if err := s.mentorRepository.Create(&mentor); err != nil {
		s.Error(err)
	}

	if err := s.mock.ExpectationsWereMet(); err != nil {
		s.Error(err)
	}
}

func (s *suiteMentor) TestFindByIdUser() {
	birthDate := time.Date(2022, 8, 12, 0, 0, 0, 0, time.Local)

	mentor := sqlmock.NewRows([]string{"id", "user_id", "fullname", "phone", "role", "jobs", "gender", "birth_place", "birth_date", "address", "profile_picture", "created_at", "updated_at"}).
		AddRow("MID1", "UID1", "Mentor 1", "0857754321", "mentor", "frontend", "laki-laki", "bogor", birthDate, "Jl Bungur", "https://examples.com/to/bucket", time.Now(), time.Now())

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM mentors WHERE user_id = ?")).
		WithArgs("UID1").
		WillReturnRows(mentor)

	if _, err := s.mentorRepository.FindByIdUser("UID1"); err != nil {
		s.Error(err)
	}

	if err := s.mock.ExpectationsWereMet(); err != nil {
		s.Error(err)
	}
}

func TestSuiteMentor(t *testing.T) {
	suite.Run(t, new(suiteMentor))
}
