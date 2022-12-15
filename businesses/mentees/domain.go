package mentees

import (
	"mime/multipart"
	"time"

	"github.com/Kelompok14-LMS/backend-go/businesses/users"
)

type Domain struct {
	ID                 string
	UserId             string
	Fullname           string
	Phone              string
	Role               string
	BirthDate          time.Time
	Address            string
	ProfilePicture     string
	ProfilePictureFile *multipart.FileHeader
	User               users.Domain
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

type MenteeAuth struct {
	Email    string
	Password string
}

type MenteeRegister struct {
	Fullname string
	Phone    string
	Email    string
	Password string
	OTP      string
}

type MenteeForgotPassword struct {
	Email            string
	Password         string
	RepeatedPassword string
	OTP              string
}

type Repository interface {
	// Create repository create mentee
	Create(menteeDomain *Domain) error

	// FindAll repository find all mentees
	FindAll() (*[]Domain, error)

	// FindById repository find mentee by id
	FindById(id string) (*Domain, error)

	// FindByIdUser repository find mentee by id user
	FindByIdUser(userId string) (*Domain, error)

	// repository find mentees by course
	FindByCourse(courseId string) (*[]Domain, error)

	// repository count total mentees by course
	CountByCourse(courseId string) (int64, error)

	// Update repository edit data mentee
	Update(id string, menteeDomain *Domain) error
}

type Usecase interface {
	// Register usecase mentee register
	Register(menteeAuth *MenteeAuth) error

	// VerifyRegister usecase verify register mentee
	VerifyRegister(menteeRegister *MenteeRegister) error

	// ForgotPassword usecase mentee verify forgot password
	ForgotPassword(forgotPassword *MenteeForgotPassword) error

	// Login usecase mentee login
	Login(menteeAuth *MenteeAuth) (interface{}, error)

	// FindAll usecase find all mentees
	FindAll() (*[]Domain, error)

	// FindById usecase find by id mentee
	FindById(id string) (*Domain, error)

	// usecase find mentee profile
	// MenteeProfile(menteeId string) (*Domain, error)

	// usecase find mentees by course
	FindByCourse(courseId string) (map[string]interface{}, error)

	// Update usecase edit data mentee
	Update(id string, menteeDomain *Domain) error
}
