package courses

import (
	"errors"
	"net/http"

	"github.com/Kelompok14-LMS/backend-go/businesses/courses"
	"github.com/Kelompok14-LMS/backend-go/controllers/courses/request"
	"github.com/Kelompok14-LMS/backend-go/controllers/courses/response"
	"github.com/Kelompok14-LMS/backend-go/helper"
	"github.com/Kelompok14-LMS/backend-go/pkg"
	"github.com/labstack/echo/v4"
)

type CourseController struct {
	courseUsecase courses.Usecase
}

func NewCourseController(courseUsecase courses.Usecase) *CourseController {
	return &CourseController{
		courseUsecase: courseUsecase,
	}
}

func (ctrl *CourseController) HandlerCreateCourse(c echo.Context) error {
	courseInput := request.CreateCourseInput{}

	courseInput.Thumbnail, _ = c.FormFile("thumbnail")

	if err := c.Bind(&courseInput); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrInvalidRequest.Error()))
	}

	if err := courseInput.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(err.Error()))
	}

	err := ctrl.courseUsecase.Create(courseInput.ToDomain())

	if err != nil {
		if errors.Is(err, pkg.ErrMentorNotFound) {
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrMentorNotFound.Error()))
		} else if errors.Is(err, pkg.ErrCategoryNotFound) {
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrCategoryNotFound.Error()))
		} else if errors.Is(err, pkg.ErrUnsupportedImageFile) {
			return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrUnsupportedImageFile.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(pkg.ErrInternalServerError.Error()))
		}
	}

	return c.JSON(http.StatusCreated, helper.SuccessCreatedResponse("Sukses menambahkan kursus baru", nil))
}

func (ctrl *CourseController) HandlerFindAllCourses(c echo.Context) error {
	title := c.QueryParam("keyword")

	coursesDomain, err := ctrl.courseUsecase.FindAll(title)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(pkg.ErrInternalServerError.Error()))
	}

	var courseResponse []response.FindCourses

	for _, course := range *coursesDomain {
		courseResponse = append(courseResponse, response.AllCourses(&course))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Sukses get semua kursus", courseResponse))
}

func (ctrl *CourseController) HandlerFindByIdCourse(c echo.Context) error {
	courseId := c.Param("courseId")

	courseDomain, err := ctrl.courseUsecase.FindById(courseId)

	if err != nil {
		if errors.Is(err, pkg.ErrCourseNotFound) {
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrCourseNotFound.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(pkg.ErrInternalServerError.Error()))
		}

	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Sukses get kursus berdasarkan id", response.DetailCourse(courseDomain)))
}

func (ctrl *CourseController) HandlerFindByCategory(c echo.Context) error {
	categoryId := c.Param("categoryId")

	coursesDomain, err := ctrl.courseUsecase.FindByCategory(categoryId)

	if err != nil {
		if errors.Is(err, pkg.ErrCategoryNotFound) {
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrCategoryNotFound.Error()))
		} else if errors.Is(err, pkg.ErrCourseNotFound) {
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrCourseNotFound.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(pkg.ErrInternalServerError.Error()))
		}
	}

	var coursesResponse []response.FindCourses

	for _, course := range *coursesDomain {
		coursesResponse = append(coursesResponse, response.AllCourses(&course))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Sukses get kursus berdasarkan kategori", coursesResponse))
}

func (ctrl *CourseController) HandlerFindByMentor(c echo.Context) error {
	mentorId := c.Param("mentorId")

	coursesDomain, err := ctrl.courseUsecase.FindByMentor(mentorId)

	if err != nil {
		if errors.Is(err, pkg.ErrCategoryNotFound) {
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrMentorNotFound.Error()))
		} else if errors.Is(err, pkg.ErrCourseNotFound) {
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrCourseNotFound.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(pkg.ErrInternalServerError.Error()))
		}
	}

	var coursesResponse []response.FindCourses

	for _, course := range *coursesDomain {
		coursesResponse = append(coursesResponse, response.AllCourses(&course))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Sukses get kursus berdasarkan mentor", coursesResponse))
}

func (ctrl *CourseController) HandlerFindByPopular(c echo.Context) error {
	coursesDomain, err := ctrl.courseUsecase.FindByPopular()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(pkg.ErrInternalServerError.Error()))
	}

	var coursesResponse []response.FindCourses

	for _, course := range coursesDomain {
		coursesResponse = append(coursesResponse, response.AllCourses(&course))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Sukses get kursus berdasarkan paling populer", coursesResponse))
}

func (ctrl *CourseController) HandlerUpdateCourse(c echo.Context) error {
	courseId := c.Param("courseId")
	courseInput := request.UpdateCourseInput{}

	// get the thumbnail object file from request body
	thumbnail, _ := c.FormFile("thumbnail")

	if thumbnail != nil {
		courseInput.Thumbnail = thumbnail

		if err := c.Bind(&courseInput); err != nil {
			return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrInvalidRequest.Error()))
		}
	} else {
		courseInput.CategoryId = c.FormValue("category_id")
		courseInput.Title = c.FormValue("title")
		courseInput.Description = c.FormValue("description")
		courseInput.Thumbnail = nil
	}

	if err := courseInput.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(err.Error()))
	}

	err := ctrl.courseUsecase.Update(courseId, courseInput.ToDomain())

	if err != nil {
		if errors.Is(err, pkg.ErrCategoryNotFound) {
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrCategoryNotFound.Error()))
		} else if errors.Is(err, pkg.ErrCourseNotFound) {
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrCourseNotFound.Error()))
		} else if errors.Is(err, pkg.ErrUnsupportedImageFile) {
			return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(pkg.ErrUnsupportedImageFile.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(pkg.ErrInternalServerError.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Sukses update kursus", nil))
}

func (ctrl *CourseController) HandlerSoftDeleteCourse(c echo.Context) error {
	courseId := c.Param("courseId")

	err := ctrl.courseUsecase.Delete(courseId)

	if err != nil {
		if errors.Is(err, pkg.ErrCourseNotFound) {
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(pkg.ErrCourseNotFound.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(pkg.ErrInternalServerError.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Kursus dihapus", nil))
}
