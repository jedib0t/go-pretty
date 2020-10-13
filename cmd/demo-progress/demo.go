package main

import (
	"github.com/jedib0t/go-pretty/v6/progress"
	"time"
)

func main() {
	pw := progress.NewWriter()
	pw.ShowOverallTracker(false)
	pw.ShowETA(true)
	pw.ShowTime(true)
	pw.ShowTracker(true)
	pw.ShowValue(true)
	pw.SetSortBy(progress.SortByPercentDsc)
	pw.SetStyle(progress.StyleDefault)
	pw.SetTrackerPosition(progress.PositionRight)
	pw.SetUpdateFrequency(time.Millisecond * 50)
	pw.Style().Colors = progress.StyleColorsExample
	defer pw.Stop()

	tracker := progress.Tracker{Message: "Downloading", Total: 1024 * 1024 * 1024, Units: progress.UnitsBytes}
	pw.AppendTracker(&tracker)
	go pw.Render()

	read := int64(0)
	for {
		read += 10240
		tracker.SetValue(read)
		time.Sleep(5 * time.Microsecond)
		if read >= 1024*1024*1024 {
			break
		}
	}

	tracker.MarkAsDone()
	time.Sleep(100 * time.Millisecond)
}
