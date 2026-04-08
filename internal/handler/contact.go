package handler

import (
	"github.com/labstack/echo/v4"

	"ccassociation/templates/pages"
)

func (h *Handler) Contact(c echo.Context) error {
	return pages.Contact().Render(c.Request().Context(), c.Response().Writer)
}
