package category

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

func (h *handler) Get(c echo.Context) error {
	cc := c.(*abstraction.Context)

	result, err := h.service.Find(cc)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result.Datas).Send(c)

}

func (h *handler) Create(c echo.Context) error {
	cc := c.(*abstraction.Context)

	payload := new(dto.CategoryCreateRequest)
	if err = c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	if err = c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}

	data, err := h.service.Create(cc, payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(data).Send(c)
}
