package dto

import "lms-api/internal/model"

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
