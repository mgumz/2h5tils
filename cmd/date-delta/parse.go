package main

import (
	"fmt"
	"strings"
	"time"
)

// <time-a>[ - ]<time-b>[ in <unit>]
func parseCmdLine(args string) (a, b time.Time, unit Unit, err error) {

	b, unit = time.Now().UTC(), Seconds

	parts := strings.SplitN(args, "in", 2)
	if len(parts) > 1 {
		unit, err = parseUnit(strings.TrimSpace(parts[1]))
		if err != nil {
			return
		}
	}

	parts = strings.SplitN(parts[0], " - ", 2)
	if len(parts) < 1 {
		return
	}

	a, err = parseDateTime(strings.TrimSpace(parts[0]))
	if err != nil || len(parts) == 1 {
		return
	}

	b, err = parseDateTime(strings.TrimSpace(parts[1]))
	if err != nil {
		return
	}

	return a, b, unit, nil
}

func parseDateTime(dt string) (time.Time, error) {

	t := time.Now().UTC()

	// times without date
	if nt, err := time.Parse(":04", dt); err == nil {
		return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), nt.Minute(), 0, 0, time.UTC), nil
	}
	if nt, err := time.Parse("15:04", dt); err == nil {
		return time.Date(t.Year(), t.Month(), t.Day(), nt.Hour(), nt.Minute(), 0, 0, time.UTC), nil
	}
	if nt, err := time.Parse("15:04:05", dt); err == nil {
		return time.Date(t.Year(), t.Month(), t.Day(), nt.Hour(), nt.Minute(), nt.Second(), 0, time.UTC), nil
	}
	if nt, err := time.Parse(time.Kitchen, dt); err == nil {
		return time.Date(t.Year(), t.Month(), t.Day(), nt.Hour(), nt.Minute(), 0, 0, time.UTC), nil
	}

	dt = strings.Map(func(r rune) rune {
		switch r {
		case '/':
			return '-'
		case 'T':
			return ' '
		default:
			return r
		}
	}, dt)

	// datetime formats
	formats := []string{
		"2006-01-02",
		"2006-01-02 03:04",
		"2006-01-02 03:04:05",
		time.RFC3339,
	}
	for _, f := range formats {
		if nt, err := time.Parse(f, dt); err == nil {
			return nt, nil
		}
	}

	return t, fmt.Errorf("can't parse %q", dt)
}
