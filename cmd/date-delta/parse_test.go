package main

import (
	"testing"
	"time"
)

func TestParseCmdLine(t *testing.T) {

	type Expected struct {
		TimeA time.Time
		TimeB time.Time
		Unit  Unit
	}

	now := time.Now().UTC()

	fixture := []struct {
		Input    string
		Expected Expected
	}{
		{"2023-04-01", Expected{
			time.Date(2023, 04, 01, 0, 0, 0, 0, time.UTC),
			now,
			Seconds,
		}},
		{"2023-04-01 in minutes", Expected{
			time.Date(2023, 04, 01, 0, 0, 0, 0, time.UTC),
			now,
			Minutes,
		}},
		{"2023-04-01 - 2023-04-02", Expected{
			time.Date(2023, 04, 01, 0, 0, 0, 0, time.UTC),
			time.Date(2023, 04, 02, 0, 0, 0, 0, time.UTC),
			Seconds,
		}},
		{"2023-04-01 - 2023/04/02", Expected{
			time.Date(2023, 04, 01, 0, 0, 0, 0, time.UTC),
			time.Date(2023, 04, 02, 0, 0, 0, 0, time.UTC),
			Seconds,
		}},
		{"15:04", Expected{
			time.Date(now.Year(), now.Month(), now.Day(), 15, 4, 0, 0, time.UTC),
			now,
			Seconds,
		}},
	}

	for i, f := range fixture {

		a, b, u, err := parseCmdLine(f.Input)
		i++

		if err != nil {
			t.Fatalf("%d: \"%s\": error: %s", i, f.Input, err)
		}
		if da := a.Sub(f.Expected.TimeA).Abs(); da >= 1*time.Second {
			t.Fatalf("%d timeA: %d \"%s\": expected %q, got %q", i, da*time.Second, f.Input, f.Expected.TimeA, a)
		}
		if db := b.Sub(f.Expected.TimeB).Abs(); db >= 1*time.Second {
			t.Fatalf("%d timeB: %d \"%s\": expected %q, got %q", i, db*time.Second, f.Input, f.Expected.TimeB, b)
		}
		if u != f.Expected.Unit {
			t.Fatalf("%d unit: %s: expected %q, got %q", i, f.Input, f.Expected.Unit, u)
		}

		t.Logf("%d: \"%s\": a: %s, b: %s, u: %d", i, f.Input, a, b, u)
	}

}
