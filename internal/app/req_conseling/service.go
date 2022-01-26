package reqconseling

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
	Find(ctx *abstraction.Context) (*dto.ReqCounselingGetResponse, error)
	Create(ctx *abstraction.Context, payload *dto.ReqCounselingCreateRequest) (*dto.ReqCounselingCreateResponse, error)
}

type service struct {
	Repository repository.ReqCounseling
	Db         *gorm.DB
}

func NewService(f *factory.Factory) *service {
	repository := f.ReqCounselingRepository
	db := f.Db
	return &service{repository, db}
}

func (s *service) Find(ctx *abstraction.Context) (*dto.ReqCounselingGetResponse, error) {
	var result *dto.ReqCounselingGetResponse
	var datas *[]model.ReqCounselingEntityModel

	datas, err := s.Repository.FindAll(ctx)
	if err != nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result = &dto.ReqCounselingGetResponse{
		Datas: *datas,
	}

	return result, nil
}

func (s *service) Create(ctx *abstraction.Context, payload *dto.ReqCounselingCreateRequest) (*dto.ReqCounselingCreateResponse, error) {
	var result *dto.ReqCounselingCreateResponse
	var data *model.ReqCounselingEntityModel

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		data, err = s.Repository.Create(ctx, &payload.ReqCounselingEntity)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		return nil
	}); err != nil {
		return result, err

	}

	result = &dto.ReqCounselingCreateResponse{
		ReqCounselingEntityModel: *data,
	}

	return result, nil
}
