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

	_assignmentUsecase "github.com/Kelompok14-LMS/backend-go/businesses/assignments"
	_assignmentController "github.com/Kelompok14-LMS/backend-go/controllers/assignments"

	_materialUsecase "github.com/Kelompok14-LMS/backend-go/businesses/materials"
	_materialController "github.com/Kelompok14-LMS/backend-go/controllers/materials"

	_menteeCoursesUsecase "github.com/Kelompok14-LMS/backend-go/businesses/menteeCourses"
	_menteeCoursesController "github.com/Kelompok14-LMS/backend-go/controllers/menteeCourses"

	_menteeProgressesUsecase "github.com/Kelompok14-LMS/backend-go/businesses/menteeProgresses"
	_menteeProgressController "github.com/Kelompok14-LMS/backend-go/controllers/menteeProgresses"

	_detailCourseUsecase "github.com/Kelompok14-LMS/backend-go/businesses/detailCourse"
	_detailCourseController "github.com/Kelompok14-LMS/backend-go/controllers/detailCourse"
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
	mentorUsecase := _mentorUsecase.NewMentorUsecase(mentorRepository, userRepository, routeConfig.JWTConfig, routeConfig.StorageConfig, routeConfig.Mailer)
	mentorController := _mentorController.NewMentorController(mentorUsecase, routeConfig.JWTConfig)

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

	// Inject the dependency to assignment
	assignmentRepository := _driverFactory.NewAssignmentRepository(routeConfig.MySQLDB)
	assignmentUsecase := _assignmentUsecase.NewAssignmentUsecase(assignmentRepository, courseRepository)
	assignmentController := _assignmentController.NewAssignmentsController(assignmentUsecase)

	// Inject the dependency to material
	materialRepository := _driverFactory.NewMaterialRepository(routeConfig.MySQLDB)
	materialUsecase := _materialUsecase.NewMaterialUsecase(materialRepository, moduleRepository, routeConfig.StorageConfig)
	materialController := _materialController.NewMaterialController(materialUsecase)

	// Inject the dependency to menteeProgress
	menteeProgressRepository := _driverFactory.NewMenteeProgressRepository(routeConfig.MySQLDB)
	menteeProgressUsecase := _menteeProgressesUsecase.NewMenteeProgressUsecase(menteeProgressRepository, menteeRepository, courseRepository, materialRepository)
	menteeProgressController := _menteeProgressController.NewMenteeProgressController(menteeProgressUsecase)

	// Inject the dependency to menteeCourse
	menteeCourseRepository := _driverFactory.NewMenteeCourseRepository(routeConfig.MySQLDB)
	menteeCourseUsecase := _menteeCoursesUsecase.NewMenteeCourseUsecase(menteeCourseRepository, menteeRepository, courseRepository, materialRepository, menteeProgressRepository)
	menteeCourseController := _menteeCoursesController.NewMenteeCourseController(menteeCourseUsecase)

	detailCourseUsecase := _detailCourseUsecase.NewDetailCourseUsecase(courseRepository, moduleRepository, materialRepository, assignmentRepository)
	detailCourseController := _detailCourseController.NewDetailCourseController(detailCourseUsecase)

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
	auth.POST("/mentor/forgot-password", mentorController.HandlerForgotPassword)

	// mentor routes
	mentor := v1.Group("/mentors", authMiddleware.IsAuthenticated(), authMiddleware.IsMentor)
	mentor.GET("", mentorController.HandlerFindAll)
	mentor.GET("/profile", mentorController.HandlerProfileMentor)
	mentor.PUT("/:mentorId/update-password", mentorController.HandlerUpdatePassword)
	mentor.GET("/:mentorId", mentorController.HandlerFindByID)
	mentor.PUT("/:mentorId", mentorController.HandlerUpdateProfile)

	// mentee routes
	mentee := v1.Group("/mentees", authMiddleware.IsAuthenticated(), authMiddleware.IsMentee)
	mentee.POST("/progress", menteeProgressController.HandlerAddProgress, authMiddleware.IsAuthenticated(), authMiddleware.IsMentee)
	mentee.GET("/:menteeId/courses", menteeCourseController.HandlerFindMenteeCourses)
	mentee.GET("/:menteeId/courses/:courseId", menteeCourseController.HandlerCheckEnrollmentCourse)

	//	category routes
	cat := v1.Group("/categories")
	cat.POST("", categoryController.HandlerCreateCategory, authMiddleware.IsAuthenticated(), authMiddleware.IsMentor)
	cat.GET("", categoryController.HandlerFindAllCategories)
	cat.GET("/:categoryId", categoryController.HandlerFindByIdCategory)
	cat.PUT("/:categoryId", categoryController.HandlerUpdateCategory, authMiddleware.IsAuthenticated(), authMiddleware.IsMentor)

	// course routes
	course := v1.Group("/courses")
	course.POST("", courseController.HandlerCreateCourse, authMiddleware.IsAuthenticated(), authMiddleware.IsMentor)
	course.GET("", courseController.HandlerFindAllCourses)
	course.POST("/enroll-course", menteeCourseController.HandlerEnrollCourse, authMiddleware.IsAuthenticated())
	course.GET("/categories/:categoryId", courseController.HandlerFindByCategory)
	course.GET("/mentors/:mentorId", courseController.HandlerFindByMentor)
	course.GET("/:courseId/details", detailCourseController.HandlerDetailCourse)
	course.GET("/:courseId", courseController.HandlerFindByIdCourse)
	course.PUT("/:courseId", courseController.HandlerUpdateCourse, authMiddleware.IsAuthenticated(), authMiddleware.IsMentor)
	course.DELETE("/:courseId", courseController.HandlerSoftDeleteCourse, authMiddleware.IsAuthenticated(), authMiddleware.IsMentor)

	// module routes
	module := v1.Group("/modules")
	module.POST("", moduleController.HandlerCreateModule, authMiddleware.IsAuthenticated(), authMiddleware.IsMentor)
	module.GET("/:moduleId", moduleController.HandlerFindByIdModule)
	module.PUT("/:moduleId", moduleController.HandlerUpdateModule)
	module.DELETE("/:moduleId", moduleController.HandlerDeleteModule)
	module.PUT("/:moduleId", moduleController.HandlerUpdateModule, authMiddleware.IsAuthenticated(), authMiddleware.IsMentor)
	module.DELETE("/:moduleId", moduleController.HandlerDeleteModule, authMiddleware.IsAuthenticated(), authMiddleware.IsMentor)

	// assignment routes
	assignment := v1.Group("/assignments")
	assignment.POST("", assignmentController.HandlerCreateAssignment)
	assignment.GET("/:assignmentId", assignmentController.HandlerFindByIdAssignment)
	assignment.GET("/course/:courseid", assignmentController.HandlerFindByCourse)
	assignment.PUT("/:assignmentId", assignmentController.HandlerUpdateAssignment)
	assignment.DELETE("/:assignmentId", assignmentController.HandlerDeleteAssignment)

	// material routes
	material := v1.Group("/materials")
	material.POST("", materialController.HandlerCreateMaterial, authMiddleware.IsAuthenticated(), authMiddleware.IsMentor)
	material.DELETE("/modules/:moduleId", materialController.HandlerSoftDeleteMaterialByModule, authMiddleware.IsAuthenticated(), authMiddleware.IsMentor)
	material.GET("/:materialId", materialController.HandlerFindByIdMaterial)
	material.PUT("/:materialId", materialController.HandlerUpdateMaterial, authMiddleware.IsAuthenticated(), authMiddleware.IsMentor)
	material.DELETE("/:materialId", materialController.HandlerSoftDeleteMaterial, authMiddleware.IsAuthenticated(), authMiddleware.IsMentor)
}
