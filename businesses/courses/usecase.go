package courses

import (
	"context"
	"mime/multipart"

	"github.com/Kelompok14-LMS/backend-go/businesses/categories"
	"github.com/Kelompok14-LMS/backend-go/businesses/mentors"
	"github.com/Kelompok14-LMS/backend-go/helper"
	"github.com/google/uuid"
)

type courseUsecase struct {
	courseRepository   Repository
	mentorRepository   mentors.Repository
	categoryRepository categories.Repository
	storage            *helper.StorageConfig
}

func NewCourseUsecase(
	courseRepository Repository,
	mentorRepository mentors.Repository,
	categoryRepository categories.Repository,
	storage *helper.StorageConfig,
) Usecase {
	return courseUsecase{
		courseRepository:   courseRepository,
		mentorRepository:   mentorRepository,
		categoryRepository: categoryRepository,
		storage:            storage,
	}
}

func (cu courseUsecase) Create(courseDomain *Domain) error {
	if _, err := cu.mentorRepository.FindById(courseDomain.MentorId); err != nil {
		return err
	}

	if _, err := cu.categoryRepository.FindById(courseDomain.CategoryId); err != nil {
		return err
	}

	var err error

	ctx := context.Background()
	filename := courseDomain.ThumbnailFileHeader.Filename

	var thumbnailFile multipart.File
	thumbnailFile, err = courseDomain.ThumbnailFileHeader.Open()

	if err != nil {
		return err
	}

	defer thumbnailFile.Close()

	var thumbnailUrl string
	thumbnailUrl, err = cu.storage.UploadImage(ctx, filename, thumbnailFile)

	if err != nil {
		return err
	}

	id := uuid.NewString()

	course := Domain{
		ID:          id,
		MentorId:    courseDomain.MentorId,
		CategoryId:  courseDomain.CategoryId,
		Title:       courseDomain.Title,
		Description: courseDomain.Description,
		Thumbnail:   thumbnailUrl,
	}

	err = cu.courseRepository.Create(&course)

	if err != nil {
		return err
	}

	return nil
}

func (cu courseUsecase) FindAll(keyword string) (*[]Domain, error) {
	courses, err := cu.courseRepository.FindAll(keyword)

	if err != nil {
		return nil, err
	}

	return courses, nil
}

func (cu courseUsecase) FindById(id string) (*Domain, error) {
	course, err := cu.courseRepository.FindById(id)

	if err != nil {
		return nil, err
	}

	return course, nil
}

func (cu courseUsecase) FindByCategory(categoryId string) (*[]Domain, error) {
	if _, err := cu.categoryRepository.FindById(categoryId); err != nil {
		return nil, err
	}

	courses, err := cu.courseRepository.FindByCategory(categoryId)

	if err != nil {
		return nil, err
	}

	return courses, nil
}

func (cu courseUsecase) Update(id string, courseDomain *Domain) error {
	if _, err := cu.categoryRepository.FindById(courseDomain.CategoryId); err != nil {
		return err
	}

	var err error

	course := &Domain{}
	course, err = cu.courseRepository.FindById(id)

	if err != nil {
		return err
	}

	var thumbnailUrl string

	// check if user update the image, do the process
	if courseDomain.ThumbnailFileHeader != nil {
		ctx := context.Background()

		if err := cu.storage.DeleteObject(ctx, course.Thumbnail); err != nil {
			return err
		}

		filename := courseDomain.ThumbnailFileHeader.Filename

		var thumbnailFile multipart.File
		thumbnailFile, err = courseDomain.ThumbnailFileHeader.Open()

		if err != nil {
			return err
		}

		thumbnailUrl, err = cu.storage.UploadImage(ctx, filename, thumbnailFile)

		if err != nil {
			return err
		}
	}

	updatedCourse := Domain{
		CategoryId:  courseDomain.CategoryId,
		Title:       courseDomain.Title,
		Description: courseDomain.Description,
		Thumbnail:   thumbnailUrl,
	}

	err = cu.courseRepository.Update(id, &updatedCourse)

	if err != nil {
		return err
	}

	return nil
}

func (cu courseUsecase) Delete(id string) error {
	if _, err := cu.courseRepository.FindById(id); err != nil {
		return err
	}

	err := cu.courseRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
