package repository

import (
	"lms-api/internal/abstraction"
	"lms-api/internal/model"

	"gorm.io/gorm"
)

type Review interface {
	Create(ctx *abstraction.Context, e *model.ReviewEntity) (*model.ReviewEntityModel, error)
	FindAll(ctx *abstraction.Context) (*[]model.ReviewEntityModel, error)
	FindByID(ctx *abstraction.Context, id *int) (*model.ReviewEntityModel, error)
	Update(ctx *abstraction.Context, id *int, e *model.ReviewEntity) (*model.ReviewEntityModel, error)
	Delete(ctx *abstraction.Context, id *int, e *model.ReviewEntityModel) (*model.ReviewEntityModel, error)
}

type review struct {
	abstraction.Repository
}

func NewReview(db *gorm.DB) *review {
	return &review{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *review) FindByID(ctx *abstraction.Context, id *int) (*model.ReviewEntityModel, error) {
	conn := r.CheckTrx(ctx)
	var data model.ReviewEntityModel

	err := conn.Where("id = ?", id).First(&data).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *review) Update(ctx *abstraction.Context, id *int, e *model.ReviewEntity) (*model.ReviewEntityModel, error) {
	conn := r.CheckTrx(ctx)
	var data model.ReviewEntityModel
	data.Context = ctx

	err := conn.Where("id = ?", id).First(&data).Error
	if err != nil {
		return nil, err
	}

	data.ReviewEntity = *e

	err = conn.Model(&data).Updates(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil

}

func (r *review) Create(ctx *abstraction.Context, e *model.ReviewEntity) (*model.ReviewEntityModel, error) {
	conn := r.CheckTrx(ctx)
	var data model.ReviewEntityModel
	data.ReviewEntity = *e
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

func (r *review) FindAll(ctx *abstraction.Context) (*[]model.ReviewEntityModel, error) {
	conn := r.CheckTrx(ctx)
	var datas []model.ReviewEntityModel

	err := conn.Find(&datas).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return &datas, err
	}

	return &datas, nil

}

func (r *review) Delete(ctx *abstraction.Context, id *int, e *model.ReviewEntityModel) (*model.ReviewEntityModel, error) {
	conn := r.CheckTrx(ctx)
	err := conn.Where("id = ?", id).Delete(e).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}

	return e, nil
}
