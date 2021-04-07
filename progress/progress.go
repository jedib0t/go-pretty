package progress

import (
	"io"
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
	autoStop              bool
	done                  chan bool
	lengthTracker         int
	lengthProgress        int
	outputWriter          io.Writer
	hideTime              bool
	hideTracker           bool
	hideValue             bool
	hidePercentage        bool
	messageWidth          int
	numTrackersExpected   int64
	overallTracker        *Tracker
	overallTrackerMutex   sync.RWMutex
	renderInProgress      bool
	renderInProgressMutex sync.RWMutex
	showETA               bool
	showOverallTracker    bool
	sortBy                SortBy
	style                 *Style
	trackerPosition       Position
	trackersActive        []*Tracker
	trackersActiveMutex   sync.RWMutex
	trackersDone          []*Tracker
	trackersDoneMutex     sync.RWMutex
	trackersInQueue       []*Tracker
	trackersInQueueMutex  sync.RWMutex
	updateFrequency       time.Duration
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
func (p *Progress) AppendTracker(t *Tracker) {
	t.start()

	p.overallTrackerMutex.Lock()
	if p.overallTracker == nil {
		p.overallTracker = &Tracker{Total: 1}
		if p.numTrackersExpected > 0 {
			p.overallTracker.Total = p.numTrackersExpected * 100
		}
		p.overallTracker.start()
	}
	p.trackersInQueueMutex.Lock()
	p.trackersInQueue = append(p.trackersInQueue, t)
	p.trackersInQueueMutex.Unlock()
	p.overallTracker.mutex.Lock()
	if p.overallTracker.Total < int64(p.Length())*100 {
		p.overallTracker.Total = int64(p.Length()) * 100
	}
	p.overallTracker.mutex.Unlock()
	p.overallTrackerMutex.Unlock()
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
	p.renderInProgressMutex.RLock()
	defer p.renderInProgressMutex.RUnlock()

	return p.renderInProgress
}

// Length returns the number of Trackers tracked overall.
func (p *Progress) Length() int {
	p.trackersActiveMutex.RLock()
	p.trackersDoneMutex.RLock()
	p.trackersInQueueMutex.RLock()
	out := len(p.trackersInQueue) + len(p.trackersActive) + len(p.trackersDone)
	p.trackersInQueueMutex.RUnlock()
	p.trackersDoneMutex.RUnlock()
	p.trackersActiveMutex.RUnlock()

	return out
}

// LengthActive returns the number of Trackers actively tracked (not done yet).
func (p *Progress) LengthActive() int {
	p.trackersActiveMutex.RLock()
	p.trackersInQueueMutex.RLock()
	out := len(p.trackersInQueue) + len(p.trackersActive)
	p.trackersInQueueMutex.RUnlock()
	p.trackersActiveMutex.RUnlock()

	return out
}

// LengthDone returns the number of Trackers that are done tracking.
func (p *Progress) LengthDone() int {
	p.trackersDoneMutex.RLock()
	out := len(p.trackersDone)
	p.trackersDoneMutex.RUnlock()

	return out
}

// LengthInQueue returns the number of Trackers in queue to be actively tracked
// (not tracking yet).
func (p *Progress) LengthInQueue() int {
	p.trackersInQueueMutex.RLock()
	out := len(p.trackersInQueue)
	p.trackersInQueueMutex.RUnlock()

	return out
}

// SetAutoStop toggles the auto-stop functionality. Auto-stop set to true would
// mean that the Render() function will automatically stop once all currently
// active Trackers reach their final states. When set to false, the client code
// will have to call Progress.Stop() to stop the Render() logic. Default: false.
func (p *Progress) SetAutoStop(autoStop bool) {
	p.autoStop = autoStop
}

// SetMessageWidth sets the (printed) length of the tracker message. Any message
// longer the specified width will be snipped abruptly. Any message shorter than
// the specified width will be padded with spaces.
func (p *Progress) SetMessageWidth(width int) {
	p.messageWidth = width
}

// SetNumTrackersExpected sets the expected number of trackers to be tracked.
// This helps calculate the overall progress with better accuracy.
func (p *Progress) SetNumTrackersExpected(numTrackers int) {
	p.numTrackersExpected = int64(numTrackers)
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

// ShowETA toggles showing the ETA for all individual trackers.
func (p *Progress) ShowETA(show bool) {
	p.showETA = show
}

// ShowPercentage toggles showing the Percent complete for each Tracker.
func (p *Progress) ShowPercentage(show bool) {
	p.hidePercentage = !show
}

// ShowOverallTracker toggles showing the Overall progress tracker with an ETA.
func (p *Progress) ShowOverallTracker(show bool) {
	p.showOverallTracker = show
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
	if p.IsRenderInProgress() {
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

// renderHint has hints for the Render*() logic
type renderHint struct {
	hideTime         bool // hide the time
	hideValue        bool // hide the value
	isOverallTracker bool // is the Overall Progress tracker
}
