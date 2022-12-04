package routes

import (
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"github.com/Kelompok14-LMS/backend-go/app/middlewares"
	_driverFactory "github.com/Kelompok14-LMS/backend-go/drivers"
	"github.com/Kelompok14-LMS/backend-go/helper"
	"github.com/Kelompok14-LMS/backend-go/pkg"
	"github.com/Kelompok14-LMS/backend-go/utils"

	_menteeUsecase "github.com/Kelompok14-LMS/backend-go/businesses/mentees"
	_menteeController "github.com/Kelompok14-LMS/backend-go/controllers/mentees"

	_mentorUsecase "github.com/Kelompok14-LMS/backend-go/businesses/mentors"
	_mentorController "github.com/Kelompok14-LMS/backend-go/controllers/mentors"

	_otpUsecase "github.com/Kelompok14-LMS/backend-go/businesses/otp"
	_otpController "github.com/Kelompok14-LMS/backend-go/controllers/otp"

	_categoryUsecase "github.com/Kelompok14-LMS/backend-go/businesses/categories"
	_categoryController "github.com/Kelompok14-LMS/backend-go/controllers/categories"

	_courseUsecase "github.com/Kelompok14-LMS/backend-go/businesses/courses"
	_courseController "github.com/Kelompok14-LMS/backend-go/controllers/courses"

	_moduleUsecase "github.com/Kelompok14-LMS/backend-go/businesses/modules"
	_moduleController "github.com/Kelompok14-LMS/backend-go/controllers/modules"
)

type RouteConfig struct {
	// echo top level instance
	Echo *echo.Echo

	// mysql conn
	MySQLDB *gorm.DB

	// redis conn
	RedisDB *redis.Client

	// JWT config
	JWTConfig *utils.JWTConfig

	// mail config
	Mailer *pkg.MailerConfig

	// cloud storage config
	StorageConfig *helper.StorageConfig
}

func (routeConfig *RouteConfig) New() {
	// setup api v1
	v1 := routeConfig.Echo.Group("/api/v1")

	// setup auth middleware
	authMiddleware := middlewares.NewAuthMiddleware(routeConfig.JWTConfig)

	// Inject the dependency to user
	userRepository := _driverFactory.NewUserRepository(routeConfig.MySQLDB)

	// Inject the dependency to otp
	otpRepository := _driverFactory.NewOTPRepository(routeConfig.RedisDB)
	otpUsecase := _otpUsecase.NewOTPUsecase(otpRepository, userRepository, routeConfig.Mailer)
	otpController := _otpController.NewOTPController(otpUsecase)

	// Inject the dependency to mentee
	menteeRepository := _driverFactory.NewMenteeRepository(routeConfig.MySQLDB)
	menteeUsecase := _menteeUsecase.NewMenteeUsecase(menteeRepository, userRepository, otpRepository, routeConfig.JWTConfig, routeConfig.Mailer)
	menteeController := _menteeController.NewMenteeController(menteeUsecase)

	// Inject the dependency to mentor
	mentorRepository := _driverFactory.NewMentorRepository(routeConfig.MySQLDB)
	mentorUsecase := _mentorUsecase.NewMentorUsecase(mentorRepository, userRepository, routeConfig.JWTConfig)
	mentorController := _mentorController.NewMentorController(mentorUsecase)

	// Inject the dependency to category
	categoryRepository := _driverFactory.NewCategoryRepository(routeConfig.MySQLDB)
	categoryUsecase := _categoryUsecase.NewCategoryUsecase(categoryRepository)
	categoryController := _categoryController.NewCategoryController(categoryUsecase)

	// Inject the dependency to course
	courseRepository := _driverFactory.NewCourseRepository(routeConfig.MySQLDB)
	courseUsecase := _courseUsecase.NewCourseUsecase(courseRepository, mentorRepository, categoryRepository, routeConfig.StorageConfig)
	courseController := _courseController.NewCourseController(courseUsecase)

	// Inject the dependency to module
	moduleRepository := _driverFactory.NewModuleRepository(routeConfig.MySQLDB)
	moduleUsecase := _moduleUsecase.NewModuleUsecase(moduleRepository, courseRepository)
	moduleController := _moduleController.NewModuleController(moduleUsecase)

	// authentication routes
	auth := v1.Group("/auth")
	auth.POST("/mentee/login", menteeController.HandlerLoginMentee)
	auth.POST("/mentee/register", menteeController.HandlerRegisterMentee)
	auth.POST("/mentee/register/verify", menteeController.HandlerVerifyRegisterMentee)
	auth.POST("/forgot-password", menteeController.HandlerForgotPassword)
	auth.POST("/send-otp", otpController.HandlerSendOTP)
	auth.POST("/check-otp", otpController.HandlerCheckOTP)
	auth.POST("/mentor/login", mentorController.HandlerLoginMentor)
	auth.POST("/mentor/register", mentorController.HandlerRegisterMentor)

	mentor := v1.Group("/mentors", authMiddleware.IsAuthenticated(), authMiddleware.IsMentor)
	mentor.GET("", mentorController.HandlerFindAll)
	mentor.PUT("/:mentorId/update-password", mentorController.HandlerUpdatePassword)
	mentor.GET("/:mentorId", mentorController.HandlerFindByCurrentMentor)
	mentor.GET("/:mentorId", mentorController.HandlerFindByID)

	// mentee routes
	// m := v1.Group("/mentees")

	//	category routes
	cat := v1.Group("/categories")
	cat.POST("", categoryController.HandlerCreateCategory)
	cat.GET("", categoryController.HandlerFindAllCategories)
	cat.GET("/:categoryId", categoryController.HandlerFindByIdCategory)
	cat.PUT("/:categoryId", categoryController.HandlerUpdateCategory)

	// course routes
	course := v1.Group("/courses")
	course.POST("", courseController.HandlerCreateCourse)
	course.GET("", courseController.HandlerFindAllCourses)
	course.GET("/categories/:categoryId", courseController.HandlerFindByCategory)
	course.GET("/:courseId", courseController.HandlerFindByIdCourse)
	course.PUT("/:courseId", courseController.HandlerUpdateCourse)
	course.DELETE("/:courseId", courseController.HandlerSoftDeleteCourse)

	// module routes
	module := v1.Group("/modules")
	module.POST("", moduleController.HandlerCreateModule)
	module.GET("/:moduleId", moduleController.HandlerFindByIdModule)
	module.PUT("/:moduleId", moduleController.HandlerUpdateModule)
	module.DELETE("/:moduleId", moduleController.HandlerDeleteModule)
}
