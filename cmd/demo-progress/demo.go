package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/jedib0t/go-pretty/v6/progress"
	"github.com/jedib0t/go-pretty/v6/text"
)

var (
	flagAutoStop           = flag.Bool("auto-stop", false, "Auto-stop rendering?")
	flagHideETA            = flag.Bool("hide-eta", false, "Hide the ETA?")
	flagHideETAOverall     = flag.Bool("hide-eta-overall", false, "Hide the ETA in the overall tracker?")
	flagHideOverallTracker = flag.Bool("hide-overall", false, "Hide the Overall Tracker?")
	flagHidePercentage     = flag.Bool("hide-percentage", false, "Hide the progress percent?")
	flagHideTime           = flag.Bool("hide-time", false, "Hide the time taken?")
	flagHideValue          = flag.Bool("hide-value", false, "Hide the tracker value?")
	flagNumTrackers        = flag.Int("num-trackers", 13, "Number of Trackers")
	flagShowSpeed          = flag.Bool("show-speed", false, "Show the tracker speed?")
	flagShowSpeedOverall   = flag.Bool("show-speed-overall", false, "Show the overall tracker speed?")
	flagShowPinned         = flag.Bool("show-pinned", false, "Show a pinned message?")
	flagRandomFail         = flag.Bool("rnd-fail", false, "Introduce random failures in tracking")
	flagRandomDefer        = flag.Bool("rnd-defer", false, "Introduce random deferred starts")
	flagRandomRemove       = flag.Bool("rnd-remove", false, "Introduce random remove of trackers on completion")
	flagRandomLogs         = flag.Bool("rnd-logs", false, "Output random logs in the middle of tracking")
	flagCustomRender       = flag.Bool("custom-render", false, "Use custom render functions with rainbow colors")

	messageColors = []text.Color{
		text.FgRed,
		text.FgGreen,
		text.FgYellow,
		text.FgBlue,
		text.FgMagenta,
		text.FgCyan,
		text.FgWhite,
	}
	rng       = rand.New(rand.NewSource(time.Now().UnixNano()))
	timeStart = time.Now()
)

func getMessage(idx int64, units *progress.Units) string {
	var message string
	switch units {
	case &progress.UnitsBytes:
		message = fmt.Sprintf("Downloading File    #%3d", idx)
	case &progress.UnitsCurrencyDollar, &progress.UnitsCurrencyEuro, &progress.UnitsCurrencyPound:
		message = fmt.Sprintf("Transferring Amount #%3d", idx)
	default:
		message = fmt.Sprintf("Calculating Total   #%3d", idx)
	}
	return message
}

func getUnits(idx int64) *progress.Units {
	var units *progress.Units
	switch {
	case idx%5 == 0:
		units = &progress.UnitsCurrencyPound
	case idx%4 == 0:
		units = &progress.UnitsCurrencyDollar
	case idx%3 == 0:
		units = &progress.UnitsBytes
	default:
		units = &progress.UnitsDefault
	}
	return units
}

func trackSomething(pw progress.Writer, idx int64, updateMessage bool) {
	total := idx * idx * idx * 250
	incrementPerCycle := idx * int64(*flagNumTrackers) * 250

	units := getUnits(idx)
	message := getMessage(idx, units)
	tracker := progress.Tracker{
		DeferStart:         *flagRandomDefer && rng.Float64() < 0.5,
		Message:            message,
		RemoveOnCompletion: *flagRandomRemove && rng.Float64() < 0.25,
		Total:              total,
		Units:              *units,
	}
	if idx == int64(*flagNumTrackers) {
		tracker.Total = 0
	}

	pw.AppendTracker(&tracker)

	if tracker.DeferStart {
		time.Sleep(3 * time.Second)
		tracker.Start()
	}

	ticker := time.Tick(time.Millisecond * 500)
	updateTicker := time.Tick(time.Millisecond * 250)
	for !tracker.IsDone() {
		select {
		case <-ticker:
			tracker.Increment(incrementPerCycle)
			if idx == int64(*flagNumTrackers) && tracker.Value() >= total {
				tracker.MarkAsDone()
			} else if *flagRandomFail && rand.Float64() < 0.1 {
				tracker.MarkAsErrored()
			}
			pw.SetPinnedMessages(
				fmt.Sprintf(">> Current Time: %-32s", time.Now().Format(time.RFC3339)),
				fmt.Sprintf(">>   Total Time: %-32s", time.Since(timeStart).Round(time.Millisecond)),
			)
		case <-updateTicker:
			if updateMessage {
				rndIdx := rand.Intn(len(messageColors))
				if rndIdx == len(messageColors) {
					rndIdx--
				}
				tracker.UpdateMessage(messageColors[rndIdx].Sprint(message))
			}
		}
	}
}

