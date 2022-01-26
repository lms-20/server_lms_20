package mycourse

import (
	"lms-api/internal/middleware"

	"github.com/labstack/echo/v4"
)

func (h *handler) Route(g *echo.Group) {
	g.GET("", h.GetByID, middleware.Authentication)
	g.POST("", h.Create, middleware.Authentication)
}
