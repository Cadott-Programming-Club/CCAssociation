package events

import (
	"bytes"
	"net/url"
	"strings"
	"testing"
	"time"
)

func mustTime(t *testing.T, s string) time.Time {
	t.Helper()
	v, err := time.Parse(time.RFC3339, s)
	if err != nil {
		t.Fatalf("parse %q: %v", s, err)
	}
	return v
}

func TestLoad_ParsesTZAndSorts(t *testing.T) {
	raw := []byte(`{"events":[
		{"slug":"b","name":"B","startDate":"2026-08-01T10:00:00-05:00","endDate":"2026-08-01T11:00:00-05:00"},
		{"slug":"a","name":"A","startDate":"2026-07-23T17:00:00-05:00","endDate":"2026-07-26T22:00:00-05:00"},
		{"slug":"c","name":"C","startDate":"2026-12-12T17:00:00-06:00","endDate":"2026-12-12T21:00:00-06:00"}
	]}`)

	got, err := Load(raw)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if len(got) != 3 {
		t.Fatalf("got %d events, want 3", len(got))
	}
	wantOrder := []string{"a", "b", "c"}
	for i, want := range wantOrder {
		if got[i].Slug != want {
			t.Errorf("position %d: got slug %q, want %q", i, got[i].Slug, want)
		}
	}
	// Timezone should be preserved on parse: -05:00 offset for entry "a".
	if _, offset := got[0].StartDate.Zone(); offset != -5*3600 {
		t.Errorf("entry a startDate offset = %d, want -18000", offset)
	}
}

func TestLoad_RejectsDuplicateSlugs(t *testing.T) {
	raw := []byte(`{"events":[
		{"slug":"x","name":"X","startDate":"2026-01-01T00:00:00Z"},
		{"slug":"x","name":"X2","startDate":"2026-02-01T00:00:00Z"}
	]}`)
	if _, err := Load(raw); err == nil {
		t.Fatal("expected duplicate-slug error, got nil")
	}
}

func TestLoad_RejectsMissingSlug(t *testing.T) {
	raw := []byte(`{"events":[{"name":"NoSlug","startDate":"2026-01-01T00:00:00Z"}]}`)
	if _, err := Load(raw); err == nil {
		t.Fatal("expected missing-slug error, got nil")
	}
}

func TestLoad_RejectsMissingStart(t *testing.T) {
	raw := []byte(`{"events":[{"slug":"x","name":"X"}]}`)
	if _, err := Load(raw); err == nil {
		t.Fatal("expected missing-start error, got nil")
	}
}

func TestUpcoming_FiltersPastAndCaps(t *testing.T) {
	events := []Event{
		{Slug: "past", StartDate: mustTime(t, "2026-01-01T10:00:00Z"), EndDate: mustTime(t, "2026-01-01T11:00:00Z")},
		{Slug: "soon", StartDate: mustTime(t, "2026-06-01T10:00:00Z"), EndDate: mustTime(t, "2026-06-01T11:00:00Z")},
		{Slug: "later", StartDate: mustTime(t, "2026-07-01T10:00:00Z"), EndDate: mustTime(t, "2026-07-01T11:00:00Z")},
		{Slug: "latest", StartDate: mustTime(t, "2026-08-01T10:00:00Z"), EndDate: mustTime(t, "2026-08-01T11:00:00Z")},
	}
	now := mustTime(t, "2026-05-20T00:00:00Z")
	got := Upcoming(events, now, 2)
	if len(got) != 2 {
		t.Fatalf("got %d, want 2", len(got))
	}
	if got[0].Slug != "soon" || got[1].Slug != "later" {
		t.Errorf("got slugs [%s %s], want [soon later]", got[0].Slug, got[1].Slug)
	}
}

