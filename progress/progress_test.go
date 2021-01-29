package progress

import (
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

	p.SetUpdateFrequency(time.Duration(time.Second))
	assert.Equal(t, time.Duration(time.Second), p.updateFrequency)
}

func TestProgress_ShowOverallTracker(t *testing.T) {
	p := Progress{}
	assert.False(t, p.showOverallTracker)

	p.ShowOverallTracker(true)
	assert.True(t, p.showOverallTracker)
}

func TestProgress_ShowPercentage(t *testing.T) {
	p := Progress{}
	assert.False(t, p.hidePercentage)

	p.ShowPercentage(false)
	assert.True(t, p.hidePercentage)
}

func TestProgress_ShowTime(t *testing.T) {
	p := Progress{}
	assert.False(t, p.hideTime)

	p.ShowTime(false)
	assert.True(t, p.hideTime)
}

func TestProgress_ShowTracker(t *testing.T) {
	p := Progress{}
	assert.False(t, p.hideTracker)

	p.ShowTracker(false)
	assert.True(t, p.hideTracker)
}

func TestProgress_ShowValue(t *testing.T) {
	p := Progress{}
	assert.False(t, p.hideValue)

	p.ShowValue(false)
	assert.True(t, p.hideValue)
}

func TestProgress_Stop(t *testing.T) {
	doneChannel := make(chan bool, 1)

	p := Progress{}
	p.done = doneChannel
	p.renderInProgress = true
	p.Stop()
	assert.True(t, <-doneChannel)
}

func TestProgress_Style(t *testing.T) {
	p := Progress{}
	assert.Nil(t, p.style)

	assert.NotNil(t, p.Style())
	assert.Equal(t, StyleDefault.Name, p.Style().Name)
}
