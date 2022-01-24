package model

import (
	"lms-api/internal/abstraction"
	"lms-api/pkg/constant"
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
	User   UserEntityModel   `json:"user,omitempty" gorm:"foreignKey:UserID;"`

	// contexts
	Context *abstraction.Context `json:"-" gorm:"-"`
}

func (MyCourseEntityModel) TableName() string {
	return "mycourses"
}

func (m *MyCourseEntityModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = constant.DB_DEFAULT_CREATED_BY
	return
}

func (m *MyCourseEntityModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	m.ModifiedBy = &m.Context.Auth.Name
	return
}
