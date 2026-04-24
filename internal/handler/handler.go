package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"ccassociation/internal/config"
	"ccassociation/internal/database"
	"ccassociation/templates/pages"
)

type Handler struct {
	cfg *config.Config
	db  *database.DB
}

func New(cfg *config.Config, db *database.DB) *Handler {
	return &Handler{
		cfg: cfg,
		db:  db,
	}
}

func (h *Handler) RegisterRoutes(e *echo.Echo) {
	// Static files with cache headers
	e.Static("/static", "static")

	// SEO files
	e.GET("/robots.txt", func(c echo.Context) error {
		return c.File("static/robots.txt")
	})
	e.GET("/sitemap.xml", func(c echo.Context) error {
		c.Response().Header().Set("Content-Type", "application/xml")
		return c.File("static/sitemap.xml")
	})

	// Health check
	e.GET("/health", h.Health)

	// Pages
	e.GET("/", h.Home)
	e.GET("/events", h.Events)
	e.GET("/events/nabor-days", h.EventNaborDays)
	e.GET("/events/community-bingos", h.EventCommunityBingos)
	e.GET("/events/sleigh-rides", h.EventSleighRides)
	e.GET("/gallery", h.Gallery)
	e.GET("/faq", h.FAQ)
	e.GET("/support", h.Support)
	e.GET("/grand-marshals", h.GrandMarshals)
	e.GET("/contact", h.Contact)

	// Custom 404
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		code := http.StatusInternalServerError
		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
		}
		if code == http.StatusNotFound {
			c.Response().WriteHeader(http.StatusNotFound)
			if err := pages.NotFound().Render(c.Request().Context(), c.Response().Writer); err != nil {
				e.DefaultHTTPErrorHandler(err, c)
			}
			return
		}
		e.DefaultHTTPErrorHandler(err, c)
	}
}
