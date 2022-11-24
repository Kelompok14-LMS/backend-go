package users

import "time"

type Domain struct {
	ID        string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Repository interface {
	Create(userDomain *Domain) error
	FindAll() (*[]Domain, error)
	FindByEmail(email string) (*[]Domain, error)
	FindByID(id string) (*Domain, error)
	Update(id string, userDomain *Domain) error
	Delete(id string) error
}

type Usecase interface {
	Create(userDomain *Domain) error
	FindAll() (*[]Domain, error)
	FindByEmail(email string) (*[]Domain, error)
	FindByID(id string) (*Domain, error)
	Update(id string, userDomain *Domain) error
	Delete(id string) error
}
