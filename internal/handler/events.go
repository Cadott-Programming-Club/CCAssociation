package handler

import (
	"ccassociation/templates/pages"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Events(c echo.Context) error {
	return pages.Events().Render(c.Request().Context(), c.Response().Writer)
}
