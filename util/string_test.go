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
	assert.Equal(t, 7, GetLongestLineLength("\x1b[33mMother\x1b[0m\nOf\nDragons"))
}

func TestInsertRuneEveryN(t *testing.T) {
	assert.Equal(t, "G-h-o-s-t", InsertRuneEveryN("Ghost", '-', 1))
	assert.Equal(t, "Gh-os-t", InsertRuneEveryN("Ghost", '-', 2))
	assert.Equal(t, "Gho-st", InsertRuneEveryN("Ghost", '-', 3))
	assert.Equal(t, "Ghos-t", InsertRuneEveryN("Ghost", '-', 4))
	assert.Equal(t, "Ghost", InsertRuneEveryN("Ghost", '-', 5))
	assert.Equal(t, "\x1b[33mG-h-o-s-t\x1b[0m", InsertRuneEveryN("\x1b[33mGhost\x1b[0m", '-', 1))
	assert.Equal(t, "\x1b[33mGh-os-t\x1b[0m", InsertRuneEveryN("\x1b[33mGhost\x1b[0m", '-', 2))
	assert.Equal(t, "\x1b[33mGho-st\x1b[0m", InsertRuneEveryN("\x1b[33mGhost\x1b[0m", '-', 3))
	assert.Equal(t, "\x1b[33mGhos-t\x1b[0m", InsertRuneEveryN("\x1b[33mGhost\x1b[0m", '-', 4))
	assert.Equal(t, "\x1b[33mGhost\x1b[0m", InsertRuneEveryN("\x1b[33mGhost\x1b[0m", '-', 5))
}

func TestRuneCountWithoutEscapeSeq(t *testing.T) {
	assert.Equal(t, 0, RuneCountWithoutEscapeSeq(""))
	assert.Equal(t, 5, RuneCountWithoutEscapeSeq("Ghost"))
	assert.Equal(t, 5, RuneCountWithoutEscapeSeq("\x1b[33mGhost\x1b[0m"))
	assert.Equal(t, 5, RuneCountWithoutEscapeSeq("\x1b[33mGhost\x1b[0"))
}

func TestTrimTextWithoutEscapeSeq(t *testing.T) {
	assert.Equal(t, "", TrimTextWithoutEscapeSeq("Ghost", 0))
	assert.Equal(t, "Gho", TrimTextWithoutEscapeSeq("Ghost", 3))
	assert.Equal(t, "Ghost", TrimTextWithoutEscapeSeq("Ghost", 6))
	assert.Equal(t, "\x1b[33mGho\x1b[0m", TrimTextWithoutEscapeSeq("\x1b[33mGhost\x1b[0m", 3))
	assert.Equal(t, "\x1b[33mGhost\x1b[0m", TrimTextWithoutEscapeSeq("\x1b[33mGhost\x1b[0m", 6))
}

func TestWrapText(t *testing.T) {
	assert.Equal(t, "", WrapText("Ghost", 0))
	assert.Equal(t, "G\nh\no\ns\nt", WrapText("Ghost", 1))
	assert.Equal(t, "Gh\nos\nt", WrapText("Ghost", 2))
	assert.Equal(t, "Gho\nst", WrapText("Ghost", 3))
	assert.Equal(t, "Ghos\nt", WrapText("Ghost", 4))
	assert.Equal(t, "Ghost", WrapText("Ghost", 5))
	assert.Equal(t, "Ghost", WrapText("Ghost", 6))
	assert.Equal(t, "Jo\nn\nSn\now", WrapText("Jon\nSnow", 2))
	assert.Equal(t, "Jo\nn\nSn\now\n", WrapText("Jon\nSnow\n", 2))
	assert.Equal(t, "Jon\nSno\nw\n", WrapText("Jon\nSnow\n", 3))
	assert.Equal(t, "\x1b[33mJon\x1b[0m\nSno\nw", WrapText("\x1b[33mJon\x1b[0m\nSnow", 3))
	assert.Equal(t, "\x1b[33mJon\x1b[0m\nSno\nw\n", WrapText("\x1b[33mJon\x1b[0m\nSnow\n", 3))

	complexIn := "+---+------+-------+------+\n| 1 | Arya | Stark | 3000 |\n+---+------+-------+------+"
	assert.Equal(t, complexIn, WrapText(complexIn, 27))
}
