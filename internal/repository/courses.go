package repository

import (
	"lms-api/internal/abstraction"
	"lms-api/internal/model"

	"gorm.io/gorm"
)

type Course interface {
	Create(ctx *abstraction.Context, e *model.CourseEntity) (*model.CourseEntityModel, error)
	FindAll(ctx *abstraction.Context) (*[]model.CourseEntityModel, error)
	Update(ctx *abstraction.Context, id *int, e *model.CourseEntity) (*model.CourseEntityModel, error)
	FindByID(ctx *abstraction.Context, id *int) (*model.CourseEntityModel, error)
	Delete(ctx *abstraction.Context, id *int, e *model.CourseEntityModel) (*model.CourseEntityModel, error)
}

type course struct {
	abstraction.Repository
}

func NewCourse(db *gorm.DB) *course {
	return &course{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *course) FindByID(ctx *abstraction.Context, id *int) (*model.CourseEntityModel, error) {
	conn := r.CheckTrx(ctx)
	var data model.CourseEntityModel

	err := conn.Where("id = ?", id).Preload("Chapters.Lessons").Preload("Category").Preload("Mentor").First(&data).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *course) Update(ctx *abstraction.Context, id *int, e *model.CourseEntity) (*model.CourseEntityModel, error) {
	conn := r.CheckTrx(ctx)
	var data model.CourseEntityModel
	data.Context = ctx

	err := conn.Where("id = ?", id).First(&data).Error
	if err != nil {
		return nil, err
	}

	data.CourseEntity = *e

	err = conn.Model(&data).Updates(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil

}

func (r *course) Create(ctx *abstraction.Context, e *model.CourseEntity) (*model.CourseEntityModel, error) {
	conn := r.CheckTrx(ctx)
	var data model.CourseEntityModel
	data.CourseEntity = *e
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

func (r *course) FindAll(ctx *abstraction.Context) (*[]model.CourseEntityModel, error) {
	conn := r.CheckTrx(ctx)
	var datas []model.CourseEntityModel

	err := conn.Preload("Mentor").Preload("Category").Find(&datas).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return &datas, err
	}

	return &datas, nil

}

func (r *course) Delete(ctx *abstraction.Context, id *int, e *model.CourseEntityModel) (*model.CourseEntityModel, error) {
	conn := r.CheckTrx(ctx)
	err := conn.Where("id = ?", id).Delete(e).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}

	return e, nil
}
