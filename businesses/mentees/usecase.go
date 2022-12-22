package mentees

import (
	"context"
	"fmt"
	"math"
	"path/filepath"
	"time"

	"github.com/Kelompok14-LMS/backend-go/businesses/otp"
	"github.com/Kelompok14-LMS/backend-go/businesses/users"
	"github.com/Kelompok14-LMS/backend-go/constants"
	"github.com/Kelompok14-LMS/backend-go/helper"
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
	storage          *helper.StorageConfig
}

func NewMenteeUsecase(
	menteeRepository Repository,
	userRepository users.Repository,
	otpRepository otp.Repository,
	jwtConfig *utils.JWTConfig,
	mailerConfig *pkg.MailerConfig,
	storage *helper.StorageConfig,
) Usecase {
	return menteeUsecase{
		menteeRepository: menteeRepository,
		userRepository:   userRepository,
		otpRepository:    otpRepository,
		jwtConfig:        jwtConfig,
		mailerConfig:     mailerConfig,
		storage:          storage,
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
	if len(menteeDomain.Password) < 6 {
		return pkg.ErrPasswordLengthInvalid
	}

	userDomain, _ := m.userRepository.FindByEmail(menteeDomain.Email)

	if userDomain != nil {
		return pkg.ErrEmailAlreadyExist
	}

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
	var err error

	mentees, err := m.menteeRepository.FindAll()

	if err != nil {
		if err == pkg.ErrMenteeNotFound {
			return nil, pkg.ErrMenteeNotFound
		}

		return nil, pkg.ErrInternalServerError
	}

	return mentees, nil
}

func (m menteeUsecase) FindById(id string) (*Domain, error) {
	mentee, err := m.menteeRepository.FindById(id)
	if err != nil {
		if err == pkg.ErrMenteeNotFound {
			return nil, pkg.ErrMenteeNotFound
		}

		return nil, pkg.ErrInternalServerError
	}

	return mentee, nil
}

func (m menteeUsecase) FindByCourse(courseId string, pagination pkg.Pagination) (*pkg.Pagination, error) {
	mentees, totalRows, err := m.menteeRepository.FindByCourse(courseId, pagination.GetLimit(), pagination.GetOffset())

	if err != nil {
		return nil, err
	}

	pagination.Result = mentees
	pagination.TotalRows = totalRows
	pagination.TotalPages = int(math.Ceil(float64(totalRows) / float64(pagination.GetLimit())))

	return &pagination, nil
}

func (m menteeUsecase) Update(ctx context.Context, id string, menteeDomain *Domain) error {
	mentee, err := m.menteeRepository.FindById(id)

	if err != nil {
		return err
	}

	var ProfilePictureURL string

	if menteeDomain.ProfilePictureFile != nil {
		if mentee.ProfilePicture != "" {
			if err := m.storage.DeleteObject(ctx, mentee.ProfilePicture); err != nil {
				return err
			}
		}

		ProfilePicture, err := menteeDomain.ProfilePictureFile.Open()
		if err != nil {
			return pkg.ErrUnsupportedImageFile
		}

		defer ProfilePicture.Close()

		extension := filepath.Ext(menteeDomain.ProfilePictureFile.Filename)

		if extension != ".jpg" && extension != ".png" && extension != ".jpeg" {
			return pkg.ErrUnsupportedImageFile
		}

		filename, _ := utils.GetFilename(menteeDomain.ProfilePictureFile.Filename)

		ProfilePictureURL, err = m.storage.UploadImage(ctx, filename, ProfilePicture)

		if err != nil {
			return err
		}
	}

	updatedMentee := Domain{
		Fullname:       menteeDomain.Fullname,
		Phone:          menteeDomain.Phone,
		BirthDate:      menteeDomain.BirthDate,
		Address:        menteeDomain.Address,
		ProfilePicture: ProfilePictureURL,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	err = m.menteeRepository.Update(id, &updatedMentee)
	if err != nil {
		if err == pkg.ErrMenteeNotFound {
			return pkg.ErrMenteeNotFound
		}

		return pkg.ErrInternalServerError
	}

	return nil
}
