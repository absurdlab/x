package timez

import "time"

// Seconds is the number of seconds.
type Seconds int64

// Duration returns the time.Duration representation of second.
func (s Seconds) Duration() time.Duration {
	return time.Second * time.Duration(s)
}

// Abs returns the absolute value of Seconds.
func (s Seconds) Abs() Seconds {
	if s < 0 {
		return -s
	}
	return s
}

// ToSeconds converts time.Duration to Seconds.
func ToSeconds(d time.Duration) Seconds {
	return Seconds(d / time.Second)
}
