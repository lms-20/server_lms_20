package auth

import (
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
	Login(ctx *abstraction.Context, payload *dto.AuthLoginRequest) (*dto.AuthLoginResponse, error)
	Register(ctx *abstraction.Context, payload *dto.AuthRegisterRequest) (*dto.AuthRegisterResponse, error)
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

func (s *service) Login(ctx *abstraction.Context, payload *dto.AuthLoginRequest) (*dto.AuthLoginResponse, error) {
	var result *dto.AuthLoginResponse

	data, err := s.Repository.FindByEmail(ctx, &payload.Email)
	if data == nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.Unauthorized, err)
	}

	if err = bcrypt.CompareHashAndPassword([]byte(data.PasswordHash), []byte(payload.Password)); err != nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	token, err := data.GenerateToken()

	if err != nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result = &dto.AuthLoginResponse{
		Token:           token,
		UserEntityModel: *data,
	}

	return result, nil
}

func (s *service) Register(ctx *abstraction.Context, payload *dto.AuthRegisterRequest) (*dto.AuthRegisterResponse, error) {
	var result *dto.AuthRegisterResponse
	var data *model.UserEntityModel

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		data, err = s.Repository.Create(ctx, &payload.UserEntity)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		return nil
	}); err != nil {
		return result, err
	}

	result = &dto.AuthRegisterResponse{
		UserEntityModel: *data,
	}

	return result, nil
}
