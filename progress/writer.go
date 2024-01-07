package progress

import (
	"io"
	"time"
)

// Writer declares the interfaces that can be used to set up and render a
// Progress tracker with one or more trackers.
type Writer interface {
	AppendTracker(tracker *Tracker)
	AppendTrackers(trackers []*Tracker)
	IsRenderInProgress() bool
	Length() int
	LengthActive() int
	LengthDone() int
	LengthInQueue() int
	Log(msg string, a ...interface{})
	SetAutoStop(autoStop bool)
	SetMessageLength(length int)
	SetNumTrackersExpected(numTrackers int)
	SetOutputWriter(output io.Writer)
	SetPinnedMessages(messages ...string)
	SetSortBy(sortBy SortBy)
	SetStyle(style Style)
	SetTerminalWidth(width int)
	SetTrackerLength(length int)
	SetTrackerPosition(position Position)
	SetUpdateFrequency(frequency time.Duration)
	Stop()
	Style() *Style
	Render()

	// Deprecated: in favor of SetMessageLength(length)
	SetMessageWidth(width int)
	// Deprecated: in favor of Style().Visibility.ETA
	ShowETA(show bool)
	// Deprecated: in favor of Style().Visibility.TrackerOverall
	ShowOverallTracker(show bool)
	// Deprecated: in favor of Style().Visibility.Percentage
	ShowPercentage(show bool)
	// Deprecated: in favor of Style().Visibility.Time
	ShowTime(show bool)
	// Deprecated: in favor of Style().Visibility.Tracker
	ShowTracker(show bool)
	// Deprecated: in favor of Style().Visibility.Value
	ShowValue(show bool)
}

// NewWriter initializes and returns a Writer.
func NewWriter() Writer {
	return &Progress{}
}
