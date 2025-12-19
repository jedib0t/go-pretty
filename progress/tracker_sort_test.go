package progress

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSortBy(t *testing.T) {
	trackers := []*Tracker{
		{Message: "Downloading File # 2", Total: 1000, value: 300},
		{Message: "Downloading File # 1", Total: 1000, value: 100},
		{Message: "Downloading File # 3", Total: 1000, value: 500},
		{Message: "Downloading File # 4", Total: 1000, value: 300, timeStart: time.Now()},
	}

	SortByNone.Sort(trackers)
	assert.Equal(t, "Downloading File # 2", trackers[0].Message)
	assert.Equal(t, "Downloading File # 1", trackers[1].Message)
	assert.Equal(t, "Downloading File # 3", trackers[2].Message)
	assert.Equal(t, "Downloading File # 4", trackers[3].Message)

	SortByMessage.Sort(trackers)
	assert.Equal(t, "Downloading File # 1", trackers[0].Message)
	assert.Equal(t, "Downloading File # 2", trackers[1].Message)
	assert.Equal(t, "Downloading File # 3", trackers[2].Message)
	assert.Equal(t, "Downloading File # 4", trackers[3].Message)

	SortByMessageDsc.Sort(trackers)
	assert.Equal(t, "Downloading File # 4", trackers[0].Message)
	assert.Equal(t, "Downloading File # 3", trackers[1].Message)
	assert.Equal(t, "Downloading File # 2", trackers[2].Message)
	assert.Equal(t, "Downloading File # 1", trackers[3].Message)

	SortByPercent.Sort(trackers)
	assert.Equal(t, "Downloading File # 1", trackers[0].Message)
	assert.Equal(t, "Downloading File # 2", trackers[1].Message)
	assert.Equal(t, "Downloading File # 4", trackers[2].Message)
	assert.Equal(t, "Downloading File # 3", trackers[3].Message)

	SortByPercentDsc.Sort(trackers)
	assert.Equal(t, "Downloading File # 3", trackers[0].Message)
	assert.Equal(t, "Downloading File # 2", trackers[1].Message) // Earlier timeStart comes first when percentages are equal
	assert.Equal(t, "Downloading File # 4", trackers[2].Message)
	assert.Equal(t, "Downloading File # 1", trackers[3].Message)

	SortByValue.Sort(trackers)
	assert.Equal(t, "Downloading File # 1", trackers[0].Message)
	assert.Equal(t, "Downloading File # 2", trackers[1].Message)
	assert.Equal(t, "Downloading File # 4", trackers[2].Message)
	assert.Equal(t, "Downloading File # 3", trackers[3].Message)

	SortByValueDsc.Sort(trackers)
	assert.Equal(t, "Downloading File # 3", trackers[0].Message)
	assert.Equal(t, "Downloading File # 2", trackers[1].Message) // Earlier timeStart comes first when values are equal
	assert.Equal(t, "Downloading File # 4", trackers[2].Message)
	assert.Equal(t, "Downloading File # 1", trackers[3].Message)
}

func TestSortByIndex(t *testing.T) {
	now := time.Now()

	t.Run("basic sorting", func(t *testing.T) {
		trackers := []*Tracker{
			{Message: "Layer 3", Total: 1000, Index: 3, timeStart: now.Add(time.Second * 3)},
			{Message: "Layer 1", Total: 1000, Index: 1, timeStart: now.Add(time.Second * 1)},
			{Message: "Layer 2", Total: 1000, Index: 2, timeStart: now.Add(time.Second * 2)},
			{Message: "Layer 4", Total: 1000, Index: 4, timeStart: now.Add(time.Second * 4)},
		}

		SortByIndex.Sort(trackers)
		assert.Equal(t, "Layer 1", trackers[0].Message)
		assert.Equal(t, "Layer 2", trackers[1].Message)
		assert.Equal(t, "Layer 3", trackers[2].Message)
		assert.Equal(t, "Layer 4", trackers[3].Message)
	})

	t.Run("with zero index", func(t *testing.T) {
		trackers := []*Tracker{
			{Message: "Layer 2", Total: 1000, Index: 2, timeStart: now.Add(time.Second * 3)},
			{Message: "Layer 0b", Total: 1000, Index: 0, timeStart: now.Add(time.Second * 2)},
			{Message: "Layer 1", Total: 1000, Index: 1, timeStart: now.Add(time.Second * 1)},
			{Message: "Layer 0a", Total: 1000, Index: 0, timeStart: now.Add(time.Second * 0)},
		}

		SortByIndex.Sort(trackers)
		// Index 0 should come first, then 1, then 2
		// Same index should maintain insertion order (timeStart as tiebreaker)
		assert.Equal(t, "Layer 0a", trackers[0].Message)
		assert.Equal(t, "Layer 0b", trackers[1].Message)
		assert.Equal(t, "Layer 1", trackers[2].Message)
		assert.Equal(t, "Layer 2", trackers[3].Message)
	})

	t.Run("with same index", func(t *testing.T) {
		trackers := []*Tracker{
			{Message: "Layer 1b", Total: 1000, Index: 1, timeStart: now.Add(time.Second * 2)},
			{Message: "Layer 1a", Total: 1000, Index: 1, timeStart: now.Add(time.Second * 1)},
			{Message: "Layer 2", Total: 1000, Index: 2, timeStart: now.Add(time.Second * 3)},
		}

		SortByIndex.Sort(trackers)
		// Same index should maintain insertion order (timeStart as tiebreaker)
		assert.Equal(t, "Layer 1a", trackers[0].Message)
		assert.Equal(t, "Layer 1b", trackers[1].Message)
		assert.Equal(t, "Layer 2", trackers[2].Message)
	})

	t.Run("descending order", func(t *testing.T) {
		trackers := []*Tracker{
			{Message: "Layer 1", Total: 1000, Index: 1, timeStart: now.Add(time.Second * 1)},
			{Message: "Layer 3", Total: 1000, Index: 3, timeStart: now.Add(time.Second * 3)},
			{Message: "Layer 2", Total: 1000, Index: 2, timeStart: now.Add(time.Second * 2)},
			{Message: "Layer 4", Total: 1000, Index: 4, timeStart: now.Add(time.Second * 4)},
		}

		SortByIndexDsc.Sort(trackers)
		// Higher indices should come first
		assert.Equal(t, "Layer 4", trackers[0].Message)
		assert.Equal(t, "Layer 3", trackers[1].Message)
		assert.Equal(t, "Layer 2", trackers[2].Message)
		assert.Equal(t, "Layer 1", trackers[3].Message)
	})

	t.Run("descending order with same index", func(t *testing.T) {
		trackers := []*Tracker{
			{Message: "Layer 2b", Total: 1000, Index: 2, timeStart: now.Add(time.Second * 3)},
			{Message: "Layer 1", Total: 1000, Index: 1, timeStart: now.Add(time.Second * 1)},
			{Message: "Layer 2a", Total: 1000, Index: 2, timeStart: now.Add(time.Second * 2)},
		}

		SortByIndexDsc.Sort(trackers)
		// Higher indices should come first
		// Same index should maintain insertion order (earlier timeStart first)
		assert.Equal(t, "Layer 2a", trackers[0].Message)
		assert.Equal(t, "Layer 2b", trackers[1].Message)
		assert.Equal(t, "Layer 1", trackers[2].Message)
	})
}
