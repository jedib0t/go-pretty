package progress

import (
	"fmt"
	"io"
	"math"
	"os"
	"sync"
	"time"
	"unicode/utf8"
)

var (
	// DefaultLengthTracker defines a sane value for a Tracker's length.
	DefaultLengthTracker = 20

	// DefaultUpdateFrequency defines a sane value for the frequency with which
	// all the Tracker's get updated on the screen.
	DefaultUpdateFrequency = time.Millisecond * 250
)

// Progress helps track progress for one or more tasks.
type Progress struct {
	autoStop             bool
	done                 chan bool
	lengthTracker        int
	lengthProgress       int
	outputWriter         io.Writer
	hideTime             bool
	hideTracker          bool
	hideValue            bool
	hidePercentage       bool
	renderInProgress     bool
	sortBy               SortBy
	style                *Style
	trackerPosition      Position
	trackersActive       []*Tracker
	trackersDone         []*Tracker
	trackersInQueue      []*Tracker
	trackersInQueueMutex sync.Mutex
	updateFrequency      time.Duration
}

// Position defines the position of the Tracker with respect to the Tracker's
// Message.
type Position int

const (
	// PositionLeft will make the Tracker be displayed first before the Message.
	PositionLeft Position = iota

	// PositionRight will make the Tracker be displayed after the Message.
	PositionRight
)

// AppendTracker appends a single Tracker for tracking. The Tracker gets added
// to a queue, which gets picked up by the Render logic in the next rendering
// cycle.
func (p *Progress) AppendTracker(tracker *Tracker) {
	p.trackersInQueueMutex.Lock()
	if tracker.Total <= 0 {
		tracker.Total = math.MaxInt64
	}
	tracker.start()
	p.trackersInQueue = append(p.trackersInQueue, tracker)
	p.trackersInQueueMutex.Unlock()
}

// AppendTrackers appends one or more Trackers for tracking.
func (p *Progress) AppendTrackers(trackers []*Tracker) {
	for _, tracker := range trackers {
		p.AppendTracker(tracker)
	}
}

// IsRenderInProgress returns true if a call to Render() was made, and is still
// in progress and has not ended yet.
func (p *Progress) IsRenderInProgress() bool {
	return p.renderInProgress
}

// Length returns the number of Trackers tracked overall.
func (p *Progress) Length() int {
	return len(p.trackersInQueue) + len(p.trackersActive) + len(p.trackersDone)
}

// LengthActive returns the number of Trackers actively tracked (not done yet).
func (p *Progress) LengthActive() int {
	return len(p.trackersInQueue) + len(p.trackersActive)
}

// SetAutoStop toggles the auto-stop functionality. Auto-stop set to true would
// mean that the Render() function will automatically stop once all currently
// active Trackers reach their final states. When set to false, the client code
// will have to call Progress.Stop() to stop the Render() logic. Default: false.
func (p *Progress) SetAutoStop(autoStop bool) {
	p.autoStop = autoStop
}

// SetOutputWriter redirects the output of Render to an io.writer object like
// os.Stdout or os.Stderr or a file. Warning: redirecting the output to a file
// may not work well as the Render() logic moves the cursor around a lot.
func (p *Progress) SetOutputWriter(writer io.Writer) {
	p.outputWriter = writer
}

// SetSortBy defines the sorting mechanism to use to sort the Active Trackers
// before rendering the. Default: no-sorting == sort-by-insertion-order.
func (p *Progress) SetSortBy(sortBy SortBy) {
	p.sortBy = sortBy
}

// SetStyle sets the Style to use for rendering.
func (p *Progress) SetStyle(style Style) {
	p.style = &style
}

// SetTrackerLength sets the text-length of all the Trackers.
func (p *Progress) SetTrackerLength(length int) {
	p.lengthTracker = length
}

// SetTrackerPosition sets the position of the tracker with respect to the
// Tracker message text.
func (p *Progress) SetTrackerPosition(position Position) {
	p.trackerPosition = position
}

// SetUpdateFrequency sets the update frequency while rendering the trackers.
// the lower the value, the more number of times the Trackers get refreshed. A
// sane value would be 250ms.
func (p *Progress) SetUpdateFrequency(frequency time.Duration) {
	p.updateFrequency = frequency
}

// ShowPercentage toggles showing the Percent complete for each Tracker.
func (p *Progress) ShowPercentage(show bool) {
	p.hidePercentage = !show
}

// ShowTime toggles showing the Time taken by each Tracker.
func (p *Progress) ShowTime(show bool) {
	p.hideTime = !show
}

// ShowTracker toggles showing the Tracker (the progress bar).
func (p *Progress) ShowTracker(show bool) {
	p.hideTracker = !show
}

// ShowValue toggles showing the actual Value of the Tracker.
func (p *Progress) ShowValue(show bool) {
	p.hideValue = !show
}

// Stop stops the Render() logic that is in progress.
func (p *Progress) Stop() {
	if p.renderInProgress {
		p.done <- true
	}
}

// Style returns the current Style.
func (p *Progress) Style() *Style {
	if p.style == nil {
		tempStyle := StyleDefault
		p.style = &tempStyle
	}
	return p.style
}

func (p *Progress) initForRender() {
	// pick a default style
	p.Style()

	// reset the signals
	p.done = make(chan bool, 1)

	// pick default lengths if no valid ones set
	if p.lengthTracker <= 0 {
		p.lengthTracker = DefaultLengthTracker
	}

	// calculate length of the actual progress bar by discount the left/right
	// border/box chars
	p.lengthProgress = p.lengthTracker -
		utf8.RuneCountInString(p.style.Chars.BoxLeft) -
		utf8.RuneCountInString(p.style.Chars.BoxRight)

	// if not output write has been set, output to STDOUT
	if p.outputWriter == nil {
		p.outputWriter = os.Stdout
	}

	// pick a sane update frequency if none set
	if p.updateFrequency <= 0 {
		p.updateFrequency = DefaultUpdateFrequency
	}
}

func (p *Progress) write(a ...interface{}) {
	p.outputWriter.Write([]byte(fmt.Sprint(a...)))
}
