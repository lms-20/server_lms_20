package users

import (
	"lms-api/internal/abstraction"
	"lms-api/internal/dto"
	"lms-api/internal/factory"
	res "lms-api/pkg/util/response"

	"github.com/labstack/echo/v4"
)

type handler struct {
	service *service
}

var err error

func NewHandler(f *factory.Factory) *handler {
	service := NewService(f)
	return &handler{service}
}

func (h *handler) Register(c echo.Context) error {
	cc := c.(*abstraction.Context)

	payload := new(dto.UserRegisterRequest)
	if err = c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	if err = c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}

	data, err := h.service.Register(cc, payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(data).Send(c)
}

func (h *handler) Login(c echo.Context) error {
	cc := c.(*abstraction.Context)

	payload := new(dto.UserLoginRequest)
	if err = c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err = c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}
	data, err := h.service.Login(cc, payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}
	return res.SuccessResponse(data).Send(c)

}

func (h *handler) GetByID(c echo.Context) error {
	cc := c.(*abstraction.Context)
	id := cc.Auth.ID
	if id == 0 {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	result, err := h.service.FindByID(cc, &id)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	// fmt.Println(cc.Auth)

	return res.SuccessResponse(result).Send(c)
}
