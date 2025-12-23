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
			case <-p.renderContext.Done():
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

func (p *Progress) collectActiveTrackers() ([]*Tracker, int64, time.Duration) {
	var allTrackers []*Tracker
	var activeTrackersProgress int64
	var maxETA time.Duration

	p.trackersDoneMutex.RLock()
	lengthDone := len(p.trackersDone)
	p.trackersDoneMutex.RUnlock()

	p.trackersActiveMutex.RLock()
	allTrackers = make([]*Tracker, 0, len(p.trackersActive)+lengthDone)
	for _, tracker := range p.trackersActive {
		if !tracker.IsDone() || !tracker.RemoveOnCompletion {
			allTrackers = append(allTrackers, tracker)
			if !tracker.IsDone() {
				activeTrackersProgress += int64(tracker.PercentDone())
				if eta := tracker.ETA(); eta > maxETA {
					maxETA = eta
				}
			}
		}
	}
	p.trackersActiveMutex.RUnlock()

	return allTrackers, activeTrackersProgress, maxETA
}

func (p *Progress) collectDoneTrackers(allTrackers *[]*Tracker) {
	p.trackersDoneMutex.RLock()
	for _, tracker := range p.trackersDone {
		if !tracker.RemoveOnCompletion {
			*allTrackers = append(*allTrackers, tracker)
		}
	}
	p.trackersDoneMutex.RUnlock()
}

func (p *Progress) consumeQueuedTrackers() {
	p.trackersInQueueMutex.Lock()
	queueLen := len(p.trackersInQueue)
	if queueLen == 0 {
		p.trackersInQueueMutex.Unlock()
		return
	}
	// copy the slice to avoid race condition - another goroutine may append
	// to p.trackersInQueue while we're appending to p.trackersActive
	queued := make([]*Tracker, len(p.trackersInQueue))
	copy(queued, p.trackersInQueue)
	p.trackersInQueue = p.trackersInQueue[:0] // reuse slice capacity
	p.trackersInQueueMutex.Unlock()

	p.trackersActiveMutex.Lock()
	p.trackersActive = append(p.trackersActive, queued...)
	p.trackersActiveMutex.Unlock()
}

func (p *Progress) endRender() {
	p.renderInProgressMutex.Lock()
	defer p.renderInProgressMutex.Unlock()

	p.renderInProgress = false
}

// extractAllTrackersInOrder extracts all trackers (both active and done) and
// sorts them together when SortByIndex is used. This allows maintaining a fixed
// order regardless of completion status.
func (p *Progress) extractAllTrackersInOrder() []*Tracker {
	// move trackers waiting in queue to the active list
	p.consumeQueuedTrackers()

	allTrackers, activeTrackersProgress, maxETA := p.collectActiveTrackers()
	p.collectDoneTrackers(&allTrackers)

	// Sort by Index (ascending or descending)
	if p.sortBy == SortByIndex || p.sortBy == SortByIndexDsc {
		p.sortBy.Sort(allTrackers)
	}

	p.updateOverallTrackerProgress(allTrackers, activeTrackersProgress, maxETA)

	return allTrackers
}

func (p *Progress) generateTrackerStr(t *Tracker, maxLen int, hint renderHint) string {
	value, total := t.valueAndTotal()
	if !hint.isOverallTracker && t.IsStarted() && (total == 0 || value > total) {
		return p.generateTrackerStrIndeterminate(maxLen)
	}
	return p.generateTrackerStrDeterminate(value, total, maxLen)
}

