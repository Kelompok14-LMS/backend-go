package detail_course_test

import (
	"github.com/Kelompok14-LMS/backend-go/businesses/assignments"
	_assignmentMock "github.com/Kelompok14-LMS/backend-go/businesses/assignments/mocks"
	"github.com/Kelompok14-LMS/backend-go/businesses/courses"
	_courseMock "github.com/Kelompok14-LMS/backend-go/businesses/courses/mocks"
	detailCourse "github.com/Kelompok14-LMS/backend-go/businesses/detailCourse"
	_materialMock "github.com/Kelompok14-LMS/backend-go/businesses/materials/mocks"
	menteeAssignments "github.com/Kelompok14-LMS/backend-go/businesses/menteeAssignments"
	_menteeAssignmentMock "github.com/Kelompok14-LMS/backend-go/businesses/menteeAssignments/mocks"
	menteeCourses "github.com/Kelompok14-LMS/backend-go/businesses/menteeCourses"
	_menteeCourseMock "github.com/Kelompok14-LMS/backend-go/businesses/menteeCourses/mocks"
	_menteeProgressMock "github.com/Kelompok14-LMS/backend-go/businesses/menteeProgresses/mocks"
	"github.com/Kelompok14-LMS/backend-go/businesses/mentees"
	_menteeMock "github.com/Kelompok14-LMS/backend-go/businesses/mentees/mocks"
	_moduleMock "github.com/Kelompok14-LMS/backend-go/businesses/modules/mocks"
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

	menteeCourseDomain     menteeCourses.Domain
	courseDomain           courses.Domain
	menteeDomain           mentees.Domain
	assignmentDomain       assignments.Domain
	menteeAssignmentDomain menteeAssignments.Domain
)
