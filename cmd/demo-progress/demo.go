package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/jedib0t/go-pretty/v6/progress"
	"github.com/jedib0t/go-pretty/v6/text"
)

var (
	autoStop    = flag.Bool("auto-stop", false, "Auto-stop rendering?")
	randomFail  = flag.Bool("rnd-fail", false, "Enable random failures")
	numTrackers = flag.Int("num-trackers", 13, "Number of Trackers")

	messageColors = []text.Color{
		text.FgRed,
		text.FgGreen,
		text.FgYellow,
		text.FgBlue,
		text.FgMagenta,
		text.FgCyan,
		text.FgWhite,
	}
)

func trackSomething(pw progress.Writer, idx int64, updateMessage bool) {
	total := idx * idx * idx * 250
	incrementPerCycle := idx * int64(*numTrackers) * 250

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

	var message string
	switch units {
	case &progress.UnitsBytes:
		message = fmt.Sprintf("Downloading File    #%3d", idx)
	case &progress.UnitsCurrencyDollar, &progress.UnitsCurrencyEuro, &progress.UnitsCurrencyPound:
		message = fmt.Sprintf("Transferring Amount #%3d", idx)
	default:
		message = fmt.Sprintf("Calculating Total   #%3d", idx)
	}
	tracker := progress.Tracker{Message: message, Total: total, Units: *units}
	if idx == int64(*numTrackers) {
		tracker.Total = 0
	}

	pw.AppendTracker(&tracker)

	ticker := time.Tick(time.Millisecond * 500)
	updateTicker := time.Tick(time.Millisecond * 250)
	for !tracker.IsDone() {
		select {
		case <-ticker:
			tracker.Increment(incrementPerCycle)
			if idx == int64(*numTrackers) && tracker.Value() >= total {
				tracker.MarkAsDone()
			} else if *randomFail && rand.Float64() < 0.1 {
				tracker.MarkAsErrored()
			}
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

func main() {
	flag.Parse()
	fmt.Printf("Tracking Progress of %d trackers ...\n\n", *numTrackers)

	// instantiate a Progress Writer and set up the options
	pw := progress.NewWriter()
	pw.SetAutoStop(*autoStop)
	pw.SetTrackerLength(25)
	pw.ShowETA(true)
	pw.ShowOverallTracker(true)
	pw.ShowTime(true)
	pw.ShowTracker(true)
	pw.ShowValue(true)
	pw.SetMessageWidth(24)
	pw.SetNumTrackersExpected(*numTrackers)
	pw.SetSortBy(progress.SortByPercentDsc)
	pw.SetStyle(progress.StyleDefault)
	pw.SetTrackerPosition(progress.PositionRight)
	pw.SetUpdateFrequency(time.Millisecond * 100)
	pw.Style().Colors = progress.StyleColorsExample
	pw.Style().Options.PercentFormat = "%4.1f%%"

	// call Render() in async mode; yes we don't have any trackers at the moment
	go pw.Render()

	// add a bunch of trackers with random parameters to demo most of the
	// features available; do this in async too like a client might do (for ex.
	// when downloading a bunch of files in parallel)
	for idx := int64(1); idx <= int64(*numTrackers); idx++ {
		go trackSomething(pw, idx, idx == int64(*numTrackers))

		// in auto-stop mode, the Render logic terminates the moment it detects
		// zero active trackers; but in a manual-stop mode, it keeps waiting and
		// is a good chance to demo trackers being added dynamically while other
		// trackers are active or done
		if !*autoStop {
			time.Sleep(time.Millisecond * 100)
		}
	}

	// wait for one or more trackers to become active (just blind-wait for a
	// second) and then keep watching until Rendering is in progress
	time.Sleep(time.Second)
	for pw.IsRenderInProgress() {
		// for manual-stop mode, stop when there are no more active trackers
		if !*autoStop && pw.LengthActive() == 0 {
			pw.Stop()
		}
		time.Sleep(time.Millisecond * 100)
	}

	fmt.Println("\nAll done!")
}