// generateTrackerStrDeterminate generates the tracker string for the case where
// the Total value is known, and the progress percentage can be calculated.
func (p *Progress) generateTrackerStrDeterminate(value int64, total int64, maxLen int) string {
	if p.style.Renderer.TrackerDeterminate != nil {
		return p.style.Renderer.TrackerDeterminate(value, total, maxLen)
	}

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

	// Use strings.Builder to avoid temporary string allocation
	var combined strings.Builder
	combined.Grow(len(pFinished) + len(pInProgress))
	combined.WriteString(pFinished)
	combined.WriteString(pInProgress)
	pFinishedStrLen := text.StringWidthWithoutEscSequences(combined.String())
	if pFinishedStrLen < maxLen {
		pUnfinished = strings.Repeat(p.style.Chars.Unfinished, maxLen-pFinishedStrLen)
	}

	return p.style.Colors.Tracker.Sprintf("%s%s%s%s%s",
		p.style.Chars.BoxLeft, pFinished, pInProgress, pUnfinished, p.style.Chars.BoxRight,
	)
}

// generateTrackerStrIndeterminate generates the tracker string for the case where
// the Total value is unknown, and the progress percentage cannot be calculated.
func (p *Progress) generateTrackerStrIndeterminate(maxLen int) string {
	if p.style.Renderer.TrackerIndeterminate != nil {
		return p.style.Renderer.TrackerIndeterminate(maxLen)
	}

	indicator := p.style.Chars.Indeterminate(maxLen)

	pUnfinished := ""
	if indicator.Position > 0 {
		pUnfinished += strings.Repeat(p.style.Chars.Unfinished, indicator.Position)
	}
	pUnfinished += indicator.Text
	if text.StringWidthWithoutEscSequences(pUnfinished) < maxLen {
		pUnfinished += strings.Repeat(p.style.Chars.Unfinished, maxLen-text.StringWidthWithoutEscSequences(pUnfinished))
	}

	return p.style.Colors.Tracker.Sprintf("%s%s%s",
		p.style.Chars.BoxLeft, pUnfinished, p.style.Chars.BoxRight,
	)
}

