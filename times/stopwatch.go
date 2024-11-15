package times

import (
	"time"
)

// Stopwatch holds start time and elapsed time to measure the duration of events
type Stopwatch struct {
	start   time.Time
	elapsed time.Duration
	running bool
}

// NewStartedStopwatch creates a new stopwatch and starts it immediately
func NewStartedStopwatch() *Stopwatch {
	sw := &Stopwatch{}
	sw.Start()
	return sw
}

// Start initializes or resumes the stopwatch.
// Returns true if the stopwatch was started successfully, false otherwise
func (sw *Stopwatch) Start() bool {
	if !sw.running {
		sw.start = time.Now()
		sw.running = true
		return true
	} else {
		return false
	}
}

// Stop stops the stopwatch and records the elapsed time.
// Returns true if the stopwatch was stopped successfully, false otherwise
func (sw *Stopwatch) Stop() bool {
	if sw.running {
		sw.elapsed += time.Since(sw.start)
		sw.running = false
		return true
	} else {
		return false
	}
}

// Elapsed returns the total duration the stopwatch has been running
func (sw *Stopwatch) Elapsed() time.Duration {
	if sw.running {
		return sw.elapsed + time.Since(sw.start)
	}
	return sw.elapsed
}

// ElapsedMillis returns the total duration the stopwatch has been running in milliseconds
func (sw *Stopwatch) ElapsedMillis() int64 {
	return sw.Elapsed().Milliseconds()
}

// IsRunning returns true if the stopwatch is currently running
func (sw *Stopwatch) IsRunning() bool {
	return sw.running
}

// Reset resets the stopwatch to its initial state
func (sw *Stopwatch) Reset() {
	sw.start = time.Time{}
	sw.elapsed = 0
	sw.running = false
}
