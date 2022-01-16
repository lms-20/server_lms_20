package category

import (
	"lms-api/internal/middleware"

	"github.com/labstack/echo/v4"
)

func (h *handler) Route(g *echo.Group) {
	g.POST("", h.Create, middleware.Authentication)
	g.GET("", h.Get)
}
