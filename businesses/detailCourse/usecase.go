package detail_course

import (
	"github.com/Kelompok14-LMS/backend-go/businesses/assignments"
	"github.com/Kelompok14-LMS/backend-go/businesses/courses"
	"github.com/Kelompok14-LMS/backend-go/businesses/materials"
	menteeProgresses "github.com/Kelompok14-LMS/backend-go/businesses/menteeProgresses"
	"github.com/Kelompok14-LMS/backend-go/businesses/mentees"
	"github.com/Kelompok14-LMS/backend-go/businesses/modules"
)

type detailCourseUsecase struct {
	menteeRepository         mentees.Repository
	courseRepository         courses.Repository
	moduleRepository         modules.Repository
	materialRepository       materials.Repository
	menteeProgressRepository menteeProgresses.Repository
	assignmentsRepository    assignments.Repository
}

func NewDetailCourseUsecase(
	menteeRepository mentees.Repository,
	courseRepository courses.Repository,
	moduleRepository modules.Repository,
	materialRepository materials.Repository,
	menteeProgressRepository menteeProgresses.Repository,
	assignmentsRepository assignments.Repository,
) Usecase {
	return detailCourseUsecase{
		menteeRepository:         menteeRepository,
		courseRepository:         courseRepository,
		moduleRepository:         moduleRepository,
		materialRepository:       materialRepository,
		menteeProgressRepository: menteeProgressRepository,
		assignmentsRepository:    assignmentsRepository,
	}
}

func (dc detailCourseUsecase) DetailCourse(courseId string) (*Domain, error) {
	course, err := dc.courseRepository.FindById(courseId)

	if err != nil {
		return nil, err
	}

	assignments, err := dc.assignmentsRepository.FindByCourseId(courseId)

	if err != nil {
		return nil, err
	}

	assignmentId := []string{}

	for _, assignment := range assignments {
		assignmentId = append(assignmentId, assignment.ID)
	}

	modules, err := dc.moduleRepository.FindByCourse(courseId)

	if err != nil {
		return nil, err
	}

	moduleIds := []string{}

	for _, module := range modules {
		moduleIds = append(moduleIds, module.ID)
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
		CourseId:    course.ID,
		CategoryId:  course.CategoryId,
		MentorId:    course.MentorId,
		Mentor:      course.Mentor.Fullname,
		Category:    course.Category.Name,
		Title:       course.Title,
		Description: course.Description,
		Thumbnail:   course.Thumbnail,
		Modules:     moduleDomain,
		CreatedAt:   course.CreatedAt,
		UpdatedAt:   course.UpdatedAt,
	}

	return &courseDomain, nil
}

func (dc detailCourseUsecase) DetailCourseEnrolled(menteeId string, courseId string) (*Domain, error) {
	course, err := dc.courseRepository.FindById(courseId)

	if err != nil {
		return nil, err
	}

	assignments, err := dc.assignmentsRepository.FindByCourseId(courseId)

	if err != nil {
		return nil, err
	}

	assignmentId := []string{}

	for _, assignment := range assignments {
		assignmentId = append(assignmentId, assignment.ID)
	}

	assignmentsDomain := make([]Assignment, len(assignments))

	for i, assignment := range assignments {
		assignmentsDomain[i].ID = assignment.ID
		assignmentsDomain[i].CourseId = assignment.CourseId
		assignmentsDomain[i].Title = assignment.Title
		assignmentsDomain[i].Description = assignment.Description
		assignmentsDomain[i].CreatedAt = assignment.CreatedAt
		assignmentsDomain[i].UpdatedAt = assignment.UpdatedAt
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

	courseDomain := Domain{
		CourseId:    course.ID,
		CategoryId:  course.CategoryId,
		MentorId:    course.MentorId,
		Mentor:      course.Mentor.Fullname,
		Category:    course.Category.Name,
		Title:       course.Title,
		Description: course.Description,
		Thumbnail:   course.Thumbnail,
		Modules:     moduleDomain,
		Assignments: assignmentsDomain,
		CreatedAt:   course.CreatedAt,
		UpdatedAt:   course.UpdatedAt,
	}

	return &courseDomain, nil
}
