package repository

import (
	"fmt"
	"lms-api/internal/abstraction"
	"lms-api/internal/model"

	"gorm.io/gorm"
)

type Category interface {
	Create(ctx *abstraction.Context, data *model.CategoryEntity) (*model.CategoryEntityModel, error)
	FindAll(ctx *abstraction.Context) (*[]model.CategoryEntityModel, error)
}

type category struct {
	abstraction.Repository
}

func NewCategory(db *gorm.DB) *category {
	return &category{
		abstraction.Repository{
			Db: db,
		},
	}
}
func (r *category) FindAll(ctx *abstraction.Context) (*[]model.CategoryEntityModel, error) {
	conn := r.CheckTrx(ctx)
	var datas []model.CategoryEntityModel

	err := conn.Find(&datas).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return &datas, err
	}

	return &datas, nil

}

func (r *category) Create(ctx *abstraction.Context, e *model.CategoryEntity) (*model.CategoryEntityModel, error) {
	conn := r.CheckTrx(ctx)
	var data model.CategoryEntityModel
	data.CategoryEntity = *e
	data.Context = ctx
	fmt.Println(ctx.Auth.Name)
	fmt.Println("========================================")
	err := conn.Create(&data).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	err = conn.Model(&data).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}
