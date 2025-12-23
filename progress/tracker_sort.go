package progress

import "sort"

// SortBy helps sort a list of Trackers by various means.
type SortBy int

const (
	// SortByNone doesn't do any sorting == sort by insertion order.
	SortByNone SortBy = iota

	// SortByIndex sorts by the Index field in ascending order. When this is used,
	// trackers are rendered in Index order regardless of completion status (done
	// and active trackers are merged and sorted together). Index 0 comes first.
	SortByIndex

	// SortByIndexDsc sorts by the Index field in descending order. When this is used,
	// trackers are rendered in Index order regardless of completion status (done
	// and active trackers are merged and sorted together). Higher indices come first.
	SortByIndexDsc

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
func (sb SortBy) Sort(trackers []*Tracker) {
	switch sb {
	case SortByIndex:
		sort.Sort(sortByIndex(trackers))
	case SortByIndexDsc:
		sort.Stable(sortByIndexDsc(trackers))
	case SortByMessage:
		sort.Sort(sortByMessage(trackers))
	case SortByMessageDsc:
		sort.Sort(sortDsc{sortByMessage(trackers)})
	case SortByPercent:
		sort.Stable(sortByPercent(trackers))
	case SortByPercentDsc:
		sort.Stable(sortByPercentDsc(trackers))
	case SortByValue:
		sort.Stable(sortByValue(trackers))
	case SortByValueDsc:
		sort.Stable(sortByValueDsc(trackers))
	default:
		// no sort
	}
}

type sortByIndex []*Tracker

func (sb sortByIndex) Len() int      { return len(sb) }
func (sb sortByIndex) Swap(i, j int) { sb[i], sb[j] = sb[j], sb[i] }
func (sb sortByIndex) Less(i, j int) bool {
	if sb[i].Index == sb[j].Index {
		// Same index: maintain insertion order (use timeStart as tiebreaker)
		return sb[i].timeStartValue().Before(sb[j].timeStartValue())
	}
	return sb[i].Index < sb[j].Index
}

type sortByIndexDsc []*Tracker

func (sb sortByIndexDsc) Len() int      { return len(sb) }
func (sb sortByIndexDsc) Swap(i, j int) { sb[i], sb[j] = sb[j], sb[i] }
func (sb sortByIndexDsc) Less(i, j int) bool {
	if sb[i].Index == sb[j].Index {
		// Same index: maintain insertion order (earlier timeStart first)
		return sb[i].timeStartValue().Before(sb[j].timeStartValue())
	}
	// Reverse: higher index comes first
	return sb[i].Index > sb[j].Index
}

type sortByMessage []*Tracker

func (sb sortByMessage) Len() int           { return len(sb) }
func (sb sortByMessage) Swap(i, j int)      { sb[i], sb[j] = sb[j], sb[i] }
func (sb sortByMessage) Less(i, j int) bool { return sb[i].message() < sb[j].message() }

type sortByPercent []*Tracker

func (sb sortByPercent) Len() int      { return len(sb) }
func (sb sortByPercent) Swap(i, j int) { sb[i], sb[j] = sb[j], sb[i] }
func (sb sortByPercent) Less(i, j int) bool {
	if sb[i].PercentDone() == sb[j].PercentDone() {
		// When percentages are equal, preserve insertion order (earlier timeStart first)
		return sb[i].timeStartValue().Before(sb[j].timeStartValue())
	}
	return sb[i].PercentDone() < sb[j].PercentDone()
}

type sortByPercentDsc []*Tracker

func (sb sortByPercentDsc) Len() int      { return len(sb) }
func (sb sortByPercentDsc) Swap(i, j int) { sb[i], sb[j] = sb[j], sb[i] }
func (sb sortByPercentDsc) Less(i, j int) bool {
	if sb[i].PercentDone() == sb[j].PercentDone() {
		// When percentages are equal, preserve insertion order (earlier timeStart first)
		return sb[i].timeStartValue().Before(sb[j].timeStartValue())
	}
	// Reverse: higher percentage comes first
	return sb[i].PercentDone() > sb[j].PercentDone()
}

type sortByValue []*Tracker

func (sb sortByValue) Len() int      { return len(sb) }
func (sb sortByValue) Swap(i, j int) { sb[i], sb[j] = sb[j], sb[i] }
func (sb sortByValue) Less(i, j int) bool {
	valueI := sb[i].Value()
	valueJ := sb[j].Value()
	if valueI == valueJ {
		return sb[i].timeStartValue().Before(sb[j].timeStartValue())
	}
	return valueI < valueJ
}

type sortByValueDsc []*Tracker

func (sb sortByValueDsc) Len() int      { return len(sb) }
func (sb sortByValueDsc) Swap(i, j int) { sb[i], sb[j] = sb[j], sb[i] }
func (sb sortByValueDsc) Less(i, j int) bool {
	valueI := sb[i].Value()
	valueJ := sb[j].Value()
	if valueI == valueJ {
		// When values are equal, preserve insertion order (earlier timeStart first)
		return sb[i].timeStartValue().Before(sb[j].timeStartValue())
	}
	// Reverse: higher value comes first
	return valueI > valueJ
}

type sortDsc struct{ sort.Interface }

func (sd sortDsc) Less(i, j int) bool {
	// Reverse the comparison for descending order
	// When elements are equal (both Less calls return false), preserve insertion order
	return !sd.Interface.Less(i, j) && sd.Interface.Less(j, i)
}
