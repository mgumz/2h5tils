package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const (
    errFmtCantParse string = "can't parse %q"
)

func main() {

	if len(os.Args) < 2 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		printUsage()
		os.Exit(0)
	}
    if os.Args[1] == "-v" || os.Args[1] == "--version" {
		printVersion()
		os.Exit(0)
    }

	spec := strings.TrimSpace(os.Args[1])
	t := time.Now()
	when, err := parse(spec, t)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	dur := when.Sub(t)

	fmt.Println("go on at", when, "; in", dur)
	time.Sleep(dur)
}

func parse(spec string, t time.Time) (time.Time, error) {
	if len(spec) == 0 {
		return t, fmt.Errorf(errFmtCantParse, spec)
	}

	// shortcuts
	if spec[0] == '@' {
		return parseAtSpec(spec, t)
	}

	// try duration
	if d, _ := time.ParseDuration(spec); d > 0 {
		return t.Add(d), nil
	}

	return parseAsTime(spec, t)
}

func parseAtSpec(spec string, t time.Time) (time.Time, error) {
	switch spec {
	case "@hour":
		return t.Add(time.Hour).Truncate(time.Hour), nil
	case "@minute":
		return t.Add(time.Minute).Truncate(time.Minute), nil
	case "@tens":
		return t.Add(10 * time.Minute).Truncate(10 * time.Minute), nil
	case "@quarter":
		return t.Add(15 * time.Minute).Truncate(15 * time.Minute), nil
	case "@midnight":
		return t.Add(24 * time.Hour).Truncate(24 * time.Hour), nil
	case "@noon":
		return parseAsTime("12:00", t)
	}

	return time.Time{}, fmt.Errorf(errFmtCantParse, spec)
}

func parseAsTime(spec string, t time.Time) (time.Time, error) {
	if nt, err := time.Parse(":04", spec); err == nil {
		return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), nt.Minute(), 0, 0, t.Location()), nil
	}
	if nt, err := time.Parse("15:04", spec); err == nil {
		return time.Date(t.Year(), t.Month(), t.Day(), nt.Hour(), nt.Minute(), 0, 0, t.Location()), nil
	}
	if nt, err := time.Parse("15:04:05", spec); err == nil {
		return time.Date(t.Year(), t.Month(), t.Day(), nt.Hour(), nt.Minute(), nt.Second(), 0, t.Location()), nil
	}
	if nt, err := time.Parse(time.Kitchen, spec); err == nil {
		return time.Date(t.Year(), t.Month(), t.Day(), nt.Hour(), nt.Minute(), 0, 0, t.Location()), nil
	}
	if nt, err := time.Parse(time.RFC3339, spec); err == nil {
		return nt, nil
	}
	return t, fmt.Errorf(errFmtCantParse, spec)
}
