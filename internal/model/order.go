package model

import (
	"lms-api/internal/abstraction"
	"lms-api/pkg/util/date"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type MetadataEntity struct {
	CourseID        int    `json:"course_id"`
	CoursePrice     int    `json:"course_price"`
	CourseName      string `json:"course_name"`
	CourseThumbnail string `json:"course_thumbnail"`
	CourseLevel     string `json:"course_level"`
}

type OrderEntity struct {
	Status   string         `json:"status"`
	CourseID int            `json:"course_id" validate:"required"`
	UserID   int            `json:"user_id" validate:"required"`
	SnapURL  string         `json:"snap_url"`
	Metadata datatypes.JSON `json:"metadata"`
}

type PaymentLogEntity struct {
	Status      string         `json:"status"`
	PaymentType string         `json:"payment_type"`
	RawResponse datatypes.JSON `json:"raw_response"`
	OrderID     int            `json:"order_id"`
}

type OrderEntityModel struct {
	// abstraction
	abstraction.Entity

	//entity
	OrderEntity

	// relationals
	PaymentLog PaymentLogEntityModel `json:"payment_log" gorm:"foreignKey:OrderID"`

	// contexts
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type PaymentLogEntityModel struct {
	// abstraction
	abstraction.Entity

	//entity
	PaymentLogEntity

	// contexts
	Context *abstraction.Context `json:"-" gorm:"-"`
}

func (OrderEntityModel) TableName() string {
	return "orders"
}

func (m *OrderEntityModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = m.Context.Auth.Name
	m.UserID = m.Context.Auth.ID
	return
}

func (m *OrderEntityModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	m.ModifiedBy = &m.Context.Auth.Name
	m.UserID = m.Context.Auth.ID
	return
}

func (PaymentLogEntityModel) TableName() string {
	return "payment_logs"
}

func (m *PaymentLogEntityModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = m.Context.Auth.Name
	return
}

func (m *PaymentLogEntityModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	m.ModifiedBy = &m.Context.Auth.Name
	return
}
