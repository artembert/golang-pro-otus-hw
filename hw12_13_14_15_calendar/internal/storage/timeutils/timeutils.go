package timeutils

import (
	"time"
)

// DT common values ...
const (
	beginningOfDayHour   = 0
	beginningOfDayMinute = 0
	beginningOfDaySecond = 0

	endOfDayHour   = 23
	endOfDayMinute = 59
	endOfDaySecond = 59
)

// Date ...
const (
	DaysInWeek  = 7
	DaysInMonth = 30
)

// BeginningOfDay returns the beginning of the day (00:00:00:00000) ...
func BeginningOfDay(t time.Time) time.Time {
	y, m, d := t.Date()

	return time.Date(y, m, d,
		beginningOfDayHour, beginningOfDayMinute, beginningOfDaySecond, beginningOfDaySecond, t.Location())
}

// EndOfDay returns the end of the day (23:59:59:99999) ...
func EndOfDay(t time.Time) time.Time {
	y, m, d := t.Date()

	return time.Date(y, m, d,
		endOfDayHour, endOfDayMinute, endOfDaySecond, int(time.Second-time.Nanosecond), t.Location())
}
