package users

import (
	"database/sql/driver"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Kelompok14-LMS/backend-go/businesses/users"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type anytime struct{}

func (a anytime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

type suiteUser struct {
	suite.Suite
	mock           sqlmock.Sqlmock
	userRepository users.Repository
}

func (s *suiteUser) SetupSuite() {
	db, mock, err := sqlmock.New()

	s.NoError(err)

	s.mock = mock

	dbGorm, _ := gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      db,
	}))

	s.userRepository = NewMySQLRepository(dbGorm)
}

func (s *suiteUser) TestCreate_Success() {
	user := users.Domain{
		ID:        "UID1",
		Email:     "mentee@gmail.com",
		Password:  "123456",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`id`,`email`,`password`,`created_at`,`updated_at`) VALUES (?,?,?,?,?)")).
		WithArgs("UID1", "mentee@gmail.com", "123456", time.Now(), time.Now()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	err := s.userRepository.Create(&user)

	if err != nil {
		s.Error(err)
	}

	s.Nil(err)

	if err := s.mock.ExpectationsWereMet(); err != nil {
		s.Error(err)
	}
}

//func (s *suiteUser) TestCreate_Failed() {
//	s.T().Run("test create failed", func(t *testing.T) {
//		user := users.Domain{
//			ID:        "UID1",
//			Password:  "123456",
//			CreatedAt: time.Now(),
//			UpdatedAt: time.Now(),
//		}
//
//		s.mock.ExpectBegin()
//		s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`id`,`password`,`created_at`,`updated_at`) VALUES (?,?,?,?)")).
//			WithArgs("UID1", "123456", time.Now(), time.Now()).
//			WillReturnError(fmt.Errorf("error occured"))
//		s.mock.ExpectRollback()
//
//		err := s.userRepository.Create(&user)
//
//		//if err != nil {
//		//s.Error(err)
//		//}
//		t.Error(err)
//
//		//if err := s.mock.ExpectationsWereMet(); err != nil {
//		//	s.Error(err)
//		//}
//	})
//}

func (s *suiteUser) TestFindByEmail() {
	user := sqlmock.NewRows([]string{"id", "email", "password", "created_at", "updated_at"}).
		AddRow("UID1", "mentee@gmail.com", "hashedpassword", time.Now(), time.Now())

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE email = ? ORDER BY `users`.`id` LIMIT 1")).
		WithArgs("mentee@gmail.com").
		WillReturnRows(user)

	result, err := s.userRepository.FindByEmail("mentee@gmail.com")

	if err != nil {
		s.Error(err)
	}

	s.Nil(err)
	s.NotNil(result)

	if err := s.mock.ExpectationsWereMet(); err != nil {
		s.Error(err)
	}
}

func (s *suiteUser) TestFindById() {
	user := sqlmock.NewRows([]string{"id", "email", "password", "created_at", "updated_at"}).
		AddRow("UID1", "mentee@gmail.com", "hashedpassword", time.Now(), time.Now())

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE id = ? ORDER BY `users`.`id` LIMIT 1")).
		WithArgs("UID1").
		WillReturnRows(user)

	result, err := s.userRepository.FindById("UID1")

	if err != nil {
		s.Error(err)
	}

	s.Nil(err)
	s.NotNil(result)

	if err := s.mock.ExpectationsWereMet(); err != nil {
		s.Error(err)
	}
}

func (s *suiteUser) TestUpdate() {
	user := users.Domain{
		Password:  "updatedhashedpassword",
		UpdatedAt: time.Now(),
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `users` SET `password`=?,`updated_at`=? WHERE id = ?")).
		WithArgs("updatedhashedpassword", anytime{}, "UID1").
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	err := s.userRepository.Update("UID1", &user)

	if err != nil {
		s.Error(err)
	}

	s.Nil(err)

	if err := s.mock.ExpectationsWereMet(); err != nil {
		s.Error(err)
	}
}

func TestSuiteUser(t *testing.T) {
	suite.Run(t, new(suiteUser))
}
