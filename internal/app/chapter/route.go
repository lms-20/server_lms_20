package chapter

import (
	"lms-api/internal/middleware"

	"github.com/labstack/echo/v4"
)

func (h *handler) Route(g *echo.Group) {
	g.GET("", h.Get)
	g.POST("", h.Create, middleware.Authentication)
	g.PUT("/:id", h.Update, middleware.Authentication)
	g.GET("/:id", h.GetByID)
	g.DELETE("/:id", h.Delete, middleware.Authentication)
}
