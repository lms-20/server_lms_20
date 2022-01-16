package dto

import (
	"lms-api/internal/model"
	res "lms-api/pkg/util/response"
)

// Get
type MentorGetResponse struct {
	Datas []model.MentorEntityModel
}

type MentorGetResponseDoc struct {
	Body struct {
		Meta res.Meta          `json:"meta"`
		Data MentorGetResponse `json:"data"`
	} `json:"body"`
}

// GetByID
type MentorGetByIDRequest struct {
	ID int `param:"id" validate:"required,numeric"`
}

type MentorGetByIDResponse struct {
	model.MentorEntityModel
}

type MentorGetByIDResponseDoc struct {
	Body struct {
		Meta res.Meta              `json:"meta"`
		Data MentorGetByIDResponse `json:"data"`
	} `json:"body"`
}

// create
type MentorCreateRequest struct {
	model.MentorEntity
}

type MentorCreateResponse struct {
	model.MentorEntityModel
}

type MentorCreateResponseDoc struct {
	Body struct {
		Meta res.Meta             `json:"meta"`
		Data MentorCreateResponse `json:"data"`
	} `json:"body"`
}

// update
type paramID struct {
	ID int `json:"id" param:"id"`
}
type MentorUpdateRequest struct {
	ID int `param:"id"`
	model.MentorEntity
}

type MentorUpdateResponse struct {
	model.MentorEntityModel
}

type MentorUpdateResponseDoc struct {
	Body struct {
		Meta res.Meta             `json:"meta"`
		Data MentorUpdateResponse `json:"data"`
	} `json:"body"`
}

// delete
type MentorDeleteRequest struct {
	ID int `param:"id" validateL:"required,numeric"`
}

type MentorDeleteResponse struct {
	model.MentorEntityModel
}

type MentorDeleteResponseDoc struct {
	Body struct {
		Meta res.Meta             `json:"meta"`
		Data MentorDeleteResponse `json:"data"`
	} `json:"body"`
}
