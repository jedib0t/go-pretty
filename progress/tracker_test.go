package progress

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"time"
)

func TestTracker_Increment(t *testing.T) {
	tracker := Tracker{Total: 100}
	assert.Equal(t, int64(0), tracker.value)
	assert.Equal(t, int64(100), tracker.Total)

	tracker.Increment(10)
	assert.Equal(t, int64(10), tracker.value)
	assert.Equal(t, int64(100), tracker.Total)

	tracker.Increment(100)
	assert.Equal(t, int64(110), tracker.value)
	assert.Equal(t, int64(110), tracker.Total)
	assert.False(t, tracker.timeStop.IsZero())
	assert.True(t, tracker.IsDone())
}

func TestTracker_IsDone(t *testing.T) {
	tracker := Tracker{Total: 10}
	assert.False(t, tracker.IsDone())

	tracker.Increment(10)
	assert.True(t, tracker.IsDone())
}

func TestTracker_MarkAsDone(t *testing.T) {
	tracker := Tracker{}
	assert.False(t, tracker.IsDone())
	assert.True(t, tracker.timeStop.IsZero())

	tracker.MarkAsDone()
	assert.True(t, tracker.IsDone())
	assert.False(t, tracker.timeStop.IsZero())
}

func TestTracker_PercentDone(t *testing.T) {
	tracker := Tracker{Total: 100}
	assert.Equal(t, 0.00, tracker.PercentDone())

	for idx := 1; idx <= 100; idx++ {
		tracker.Increment(1)
		assert.Equal(t, float64(idx), tracker.PercentDone())
	}
}

func TestTracker_Reset(t *testing.T) {
	tracker := Tracker{Total: 100}
	assert.False(t, tracker.done)
	assert.Equal(t, time.Time{}, tracker.timeStart)
	assert.Equal(t, time.Time{}, tracker.timeStop)
	assert.Equal(t, int64(0), tracker.value)

	tracker.start()
	tracker.Increment(tracker.Total)
	tracker.stop()
	assert.True(t, tracker.done)
	assert.NotEqual(t, time.Time{}, tracker.timeStart)
	assert.NotEqual(t, time.Time{}, tracker.timeStop)
	assert.Equal(t, tracker.Total, tracker.value)

	tracker.Reset()
	assert.False(t, tracker.done)
	assert.Equal(t, time.Time{}, tracker.timeStart)
	assert.Equal(t, time.Time{}, tracker.timeStop)
	assert.Equal(t, int64(0), tracker.value)
}

