package progress

import (
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type outputWriter struct {
	Text strings.Builder
}

func (rc *outputWriter) Write(p []byte) (n int, err error) {
	return rc.Text.Write(p)
}

func (rc *outputWriter) String() string {
	return rc.Text.String()
}

func generateWriter() Writer {
	pw := NewWriter()
	pw.SetAutoStop(false)
	pw.SetNumTrackersExpected(1)
	pw.SetSortBy(SortByNone)
	pw.SetStyle(StyleDefault)
	pw.SetTrackerLength(25)
	pw.SetTrackerPosition(PositionRight)
	pw.SetUpdateFrequency(time.Millisecond * 50)
	pw.ShowOverallTracker(false)
	pw.ShowPercentage(true)
	pw.ShowTime(true)
	pw.ShowTracker(true)
	pw.ShowValue(true)
	pw.Style().Colors = StyleColors{}
	pw.Style().Options = StyleOptionsDefault
	return pw
}

func trackSomething(pw Writer, tracker *Tracker) {
	pw.AppendTracker(tracker)

	incrementPerCycle := tracker.Total / 3

	c := time.Tick(time.Millisecond * 100)
	for !tracker.IsDone() {
		select {
		case <-c:
			if tracker.value+incrementPerCycle > tracker.Total {
				tracker.Increment(tracker.Total - tracker.value)
			} else {
				tracker.Increment(incrementPerCycle)
			}
		}
	}
}

func renderAndWait(pw Writer, autoStop bool) {
	go pw.Render()
	time.Sleep(time.Millisecond * 100)
	for pw.IsRenderInProgress() {
		if pw.LengthActive() == 0 {
			break
		}
		time.Sleep(time.Millisecond * 100)
	}
	if !autoStop {
		pw.Stop()
	}
}

func TestProgress_RenderNothing(t *testing.T) {
	renderOutput := outputWriter{}

	pw := generateWriter()
	pw.SetOutputWriter(&renderOutput)

	go pw.Render()
	time.Sleep(time.Second)
	pw.Stop()
	time.Sleep(time.Second)

	assert.Empty(t, renderOutput.String())
}

func TestProgress_RenderSomeTrackers_OnLeftSide(t *testing.T) {
	renderOutput := outputWriter{}

	pw := generateWriter()
	pw.SetOutputWriter(&renderOutput)
	pw.SetTrackerPosition(PositionLeft)
	go trackSomething(pw, &Tracker{Message: "Calculation Total   # 1", Total: 1000, Units: UnitsDefault})
	go trackSomething(pw, &Tracker{Message: "Downloading File    # 2", Total: 1000, Units: UnitsBytes})
	go trackSomething(pw, &Tracker{Message: "Transferring Amount # 3", Total: 1000, Units: UnitsCurrencyDollar})
	renderAndWait(pw, false)

	expectedOutPatterns := []*regexp.Regexp{
		regexp.MustCompile(`\x1b\[K\d+\.\d+% \[[#.]{23}] \[\d+ in [\d.]+ms] \.\.\. Calculation Total   # 1`),
		regexp.MustCompile(`\x1b\[K\d+\.\d+% \[[#.]{23}] \[\d+B in [\d.]+ms] \.\.\. Downloading File    # 2`),
		regexp.MustCompile(`\x1b\[K\d+\.\d+% \[[#.]{23}] \[\$\d+ in [\d.]+ms] \.\.\. Transferring Amount # 3`),
		regexp.MustCompile(`\x1b\[KCalculation Total   # 1 \.\.\. done! \[\d+\.\d+K in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KDownloading File    # 2 \.\.\. done! \[\d+\.\d+KB in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KTransferring Amount # 3 \.\.\. done! \[\$\d+\.\d+K in [\d.]+ms]`),
	}
	out := renderOutput.String()
	for _, expectedOutPattern := range expectedOutPatterns {
		if !expectedOutPattern.MatchString(out) {
			assert.Fail(t, "Failed to find a pattern in the Output.", expectedOutPattern.String())
		}
	}
}

func TestProgress_RenderSomeTrackers_OnRightSide(t *testing.T) {
	renderOutput := outputWriter{}

	pw := generateWriter()
	pw.SetOutputWriter(&renderOutput)
	pw.SetTrackerPosition(PositionRight)
	go trackSomething(pw, &Tracker{Message: "Calculation Total   # 1", Total: 1000, Units: UnitsDefault})
	go trackSomething(pw, &Tracker{Message: "Downloading File    # 2", Total: 1000, Units: UnitsBytes})
	go trackSomething(pw, &Tracker{Message: "Transferring Amount # 3", Total: 1000, Units: UnitsCurrencyDollar})
	renderAndWait(pw, false)

	expectedOutPatterns := []*regexp.Regexp{
		regexp.MustCompile(`\x1b\[KCalculation Total   # 1 \.\.\. \d+\.\d+% \[[#.]{23}] \[\d+ in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KDownloading File    # 2 \.\.\. \d+\.\d+% \[[#.]{23}] \[\d+B in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KTransferring Amount # 3 \.\.\. \d+\.\d+% \[[#.]{23}] \[\$\d+ in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KCalculation Total   # 1 \.\.\. done! \[\d+\.\d+K in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KDownloading File    # 2 \.\.\. done! \[\d+\.\d+KB in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KTransferring Amount # 3 \.\.\. done! \[\$\d+\.\d+K in [\d.]+ms]`),
	}
	out := renderOutput.String()
	for _, expectedOutPattern := range expectedOutPatterns {
		if !expectedOutPattern.MatchString(out) {
			assert.Fail(t, "Failed to find a pattern in the Output.", expectedOutPattern.String())
		}
	}
}

func TestProgress_RenderSomeTrackers_WithAutoStop(t *testing.T) {
	renderOutput := outputWriter{}

	pw := generateWriter()
	pw.SetAutoStop(true)
	pw.SetOutputWriter(&renderOutput)
	pw.SetTrackerPosition(PositionRight)
	go trackSomething(pw, &Tracker{Message: "Calculation Total   # 1", Total: 1000, Units: UnitsDefault})
	go trackSomething(pw, &Tracker{Message: "Downloading File    # 2", Total: 1000, Units: UnitsBytes})
	go trackSomething(pw, &Tracker{Message: "Transferring Amount # 3", Total: 1000, Units: UnitsCurrencyDollar})
	renderAndWait(pw, true)

	expectedOutPatterns := []*regexp.Regexp{
		regexp.MustCompile(`\x1b\[KCalculation Total   # 1 \.\.\. \d+\.\d+% \[[#.]{23}] \[\d+ in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KDownloading File    # 2 \.\.\. \d+\.\d+% \[[#.]{23}] \[\d+B in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KTransferring Amount # 3 \.\.\. \d+\.\d+% \[[#.]{23}] \[\$\d+ in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KCalculation Total   # 1 \.\.\. done! \[\d+\.\d+K in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KDownloading File    # 2 \.\.\. done! \[\d+\.\d+KB in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KTransferring Amount # 3 \.\.\. done! \[\$\d+\.\d+K in [\d.]+ms]`),
	}
	out := renderOutput.String()
	for _, expectedOutPattern := range expectedOutPatterns {
		if !expectedOutPattern.MatchString(out) {
			assert.Fail(t, "Failed to find a pattern in the Output.", expectedOutPattern.String())
		}
	}
}

func TestProgress_RenderSomeTrackers_WithLineWidth1(t *testing.T) {
	renderOutput := outputWriter{}

	pw := generateWriter()
	pw.SetMessageWidth(5)
	pw.SetOutputWriter(&renderOutput)
	pw.SetTrackerPosition(PositionRight)
	go trackSomething(pw, &Tracker{Message: "Calculation Total   # 1", Total: 1000, Units: UnitsDefault})
	go trackSomething(pw, &Tracker{Message: "Downloading File    # 2", Total: 1000, Units: UnitsBytes})
	go trackSomething(pw, &Tracker{Message: "Transferring Amount # 3", Total: 1000, Units: UnitsCurrencyDollar})
	renderAndWait(pw, false)

	expectedOutPatterns := []*regexp.Regexp{
		regexp.MustCompile(`\x1b\[KCalc~ \.\.\. \d+\.\d+% \[[#.]{23}] \[\d+ in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KDown~ \.\.\. \d+\.\d+% \[[#.]{23}] \[\d+B in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KTran~ \.\.\. \d+\.\d+% \[[#.]{23}] \[\$\d+ in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KCalc~ \.\.\. done! \[\d+\.\d+K in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KDown~ \.\.\. done! \[\d+\.\d+KB in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KTran~ \.\.\. done! \[\$\d+\.\d+K in [\d.]+ms]`),
	}
	out := renderOutput.String()
	for _, expectedOutPattern := range expectedOutPatterns {
		if !expectedOutPattern.MatchString(out) {
			assert.Fail(t, "Failed to find a pattern in the Output.", expectedOutPattern.String())
		}
	}
}

func TestProgress_RenderSomeTrackers_WithLineWidth2(t *testing.T) {
	renderOutput := outputWriter{}

	pw := generateWriter()
	pw.SetMessageWidth(50)
	pw.SetOutputWriter(&renderOutput)
	pw.SetTrackerPosition(PositionRight)
	go trackSomething(pw, &Tracker{Message: "Calculation Total   # 1", Total: 1000, Units: UnitsDefault})
	go trackSomething(pw, &Tracker{Message: "Downloading File    # 2", Total: 1000, Units: UnitsBytes})
	go trackSomething(pw, &Tracker{Message: "Transferring Amount # 3", Total: 1000, Units: UnitsCurrencyDollar})
	renderAndWait(pw, false)

	expectedOutPatterns := []*regexp.Regexp{
		regexp.MustCompile(`\x1b\[KCalculation Total   # 1\s{28}\.\.\. \d+\.\d+% \[[#.]{23}] \[\d+ in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KDownloading File    # 2\s{28}\.\.\. \d+\.\d+% \[[#.]{23}] \[\d+B in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KTransferring Amount # 3\s{28}\.\.\. \d+\.\d+% \[[#.]{23}] \[\$\d+ in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KCalculation Total   # 1\s{28}\.\.\. done! \[\d+\.\d+K in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KDownloading File    # 2\s{28}\.\.\. done! \[\d+\.\d+KB in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KTransferring Amount # 3\s{28}\.\.\. done! \[\$\d+\.\d+K in [\d.]+ms]`),
	}
	out := renderOutput.String()
	for _, expectedOutPattern := range expectedOutPatterns {
		if !expectedOutPattern.MatchString(out) {
			assert.Fail(t, "Failed to find a pattern in the Output.", expectedOutPattern.String())
		}
	}
}

func TestProgress_RenderSomeTrackers_WithOverallTracker(t *testing.T) {
	renderOutput := outputWriter{}

	pw := generateWriter()
	pw.SetOutputWriter(&renderOutput)
	pw.SetTrackerPosition(PositionRight)
	pw.ShowOverallTracker(true)
	pw.Style().Options.TimeOverallPrecision = time.Millisecond
	go trackSomething(pw, &Tracker{Message: "Calculation Total   # 1", Total: 1000, Units: UnitsDefault})
	go trackSomething(pw, &Tracker{Message: "Downloading File    # 2", Total: 1000, Units: UnitsBytes})
	go trackSomething(pw, &Tracker{Message: "Transferring Amount # 3", Total: 1000, Units: UnitsCurrencyDollar})
	renderAndWait(pw, false)

	expectedOutPatterns := []*regexp.Regexp{
		regexp.MustCompile(`\x1b\[KCalculation Total   # 1 \.\.\. \d+\.\d+% \[[#.]{23}] \[\d+ in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KDownloading File    # 2 \.\.\. \d+\.\d+% \[[#.]{23}] \[\d+B in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KTransferring Amount # 3 \.\.\. \d+\.\d+% \[[#.]{23}] \[\$\d+ in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KCalculation Total   # 1 \.\.\. done! \[\d+\.\d+K in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KDownloading File    # 2 \.\.\. done! \[\d+\.\d+KB in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KTransferring Amount # 3 \.\.\. done! \[\$\d+\.\d+K in [\d.]+ms]`),
		regexp.MustCompile(`\[[\d.ms]+; ~ETA: [\d.ms]+`),
	}
	out := renderOutput.String()
	for _, expectedOutPattern := range expectedOutPatterns {
		if !expectedOutPattern.MatchString(out) {
			assert.Fail(t, "Failed to find a pattern in the Output.", expectedOutPattern.String())
		}
	}
}
