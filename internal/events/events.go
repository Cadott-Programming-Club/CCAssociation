// Package events loads, indexes, and serves the JSON-driven event calendar
// data embedded in the binary at build time.
package events

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"time"
)

//go:embed events.json
var embedded []byte

// EmbeddedJSON returns the raw embedded events JSON bytes.
func EmbeddedJSON() []byte { return embedded }

// Location is an event venue.
type Location struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	MapsURL string `json:"mapsURL"`
}

// Event is one occurrence of a community event. Field names mirror
// schema.org/Event so the JSON can be reused for JSON-LD.
type Event struct {
	Slug        string    `json:"slug"`
	Name        string    `json:"name"`
	StartDate   time.Time `json:"startDate"`
	EndDate     time.Time `json:"endDate"`
	AllDay      bool      `json:"allDay"`
	Category    string    `json:"category"`
	Color       string    `json:"color"`
	Summary     string    `json:"summary"`
	Description string    `json:"description"`
	Location    Location  `json:"location"`
	Image       string    `json:"image"`
	DetailURL   string    `json:"detailURL,omitempty"`
	FacebookURL string    `json:"facebookURL,omitempty"`
	Tags        []string  `json:"tags,omitempty"`
}

// URL returns the canonical detail URL for this event. Signature events
// keep their hand-coded detail pages via DetailURL; one-off events fall
// back to the auto-generated /events/c/<slug>.
func (e Event) URL() string {
	if e.DetailURL != "" {
		return e.DetailURL
	}
	return "/events/c/" + e.Slug
}

// ICSURL returns the per-event .ics download path.
func (e Event) ICSURL() string {
	return "/events/c/" + e.Slug + ".ics"
}

// IsAutoDetail reports whether this event uses the auto-generated detail
// page (true) vs. an existing hand-coded page (false).
func (e Event) IsAutoDetail() bool { return e.DetailURL == "" }

// jsonRoot mirrors the top-level shape of events.json.
type jsonRoot struct {
	Events []Event `json:"events"`
}

// Load parses raw JSON bytes, validates basic shape, and returns events
// sorted ascending by StartDate.
func Load(data []byte) ([]Event, error) {
	var root jsonRoot
	if err := json.Unmarshal(data, &root); err != nil {
		return nil, fmt.Errorf("events: parse JSON: %w", err)
	}
	seen := make(map[string]struct{}, len(root.Events))
	for i, e := range root.Events {
		if e.Slug == "" {
			return nil, fmt.Errorf("events: entry %d missing slug", i)
		}
		if _, dup := seen[e.Slug]; dup {
			return nil, fmt.Errorf("events: duplicate slug %q", e.Slug)
		}
		seen[e.Slug] = struct{}{}
		if e.StartDate.IsZero() {
			return nil, fmt.Errorf("events: %s missing startDate", e.Slug)
		}
		if e.EndDate.IsZero() {
			root.Events[i].EndDate = e.StartDate
		}
	}
	sort.SliceStable(root.Events, func(i, j int) bool {
		return root.Events[i].StartDate.Before(root.Events[j].StartDate)
	})
	return root.Events, nil
}

// Upcoming returns up to n future events (EndDate >= now) sorted by StartDate.
// Input must already be sorted by Load.
func Upcoming(events []Event, now time.Time, n int) []Event {
	if n <= 0 {
		return nil
	}
	out := make([]Event, 0, n)
	for _, e := range events {
		if !e.EndDate.Before(now) {
			out = append(out, e)
			if len(out) == n {
				break
			}
		}
	}
	return out
}

// Find returns the event with the given slug, or false if not found.
func Find(events []Event, slug string) (Event, bool) {
	for _, e := range events {
		if e.Slug == slug {
			return e, true
		}
	}
	return Event{}, false
}

// ByDay returns events whose [StartDate, EndDate] range covers the given
// calendar day in loc. day is interpreted as midnight-to-midnight in loc.
func ByDay(events []Event, day time.Time, loc *time.Location) []Event {
	if loc == nil {
		loc = day.Location()
	}
	dayStart := time.Date(day.Year(), day.Month(), day.Day(), 0, 0, 0, 0, loc)
	dayEnd := dayStart.Add(24 * time.Hour)
	var out []Event
	for _, e := range events {
		if e.StartDate.Before(dayEnd) && !e.EndDate.Before(dayStart) {
			out = append(out, e)
		}
	}
	return out
}

// Month returns the subset of events that overlap the calendar grid
// rendered for (year, month) in loc. The grid is a Sunday-start 6-row
// span covering up to 42 days, which is why we widen the window past the
// strict month boundaries.
func Month(events []Event, year int, month time.Month, loc *time.Location) []Event {
	if loc == nil {
		loc = time.UTC
	}
	gridStart, gridEnd := MonthGridRange(year, month, loc)
	var out []Event
	for _, e := range events {
		if e.StartDate.Before(gridEnd) && !e.EndDate.Before(gridStart) {
			out = append(out, e)
		}
	}
	return out
}

