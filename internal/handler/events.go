package handler

import (
	"github.com/labstack/echo/v4"

	"ccassociation/templates/pages"
)

func (h *Handler) Events(c echo.Context) error {
	return pages.Events().Render(c.Request().Context(), c.Response().Writer)
}

func (h *Handler) EventNaborDays(c echo.Context) error {
	return pages.EventNaborDays().Render(c.Request().Context(), c.Response().Writer)
}

func (h *Handler) EventCommunityBingos(c echo.Context) error {
	return pages.EventCommunityBingos().Render(c.Request().Context(), c.Response().Writer)
}

func (h *Handler) EventWagonRides(c echo.Context) error {
	return pages.EventWagonRides().Render(c.Request().Context(), c.Response().Writer)
}
