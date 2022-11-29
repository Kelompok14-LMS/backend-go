package drivers

import (
	otpDomain "github.com/Kelompok14-LMS/backend-go/businesses/otp"
	otpDB "github.com/Kelompok14-LMS/backend-go/drivers/redis/otp"
	"github.com/go-redis/redis/v8"

	menteeDomain "github.com/Kelompok14-LMS/backend-go/businesses/mentees"
	menteeDB "github.com/Kelompok14-LMS/backend-go/drivers/mysql/mentees"

	userDomain "github.com/Kelompok14-LMS/backend-go/businesses/users"
	userDB "github.com/Kelompok14-LMS/backend-go/drivers/mysql/users"
	"gorm.io/gorm"
)

func NewOTPRepository(client *redis.Client) otpDomain.Repository {
	return otpDB.NewRedisRepository(client)
}

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewMySQLRepository(conn)
}

func NewMenteeRepository(conn *gorm.DB) menteeDomain.Repository {
	return menteeDB.NewSQLRepository(conn)
}
