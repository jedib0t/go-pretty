package progress

import (
	"fmt"
	"sort"
	"time"
)

// Tracker helps track the progress of a single task. The way to use it is to
// instantiate a Tracker with a valid Message, a valid (expected) Total, and
// Units values. This should then be fed to the Progress Writer with the
// Writer.AppendTracker() method. When the task that is being done has progress,
// increment the value using the Tracker.Increment(value) method.
type Tracker struct {
	// Message should contain a short description of the "task"
	Message string
	// Total should be set to the (expected) Total/Final value to be reached
	Total int64
	// Units defines the type of the "value" being tracked
	Units Units

	done      bool
	timeStart time.Time
	timeStop  time.Time
	value     int64
}

// Increment updates the current value of the task being tracked.
func (t *Tracker) Increment(value int64) {
	if !t.done {
		t.value += value
		if t.Total > 0 && t.value >= t.Total {
			t.stop()
		}
	}
}

// IsDone returns true if the tracker is done (value has reached the expected
// Total set during initialization).
func (t *Tracker) IsDone() bool {
	return t.done
}

// MarkAsDone forces completion of the tracker by updating the current value as
// the expected Total value.
func (t *Tracker) MarkAsDone() {
	t.Total = t.value
	t.stop()
}

// PercentDone returns the currently completed percentage value.
func (t *Tracker) PercentDone() float64 {
	return float64(t.value) * 100.0 / float64(t.Total)
}

// Reset resets the tracker to its initial state.
func (t *Tracker) Reset() {
	t.done = false
	t.timeStart = time.Time{}
	t.timeStop = time.Time{}
	t.value = 0
}

func (t *Tracker) start() {
	t.done = false
	t.timeStart = time.Now()
}

func (t *Tracker) stop() {
	t.done = true
	t.timeStop = time.Now()
	if t.value > t.Total {
		t.Total = t.value
	}
}

// Units defines the "type" of the value being tracked by the Tracker.
type Units int

const (
	// UnitsDefault doesn't define any units. The value will be treated as any
	// other number.
	UnitsDefault Units = iota

	// UnitsBytes defines the value as a storage unit. Values will be converted
	// and printed in one of these forms: B, KB, MB, GB, TB, PB
	UnitsBytes

	// UnitsCurrencyDollar defines the value as a Dollar amount. Values will be
	// converted and printed in one of these forms: $x.yz, $x.yzK, $x.yzM,
	// $x.yzB, $x.yzT
	UnitsCurrencyDollar

	// UnitsCurrencyEuro defines the value as a Euro amount. Values will be
	// converted and printed in one of these forms: ₠x.yz, ₠x.yzK, ₠x.yzM,
	// ₠x.yzB, ₠x.yzT
	UnitsCurrencyEuro

	// UnitsCurrencyPound defines the value as a Pound amount. Values will be
	// converted and printed in one of these forms: £x.yz, £x.yzK, £x.yzM,
	// £x.yzB, £x.yzT
	UnitsCurrencyPound
)

// Sprint prints the value as defined by the Units.
func (tu Units) Sprint(value int64) string {
	switch tu {
	case UnitsBytes:
		return tu.sprintBytes(value)
	case UnitsCurrencyDollar:
		return "$" + tu.sprintAll(value)
	case UnitsCurrencyEuro:
		return "₠" + tu.sprintAll(value)
	case UnitsCurrencyPound:
		return "£" + tu.sprintAll(value)
	default:
		return tu.sprintAll(value)
	}
}

func (tu Units) sprintAll(value int64) string {
	if value < 1000 {
		return fmt.Sprintf("%d", value)
	} else if value < 1000000 {
		return fmt.Sprintf("%.2fK", float64(value)/1000.0)
	} else if value < 1000000000 {
		return fmt.Sprintf("%.2fM", float64(value)/1000000.0)
	} else if value < 1000000000000 {
		return fmt.Sprintf("%.2fB", float64(value)/1000000000.0)
	} else if value < 1000000000000000 {
		return fmt.Sprintf("%.2fT", float64(value)/1000000000000.0)
	}
	return fmt.Sprintf("%.2fQ", float64(value)/1000000000000000.0)
}

