package lessons

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
	Create(ctx *abstraction.Context, payload *dto.LessonCreateRequest) (*dto.LessonCreateResponse, error)
	Find(ctx *abstraction.Context) (*dto.LessonGetResponse, error)
	Update(ctx *abstraction.Context, payload *dto.LessonUpdateRequest) (*dto.LessonUpdateResponse, error)
	FindByID(ctx *abstraction.Context, id *int) (*dto.LessonGetByIDResponse, error)
}

type service struct {
	Repository repository.Lesson
	Db         *gorm.DB
}

func NewService(f *factory.Factory) *service {
	repository := f.LessonRepository
	db := f.Db
	return &service{repository, db}
}

func (s *service) FindByID(ctx *abstraction.Context, id *int) (*dto.LessonGetByIDResponse, error) {
	var result *dto.LessonGetByIDResponse

	data, err := s.Repository.FindByID(ctx, id)
	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return result, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	result = &dto.LessonGetByIDResponse{
		LessonEntityModel: *data,
	}

	return result, nil
}

func (s *service) Find(ctx *abstraction.Context) (*dto.LessonGetResponse, error) {
	var result *dto.LessonGetResponse
	var datas *[]model.LessonEntityModel

	datas, err := s.Repository.FindAll(ctx)
	if err != nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result = &dto.LessonGetResponse{
		Datas: *datas,
	}

	return result, nil
}

func (s *service) Create(ctx *abstraction.Context, payload *dto.LessonCreateRequest) (*dto.LessonCreateResponse, error) {
	var result *dto.LessonCreateResponse
	var data *model.LessonEntityModel

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		data, err = s.Repository.Create(ctx, &payload.LessonEntity)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		return nil
	}); err != nil {
		return result, err

	}

	result = &dto.LessonCreateResponse{
		LessonEntityModel: *data,
	}

	return result, nil
}

func (s *service) Update(ctx *abstraction.Context, payload *dto.LessonUpdateRequest) (*dto.LessonUpdateResponse, error) {
	var result *dto.LessonUpdateResponse
	var data *model.LessonEntityModel

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		_, err := s.Repository.FindByID(ctx, &payload.ID)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}

		data, err = s.Repository.Update(ctx, &payload.ID, &payload.LessonEntity)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.Duplicate, err)
		}
		return nil
	}); err != nil {
		return result, err
	}

	result = &dto.LessonUpdateResponse{
		LessonEntityModel: *data,
	}
	return result, nil

}

func (s *service) Delete(ctx *abstraction.Context, id *int) (*dto.LessonDeleteResponse, error) {
	var result *dto.LessonDeleteResponse
	var data *model.LessonEntityModel

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

	result = &dto.LessonDeleteResponse{
		LessonEntityModel: *data,
	}

	return result, nil
}
