package events

import (
	"net/url"
	"strings"
	"time"
)

// CanonicalURL returns the canonical absolute URL for an event.
func CanonicalURL(e Event) string { return canonicalEventURL(e) }

// GoogleCalendarURL builds a "create new event" link for Google Calendar
// per https://add-to-calendar-pro.com/articles/ics-format-google-calendar.
func GoogleCalendarURL(e Event) string {
	q := url.Values{}
	q.Set("action", "TEMPLATE")
	q.Set("text", e.Name)
	q.Set("dates", googleDates(e))
	desc := e.Summary
	if e.Description != "" {
		desc = e.Description
	}
	if desc != "" {
		q.Set("details", desc)
	}
	if loc := strings.TrimSpace(e.Location.Name + " " + e.Location.Address); loc != "" {
		q.Set("location", loc)
	}
	return "https://calendar.google.com/calendar/render?" + q.Encode()
}

// googleDates returns the dates= value for a Google Calendar template URL.
func googleDates(e Event) string {
	if e.AllDay {
		start := e.StartDate.Format("20060102")
		end := e.EndDate
		if end.IsZero() || !end.After(e.StartDate) {
			end = e.StartDate.AddDate(0, 0, 1)
		} else {
			end = end.AddDate(0, 0, 1)
		}
		return start + "/" + end.Format("20060102")
	}
	start := e.StartDate.UTC().Format("20060102T150405Z")
	end := e.EndDate
	if end.IsZero() || !end.After(e.StartDate) {
		end = e.StartDate.Add(1 * time.Hour)
	}
	return start + "/" + end.UTC().Format("20060102T150405Z")
}

// FacebookShareURL returns the Facebook share dialog URL for a public
// event page URL. Facebook scrapes the destination's Open Graph tags for
// the preview card.
func FacebookShareURL(publicURL string) string {
	q := url.Values{}
	q.Set("u", publicURL)
	return "https://www.facebook.com/sharer/sharer.php?" + q.Encode()
}

// MailtoShareURL builds a mailto: link prefilled with the event subject
// and a short body containing the summary and public URL.
func MailtoShareURL(e Event, publicURL string) string {
	q := url.Values{}
	q.Set("subject", e.Name)
	body := e.Summary
	if body != "" {
		body += "\n\n"
	}
	body += publicURL
	q.Set("body", body)
	return "mailto:?" + q.Encode()
}
