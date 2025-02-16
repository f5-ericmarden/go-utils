package util

import (
	"path/filepath"
	"time"
)

type TimeRange struct {
	Months int
	Weeks  int
	Days   int
}

// Given some strings and a range of time from now
// filters those strings based on whether or not the date parsed from those
// strings fall within the range of time from now
func TFilter(filnames []string, tr *TimeRange, ignore bool) []string {
	var results []string
	// now sure why there is a one off error of 2 here... maybe timezones?
	timeHoriz := time.Now().AddDate(-tr.Months, -tr.Weeks, -tr.Days-2)
	for _, filename := range filnames {
		abs, _ := filepath.Abs(filename)
		t, err := parseDateFileName(filename)

		if ignore && err != nil {
			continue
		}

		if t.After(timeHoriz) {
			results = append(results, abs)
		}
	}
	return results
}