func TestUnits_Sprint(t *testing.T) {
	assert.Equal(t, "1", UnitsDefault.Sprint(1))
	assert.Equal(t, "1.50K", UnitsDefault.Sprint(1500))
	assert.Equal(t, "1.50M", UnitsDefault.Sprint(1500000))
	assert.Equal(t, "1.50B", UnitsDefault.Sprint(1500000000))
	assert.Equal(t, "1.50T", UnitsDefault.Sprint(1500000000000))
	assert.Equal(t, "1.50Q", UnitsDefault.Sprint(1500000000000000))
	assert.Equal(t, "1500.00Q", UnitsDefault.Sprint(1500000000000000000))

	assert.Equal(t, "1B", UnitsBytes.Sprint(1))
	assert.Equal(t, "1.50KB", UnitsBytes.Sprint(1500))
	assert.Equal(t, "1.50MB", UnitsBytes.Sprint(1500000))
	assert.Equal(t, "1.50GB", UnitsBytes.Sprint(1500000000))
	assert.Equal(t, "1.50TB", UnitsBytes.Sprint(1500000000000))
	assert.Equal(t, "1.50PB", UnitsBytes.Sprint(1500000000000000))
	assert.Equal(t, "1500.00PB", UnitsBytes.Sprint(1500000000000000000))

	assert.Equal(t, "$1", UnitsCurrencyDollar.Sprint(1))
	assert.Equal(t, "$1.50K", UnitsCurrencyDollar.Sprint(1500))
	assert.Equal(t, "$1.50M", UnitsCurrencyDollar.Sprint(1500000))
	assert.Equal(t, "$1.50B", UnitsCurrencyDollar.Sprint(1500000000))
	assert.Equal(t, "$1.50T", UnitsCurrencyDollar.Sprint(1500000000000))
	assert.Equal(t, "$1.50Q", UnitsCurrencyDollar.Sprint(1500000000000000))
	assert.Equal(t, "$1500.00Q", UnitsCurrencyDollar.Sprint(1500000000000000000))

	assert.Equal(t, "₠1", UnitsCurrencyEuro.Sprint(1))
	assert.Equal(t, "₠1.50K", UnitsCurrencyEuro.Sprint(1500))
	assert.Equal(t, "₠1.50M", UnitsCurrencyEuro.Sprint(1500000))
	assert.Equal(t, "₠1.50B", UnitsCurrencyEuro.Sprint(1500000000))
	assert.Equal(t, "₠1.50T", UnitsCurrencyEuro.Sprint(1500000000000))
	assert.Equal(t, "₠1.50Q", UnitsCurrencyEuro.Sprint(1500000000000000))
	assert.Equal(t, "₠1500.00Q", UnitsCurrencyEuro.Sprint(1500000000000000000))

	assert.Equal(t, "£1", UnitsCurrencyPound.Sprint(1))
	assert.Equal(t, "£1.50K", UnitsCurrencyPound.Sprint(1500))
	assert.Equal(t, "£1.50M", UnitsCurrencyPound.Sprint(1500000))
	assert.Equal(t, "£1.50B", UnitsCurrencyPound.Sprint(1500000000))
	assert.Equal(t, "£1.50T", UnitsCurrencyPound.Sprint(1500000000000))
	assert.Equal(t, "£1.50Q", UnitsCurrencyPound.Sprint(1500000000000000))
	assert.Equal(t, "£1500.00Q", UnitsCurrencyPound.Sprint(1500000000000000000))
}

func TestSortBy(t *testing.T) {
	trackers := []*Tracker{
		{Message: "Downloading File # 2", Total: 1000, value: 300},
		{Message: "Downloading File # 1", Total: 1000, value: 100},
		{Message: "Downloading File # 3", Total: 1000, value: 500},
	}

	SortByNone.Sort(trackers)
	assert.Equal(t, "Downloading File # 2", trackers[0].Message)
	assert.Equal(t, "Downloading File # 1", trackers[1].Message)
	assert.Equal(t, "Downloading File # 3", trackers[2].Message)

	SortByMessage.Sort(trackers)
	assert.Equal(t, "Downloading File # 1", trackers[0].Message)
	assert.Equal(t, "Downloading File # 2", trackers[1].Message)
	assert.Equal(t, "Downloading File # 3", trackers[2].Message)

	SortByMessageDsc.Sort(trackers)
	assert.Equal(t, "Downloading File # 3", trackers[0].Message)
	assert.Equal(t, "Downloading File # 2", trackers[1].Message)
	assert.Equal(t, "Downloading File # 1", trackers[2].Message)

	SortByPercent.Sort(trackers)
	assert.Equal(t, "Downloading File # 1", trackers[0].Message)
	assert.Equal(t, "Downloading File # 2", trackers[1].Message)
	assert.Equal(t, "Downloading File # 3", trackers[2].Message)

	SortByPercentDsc.Sort(trackers)
	assert.Equal(t, "Downloading File # 3", trackers[0].Message)
	assert.Equal(t, "Downloading File # 2", trackers[1].Message)
	assert.Equal(t, "Downloading File # 1", trackers[2].Message)

	SortByValue.Sort(trackers)
	assert.Equal(t, "Downloading File # 1", trackers[0].Message)
	assert.Equal(t, "Downloading File # 2", trackers[1].Message)
	assert.Equal(t, "Downloading File # 3", trackers[2].Message)

	SortByValueDsc.Sort(trackers)
	assert.Equal(t, "Downloading File # 3", trackers[0].Message)
	assert.Equal(t, "Downloading File # 2", trackers[1].Message)
	assert.Equal(t, "Downloading File # 1", trackers[2].Message)
}
