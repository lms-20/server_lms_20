package model

import (
	"lms-api/internal/abstraction"
	"lms-api/pkg/util/date"

	"gorm.io/gorm"
)

type ReqCounselingEntity struct {
	CourseID    int    `json:"course_id" validate:"required"`
	UserID      int    `json:"user_id"`
	Section     string `json:"section" validate:"required"`
	Description string `json:"description" validate:"required"`
	Goal        string `json:"goal" validate:"required"`
}

type ReqCounselingEntityModel struct {
	// abstraction
	abstraction.Entity

	//entity
	ReqCounselingEntity

	// relational
	User   UserEntityModel   `json:"user" gorm:"foreignKey:UserID;"`
	Course CourseEntityModel `json:"course" gorm:"foreignKey:CourseID;"`

	// contexts
	Context *abstraction.Context `json:"-" gorm:"-"`
}

func (ReqCounselingEntityModel) TableName() string {
	return "request_counselings"
}

func (m *ReqCounselingEntityModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.UserID = m.Context.Auth.ID
	m.CreatedBy = m.Context.Auth.Name
	return
}

func (m *ReqCounselingEntityModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	m.UserID = m.Context.Auth.ID
	m.ModifiedBy = &m.Context.Auth.Name
	return
}
