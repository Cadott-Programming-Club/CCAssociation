package handler

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"

	"ccassociation/internal/events"
	"ccassociation/templates/components"
	"ccassociation/templates/pages"
)

// chicagoLoc is the calendar's display timezone. Falls back to UTC if
// the tzdata isn't available in the build (vendored on stdlib for Go 1.20+).
var chicagoLoc = func() *time.Location {
	if loc, err := time.LoadLocation("America/Chicago"); err == nil {
		return loc
	}
	return time.UTC
}()

// EventsCalendar renders the month + list view at /events/calendar.
// Query params:
//
//	m=YYYY-MM (default = current month)
//	view=month|list (default = month; CSS chooses list on small screens)
func (h *Handler) EventsCalendar(c echo.Context) error {
	now := time.Now().In(chicagoLoc)
	year, month := now.Year(), now.Month()

	if mq := c.QueryParam("m"); mq != "" {
		if t, err := time.ParseInLocation("2006-01", mq, chicagoLoc); err == nil {
			year, month = t.Year(), t.Month()
		}
	}
	view := c.QueryParam("view")
	if view != "list" {
		view = "month"
	}

	monthEvents := events.Month(h.events, year, month, chicagoLoc)
	gridStart, gridEnd := events.MonthGridRange(year, month, chicagoLoc)
	listGroups := events.GroupByDay(h.events, now.Truncate(24*time.Hour), gridEnd.AddDate(0, 6, 0), chicagoLoc)

	data := pages.CalendarData{
		Year:        year,
		Month:       month,
		Today:       now,
		View:        view,
		GridStart:   gridStart,
		MonthEvents: monthEvents,
		ListGroups:  listGroups,
		Location:    chicagoLoc,
	}
	return pages.EventCalendar(data).Render(c.Request().Context(), c.Response().Writer)
}

// EventsDay returns just the modal body for a single day, loaded via
// HTMX. Path: /events/calendar/day/:date where :date = YYYY-MM-DD.
func (h *Handler) EventsDay(c echo.Context) error {
	day, err := time.ParseInLocation("2006-01-02", c.Param("date"), chicagoLoc)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid date")
	}
	matches := events.ByDay(h.events, day, chicagoLoc)
	return components.EventModal(day, matches).Render(c.Request().Context(), c.Response().Writer)
}

// EventDetailOrICS serves either the auto-generated detail page or the
// per-event .ics file, depending on the path extension. Signature events
// (those with DetailURL set) redirect to their hand-coded page.
func (h *Handler) EventDetailOrICS(c echo.Context) error {
	slug := c.Param("slug")
	if strings.HasSuffix(slug, ".ics") {
		return h.eventICS(c, strings.TrimSuffix(slug, ".ics"))
	}
	return h.eventDetail(c, slug)
}

func (h *Handler) eventDetail(c echo.Context, slug string) error {
	e, ok := events.Find(h.events, slug)
	if !ok {
		return echo.NewHTTPError(http.StatusNotFound, "event not found")
	}
	if e.DetailURL != "" {
		return c.Redirect(http.StatusMovedPermanently, e.DetailURL)
	}
	return pages.EventDetail(e).Render(c.Request().Context(), c.Response().Writer)
}

func (h *Handler) eventICS(c echo.Context, slug string) error {
	e, ok := events.Find(h.events, slug)
	if !ok {
		return echo.NewHTTPError(http.StatusNotFound, "event not found")
	}
	resp := c.Response()
	resp.Header().Set("Content-Type", "text/calendar; charset=utf-8")
	resp.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s.ics"`, e.Slug))
	return events.WriteICS(resp.Writer, e)
}
