package note

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
	Create(ctx *abstraction.Context, payload *dto.NoteCreateRequest) (*dto.NoteCreateResponse, error)
	Find(ctx *abstraction.Context) (*dto.NoteGetResponse, error)
	Update(ctx *abstraction.Context, payload *dto.NoteUpdateRequest) (*dto.NoteUpdateResponse, error)
	FindByID(ctx *abstraction.Context, id *int) (*dto.NoteGetByIDResponse, error)
}

type service struct {
	Repository repository.Note
	Db         *gorm.DB
}

func NewService(f *factory.Factory) *service {
	repository := f.NoteRepository
	db := f.Db
	return &service{repository, db}
}

func (s *service) FindByID(ctx *abstraction.Context, id *int) (*dto.NoteGetByIDResponse, error) {
	var result *dto.NoteGetByIDResponse

	data, err := s.Repository.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return result, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}

		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	result = &dto.NoteGetByIDResponse{
		NoteEntityModel: *data,
	}

	return result, nil
}

func (s *service) Find(ctx *abstraction.Context) (*dto.NoteGetResponse, error) {
	var result *dto.NoteGetResponse
	var datas *[]model.NoteEntityModel

	datas, err := s.Repository.FindAll(ctx)
	if err != nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result = &dto.NoteGetResponse{
		Datas: *datas,
	}

	return result, nil
}

func (s *service) Create(ctx *abstraction.Context, payload *dto.NoteCreateRequest) (*dto.NoteCreateResponse, error) {
	var result *dto.NoteCreateResponse
	var data *model.NoteEntityModel

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		data, err = s.Repository.Create(ctx, &payload.NoteEntity)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		return nil
	}); err != nil {
		return result, err

	}

	result = &dto.NoteCreateResponse{
		NoteEntityModel: *data,
	}

	return result, nil
}

func (s *service) Update(ctx *abstraction.Context, payload *dto.NoteUpdateRequest) (*dto.NoteUpdateResponse, error) {
	var result *dto.NoteUpdateResponse
	var data *model.NoteEntityModel

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		_, err := s.Repository.FindByID(ctx, &payload.ID)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}

		data, err = s.Repository.Update(ctx, &payload.ID, &payload.NoteEntity)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.Duplicate, err)
		}
		return nil
	}); err != nil {
		return result, err
	}

	result = &dto.NoteUpdateResponse{
		NoteEntityModel: *data,
	}
	return result, nil

}

func (s *service) Delete(ctx *abstraction.Context, id *int) (*dto.NoteDeleteResponse, error) {
	var result *dto.NoteDeleteResponse
	var data *model.NoteEntityModel

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

	result = &dto.NoteDeleteResponse{
		NoteEntityModel: *data,
	}

	return result, nil
}
