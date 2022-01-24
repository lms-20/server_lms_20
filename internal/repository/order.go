package repository

import (
	"lms-api/internal/abstraction"
	"lms-api/internal/model"

	"gorm.io/gorm"
)

type Order interface {
	Create(ctx *abstraction.Context, course_id *int) (*model.OrderEntityModel, error)
	Update(ctx *abstraction.Context, order_id *int, e *model.OrderEntity) (*model.OrderEntityModel, error)
	FindByID(ctx *abstraction.Context, id *int) (*model.OrderEntityModel, error)
	FindByUserID(ctx *abstraction.Context, user_id *int) (*[]model.OrderEntityModel, error)
	CreateMyCourse(ctx *abstraction.Context, userID *int, courseID *int) error
}

type order struct {
	abstraction.Repository
}

func NewOrder(db *gorm.DB) *order {
	return &order{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *order) FindByUserID(ctx *abstraction.Context, user_id *int) (*[]model.OrderEntityModel, error) {
	conn := r.CheckTrx(ctx)
	var datas []model.OrderEntityModel

	err := conn.Where("user_id = ?", user_id).Find(&datas).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return &datas, nil
}

func (r *order) CreateMyCourse(ctx *abstraction.Context, userID *int, courseID *int) error {
	conn := r.CheckTrx(ctx)
	var data model.MyCourseEntityModel
	data.Context = ctx
	data.CourseID = *courseID
	data.UserID = *userID

	err := conn.Create(&data).Error
	if err != nil {
		return err
	}
	err = conn.Model(&data).First(&data).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *order) FindByID(ctx *abstraction.Context, id *int) (*model.OrderEntityModel, error) {
	conn := r.CheckTrx(ctx)
	var data model.OrderEntityModel

	err := conn.Where("id = ?", id).First(&data).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *order) Create(ctx *abstraction.Context, course_id *int) (*model.OrderEntityModel, error) {
	conn := r.CheckTrx(ctx)
	var data model.OrderEntityModel
	data.OrderEntity.CourseID = *course_id
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

func (r *order) Update(ctx *abstraction.Context, order_id *int, e *model.OrderEntity) (*model.OrderEntityModel, error) {
	conn := r.CheckTrx(ctx)
	var data model.OrderEntityModel
	data.Context = ctx

	err := conn.Where("id = ?", order_id).First(&data).Error
	if err != nil {
		return nil, err
	}

	data.OrderEntity = *e

	err = conn.Model(&data).Updates(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil

}
