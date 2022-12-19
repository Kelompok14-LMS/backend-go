package detail_course

import (
	"github.com/Kelompok14-LMS/backend-go/businesses/assignments"
	"github.com/Kelompok14-LMS/backend-go/businesses/courses"
	"github.com/Kelompok14-LMS/backend-go/businesses/materials"
	menteeAssignments "github.com/Kelompok14-LMS/backend-go/businesses/menteeAssignments"
	menteeCourses "github.com/Kelompok14-LMS/backend-go/businesses/menteeCourses"
	menteeProgresses "github.com/Kelompok14-LMS/backend-go/businesses/menteeProgresses"
	"github.com/Kelompok14-LMS/backend-go/businesses/mentees"
	"github.com/Kelompok14-LMS/backend-go/businesses/modules"
)

type detailCourseUsecase struct {
	menteeRepository           mentees.Repository
	courseRepository           courses.Repository
	moduleRepository           modules.Repository
	materialRepository         materials.Repository
	menteeProgressRepository   menteeProgresses.Repository
	assignmentsRepository      assignments.Repository
	menteeAssignmentRepository menteeAssignments.Repository
	menteeCourse               menteeCourses.Repository
}

func NewDetailCourseUsecase(
	menteeRepository mentees.Repository,
	courseRepository courses.Repository,
	moduleRepository modules.Repository,
	materialRepository materials.Repository,
	menteeProgressRepository menteeProgresses.Repository,
	assignmentsRepository assignments.Repository,
	menteeAssignmentRepository menteeAssignments.Repository,
	menteeCourse menteeCourses.Repository,
) Usecase {
	return detailCourseUsecase{
		menteeRepository:           menteeRepository,
		courseRepository:           courseRepository,
		moduleRepository:           moduleRepository,
		materialRepository:         materialRepository,
		menteeProgressRepository:   menteeProgressRepository,
		assignmentsRepository:      assignmentsRepository,
		menteeAssignmentRepository: menteeAssignmentRepository,
		menteeCourse:               menteeCourse,
	}
}

func (dc detailCourseUsecase) DetailCourse(courseId string) (*Domain, error) {
	course, err := dc.courseRepository.FindById(courseId)

	if err != nil {
		return nil, err
	}

	modules, _ := dc.moduleRepository.FindByCourse(courseId)

	moduleIds := []string{}

	for _, module := range modules {
		moduleIds = append(moduleIds, module.ID)
	}

	assignment, _ := dc.assignmentsRepository.FindByCourseId(courseId)

	assignmentDomain := Assignment{
		ID:          assignment.ID,
		CourseId:    assignment.CourseId,
		Title:       assignment.Title,
		Description: assignment.Description,
		CreatedAt:   assignment.CreatedAt,
		UpdatedAt:   assignment.UpdatedAt,
	}

	materials, _ := dc.materialRepository.FindByModule(moduleIds)

	materialDomain := make([]Material, len(materials))

	for i, material := range materials {
		materialDomain[i].MaterialId = material.ID
		materialDomain[i].ModuleId = material.ModuleId
		materialDomain[i].Title = material.Title
		materialDomain[i].URL = material.URL
		materialDomain[i].Description = material.Description
		materialDomain[i].CreatedAt = material.CreatedAt
		materialDomain[i].UpdatedAt = material.UpdatedAt
	}

	moduleDomain := make([]Module, len(modules))

	for i, module := range modules {
		moduleDomain[i].ModuleId = module.ID
		moduleDomain[i].CourseId = module.CourseId
		moduleDomain[i].Title = module.Title
		moduleDomain[i].Description = module.Description
		moduleDomain[i].CreatedAt = module.CreatedAt
		moduleDomain[i].UpdatedAt = module.UpdatedAt
	}

	for i, module := range moduleDomain {
		for _, material := range materialDomain {
			if module.ModuleId == material.ModuleId {
				moduleDomain[i].Materials = append(moduleDomain[i].Materials, material)
			}
		}
	}

	courseDomain := Domain{
		CourseId:     course.ID,
		CategoryId:   course.CategoryId,
		MentorId:     course.MentorId,
		Mentor:       course.Mentor.Fullname,
		Category:     course.Category.Name,
		Title:        course.Title,
		Description:  course.Description,
		Thumbnail:    course.Thumbnail,
		TotalReviews: course.TotalReviews,
		Rating:       course.Rating,
		Modules:      moduleDomain,
		Assignment:   assignmentDomain,
		CreatedAt:    course.CreatedAt,
		UpdatedAt:    course.UpdatedAt,
	}

	return &courseDomain, nil
}

