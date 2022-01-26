package dto

import (
	"lms-api/internal/model"
	res "lms-api/pkg/util/response"
)

// Get
type ReqCounselingGetResponse struct {
	Datas []model.ReqCounselingEntityModel
}

type ReqCounselingGetResponseDoc struct {
	Body struct {
		Meta res.Meta                 `json:"meta"`
		Data ReqCounselingGetResponse `json:"data"`
	} `json:"body"`
}

// GetByID
type ReqCounselingGetByIDRequest struct {
	ID int `param:"id" validate:"required,numeric"`
}

type ReqCounselingGetByIDResponse struct {
	model.ReqCounselingEntityModel
}

type ReqCounselingGetByIDResponseDoc struct {
	Body struct {
		Meta res.Meta                     `json:"meta"`
		Data ReqCounselingGetByIDResponse `json:"data"`
	} `json:"body"`
}

// create
type ReqCounselingCreateRequest struct {
	model.ReqCounselingEntity
}

type ReqCounselingCreateResponse struct {
	model.ReqCounselingEntityModel
}

type ReqCounselingCreateResponseDoc struct {
	Body struct {
		Meta res.Meta                    `json:"meta"`
		Data ReqCounselingCreateResponse `json:"data"`
	} `json:"body"`
}

// update
type ReqCounselingUpdateRequest struct {
	ID int `param:"id"`
	model.ReqCounselingEntity
}

type ReqCounselingUpdateResponse struct {
	model.ReqCounselingEntityModel
}

type ReqCounselingUpdateResponseDoc struct {
	Body struct {
		Meta res.Meta                    `json:"meta"`
		Data ReqCounselingUpdateResponse `json:"data"`
	} `json:"body"`
}

// delete
type ReqCounselingDeleteRequest struct {
	ID int `param:"id" validateL:"required,numeric"`
}

type ReqCounselingDeleteResponse struct {
	model.ReqCounselingEntityModel
}

type ReqCounselingDeleteResponseDoc struct {
	Body struct {
		Meta res.Meta                    `json:"meta"`
		Data ReqCounselingDeleteResponse `json:"data"`
	} `json:"body"`
}
