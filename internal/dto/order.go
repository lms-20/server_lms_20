package dto

import "lms-api/internal/model"

// GetByID
type OrderGetByUserIDRequest struct {
	ID int
}
type OrderGetByUserIDResponse struct {
	Datas []model.OrderEntityModel
}

// create
type OrderCreateRequest struct {
	CourseID int `json:"course_id" validate:"true"`
}

type OrderCreateResponse struct {
	model.OrderEntityModel
}

// update
type OrderUpdateRequest struct {
	ID int `param:"id"`
	model.OrderEntity
}

type OrderUpdateResponse struct {
	model.OrderEntityModel
}

// transaction notification
type TransactionNotificationRequest struct {
	TransactionStatus string `json:"transaction_status"`
	OrderID           string `json:"order_id"`
	PaymentType       string `json:"payment_type"`
	FraudStatus       string `json:"fraud_status"`
}

type TransactionNotificationResponse struct {
	UserID   int `json:"user_id"`
	CourseID int `json:"course_id"`
}
