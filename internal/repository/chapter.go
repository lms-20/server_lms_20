package repository

import (
	"lms-api/internal/abstraction"
	"lms-api/internal/model"

	"gorm.io/gorm"
)

type Chapter interface {
	Create(ctx *abstraction.Context, e *model.ChapterEntity) (*model.ChapterEntityModel, error)
	FindAll(ctx *abstraction.Context) (*[]model.ChapterEntityModel, error)
	Update(ctx *abstraction.Context, id *int, e *model.ChapterEntity) (*model.ChapterEntityModel, error)
	FindByID(ctx *abstraction.Context, id *int) (*model.ChapterEntityModel, error)
	Delete(ctx *abstraction.Context, id *int, e *model.ChapterEntityModel) (*model.ChapterEntityModel, error)
}

type chapter struct {
	abstraction.Repository
}

func NewChapter(db *gorm.DB) *chapter {
	return &chapter{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *chapter) FindByID(ctx *abstraction.Context, id *int) (*model.ChapterEntityModel, error) {
	conn := r.CheckTrx(ctx)
	var data model.ChapterEntityModel

	err := conn.Where("id = ?", id).First(&data).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *chapter) Update(ctx *abstraction.Context, id *int, e *model.ChapterEntity) (*model.ChapterEntityModel, error) {
	conn := r.CheckTrx(ctx)
	var data model.ChapterEntityModel
	data.Context = ctx

	err := conn.Where("id = ?", id).First(&data).Error
	if err != nil {
		return nil, err
	}

	data.ChapterEntity = *e

	err = conn.Model(&data).Updates(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil

}

func (r *chapter) Create(ctx *abstraction.Context, e *model.ChapterEntity) (*model.ChapterEntityModel, error) {
	conn := r.CheckTrx(ctx)
	var data model.ChapterEntityModel
	data.ChapterEntity = *e
	data.Context = ctx
	// fmt.Println(ctx.Auth.Name)
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

func (r *chapter) FindAll(ctx *abstraction.Context) (*[]model.ChapterEntityModel, error) {
	conn := r.CheckTrx(ctx)
	var datas []model.ChapterEntityModel

	err := conn.Find(&datas).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return &datas, err
	}

	return &datas, nil

}

func (r *chapter) Delete(ctx *abstraction.Context, id *int, e *model.ChapterEntityModel) (*model.ChapterEntityModel, error) {
	conn := r.CheckTrx(ctx)
	err := conn.Where("id = ?", id).Delete(e).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}

	return e, nil
}
