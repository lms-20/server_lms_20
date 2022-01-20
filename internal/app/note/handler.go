package note

import (
	"fmt"
	"lms-api/internal/abstraction"
	"lms-api/internal/dto"
	"lms-api/internal/factory"
	res "lms-api/pkg/util/response"
	"strconv"

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

	return res.CustomSuccessBuilder(200, result.Datas, "Get Notes Success", nil).Send(c)
}

func (h *handler) Create(c echo.Context) error {
	cc := c.(*abstraction.Context)

	payload := new(dto.NoteCreateRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}
	result, err := h.service.Create(cc, payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}

func (h *handler) Update(c echo.Context) error {
	cc := c.(*abstraction.Context)
	id := c.Param("id")
	payload := new(dto.NoteUpdateRequest)
	if err := c.Bind(&payload.NoteEntity); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}
	payload.ID, _ = strconv.Atoi(id)

	result, err := h.service.Update(cc, payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}

func (h *handler) GetByID(c echo.Context) error {
	cc := c.(*abstraction.Context)
	id, _ := strconv.Atoi(c.Param("id"))
	result, err := h.service.FindByID(cc, &id)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}
	return res.SuccessResponse(result).Send(c)
}

func (h *handler) Delete(c echo.Context) error {
	cc := c.(*abstraction.Context)
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := h.service.Delete(cc, &id)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}
	message := fmt.Sprintf("success delete note id : %d", id)
	return res.CustomSuccessBuilder(200, nil, message, nil).Send(c)
}
