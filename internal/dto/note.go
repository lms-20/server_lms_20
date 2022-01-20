package dto

import (
	"lms-api/internal/model"
	res "lms-api/pkg/util/response"
)

// Get
type NoteGetResponse struct {
	Datas []model.NoteEntityModel
}

type NoteGetResponseDoc struct {
	Body struct {
		Meta res.Meta        `json:"meta"`
		Data NoteGetResponse `json:"data"`
	} `json:"body"`
}

// GetByID
type NoteGetByIDRequest struct {
	ID int `param:"id" validate:"required,numeric"`
}

type NoteGetByIDResponse struct {
	model.NoteEntityModel
}

type NoteGetByIDResponseDoc struct {
	Body struct {
		Meta res.Meta            `json:"meta"`
		Data NoteGetByIDResponse `json:"data"`
	} `json:"body"`
}

// create
type NoteCreateRequest struct {
	model.NoteEntity
}

type NoteCreateResponse struct {
	model.NoteEntityModel
}

type NoteCreateResponseDoc struct {
	Body struct {
		Meta res.Meta           `json:"meta"`
		Data NoteCreateResponse `json:"data"`
	} `json:"body"`
}

type NoteUpdateRequest struct {
	ID int `param:"id"`
	model.NoteEntity
}

type NoteUpdateResponse struct {
	model.NoteEntityModel
}

type NoteUpdateResponseDoc struct {
	Body struct {
		Meta res.Meta           `json:"meta"`
		Data NoteUpdateResponse `json:"data"`
	} `json:"body"`
}

// delete
type NoteDeleteRequest struct {
	ID int `param:"id" validateL:"required,numeric"`
}

type NoteDeleteResponse struct {
	model.NoteEntityModel
}

type NoteDeleteResponseDoc struct {
	Body struct {
		Meta res.Meta           `json:"meta"`
		Data NoteDeleteResponse `json:"data"`
	} `json:"body"`
}
