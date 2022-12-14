package drivers

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"

	otpDomain "github.com/Kelompok14-LMS/backend-go/businesses/otp"
	otpDB "github.com/Kelompok14-LMS/backend-go/drivers/redis/otp"

	menteeDomain "github.com/Kelompok14-LMS/backend-go/businesses/mentees"
	menteeDB "github.com/Kelompok14-LMS/backend-go/drivers/mysql/mentees"

	mentorsDomain "github.com/Kelompok14-LMS/backend-go/businesses/mentors"
	mentorsDB "github.com/Kelompok14-LMS/backend-go/drivers/mysql/mentors"

	userDomain "github.com/Kelompok14-LMS/backend-go/businesses/users"
	userDB "github.com/Kelompok14-LMS/backend-go/drivers/mysql/users"

	categoryDomain "github.com/Kelompok14-LMS/backend-go/businesses/categories"
	categoryDB "github.com/Kelompok14-LMS/backend-go/drivers/mysql/categories"

	courseDomain "github.com/Kelompok14-LMS/backend-go/businesses/courses"
	courseDB "github.com/Kelompok14-LMS/backend-go/drivers/mysql/courses"

	moduleDomain "github.com/Kelompok14-LMS/backend-go/businesses/modules"
	moduleDB "github.com/Kelompok14-LMS/backend-go/drivers/mysql/modules"

	assignmentDomain "github.com/Kelompok14-LMS/backend-go/businesses/assignments"
	assignmentDB "github.com/Kelompok14-LMS/backend-go/drivers/mysql/assignments"

	materialDomain "github.com/Kelompok14-LMS/backend-go/businesses/materials"
	materialDB "github.com/Kelompok14-LMS/backend-go/drivers/mysql/materials"

	menteeCoursesDomain "github.com/Kelompok14-LMS/backend-go/businesses/menteeCourses"
	menteeCoursesDB "github.com/Kelompok14-LMS/backend-go/drivers/mysql/menteeCourses"

	menteeProgressesDomain "github.com/Kelompok14-LMS/backend-go/businesses/menteeProgresses"
	menteeProgressesDB "github.com/Kelompok14-LMS/backend-go/drivers/mysql/menteeProgresses"

	menteeAssignmentsDomain "github.com/Kelompok14-LMS/backend-go/businesses/menteeAssignments"
	menteeAssignmentsDB "github.com/Kelompok14-LMS/backend-go/drivers/mysql/menteeAssignments"
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

func NewMentorRepository(conn *gorm.DB) mentorsDomain.Repository {
	return mentorsDB.NewSQLRepository(conn)
}

func NewCategoryRepository(conn *gorm.DB) categoryDomain.Repository {
	return categoryDB.NewSQLRepository(conn)
}

func NewCourseRepository(conn *gorm.DB) courseDomain.Repository {
	return courseDB.NewSQLRepository(conn)
}

func NewModuleRepository(conn *gorm.DB) moduleDomain.Repository {
	return moduleDB.NewSQLRepository(conn)
}

func NewAssignmentRepository(conn *gorm.DB) assignmentDomain.Repository {
	return assignmentDB.NewSQLRepository(conn)
}

func NewMaterialRepository(conn *gorm.DB) materialDomain.Repository {
	return materialDB.NewSQLRepository(conn)
}

func NewMenteeCourseRepository(conn *gorm.DB) menteeCoursesDomain.Repository {
	return menteeCoursesDB.NewSQLRepository(conn)
}

func NewMenteeProgressRepository(conn *gorm.DB) menteeProgressesDomain.Repository {
	return menteeProgressesDB.NewSQLRepository(conn)
}

func NewMenteeAssignmentRepository(conn *gorm.DB) menteeAssignmentsDomain.Repository {
	return menteeAssignmentsDB.NewSQLRepository(conn)
}
