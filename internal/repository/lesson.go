package repository

import (
	"lms-api/internal/abstraction"
	"lms-api/internal/model"

	"gorm.io/gorm"
)

type Lesson interface {
	Create(ctx *abstraction.Context, e *model.LessonEntity) (*model.LessonEntityModel, error)
	FindAll(ctx *abstraction.Context) (*[]model.LessonEntityModel, error)
	Update(ctx *abstraction.Context, id *int, e *model.LessonEntity) (*model.LessonEntityModel, error)
	FindByID(ctx *abstraction.Context, id *int) (*model.LessonEntityModel, error)
	Delete(ctx *abstraction.Context, id *int, e *model.LessonEntityModel) (*model.LessonEntityModel, error)
}

type lesson struct {
	abstraction.Repository
}

func NewLesson(db *gorm.DB) *lesson {
	return &lesson{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *lesson) FindByID(ctx *abstraction.Context, id *int) (*model.LessonEntityModel, error) {
	conn := r.CheckTrx(ctx)
	var data model.LessonEntityModel

	err := conn.Where("id = ?", id).First(&data).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *lesson) Update(ctx *abstraction.Context, id *int, e *model.LessonEntity) (*model.LessonEntityModel, error) {
	conn := r.CheckTrx(ctx)
	var data model.LessonEntityModel
	data.Context = ctx

	err := conn.Where("id = ?", id).First(&data).Error
	if err != nil {
		return nil, err
	}

	data.LessonEntity = *e

	err = conn.Model(&data).Updates(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil

}

func (r *lesson) Create(ctx *abstraction.Context, e *model.LessonEntity) (*model.LessonEntityModel, error) {
	conn := r.CheckTrx(ctx)
	var data model.LessonEntityModel
	data.LessonEntity = *e
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

func (r *lesson) FindAll(ctx *abstraction.Context) (*[]model.LessonEntityModel, error) {
	conn := r.CheckTrx(ctx)
	var datas []model.LessonEntityModel

	err := conn.Find(&datas).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return &datas, err
	}

	return &datas, nil

}

func (r *lesson) Delete(ctx *abstraction.Context, id *int, e *model.LessonEntityModel) (*model.LessonEntityModel, error) {
	conn := r.CheckTrx(ctx)
	err := conn.Where("id = ?", id).Delete(e).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}

	return e, nil
}
