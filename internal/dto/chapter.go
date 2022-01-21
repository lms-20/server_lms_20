package dto

import (
	"lms-api/internal/model"
	res "lms-api/pkg/util/response"
)

// Get
type ChapterGetResponse struct {
	Datas []model.ChapterEntityModel
}

type ChapterGetResponseDoc struct {
	Body struct {
		Meta res.Meta           `json:"meta"`
		Data ChapterGetResponse `json:"data"`
	} `json:"body"`
}

// GetByID
type ChapterGetByIDRequest struct {
	ID int `param:"id" validate:"required,numeric"`
}

type ChapterGetByIDResponse struct {
	model.ChapterEntityModel
}

type ChapterGetByIDResponseDoc struct {
	Body struct {
		Meta res.Meta               `json:"meta"`
		Data ChapterGetByIDResponse `json:"data"`
	} `json:"body"`
}

// create
type ChapterCreateRequest struct {
	model.ChapterEntity
}

type ChapterCreateResponse struct {
	model.ChapterEntityModel
}

type ChapterCreateResponseDoc struct {
	Body struct {
		Meta res.Meta              `json:"meta"`
		Data ChapterCreateResponse `json:"data"`
	} `json:"body"`
}

// update

type ChapterUpdateRequest struct {
	ID int `param:"id"`
	model.ChapterEntity
}

type ChapterUpdateResponse struct {
	model.ChapterEntityModel
}

type ChapterUpdateResponseDoc struct {
	Body struct {
		Meta res.Meta              `json:"meta"`
		Data ChapterUpdateResponse `json:"data"`
	} `json:"body"`
}

// delete
type ChapterDeleteRequest struct {
	ID int `param:"id" validateL:"required,numeric"`
}

type ChapterDeleteResponse struct {
	model.ChapterEntityModel
}

type ChapterDeleteResponseDoc struct {
	Body struct {
		Meta res.Meta              `json:"meta"`
		Data ChapterDeleteResponse `json:"data"`
	} `json:"body"`
}
