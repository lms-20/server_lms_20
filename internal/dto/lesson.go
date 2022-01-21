package dto

import (
	"lms-api/internal/model"
	res "lms-api/pkg/util/response"
)

// Get
type LessonGetResponse struct {
	Datas []model.LessonEntityModel
}

type LessonGetResponseDoc struct {
	Body struct {
		Meta res.Meta          `json:"meta"`
		Data LessonGetResponse `json:"data"`
	} `json:"body"`
}

// GetByID
type LessonGetByIDRequest struct {
	ID int `param:"id" validate:"required,numeric"`
}

type LessonGetByIDResponse struct {
	model.LessonEntityModel
}

type LessonGetByIDResponseDoc struct {
	Body struct {
		Meta res.Meta              `json:"meta"`
		Data LessonGetByIDResponse `json:"data"`
	} `json:"body"`
}

// create
type LessonCreateRequest struct {
	model.LessonEntity
}

type LessonCreateResponse struct {
	model.LessonEntityModel
}

type LessonCreateResponseDoc struct {
	Body struct {
		Meta res.Meta             `json:"meta"`
		Data LessonCreateResponse `json:"data"`
	} `json:"body"`
}

type LessonUpdateRequest struct {
	ID int `param:"id"`
	model.LessonEntity
}

type LessonUpdateResponse struct {
	model.LessonEntityModel
}

type LessonUpdateResponseDoc struct {
	Body struct {
		Meta res.Meta             `json:"meta"`
		Data LessonUpdateResponse `json:"data"`
	} `json:"body"`
}

// delete
type LessonDeleteRequest struct {
	ID int `param:"id" validateL:"required,numeric"`
}

type LessonDeleteResponse struct {
	model.LessonEntityModel
}

type LessonDeleteResponseDoc struct {
	Body struct {
		Meta res.Meta             `json:"meta"`
		Data LessonDeleteResponse `json:"data"`
	} `json:"body"`
}
