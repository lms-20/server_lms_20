package repository

import (
	"lms-api/internal/abstraction"
	"lms-api/internal/model"
	"regexp"

	"gorm.io/gorm"
)

type User interface {
	FindByEmail(ctx *abstraction.Context, email *string) (*model.UserEntityModel, error)
	Create(ctx *abstraction.Context, data *model.UserEntity) (*model.UserEntityModel, error)
	FindByID(ctx *abstraction.Context, id *int) (*model.UserEntityModel, error)
	checkTrx(ctx *abstraction.Context) *gorm.DB
}

type user struct {
	abstraction.Repository
}

func NewUser(db *gorm.DB) *user {
	return &user{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *user) FindByEmail(ctx *abstraction.Context, email *string) (*model.UserEntityModel, error) {
	conn := r.checkTrx(ctx)

	var data model.UserEntityModel
	err := conn.Where("email = ?", email).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *user) Create(ctx *abstraction.Context, e *model.UserEntity) (*model.UserEntityModel, error) {
	conn := r.checkTrx(ctx)

	var data model.UserEntityModel
	data.UserEntity = *e

	var regex, _ = regexp.Compile(`^[a-zA-Z0-9._%+-]+@alterra\.com$`)

	var isMatch = regex.MatchString(data.UserEntity.Email)
	if isMatch {
		data.UserEntity.Role = "employee"
	} else {
		data.UserEntity.Role = "student"
	}

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

func (r *user) checkTrx(ctx *abstraction.Context) *gorm.DB {
	if ctx.Trx != nil {
		return ctx.Trx.Db
	}
	return r.Db
}

func (r *user) FindByID(ctx *abstraction.Context, id *int) (*model.UserEntityModel, error) {
	conn := r.checkTrx(ctx)

	var data model.UserEntityModel
	err := conn.Where("id = ?", id).First(&data).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}
