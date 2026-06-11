package progress

import (
	"io"
	"testing"
	"time"
)

func BenchmarkProgress_Render(b *testing.B) {
	track := func(pw Writer, tracker *Tracker) {
		tracker.Reset()
		pw.AppendTracker(tracker)
		parts := 4
		for i := 0; i < parts; i++ {
			tracker.Increment(tracker.Total / int64(parts))
		}
	}
	trackers := []*Tracker{
		{Message: "Calculating Total   # 1", Total: 1000, Units: UnitsDefault},
		{Message: "Downloading File    # 2", Total: 1000, Units: UnitsBytes},
		{Message: "Transferring Amount # 3", Total: 1000, Units: UnitsCurrencyDollar},
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		pw := NewWriter()
		pw.SetAutoStop(true)
		pw.SetOutputWriter(io.Discard)
		// set a very short update frequency for faster benchmark execution
		pw.SetUpdateFrequency(time.Millisecond)
		for _, tracker := range trackers {
			go track(pw, tracker)
		}
		// render once to test rendering performance
		pw.Render()
	}
}
