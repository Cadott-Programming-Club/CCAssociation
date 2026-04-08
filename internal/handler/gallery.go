package handler

import (
	"ccassociation/templates/pages"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Gallery(c echo.Context) error {
	return pages.Gallery().Render(c.Request().Context(), c.Response().Writer)
}
