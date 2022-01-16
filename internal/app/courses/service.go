package courses

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
	Create(ctx *abstraction.Context, payload *dto.CourseCreateRequest) (*dto.CourseCreateResponse, error)
	Find(ctx *abstraction.Context) (*dto.CourseGetResponse, error)
	Update(ctx *abstraction.Context, payload *dto.CourseUpdateRequest) (*dto.CourseUpdateResponse, error)
	FindByID(ctx *abstraction.Context, id *int) (*dto.CourseGetByIDResponse, error)
}

type service struct {
	Repository repository.Course
	Db         *gorm.DB
}

func NewService(f *factory.Factory) *service {
	repository := f.CourseRepository
	db := f.Db
	return &service{repository, db}
}

func (s *service) FindByID(ctx *abstraction.Context, id *int) (*dto.CourseGetByIDResponse, error) {
	var result *dto.CourseGetByIDResponse

	data, err := s.Repository.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return result, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	result = &dto.CourseGetByIDResponse{
		CourseEntityModel: *data,
	}

	return result, nil
}

func (s *service) Find(ctx *abstraction.Context) (*dto.CourseGetResponse, error) {
	var result *dto.CourseGetResponse
	var datas *[]model.CourseEntityModel

	datas, err := s.Repository.FindAll(ctx)
	if err != nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result = &dto.CourseGetResponse{
		Datas: *datas,
	}

	return result, nil
}

func (s *service) Create(ctx *abstraction.Context, payload *dto.CourseCreateRequest) (*dto.CourseCreateResponse, error) {
	var result *dto.CourseCreateResponse
	var data *model.CourseEntityModel

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		data, err = s.Repository.Create(ctx, &payload.CourseEntity)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		return nil
	}); err != nil {
		return result, err

	}

	result = &dto.CourseCreateResponse{
		CourseEntityModel: *data,
	}

	return result, nil
}

func (s *service) Update(ctx *abstraction.Context, payload *dto.CourseUpdateRequest) (*dto.CourseUpdateResponse, error) {
	var result *dto.CourseUpdateResponse
	var data *model.CourseEntityModel

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		_, err := s.Repository.FindByID(ctx, &payload.ID)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}

		data, err = s.Repository.Update(ctx, &payload.ID, &payload.CourseEntity)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.Duplicate, err)
		}
		return nil
	}); err != nil {
		return result, err
	}

	result = &dto.CourseUpdateResponse{
		CourseEntityModel: *data,
	}
	return result, nil

}

func (s *service) Delete(ctx *abstraction.Context, id *int) (*dto.CourseDeleteResponse, error) {
	var result *dto.CourseDeleteResponse
	var data *model.CourseEntityModel

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

	result = &dto.CourseDeleteResponse{
		CourseEntityModel: *data,
	}

	return result, nil
}
