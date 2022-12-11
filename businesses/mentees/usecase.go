package mentees

import (
	"context"
	"fmt"
	"time"

	"github.com/Kelompok14-LMS/backend-go/businesses/otp"
	"github.com/Kelompok14-LMS/backend-go/businesses/users"
	"github.com/Kelompok14-LMS/backend-go/constants"
	"github.com/Kelompok14-LMS/backend-go/pkg"
	"github.com/Kelompok14-LMS/backend-go/utils"
	"github.com/google/uuid"
)

type menteeUsecase struct {
	menteeRepository Repository
	userRepository   users.Repository
	otpRepository    otp.Repository
	jwtConfig        *utils.JWTConfig
	mailerConfig     *pkg.MailerConfig
}

func NewMenteeUsecase(
	menteeRepository Repository,
	userRepository users.Repository,
	otpRepository otp.Repository,
	jwtConfig *utils.JWTConfig,
	mailerConfig *pkg.MailerConfig,
) Usecase {
	return menteeUsecase{
		menteeRepository: menteeRepository,
		userRepository:   userRepository,
		otpRepository:    otpRepository,
		jwtConfig:        jwtConfig,
		mailerConfig:     mailerConfig,
	}
}

func (m menteeUsecase) Register(menteeAuth *MenteeAuth) error {
	if len(menteeAuth.Password) < 6 {
		return pkg.ErrPasswordLengthInvalid
	}

	user, _ := m.userRepository.FindByEmail(menteeAuth.Email)

	if user != nil {
		return pkg.ErrEmailAlreadyExist
	}

	newOTP := pkg.GenerateOTP(4)

	var err error

	ctx := context.Background()

	err = m.otpRepository.Save(ctx, menteeAuth.Email, newOTP, constants.TIME_TO_LIVE)

	if err != nil {
		return err
	}

	message := fmt.Sprintf("OTP: %s", newOTP)
	subject := "Verification Registering Eduworld"

	_ = m.mailerConfig.SendMail(menteeAuth.Email, subject, message)

	return nil
}

func (m menteeUsecase) VerifyRegister(menteeDomain *MenteeRegister) error {
	var err error

	ctx := context.Background()
	var validOTP string

	validOTP, err = m.otpRepository.Get(ctx, menteeDomain.Email)

	if err != nil {
		return err
	}

	if validOTP != menteeDomain.OTP {
		return pkg.ErrOTPNotMatch
	}

	userId := uuid.NewString()
	hashedPassword := utils.HashPassword(menteeDomain.Password)

	user := users.Domain{
		ID:        userId,
		Email:     menteeDomain.Email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = m.userRepository.Create(&user)

	if err != nil {
		return err
	}

	menteeId := uuid.NewString()

	mentee := Domain{
		ID:        menteeId,
		UserId:    userId,
		Fullname:  menteeDomain.Fullname,
		Phone:     menteeDomain.Phone,
		Role:      "mentee",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = m.menteeRepository.Create(&mentee)

	if err != nil {
		return err
	}

	return nil
}

func (m menteeUsecase) ForgotPassword(forgotPassword *MenteeForgotPassword) error {
	var err error

	if len(forgotPassword.Password) < 6 || len(forgotPassword.RepeatedPassword) < 6 {
		return pkg.ErrPasswordLengthInvalid
	}

	var user *users.Domain
	user, err = m.userRepository.FindByEmail(forgotPassword.Email)

	if err != nil {
		return err
	}

	ctx := context.Background()
	var result string

	result, err = m.otpRepository.Get(ctx, forgotPassword.Email)

	if err != nil {
		return err
	}

	if result != forgotPassword.OTP {
		return pkg.ErrOTPNotMatch
	}

	if forgotPassword.Password != forgotPassword.RepeatedPassword {
		return pkg.ErrPasswordNotMatch
	}

	hashPassword := utils.HashPassword(forgotPassword.RepeatedPassword)

	updatedUser := users.Domain{
		Password: hashPassword,
	}

	err = m.userRepository.Update(user.ID, &updatedUser)

	if err != nil {
		return err
	}

	return nil
}

func (m menteeUsecase) Login(menteeAuth *MenteeAuth) (interface{}, error) {
	if len(menteeAuth.Password) < 6 {
		return nil, pkg.ErrPasswordLengthInvalid
	}

	var err error

	var user *users.Domain
	user, err = m.userRepository.FindByEmail(menteeAuth.Email)

	if err != nil {
		return nil, err
	}

	ok := utils.ComparePassword(user.Password, menteeAuth.Password)
	if !ok {
		return nil, pkg.ErrUserNotFound
	}

	var mentee *Domain
	mentee, err = m.menteeRepository.FindByIdUser(user.ID)

	if err != nil {
		return nil, err
	}

	var token string
	exp := time.Now().Add(6 * time.Hour)

	token, err = m.jwtConfig.GenerateToken(user.ID, mentee.ID, mentee.Role, exp)

	if err != nil {
		return nil, err
	}

	data := map[string]interface{}{
		"token":   token,
		"expires": exp,
	}

	return data, nil
}

func (m menteeUsecase) FindAll() (*[]Domain, error) {
	//TODO implement me
	panic("implement me")
}

func (m menteeUsecase) FindById(id string) (*Domain, error) {
	//TODO implement me
	panic("implement me")
}

func (m menteeUsecase) Update(id string, userDomain *Domain) error {
	//TODO implement me
	panic("implement me")
}
