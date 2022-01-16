package model

import (
	"lms-api/internal/abstraction"
	"lms-api/pkg/util/date"

	"gorm.io/gorm"
)

type MentorEntity struct {
	Name       string `json:"name" validate:"required"`
	Profile    string `json:"profile" validate:"required"`
	Profession string `json:"profession" validate:"required"`
	Email      string `json:"email" validate:"required,email" gorm:"index:idx_mentor_email,unique"`
}

type MentorEntityModel struct {
	// abstraction
	abstraction.Entity

	//entity
	MentorEntity

	// context
	Context *abstraction.Context `json:"-" gorm:"-"`
}

func (MentorEntityModel) TableName() string {
	return "mentors"
}

func (m *MentorEntityModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = m.Context.Auth.Name
	return
}

func (m *MentorEntityModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	m.ModifiedBy = &m.Context.Auth.Name
	return
}
