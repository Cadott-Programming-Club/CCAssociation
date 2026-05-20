package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"ccassociation/internal/events"
	"ccassociation/templates/pages"
)

func (h *Handler) Health(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}

func (h *Handler) Home(c echo.Context) error {
	upcoming := events.Upcoming(h.events, time.Now(), 4)
	return pages.Home(upcoming).Render(c.Request().Context(), c.Response().Writer)
}
