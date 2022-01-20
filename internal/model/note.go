package model

import (
	"lms-api/internal/abstraction"
	"lms-api/pkg/util/date"

	"gorm.io/gorm"
)

type NoteEntity struct {
	Note     string `json:"note" validate:"required"`
	LessonID int    `json:"lesson_id" validate:"required"`
}

type NoteEntityModel struct {
	// abstraction
	abstraction.Entity

	//entity
	NoteEntity

	// relationals
	UserID int               `json:"user_id"`
	User   UserEntityModel   `json:"user" gorm:"foreignKey:UserID"`
	Lesson LessonEntityModel `json:"lesson" gorm:"foreignKey:LessonID"`

	// contexts
	Context *abstraction.Context `json:"-" gorm:"-"`
}

func (NoteEntityModel) TableName() string {
	return "notes"
}

func (m *NoteEntityModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.UserID = m.Context.Auth.ID
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = m.Context.Auth.Name
	return
}

func (m *NoteEntityModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	m.ModifiedBy = &m.Context.Auth.Name
	return
}
