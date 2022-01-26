package mycourse

import (
	"lms-api/internal/abstraction"
	"lms-api/internal/app/courses"
	"lms-api/internal/app/order"

	"lms-api/internal/dto"
	"lms-api/internal/factory"
	res "lms-api/pkg/util/response"

	"github.com/labstack/echo/v4"
)

var err error

type handler struct {
	service       *service
	serviceCourse courses.Service
	serviceOrder  order.Service
}

func NewHandler(f *factory.Factory) *handler {
	service := NewService(f)
	serviceCourse := courses.NewService(f)
	serviceOrder := order.NewService(f)
	return &handler{service: service, serviceCourse: serviceCourse, serviceOrder: serviceOrder}
}

func (h *handler) Create(c echo.Context) error {
	cc := c.(*abstraction.Context)

	payload := new(dto.MyCourseCreateRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}

	course, err := h.serviceCourse.FindByID(cc, &payload.CourseID)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	if course.Type == "premium" && cc.Auth.Role == "student" {
		order, err := h.serviceOrder.Create(cc, &dto.OrderCreateRequest{CourseID: course.ID})
		if err != nil {
			return res.ErrorResponse(err).Send(c)
		}

		newOrder, err := h.serviceOrder.Update(cc, &dto.OrderUpdateRequest{ID: order.ID, OrderEntity: order.OrderEntity}, &dto.CourseGetByIDResponse{CourseEntityModel: course.CourseEntityModel})
		if err != nil {
			return res.ErrorResponse(err).Send(c)
		}

		return res.SuccessResponse(newOrder).Send(c)

	} else {
		result, err := h.service.Create(cc, payload)
		if err != nil {
			return res.ErrorResponse(err).Send(c)
		}

		return res.SuccessResponse(result).Send(c)
	}

}

func (h *handler) GetByID(c echo.Context) error {
	cc := c.(*abstraction.Context)
	id := cc.Auth.ID
	result, err := h.service.FindByID(cc, &id)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}
	return res.SuccessResponse(result.Datas).Send(c)
}
