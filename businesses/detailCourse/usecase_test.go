package detail_course_test

import (
	"testing"

	"github.com/Kelompok14-LMS/backend-go/businesses/assignments"
	_assignmentMock "github.com/Kelompok14-LMS/backend-go/businesses/assignments/mocks"
	"github.com/Kelompok14-LMS/backend-go/businesses/courses"
	_courseMock "github.com/Kelompok14-LMS/backend-go/businesses/courses/mocks"
	detailCourse "github.com/Kelompok14-LMS/backend-go/businesses/detailCourse"
	"github.com/Kelompok14-LMS/backend-go/businesses/materials"
	_materialMock "github.com/Kelompok14-LMS/backend-go/businesses/materials/mocks"
	menteeAssignments "github.com/Kelompok14-LMS/backend-go/businesses/menteeAssignments"
	_menteeAssignmentMock "github.com/Kelompok14-LMS/backend-go/businesses/menteeAssignments/mocks"
	menteeCourses "github.com/Kelompok14-LMS/backend-go/businesses/menteeCourses"
	_menteeCourseMock "github.com/Kelompok14-LMS/backend-go/businesses/menteeCourses/mocks"
	menteeProgresses "github.com/Kelompok14-LMS/backend-go/businesses/menteeProgresses"
	_menteeProgressMock "github.com/Kelompok14-LMS/backend-go/businesses/menteeProgresses/mocks"
	"github.com/Kelompok14-LMS/backend-go/businesses/mentees"
	_menteeMock "github.com/Kelompok14-LMS/backend-go/businesses/mentees/mocks"
	"github.com/Kelompok14-LMS/backend-go/businesses/modules"
	_moduleMock "github.com/Kelompok14-LMS/backend-go/businesses/modules/mocks"
	"github.com/Kelompok14-LMS/backend-go/pkg"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	menteeRepository           _menteeMock.Repository
	courseRepository           _courseMock.Repository
	moduleRepository           _moduleMock.Repository
	materialRepository         _materialMock.Repository
	menteeProgressRepository   _menteeProgressMock.Repository
	assignmentRepository       _assignmentMock.Repository
	menteeAssignmentRepository _menteeAssignmentMock.Repository
	menteeCourseRepository     _menteeCourseMock.Repository

	detailCourseService detailCourse.Usecase

	menteeDomain           mentees.Domain
	courseDomain           courses.Domain
	assignmentDomain       assignments.Domain
	moduleDomain           modules.Domain
	materialDomain         materials.Domain
	menteeCourseDomain     menteeCourses.Domain
	menteeProgressDomain   menteeProgresses.Domain
	menteeAssignmentDomain menteeAssignments.Domain
)

func TestMain(m *testing.M) {
	detailCourseService = detailCourse.NewDetailCourseUsecase(
		&menteeRepository,
		&courseRepository,
		&moduleRepository,
		&materialRepository,
		&menteeProgressRepository,
		&assignmentRepository,
		&menteeAssignmentRepository,
		&menteeCourseRepository,
	)

	menteeDomain = mentees.Domain{
		ID:             uuid.NewString(),
		UserId:         uuid.NewString(),
		Fullname:       "test",
		Phone:          "test",
		Role:           "mentee",
		BirthDate:      "test",
		Address:        "test",
		ProfilePicture: "test.com",
	}

	courseDomain = courses.Domain{
		ID:           uuid.NewString(),
		MentorId:     uuid.NewString(),
		CategoryId:   uuid.NewString(),
		Title:        "test",
		Description:  "test",
		Thumbnail:    "test.com",
		TotalReviews: 100,
		Rating:       5,
	}

	assignmentDomain = assignments.Domain{
		ID:          uuid.NewString(),
		CourseId:    courseDomain.ID,
		Title:       "test",
		Description: "test",
	}

	moduleDomain = modules.Domain{
		ID:          uuid.NewString(),
		CourseId:    courseDomain.ID,
		Title:       "test",
		Description: "test",
	}

	materialDomain = materials.Domain{
		ID:          uuid.NewString(),
		CourseId:    courseDomain.ID,
		ModuleId:    moduleDomain.ID,
		Title:       "test",
		URL:         "test.com",
		Description: "test",
	}

	menteeCourseDomain = menteeCourses.Domain{
		ID:             uuid.NewString(),
		MenteeId:       menteeDomain.ID,
		CourseId:       courseDomain.ID,
		Status:         "ongoing",
		Reviewed:       false,
		ProgressCount:  8,
		TotalMaterials: 10,
	}

	menteeProgressDomain = menteeProgresses.Domain{
		ID:         uuid.NewString(),
		MenteeId:   menteeDomain.ID,
		CourseId:   courseDomain.ID,
		MaterialId: materialDomain.ID,
		Completed:  false,
	}

	menteeAssignmentDomain = menteeAssignments.Domain{
		ID:             uuid.NewString(),
		MenteeId:       menteeDomain.ID,
		AssignmentId:   assignmentDomain.ID,
		Name:           menteeDomain.Fullname,
		ProfilePicture: menteeDomain.ProfilePicture,
		AssignmentURL:  "test.com",
		Grade:          0,
		Completed:      false,
	}

	m.Run()
}

