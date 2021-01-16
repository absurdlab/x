package timez

import "time"

// Timestamp is the number of seconds since Unix epoch.
type Timestamp int64

// Time returns the time.Time representation of the Timestamp.
func (t Timestamp) Time() time.Time {
	return time.Unix(int64(t), 0)
}

// Ref returns the reference to the Timestamp.
func (t Timestamp) Ref() *Timestamp {
	return &t
}

// Now returns the current Timestamp
func Now() Timestamp {
	return ToTimestamp(time.Now())
}

// ToTimestamp returns the UTC Timestamp.
func ToTimestamp(t time.Time) Timestamp {
	return Timestamp(t.Round(time.Second).Unix())
}
