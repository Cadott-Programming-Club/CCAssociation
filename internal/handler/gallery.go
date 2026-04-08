package handler

import (
	"github.com/labstack/echo/v4"

	"ccassociation/templates/pages"
)

func (h *Handler) Gallery(c echo.Context) error {
	return pages.Gallery().Render(c.Request().Context(), c.Response().Writer)
}
