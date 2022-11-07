package progress

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/jedib0t/go-pretty/v6/text"
)

// Render renders the Progress tracker and handles all existing trackers and
// those that are added dynamically while render is in progress.
func (p *Progress) Render() {
	if p.beginRender() {
		p.initForRender()

		lastRenderLength := 0
		ticker := time.NewTicker(p.updateFrequency)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				lastRenderLength = p.renderTrackers(lastRenderLength)
			case <-p.done:
				// always render the current state before finishing render in
				// case it hasn't been shown yet
				p.renderTrackers(lastRenderLength)
				p.endRender()
				return
			}
		}
	}
}

func (p *Progress) beginRender() bool {
	p.renderInProgressMutex.Lock()
	defer p.renderInProgressMutex.Unlock()

	if p.renderInProgress {
		return false
	}
	p.renderInProgress = true
	return true
}

func (p *Progress) consumeQueuedTrackers() {
	if p.LengthInQueue() > 0 {
		p.trackersActiveMutex.Lock()
		p.trackersInQueueMutex.Lock()
		p.trackersActive = append(p.trackersActive, p.trackersInQueue...)
		p.trackersInQueue = make([]*Tracker, 0)
		p.trackersInQueueMutex.Unlock()
		p.trackersActiveMutex.Unlock()
	}
}

func (p *Progress) endRender() {
	p.renderInProgressMutex.Lock()
	defer p.renderInProgressMutex.Unlock()

	p.renderInProgress = false
}

func (p *Progress) extractDoneAndActiveTrackers() ([]*Tracker, []*Tracker) {
	// move trackers waiting in queue to the active list
	p.consumeQueuedTrackers()

	// separate the active and done trackers
	var trackersActive, trackersDone []*Tracker
	var activeTrackersProgress int64
	p.trackersActiveMutex.RLock()
	for _, tracker := range p.trackersActive {
		if !tracker.IsDone() {
			trackersActive = append(trackersActive, tracker)
			activeTrackersProgress += int64(tracker.PercentDone())
		} else {
			trackersDone = append(trackersDone, tracker)
		}
	}
	p.trackersActiveMutex.RUnlock()
	p.sortBy.Sort(trackersDone)
	p.sortBy.Sort(trackersActive)

	// calculate the overall tracker's progress value
	p.overallTracker.value = int64(p.LengthDone()+len(trackersDone)) * 100
	p.overallTracker.value += activeTrackersProgress
	if len(trackersActive) == 0 {
		p.overallTracker.MarkAsDone()
	}
	return trackersActive, trackersDone
}

func (p *Progress) generateTrackerStr(t *Tracker, maxLen int, hint renderHint) string {
	value, total := t.valueAndTotal()
	if !hint.isOverallTracker && (total == 0 || value > total) {
		return p.generateTrackerStrIndeterminate(maxLen)
	}
	return p.generateTrackerStrDeterminate(value, total, maxLen)
}

// generateTrackerStrDeterminate generates the tracker string for the case where
// the Total value is known, and the progress percentage can be calculated.
func (p *Progress) generateTrackerStrDeterminate(value int64, total int64, maxLen int) string {
	pFinishedDots, pFinishedDotsFraction := 0.0, 0.0
	pDotValue := float64(total) / float64(maxLen)
	if pDotValue > 0 {
		pFinishedDots = float64(value) / pDotValue
		pFinishedDotsFraction = pFinishedDots - float64(int(pFinishedDots))
	}
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
	pFinishedStrLen := text.RuneWidthWithoutEscSequences(pFinished + pInProgress)
	if pFinishedStrLen < maxLen {
		pUnfinished = strings.Repeat(p.style.Chars.Unfinished, maxLen-pFinishedStrLen)
	}

	return p.style.Colors.Tracker.Sprintf("%s%s%s%s%s",
		p.style.Chars.BoxLeft, pFinished, pInProgress, pUnfinished, p.style.Chars.BoxRight,
	)
}

