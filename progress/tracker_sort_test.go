package progress

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
