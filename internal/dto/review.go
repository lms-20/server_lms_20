package dto

import (
	"lms-api/internal/model"
	res "lms-api/pkg/util/response"
)

// Get
type ReviewGetResponse struct {
	Datas []model.ReviewEntityModel
}

type ReviewGetResponseDoc struct {
	Body struct {
		Meta res.Meta          `json:"meta"`
		Data ReviewGetResponse `json:"data"`
	} `json:"body"`
}

// GetByID
type ReviewGetByIDRequest struct {
	ID int `param:"id" validate:"required,numeric"`
}

type ReviewGetByIDResponse struct {
	model.ReviewEntityModel
}

type ReviewGetByIDResponseDoc struct {
	Body struct {
		Meta res.Meta              `json:"meta"`
		Data ReviewGetByIDResponse `json:"data"`
	} `json:"body"`
}

// create
type ReviewCreateRequest struct {
	model.ReviewEntity
}

type ReviewCreateResponse struct {
	model.ReviewEntityModel
}

type ReviewCreateResponseDoc struct {
	Body struct {
		Meta res.Meta             `json:"meta"`
		Data ReviewCreateResponse `json:"data"`
	} `json:"body"`
}

// update
type ReviewUpdateRequest struct {
	ID int `param:"id"`
	model.ReviewEntity
}

type ReviewUpdateResponse struct {
	model.ReviewEntityModel
}

type ReviewUpdateResponseDoc struct {
	Body struct {
		Meta res.Meta             `json:"meta"`
		Data ReviewUpdateResponse `json:"data"`
	} `json:"body"`
}

// delete
type ReviewDeleteRequest struct {
	ID int `param:"id" validateL:"required,numeric"`
}

type ReviewDeleteResponse struct {
	model.ReviewEntityModel
}

type ReviewDeleteResponseDoc struct {
	Body struct {
		Meta res.Meta             `json:"meta"`
		Data ReviewDeleteResponse `json:"data"`
	} `json:"body"`
}
