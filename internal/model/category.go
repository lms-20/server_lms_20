package model

import (
	"lms-api/internal/abstraction"
	"lms-api/pkg/util/date"

	"gorm.io/gorm"
)

type CategoryEntity struct {
	Name string `json:"name" validate:"required" gorm:"index:idx_category_name,unique"`
}

type CategoryEntityModel struct {
	// abstraction
	abstraction.Entity

	// entity
	CategoryEntity

	// context
	Context *abstraction.Context `json:"-" gorm:"-"`
}

func (CategoryEntityModel) TableName() string {
	return "categories"
}

func (m *CategoryEntityModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = m.Context.Auth.Name

	return
}

func (m *CategoryEntityModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	m.ModifiedBy = &m.Context.Auth.Name
	return
}
