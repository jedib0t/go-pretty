package progress

import (
	"math"
	"strings"
	"time"

	"github.com/jedib0t/go-pretty/util"
)

// Render renders the Progress tracker and handles all existing trackers and
// those that are added dynamically while render is in progress.
func (p *Progress) Render() {
	if !p.renderInProgress {
		p.initForRender()

		c := time.Tick(p.updateFrequency)
		for p.renderInProgress = true; p.renderInProgress; {
			select {
			case <-c:
				if len(p.trackersInQueue) > 0 || len(p.trackersActive) > 0 {
					p.renderTrackers()
				}
			case <-p.done:
				p.renderInProgress = false
			}
		}
	}
}

func (p *Progress) renderTrackers() {
	// move up N times based on the number of active trackers
	if len(p.trackersActive) > 0 {
		p.write(util.CursorUp.Sprintn(len(p.trackersActive)))
	}

	// move trackers waiting in queue to the active list
	if len(p.trackersInQueue) > 0 {
		p.trackersInQueueMutex.Lock()
		p.trackersActive = append(p.trackersActive, p.trackersInQueue...)
		p.trackersInQueue = []*Tracker{}
		p.trackersInQueueMutex.Unlock()
	}

	// render the finished trackers and move them to the "done" list
	for idx, tracker := range p.trackersActive {
		if tracker.IsDone() {
			p.renderTracker(tracker)
			if idx < len(p.trackersActive) {
				p.trackersActive = append(p.trackersActive[:idx], p.trackersActive[idx+1:]...)
			}
			p.trackersDone = append(p.trackersDone, tracker)
		}
	}

	// sort and render the active trackers
	p.sortBy.Sort(p.trackersActive)
	for _, tracker := range p.trackersActive {
		p.renderTracker(tracker)
	}

	// stop if auto stop is enabled and there are no more active trackers
	if p.autoStop && len(p.trackersInQueue) == 0 && len(p.trackersActive) == 0 {
		p.done <- true
	}
}

func (p *Progress) renderTracker(t *Tracker) {
	p.write(util.EraseLine.Sprint())

	pDotValue := float64(t.Total) / float64(p.lengthProgress)
	pFinishedDots := float64(t.value) / pDotValue
	pFinishedLen := int(math.Ceil(pFinishedDots))
	pUnfinishedLen := p.lengthProgress - pFinishedLen

	var pFinished, pInProgress, pUnfinished string
	if pFinishedLen > 0 {
		pFinished = strings.Repeat(p.style.Chars.Finished, pFinishedLen-1)
	}
	if pUnfinishedLen > 0 {
		pUnfinished = strings.Repeat(p.style.Chars.Unfinished, pUnfinishedLen)
	}

	pFinishedDecimals := pFinishedDots - float64(int(pFinishedDots))
	if pFinishedDecimals > 0.75 {
		pInProgress = p.style.Chars.Finished75
	} else if pFinishedDecimals > 0.50 {
		pInProgress = p.style.Chars.Finished50
	} else if pFinishedDecimals > 0.25 {
		pInProgress = p.style.Chars.Finished25
	} else {
		pInProgress = p.style.Chars.Unfinished
	}

	if t.IsDone() {
		p.renderTrackerDone(t)
	} else {
		p.renderTrackerProgress(t, p.style.Colors.Tracker.Sprintf("%s%s%s%s%s",
			p.style.Chars.BoxLeft, pFinished, pInProgress, pUnfinished, p.style.Chars.BoxRight,
		))
	}
}

func (p *Progress) renderTrackerDone(t *Tracker) {
	p.write(p.style.Colors.Message.Sprint(t.Message))
	p.write(" " + p.style.Options.MessageTrackerSeparator + " ")
	p.write(p.style.Colors.Done.Sprint(p.style.Options.DoneString))
	p.renderTrackerValueAndTime(t)
	p.write("\n")
}

func (p *Progress) renderTrackerProgress(t *Tracker, trackerStr string) {
	if p.trackerPosition == PositionRight {
		p.write(p.style.Colors.Message.Sprint(t.Message))
		p.write(" " + p.style.Options.MessageTrackerSeparator + " ")
		p.renderTrackerPercentage(t)
		if !p.hideTracker {
			p.write(" " + p.style.Colors.Tracker.Sprint(trackerStr))
		}
		p.renderTrackerValueAndTime(t)
		p.write("\n")
	} else {
		p.renderTrackerPercentage(t)
		if !p.hideTracker {
			p.write(" " + p.style.Colors.Tracker.Sprint(trackerStr))
		}
		p.renderTrackerValueAndTime(t)
		p.write(" " + p.style.Options.MessageTrackerSeparator + " ")
		p.write(p.style.Colors.Message.Sprint(t.Message))
		p.write("\n")
	}
}

func (p *Progress) renderTrackerPercentage(t *Tracker) {
	if !p.hidePercentage {
		p.write(p.style.Colors.Percent.Sprintf(p.style.Options.PercentFormat, t.PercentDone()))
	}
}

func (p *Progress) renderTrackerValueAndTime(t *Tracker) {
	if !p.hideValue || !p.hideTime {
		var out strings.Builder
		out.WriteString(" [")
		if !p.hideValue {
			out.WriteString(p.style.Colors.Value.Sprint(t.Units.Sprint(t.value)))
		}
		if !p.hideValue && !p.hideTime {
			out.WriteString(" ")
		}
		if !p.hideTime {
			out.WriteString("in ")
			if t.IsDone() {
				out.WriteString(p.style.Colors.Time.Sprint(
					t.timeStop.Sub(t.timeStart).Round(p.style.Options.TimeDonePrecision)))
			} else {
				out.WriteString(p.style.Colors.Time.Sprint(
					time.Since(t.timeStart).Round(p.style.Options.TimeInProgressPrecision)))
			}
		}
		out.WriteString("]")

		p.write(p.style.Colors.Stats.Sprint(out.String()))
	}
}
