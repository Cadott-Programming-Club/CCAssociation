package handler

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"

	"ccassociation/internal/config"
	"ccassociation/internal/database"
	"ccassociation/internal/events"
	"ccassociation/templates/pages"
)

type Handler struct {
	cfg    *config.Config
	db     *database.DB
	events []events.Event
}

func New(cfg *config.Config, db *database.DB) *Handler {
	loaded, err := events.Load(events.EmbeddedJSON())
	if err != nil {
		slog.Error("events: failed to load embedded calendar", "error", err)
	}
	return &Handler{
		cfg:    cfg,
		db:     db,
		events: loaded,
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
	e.GET("/events/calendar", h.EventsCalendar)
	e.GET("/events/calendar/day/:date", h.EventsDay)
	e.GET("/events/c/:slug", h.EventDetailOrICS)
	e.GET("/events/nabor-days", h.EventNaborDays)
	e.GET("/events/community-bingos", h.EventCommunityBingos)
	e.GET("/events/wagon-rides", h.EventWagonRides)
	// Permanent redirect from the old sleigh-rides slug to wagon-rides
	e.GET("/events/sleigh-rides", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/events/wagon-rides")
	})
	e.GET("/gallery", h.Gallery)
	e.GET("/faq", h.FAQ)
	e.GET("/support", h.Support)
	e.GET("/grand-marshals", h.GrandMarshals)
	e.GET("/contact", h.Contact)

	// 301 redirects from the old Weebly URLs Google has indexed. The
	// first four (frequently-asked-questions, relocation, index,
	// meeting--member-information) are the exact slugs from the
	// 2026-05-11 Search Console "Not found (404)" report. The
	// remaining entries are belt-and-suspenders mappings for plausible
	// alternate paths so bookmarks and old inbound links still work.
	for from, to := range map[string]string{
		"/index.html":                       "/",
		"/home":                             "/",
		"/home.html":                        "/",
		"/frequently-asked-questions.html":  "/faq",
		"/faq.html":                         "/faq",
		"/meeting--member-information.html": "/contact",
		"/meeting-member-information.html":  "/contact",
		"/contact.html":                     "/contact",
		"/relocation.html":                  "/",
		"/events.html":                      "/events",
		"/gallery.html":                     "/gallery",
		"/nabor-days.html":                  "/events/nabor-days",
		"/community-bingos.html":            "/events/community-bingos",
		"/bingo.html":                       "/events/community-bingos",
		"/sleigh-rides.html":                "/events/wagon-rides",
		"/wagon-rides.html":                 "/events/wagon-rides",
		"/grand-marshals.html":              "/grand-marshals",
		"/support.html":                     "/support",
		"/donate.html":                      "/support",
	} {
		dest := to
		e.GET(from, func(c echo.Context) error {
			return c.Redirect(http.StatusMovedPermanently, dest)
		})
	}

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
