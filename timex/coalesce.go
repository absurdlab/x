package timex

import "time"

// CoalesceTimestamp returns the first non-zero timestamp from the elements, or
// return a zero-valued UTC timestamp.
func CoalesceTimestamp(elem ...int64) int64 {
	for _, each := range elem {
		if !time.Unix(each, 0).IsZero() {
			return each
		}
	}
	return time.Time{}.UTC().Unix()
}

// Coalesce returns the first non-zero time, or a zero UTC time.
func Coalesce(times ...time.Time) time.Time {
	for _, each := range times {
		if !each.IsZero() {
			return each
		}
	}
	return time.Time{}.UTC()
}
