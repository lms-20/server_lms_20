package dto

import (
	"lms-api/internal/model"
	res "lms-api/pkg/util/response"
)

// GetByID
type UserGetByIDRequest struct {
	ID int
}
type UserGetByIDResponse struct {
	model.UserEntityModel
}
type UserGetByIDResponseDoc struct {
	Body struct {
		Meta res.Meta            `json:"meta"`
		Data UserGetByIDResponse `json:"data"`
	} `json:"body"`
}

//Register
type UserRegisterRequest struct {
	model.UserEntity
}

type UserRegisterResponse struct {
	model.UserEntityModel
}

type UserRegisterResponseDoc struct {
	Body struct {
		Meta res.Meta             `json:"meta"`
		Data UserRegisterResponse `json:"data"`
	} `json:"body"`
}

// login
type UserLoginRequest struct {
	Email    string `json:"emailAddress" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserLoginResponse struct {
	model.UserEntityModel
	Token string `json:"token"`
}

type UserLoginResponseDoc struct {
	Body struct {
		Meta res.Meta         `json:"meta"`
		Data UserLoginRequest `json:"data"`
	} `json:"body"`
}
