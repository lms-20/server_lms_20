package reqcourse

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
	Find(ctx *abstraction.Context) (*dto.ReqCourseGetResponse, error)
	Create(ctx *abstraction.Context, payload *dto.ReqCourseCreateRequest) (*dto.ReqCourseCreateResponse, error)
}

type service struct {
	Repository repository.ReqCourse
	Db         *gorm.DB
}

func NewService(f *factory.Factory) *service {
	repository := f.ReqCourseRepository
	db := f.Db
	return &service{repository, db}
}

func (s *service) Find(ctx *abstraction.Context) (*dto.ReqCourseGetResponse, error) {
	var result *dto.ReqCourseGetResponse
	var datas *[]model.ReqCourseEntityModel

	datas, err := s.Repository.FindAll(ctx)
	if err != nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result = &dto.ReqCourseGetResponse{
		Datas: *datas,
	}

	return result, nil
}

func (s *service) Create(ctx *abstraction.Context, payload *dto.ReqCourseCreateRequest) (*dto.ReqCourseCreateResponse, error) {
	var result *dto.ReqCourseCreateResponse
	var data *model.ReqCourseEntityModel

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		data, err = s.Repository.Create(ctx, &payload.ReqCourseEntity)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		return nil
	}); err != nil {
		return result, err

	}

	result = &dto.ReqCourseCreateResponse{
		ReqCourseEntityModel: *data,
	}

	return result, nil
}
