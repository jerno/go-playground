package nextBusinessDay

import (
	"time"
)

func isBusinessDay(date time.Time) bool {
	wday := date.Weekday()
	if wday == time.Saturday || wday == time.Sunday {
		return false
	}

	return true
}

func NextBusinessDay(date time.Time) time.Time {
	const day = 24 * time.Hour
	for {
		date = date.Add(day)
		if isBusinessDay(date) {
			break
		}
	}

	return date
}
