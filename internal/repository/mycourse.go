package repository

import (
	"lms-api/internal/abstraction"
	"lms-api/internal/model"

	"gorm.io/gorm"
)

type MyCourse interface {
	Create(ctx *abstraction.Context, e *model.MyCourseEntity) (*model.MyCourseEntityModel, error)
	FindByID(ctx *abstraction.Context, id *int) (*[]model.MyCourseEntityModel, error)
}

type mycourse struct {
	abstraction.Repository
}

func NewMyCourse(db *gorm.DB) *mycourse {
	return &mycourse{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *mycourse) FindByID(ctx *abstraction.Context, id *int) (*[]model.MyCourseEntityModel, error) {
	conn := r.CheckTrx(ctx)
	var datas []model.MyCourseEntityModel

	err := conn.Where("user_id = ?", id).Preload("Course").Find(&datas).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return &datas, nil
}

func (r *mycourse) Create(ctx *abstraction.Context, e *model.MyCourseEntity) (*model.MyCourseEntityModel, error) {
	conn := r.CheckTrx(ctx)
	var data model.MyCourseEntityModel
	data.MyCourseEntity = *e
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
