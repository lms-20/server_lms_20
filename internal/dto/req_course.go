package dto

import (
	"lms-api/internal/model"
	res "lms-api/pkg/util/response"
)

// Get
type ReqCourseGetResponse struct {
	Datas []model.ReqCourseEntityModel
}

type ReqCourseGetResponseDoc struct {
	Body struct {
		Meta res.Meta             `json:"meta"`
		Data ReqCourseGetResponse `json:"data"`
	} `json:"body"`
}

// GetByID
type ReqCourseGetByIDRequest struct {
	ID int `param:"id" validate:"required,numeric"`
}

type ReqCourseGetByIDResponse struct {
	model.ReqCourseEntityModel
}

type ReqCourseGetByIDResponseDoc struct {
	Body struct {
		Meta res.Meta                 `json:"meta"`
		Data ReqCourseGetByIDResponse `json:"data"`
	} `json:"body"`
}

// create
type ReqCourseCreateRequest struct {
	model.ReqCourseEntity
}

type ReqCourseCreateResponse struct {
	model.ReqCourseEntityModel
}

type ReqCourseCreateResponseDoc struct {
	Body struct {
		Meta res.Meta                `json:"meta"`
		Data ReqCourseCreateResponse `json:"data"`
	} `json:"body"`
}

// update
type ReqCourseUpdateRequest struct {
	ID int `param:"id"`
	model.ReqCourseEntity
}

type ReqCourseUpdateResponse struct {
	model.ReqCourseEntityModel
}

type ReqCourseUpdateResponseDoc struct {
	Body struct {
		Meta res.Meta                `json:"meta"`
		Data ReqCourseUpdateResponse `json:"data"`
	} `json:"body"`
}

// delete
type ReqCourseDeleteRequest struct {
	ID int `param:"id" validateL:"required,numeric"`
}

type ReqCourseDeleteResponse struct {
	model.ReqCourseEntityModel
}

type ReqCourseDeleteResponseDoc struct {
	Body struct {
		Meta res.Meta                `json:"meta"`
		Data ReqCourseDeleteResponse `json:"data"`
	} `json:"body"`
}
