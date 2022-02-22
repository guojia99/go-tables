package tables

import (
	"fmt"
	"time"
)

const (
	second = 1
	minute = second * 60
	hour   = minute * 60
	day    = hour * 24
	month  = day * 30
	year   = month * 12
)

var timeFormal = map[bool]map[string]string{
	false: {
		"second": " second ago",
		"minute": " minute ago",
		"hour":   " hour ago",
		"day":    " days ago",
		"month":  " month ago",
		"year":   " year ago",
	},
	true: {
		"second": " second after",
		"minute": " minute after",
		"hour":   " hour after",
		"day":    " days after",
		"month":  " month after",
		"year":   " year after",
	},
}

func getTimeFormal(isFuture bool, timeF string, stamp int64) string {
	return fmt.Sprint(stamp, timeFormal[isFuture][timeF])
}

func DefaultSerializationTime(t time.Time) string {
	timeDifference := int64(time.Now().Second() - t.Second())

	isFuture := false
	if timeDifference < 0 {
		isFuture = true
	}

	if timeDifference <= minute {
		return getTimeFormal(isFuture, "second", timeDifference)
	}

	if timeDifference <= hour {
		return getTimeFormal(isFuture, "minute", timeDifference/minute)
	}

	if timeDifference <= day*2 {
		return getTimeFormal(isFuture, "hour", timeDifference/hour)
	}

	if timeDifference <= month {
		return getTimeFormal(isFuture, "day", timeDifference/day)
	}

	if timeDifference <= year {
		return getTimeFormal(isFuture, "month", timeDifference/month)
	}

	return getTimeFormal(isFuture, "year", timeDifference/year)
}

// EnableTimeEngine
func (t *Table) EnableTimeEngine() *Table {
	t.useTimeEngine = true
	return t
}

// DisableTimeEngine
func (t *Table) DisableTimeEngine() *Table {
	t.useTimeEngine = false
	return t
}

// SetTimeEngine
func (t *Table) SetTimeEngine(in func(time time.Time) string) *Table {
	t.timeEngineFunc = in
	t.useTimeEngine = true
	return t
}
