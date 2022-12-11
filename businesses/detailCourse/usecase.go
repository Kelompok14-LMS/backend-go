package detail_course

import (
	"github.com/Kelompok14-LMS/backend-go/businesses/courses"
	"github.com/Kelompok14-LMS/backend-go/businesses/materials"
	"github.com/Kelompok14-LMS/backend-go/businesses/modules"
)

type detailCourseUsecase struct {
	courseRepository   courses.Repository
	moduleRepository   modules.Repository
	materialRepository materials.Repository
}

func NewDetailCourseUsecase(
	courseRepository courses.Repository,
	moduleRepository modules.Repository,
	materialRepository materials.Repository,
) Usecase {
	return detailCourseUsecase{
		courseRepository:   courseRepository,
		moduleRepository:   moduleRepository,
		materialRepository: materialRepository,
	}
}

func (dc detailCourseUsecase) DetailCourse(courseId string) (*Domain, error) {
	course, err := dc.courseRepository.FindById(courseId)

	if err != nil {
		return nil, err
	}

	modules, err := dc.moduleRepository.FindByCourse(courseId)

	if err != nil {
		return nil, err
	}

	moduleIds := []string{}

	for _, module := range modules {
		moduleIds = append(moduleIds, module.ID)
	}

	materials, err := dc.materialRepository.FindByModule(moduleIds)

	if err != nil {
		return nil, err
	}

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
