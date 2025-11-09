package progress

import (
	"context"
	"math"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestProgress_AppendTracker(t *testing.T) {
	p := Progress{}
	assert.Equal(t, 0, len(p.trackersInQueue))

	tracker := &Tracker{}
	assert.Equal(t, int64(0), tracker.Total)
	p.AppendTracker(tracker)
	assert.Equal(t, 1, len(p.trackersInQueue))
	assert.Equal(t, int64(0), tracker.Total)

	tracker2 := &Tracker{Total: -1}
	assert.Equal(t, int64(-1), tracker2.Total)
	p.AppendTracker(tracker2)
	assert.Equal(t, 2, len(p.trackersInQueue))
	assert.Equal(t, int64(math.MaxInt64), tracker2.Total)
}

func TestProgress_AppendTrackers(t *testing.T) {
	p := Progress{}
	assert.Equal(t, 0, len(p.trackersInQueue))

	p.AppendTrackers([]*Tracker{{}, {}})
	assert.Equal(t, 2, len(p.trackersInQueue))
}

func TestProgress_IsRenderInProgress(t *testing.T) {
	p := Progress{}
	assert.False(t, p.IsRenderInProgress())

	p.renderInProgress = true
	assert.True(t, p.IsRenderInProgress())
}

func TestProgress_Length(t *testing.T) {
	p := Progress{}
	assert.Equal(t, 0, p.Length())

	p.trackersActive = append(p.trackersActive, &Tracker{})
	assert.Equal(t, 1, p.Length())
	p.trackersInQueue = append(p.trackersInQueue, &Tracker{})
	assert.Equal(t, 2, p.Length())
	p.trackersDone = append(p.trackersDone, &Tracker{})
	assert.Equal(t, 3, p.Length())
}

func TestProgress_LengthActive(t *testing.T) {
	p := Progress{}
	assert.Equal(t, 0, p.Length())
	assert.Equal(t, 0, p.LengthActive())

	p.trackersActive = append(p.trackersActive, &Tracker{})
	assert.Equal(t, 1, p.Length())
	assert.Equal(t, 1, p.LengthActive())
	p.trackersInQueue = append(p.trackersInQueue, &Tracker{})
	assert.Equal(t, 2, p.Length())
	assert.Equal(t, 2, p.LengthActive())
}

func TestProgress_LengthDone(t *testing.T) {
	p := Progress{}
	assert.Equal(t, 0, p.Length())
	assert.Equal(t, 0, p.LengthDone())

	p.trackersDone = append(p.trackersDone, &Tracker{})
	assert.Equal(t, 1, p.Length())
	assert.Equal(t, 1, p.LengthDone())
}

func TestProgress_LengthInQueue(t *testing.T) {
	p := Progress{}
	assert.Equal(t, 0, p.Length())
	assert.Equal(t, 0, p.LengthInQueue())

	p.trackersInQueue = append(p.trackersInQueue, &Tracker{})
	assert.Equal(t, 1, p.Length())
	assert.Equal(t, 1, p.LengthInQueue())
}

func TestProgress_Log(t *testing.T) {
	p := Progress{}
	assert.Len(t, p.logsToRender, 0)

	p.Log("testing log")
	assert.Len(t, p.logsToRender, 1)
}

func TestProgress_SetAutoStop(t *testing.T) {
	p := Progress{}
	assert.False(t, p.autoStop)

	p.SetAutoStop(true)
	assert.True(t, p.autoStop)
}

func TestProgress_SetNumTrackersExpected(t *testing.T) {
	p := Progress{}
	assert.Equal(t, int64(0), p.numTrackersExpected)

	p.SetNumTrackersExpected(5)
	assert.Equal(t, int64(5), p.numTrackersExpected)
}

func TestProgress_SetOutputWriter(t *testing.T) {
	p := Progress{}
	assert.Nil(t, p.outputWriter)

	p.SetOutputWriter(os.Stdout)
	assert.Equal(t, os.Stdout, p.outputWriter)
}

func TestProgress_SetPinnedMessages(t *testing.T) {
	p := Progress{}
	assert.Nil(t, p.pinnedMessages)

	p.SetPinnedMessages("pin1", "pin2")
	assert.Equal(t, []string{"pin1", "pin2"}, p.pinnedMessages)
}

func TestProgress_SetSortBy(t *testing.T) {
	p := Progress{}
	assert.Zero(t, p.sortBy)

	p.SetSortBy(SortByMessage)
	assert.Equal(t, SortByMessage, p.sortBy)
}

func TestProgress_SetStyle(t *testing.T) {
	p := Progress{}
	assert.Nil(t, p.style)

	p.SetStyle(StyleCircle)
	assert.Equal(t, StyleCircle.Name, p.Style().Name)
}

func TestProgress_SetMessageLength(t *testing.T) {
	p := Progress{}
	assert.Equal(t, 0, p.lengthMessage)

	p.SetMessageLength(80)
	assert.Equal(t, 80, p.lengthMessage)
	p.SetMessageWidth(81)
	assert.Equal(t, 81, p.lengthMessage)
}

func TestProgress_SetTrackerLength(t *testing.T) {
	p := Progress{}
	assert.Equal(t, 0, p.lengthTracker)

	p.initForRender()
	assert.Equal(t, DefaultLengthTracker, p.lengthTracker)

	p.SetTrackerLength(80)
	assert.Equal(t, 80, p.lengthTracker)
}

func TestProgress_SetTrackerPosition(t *testing.T) {
	p := Progress{}
	assert.Equal(t, PositionLeft, p.trackerPosition)

	p.SetTrackerPosition(PositionRight)
	assert.Equal(t, PositionRight, p.trackerPosition)
}

func TestProgress_SetUpdateFrequency(t *testing.T) {
	p := Progress{}
	assert.Equal(t, time.Duration(0), p.updateFrequency)

	p.initForRender()
	assert.Equal(t, DefaultUpdateFrequency, p.updateFrequency)

	p.SetUpdateFrequency(time.Second)
	assert.Equal(t, time.Second, p.updateFrequency)
}

func TestProgress_ShowETA(t *testing.T) {
	p := Progress{}
	assert.False(t, p.Style().Visibility.ETA)

	p.ShowETA(true)
	assert.True(t, p.Style().Visibility.ETA)
}

func TestProgress_ShowOverallTracker(t *testing.T) {
	p := Progress{}
	assert.False(t, p.Style().Visibility.TrackerOverall)

	p.ShowOverallTracker(true)
	assert.True(t, p.Style().Visibility.TrackerOverall)
}

func TestProgress_ShowPercentage(t *testing.T) {
	p := Progress{}
	assert.True(t, p.Style().Visibility.Percentage)

	p.ShowPercentage(false)
	assert.False(t, p.Style().Visibility.Percentage)
}

func TestProgress_ShowTime(t *testing.T) {
	p := Progress{}
	assert.True(t, p.Style().Visibility.Time)

	p.ShowTime(false)
	assert.False(t, p.Style().Visibility.Time)
}

func TestProgress_ShowTracker(t *testing.T) {
	p := Progress{}
	assert.True(t, p.Style().Visibility.Tracker)

	p.ShowTracker(false)
	assert.False(t, p.Style().Visibility.Tracker)
}

func TestProgress_ShowValue(t *testing.T) {
	p := Progress{}
	assert.True(t, p.Style().Visibility.Value)

	p.ShowValue(false)
	assert.False(t, p.Style().Visibility.Value)
}

func TestProgress_Stop(t *testing.T) {
	p := Progress{}
	p.renderContext, p.renderContextCancel = context.WithCancel(context.Background())
	p.renderInProgress = true
	p.Stop()
	assert.NotNil(t, <-p.renderContext.Done())
}

func TestProgress_Style(t *testing.T) {
	p := Progress{}
	assert.Nil(t, p.style)

	assert.NotNil(t, p.Style())
	assert.Equal(t, StyleDefault.Name, p.Style().Name)
}

func TestProgress_OverallTrackerDisappearsCase(t *testing.T) {
	p := &Progress{}
	p.overallTracker = &Tracker{Total: 1, done: true}
	p.AppendTracker(&Tracker{Total: 1})

	assert.Equal(t, false, p.overallTracker.IsDone())
}

func TestProgress_watchTerminalSize(t *testing.T) {
	p := &Progress{}
	// Set up a cancellable context for watchTerminalSize
	p.renderContext, p.renderContextCancel = context.WithCancel(context.Background())

	// Call watchTerminalSize in a goroutine
	done := make(chan bool)
	go func() {
		p.watchTerminalSize()
		done <- true
	}()

	// Wait a bit to let the ticker fire at least once (covers case <-ticker.C)
	time.Sleep(150 * time.Millisecond)
	p.renderContextCancel()

	// Wait for the goroutine to exit (with timeout)
	select {
	case <-done:
		// Success: goroutine exited after context cancellation
	case <-time.After(1 * time.Second):
		t.Fatal("watchTerminalSize goroutine did not exit after context cancellation")
	}

	// Verify the context was cancelled
	select {
	case <-p.renderContext.Done():
		// Context is cancelled, which is expected
	default:
		t.Fatal("Expected context to be cancelled")
	}
}
