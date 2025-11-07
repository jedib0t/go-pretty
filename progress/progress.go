package progress

import (
	"context"
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"github.com/jedib0t/go-pretty/v6/text"
	"golang.org/x/term"
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
	autoStop                 bool
	lengthMessage            int
	lengthProgress           int
	lengthProgressOverall    int
	lengthTracker            int
	logsToRender             []string
	logsToRenderMutex        sync.RWMutex
	numTrackersExpected      int64
	outputWriter             io.Writer
	overallTracker           *Tracker
	overallTrackerMutex      sync.RWMutex
	pinnedMessages           []string
	pinnedMessageMutex       sync.RWMutex
	pinnedMessageNumLines    int
	renderContext            context.Context
	renderContextCancel      context.CancelFunc
	renderContextCancelMutex sync.Mutex
	renderInProgress         bool
	renderInProgressMutex    sync.RWMutex
	sortBy                   SortBy
	style                    *Style
	terminalWidth            int
	terminalWidthMutex       sync.RWMutex
	terminalWidthOverride    int
	trackerPosition          Position
	trackersActive           []*Tracker
	trackersActiveMutex      sync.RWMutex
	trackersDone             []*Tracker
	trackersDoneMutex        sync.RWMutex
	trackersInQueue          []*Tracker
	trackersInQueueMutex     sync.RWMutex
	updateFrequency          time.Duration
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
	if !t.DeferStart {
		t.start()
	}
	p.overallTrackerMutex.Lock()
	defer p.overallTrackerMutex.Unlock()

	if p.overallTracker == nil {
		p.overallTracker = &Tracker{Total: 1}
		if p.numTrackersExpected > 0 {
			p.overallTracker.Total = p.numTrackersExpected * 100
		}
		p.overallTracker.start()
	}

	// append the tracker to the "in-queue" list
	p.trackersInQueueMutex.Lock()
	p.trackersInQueue = append(p.trackersInQueue, t)
	p.trackersInQueueMutex.Unlock()

	// update the expected total progress since we are appending a new tracker
	p.overallTracker.UpdateTotal(int64(p.Length()) * 100)
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

// Log appends a log to display above the active progress bars during the next
// refresh.
func (p *Progress) Log(msg string, a ...interface{}) {
	if len(a) > 0 {
		msg = fmt.Sprintf(msg, a...)
	}
	p.logsToRenderMutex.Lock()
	p.logsToRender = append(p.logsToRender, msg)
	p.logsToRenderMutex.Unlock()
}

// SetAutoStop toggles the auto-stop functionality. Auto-stop set to true would
// mean that the Render() function will automatically stop once all currently
// active Trackers reach their final states. When set to false, the client code
// will have to call Progress.Stop() to stop the Render() logic. Default: false.
func (p *Progress) SetAutoStop(autoStop bool) {
	p.autoStop = autoStop
}

// SetMessageLength sets the (printed) length of the tracker message. Any
// message longer the specified length will be snipped. Any message shorter than
// the specified width will be padded with spaces.
func (p *Progress) SetMessageLength(length int) {
	p.lengthMessage = length
}

// SetMessageWidth sets the (printed) length of the tracker message. Any message
// longer the specified width will be snipped. Any message shorter than the
// specified width will be padded with spaces.
// Deprecated: in favor of SetMessageLength(length)
func (p *Progress) SetMessageWidth(width int) {
	p.lengthMessage = width
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

// SetPinnedMessages sets message(s) pinned above all the trackers of the
// progress bar. This method can be used to overwrite all the pinned messages.
// Call this function without arguments to "clear" the pinned messages.
func (p *Progress) SetPinnedMessages(messages ...string) {
	p.pinnedMessageMutex.Lock()
	defer p.pinnedMessageMutex.Unlock()

	p.pinnedMessages = messages
}

// SetSortBy defines the sorting mechanism to use to sort the Active Trackers
// before rendering. Default: no-sorting == sort-by-insertion-order.
func (p *Progress) SetSortBy(sortBy SortBy) {
	p.sortBy = sortBy
}

// SetStyle sets the Style to use for rendering.
func (p *Progress) SetStyle(style Style) {
	p.style = &style
}

// SetTerminalWidth sets up a sticky terminal width and prevents the Progress
// Writer from polling for the real width during render.
func (p *Progress) SetTerminalWidth(width int) {
	p.terminalWidthOverride = width
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
// the lower the value, the more frequently the Trackers get refreshed. A
// sane value would be 250ms.
func (p *Progress) SetUpdateFrequency(frequency time.Duration) {
	p.updateFrequency = frequency
}

// ShowETA toggles showing the ETA for all individual trackers.
// Deprecated: in favor of Style().Visibility.ETA
func (p *Progress) ShowETA(show bool) {
	p.Style().Visibility.ETA = show
}

// ShowPercentage toggles showing the Percent complete for each Tracker.
// Deprecated: in favor of Style().Visibility.Percentage
func (p *Progress) ShowPercentage(show bool) {
	p.Style().Visibility.Percentage = show
}

// ShowOverallTracker toggles showing the Overall progress tracker with an ETA.
// Deprecated: in favor of Style().Visibility.TrackerOverall
func (p *Progress) ShowOverallTracker(show bool) {
	p.Style().Visibility.TrackerOverall = show
}

// ShowTime toggles showing the Time taken by each Tracker.
// Deprecated: in favor of Style().Visibility.Time
func (p *Progress) ShowTime(show bool) {
	p.Style().Visibility.Time = show
}

// ShowTracker toggles showing the Tracker (the progress bar).
// Deprecated: in favor of Style().Visibility.Tracker
func (p *Progress) ShowTracker(show bool) {
	p.Style().Visibility.Tracker = show
}

// ShowValue toggles showing the actual Value of the Tracker.
// Deprecated: in favor of Style().Visibility.Value
func (p *Progress) ShowValue(show bool) {
	p.Style().Visibility.Value = show
}

// Stop stops the Render() logic that is in progress.
func (p *Progress) Stop() {
	p.renderContextCancelMutex.Lock()
	defer p.renderContextCancelMutex.Unlock()

	if p.renderContextCancel != nil {
		p.renderContextCancel()
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

func (p *Progress) getTerminalWidth() int {
	p.terminalWidthMutex.RLock()
	defer p.terminalWidthMutex.RUnlock()

	if p.terminalWidthOverride > 0 {
		return p.terminalWidthOverride
	}
	return p.terminalWidth
}

func (p *Progress) initForRender() {
	// reset the signals
	p.renderContextCancelMutex.Lock()
	p.renderContext, p.renderContextCancel = context.WithCancel(context.Background())
	p.renderContextCancelMutex.Unlock()

	// pick a default style
	p.Style()
	if p.style.Options.SpeedOverallFormatter == nil {
		p.style.Options.SpeedOverallFormatter = FormatNumber
	}

	// pick default lengths if no valid ones set
	if p.lengthTracker <= 0 {
		p.lengthTracker = DefaultLengthTracker
	}

	// calculate length of the actual progress bar by discounting the left/right
	// border/box chars
	p.lengthProgress = p.lengthTracker -
		text.StringWidthWithoutEscSequences(p.style.Chars.BoxLeft) -
		text.StringWidthWithoutEscSequences(p.style.Chars.BoxRight)
	p.lengthProgressOverall = p.lengthMessage +
		text.StringWidthWithoutEscSequences(p.style.Options.Separator) +
		p.lengthProgress + 1
	if p.style.Visibility.Percentage {
		p.lengthProgressOverall += text.StringWidthWithoutEscSequences(
			fmt.Sprintf(p.style.Options.PercentFormat, 0.0),
		)
	}

	// if not output write has been set, output to STDOUT
	if p.outputWriter == nil {
		p.outputWriter = os.Stdout
	}

	// pick a sane update frequency if none set
	if p.updateFrequency <= 0 {
		p.updateFrequency = DefaultUpdateFrequency
	}

	if p.outputWriter == os.Stdout {
		// get the current terminal size for preventing roll-overs, and do this in a
		// background loop until end of render. This only works if the output writer is STDOUT.
		go p.watchTerminalSize() // needs p.updateFrequency
	}
}

func (p *Progress) updateTerminalSize() {
	p.terminalWidthMutex.Lock()
	defer p.terminalWidthMutex.Unlock()

	p.terminalWidth, _, _ = term.GetSize(int(os.Stdout.Fd()))
}

func (p *Progress) watchTerminalSize() {
	// once
	p.updateTerminalSize()
	// until end of time
	ticker := time.NewTicker(time.Second / 10)
	for {
		select {
		case <-ticker.C:
			p.updateTerminalSize()
		case <-p.renderContext.Done():
			return
		}
	}
}

// renderHint has hints for the Render*() logic
type renderHint struct {
	hideTime         bool // hide the time
	hideValue        bool // hide the value
	isOverallTracker bool // is the Overall Progress tracker
	terminalWidth    int  // cached terminal width for this render cycle
}
