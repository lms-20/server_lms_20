package repository

import (
	"lms-api/internal/abstraction"
	"lms-api/internal/model"

	"gorm.io/gorm"
)

type ReqCourse interface {
	Create(ctx *abstraction.Context, e *model.ReqCourseEntity) (*model.ReqCourseEntityModel, error)
	FindAll(ctx *abstraction.Context) (*[]model.ReqCourseEntityModel, error)
}

type reqCourse struct {
	abstraction.Repository
}

func NewReqCourse(db *gorm.DB) *reqCourse {
	return &reqCourse{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *reqCourse) Create(ctx *abstraction.Context, e *model.ReqCourseEntity) (*model.ReqCourseEntityModel, error) {
	conn := r.CheckTrx(ctx)
	var data model.ReqCourseEntityModel
	data.ReqCourseEntity = *e
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

func (r *reqCourse) FindAll(ctx *abstraction.Context) (*[]model.ReqCourseEntityModel, error) {
	conn := r.CheckTrx(ctx)
	var datas []model.ReqCourseEntityModel

	err := conn.Preload("User").Find(&datas).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return &datas, err
	}
	return &datas, nil
}
