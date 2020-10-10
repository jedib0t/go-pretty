package progress

import (
	"io"
	"time"
)

// Writer declares the interfaces that can be used to setup and render a
// Progress tracker with one or more trackers.
type Writer interface {
	AppendTracker(tracker *Tracker)
	AppendTrackers(trackers []*Tracker)
	IsRenderInProgress() bool
	Length() int
	LengthActive() int
	LengthDone() int
	LengthInQueue() int
	SetAutoStop(autoStop bool)
	SetMessageWidth(width int)
	SetNumTrackersExpected(numTrackers int)
	SetOutputWriter(output io.Writer)
	SetSortBy(sortBy SortBy)
	SetStyle(style Style)
	SetTrackerLength(length int)
	SetTrackerPosition(position Position)
	ShowETA(show bool)
	ShowOverallTracker(show bool)
	ShowPercentage(show bool)
	ShowTime(show bool)
	ShowTracker(show bool)
	ShowValue(show bool)
	SetUpdateFrequency(frequency time.Duration)
	Stop()
	Style() *Style
	Render()
}

// NewWriter initializes and returns a Writer.
func NewWriter() Writer {
	return &Progress{}
}
