package dto

import (
	"lms-api/internal/model"
	res "lms-api/pkg/util/response"
)

// Get
type MyCourseGetResponse struct {
	Datas []model.MyCourseEntityModel
}

type MyCourseGetResponseDoc struct {
	Body struct {
		Meta res.Meta            `json:"meta"`
		Data MyCourseGetResponse `json:"data"`
	} `json:"body"`
}

// GetByID
type MyCourseGetByIDRequest struct {
	ID int `param:"id" validate:"required,numeric"`
}

type MyCourseGetByIDResponse struct {
	model.MyCourseEntityModel
}

type MyCourseGetByIDResponseDoc struct {
	Body struct {
		Meta res.Meta                `json:"meta"`
		Data MyCourseGetByIDResponse `json:"data"`
	} `json:"body"`
}

// create
type MyCourseCreateRequest struct {
	model.MyCourseEntity
}

type MyCourseCreateResponse struct {
	model.MyCourseEntityModel
}

type MyCourseCreateResponseDoc struct {
	Body struct {
		Meta res.Meta               `json:"meta"`
		Data MyCourseCreateResponse `json:"data"`
	} `json:"body"`
}

// update
type MyCourseUpdateRequest struct {
	ID int `param:"id"`
	model.MyCourseEntity
}

type MyCourseUpdateResponse struct {
	model.MyCourseEntityModel
}

type MyCourseUpdateResponseDoc struct {
	Body struct {
		Meta res.Meta               `json:"meta"`
		Data MyCourseUpdateResponse `json:"data"`
	} `json:"body"`
}

// delete
type MyCourseDeleteRequest struct {
	ID int `param:"id" validateL:"required,numeric"`
}

type MyCourseDeleteResponse struct {
	model.MyCourseEntityModel
}

type MyCourseDeleteResponseDoc struct {
	Body struct {
		Meta res.Meta               `json:"meta"`
		Data MyCourseDeleteResponse `json:"data"`
	} `json:"body"`
}
