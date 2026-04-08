package handler

import (
	"github.com/labstack/echo/v4"

	"ccassociation/templates/pages"
)

func (h *Handler) FAQ(c echo.Context) error {
	return pages.FAQ().Render(c.Request().Context(), c.Response().Writer)
}
