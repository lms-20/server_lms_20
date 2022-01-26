package repository

import (
	"lms-api/internal/abstraction"
	"lms-api/internal/model"

	"gorm.io/gorm"
)

type ReqCounseling interface {
	Create(ctx *abstraction.Context, e *model.ReqCounselingEntity) (*model.ReqCounselingEntityModel, error)
	FindAll(ctx *abstraction.Context) (*[]model.ReqCounselingEntityModel, error)
}

type reqCounseling struct {
	abstraction.Repository
}

func NewReqCounseling(db *gorm.DB) *reqCounseling {
	return &reqCounseling{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *reqCounseling) Create(ctx *abstraction.Context, e *model.ReqCounselingEntity) (*model.ReqCounselingEntityModel, error) {
	conn := r.CheckTrx(ctx)
	var data model.ReqCounselingEntityModel
	data.ReqCounselingEntity = *e
	data.Context = ctx

	err := conn.Create(&data).Error
	if err != nil {
		return nil, err
	}
	err = conn.Model(&data).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *reqCounseling) FindAll(ctx *abstraction.Context) (*[]model.ReqCounselingEntityModel, error) {
	conn := r.CheckTrx(ctx)
	var datas []model.ReqCounselingEntityModel

	err := conn.Preload("User").Find(&datas).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return &datas, err
	}

	return &datas, nil

}
