package response

import detailCourse "github.com/Kelompok14-LMS/backend-go/businesses/detailCourse"

func FullDetailCourseEnrolled(domain *detailCourse.Domain) *Course {
	modules := make([]Module, len(domain.Modules))

	for i, module := range domain.Modules {
		modules[i].ModuleId = module.ModuleId
		modules[i].CourseId = module.CourseId
		modules[i].Title = module.Title
		modules[i].Description = module.Description
		modules[i].CreatedAt = module.CreatedAt
		modules[i].UpdatedAt = module.UpdatedAt
	}

	assignments := make([]Assignment, len(domain.Assignments))

	for i, assignment := range domain.Assignments {
		assignments[i].AssignmentID = assignment.ID
		assignments[i].CourseId = assignment.CourseId
		assignments[i].Title = assignment.Title
		assignments[i].Description = assignment.Description
		assignments[i].CreatedAt = assignment.CreatedAt
		assignments[i].UpdatedAt = assignment.UpdatedAt
	}

	for i, module := range modules {
		module.Materials = make([]Material, len(domain.Modules[i].Materials))

		for j, material := range domain.Modules[i].Materials {
			if module.ModuleId == material.ModuleId {
				module.Materials[j].MaterialId = material.MaterialId
				module.Materials[j].ModuleId = material.ModuleId
				module.Materials[j].Title = material.Title
				module.Materials[j].URL = material.URL
				module.Materials[j].Description = material.Description
				module.Materials[j].Completed = material.Completed
				module.Materials[j].CreatedAt = material.CreatedAt
				module.Materials[j].UpdatedAt = material.UpdatedAt

				modules[i].Materials = append(modules[i].Materials, module.Materials[j])
			}
		}
	}

	return &Course{
		CourseId:    domain.CourseId,
		CategoryId:  domain.CategoryId,
		MentorId:    domain.MentorId,
		Mentor:      domain.Mentor,
		Category:    domain.Category,
		Title:       domain.Title,
		Description: domain.Description,
		Thumbnail:   domain.Thumbnail,
		Modules:     modules,
		Assignments: assignments,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}
