package progress

import (
	"time"
)

// Tracker helps track the progress of a single task. The way to use it is to
// instantiate a Tracker with a valid Message, a valid (expected) Total, and
// Units values. This should then be fed to the Progress Writer with the
// Writer.AppendTracker() method. When the task that is being done has progress,
// increment the value using the Tracker.Increment(value) method.
type Tracker struct {
	// Message should contain a short description of the "task"
	Message string
	// ExpectedDuration tells how long this task is expected to take; and will
	// be used in calculation of the ETA value
	ExpectedDuration time.Duration
	// Total should be set to the (expected) Total/Final value to be reached
	Total int64
	// Units defines the type of the "value" being tracked
	Units Units

	done      bool
	timeStart time.Time
	timeStop  time.Time
	value     int64
}

// ETA returns the expected time of "arrival" or completion of this tracker. It
// is an estimate and is not guaranteed.
func (t *Tracker) ETA() time.Duration {
	timeTaken := time.Since(t.timeStart)
	if t.ExpectedDuration > time.Duration(0) && t.ExpectedDuration > timeTaken {
		return t.ExpectedDuration - timeTaken
	}

	pDone := int64(t.PercentDone())
	if pDone == 0 {
		return time.Duration(0)
	}
	return time.Duration((int64(timeTaken) / pDone) * (100 - pDone))
}

// Increment updates the current value of the task being tracked.
func (t *Tracker) Increment(value int64) {
	if !t.done {
		t.value += value
		if t.Total > 0 && t.value >= t.Total {
			t.stop()
		}
	}
}

// IsDone returns true if the tracker is done (value has reached the expected
// Total set during initialization).
func (t *Tracker) IsDone() bool {
	return t.done
}

// MarkAsDone forces completion of the tracker by updating the current value as
// the expected Total value.
func (t *Tracker) MarkAsDone() {
	t.Total = t.value
	t.stop()
}

// PercentDone returns the currently completed percentage value.
func (t *Tracker) PercentDone() float64 {
	if t.Total == 0 {
		return 0
	}
	return float64(t.value) * 100.0 / float64(t.Total)
}

// Reset resets the tracker to its initial state.
func (t *Tracker) Reset() {
	t.done = false
	t.timeStart = time.Time{}
	t.timeStop = time.Time{}
	t.value = 0
}

// SetValue sets the value of the tracker and re-calculates if the tracker is
// "done".
func (t *Tracker) SetValue(value int64) {
	t.done = false
	t.timeStop = time.Time{}
	t.value = 0
	t.Increment(value)
}

func (t *Tracker) start() {
	t.done = false
	t.timeStart = time.Now()
}

func (t *Tracker) stop() {
	t.done = true
	t.timeStop = time.Now()
	if t.value > t.Total {
		t.Total = t.value
	}
}
