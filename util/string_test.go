package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLongestLineLength(t *testing.T) {
	assert.Equal(t, 0, GetLongestLineLength(""))
	assert.Equal(t, 0, GetLongestLineLength("\n\n"))
	assert.Equal(t, 5, GetLongestLineLength("Ghost"))
	assert.Equal(t, 6, GetLongestLineLength("Winter\nIs\nComing"))
	assert.Equal(t, 7, GetLongestLineLength("Mother\nOf\nDragons"))
}

func TestInsertRuneEveryN(t *testing.T) {
	assert.Equal(t, "G-h-o-s-t", InsertRuneEveryN("Ghost", '-', 1))
	assert.Equal(t, "Gh-os-t", InsertRuneEveryN("Ghost", '-', 2))
	assert.Equal(t, "Gho-st", InsertRuneEveryN("Ghost", '-', 3))
	assert.Equal(t, "Ghos-t", InsertRuneEveryN("Ghost", '-', 4))
	assert.Equal(t, "Ghost", InsertRuneEveryN("Ghost", '-', 5))
}

func TestWrapText(t *testing.T) {
	assert.Equal(t, "G\nh\no\ns\nt", WrapText("Ghost", 1))
	assert.Equal(t, "Gh\nos\nt", WrapText("Ghost", 2))
	assert.Equal(t, "Gho\nst", WrapText("Ghost", 3))
	assert.Equal(t, "Ghos\nt", WrapText("Ghost", 4))
	assert.Equal(t, "Ghost", WrapText("Ghost", 5))
	assert.Equal(t, "Ghost", WrapText("Ghost", 6))
	assert.Equal(t, "Jo\nn\nSn\now", WrapText("Jon\nSnow", 2))
	assert.Equal(t, "Jo\nn\nSn\now\n", WrapText("Jon\nSnow\n", 2))
}
