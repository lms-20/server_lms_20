package model

import (
	"lms-api/internal/abstraction"
	"lms-api/pkg/util/date"

	"gorm.io/gorm"
)

type LessonEntity struct {
	Name      string `json:"name" validate:"required"`
	Video     string `json:"video" validate:"required"`
	ChapterID int    `json:"chapter_id" validate:"required"`
}

type LessonEntityModel struct {
	// abstraction
	abstraction.Entity

	//entity
	LessonEntity

	// contexts
	Context *abstraction.Context `json:"-" gorm:"-"`
}

func (LessonEntityModel) TableName() string {
	return "lessons"
}

func (m *LessonEntityModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = m.Context.Auth.Name
	return
}

func (m *LessonEntityModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	m.ModifiedBy = &m.Context.Auth.Name
	return
}
