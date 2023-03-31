package main

import (
	"fmt"
	"strings"
	"time"
)

type Unit int

const (
	Seconds = 1 + iota
	Minutes
	Hours
	Days
	Months
	Years

	fmtOut         string = "%.2f %s"
	errUnknownUnit string = "unknown unit \"%s\""
)

var unitLU = map[Unit]struct {
	Factor       time.Duration
	UnitSingular string
	UnitPlural   string
}{
	Seconds: {time.Second * 1, "second", "seconds"},
	Minutes: {time.Second * 60, "minute", "minutes"},
	Hours:   {time.Second * 60 * 60, "hour", "hours"},
	Days:    {time.Second * 3600 * 24, "day", "days"},
	Months:  {time.Second * 3600 * 24 * 30, "month", "months"},
}

func (unit Unit) durationAsString(dur time.Duration) string {

	u := unitLU[unit]
	d := float64(dur / u.Factor)
	us := u.UnitPlural
	if d <= 1 {
		us = u.UnitSingular
	}

	s := fmt.Sprintf(fmtOut, d, us)

	return s
}

func parseUnit(unit string) (Unit, error) {

	unit = strings.TrimSpace(unit)
	unit = strings.ToLower(unit)

	switch unit {
	case "s", "secs", "seconds":
		return Seconds, nil
	case "min", "mins", "minutes":
		return Minutes, nil
	case "h", "hours":
		return Hours, nil
	case "d", "days":
		return Days, nil
	case "months":
		return Months, nil
	}

	err := fmt.Errorf(errUnknownUnit, unit)

	return 0, err
}