func (p *Progress) moveCursorToTheTop(out *strings.Builder) {
	// Count all trackers that will be rendered (both done and active)
	var numTrackersToRender int

	// Count active trackers (excluding those with RemoveOnCompletion)
	p.trackersActiveMutex.RLock()
	for _, tracker := range p.trackersActive {
		if !tracker.RemoveOnCompletion {
			numTrackersToRender++
		}
	}
	p.trackersActiveMutex.RUnlock()

	// Count done trackers (excluding those with RemoveOnCompletion)
	p.trackersDoneMutex.RLock()
	for _, tracker := range p.trackersDone {
		if !tracker.RemoveOnCompletion {
			numTrackersToRender++
		}
	}
	p.trackersDoneMutex.RUnlock()

	numLinesToMoveUp := numTrackersToRender
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

func (p *Progress) renderPinnedMessages(out *strings.Builder, hint renderHint) {
	p.pinnedMessageMutex.RLock()
	defer p.pinnedMessageMutex.RUnlock()

	numLines := len(p.pinnedMessages)
	for _, msg := range p.pinnedMessages {
		msg = strings.TrimSpace(msg)
		msg = p.style.Colors.Pinned.Sprint(msg)
		if hint.terminalWidth > 0 {
			msg = text.Trim(msg, hint.terminalWidth)
		}
		out.WriteString(msg)
		out.WriteRune('\n')

		numLines += strings.Count(msg, "\n")
	}
	p.pinnedMessageNumLines = numLines
}

func (p *Progress) renderTracker(out *strings.Builder, t *Tracker, hint renderHint) {
	message := t.message()
	// Optimize: only process if message contains tabs or carriage returns
	if strings.ContainsAny(message, "\t\r") {
		message = strings.ReplaceAll(message, "\t", "    ")
		message = strings.ReplaceAll(message, "\r", "")
	}
	if p.lengthMessage > 0 {
		messageLen := text.StringWidthWithoutEscSequences(message)
		if messageLen < p.lengthMessage {
			message = text.Pad(message, p.lengthMessage, ' ')
		} else {
			message = text.Snip(message, p.lengthMessage, p.style.Options.SnipIndicator)
		}
	}

	tOut := &strings.Builder{}
	tOut.Grow(p.lengthProgressOverall)
	if hint.isOverallTracker {
		if !t.IsDone() {
			hint := renderHint{hideValue: true, isOverallTracker: true, terminalWidth: hint.terminalWidth}
			p.renderTrackerProgress(tOut, t, message, p.generateTrackerStr(t, p.lengthProgressOverall, hint), hint)
		}
	} else {
		if t.IsDone() {
			p.renderTrackerDone(tOut, t, message)
		} else {
			hint := renderHint{hideTime: !p.style.Visibility.Time, hideValue: !p.style.Visibility.Value, terminalWidth: hint.terminalWidth}
			p.renderTrackerProgress(tOut, t, message, p.generateTrackerStr(t, p.lengthProgress, hint), hint)
		}
	}

	outStr := tOut.String()
	if hint.terminalWidth > 0 {
		outStr = text.Trim(outStr, hint.terminalWidth)
	}
	out.WriteString(outStr)
	out.WriteRune('\n')
}

func (p *Progress) renderTrackerDone(out *strings.Builder, t *Tracker, message string) {
	if !t.RemoveOnCompletion {
		out.WriteString(p.style.Colors.Message.Sprint(message))
		out.WriteString(p.style.Colors.Message.Sprint(p.style.Options.Separator))
		if !t.IsErrored() {
			out.WriteString(p.style.Colors.Message.Sprint(p.style.Options.DoneString))
		} else {
			out.WriteString(p.style.Colors.Error.Sprint(p.style.Options.ErrorString))
		}
		p.renderTrackerStats(out, t, renderHint{hideTime: !p.style.Visibility.Time, hideValue: !p.style.Visibility.Value})
	}
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
	} else if p.trackerPosition == PositionRight {
		p.renderTrackerMessage(out, t, message)
		out.WriteString(p.style.Colors.Message.Sprint(p.style.Options.Separator))
		p.renderTrackerPercentage(out, t)
		if p.style.Visibility.Tracker {
			out.WriteString(p.style.Colors.Tracker.Sprint(" " + trackerStr))
		}
		p.renderTrackerStats(out, t, hint)
	} else {
		p.renderTrackerPercentage(out, t)
		if p.style.Visibility.Tracker {
			out.WriteString(p.style.Colors.Tracker.Sprint(" " + trackerStr))
		}
		p.renderTrackerStats(out, t, hint)
		out.WriteString(p.style.Colors.Message.Sprint(p.style.Options.Separator))
		p.renderTrackerMessage(out, t, message)
	}
}

func (p *Progress) renderTrackers(lastRenderLength int) int {
	if p.LengthActive() == 0 {
		return 0
	}

	// Cache terminal width once per render cycle to avoid repeated mutex locks
	terminalWidth := p.getTerminalWidth()
	hint := renderHint{terminalWidth: terminalWidth}

	// buffer all output into a strings.Builder object
	var out strings.Builder
	out.Grow(lastRenderLength)

	// move up N times based on the number of active trackers
	if lastRenderLength > 0 {
		p.moveCursorToTheTop(&out)
	}

	// render the trackers that are done, and then the ones that are active
	p.renderTrackersDoneAndActive(&out, hint)

	// render the overall tracker
	if p.style.Visibility.TrackerOverall {
		overallHint := renderHint{isOverallTracker: true, terminalWidth: terminalWidth}
		p.renderTracker(&out, p.overallTracker, overallHint)
	}

	// write the text to the output writer
	p.outputWriterMutex.Lock()
	_, _ = p.outputWriter.Write([]byte(out.String()))
	p.outputWriterMutex.Unlock()

	// stop if auto stop is enabled and there are no more active trackers
	if p.autoStop && p.LengthActive() == 0 {
		p.renderContextCancelMutex.Lock()
		if p.renderContextCancel != nil {
			p.renderContextCancel()
		}
		p.renderContextCancelMutex.Unlock()
	}

	return out.Len()
}