// MonthGridRange returns the [start, end) range covered by the Sunday-
// start 6-row calendar grid for (year, month) in loc.
func MonthGridRange(year int, month time.Month, loc *time.Location) (time.Time, time.Time) {
	first := time.Date(year, month, 1, 0, 0, 0, 0, loc)
	// Walk back to Sunday.
	offset := int(first.Weekday()) // Sunday=0
	gridStart := first.AddDate(0, 0, -offset)
	gridEnd := gridStart.AddDate(0, 0, 42) // 6 rows of 7 days
	return gridStart, gridEnd
}

// GroupByDay groups events by local-date string (YYYY-MM-DD) in loc and
// returns an ordered list of days plus the per-day events. Multi-day
// events appear under every day they cover within [from, to).
func GroupByDay(events []Event, from, to time.Time, loc *time.Location) []DayGroup {
	if loc == nil {
		loc = time.UTC
	}
	if !from.Before(to) {
		return nil
	}
	dayCount := int(to.Sub(from) / (24 * time.Hour))
	groups := make([]DayGroup, 0, dayCount)
	for i := 0; i < dayCount; i++ {
		day := from.AddDate(0, 0, i)
		day = time.Date(day.Year(), day.Month(), day.Day(), 0, 0, 0, 0, loc)
		matches := ByDay(events, day, loc)
		if len(matches) > 0 {
			groups = append(groups, DayGroup{Date: day, Events: matches})
		}
	}
	return groups
}

// DayGroup is a single day's bucket of events.
type DayGroup struct {
	Date   time.Time
	Events []Event
}

// ColorClass returns a CSS class fragment for the event's color, falling
// back to teal when the color is unknown.
func (e Event) ColorClass() string {
	c := strings.ToLower(e.Color)
	switch c {
	case "gold", "teal", "purple", "coral", "lime":
		return c
	}
	return "teal"
}

// AccentTextClass returns the Tailwind text-* class matching the event's
// accent color — used for the small uppercase date/location label that
// sits above the title on event cards (matches the /events page pattern).
func (e Event) AccentTextClass() string {
	switch e.ColorClass() {
	case "gold":
		return "text-cca-gold-500"
	case "purple":
		return "text-cca-purple"
	case "coral":
		return "text-cca-coral"
	case "lime":
		return "text-cca-lime"
	}
	return "text-cca-teal-600"
}

// AccentBgClass returns the Tailwind bg-* class matching the event's
// accent color — used as a fallback placeholder when no image is set.
func (e Event) AccentBgClass() string {
	switch e.ColorClass() {
	case "gold":
		return "bg-cca-gold-500"
	case "purple":
		return "bg-cca-purple"
	case "coral":
		return "bg-cca-coral"
	case "lime":
		return "bg-cca-lime"
	}
	return "bg-cca-teal-600"
}

// DateRangeLabel returns a compact human range like "JUL 23" or
// "JUL 23–26" (uppercase, en-dash) matching the site's uppercase
// section-label convention.
func (e Event) DateRangeLabel() string {
	s := e.StartDate
	t := e.EndDate
	if t.IsZero() || sameDay(s, t) {
		return strings.ToUpper(s.Format("Jan 2"))
	}
	if s.Year() == t.Year() && s.Month() == t.Month() {
		return strings.ToUpper(s.Format("Jan 2")) + "–" + t.Format("2")
	}
	return strings.ToUpper(s.Format("Jan 2") + "–" + t.Format("Jan 2"))
}

// TimeLabel returns the event's time as a compact string: "All day",
// "5:00 PM", or "5:00 PM – 10:00 PM".
func (e Event) TimeLabel() string {
	if e.AllDay {
		return "All day"
	}
	if e.EndDate.IsZero() || e.StartDate.Equal(e.EndDate) {
		return e.StartDate.Format("3:04 PM")
	}
	if sameDay(e.StartDate, e.EndDate) {
		return e.StartDate.Format("3:04 PM") + " – " + e.EndDate.Format("3:04 PM")
	}
	return e.StartDate.Format("3:04 PM")
}

// WeekdayRangeLabel returns "Thursday" or "Thursday–Sunday" for the
// event's date range. Used as a secondary line on cards.
func (e Event) WeekdayRangeLabel() string {
	s := e.StartDate
	t := e.EndDate
	if t.IsZero() || sameDay(s, t) {
		return s.Format("Monday")
	}
	return s.Format("Monday") + "–" + t.Format("Monday")
}

func sameDay(a, b time.Time) bool {
	return a.Year() == b.Year() && a.YearDay() == b.YearDay()
}