func TestUpcoming_IncludesEventsEndingExactlyAtNow(t *testing.T) {
	events := []Event{
		{Slug: "running", StartDate: mustTime(t, "2026-05-19T10:00:00Z"), EndDate: mustTime(t, "2026-05-21T22:00:00Z")},
	}
	now := mustTime(t, "2026-05-20T00:00:00Z")
	got := Upcoming(events, now, 4)
	if len(got) != 1 {
		t.Fatalf("in-progress event filtered out: %#v", got)
	}
}

func TestByDay_OverlapsMultiDayEvent(t *testing.T) {
	naborStart := mustTime(t, "2026-07-23T17:00:00-05:00")
	naborEnd := mustTime(t, "2026-07-26T22:00:00-05:00")
	events := []Event{{Slug: "nabor", StartDate: naborStart, EndDate: naborEnd}}
	loc := naborStart.Location()

	for _, day := range []time.Time{
		time.Date(2026, 7, 23, 12, 0, 0, 0, loc),
		time.Date(2026, 7, 24, 12, 0, 0, 0, loc),
		time.Date(2026, 7, 25, 12, 0, 0, 0, loc),
		time.Date(2026, 7, 26, 12, 0, 0, 0, loc),
	} {
		got := ByDay(events, day, loc)
		if len(got) != 1 {
			t.Errorf("day %s: got %d, want 1", day.Format("2006-01-02"), len(got))
		}
	}
	if got := ByDay(events, time.Date(2026, 7, 22, 12, 0, 0, 0, loc), loc); len(got) != 0 {
		t.Errorf("day before: got %d, want 0", len(got))
	}
	if got := ByDay(events, time.Date(2026, 7, 27, 12, 0, 0, 0, loc), loc); len(got) != 0 {
		t.Errorf("day after: got %d, want 0", len(got))
	}
}

func TestMonth_IncludesGridSpillover(t *testing.T) {
	// Event entirely in early July; we render the JUNE grid, which spills
	// into early July via the last row of the Sunday-start grid.
	events := []Event{{
		Slug:      "july-early",
		StartDate: mustTime(t, "2026-07-02T10:00:00Z"),
		EndDate:   mustTime(t, "2026-07-02T12:00:00Z"),
	}}
	got := Month(events, 2026, time.June, time.UTC)
	if len(got) != 1 {
		t.Errorf("expected event to spill into June grid; got %d", len(got))
	}
}

func TestMonthGridRange_SundayStart(t *testing.T) {
	// July 2026: July 1 is a Wednesday. Grid should start on Sunday June 28.
	start, end := MonthGridRange(2026, time.July, time.UTC)
	if want := time.Date(2026, 6, 28, 0, 0, 0, 0, time.UTC); !start.Equal(want) {
		t.Errorf("grid start = %s, want %s", start, want)
	}
	if d := end.Sub(start); d != 42*24*time.Hour {
		t.Errorf("grid span = %s, want 42d", d)
	}
}

func TestFind(t *testing.T) {
	events := []Event{{Slug: "a"}, {Slug: "b"}}
	if got, ok := Find(events, "b"); !ok || got.Slug != "b" {
		t.Errorf("Find(b) = (%v, %v)", got, ok)
	}
	if _, ok := Find(events, "missing"); ok {
		t.Error("Find(missing) reported found")
	}
}

func TestEvent_URLAndAutoDetail(t *testing.T) {
	signature := Event{Slug: "nabor-days-2026", DetailURL: "/events/nabor-days"}
	oneOff := Event{Slug: "church-dinner"}
	if got := signature.URL(); got != "/events/nabor-days" {
		t.Errorf("signature.URL() = %q", got)
	}
	if got := oneOff.URL(); got != "/events/c/church-dinner" {
		t.Errorf("oneOff.URL() = %q", got)
	}
	if signature.IsAutoDetail() {
		t.Error("signature.IsAutoDetail() = true, want false")
	}
	if !oneOff.IsAutoDetail() {
		t.Error("oneOff.IsAutoDetail() = false, want true")
	}
	if got := oneOff.ICSURL(); got != "/events/c/church-dinner.ics" {
		t.Errorf("ICSURL = %q", got)
	}
}

