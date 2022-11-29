package routes

import (
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	_driverFactory "github.com/Kelompok14-LMS/backend-go/drivers"
	"github.com/Kelompok14-LMS/backend-go/pkg"
	"github.com/Kelompok14-LMS/backend-go/utils"

	_menteeUsecase "github.com/Kelompok14-LMS/backend-go/businesses/mentees"
	_menteeController "github.com/Kelompok14-LMS/backend-go/controllers/mentees"

	_otpUsecase "github.com/Kelompok14-LMS/backend-go/businesses/otp"
	_otpController "github.com/Kelompok14-LMS/backend-go/controllers/otp"

	_categoryUsecase "github.com/Kelompok14-LMS/backend-go/businesses/categories"
	_categoryController "github.com/Kelompok14-LMS/backend-go/controllers/categories"
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
}

func (routeConfig *RouteConfig) New() {
	// setup api v1
	v1 := routeConfig.Echo.Group("/api/v1")

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

	// Inject the dependency to category
	categoryRepository := _driverFactory.NewCategoryRepository(routeConfig.MySQLDB)
	categoryUsecase := _categoryUsecase.NewCategoryUsecase(categoryRepository)
	categoryController := _categoryController.NewCategoryController(categoryUsecase)

	// authentication routes
	auth := v1.Group("/auth")
	auth.POST("/mentee/login", menteeController.HandlerLoginMentee)
	auth.POST("/mentee/register", menteeController.HandlerRegisterMentee)
	auth.POST("/mentee/register/verify", menteeController.HandlerVerifyRegisterMentee)
	auth.POST("/forgot-password", menteeController.HandlerForgotPassword)
	auth.POST("/send-otp", otpController.HandlerSendOTP)
	auth.POST("/check-otp", otpController.HandlerCheckOTP)

	// mentee routes
	// m := v1.Group("/mentees")

	//	category routes
	cat := v1.Group("/categories")
	cat.POST("", categoryController.HandlerCreateCategory)
	cat.GET("", categoryController.HandlerFindAllCategories)
	cat.GET("/:categoryId", categoryController.HandlerFindByIdCategory)
	cat.PUT("/:categoryId", categoryController.HandlerUpdateCategory)
}