// generateTrackerStrDeterminate generates the tracker string for the case where
// the Total value is unknown, and the progress percentage cannot be calculated.
func (p *Progress) generateTrackerStrIndeterminate(maxLen int) string {
	indicator := p.style.Chars.Indeterminate(maxLen)

	pUnfinished := ""
	if indicator.Position > 0 {
		pUnfinished += strings.Repeat(p.style.Chars.Unfinished, indicator.Position)
	}
	pUnfinished += indicator.Text
	if text.RuneWidthWithoutEscSequences(pUnfinished) < maxLen {
		pUnfinished += strings.Repeat(p.style.Chars.Unfinished, maxLen-text.RuneWidthWithoutEscSequences(pUnfinished))
	}

	return p.style.Colors.Tracker.Sprintf("%s%s%s",
		p.style.Chars.BoxLeft, string(pUnfinished), p.style.Chars.BoxRight,
	)
}

func (p *Progress) moveCursorToTheTop(out *strings.Builder) {
	numLinesToMoveUp := len(p.trackersActive)
	if p.style.Visibility.TrackerOverall && p.overallTracker != nil && !p.overallTracker.IsDone() {
		numLinesToMoveUp++
	}
	if p.style.Visibility.Pinned {
		numLinesToMoveUp += p.pinnedMessageNumLines
	}
	for numLinesToMoveUp > 0 {
		out.WriteString(text.CursorUp.Sprint())
		out.WriteString(text.EraseLine.Sprint())
		numLinesToMoveUp--
	}
}

func (p *Progress) renderPinnedMessages(out *strings.Builder) {
	p.pinnedMessageMutex.RLock()
	defer p.pinnedMessageMutex.RUnlock()

	numLines := len(p.pinnedMessages)
	for _, msg := range p.pinnedMessages {
		msg = strings.TrimSpace(msg)
		out.WriteString(p.style.Colors.Pinned.Sprint(msg))
		out.WriteRune('\n')

		numLines += strings.Count(msg, "\n")
	}
	p.pinnedMessageNumLines = numLines
}

func (p *Progress) renderTracker(out *strings.Builder, t *Tracker, hint renderHint) {
	message := t.message()
	if strings.Contains(message, "\t") {
		message = strings.Replace(message, "\t", "    ", -1)
	}
	if strings.Contains(message, "\r") {
		message = strings.Replace(message, "\r", "", -1)
	}
	if p.messageWidth > 0 {
		messageLen := text.RuneWidthWithoutEscSequences(message)
		if messageLen < p.messageWidth {
			message = text.Pad(message, p.messageWidth, ' ')
		} else {
			message = text.Snip(message, p.messageWidth, p.style.Options.SnipIndicator)
		}
	}

	if hint.isOverallTracker {
		if !t.IsDone() {
			hint := renderHint{hideValue: true, isOverallTracker: true}
			p.renderTrackerProgress(out, t, message, p.generateTrackerStr(t, p.lengthProgressOverall, hint), hint)
		}
	} else {
		if t.IsDone() {
			p.renderTrackerDone(out, t, message)
		} else {
			hint := renderHint{hideTime: !p.style.Visibility.Time, hideValue: !p.style.Visibility.Value}
			p.renderTrackerProgress(out, t, message, p.generateTrackerStr(t, p.lengthProgress, hint), hint)
		}
	}
}

func (p *Progress) renderTrackerDone(out *strings.Builder, t *Tracker, message string) {
	out.WriteString(p.style.Colors.Message.Sprint(message))
	out.WriteString(p.style.Colors.Message.Sprint(p.style.Options.Separator))
	if !t.IsErrored() {
		out.WriteString(p.style.Colors.Message.Sprint(p.style.Options.DoneString))
	} else {
		out.WriteString(p.style.Colors.Error.Sprint(p.style.Options.ErrorString))
	}
	p.renderTrackerStats(out, t, renderHint{hideTime: !p.style.Visibility.Time, hideValue: !p.style.Visibility.Value})
	out.WriteRune('\n')
}