func TestDetailCourse(t *testing.T) {
	t.Run("Test Detail Course | Success get detail course", func(t *testing.T) {
		courseRepository.Mock.On("FindById", courseDomain.ID).Return(&courseDomain, nil).Once()

		moduleRepository.Mock.On("FindByCourse", courseDomain.ID).Return([]modules.Domain{moduleDomain}, nil).Once()

		assignmentRepository.Mock.On("FindByCourseId", courseDomain.ID).Return(&assignmentDomain, nil).Once()

		materialRepository.Mock.On("FindByModule", []string{moduleDomain.ID}).Return([]materials.Domain{materialDomain}, nil).Once()

		result, err := detailCourseService.DetailCourse(courseDomain.ID)

		assert.NoError(t, err)
		assert.NotEmpty(t, result)
	})

	t.Run("Test Detail Course | Failed get detail course | Course not found", func(t *testing.T) {
		courseRepository.Mock.On("FindById", courseDomain.ID).Return(nil, pkg.ErrCourseNotFound).Once()

		result, err := detailCourseService.DetailCourse(courseDomain.ID)

		assert.Error(t, err)
		assert.Empty(t, result)
	})
}

func TestDetailCourseEnrolled(t *testing.T) {
	t.Run("Test Detail Course Enrolled | Success get detail course enrolled", func(t *testing.T) {
		menteeCourseRepository.Mock.On("CheckEnrollment", menteeDomain.ID, courseDomain.ID).Return(&menteeCourseDomain, nil).Once()

		courseRepository.Mock.On("FindById", courseDomain.ID).Return(&courseDomain, nil).Once()

		menteeRepository.Mock.On("FindById", menteeDomain.ID).Return(&menteeDomain, nil).Once()

		assignmentRepository.Mock.On("FindByCourseId", courseDomain.ID).Return(&assignmentDomain, nil).Once()

		menteeAssignmentRepository.Mock.On("FindByCourse", menteeDomain.ID, courseDomain.ID).Return(&menteeAssignmentDomain, nil).Once()

		moduleRepository.Mock.On("FindByCourse", courseDomain.ID).Return([]modules.Domain{moduleDomain}, nil).Once()

		materialRepository.Mock.On("FindByModule", []string{moduleDomain.ID}).Return([]materials.Domain{materialDomain}, nil).Once()

		menteeProgressRepository.Mock.On("FindByMentee", menteeDomain.ID, courseDomain.ID).Return([]menteeProgresses.Domain{menteeProgressDomain}, nil).Once()

		materialRepository.Mock.On("CountByCourse", []string{courseDomain.ID}).Return([]int64{menteeCourseDomain.TotalMaterials}, nil).Once()

		menteeProgressRepository.Mock.On("Count", menteeDomain.ID, courseDomain.Title, menteeCourseDomain.Status).Return([]int64{menteeCourseDomain.ProgressCount}, nil).Once()

		result, err := detailCourseService.DetailCourseEnrolled(menteeDomain.ID, courseDomain.ID)

		assert.NoError(t, err)
		assert.NotEmpty(t, result)
	})

	t.Run("Test Detail Course Enrolled | Failed get detail course enrolled | Course enrollment not found", func(t *testing.T) {
		menteeCourseRepository.Mock.On("CheckEnrollment", menteeDomain.ID, courseDomain.ID).Return(nil, pkg.ErrNoEnrolled).Once()

		result, err := detailCourseService.DetailCourseEnrolled(menteeDomain.ID, courseDomain.ID)

		assert.Error(t, err)
		assert.Empty(t, result)
	})

	t.Run("Test Detail Course Enrolled | Failed get detail course enrolled | Course not found", func(t *testing.T) {
		menteeCourseRepository.Mock.On("CheckEnrollment", menteeDomain.ID, courseDomain.ID).Return(&menteeCourseDomain, nil).Once()

		courseRepository.Mock.On("FindById", courseDomain.ID).Return(nil, pkg.ErrCourseNotFound).Once()

		result, err := detailCourseService.DetailCourseEnrolled(menteeDomain.ID, courseDomain.ID)

		assert.Error(t, err)
		assert.Empty(t, result)
	})

	t.Run("Test Detail Course Enrolled | Failed get detail course enrolled | Mentee not found", func(t *testing.T) {
		menteeCourseRepository.Mock.On("CheckEnrollment", menteeDomain.ID, courseDomain.ID).Return(&menteeCourseDomain, nil).Once()

		courseRepository.Mock.On("FindById", courseDomain.ID).Return(&courseDomain, nil).Once()

		menteeRepository.Mock.On("FindById", menteeDomain.ID).Return(nil, pkg.ErrMenteeNotFound).Once()

		result, err := detailCourseService.DetailCourseEnrolled(menteeDomain.ID, courseDomain.ID)

		assert.Error(t, err)
		assert.Empty(t, result)
	})
}
