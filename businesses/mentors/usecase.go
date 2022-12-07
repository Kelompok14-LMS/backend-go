package mentors

import (
	"context"
	"fmt"
	"time"

	"github.com/Kelompok14-LMS/backend-go/businesses/users"
	"github.com/Kelompok14-LMS/backend-go/helper"
	"github.com/Kelompok14-LMS/backend-go/pkg"
	"github.com/Kelompok14-LMS/backend-go/utils"
	"github.com/google/uuid"
)

type mentorUsecase struct {
	mentorsRepository Repository
	userRepository    users.Repository
	jwtConfig         *utils.JWTConfig
	storage           *helper.StorageConfig
	mailerConfig      *pkg.MailerConfig
}

func NewMentorUsecase(mentorsRepository Repository, userRepository users.Repository, jwtConfig *utils.JWTConfig, storage *helper.StorageConfig, mailerConfig *pkg.MailerConfig) Usecase {
	return mentorUsecase{
		mentorsRepository: mentorsRepository,
		userRepository:    userRepository,
		jwtConfig:         jwtConfig,
		storage:           storage,
		mailerConfig:      mailerConfig,
	}
}

func (m mentorUsecase) Register(mentorDomain *MentorRegister) error {
	var err error

	if len(mentorDomain.Password) < 6 {
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
		Fullname:  mentorDomain.Fullname,
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

	var user *users.Domain
	user, err = m.userRepository.FindByEmail(mentorAuth.Email)

	if err != nil {
		return nil, err
	}

	ok := utils.ComparePassword(user.Password, mentorAuth.Password)
	if !ok {
		return nil, pkg.ErrUserNotFound
	}

	var mentor *Domain
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

	if len(updatePassword.NewPassword) < 8 {
		return pkg.ErrPasswordLengthInvalid
	}

	oldPassword, err := m.userRepository.FindById(updatePassword.UserID)

	if err != nil {
		return pkg.ErrUserNotFound
	}

	ok := utils.ComparePassword(oldPassword.Password, updatePassword.OldPassword)
	if !ok {
		return pkg.ErrPasswordNotMatch
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
		if err == pkg.ErrMentorNotFound {
			return nil, pkg.ErrMentorNotFound
		}

		return nil, pkg.ErrInternalServerError
	}

	return mentor, nil
}

func (m mentorUsecase) FindById(id string) (*Domain, error) {

	mentor, err := m.mentorsRepository.FindById(id)
	if err != nil {
		if err == pkg.ErrMentorNotFound {
			return nil, pkg.ErrMentorNotFound
		}

		return nil, pkg.ErrInternalServerError
	}

	return mentor, nil
}

func (m mentorUsecase) Update(updateMentor *MentorUpdateProfile) error {

	_, err := m.userRepository.FindById(updateMentor.UserID)

	if err != nil {
		return err
	}
	user := users.Domain{
		ID:        updateMentor.UserID,
		Email:     updateMentor.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = m.userRepository.Update(user.ID, &user)
	if err != nil {
		return err
	}

	mentor, err := m.mentorsRepository.FindById(updateMentor.ID)

	if err != nil {
		return err
	}

	var ProfilePictureURL string

	if updateMentor.ProfilePictureFile != nil {
		ctx := context.Background()

		if err := m.storage.DeleteObject(ctx, mentor.ProfilePicture); err != nil {
			return err
		}

		filename := updateMentor.ProfilePictureFile.Filename

		ProfilePicture, err := updateMentor.ProfilePictureFile.Open()

		if err != nil {
			return err
		}

		defer ProfilePicture.Close()

		ProfilePictureURL, err = m.storage.UploadImage(ctx, filename, ProfilePicture)

		if err != nil {
			return err
		}
	}

	updatedMentor := Domain{

		Fullname:       updateMentor.Fullname,
		Phone:          updateMentor.Phone,
		Jobs:           updateMentor.Jobs,
		Gender:         updateMentor.Gender,
		BirthPlace:     updateMentor.BirthPlace,
		BirthDate:      updateMentor.BirthDate,
		Address:        updateMentor.Address,
		ProfilePicture: ProfilePictureURL,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	err = m.mentorsRepository.Update(updateMentor.ID, &updatedMentor)
	if err != nil {
		if err == pkg.ErrMentorNotFound {
			return pkg.ErrMentorNotFound
		}

		return pkg.ErrInternalServerError
	}

	return nil
}

func (m mentorUsecase) ForgotPassword(forgotPassword *MentorForgotPassword) error {
	var err error

	user, err := m.userRepository.FindByEmail(forgotPassword.Email)

	if err != nil {
		return err
	}

	randomPassword := pkg.GenerateOTP(8)
	hashPassword := utils.HashPassword(randomPassword)

	updatedUser := users.Domain{
		Password: hashPassword,
	}

	err = m.userRepository.Update(user.ID, &updatedUser)

	if err != nil {
		return err
	}

	message := fmt.Sprintf("Password baru anda: %s", randomPassword)
	subject := "Reset Password Eduworld"

	_ = m.mailerConfig.SendMail(user.Email, subject, message)

	return nil
}
