package repository

import (
	"lms-api/internal/abstraction"
	"lms-api/internal/model"

	"gorm.io/gorm"
)

type Note interface {
	Create(ctx *abstraction.Context, e *model.NoteEntity) (*model.NoteEntityModel, error)
	FindAll(ctx *abstraction.Context) (*[]model.NoteEntityModel, error)
	Update(ctx *abstraction.Context, id *int, e *model.NoteEntity) (*model.NoteEntityModel, error)
	FindByID(ctx *abstraction.Context, id *int) (*model.NoteEntityModel, error)
	Delete(ctx *abstraction.Context, id *int, e *model.NoteEntityModel) (*model.NoteEntityModel, error)
}

type note struct {
	abstraction.Repository
}

func NewNote(db *gorm.DB) *note {
	return &note{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *note) FindByID(ctx *abstraction.Context, id *int) (*model.NoteEntityModel, error) {
	conn := r.CheckTrx(ctx)
	var data model.NoteEntityModel

	err := conn.Where("id = ?", id).First(&data).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *note) Update(ctx *abstraction.Context, id *int, e *model.NoteEntity) (*model.NoteEntityModel, error) {
	conn := r.CheckTrx(ctx)
	var data model.NoteEntityModel
	data.Context = ctx

	err := conn.Where("id = ?", id).First(&data).Error
	if err != nil {
		return nil, err
	}

	data.NoteEntity = *e

	err = conn.Model(&data).Updates(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil

}

func (r *note) Create(ctx *abstraction.Context, e *model.NoteEntity) (*model.NoteEntityModel, error) {
	conn := r.CheckTrx(ctx)
	var data model.NoteEntityModel
	data.NoteEntity = *e
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

func (r *note) FindAll(ctx *abstraction.Context) (*[]model.NoteEntityModel, error) {
	conn := r.CheckTrx(ctx)
	var datas []model.NoteEntityModel

	err := conn.Find(&datas).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return &datas, err
	}

	return &datas, nil

}

func (r *note) Delete(ctx *abstraction.Context, id *int, e *model.NoteEntityModel) (*model.NoteEntityModel, error) {
	conn := r.CheckTrx(ctx)
	err := conn.Where("id = ?", id).Delete(e).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}

	return e, nil
}
