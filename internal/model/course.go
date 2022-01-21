package model

import (
	"lms-api/internal/abstraction"
)

type levelType string

const (
	all_level    levelType = "all-level"
	beginner     levelType = "beginner"
	intermediete levelType = "intermediete"
	advance      levelType = "advance"
)

type typeType string

const (
	free    typeType = "free"
	premium typeType = "premium"
)

type statusType string

const (
	draft     statusType = "draft"
	published statusType = "published"
)

type CourseEntity struct {
	Name        string     `json:"name" validate:"required"`
	Kode        string     `json:"kode" validate:"required"`
	Certificate bool       `json:"certificate" validate:"required"`
	Thumbnail   string     `json:"thumbnail" validate:"required"`
	Type        typeType   `json:"type" sql:"type:typeType" validate:"required"`
	Status      statusType `json:"status" sql:"type:statusType" validate:"required"`
	Price       int        `json:"price"  validate:"required"`
	Level       levelType  `json:"level" sql:"type:levelType"  validate:"required"`
	Description string     `json:"description"  validate:"required"`
	MentorID    int        `json:"mentor_id"  validate:"required"`
	CategoryID  int        `json:"category_id"  validate:"required"`
}

type CourseEntityModel struct {
	// abstraction
	abstraction.Entity

	// entity
	CourseEntity

	// relationals
	Mentor   MentorEntityModel    `json:"mentor" gorm:"foreignKey:MentorID"`
	Category CategoryEntityModel  `json:"category" gorm:"foreignKey:CategoryID"`
	Chapters []ChapterEntityModel `json:"chapters" gorm:"foreignKey:CourseID;"`
	Reviews  []ReviewEntityModel  `json:"reviews" gorm:"foreignKey:CourseID;"`

	// context
	Context *abstraction.Context `json:"-" gorm:"-"`
}

func (CourseEntity) TableName() string {
	return "courses"
}

//   id int [pk, increment]
//   category_id int [ref: > categories.id]
//   name varchar
//   kode string
//   certificate tinyint
//   thumbnail varchar
//   type course_type
//   status course_status
//   price int
//   level course_level
//   description longtext
//   mentor_id int [ref: > mentors.id]
//   created_at datetime
//   updated_at datetime
