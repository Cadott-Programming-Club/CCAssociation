package handler

import (
	"github.com/labstack/echo/v4"

	"ccassociation/templates/pages"
)

func (h *Handler) Events(c echo.Context) error {
	return pages.Events().Render(c.Request().Context(), c.Response().Writer)
}
