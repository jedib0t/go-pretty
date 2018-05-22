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
		lastRenderLength := 0
		for p.renderInProgress = true; p.renderInProgress; {
			select {
			case <-c:
				if len(p.trackersInQueue) > 0 || len(p.trackersActive) > 0 {
					lastRenderLength = p.renderTrackers(lastRenderLength)
				}
			case <-p.done:
				p.renderInProgress = false
			}
		}
	}
}

func (p *Progress) renderTrackers(lastRenderLength int) int {
	// buffer all output into a strings.Builder object
	var out strings.Builder
	out.Grow(lastRenderLength)

	// move up N times based on the number of active trackers
	if len(p.trackersActive) > 0 {
		out.WriteString(util.CursorUp.Sprintn(len(p.trackersActive)))
	}

	// move trackers waiting in queue to the active list
	if len(p.trackersInQueue) > 0 {
		p.trackersInQueueMutex.Lock()
		p.trackersActive = append(p.trackersActive, p.trackersInQueue...)
		p.trackersInQueue = make([]*Tracker, 0)
		p.trackersInQueueMutex.Unlock()
	}

	// render the finished trackers and move them to the "done" list
	for idx, tracker := range p.trackersActive {
		if tracker.IsDone() {
			p.renderTracker(&out, tracker)
			if idx < len(p.trackersActive) {
				p.trackersActive = append(p.trackersActive[:idx], p.trackersActive[idx+1:]...)
			}
			p.trackersDone = append(p.trackersDone, tracker)
		}
	}

	// sort and render the active trackers
	p.sortBy.Sort(p.trackersActive)
	for _, tracker := range p.trackersActive {
		p.renderTracker(&out, tracker)
	}

	// write the text to the output writer
	p.outputWriter.Write([]byte(out.String()))

	// stop if auto stop is enabled and there are no more active trackers
	if p.autoStop && len(p.trackersInQueue) == 0 && len(p.trackersActive) == 0 {
		p.done <- true
	}

	return out.Len()
}

func (p *Progress) renderTracker(out *strings.Builder, t *Tracker) {
	out.WriteString(util.EraseLine.Sprint())

	if t.IsDone() {
		p.renderTrackerDone(out, t)
	} else {
		pDotValue := float64(t.Total) / float64(p.lengthProgress)
		pFinishedDots := float64(t.value) / pDotValue
		pFinishedDotsFraction := pFinishedDots - float64(int(pFinishedDots))
		pFinishedLen := int(math.Ceil(pFinishedDots))
		pUnfinishedLen := p.lengthProgress - pFinishedLen

		var pFinished, pUnfinished string
		if pFinishedLen > 0 {
			pFinished = strings.Repeat(p.style.Chars.Finished, pFinishedLen-1)
		}
		if pUnfinishedLen > 0 {
			pUnfinished = strings.Repeat(p.style.Chars.Unfinished, pUnfinishedLen)
		}

		pInProgress := p.style.Chars.Unfinished
		if pFinishedDotsFraction > 0.75 {
			pInProgress = p.style.Chars.Finished75
		} else if pFinishedDotsFraction > 0.50 {
			pInProgress = p.style.Chars.Finished50
		} else if pFinishedDotsFraction > 0.25 {
			pInProgress = p.style.Chars.Finished25
		}

		p.renderTrackerProgress(out, t, p.style.Colors.Tracker.Sprintf("%s%s%s%s%s",
			p.style.Chars.BoxLeft, pFinished, pInProgress, pUnfinished, p.style.Chars.BoxRight,
		))
	}
}

func (p *Progress) renderTrackerDone(out *strings.Builder, t *Tracker) {
	out.WriteString(p.style.Colors.Message.Sprint(t.Message))
	out.WriteString(p.style.Colors.Message.Sprint(p.style.Options.Separator))
	out.WriteString(p.style.Colors.Message.Sprint(p.style.Options.DoneString))
	p.renderTrackerStats(out, t)
	out.WriteRune('\n')
}

func (p *Progress) renderTrackerProgress(out *strings.Builder, t *Tracker, trackerStr string) {
	if p.trackerPosition == PositionRight {
		out.WriteString(p.style.Colors.Message.Sprint(t.Message))
		out.WriteString(p.style.Colors.Message.Sprint(p.style.Options.Separator))
		p.renderTrackerPercentage(out, t)
		if !p.hideTracker {
			out.WriteRune(' ')
			out.WriteString(p.style.Colors.Tracker.Sprint(trackerStr))
		}
		p.renderTrackerStats(out, t)
		out.WriteString("\n")
	} else {
		p.renderTrackerPercentage(out, t)
		if !p.hideTracker {
			out.WriteRune(' ')
			out.WriteString(p.style.Colors.Tracker.Sprint(trackerStr))
		}
		p.renderTrackerStats(out, t)
		out.WriteString(p.style.Colors.Message.Sprint(p.style.Options.Separator))
		out.WriteString(p.style.Colors.Message.Sprint(t.Message))
		out.WriteString(p.style.Colors.Message.Sprint(t.Message))
		out.WriteRune(' ')
	}
}

func (p *Progress) renderTrackerPercentage(out *strings.Builder, t *Tracker) {
	if !p.hidePercentage {
		out.WriteString(p.style.Colors.Percent.Sprintf(p.style.Options.PercentFormat, t.PercentDone()))
	}
}

func (p *Progress) renderTrackerStats(out *strings.Builder, t *Tracker) {
	if !p.hideValue || !p.hideTime {
		var outStats strings.Builder
		outStats.WriteString(" [")
		if !p.hideValue {
			outStats.WriteString(p.style.Colors.Value.Sprint(t.Units.Sprint(t.value)))
		}
		if !p.hideValue && !p.hideTime {
			outStats.WriteRune(' ')
		}
		if !p.hideTime {
			outStats.WriteString("in ")
			if t.IsDone() {
				outStats.WriteString(p.style.Colors.Time.Sprint(
					t.timeStop.Sub(t.timeStart).Round(p.style.Options.TimeDonePrecision)))
			} else {
				outStats.WriteString(p.style.Colors.Time.Sprint(
					time.Since(t.timeStart).Round(p.style.Options.TimeInProgressPrecision)))
			}
		}
		outStats.WriteRune(']')

		out.WriteString(p.style.Colors.Stats.Sprint(outStats.String()))
	}
}
