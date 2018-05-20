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
	SetAutoStop(autoStop bool)
	SetOutputWriter(output io.Writer)
	SetSortBy(sortBy SortBy)
	SetStyle(style Style)
	SetTrackerLength(length int)
	SetTrackerPosition(position Position)
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
