package users

import (
	"errors"
	"lms-api/internal/abstraction"
	"lms-api/internal/dto"
	"lms-api/internal/factory"
	"lms-api/internal/model"
	"lms-api/internal/repository"
	res "lms-api/pkg/util/response"
	"lms-api/pkg/util/trxmanager"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Service interface {
	Register(ctx *abstraction.Context, payload *dto.UserRegisterRequest) (*dto.UserRegisterResponse, error)
	Login(ctx *abstraction.Context, payload *dto.UserLoginRequest) (*dto.UserLoginResponse, error)
	FindByID(ctx *abstraction.Context, ID *int) (*dto.UserGetByIDResponse, error)
}

type service struct {
	Repository repository.User
	Db         *gorm.DB
}

func NewService(f *factory.Factory) *service {
	repository := f.UserRepository
	db := f.Db
	return &service{repository, db}
}

func (s *service) Register(ctx *abstraction.Context, payload *dto.UserRegisterRequest) (*dto.UserRegisterResponse, error) {
	var result *dto.UserRegisterResponse
	var data *model.UserEntityModel

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		data, err = s.Repository.Create(ctx, &payload.UserEntity)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.Duplicate, err)
		}
		return nil
	}); err != nil {
		return result, err
	}

	result = &dto.UserRegisterResponse{
		UserEntityModel: *data,
	}

	return result, nil
}

func (s *service) Login(ctx *abstraction.Context, payload *dto.UserLoginRequest) (*dto.UserLoginResponse, error) {
	var result *dto.UserLoginResponse

	data, err := s.Repository.FindByEmail(ctx, &payload.Email)
	if data == nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.Unauthorized, err)
	}

	if err = bcrypt.CompareHashAndPassword([]byte(data.PasswordHash), []byte(payload.Password)); err != nil {
		// return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
		return result, res.CustomErrorBuilder(402, "wrong password", "wrong password")
	}

	token, err := data.GenerateToken()
	if err != nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result = &dto.UserLoginResponse{
		Token:           token,
		UserEntityModel: *data,
	}

	return result, nil

}

func (s *service) FindByID(ctx *abstraction.Context, ID *int) (*dto.UserGetByIDResponse, error) {
	var result *dto.UserGetByIDResponse

	data, err := s.Repository.FindByID(ctx, ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return result, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result = &dto.UserGetByIDResponse{
		UserEntityModel: *data,
	}
	return result, err

}
