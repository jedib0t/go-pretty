package progress

import (
	"math"
	"strings"
	"time"

	"github.com/jedib0t/go-pretty/text"
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
	if lastRenderLength > 0 {
		p.moveCursorToTheTop(&out)
	}

	// find the currently "active" and "done" trackers
	trackersActive, trackersDone := p.extractDoneAndActiveTrackers()

	// sort and render the done trackers
	for _, tracker := range trackersDone {
		p.renderTracker(&out, tracker, renderHint{})
	}
	p.trackersDone = append(p.trackersDone, trackersDone...)

	// sort and render the active trackers
	for _, tracker := range trackersActive {
		p.renderTracker(&out, tracker, renderHint{})
	}
	p.trackersActive = trackersActive

	// render the overall tracker
	p.renderTracker(&out, p.overallTracker, renderHint{isOverallTracker: true})

	// write the text to the output writer
	p.outputWriter.Write([]byte(out.String()))

	// stop if auto stop is enabled and there are no more active trackers
	if p.autoStop && len(p.trackersInQueue) == 0 && len(p.trackersActive) == 0 {
		p.done <- true
	}

	return out.Len()
}

func (p *Progress) consumeQueuedTrackers() {
	if len(p.trackersInQueue) > 0 {
		p.trackersInQueueMutex.Lock()
		p.trackersActive = append(p.trackersActive, p.trackersInQueue...)
		p.trackersInQueue = make([]*Tracker, 0)
		p.trackersInQueueMutex.Unlock()
	}
}

func (p *Progress) extractDoneAndActiveTrackers() ([]*Tracker, []*Tracker) {
	// move trackers waiting in queue to the active list
	p.consumeQueuedTrackers()

	// separate the active and done trackers
	var trackersActive, trackersDone []*Tracker
	var activeTrackersProgress int64
	for _, tracker := range p.trackersActive {
		if !tracker.IsDone() {
			trackersActive = append(trackersActive, tracker)
			activeTrackersProgress += int64(tracker.PercentDone())
		} else {
			trackersDone = append(trackersDone, tracker)
		}
	}
	p.sortBy.Sort(trackersDone)
	p.sortBy.Sort(trackersActive)

	// calculate the overall tracker's progress value
	p.overallTracker.value = int64(len(p.trackersDone)+len(trackersDone)) * 100
	p.overallTracker.value += activeTrackersProgress
	if len(trackersActive) == 0 {
		p.overallTracker.MarkAsDone()
	}
	return trackersActive, trackersDone
}

func (p *Progress) generateTrackerStr(t *Tracker, maxLen int) string {
	pDotValue := float64(t.Total) / float64(maxLen)
	pFinishedDots := float64(t.value) / pDotValue
	pFinishedDotsFraction := pFinishedDots - float64(int(pFinishedDots))
	pFinishedLen := int(math.Floor(pFinishedDots))

	var pFinished, pInProgress, pUnfinished string
	if pFinishedLen > 0 {
		pFinished = strings.Repeat(p.style.Chars.Finished, pFinishedLen)
	}
	pInProgress = p.style.Chars.Unfinished
	if pFinishedDotsFraction >= 0.75 {
		pInProgress = p.style.Chars.Finished75
	} else if pFinishedDotsFraction >= 0.50 {
		pInProgress = p.style.Chars.Finished50
	} else if pFinishedDotsFraction >= 0.25 {
		pInProgress = p.style.Chars.Finished25
	} else if pFinishedDotsFraction == 0 {
		pInProgress = ""
	}
	pFinishedStrLen := text.RuneCountWithoutEscapeSeq(pFinished + pInProgress)
	if pFinishedStrLen < maxLen {
		pUnfinished = strings.Repeat(p.style.Chars.Unfinished, maxLen-pFinishedStrLen)
	}

	return p.style.Colors.Tracker.Sprintf("%s%s%s%s%s",
		p.style.Chars.BoxLeft, pFinished, pInProgress, pUnfinished, p.style.Chars.BoxRight,
	)
}

func (p *Progress) moveCursorToTheTop(out *strings.Builder) {
	numLinesToMoveUp := len(p.trackersActive)
	if p.showOverallTracker && p.overallTracker != nil && !p.overallTracker.IsDone() {
		numLinesToMoveUp++
	}
	if numLinesToMoveUp > 0 {
		out.WriteString(text.CursorUp.Sprintn(numLinesToMoveUp))
	}
}

