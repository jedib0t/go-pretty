package progress

import (
	"math"
	"sync"
	"time"
)

// Tracker helps track the progress of a single task. The way to use it is to
// instantiate a Tracker with a valid Message, a valid (expected) Total, and
// Units values. This should then be fed to the Progress Writer with the
// Writer.AppendTracker() method. When the task that is being done has progress,
// increment the value using the Tracker.Increment(value) method.
type Tracker struct {
	// Message should contain a short description of the "task"; please note
	// that this should NOT be updated in the middle of progress - you should
	// instead use UpdateMessage() to do this safely without hitting any race
	// conditions
	Message string
	// ExpectedDuration tells how long this task is expected to take; and will
	// be used in calculation of the ETA value
	ExpectedDuration time.Duration
	// Total should be set to the (expected) Total/Final value to be reached
	Total int64
	// Units defines the type of the "value" being tracked
	Units Units

	done      bool
	err       bool
	mutex     sync.RWMutex
	timeStart time.Time
	timeStop  time.Time
	value     int64
}

// ETA returns the expected time of "arrival" or completion of this tracker. It
// is an estimate and is not guaranteed.
func (t *Tracker) ETA() time.Duration {
	t.mutex.RLock()
	defer t.mutex.RUnlock()

	timeTaken := time.Since(t.timeStart)
	if t.ExpectedDuration > time.Duration(0) && t.ExpectedDuration > timeTaken {
		return t.ExpectedDuration - timeTaken
	}

	pDone := int64(t.percentDoneWithoutLock())
	if pDone == 0 {
		return time.Duration(0)
	}
	return time.Duration((int64(timeTaken) / pDone) * (100 - pDone))
}

// Increment updates the current value of the task being tracked.
func (t *Tracker) Increment(value int64) {
	t.mutex.Lock()
	t.incrementWithoutLock(value)
	t.mutex.Unlock()
}

// IncrementWithError updates the current value of the task being tracked and
// marks that an error occurred.
func (t *Tracker) IncrementWithError(value int64) {
	t.mutex.Lock()
	t.incrementWithoutLock(value)
	t.err = true
	t.mutex.Unlock()
}

// IsDone returns true if the tracker is done (value has reached the expected
// Total set during initialization).
func (t *Tracker) IsDone() bool {
	t.mutex.RLock()
	defer t.mutex.RUnlock()

	return t.done
}

// IsErrored true if an error was set with IncrementWithError or MarkAsErrored.
func (t *Tracker) IsErrored() bool {
	t.mutex.RLock()
	defer t.mutex.RUnlock()

	return t.err
}

// IsIndeterminate returns true if the tracker is indeterminate; i.e., the total
// is unknown and it is impossible to auto-calculate if tracking is done.
func (t *Tracker) IsIndeterminate() bool {
	t.mutex.RLock()
	defer t.mutex.RUnlock()

	return t.Total == 0
}

// MarkAsDone forces completion of the tracker by updating the current value as
// the expected Total value.
func (t *Tracker) MarkAsDone() {
	t.mutex.Lock()
	t.Total = t.value
	t.stop()
	t.mutex.Unlock()
}

// MarkAsErrored forces completion of the tracker by updating the current value as
// the expected Total value, and recording as error.
func (t *Tracker) MarkAsErrored() {
	t.mutex.Lock()
	// only update error if not done and if not previously set
	if !t.done {
		t.Total = t.value
		t.err = true
		t.stop()
	}
	t.mutex.Unlock()
}

// PercentDone returns the currently completed percentage value.
func (t *Tracker) PercentDone() float64 {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	return t.percentDoneWithoutLock()
}

func (t *Tracker) percentDoneWithoutLock() float64 {
	if t.Total == 0 {
		return 0
	}
	return float64(t.value) * 100.0 / float64(t.Total)
}

// Reset resets the tracker to its initial state.
func (t *Tracker) Reset() {
	t.mutex.Lock()
	t.done = false
	t.err = false
	t.timeStart = time.Time{}
	t.timeStop = time.Time{}
	t.value = 0
	t.mutex.Unlock()
}

// SetValue sets the value of the tracker and re-calculates if the tracker is
// "done".
func (t *Tracker) SetValue(value int64) {
	t.mutex.Lock()
	t.done = false
	t.timeStop = time.Time{}
	t.value = 0
	t.incrementWithoutLock(value)
	t.mutex.Unlock()
}

// UpdateMessage updates the message string.
func (t *Tracker) UpdateMessage(msg string) {
	t.mutex.Lock()
	t.Message = msg
	t.mutex.Unlock()
}

// Value returns the current value of the tracker.
func (t *Tracker) Value() int64 {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	return t.value
}

func (t *Tracker) message() string {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	return t.Message
}

func (t *Tracker) valueAndTotal() (int64, int64) {
	t.mutex.RLock()
	value := t.value
	total := t.Total
	t.mutex.RUnlock()
	return value, total
}

func (t *Tracker) incrementWithoutLock(value int64) {
	if !t.done {
		t.value += value
		if t.Total > 0 && t.value >= t.Total {
			t.stop()
		}
	}
}

func (t *Tracker) start() {
	t.mutex.Lock()
	if t.Total < 0 {
		t.Total = math.MaxInt64
	}
	t.done = false
	t.err = false
	t.timeStart = time.Now()
	t.mutex.Unlock()
}

// this must be called with the mutex held with a write lock
func (t *Tracker) stop() {
	t.done = true
	t.timeStop = time.Now()
	if t.value > t.Total {
		t.Total = t.value
	}
}
