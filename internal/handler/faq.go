package handler

import (
	"ccassociation/templates/pages"

	"github.com/labstack/echo/v4"
)

func (h *Handler) FAQ(c echo.Context) error {
	return pages.FAQ().Render(c.Request().Context(), c.Response().Writer)
}
