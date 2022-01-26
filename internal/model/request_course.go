package model

import (
	"lms-api/internal/abstraction"
	"lms-api/pkg/util/date"

	"gorm.io/gorm"
)

type ReqCourseEntity struct {
	UserID      int    `json:"user_id"`
	Description string `json:"description" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Goal        string `json:"goal" validate:"required"`
}

type ReqCourseEntityModel struct {
	// abstraction
	abstraction.Entity

	//entity
	ReqCourseEntity

	// relational
	User UserEntityModel `json:"user" gorm:"foreignKey:UserID;"`

	// contexts
	Context *abstraction.Context `json:"-" gorm:"-"`
}

func (ReqCourseEntityModel) TableName() string {
	return "request_courses"
}

func (m *ReqCourseEntityModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.UserID = m.Context.Auth.ID
	m.CreatedBy = m.Context.Auth.Name
	return
}

func (m *ReqCourseEntityModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	m.UserID = m.Context.Auth.ID
	m.ModifiedBy = &m.Context.Auth.Name
	return
}
