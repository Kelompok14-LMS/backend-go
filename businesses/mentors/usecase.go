package mentors

import (
	"time"

	"github.com/Kelompok14-LMS/backend-go/businesses/users"
	"github.com/Kelompok14-LMS/backend-go/pkg"
	"github.com/Kelompok14-LMS/backend-go/utils"
	"github.com/google/uuid"
)

type mentorUsecase struct {
	mentorsRepository Repository
	userRepository    users.Repository
	jwtConfig         *utils.JWTConfig
}

func NewMentorUsecase(mentorsRepository Repository, userRepository users.Repository, jwtConfig *utils.JWTConfig) Usecase {
	return mentorUsecase{
		mentorsRepository: mentorsRepository,
		userRepository:    userRepository,
		jwtConfig:         jwtConfig,
	}
}

func (m mentorUsecase) Register(mentorDomain *MentorRegister) error {
	var err error

	if mentorDomain.Password != mentorDomain.RepeatedPassword {
		return pkg.ErrPasswordNotMatch
	}

	userId := uuid.NewString()
	hashedPassword := utils.HashPassword(mentorDomain.RepeatedPassword)

	user := users.Domain{
		ID:        userId,
		Email:     mentorDomain.Email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = m.userRepository.Create(&user)

	if err != nil {
		return err
	}

	mentorId := uuid.NewString()

	mentor := Domain{
		ID:        mentorId,
		UserId:    userId,
		FullName:  mentorDomain.FullName,
		Role:      "mentor",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = m.mentorsRepository.Create(&mentor)

	if err != nil {
		return err
	}

	return nil
}

func (m mentorUsecase) Login(mentorAuth *MentorAuth) (*string, error) {
	if len(mentorAuth.Password) < 6 {
		return nil, pkg.ErrPasswordLengthInvalid
	}

	var err error

	user := &users.Domain{}
	user, err = m.userRepository.FindByEmail(mentorAuth.Email)

	if err != nil {
		return nil, err
	}

	ok := utils.ComparePassword(user.Password, mentorAuth.Password)
	if !ok {
		return nil, pkg.ErrUserNotFound
	}

	mentor := &Domain{}
	mentor, err = m.mentorsRepository.FindByIdUser(user.ID)

	if err != nil {
		return nil, err
	}

	var token string
	token, err = m.jwtConfig.GenerateToken(user.ID, mentor.ID, mentor.Role)

	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (m mentorUsecase) FindAll() (*[]Domain, error) {
	//TODO implement me
	panic("implement me")
}

func (m mentorUsecase) FindById(id string) (*Domain, error) {
	//TODO implement me
	panic("implement me")
}

func (m mentorUsecase) Update(id string, userDomain *Domain) error {
	//TODO implement me
	panic("implement me")
}
