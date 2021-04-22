package progress

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type outputWriter struct {
	Text strings.Builder
}

func (w *outputWriter) Write(p []byte) (n int, err error) {
	return w.Text.Write(p)
}

func (w *outputWriter) String() string {
	return w.Text.String()
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
	incrementPerCycle := tracker.Total / 3

	pw.AppendTracker(tracker)

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

func trackSomethingErrored(pw Writer, tracker *Tracker) {
	incrementPerCycle := tracker.Total / 3
	total := tracker.Total
	tracker.Total = 0

	pw.AppendTracker(tracker)

	c := time.Tick(time.Millisecond * 100)
	for !tracker.IsDone() {
		select {
		case <-c:
			if tracker.value+incrementPerCycle > total {
				tracker.MarkAsErrored()
			} else {
				tracker.IncrementWithError(incrementPerCycle)
			}
		}
	}
}

func trackSomethingIndeterminate(pw Writer, tracker *Tracker) {
	incrementPerCycle := tracker.Total / 3
	total := tracker.Total
	tracker.Total = 0

	pw.AppendTracker(tracker)

	c := time.Tick(time.Millisecond * 100)
	for !tracker.IsDone() {
		select {
		case <-c:
			if tracker.value+incrementPerCycle > total {
				tracker.Increment(total - tracker.value)
			} else {
				tracker.Increment(incrementPerCycle)
			}
			if tracker.Value() >= total {
				tracker.MarkAsDone()
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

func showOutputOnFailure(t *testing.T, out string) {
	if t.Failed() {
		lines := strings.Split(out, "\n")
		sort.Strings(lines)
		for _, line := range lines {
			fmt.Printf("%#v,\n", line)
		}
	}
}

func TestProgress_generateTrackerStr(t *testing.T) {
	pw := Progress{}
	pw.Style().Chars = StyleChars{
		BoxLeft:    "",
		BoxRight:   "",
		Finished:   "#",
		Finished25: "1",
		Finished50: "2",
		Finished75: "3",
		Unfinished: ".",
	}

	expectedTrackerStrMap := map[int64]string{
		0:   "..........",
		1:   "..........",
		2:   "..........",
		3:   "1.........",
		4:   "1.........",
		5:   "2.........",
		6:   "2.........",
		7:   "2.........",
		8:   "3.........",
		9:   "3.........",
		10:  "#.........",
		11:  "#.........",
		12:  "#.........",
		13:  "#1........",
		14:  "#1........",
		15:  "#2........",
		16:  "#2........",
		17:  "#2........",
		18:  "#3........",
		19:  "#3........",
		20:  "##........",
		21:  "##........",
		22:  "##........",
		23:  "##1.......",
		24:  "##1.......",
		25:  "##2.......",
		26:  "##2.......",
		27:  "##2.......",
		28:  "##3.......",
		29:  "##3.......",
		30:  "###.......",
		31:  "###.......",
		32:  "###.......",
		33:  "###1......",
		34:  "###1......",
		35:  "###2......",
		36:  "###2......",
		37:  "###2......",
		38:  "###3......",
		39:  "###3......",
		40:  "####......",
		41:  "####......",
		42:  "####......",
		43:  "####1.....",
		44:  "####1.....",
		45:  "####2.....",
		46:  "####2.....",
		47:  "####2.....",
		48:  "####3.....",
		49:  "####3.....",
		50:  "#####.....",
		51:  "#####.....",
		52:  "#####.....",
		53:  "#####1....",
		54:  "#####1....",
		55:  "#####2....",
		56:  "#####2....",
		57:  "#####2....",
		58:  "#####3....",
		59:  "#####3....",
		60:  "######....",
		61:  "######....",
		62:  "######....",
		63:  "######1...",
		64:  "######1...",
		65:  "######2...",
		66:  "######2...",
		67:  "######2...",
		68:  "######3...",
		69:  "######3...",
		70:  "#######...",
		71:  "#######...",
		72:  "#######...",
		73:  "#######1..",
		74:  "#######1..",
		75:  "#######2..",
		76:  "#######2..",
		77:  "#######2..",
		78:  "#######3..",
		79:  "#######3..",
		80:  "########..",
		81:  "########..",
		82:  "########..",
		83:  "########1.",
		84:  "########1.",
		85:  "########2.",
		86:  "########2.",
		87:  "########2.",
		88:  "########3.",
		89:  "########3.",
		90:  "#########.",
		91:  "#########.",
		92:  "#########.",
		93:  "#########1",
		94:  "#########1",
		95:  "#########2",
		96:  "#########2",
		97:  "#########2",
		98:  "#########3",
		99:  "#########3",
		100: "##########",
	}

	finalOutput := strings.Builder{}
	tr := Tracker{Total: 100}
	for value := int64(0); value <= 100; value++ {
		tr.value = value
		actualStr := pw.generateTrackerStr(&tr, 10, renderHint{})
		if expectedStr, ok := expectedTrackerStrMap[value]; ok {
			assert.Equal(t, expectedStr, actualStr, "value=%d", value)
		}
		finalOutput.WriteString(fmt.Sprintf(" %d: \"%s\",\n", value, actualStr))
	}
	if t.Failed() {
		fmt.Println(finalOutput.String())
	}
}

func TestProgress_generateTrackerStr_Indeterminate(t *testing.T) {
	pw := Progress{}
	pw.Style().Chars = StyleChars{
		BoxLeft:       "",
		BoxRight:      "",
		Finished:      "#",
		Finished25:    "1",
		Finished50:    "2",
		Finished75:    "3",
		Indeterminate: indeterminateIndicatorMovingBackAndForth("<=>"),
		Unfinished:    ".",
	}

	expectedTrackerStrMap := map[int64]string{
		0:   "<=>.......",
		1:   ".<=>......",
		2:   "..<=>.....",
		3:   "...<=>....",
		4:   "....<=>...",
		5:   ".....<=>..",
		6:   "......<=>.",
		7:   ".......<=>",
		8:   "......<=>.",
		9:   ".....<=>..",
		10:  "....<=>...",
		11:  "...<=>....",
		12:  "..<=>.....",
		13:  ".<=>......",
		14:  "<=>.......",
		15:  ".<=>......",
		16:  "..<=>.....",
		17:  "...<=>....",
		18:  "....<=>...",
		19:  ".....<=>..",
		20:  "......<=>.",
		21:  ".......<=>",
		22:  "......<=>.",
		23:  ".....<=>..",
		24:  "....<=>...",
		25:  "...<=>....",
		26:  "..<=>.....",
		27:  ".<=>......",
		28:  "<=>.......",
		29:  ".<=>......",
		30:  "..<=>.....",
		31:  "...<=>....",
		32:  "....<=>...",
		33:  ".....<=>..",
		34:  "......<=>.",
		35:  ".......<=>",
		36:  "......<=>.",
		37:  ".....<=>..",
		38:  "....<=>...",
		39:  "...<=>....",
		40:  "..<=>.....",
		41:  ".<=>......",
		42:  "<=>.......",
		43:  ".<=>......",
		44:  "..<=>.....",
		45:  "...<=>....",
		46:  "....<=>...",
		47:  ".....<=>..",
		48:  "......<=>.",
		49:  ".......<=>",
		50:  "......<=>.",
		51:  ".....<=>..",
		52:  "....<=>...",
		53:  "...<=>....",
		54:  "..<=>.....",
		55:  ".<=>......",
		56:  "<=>.......",
		57:  ".<=>......",
		58:  "..<=>.....",
		59:  "...<=>....",
		60:  "....<=>...",
		61:  ".....<=>..",
		62:  "......<=>.",
		63:  ".......<=>",
		64:  "......<=>.",
		65:  ".....<=>..",
		66:  "....<=>...",
		67:  "...<=>....",
		68:  "..<=>.....",
		69:  ".<=>......",
		70:  "<=>.......",
		71:  ".<=>......",
		72:  "..<=>.....",
		73:  "...<=>....",
		74:  "....<=>...",
		75:  ".....<=>..",
		76:  "......<=>.",
		77:  ".......<=>",
		78:  "......<=>.",
		79:  ".....<=>..",
		80:  "....<=>...",
		81:  "...<=>....",
		82:  "..<=>.....",
		83:  ".<=>......",
		84:  "<=>.......",
		85:  ".<=>......",
		86:  "..<=>.....",
		87:  "...<=>....",
		88:  "....<=>...",
		89:  ".....<=>..",
		90:  "......<=>.",
		91:  ".......<=>",
		92:  "......<=>.",
		93:  ".....<=>..",
		94:  "....<=>...",
		95:  "...<=>....",
		96:  "..<=>.....",
		97:  ".<=>......",
		98:  "<=>.......",
		99:  ".<=>......",
		100: "..<=>.....",
	}

	finalOutput := strings.Builder{}
	tr := Tracker{Total: 0}
	for value := int64(0); value <= 100; value++ {
		tr.value = value
		actualStr := pw.generateTrackerStr(&tr, 10, renderHint{})
		if expectedStr, ok := expectedTrackerStrMap[value]; ok {
			assert.Equal(t, expectedStr, actualStr, "value=%d", value)
		}
		finalOutput.WriteString(fmt.Sprintf(" %d: \"%s\",\n", value, actualStr))
	}
	if t.Failed() {
		fmt.Println(finalOutput.String())
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
	go trackSomething(pw, &Tracker{Message: "Calculating Total   # 1\r", Total: 1000, Units: UnitsDefault})
	go trackSomething(pw, &Tracker{Message: "Downloading File\t# 2", Total: 1000, Units: UnitsBytes})
	go trackSomething(pw, &Tracker{Message: "Transferring Amount # 3", Total: 1000, Units: UnitsCurrencyDollar})
	renderAndWait(pw, false)

	expectedOutPatterns := []*regexp.Regexp{
		regexp.MustCompile(`\x1b\[K\d+\.\d+% \[[#.]{23}] \[\d+ in [\d.]+ms] \.\.\. Calculating Total   # 1`),
		regexp.MustCompile(`\x1b\[K\d+\.\d+% \[[#.]{23}] \[\d+B in [\d.]+ms] \.\.\. Downloading File    # 2`),
		regexp.MustCompile(`\x1b\[K\d+\.\d+% \[[#.]{23}] \[\$\d+ in [\d.]+ms] \.\.\. Transferring Amount # 3`),
		regexp.MustCompile(`\x1b\[KCalculating Total   # 1 \.\.\. done! \[\d+\.\d+K in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KDownloading File    # 2 \.\.\. done! \[\d+\.\d+KB in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KTransferring Amount # 3 \.\.\. done! \[\$\d+\.\d+K in [\d.]+ms]`),
	}
	out := renderOutput.String()
	for _, expectedOutPattern := range expectedOutPatterns {
		if !expectedOutPattern.MatchString(out) {
			assert.Fail(t, "Failed to find a pattern in the Output.", expectedOutPattern.String())
		}
	}
	showOutputOnFailure(t, out)
}

func TestProgress_RenderSomeTrackers_OnRightSide(t *testing.T) {
	renderOutput := outputWriter{}

	pw := generateWriter()
	pw.SetOutputWriter(&renderOutput)
	pw.SetTrackerPosition(PositionRight)
	go trackSomething(pw, &Tracker{Message: "Calculating Total   # 1\r", Total: 1000, Units: UnitsDefault})
	go trackSomething(pw, &Tracker{Message: "Downloading File\t# 2", Total: 1000, Units: UnitsBytes})
	go trackSomething(pw, &Tracker{Message: "Transferring Amount # 3", Total: 1000, Units: UnitsCurrencyDollar})
	renderAndWait(pw, false)

	expectedOutPatterns := []*regexp.Regexp{
		regexp.MustCompile(`\x1b\[KCalculating Total   # 1 \.\.\. \d+\.\d+% \[[#.]{23}] \[\d+ in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KDownloading File    # 2 \.\.\. \d+\.\d+% \[[#.]{23}] \[\d+B in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KTransferring Amount # 3 \.\.\. \d+\.\d+% \[[#.]{23}] \[\$\d+ in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KCalculating Total   # 1 \.\.\. done! \[\d+\.\d+K in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KDownloading File    # 2 \.\.\. done! \[\d+\.\d+KB in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KTransferring Amount # 3 \.\.\. done! \[\$\d+\.\d+K in [\d.]+ms]`),
	}
	out := renderOutput.String()
	for _, expectedOutPattern := range expectedOutPatterns {
		if !expectedOutPattern.MatchString(out) {
			assert.Fail(t, "Failed to find a pattern in the Output.", expectedOutPattern.String())
		}
	}
	showOutputOnFailure(t, out)
}

func TestProgress_RenderSomeTrackers_WithAutoStop(t *testing.T) {
	renderOutput := outputWriter{}

	pw := generateWriter()
	pw.SetAutoStop(true)
	pw.SetOutputWriter(&renderOutput)
	pw.SetTrackerPosition(PositionRight)
	go trackSomething(pw, &Tracker{Message: "Calculating Total   # 1\r", Total: 1000, Units: UnitsDefault})
	go trackSomething(pw, &Tracker{Message: "Downloading File\t# 2", Total: 1000, Units: UnitsBytes})
	go trackSomething(pw, &Tracker{Message: "Transferring Amount # 3", Total: 1000, Units: UnitsCurrencyDollar})
	renderAndWait(pw, true)

	expectedOutPatterns := []*regexp.Regexp{
		regexp.MustCompile(`\x1b\[KCalculating Total   # 1 \.\.\. \d+\.\d+% \[[#.]{23}] \[\d+ in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KDownloading File    # 2 \.\.\. \d+\.\d+% \[[#.]{23}] \[\d+B in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KTransferring Amount # 3 \.\.\. \d+\.\d+% \[[#.]{23}] \[\$\d+ in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KCalculating Total   # 1 \.\.\. done! \[\d+\.\d+K in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KDownloading File    # 2 \.\.\. done! \[\d+\.\d+KB in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KTransferring Amount # 3 \.\.\. done! \[\$\d+\.\d+K in [\d.]+ms]`),
	}
	out := renderOutput.String()
	for _, expectedOutPattern := range expectedOutPatterns {
		if !expectedOutPattern.MatchString(out) {
			assert.Fail(t, "Failed to find a pattern in the Output.", expectedOutPattern.String())
		}
	}
	showOutputOnFailure(t, out)
}

func TestProgress_RenderSomeTrackers_WithError(t *testing.T) {
	renderOutput := outputWriter{}

	pw := generateWriter()
	pw.SetOutputWriter(&renderOutput)
	go trackSomething(pw, &Tracker{Message: "Calculating Total   # 1\r", Total: 1000, Units: UnitsDefault})
	go trackSomething(pw, &Tracker{Message: "Downloading File\t# 2", Total: 1000, Units: UnitsBytes})
	go trackSomethingErrored(pw, &Tracker{Message: "Transferring Amount # 3", Total: 1000, Units: UnitsCurrencyDollar})
	renderAndWait(pw, false)

	expectedOutPatterns := []*regexp.Regexp{
		regexp.MustCompile(`\x1b\[KCalculating Total   # 1 \.\.\. \d+\.\d+% \[[#.]{23}] \[\d+ in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KDownloading File    # 2 \.\.\. \d+\.\d+% \[[#.]{23}] \[\d+B in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KTransferring Amount # 3 \.\.\.  \?\?\?  \[[<#>.]{23}] \[\$\d+ in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KCalculating Total   # 1 \.\.\. done! \[\d+\.\d+K in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KDownloading File    # 2 \.\.\. done! \[\d+\.\d+KB in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KTransferring Amount # 3 \.\.\. fail! \[\$\d+ in [\d.]+ms]`),
	}
	out := renderOutput.String()
	for _, expectedOutPattern := range expectedOutPatterns {
		if !expectedOutPattern.MatchString(out) {
			assert.Fail(t, "Failed to find a pattern in the Output.", expectedOutPattern.String())
		}
	}
	showOutputOnFailure(t, out)
}

func TestProgress_RenderSomeTrackers_WithIndeterminateTracker(t *testing.T) {
	renderOutput := outputWriter{}

	pw := generateWriter()
	pw.SetOutputWriter(&renderOutput)
	go trackSomething(pw, &Tracker{Message: "Calculating Total   # 1\r", Total: 1000, Units: UnitsDefault})
	go trackSomething(pw, &Tracker{Message: "Downloading File\t# 2", Total: 1000, Units: UnitsBytes})
	go trackSomethingIndeterminate(pw, &Tracker{Message: "Transferring Amount # 3", Total: 1000, Units: UnitsCurrencyDollar})
	renderAndWait(pw, false)

	expectedOutPatterns := []*regexp.Regexp{
		regexp.MustCompile(`\x1b\[KCalculating Total   # 1 \.\.\. \d+\.\d+% \[[#.]{23}] \[\d+ in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KDownloading File    # 2 \.\.\. \d+\.\d+% \[[#.]{23}] \[\d+B in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KTransferring Amount # 3 \.\.\.  \?\?\?  \[[<#>.]{23}] \[\$\d+ in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KCalculating Total   # 1 \.\.\. done! \[\d+\.\d+K in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KDownloading File    # 2 \.\.\. done! \[\d+\.\d+KB in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KTransferring Amount # 3 \.\.\. done! \[\$\d+\.\d+K in [\d.]+ms]`),
	}
	out := renderOutput.String()
	for _, expectedOutPattern := range expectedOutPatterns {
		if !expectedOutPattern.MatchString(out) {
			assert.Fail(t, "Failed to find a pattern in the Output.", expectedOutPattern.String())
		}
	}
	showOutputOnFailure(t, out)
}

func TestProgress_RenderSomeTrackers_WithLineWidth1(t *testing.T) {
	renderOutput := outputWriter{}

	pw := generateWriter()
	pw.SetMessageWidth(5)
	pw.SetOutputWriter(&renderOutput)
	pw.SetTrackerPosition(PositionRight)
	go trackSomething(pw, &Tracker{Message: "Calculating Total   # 1\r", Total: 1000, Units: UnitsDefault})
	go trackSomething(pw, &Tracker{Message: "Downloading File\t# 2", Total: 1000, Units: UnitsBytes})
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
	showOutputOnFailure(t, out)
}

func TestProgress_RenderSomeTrackers_WithLineWidth2(t *testing.T) {
	renderOutput := outputWriter{}

	pw := generateWriter()
	pw.SetMessageWidth(50)
	pw.SetOutputWriter(&renderOutput)
	pw.SetTrackerPosition(PositionRight)
	go trackSomething(pw, &Tracker{Message: "Calculating Total   # 1\r", Total: 1000, Units: UnitsDefault})
	go trackSomething(pw, &Tracker{Message: "Downloading File\t# 2", Total: 1000, Units: UnitsBytes})
	go trackSomething(pw, &Tracker{Message: "Transferring Amount # 3", Total: 1000, Units: UnitsCurrencyDollar})
	renderAndWait(pw, false)

	expectedOutPatterns := []*regexp.Regexp{
		regexp.MustCompile(`\x1b\[KCalculating Total   # 1\s{28}\.\.\. \d+\.\d+% \[[#.]{23}] \[\d+ in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KDownloading File    # 2\s{28}\.\.\. \d+\.\d+% \[[#.]{23}] \[\d+B in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KTransferring Amount # 3\s{28}\.\.\. \d+\.\d+% \[[#.]{23}] \[\$\d+ in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KCalculating Total   # 1\s{28}\.\.\. done! \[\d+\.\d+K in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KDownloading File    # 2\s{28}\.\.\. done! \[\d+\.\d+KB in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KTransferring Amount # 3\s{28}\.\.\. done! \[\$\d+\.\d+K in [\d.]+ms]`),
	}
	out := renderOutput.String()
	for _, expectedOutPattern := range expectedOutPatterns {
		if !expectedOutPattern.MatchString(out) {
			assert.Fail(t, "Failed to find a pattern in the Output.", expectedOutPattern.String())
		}
	}
	showOutputOnFailure(t, out)
}

func TestProgress_RenderSomeTrackers_WithOverallTracker(t *testing.T) {
	renderOutput := outputWriter{}

	pw := generateWriter()
	pw.SetOutputWriter(&renderOutput)
	pw.SetTrackerPosition(PositionRight)
	pw.ShowOverallTracker(true)
	pw.Style().Options.TimeOverallPrecision = time.Millisecond
	go trackSomething(pw, &Tracker{Message: "Calculating Total   # 1\r", Total: 1000, Units: UnitsDefault})
	go trackSomething(pw, &Tracker{Message: "Downloading File\t# 2", Total: 1000, Units: UnitsBytes})
	go trackSomething(pw, &Tracker{Message: "Transferring Amount # 3", Total: 1000, Units: UnitsCurrencyDollar})
	renderAndWait(pw, false)

	expectedOutPatterns := []*regexp.Regexp{
		regexp.MustCompile(`\x1b\[KCalculating Total   # 1 \.\.\. \d+\.\d+% \[[#.]{23}] \[\d+ in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KDownloading File    # 2 \.\.\. \d+\.\d+% \[[#.]{23}] \[\d+B in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KTransferring Amount # 3 \.\.\. \d+\.\d+% \[[#.]{23}] \[\$\d+ in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KCalculating Total   # 1 \.\.\. done! \[\d+\.\d+K in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KDownloading File    # 2 \.\.\. done! \[\d+\.\d+KB in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KTransferring Amount # 3 \.\.\. done! \[\$\d+\.\d+K in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[K\[[.#]+] \[[\d.ms]+; ~ETA: [\d.ms]+`),
	}
	out := renderOutput.String()
	for _, expectedOutPattern := range expectedOutPatterns {
		if !expectedOutPattern.MatchString(out) {
			assert.Fail(t, "Failed to find a pattern in the Output.", expectedOutPattern.String())
		}
	}
	showOutputOnFailure(t, out)
}

func TestProgress_RenderSomeTrackers_WithoutOverallTracker_WithETA(t *testing.T) {
	renderOutput := outputWriter{}

	pw := generateWriter()
	pw.SetOutputWriter(&renderOutput)
	pw.SetTrackerPosition(PositionRight)
	pw.ShowETA(true)
	pw.ShowOverallTracker(false)
	pw.Style().Options.ETAPrecision = time.Millisecond
	go trackSomething(pw, &Tracker{Message: "Calculating Total   # 1\r", Total: 1000, Units: UnitsDefault})
	go trackSomething(pw, &Tracker{Message: "Downloading File\t# 2", Total: 1000, Units: UnitsBytes})
	go trackSomething(pw, &Tracker{Message: "Transferring Amount # 3", Total: 1000, Units: UnitsCurrencyDollar})
	renderAndWait(pw, false)

	expectedOutPatterns := []*regexp.Regexp{
		regexp.MustCompile(`\x1b\[KCalculating Total   # 1 \.\.\. \d+\.\d+% \[[#.]{23}] \[\d+ in [\d.]+ms; ~ETA: [\d]+ms]`),
		regexp.MustCompile(`\x1b\[KDownloading File    # 2 \.\.\. \d+\.\d+% \[[#.]{23}] \[\d+B in [\d.]+ms; ~ETA: [\d]+ms]`),
		regexp.MustCompile(`\x1b\[KTransferring Amount # 3 \.\.\. \d+\.\d+% \[[#.]{23}] \[\$\d+ in [\d.]+ms; ~ETA: [\d]+ms]`),
		regexp.MustCompile(`\x1b\[KCalculating Total   # 1 \.\.\. done! \[\d+\.\d+K in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KDownloading File    # 2 \.\.\. done! \[\d+\.\d+KB in [\d.]+ms]`),
		regexp.MustCompile(`\x1b\[KTransferring Amount # 3 \.\.\. done! \[\$\d+\.\d+K in [\d.]+ms]`),
	}
	out := renderOutput.String()
	for _, expectedOutPattern := range expectedOutPatterns {
		if !expectedOutPattern.MatchString(out) {
			assert.Fail(t, "Failed to find a pattern in the Output.", expectedOutPattern.String())
		}
	}
	showOutputOnFailure(t, out)
}
