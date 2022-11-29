package drivers

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"

	otpDomain "github.com/Kelompok14-LMS/backend-go/businesses/otp"
	otpDB "github.com/Kelompok14-LMS/backend-go/drivers/redis/otp"

	menteeDomain "github.com/Kelompok14-LMS/backend-go/businesses/mentees"
	menteeDB "github.com/Kelompok14-LMS/backend-go/drivers/mysql/mentees"

	userDomain "github.com/Kelompok14-LMS/backend-go/businesses/users"
	userDB "github.com/Kelompok14-LMS/backend-go/drivers/mysql/users"

	categoryDomain "github.com/Kelompok14-LMS/backend-go/businesses/categories"
	categoryDB "github.com/Kelompok14-LMS/backend-go/drivers/mysql/categories"
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

func NewCategoryRepository(conn *gorm.DB) categoryDomain.Repository {
	return categoryDB.NewSQLRepository(conn)
}