func (p *Progress) renderTrackerMessage(out *strings.Builder, t *Tracker, message string) {
	if !t.IsErrored() {
		out.WriteString(p.style.Colors.Message.Sprint(message))
	} else {
		out.WriteString(p.style.Colors.Error.Sprint(message))
	}
}

func (p *Progress) renderTrackerPercentage(out *strings.Builder, t *Tracker) {
	if p.style.Visibility.Percentage {
		var percentageStr string
		if t.IsIndeterminate() {
			percentageStr = p.style.Options.PercentIndeterminate
		} else {
			percentageStr = fmt.Sprintf(p.style.Options.PercentFormat, t.PercentDone())
		}
		out.WriteString(p.style.Colors.Percent.Sprint(percentageStr))
	}
}

func (p *Progress) renderTrackerProgress(out *strings.Builder, t *Tracker, message string, trackerStr string, hint renderHint) {
	if hint.isOverallTracker {
		out.WriteString(p.style.Colors.Tracker.Sprint(trackerStr))
		p.renderTrackerStats(out, t, hint)
		out.WriteRune('\n')
	} else if p.trackerPosition == PositionRight {
		p.renderTrackerMessage(out, t, message)
		out.WriteString(p.style.Colors.Message.Sprint(p.style.Options.Separator))
		p.renderTrackerPercentage(out, t)
		if p.style.Visibility.Tracker {
			out.WriteString(p.style.Colors.Tracker.Sprint(" " + trackerStr))
		}
		p.renderTrackerStats(out, t, hint)
		out.WriteRune('\n')
	} else {
		p.renderTrackerPercentage(out, t)
		if p.style.Visibility.Tracker {
			out.WriteString(p.style.Colors.Tracker.Sprint(" " + trackerStr))
		}
		p.renderTrackerStats(out, t, hint)
		out.WriteString(p.style.Colors.Message.Sprint(p.style.Options.Separator))
		p.renderTrackerMessage(out, t, message)
		out.WriteRune('\n')
	}
}

func (p *Progress) renderTrackers(lastRenderLength int) int {
	if p.LengthActive() == 0 {
		return 0
	}

	// buffer all output into a strings.Builder object
	var out strings.Builder
	out.Grow(lastRenderLength)

	// move up N times based on the number of active trackers
	if lastRenderLength > 0 {
		p.moveCursorToTheTop(&out)
	}

	// render the trackers that are done, and then the ones that are active
	p.renderTrackersDoneAndActive(&out)

	// render the overall tracker
	if p.style.Visibility.TrackerOverall {
		p.renderTracker(&out, p.overallTracker, renderHint{isOverallTracker: true})
	}

	// write the text to the output writer
	_, _ = p.outputWriter.Write([]byte(out.String()))

	// stop if auto stop is enabled and there are no more active trackers
	if p.autoStop && p.LengthActive() == 0 {
		p.done <- true
	}

	return out.Len()
}

func (p *Progress) renderTrackersDoneAndActive(out *strings.Builder) {
	// find the currently "active" and "done" trackers
	trackersActive, trackersDone := p.extractDoneAndActiveTrackers()

	// sort and render the done trackers
	for _, tracker := range trackersDone {
		p.renderTracker(out, tracker, renderHint{})
	}
	p.trackersDoneMutex.Lock()
	p.trackersDone = append(p.trackersDone, trackersDone...)
	p.trackersDoneMutex.Unlock()

	// render all the logs received and flush them out
	p.logsToRenderMutex.Lock()
	for _, log := range p.logsToRender {
		out.WriteString(text.EraseLine.Sprint())
		out.WriteString(log)
		out.WriteRune('\n')
	}
	p.logsToRender = nil
	p.logsToRenderMutex.Unlock()

	// render pinned messages
	if len(trackersActive) > 0 && p.style.Visibility.Pinned {
		p.renderPinnedMessages(out)
	}

	// sort and render the active trackers
	for _, tracker := range trackersActive {
		p.renderTracker(out, tracker, renderHint{})
	}
	p.trackersActiveMutex.Lock()
	p.trackersActive = trackersActive
	p.trackersActiveMutex.Unlock()
}

