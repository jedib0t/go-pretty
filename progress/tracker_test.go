package progress

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTracker_ETA(t *testing.T) {
	tracker := Tracker{Total: 100}
	assert.Equal(t, time.Duration(0), tracker.ETA())

	tracker.timeStart = time.Now()
	time.Sleep(time.Millisecond * 100)
	tracker.value = 50
	eta := tracker.ETA()
	assert.NotEqual(t, time.Duration(0), eta)
	assert.True(t, eta < time.Second)
}

func TestTracker_Increment(t *testing.T) {
	tracker := Tracker{Total: 100}
	assert.Equal(t, int64(0), tracker.value)
	assert.Equal(t, int64(100), tracker.Total)

	tracker.Increment(10)
	assert.Equal(t, int64(10), tracker.value)
	assert.Equal(t, int64(100), tracker.Total)

	tracker.Increment(100)
	assert.Equal(t, int64(110), tracker.value)
	assert.Equal(t, int64(110), tracker.Total)
	assert.False(t, tracker.timeStop.IsZero())
	assert.True(t, tracker.IsDone())
}

func TestTracker_IsDone(t *testing.T) {
	tracker := Tracker{Total: 10}
	assert.False(t, tracker.IsDone())

	tracker.Increment(10)
	assert.True(t, tracker.IsDone())
}

func TestTracker_MarkAsDone(t *testing.T) {
	tracker := Tracker{}
	assert.False(t, tracker.IsDone())
	assert.True(t, tracker.timeStop.IsZero())

	tracker.MarkAsDone()
	assert.True(t, tracker.IsDone())
	assert.False(t, tracker.timeStop.IsZero())
}

func TestTracker_PercentDone(t *testing.T) {
	tracker := Tracker{Total: 100}
	assert.Equal(t, 0.00, tracker.PercentDone())

	for idx := 1; idx <= 100; idx++ {
		tracker.Increment(1)
		assert.Equal(t, float64(idx), tracker.PercentDone())
	}
}

func TestTracker_Reset(t *testing.T) {
	tracker := Tracker{Total: 100}
	assert.False(t, tracker.done)
	assert.Equal(t, time.Time{}, tracker.timeStart)
	assert.Equal(t, time.Time{}, tracker.timeStop)
	assert.Equal(t, int64(0), tracker.value)

	tracker.start()
	tracker.Increment(tracker.Total)
	tracker.stop()
	assert.True(t, tracker.done)
	assert.NotEqual(t, time.Time{}, tracker.timeStart)
	assert.NotEqual(t, time.Time{}, tracker.timeStop)
	assert.Equal(t, tracker.Total, tracker.value)

	tracker.Reset()
	assert.False(t, tracker.done)
	assert.Equal(t, time.Time{}, tracker.timeStart)
	assert.Equal(t, time.Time{}, tracker.timeStop)
	assert.Equal(t, int64(0), tracker.value)
}
