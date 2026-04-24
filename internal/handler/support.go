package handler

import (
	"github.com/labstack/echo/v4"

	"ccassociation/templates/pages"
)

func (h *Handler) Support(c echo.Context) error {
	return pages.Support().Render(c.Request().Context(), c.Response().Writer)
}

func (h *Handler) GrandMarshals(c echo.Context) error {
	return pages.GrandMarshals().Render(c.Request().Context(), c.Response().Writer)
}
