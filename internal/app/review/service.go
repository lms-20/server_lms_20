package review

import (
	"errors"
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
	Create(ctx *abstraction.Context, payload *dto.ReviewCreateRequest) (*dto.ReviewCreateResponse, error)
	Find(ctx *abstraction.Context) (*dto.ReviewGetResponse, error)
	Update(ctx *abstraction.Context, payload *dto.ReviewUpdateRequest) (*dto.ReviewUpdateResponse, error)
	FindByID(ctx *abstraction.Context, id *int) (*dto.ReviewGetByIDResponse, error)
}

type service struct {
	Repository repository.Review
	Db         *gorm.DB
}

func NewService(f *factory.Factory) *service {
	repository := f.ReviewRepository
	db := f.Db
	return &service{repository, db}
}

func (s *service) FindByID(ctx *abstraction.Context, id *int) (*dto.ReviewGetByIDResponse, error) {
	var result *dto.ReviewGetByIDResponse

	data, err := s.Repository.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return result, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	result = &dto.ReviewGetByIDResponse{
		ReviewEntityModel: *data,
	}

	return result, nil
}

func (s *service) Find(ctx *abstraction.Context) (*dto.ReviewGetResponse, error) {
	var result *dto.ReviewGetResponse
	var datas *[]model.ReviewEntityModel

	datas, err := s.Repository.FindAll(ctx)
	if err != nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result = &dto.ReviewGetResponse{
		Datas: *datas,
	}

	return result, nil
}

func (s *service) Create(ctx *abstraction.Context, payload *dto.ReviewCreateRequest) (*dto.ReviewCreateResponse, error) {
	var result *dto.ReviewCreateResponse
	var data *model.ReviewEntityModel

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		data, err = s.Repository.Create(ctx, &payload.ReviewEntity)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		return nil
	}); err != nil {
		return result, err

	}

	result = &dto.ReviewCreateResponse{
		ReviewEntityModel: *data,
	}

	return result, nil
}

func (s *service) Update(ctx *abstraction.Context, payload *dto.ReviewUpdateRequest) (*dto.ReviewUpdateResponse, error) {
	var result *dto.ReviewUpdateResponse
	var data *model.ReviewEntityModel

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		_, err := s.Repository.FindByID(ctx, &payload.ID)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}

		data, err = s.Repository.Update(ctx, &payload.ID, &payload.ReviewEntity)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.Duplicate, err)
		}
		return nil
	}); err != nil {
		return result, err
	}

	result = &dto.ReviewUpdateResponse{
		ReviewEntityModel: *data,
	}
	return result, nil

}

func (s *service) Delete(ctx *abstraction.Context, id *int) (*dto.ReviewDeleteResponse, error) {
	var result *dto.ReviewDeleteResponse
	var data *model.ReviewEntityModel

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		data, err = s.Repository.FindByID(ctx, id)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}

		data.Context = ctx
		data, err = s.Repository.Delete(ctx, id, data)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}
		return nil
	}); err != nil {
		return result, err
	}

	result = &dto.ReviewDeleteResponse{
		ReviewEntityModel: *data,
	}

	return result, nil
}
