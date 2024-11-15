package times

import "time"

// Seconds returns the duration of s seconds.
func Seconds(s int64) time.Duration {
	return time.Duration(s) * time.Second
}

// Milliseconds returns the duration of ms milliseconds.
func Milliseconds(ms int64) time.Duration {
	return time.Duration(ms) * time.Millisecond
}

// Microseconds returns the duration of us microseconds.
func Microseconds(us int64) time.Duration {
	return time.Duration(us) * time.Microsecond
}

// Minutes returns the duration of m minutes.
func Minutes(m int64) time.Duration {
	return time.Duration(m) * time.Minute
}

// Hours returns the duration of h hours.
func Hours(h int64) time.Duration {
	return time.Duration(h) * time.Hour
}