func TestColorClass_FallsBackToTeal(t *testing.T) {
	cases := map[string]string{
		"gold":       "gold",
		"TEAL":       "teal",
		"purple":     "purple",
		"":           "teal",
		"chartreuse": "teal",
	}
	for in, want := range cases {
		got := Event{Color: in}.ColorClass()
		if got != want {
			t.Errorf("ColorClass(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestWriteICS_RFC5545Shape(t *testing.T) {
	e := Event{
		Slug:      "test-event",
		Name:      "Test, Event; with chars",
		StartDate: mustTime(t, "2026-07-23T17:00:00-05:00"),
		EndDate:   mustTime(t, "2026-07-26T22:00:00-05:00"),
		Summary:   "Line one\nline two; with semi",
		Location:  Location{Name: "Riverview Park", Address: "Cadott, WI"},
		DetailURL: "/events/nabor-days",
	}
	var buf bytes.Buffer
	if err := WriteICS(&buf, e); err != nil {
		t.Fatalf("WriteICS: %v", err)
	}
	out := buf.String()

	// CRLF endings.
	if !strings.Contains(out, "\r\n") {
		t.Error("expected CRLF line endings")
	}
	if strings.Contains(out, "\n") && !strings.Contains(out, "\r\n") {
		t.Error("found bare LF without CR")
	}

	// Required envelope.
	for _, want := range []string{
		"BEGIN:VCALENDAR\r\n",
		"VERSION:2.0\r\n",
		"PRODID:",
		"BEGIN:VEVENT\r\n",
		"UID:test-event@cadottcommunity.com\r\n",
		"DTSTAMP:",
		"DTSTART:20260723T220000Z\r\n", // UTC equivalent of 17:00-05:00
		"DTEND:20260727T030000Z\r\n",   // UTC equivalent of 22:00-05:00
		"END:VEVENT\r\n",
		"END:VCALENDAR\r\n",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("missing %q in:\n%s", want, out)
		}
	}

	// Escaping in SUMMARY: , and ; must be backslash-escaped.
	if !strings.Contains(out, `SUMMARY:Test\, Event\; with chars`) {
		t.Errorf("SUMMARY not escaped:\n%s", out)
	}
	// Escaping in DESCRIPTION: newline becomes \n literally.
	if !strings.Contains(out, `DESCRIPTION:Line one\nline two\; with semi`) {
		t.Errorf("DESCRIPTION not escaped:\n%s", out)
	}
	// LOCATION present.
	if !strings.Contains(out, "LOCATION:Riverview Park Cadott\\, WI\r\n") {
		t.Errorf("LOCATION not present/escaped:\n%s", out)
	}
	// URL is absolute.
	if !strings.Contains(out, "URL:https://www.cadottcommunity.com/events/nabor-days\r\n") {
		t.Errorf("URL not absolute:\n%s", out)
	}
}

func TestWriteICS_AllDay(t *testing.T) {
	e := Event{
		Slug:      "all-day",
		Name:      "Picnic",
		AllDay:    true,
		StartDate: time.Date(2026, 8, 15, 0, 0, 0, 0, time.UTC),
		EndDate:   time.Date(2026, 8, 15, 0, 0, 0, 0, time.UTC),
	}
	var buf bytes.Buffer
	if err := WriteICS(&buf, e); err != nil {
		t.Fatalf("WriteICS: %v", err)
	}
	out := buf.String()
	if !strings.Contains(out, "DTSTART;VALUE=DATE:20260815\r\n") {
		t.Errorf("missing date-only DTSTART:\n%s", out)
	}
	if !strings.Contains(out, "DTEND;VALUE=DATE:20260816\r\n") {
		t.Errorf("missing exclusive DTEND:\n%s", out)
	}
}

func TestGoogleCalendarURL_Format(t *testing.T) {
	e := Event{
		Slug:      "nabor-days-2026",
		Name:      "Nabor Days",
		StartDate: mustTime(t, "2026-07-23T17:00:00-05:00"),
		EndDate:   mustTime(t, "2026-07-26T22:00:00-05:00"),
		Summary:   "Four days of fun",
		Location:  Location{Name: "Riverview Park", Address: "Cadott, WI"},
	}
	got := GoogleCalendarURL(e)
	u, err := url.Parse(got)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if u.Host != "calendar.google.com" || u.Path != "/calendar/render" {
		t.Errorf("unexpected URL: %s", got)
	}
	q := u.Query()
	if q.Get("action") != "TEMPLATE" {
		t.Errorf("action = %q, want TEMPLATE", q.Get("action"))
	}
	if q.Get("text") != "Nabor Days" {
		t.Errorf("text = %q", q.Get("text"))
	}
	if q.Get("dates") != "20260723T220000Z/20260727T030000Z" {
		t.Errorf("dates = %q", q.Get("dates"))
	}
	if q.Get("location") != "Riverview Park Cadott, WI" {
		t.Errorf("location = %q", q.Get("location"))
	}
	if q.Get("details") != "Four days of fun" {
		t.Errorf("details = %q", q.Get("details"))
	}
}

func TestFacebookAndMailtoShareURL(t *testing.T) {
	e := Event{Slug: "x", Name: "Bingo Night", Summary: "Doors at 6."}
	pub := "https://www.cadottcommunity.com/events/c/x"

	fb := FacebookShareURL(pub)
	if !strings.HasPrefix(fb, "https://www.facebook.com/sharer/sharer.php?") {
		t.Errorf("unexpected fb url: %s", fb)
	}
	u, _ := url.Parse(fb)
	if u.Query().Get("u") != pub {
		t.Errorf("fb u= %q", u.Query().Get("u"))
	}

	m := MailtoShareURL(e, pub)
	if !strings.HasPrefix(m, "mailto:?") {
		t.Errorf("unexpected mailto: %s", m)
	}
	mu, _ := url.Parse(m)
	if mu.Query().Get("subject") != "Bingo Night" {
		t.Errorf("subject = %q", mu.Query().Get("subject"))
	}
	if !strings.Contains(mu.Query().Get("body"), pub) {
		t.Errorf("body missing public url: %q", mu.Query().Get("body"))
	}
}

func TestEmbeddedJSON_Loads(t *testing.T) {
	got, err := Load(EmbeddedJSON())
	if err != nil {
		t.Fatalf("Load(EmbeddedJSON): %v", err)
	}
	if len(got) < 3 {
		t.Errorf("expected at least 3 seeded events, got %d", len(got))
	}
}

func TestGroupByDay_BucketsMultiDay(t *testing.T) {
	loc := time.UTC
	events := []Event{
		{Slug: "multi", StartDate: time.Date(2026, 7, 23, 17, 0, 0, 0, loc), EndDate: time.Date(2026, 7, 25, 22, 0, 0, 0, loc)},
		{Slug: "single", StartDate: time.Date(2026, 7, 24, 12, 0, 0, 0, loc), EndDate: time.Date(2026, 7, 24, 14, 0, 0, 0, loc)},
	}
	from := time.Date(2026, 7, 22, 0, 0, 0, 0, loc)
	to := time.Date(2026, 7, 27, 0, 0, 0, 0, loc)
	groups := GroupByDay(events, from, to, loc)

	wantDays := map[string]int{
		"2026-07-23": 1, // multi only
		"2026-07-24": 2, // multi + single
		"2026-07-25": 1, // multi only
	}
	gotDays := map[string]int{}
	for _, g := range groups {
		gotDays[g.Date.Format("2006-01-02")] = len(g.Events)
	}
	for day, want := range wantDays {
		if gotDays[day] != want {
			t.Errorf("%s: got %d events, want %d", day, gotDays[day], want)
		}
	}
	if _, present := gotDays["2026-07-22"]; present {
		t.Error("empty day 2026-07-22 should be omitted")
	}
}
