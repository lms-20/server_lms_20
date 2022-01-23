package model

import (
	"lms-api/internal/abstraction"
	"lms-api/pkg/util/date"

	"gorm.io/gorm"
)

type MyCourseEntity struct {
	CourseID int `json:"course_id"`
	UserID   int `json:"user_id"`
}

type MyCourseEntityModel struct {
	// abstraction
	abstraction.Entity

	//entity
	MyCourseEntity

	// relational
	Course CourseEntityModel `json:"course" gorm:"foreignKey:CourseID"`
	User   UserEntityModel   `json:"user" gorm:"foreignKey:UserID;"`

	// contexts
	Context *abstraction.Context `json:"-" gorm:"-"`
}

func (MyCourseEntityModel) TableName() string {
	return "mycourses"
}

func (m *MyCourseEntityModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.UserID = m.Context.Auth.ID
	m.CreatedBy = m.Context.Auth.Name
	return
}

func (m *MyCourseEntityModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	m.UserID = m.Context.Auth.ID
	m.ModifiedBy = &m.Context.Auth.Name
	return
}
