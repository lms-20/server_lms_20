package repository

import (
	"lms-api/internal/abstraction"

	"lms-api/internal/model"

	"gorm.io/gorm"
)

type Mentor interface {
	Create(ctx *abstraction.Context, e *model.MentorEntity) (*model.MentorEntityModel, error)
	FindAll(ctx *abstraction.Context) (*[]model.MentorEntityModel, error)
	Update(ctx *abstraction.Context, id *int, e *model.MentorEntity) (*model.MentorEntityModel, error)
	FindByID(ctx *abstraction.Context, id *int) (*model.MentorEntityModel, error)
	Delete(ctx *abstraction.Context, id *int, e *model.MentorEntityModel) (*model.MentorEntityModel, error)
}

type mentor struct {
	abstraction.Repository
}

func NewMentor(db *gorm.DB) *mentor {
	return &mentor{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *mentor) FindByID(ctx *abstraction.Context, id *int) (*model.MentorEntityModel, error) {
	conn := r.CheckTrx(ctx)
	var data model.MentorEntityModel

	err := conn.Where("id = ?", id).First(&data).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *mentor) Update(ctx *abstraction.Context, id *int, e *model.MentorEntity) (*model.MentorEntityModel, error) {
	conn := r.CheckTrx(ctx)
	var data model.MentorEntityModel
	data.Context = ctx

	err := conn.Where("id = ?", id).First(&data).Error
	if err != nil {
		return nil, err
	}

	data.MentorEntity = *e

	err = conn.Model(&data).Updates(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil

}

func (r *mentor) Create(ctx *abstraction.Context, e *model.MentorEntity) (*model.MentorEntityModel, error) {
	conn := r.CheckTrx(ctx)
	var data model.MentorEntityModel
	data.MentorEntity = *e
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

func (r *mentor) FindAll(ctx *abstraction.Context) (*[]model.MentorEntityModel, error) {
	conn := r.CheckTrx(ctx)
	var datas []model.MentorEntityModel

	err := conn.Find(&datas).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return &datas, err
	}

	return &datas, nil

}

func (r *mentor) Delete(ctx *abstraction.Context, id *int, e *model.MentorEntityModel) (*model.MentorEntityModel, error) {
	conn := r.CheckTrx(ctx)
	err := conn.Where("id = ?", id).Delete(e).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}

	return e, nil
}