func (p *Progress) renderTracker(out *strings.Builder, t *Tracker, hint renderHint) {
	if hint.isOverallTracker && !p.showOverallTracker {
		return
	}

	out.WriteString(text.EraseLine.Sprint())
	if hint.isOverallTracker {
		if !t.IsDone() {
			trackerLen := p.messageWidth
			trackerLen += text.RuneCountWithoutEscapeSeq(p.style.Options.Separator)
			trackerLen += text.RuneCountWithoutEscapeSeq(p.style.Options.DoneString)
			trackerLen += p.lengthProgress + 1
			hint := renderHint{hideValue: true, isOverallTracker: true}
			p.renderTrackerProgress(out, t, p.generateTrackerStr(t, trackerLen), hint)
		}
	} else {
		if t.IsDone() {
			p.renderTrackerDone(out, t)
		} else {
			hint := renderHint{hideTime: p.hideTime, hideValue: p.hideValue}
			p.renderTrackerProgress(out, t, p.generateTrackerStr(t, p.lengthProgress), hint)
		}
	}
}

func (p *Progress) renderTrackerDone(out *strings.Builder, t *Tracker) {
	out.WriteString(p.style.Colors.Message.Sprint(t.Message))
	out.WriteString(p.style.Colors.Message.Sprint(p.style.Options.Separator))
	out.WriteString(p.style.Colors.Message.Sprint(p.style.Options.DoneString))
	p.renderTrackerStats(out, t, renderHint{hideTime: p.hideTime, hideValue: p.hideValue})
	out.WriteRune('\n')
}

func (p *Progress) renderTrackerProgress(out *strings.Builder, t *Tracker, trackerStr string, hint renderHint) {
	if p.messageWidth > 0 {
		t.Message = text.FixedLengthString(t.Message, p.messageWidth, p.style.Options.SnipIndicator)
	}

	if hint.isOverallTracker {
		out.WriteString(p.style.Colors.Tracker.Sprint(trackerStr))
		p.renderTrackerStats(out, t, hint)
		out.WriteRune('\n')
	} else if p.trackerPosition == PositionRight {
		out.WriteString(p.style.Colors.Message.Sprint(t.Message))
		out.WriteString(p.style.Colors.Message.Sprint(p.style.Options.Separator))
		p.renderTrackerPercentage(out, t)
		if !p.hideTracker {
			out.WriteString(p.style.Colors.Tracker.Sprint(" " + trackerStr))
		}
		p.renderTrackerStats(out, t, hint)
		out.WriteRune('\n')
	} else {
		p.renderTrackerPercentage(out, t)
		if !p.hideTracker {
			out.WriteString(p.style.Colors.Tracker.Sprint(" " + trackerStr))
		}
		p.renderTrackerStats(out, t, hint)
		out.WriteString(p.style.Colors.Message.Sprint(p.style.Options.Separator))
		out.WriteString(p.style.Colors.Message.Sprint(t.Message))
		out.WriteRune('\n')
	}
}

func (p *Progress) renderTrackerPercentage(out *strings.Builder, t *Tracker) {
	if !p.hidePercentage {
		out.WriteString(p.style.Colors.Percent.Sprintf(p.style.Options.PercentFormat, t.PercentDone()))
	}
}

func (p *Progress) renderTrackerStats(out *strings.Builder, t *Tracker, hint renderHint) {
	if !hint.hideValue || !hint.hideTime {
		var outStats strings.Builder
		outStats.WriteString(" [")
		if !hint.hideValue {
			outStats.WriteString(p.style.Colors.Value.Sprint(t.Units.Sprint(t.value)))
		}
		if !hint.hideValue && !hint.hideTime {
			outStats.WriteString(" in ")
		}
		if !hint.hideTime {
			var td, tp time.Duration
			if t.IsDone() {
				td = t.timeStop.Sub(t.timeStart)
			} else {
				td = time.Since(t.timeStart)
			}
			if hint.isOverallTracker {
				tp = p.style.Options.TimeOverallPrecision
			} else if t.IsDone() {
				tp = p.style.Options.TimeDonePrecision
			} else {
				tp = p.style.Options.TimeInProgressPrecision
			}
			outStats.WriteString(p.style.Colors.Time.Sprint(td.Round(tp)))
			if hint.isOverallTracker {
				tpO := p.style.Options.TimeOverallPrecision
				if eta := t.ETA().Round(tpO) + tpO; true || eta > tpO {
					outStats.WriteString("; ~ETA: ")
					outStats.WriteString(p.style.Colors.Time.Sprint(eta))
				}
			}
		}
		outStats.WriteRune(']')

		out.WriteString(p.style.Colors.Stats.Sprint(outStats.String()))
	}
}
