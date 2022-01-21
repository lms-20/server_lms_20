package chapter

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
	Create(ctx *abstraction.Context, payload *dto.ChapterCreateRequest) (*dto.ChapterCreateResponse, error)
	Find(ctx *abstraction.Context) (*dto.ChapterGetResponse, error)
	Update(ctx *abstraction.Context, payload *dto.ChapterUpdateRequest) (*dto.ChapterUpdateResponse, error)
	FindByID(ctx *abstraction.Context, id *int) (*dto.ChapterGetByIDResponse, error)
}

type service struct {
	Repository repository.Chapter
	Db         *gorm.DB
}

func NewService(f *factory.Factory) *service {
	repository := f.ChapterRepository
	db := f.Db
	return &service{repository, db}
}

func (s *service) FindByID(ctx *abstraction.Context, id *int) (*dto.ChapterGetByIDResponse, error) {
	var result *dto.ChapterGetByIDResponse

	data, err := s.Repository.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return result, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	result = &dto.ChapterGetByIDResponse{
		ChapterEntityModel: *data,
	}

	return result, nil
}

func (s *service) Find(ctx *abstraction.Context) (*dto.ChapterGetResponse, error) {
	var result *dto.ChapterGetResponse
	var datas *[]model.ChapterEntityModel

	datas, err := s.Repository.FindAll(ctx)
	if err != nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result = &dto.ChapterGetResponse{
		Datas: *datas,
	}

	return result, nil
}

func (s *service) Create(ctx *abstraction.Context, payload *dto.ChapterCreateRequest) (*dto.ChapterCreateResponse, error) {
	var result *dto.ChapterCreateResponse
	var data *model.ChapterEntityModel

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		data, err = s.Repository.Create(ctx, &payload.ChapterEntity)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		return nil
	}); err != nil {
		return result, err

	}

	result = &dto.ChapterCreateResponse{
		ChapterEntityModel: *data,
	}

	return result, nil
}

func (s *service) Update(ctx *abstraction.Context, payload *dto.ChapterUpdateRequest) (*dto.ChapterUpdateResponse, error) {
	var result *dto.ChapterUpdateResponse
	var data *model.ChapterEntityModel

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		_, err := s.Repository.FindByID(ctx, &payload.ID)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}

		data, err = s.Repository.Update(ctx, &payload.ID, &payload.ChapterEntity)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.Duplicate, err)
		}
		return nil
	}); err != nil {
		return result, err
	}

	result = &dto.ChapterUpdateResponse{
		ChapterEntityModel: *data,
	}
	return result, nil

}

func (s *service) Delete(ctx *abstraction.Context, id *int) (*dto.ChapterDeleteResponse, error) {
	var result *dto.ChapterDeleteResponse
	var data *model.ChapterEntityModel

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

	result = &dto.ChapterDeleteResponse{
		ChapterEntityModel: *data,
	}

	return result, nil
}
