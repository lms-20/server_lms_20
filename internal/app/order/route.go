package order

import (
	"lms-api/internal/middleware"

	"github.com/labstack/echo/v4"
)

func (h *handler) Route(g *echo.Group) {
	g.GET("", h.Get, middleware.Authentication)
	g.POST("/webhook", h.Webhook)
}
