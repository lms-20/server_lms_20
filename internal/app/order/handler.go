package order

import (
	"lms-api/internal/abstraction"
	"lms-api/internal/dto"
	"lms-api/internal/factory"
	res "lms-api/pkg/util/response"

	"github.com/labstack/echo/v4"
)

type handler struct {
	service *service
	// serviceMyCourse mycourse.Service
}

var err error

func NewHandler(f *factory.Factory) *handler {
	service := NewService(f)
	return &handler{service: service}
}

func (h *handler) Get(c echo.Context) error {

	cc := c.(*abstraction.Context)
	id := cc.Auth.ID
	if id == 0 {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	result, err := h.service.FindByUserID(cc, &id)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result.Datas).Send(c)
}

func (h *handler) Webhook(c echo.Context) error {
	cc := c.(*abstraction.Context)
	payload := new(dto.TransactionNotificationRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	order, errs := h.service.ProcessOrder(cc, payload)
	if errs != nil {
		return res.ErrorResponse(errs).Send(c)
	}

	err = h.service.PremiumAccess(cc, &order.CourseID, &order.UserID)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}
	return res.CustomSuccessBuilder(200, nil, "OK", nil).Send(c)
}
