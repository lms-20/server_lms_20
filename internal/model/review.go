package model

import (
	"lms-api/internal/abstraction"
	"lms-api/pkg/util/date"

	"gorm.io/gorm"
)

// id int [pk, increment]
// user_id int [ref: > users.id]
// course_id int [ref: > courses.id]
// rating int
// note longtext

type ReviewEntity struct {
	// relational
	UserID   int `json:"user_id"`
	CourseID int `json:"course_id"`
	// field
	Rating int    `json:"rating" validate:"required"`
	Note   string `json:"note" validate:"required"`
}

type ReviewEntityModel struct {
	// abstraction
	abstraction.Entity

	//entity
	ReviewEntity

	// relational
	User UserEntityModel `json:"user" gorm:"foreignKey:UserID;"`

	// contexts
	Context *abstraction.Context `json:"-" gorm:"-"`
}

func (ReviewEntityModel) TableName() string {
	return "reviews"
}

func (m *ReviewEntityModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.UserID = m.Context.Auth.ID
	m.CreatedBy = m.Context.Auth.Name
	return
}

func (m *ReviewEntityModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	m.ModifiedBy = &m.Context.Auth.Name
	return
}
