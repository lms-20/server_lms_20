package users

import (
	"lms-api/internal/middleware"

	"github.com/labstack/echo/v4"
)

func (h *handler) Route(g *echo.Group) {
	g.POST("/register", h.Register)
	g.POST("/login", h.Login)
	g.GET("/fetch", h.GetByID, middleware.Authentication)
}
