package mentor

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
	Create(ctx *abstraction.Context, payload *dto.MentorCreateRequest) (*dto.MentorCreateResponse, error)
	Find(ctx *abstraction.Context) (*dto.MentorGetResponse, error)
	Update(ctx *abstraction.Context, payload *dto.MentorUpdateRequest) (*dto.MentorUpdateResponse, error)
	FindByID(ctx *abstraction.Context, id *int) (*dto.MentorGetByIDResponse, error)
}

type service struct {
	Repository repository.Mentor
	Db         *gorm.DB
}

func NewService(f *factory.Factory) *service {
	repository := f.MentorRepository
	db := f.Db
	return &service{repository, db}
}

func (s *service) FindByID(ctx *abstraction.Context, id *int) (*dto.MentorGetByIDResponse, error) {
	var result *dto.MentorGetByIDResponse

	data, err := s.Repository.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return result, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	result = &dto.MentorGetByIDResponse{
		MentorEntityModel: *data,
	}

	return result, nil
}

func (s *service) Find(ctx *abstraction.Context) (*dto.MentorGetResponse, error) {
	var result *dto.MentorGetResponse
	var datas *[]model.MentorEntityModel

	datas, err := s.Repository.FindAll(ctx)
	if err != nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result = &dto.MentorGetResponse{
		Datas: *datas,
	}

	return result, nil
}

func (s *service) Create(ctx *abstraction.Context, payload *dto.MentorCreateRequest) (*dto.MentorCreateResponse, error) {
	var result *dto.MentorCreateResponse
	var data *model.MentorEntityModel

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		data, err = s.Repository.Create(ctx, &payload.MentorEntity)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		return nil
	}); err != nil {
		return result, err

	}

	result = &dto.MentorCreateResponse{
		MentorEntityModel: *data,
	}

	return result, nil
}

func (s *service) Update(ctx *abstraction.Context, payload *dto.MentorUpdateRequest) (*dto.MentorUpdateResponse, error) {
	var result *dto.MentorUpdateResponse
	var data *model.MentorEntityModel

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		_, err := s.Repository.FindByID(ctx, &payload.ID)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}

		data, err = s.Repository.Update(ctx, &payload.ID, &payload.MentorEntity)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.Duplicate, err)
		}
		return nil
	}); err != nil {
		return result, err
	}

	result = &dto.MentorUpdateResponse{
		MentorEntityModel: *data,
	}
	return result, nil

}

func (s *service) Delete(ctx *abstraction.Context, id *int) (*dto.MentorDeleteResponse, error) {
	var result *dto.MentorDeleteResponse
	var data *model.MentorEntityModel

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

	result = &dto.MentorDeleteResponse{
		MentorEntityModel: *data,
	}

	return result, nil
}
