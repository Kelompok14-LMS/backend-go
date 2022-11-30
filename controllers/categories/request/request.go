package request

import (
	"github.com/Kelompok14-LMS/backend-go/businesses/categories"
	"github.com/go-playground/validator/v10"
)

type Category struct {
	Name string `json:"name" form:"name" validate:"required"`
}

func (req *Category) ToDomain() *categories.Domain {
	return &categories.Domain{
		Name: req.Name,
	}
}

func (req *Category) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	if err != nil {
		return err
	}

	return nil
}
