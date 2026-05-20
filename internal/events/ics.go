package events

import (
	"fmt"
	"io"
	"strings"
	"time"
)

// icsProdID identifies the calendar producer in generated .ics files.
const icsProdID = "-//Cadott Community Association//Calendar//EN"

// WriteICS writes an RFC 5545 VCALENDAR/VEVENT for a single event.
// Lines use CRLF endings; SUMMARY/DESCRIPTION/LOCATION text is escaped
// per RFC 5545 § 3.3.11.
func WriteICS(w io.Writer, e Event) error {
	uid := e.Slug + "@cadottcommunity.com"
	dtstamp := time.Now().UTC().Format("20060102T150405Z")

	var b strings.Builder
	write := func(s string) { b.WriteString(s); b.WriteString("\r\n") }

	write("BEGIN:VCALENDAR")
	write("VERSION:2.0")
	write("PRODID:" + icsProdID)
	write("CALSCALE:GREGORIAN")
	write("METHOD:PUBLISH")
	write("BEGIN:VEVENT")
	write("UID:" + uid)
	write("DTSTAMP:" + dtstamp)
	if e.AllDay {
		write("DTSTART;VALUE=DATE:" + e.StartDate.Format("20060102"))
		// All-day DTEND is exclusive per RFC 5545; if not provided, end == start+1.
		end := e.EndDate
		if !end.After(e.StartDate) {
			end = e.StartDate.AddDate(0, 0, 1)
		} else {
			end = end.AddDate(0, 0, 1)
		}
		write("DTEND;VALUE=DATE:" + end.Format("20060102"))
	} else {
		write("DTSTART:" + e.StartDate.UTC().Format("20060102T150405Z"))
		end := e.EndDate
		if end.IsZero() || !end.After(e.StartDate) {
			end = e.StartDate.Add(1 * time.Hour)
		}
		write("DTEND:" + end.UTC().Format("20060102T150405Z"))
	}
	write("SUMMARY:" + icsEscape(e.Name))
	if e.Summary != "" || e.Description != "" {
		desc := e.Summary
		if e.Description != "" {
			desc = e.Description
		}
		write("DESCRIPTION:" + icsEscape(desc))
	}
	if e.Location.Name != "" || e.Location.Address != "" {
		loc := strings.TrimSpace(e.Location.Name + " " + e.Location.Address)
		write("LOCATION:" + icsEscape(loc))
	}
	if e.DetailURL != "" || e.Slug != "" {
		write("URL:" + canonicalEventURL(e))
	}
	write("END:VEVENT")
	write("END:VCALENDAR")

	if _, err := io.WriteString(w, b.String()); err != nil {
		return fmt.Errorf("write ics: %w", err)
	}
	return nil
}

// icsEscape escapes a string for RFC 5545 TEXT fields.
func icsEscape(s string) string {
	r := strings.NewReplacer(
		"\\", `\\`,
		";", `\;`,
		",", `\,`,
		"\n", `\n`,
		"\r", "",
	)
	return r.Replace(s)
}

// canonicalEventURL produces an absolute URL for the event suitable for
// the .ics URL field and share links.
func canonicalEventURL(e Event) string {
	const base = "https://www.cadottcommunity.com"
	if e.DetailURL != "" {
		if strings.HasPrefix(e.DetailURL, "http") {
			return e.DetailURL
		}
		return base + e.DetailURL
	}
	return base + "/events/c/" + e.Slug
}
