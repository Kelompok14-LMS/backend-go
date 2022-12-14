package mysql_driver

import (
	"fmt"
	"log"

	"github.com/Kelompok14-LMS/backend-go/drivers/mysql/assignments"
	"github.com/Kelompok14-LMS/backend-go/drivers/mysql/categories"
	"github.com/Kelompok14-LMS/backend-go/drivers/mysql/courses"
	"github.com/Kelompok14-LMS/backend-go/drivers/mysql/materials"
	menteeAssignments "github.com/Kelompok14-LMS/backend-go/drivers/mysql/menteeAssignments"
	menteeCourses "github.com/Kelompok14-LMS/backend-go/drivers/mysql/menteeCourses"
	menteeProgresses "github.com/Kelompok14-LMS/backend-go/drivers/mysql/menteeProgresses"
	"github.com/Kelompok14-LMS/backend-go/drivers/mysql/mentees"
	"github.com/Kelompok14-LMS/backend-go/drivers/mysql/mentors"
	"github.com/Kelompok14-LMS/backend-go/drivers/mysql/modules"
	"github.com/Kelompok14-LMS/backend-go/drivers/mysql/reviews"
	"github.com/Kelompok14-LMS/backend-go/drivers/mysql/users"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ConfigDB struct {
	MYSQL_USERNAME string
	MYSQL_PASSWORD string
	MYSQL_NAME     string
	MYSQL_HOST     string
	MYSQL_PORT     string
}

func (config *ConfigDB) InitMySQLDatabase() *gorm.DB {
	var err error

	var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.MYSQL_USERNAME,
		config.MYSQL_PASSWORD,
		config.MYSQL_HOST,
		config.MYSQL_PORT,
		config.MYSQL_NAME,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("error when connecting to the database: %s", err)
	}

	log.Println("connected to the database")

	return db
}

func DBMigrate(db *gorm.DB) {
	_ = db.AutoMigrate(
		&users.User{},
		&mentees.Mentee{},
		&mentors.Mentor{},
		&categories.Category{},
		&courses.Course{},
		&modules.Module{},
		&materials.Material{},
		&assignments.Assignment{},
		&menteeCourses.MenteeCourse{},
		&menteeProgresses.MenteeProgress{},
		&menteeAssignments.MenteeAssignment{},
		&reviews.Review{},
	)
}

func CloseDB(db *gorm.DB) error {
	database, err := db.DB()

	if err != nil {
		log.Printf("error when getting the database instance: %v", err)
		return err
	}

	if err := database.Close(); err != nil {
		log.Printf("error when closing the database connection: %v", err)
		return err
	}

	log.Println("database connection is closed")

	return nil
}