// customTrackerRender creates a progress bar using rainbow colors for determinate progress
func customTrackerRender(value int64, total int64, maxLen int) string {
	progress := float64(value) / float64(total)
	completed := int(progress * float64(maxLen))

	var result strings.Builder
	for i := 0; i < maxLen; i++ {
		if i < completed {
			// Use rainbow colors based on position in the progress bar
			colorIdx := (i * 6) / maxLen // Map position to 6 rainbow colors
			colors := []text.Color{
				text.FgRed,
				text.FgYellow,
				text.FgGreen,
				text.FgCyan,
				text.FgBlue,
				text.FgMagenta,
			}
			if colorIdx >= len(colors) {
				colorIdx = len(colors) - 1
			}
			result.WriteString(colors[colorIdx].Sprint("█"))
		} else {
			result.WriteString(text.FgHiBlack.Sprint("░"))
		}
	}

	return result.String()
}

// customTrackerIndeterminateRender creates a progress bar using rotating rainbow colors for indeterminate progress
func customTrackerIndeterminateRender(maxLen int) string {
	// For indeterminate progress, use rotating rainbow colors
	colors := []text.Color{
		text.FgRed,
		text.FgYellow,
		text.FgGreen,
		text.FgCyan,
		text.FgBlue,
		text.FgMagenta,
	}
	idx := int(time.Now().UnixNano()/100000000) % len(colors)
	return colors[idx].Sprint(strings.Repeat("█", maxLen))
}

func main() {
	flag.Parse()
	fmt.Printf("Tracking Progress of %d trackers ...\n\n", *flagNumTrackers)

	// instantiate a Progress Writer and set up the options
	pw := progress.NewWriter()
	pw.SetAutoStop(*flagAutoStop)
	pw.SetMessageLength(24)
	pw.SetNumTrackersExpected(*flagNumTrackers)
	pw.SetSortBy(progress.SortByPercentDsc)
	pw.SetStyle(progress.StyleDefault)
	pw.SetTrackerLength(25)
	pw.SetTrackerPosition(progress.PositionRight)
	pw.SetUpdateFrequency(time.Millisecond * 100)
	pw.Style().Colors = progress.StyleColorsExample
	pw.Style().Options.PercentFormat = "%4.1f%%"
	pw.Style().Visibility.ETA = !*flagHideETA
	pw.Style().Visibility.ETAOverall = !*flagHideETAOverall
	pw.Style().Visibility.Percentage = !*flagHidePercentage
	pw.Style().Visibility.Speed = *flagShowSpeed
	pw.Style().Visibility.SpeedOverall = *flagShowSpeedOverall
	pw.Style().Visibility.Time = !*flagHideTime
	pw.Style().Visibility.TrackerOverall = !*flagHideOverallTracker
	pw.Style().Visibility.Value = !*flagHideValue
	pw.Style().Visibility.Pinned = *flagShowPinned

	// set up custom render functions if flag is enabled
	if *flagCustomRender {
		pw.Style().Renderer.TrackerDeterminate = customTrackerRender
		pw.Style().Renderer.TrackerIndeterminate = customTrackerIndeterminateRender
		fmt.Println("Using custom render functions with rainbow colors!")
	}

	// call Render() in async mode; yes we don't have any trackers at the moment
	go pw.Render()

	// add a bunch of trackers with random parameters to demo most of the
	// features available; do this in async too like a client might do (for ex.
	// when downloading a bunch of files in parallel)
	for idx := int64(1); idx <= int64(*flagNumTrackers); idx++ {
		go trackSomething(pw, idx, idx == int64(*flagNumTrackers))

		// in auto-stop mode, the Render logic terminates the moment it detects
		// zero active trackers; but in a manual-stop mode, it keeps waiting and
		// is a good chance to demo trackers being added dynamically while other
		// trackers are active or done
		if !*flagAutoStop {
			time.Sleep(time.Millisecond * 100)
		}
	}

	// wait for one or more trackers to become active (just blind-wait for a
	// second) and then keep watching until Rendering is in progress
	time.Sleep(time.Second)
	messagesLogged := make(map[string]bool)
	for pw.IsRenderInProgress() {
		if *flagRandomLogs && pw.LengthDone()%3 == 0 {
			logMsg := text.Faint.Sprintf("[INFO] done with %d trackers", pw.LengthDone())
			if !messagesLogged[logMsg] {
				pw.Log(logMsg)
				messagesLogged[logMsg] = true
			}
		}

		// for manual-stop mode, stop when there are no more active trackers
		if !*flagAutoStop && pw.LengthActive() == 0 {
			pw.Stop()
		}
		time.Sleep(time.Millisecond * 100)
	}

	fmt.Println("\nAll done!")
}