func (dc detailCourseUsecase) DetailCourseEnrolled(menteeId string, courseId string) (*Domain, error) {
	menteeCourse, err := dc.menteeCourse.CheckEnrollment(menteeId, courseId)

	if err != nil {
		return nil, err
	}

	course, err := dc.courseRepository.FindById(courseId)

	if err != nil {
		return nil, err
	}

	assignment, _ := dc.assignmentsRepository.FindByCourseId(courseId)

	menteeAssignment, _ := dc.menteeAssignmentRepository.FindByCourse(menteeId, courseId)

	isCompletingAssignment := menteeAssignment != nil

	assignmentDomain := Assignment{
		ID:          assignment.ID,
		CourseId:    assignment.CourseId,
		Title:       assignment.Title,
		Description: assignment.Description,
		Completed:   isCompletingAssignment,
		CreatedAt:   assignment.CreatedAt,
		UpdatedAt:   assignment.UpdatedAt,
	}

	if _, err := dc.menteeRepository.FindById(menteeId); err != nil {
		return nil, err
	}

	modules, _ := dc.moduleRepository.FindByCourse(courseId)

	modulesIds := []string{}

	for _, module := range modules {
		modulesIds = append(modulesIds, module.ID)
	}

	materials, _ := dc.materialRepository.FindByModule(modulesIds)

	materialDomain := make([]Material, len(materials))

	for i, material := range materials {
		materialDomain[i].MaterialId = material.ID
		materialDomain[i].ModuleId = material.ModuleId
		materialDomain[i].Title = material.Title
		materialDomain[i].URL = material.URL
		materialDomain[i].Description = material.Description
		materialDomain[i].CreatedAt = material.CreatedAt
		materialDomain[i].UpdatedAt = material.UpdatedAt
	}

	progresses, _ := dc.menteeProgressRepository.FindByMentee(menteeId, courseId)

	for i := range progresses {
		if progresses[i].Completed {
			materialDomain[i].Completed = true
		}
	}

	moduleDomain := make([]Module, len(modules))

	for i, module := range modules {
		moduleDomain[i].ModuleId = module.ID
		moduleDomain[i].CourseId = module.CourseId
		moduleDomain[i].Title = module.Title
		moduleDomain[i].Description = module.Description
		moduleDomain[i].CreatedAt = module.CreatedAt
		moduleDomain[i].UpdatedAt = module.UpdatedAt
	}

	for i, module := range moduleDomain {
		for _, material := range materialDomain {
			if module.ModuleId == material.ModuleId {
				moduleDomain[i].Materials = append(moduleDomain[i].Materials, material)
			}
		}
	}

	totalMaterialsArray, _ := dc.materialRepository.CountByCourse([]string{courseId})

	var totalMaterials int64

	if len(totalMaterialsArray) != 0 {
		totalMaterials = totalMaterialsArray[0]
	}

	if assignment != nil {
		totalMaterials += 1
	}

	progressArray, _ := dc.menteeProgressRepository.Count(menteeId, course.Title, menteeCourse.Status)

	var progress int64

	if len(progressArray) != 0 {
		progress = progressArray[0]
	}

	if isCompletingAssignment {
		totalMaterials += 1
	}

	courseDomain := Domain{
		CourseId:       course.ID,
		CategoryId:     course.CategoryId,
		MentorId:       course.MentorId,
		Mentor:         course.Mentor.Fullname,
		Category:       course.Category.Name,
		Title:          course.Title,
		Description:    course.Description,
		Thumbnail:      course.Thumbnail,
		TotalReviews:   course.TotalReviews,
		Rating:         course.Rating,
		Progress:       progress,
		TotalMaterials: totalMaterials,
		Modules:        moduleDomain,
		Assignment:     assignmentDomain,
		CreatedAt:      course.CreatedAt,
		UpdatedAt:      course.UpdatedAt,
	}

	return &courseDomain, nil
}
