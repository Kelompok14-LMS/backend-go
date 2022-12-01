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

	if len(mentorDomain.Password) < 8 {
		return pkg.ErrPasswordLengthInvalid
	}

	email, _ := m.userRepository.FindByEmail(mentorDomain.Email)

	if email != nil {
		return pkg.ErrEmailAlreadyExist
	}

	userId := uuid.NewString()
	hashedPassword := utils.HashPassword(mentorDomain.Password)

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
	if len(mentorAuth.Password) < 8 {
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

func (m mentorUsecase) UpdatePassword(updatePassword *MentorUpdatePassword) error {

	oldPassword, err := m.userRepository.FindById(updatePassword.UserID)

	if err != nil {
		return pkg.ErrUserNotFound
	}

	ok := utils.ComparePassword(oldPassword.Password, updatePassword.OldPassword)
	if !ok {
		return pkg.ErrUserNotFound
	}

	hashPassword := utils.HashPassword(updatePassword.NewPassword)

	updatedUser := users.Domain{
		Password: hashPassword,
	}

	err = m.userRepository.Update(oldPassword.ID, &updatedUser)

	if err != nil {
		return err
	}

	return nil

}

func (m mentorUsecase) FindAll() (*[]Domain, error) {
	var err error

	mentor, err := m.mentorsRepository.FindAll()

	if err != nil {
		if err == pkg.ErrUserNotFound {
			return nil, pkg.ErrMentorNotFound
		}

		return nil, pkg.ErrInternalServerError
	}

	return mentor, nil
}

func (m mentorUsecase) FindById(id string) (*Domain, error) {

	mentor, err := m.mentorsRepository.FindById(id)
	if err != nil {
		if err == pkg.ErrUserNotFound {
			return nil, pkg.ErrMentorNotFound
		}

		return nil, pkg.ErrInternalServerError
	}

	return mentor, nil
}

func (m mentorUsecase) Update(updateMentor *MentorUpdateProfile) error {

	mentor, _ := m.userRepository.FindById(updateMentor.UserID)

	user := users.Domain{
		ID:        updateMentor.UserID,
		Email:     updateMentor.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := m.userRepository.Update(mentor.ID, &user)

	if err != nil {
		return err
	}

	mentorID, _ := m.mentorsRepository.FindById(updateMentor.ID)

	updatedMentor := Domain{

		FullName:       updateMentor.FullName,
		Phone:          updateMentor.Phone,
		Jobs:           updateMentor.Jobs,
		Gender:         updateMentor.Gender,
		BirthPlace:     updateMentor.BirthPlace,
		BirthDate:      updateMentor.BirthDate,
		Address:        updateMentor.Address,
		ProfilePicture: updateMentor.ProfilePicture,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	err = m.mentorsRepository.Update(mentorID.ID, &updatedMentor)
	if err != nil {
		if err == pkg.ErrUserNotFound {
			return pkg.ErrMentorNotFound
		}

		return pkg.ErrInternalServerError
	}

	return nil
}
