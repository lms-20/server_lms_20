package category

import (
	"lms-api/internal/abstraction"
	"lms-api/internal/dto"
	"lms-api/internal/factory"
	"lms-api/internal/model"
	"lms-api/internal/repository"
	res "lms-api/pkg/util/response"
	"lms-api/pkg/util/trxmanager"

	"gorm.io/gorm"
)

type Service interface {
	Create(ctx *abstraction.Context, payload *dto.CategoryCreateRequest) (*dto.CategoryCreateResponse, error)
	Find(ctx *abstraction.Context) (*dto.CategoryGetResponse, error)
	// Login(ctx *abstraction.Context, payload *dto.UserLoginRequest) (*dto.UserLoginResponse, error)
	// FindByID(ctx *abstraction.Context, ID *int) (*dto.UserGetByIDResponse, error)
}

type service struct {
	Repository repository.Category
	Db         *gorm.DB
}

func NewService(f *factory.Factory) *service {
	repository := f.CategoryRepository
	db := f.Db
	return &service{repository, db}
}

func (s *service) Create(ctx *abstraction.Context, payload *dto.CategoryCreateRequest) (*dto.CategoryCreateResponse, error) {
	var result *dto.CategoryCreateResponse
	var data *model.CategoryEntityModel

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		data, err = s.Repository.Create(ctx, &payload.CategoryEntity)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.Duplicate, err)
		}
		return nil
	}); err != nil {
		return result, err
	}

	result = &dto.CategoryCreateResponse{
		CategoryEntityModel: *data,
	}

	return result, nil
}

func (s *service) Find(ctx *abstraction.Context) (*dto.CategoryGetResponse, error) {
	var result *dto.CategoryGetResponse
	var datas *[]model.CategoryEntityModel

	datas, err := s.Repository.FindAll(ctx)
	if err != nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result = &dto.CategoryGetResponse{
		Datas: *datas,
	}

	return result, nil

}
