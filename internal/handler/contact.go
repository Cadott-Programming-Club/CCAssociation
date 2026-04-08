package handler

import (
	"ccassociation/templates/pages"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Contact(c echo.Context) error {
	return pages.Contact().Render(c.Request().Context(), c.Response().Writer)
}