func (p *Progress) renderTrackersDoneAndActive(out *strings.Builder, hint renderHint) {
	// Extract all trackers (both active and done)
	allTrackers := p.extractAllTrackersInOrder()

	// Separate done and active trackers for sorting and state management
	trackersDone, trackersActive := p.separateDoneAndActiveTrackers(allTrackers)

	// Sort trackers based on sortBy setting
	trackersToRender := p.sortTrackersForRendering(allTrackers, trackersDone, trackersActive)

	// Render all trackers in the determined order
	for _, tracker := range trackersToRender {
		p.renderTracker(out, tracker, hint)
	}

	// Update internal state
	p.trackersDoneMutex.Lock()
	// Only add newly done trackers that aren't already in trackersDone
	existingDone := make(map[*Tracker]bool)
	for _, t := range p.trackersDone {
		existingDone[t] = true
	}
	for _, t := range trackersDone {
		if !existingDone[t] {
			p.trackersDone = append(p.trackersDone, t)
		}
	}
	p.trackersDoneMutex.Unlock()

	p.trackersActiveMutex.Lock()
	p.trackersActive = trackersActive
	p.trackersActiveMutex.Unlock()

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
		p.renderPinnedMessages(out, hint)
	}
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
			timeStart := tracker.timeStartValue()
			if !timeStart.IsZero() {
				speed += float64(tracker.Value()) / time.Since(timeStart).Round(speedPrecision).Seconds()
			}
		}
		p.trackersActiveMutex.RUnlock()

		if speed > 0 {
			p.renderTrackerStatsSpeedInternal(out, p.style.Options.SpeedOverallFormatter(int64(speed)))
		}
	} else {
		timeStart := t.timeStartValue()
		if !timeStart.IsZero() {
			timeTaken := time.Since(timeStart)
			if timeTakenRounded := timeTaken.Round(speedPrecision); timeTakenRounded > speedPrecision {
				p.renderTrackerStatsSpeedInternal(out, t.Units.Sprint(int64(float64(t.Value())/timeTakenRounded.Seconds())))
			}
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
	timeStart, timeStop := t.timeStartAndStop()
	if !timeStart.IsZero() {
		if t.IsDone() {
			td = timeStop.Sub(timeStart)
		} else {
			td = time.Since(timeStart)
		}
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

func (p *Progress) separateDoneAndActiveTrackers(allTrackers []*Tracker) ([]*Tracker, []*Tracker) {
	var trackersDone, trackersActive []*Tracker
	for _, tracker := range allTrackers {
		if tracker.IsDone() {
			if !tracker.RemoveOnCompletion {
				trackersDone = append(trackersDone, tracker)
			}
		} else {
			trackersActive = append(trackersActive, tracker)
		}
	}
	return trackersDone, trackersActive
}

func (p *Progress) sortTrackersForRendering(allTrackers []*Tracker, trackersDone []*Tracker, trackersActive []*Tracker) []*Tracker {
	if p.sortBy == SortByIndex || p.sortBy == SortByIndexDsc {
		// For explicit index ordering (ascending or descending), all trackers are already sorted together
		return allTrackers
	}
	// For other sort methods, sort done and active separately, then combine
	p.sortBy.Sort(trackersDone)
	p.sortBy.Sort(trackersActive)
	// Combine: done first, then active
	return append(trackersDone, trackersActive...)
}

func (p *Progress) updateOverallTrackerProgress(allTrackers []*Tracker, activeTrackersProgress int64, maxETA time.Duration) {
	doneCount := 0
	for _, tracker := range allTrackers {
		if tracker.IsDone() {
			doneCount++
		}
	}
	p.overallTracker.value = int64(doneCount) * 100
	p.overallTracker.value += activeTrackersProgress
	p.overallTracker.minETA = maxETA
	if len(allTrackers) > 0 && doneCount == len(allTrackers) {
		p.overallTracker.MarkAsDone()
	}
}