func (p *Progress) renderTrackerStats(out *strings.Builder, t *Tracker, hint renderHint) {
	if !hint.hideValue || !hint.hideTime {
		var outStats strings.Builder
		outStats.WriteString(" [")

		if p.style.Options.SpeedPosition == PositionLeft {
			p.renderTrackerStatsSpeed(&outStats, t, hint)
		}
		if !hint.hideValue {
			outStats.WriteString(p.style.Colors.Value.Sprint(t.Units.Sprint(t.Value())))
		}
		if !hint.hideValue && !hint.hideTime {
			outStats.WriteString(" in ")
		}
		if !hint.hideTime {
			p.renderTrackerStatsTime(&outStats, t, hint)
		}
		if p.style.Options.SpeedPosition == PositionRight {
			p.renderTrackerStatsSpeed(&outStats, t, hint)
		}
		outStats.WriteRune(']')

		out.WriteString(p.style.Colors.Stats.Sprint(outStats.String()))
	}
}

func (p *Progress) renderTrackerStatsSpeed(out *strings.Builder, t *Tracker, hint renderHint) {
	if hint.isOverallTracker && !p.style.Visibility.SpeedOverall {
		return
	}
	if !hint.isOverallTracker && !p.style.Visibility.Speed {
		return
	}

	speedPrecision := p.style.Options.SpeedPrecision
	if hint.isOverallTracker {
		speed := float64(0)

		p.trackersActiveMutex.RLock()
		for _, tracker := range p.trackersActive {
			speed += float64(tracker.Value()) / time.Since(tracker.timeStart).Round(speedPrecision).Seconds()
		}
		p.trackersActiveMutex.RUnlock()

		if speed > 0 {
			p.renderTrackerStatsSpeedInternal(out, p.style.Options.SpeedOverallFormatter(int64(speed)))
		}
	} else {
		timeTaken := time.Since(t.timeStart)
		if timeTakenRounded := timeTaken.Round(speedPrecision); timeTakenRounded > speedPrecision {
			p.renderTrackerStatsSpeedInternal(out, t.Units.Sprint(int64(float64(t.Value())/timeTakenRounded.Seconds())))
		}
	}
}

func (p *Progress) renderTrackerStatsSpeedInternal(out *strings.Builder, speed string) {
	if p.style.Options.SpeedPosition == PositionRight {
		out.WriteString("; ")
	}
	out.WriteString(p.style.Colors.Speed.Sprint(speed))
	out.WriteString(p.style.Options.SpeedSuffix)
	if p.style.Options.SpeedPosition == PositionLeft {
		out.WriteString("; ")
	}
}

func (p *Progress) renderTrackerStatsTime(outStats *strings.Builder, t *Tracker, hint renderHint) {
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

	p.renderTrackerStatsETA(outStats, t, hint)
}

func (p *Progress) renderTrackerStatsETA(out *strings.Builder, t *Tracker, hint renderHint) {
	if hint.isOverallTracker && !p.style.Visibility.ETAOverall {
		return
	}
	if !hint.isOverallTracker && !p.style.Visibility.ETA {
		return
	}

	tpETA := p.style.Options.ETAPrecision
	if eta := t.ETA().Round(tpETA); hint.isOverallTracker || eta > tpETA {
		out.WriteString("; ")
		out.WriteString(p.style.Options.ETAString)
		out.WriteString(": ")
		out.WriteString(p.style.Colors.Time.Sprint(eta))
	}
}
