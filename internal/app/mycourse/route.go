package mycourse

import (
	"github.com/labstack/echo/v4"
)

func (h *handler) Route(g *echo.Group) {
	g.GET("", h.GetByID)
	g.POST("", h.Create)
}