func (tu Units) sprintBytes(value int64) string {
	if value < 1000 {
		return fmt.Sprintf("%dB", value)
	} else if value < 1000000 {
		return fmt.Sprintf("%.2fKB", float64(value)/1000.0)
	} else if value < 1000000000 {
		return fmt.Sprintf("%.2fMB", float64(value)/1000000.0)
	} else if value < 1000000000000 {
		return fmt.Sprintf("%.2fGB", float64(value)/1000000000.0)
	} else if value < 1000000000000000 {
		return fmt.Sprintf("%.2fTB", float64(value)/1000000000000.0)
	}
	return fmt.Sprintf("%.2fPB", float64(value)/1000000000000000.0)
}

// SortBy helps sort a list of Trackers by various means.
type SortBy int

const (
	// SortByNone doesn't do any sorting == sort by insertion order.
	SortByNone SortBy = iota

	// SortByMessage sorts by the Message alphabetically in ascending order.
	SortByMessage

	// SortByMessageDsc sorts by the Message alphabetically in descending order.
	SortByMessageDsc

	// SortByPercent sorts by the Percentage complete in ascending order.
	SortByPercent

	// SortByPercentDsc sorts by the Percentage complete in descending order.
	SortByPercentDsc

	// SortByValue sorts by the Value in ascending order.
	SortByValue

	// SortByValueDsc sorts by the Value in descending order.
	SortByValueDsc
)

// Sort applies the sorting method defined by SortBy.
func (ts SortBy) Sort(trackers []*Tracker) {
	switch ts {
	case SortByMessage:
		sort.Sort(sortByMessage(trackers))
	case SortByMessageDsc:
		sort.Sort(sortByMessageDsc(trackers))
	case SortByPercent:
		sort.Sort(sortByPercent(trackers))
	case SortByPercentDsc:
		sort.Sort(sortByPercentDsc(trackers))
	case SortByValue:
		sort.Sort(sortByValue(trackers))
	case SortByValueDsc:
		sort.Sort(sortByValueDsc(trackers))
	default:
		// no sort
	}
}

type sortByMessage []*Tracker

func (ta sortByMessage) Len() int           { return len(ta) }
func (ta sortByMessage) Swap(i, j int)      { ta[i], ta[j] = ta[j], ta[i] }
func (ta sortByMessage) Less(i, j int) bool { return ta[i].Message < ta[j].Message }

type sortByMessageDsc []*Tracker

func (ta sortByMessageDsc) Len() int           { return len(ta) }
func (ta sortByMessageDsc) Swap(i, j int)      { ta[i], ta[j] = ta[j], ta[i] }
func (ta sortByMessageDsc) Less(i, j int) bool { return ta[i].Message > ta[j].Message }

type sortByPercent []*Tracker

func (ta sortByPercent) Len() int           { return len(ta) }
func (ta sortByPercent) Swap(i, j int)      { ta[i], ta[j] = ta[j], ta[i] }
func (ta sortByPercent) Less(i, j int) bool { return ta[i].PercentDone() < ta[j].PercentDone() }

type sortByPercentDsc []*Tracker

func (ta sortByPercentDsc) Len() int           { return len(ta) }
func (ta sortByPercentDsc) Swap(i, j int)      { ta[i], ta[j] = ta[j], ta[i] }
func (ta sortByPercentDsc) Less(i, j int) bool { return ta[i].PercentDone() > ta[j].PercentDone() }

type sortByValue []*Tracker

func (ta sortByValue) Len() int           { return len(ta) }
func (ta sortByValue) Swap(i, j int)      { ta[i], ta[j] = ta[j], ta[i] }
func (ta sortByValue) Less(i, j int) bool { return ta[i].value < ta[j].value }

type sortByValueDsc []*Tracker

func (ta sortByValueDsc) Len() int           { return len(ta) }
func (ta sortByValueDsc) Swap(i, j int)      { ta[i], ta[j] = ta[j], ta[i] }
func (ta sortByValueDsc) Less(i, j int) bool { return ta[i].value > ta[j].value }
