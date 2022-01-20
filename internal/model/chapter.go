package model

import (
	"lms-api/internal/abstraction"
	"lms-api/pkg/util/date"

	"gorm.io/gorm"
)

type ChapterEntity struct {
	Name     string `json:"name" validate:"required"`
	CourseID int    `json:"course_id" validate:"required"`
	Order    int    `json:"order" validate:"required"`
	Link_ppt string `json:"link_ppt" validate:"required"`
}

type ChapterEntityModel struct {
	// abstraction
	abstraction.Entity

	//entity
	ChapterEntity

	// relational
	Lessons []LessonEntityModel `json:"chapters" gorm:"foreignKey:ChapterID;"`

	// context
	Context *abstraction.Context `json:"-" gorm:"-"`
}

func (ChapterEntityModel) TableName() string {
	return "chapters"
}

func (m *ChapterEntityModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = m.Context.Auth.Name
	return
}

func (m *ChapterEntityModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	m.ModifiedBy = &m.Context.Auth.Name
	return
}
