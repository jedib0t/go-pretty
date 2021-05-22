package text

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleInsertEveryN() {
	fmt.Printf("InsertEveryN(\"Ghost\", '-', 0): %#v\n", InsertEveryN("Ghost", '-', 0))
	fmt.Printf("InsertEveryN(\"Ghost\", '-', 1): %#v\n", InsertEveryN("Ghost", '-', 1))
	fmt.Printf("InsertEveryN(\"Ghost\", '-', 2): %#v\n", InsertEveryN("Ghost", '-', 2))
	fmt.Printf("InsertEveryN(\"Ghost\", '-', 3): %#v\n", InsertEveryN("Ghost", '-', 3))
	fmt.Printf("InsertEveryN(\"Ghost\", '-', 4): %#v\n", InsertEveryN("Ghost", '-', 4))
	fmt.Printf("InsertEveryN(\"Ghost\", '-', 5): %#v\n", InsertEveryN("Ghost", '-', 5))
	fmt.Printf("InsertEveryN(\"\\x1b[33mGhost\\x1b[0m\", '-', 0): %#v\n", InsertEveryN("\x1b[33mGhost\x1b[0m", '-', 0))
	fmt.Printf("InsertEveryN(\"\\x1b[33mGhost\\x1b[0m\", '-', 1): %#v\n", InsertEveryN("\x1b[33mGhost\x1b[0m", '-', 1))
	fmt.Printf("InsertEveryN(\"\\x1b[33mGhost\\x1b[0m\", '-', 2): %#v\n", InsertEveryN("\x1b[33mGhost\x1b[0m", '-', 2))
	fmt.Printf("InsertEveryN(\"\\x1b[33mGhost\\x1b[0m\", '-', 3): %#v\n", InsertEveryN("\x1b[33mGhost\x1b[0m", '-', 3))
	fmt.Printf("InsertEveryN(\"\\x1b[33mGhost\\x1b[0m\", '-', 4): %#v\n", InsertEveryN("\x1b[33mGhost\x1b[0m", '-', 4))
	fmt.Printf("InsertEveryN(\"\\x1b[33mGhost\\x1b[0m\", '-', 5): %#v\n", InsertEveryN("\x1b[33mGhost\x1b[0m", '-', 5))

	// Output: InsertEveryN("Ghost", '-', 0): "Ghost"
	// InsertEveryN("Ghost", '-', 1): "G-h-o-s-t"
	// InsertEveryN("Ghost", '-', 2): "Gh-os-t"
	// InsertEveryN("Ghost", '-', 3): "Gho-st"
	// InsertEveryN("Ghost", '-', 4): "Ghos-t"
	// InsertEveryN("Ghost", '-', 5): "Ghost"
	// InsertEveryN("\x1b[33mGhost\x1b[0m", '-', 0): "\x1b[33mGhost\x1b[0m"
	// InsertEveryN("\x1b[33mGhost\x1b[0m", '-', 1): "\x1b[33mG-h-o-s-t\x1b[0m"
	// InsertEveryN("\x1b[33mGhost\x1b[0m", '-', 2): "\x1b[33mGh-os-t\x1b[0m"
	// InsertEveryN("\x1b[33mGhost\x1b[0m", '-', 3): "\x1b[33mGho-st\x1b[0m"
	// InsertEveryN("\x1b[33mGhost\x1b[0m", '-', 4): "\x1b[33mGhos-t\x1b[0m"
	// InsertEveryN("\x1b[33mGhost\x1b[0m", '-', 5): "\x1b[33mGhost\x1b[0m"
}

func TestInsertEveryN(t *testing.T) {
	assert.Equal(t, "Ghost", InsertEveryN("Ghost", '-', 0))
	assert.Equal(t, "Gツhツoツsツt", InsertEveryN("Ghost", 'ツ', 1))
	assert.Equal(t, "G-h-o-s-t", InsertEveryN("Ghost", '-', 1))
	assert.Equal(t, "Gh-os-t", InsertEveryN("Ghost", '-', 2))
	assert.Equal(t, "Gho-st", InsertEveryN("Ghost", '-', 3))
	assert.Equal(t, "Ghos-t", InsertEveryN("Ghost", '-', 4))
	assert.Equal(t, "Ghost", InsertEveryN("Ghost", '-', 5))
	assert.Equal(t, "\x1b[33mGhost\x1b[0m", InsertEveryN("\x1b[33mGhost\x1b[0m", '-', 0))
	assert.Equal(t, "\x1b[33mGツhツoツsツt\x1b[0m", InsertEveryN("\x1b[33mGhost\x1b[0m", 'ツ', 1))
	assert.Equal(t, "\x1b[33mG-h-o-s-t\x1b[0m", InsertEveryN("\x1b[33mGhost\x1b[0m", '-', 1))
	assert.Equal(t, "\x1b[33mGh-os-t\x1b[0m", InsertEveryN("\x1b[33mGhost\x1b[0m", '-', 2))
	assert.Equal(t, "\x1b[33mGho-st\x1b[0m", InsertEveryN("\x1b[33mGhost\x1b[0m", '-', 3))
	assert.Equal(t, "\x1b[33mGhos-t\x1b[0m", InsertEveryN("\x1b[33mGhost\x1b[0m", '-', 4))
	assert.Equal(t, "\x1b[33mGhost\x1b[0m", InsertEveryN("\x1b[33mGhost\x1b[0m", '-', 5))
}

func ExampleLongestLineLen() {
	fmt.Printf("LongestLineLen(\"\"): %d\n", LongestLineLen(""))
	fmt.Printf("LongestLineLen(\"\\n\\n\"): %d\n", LongestLineLen("\n\n"))
	fmt.Printf("LongestLineLen(\"Ghost\"): %d\n", LongestLineLen("Ghost"))
	fmt.Printf("LongestLineLen(\"Ghostツ\"): %d\n", LongestLineLen("Ghostツ"))
	fmt.Printf("LongestLineLen(\"Winter\\nIs\\nComing\"): %d\n", LongestLineLen("Winter\nIs\nComing"))
	fmt.Printf("LongestLineLen(\"Mother\\nOf\\nDragons\"): %d\n", LongestLineLen("Mother\nOf\nDragons"))
	fmt.Printf("LongestLineLen(\"\\x1b[33mMother\\x1b[0m\\nOf\\nDragons\"): %d\n", LongestLineLen("\x1b[33mMother\x1b[0m\nOf\nDragons"))

	// Output: LongestLineLen(""): 0
	// LongestLineLen("\n\n"): 0
	// LongestLineLen("Ghost"): 5
	// LongestLineLen("Ghostツ"): 7
	// LongestLineLen("Winter\nIs\nComing"): 6
	// LongestLineLen("Mother\nOf\nDragons"): 7
	// LongestLineLen("\x1b[33mMother\x1b[0m\nOf\nDragons"): 7
}

func TestLongestLineLen(t *testing.T) {
	assert.Equal(t, 0, LongestLineLen(""))
	assert.Equal(t, 0, LongestLineLen("\n\n"))
	assert.Equal(t, 5, LongestLineLen("Ghost"))
	assert.Equal(t, 7, LongestLineLen("Ghostツ"))
	assert.Equal(t, 6, LongestLineLen("Winter\nIs\nComing"))
	assert.Equal(t, 7, LongestLineLen("Mother\nOf\nDragons"))
	assert.Equal(t, 7, LongestLineLen("\x1b[33mMother\x1b[0m\nOf\nDragons"))
}

func ExamplePad() {
	fmt.Printf("%#v\n", Pad("Ghost", 0, ' '))
	fmt.Printf("%#v\n", Pad("Ghost", 3, ' '))
	fmt.Printf("%#v\n", Pad("Ghost", 5, ' '))
	fmt.Printf("%#v\n", Pad("\x1b[33mGhost\x1b[0m", 7, '-'))
	fmt.Printf("%#v\n", Pad("\x1b[33mGhost\x1b[0m", 10, '.'))

	// Output: "Ghost"
	// "Ghost"
	// "Ghost"
	// "\x1b[33mGhost\x1b[0m--"
	// "\x1b[33mGhost\x1b[0m....."
}

func TestPad(t *testing.T) {
	assert.Equal(t, "Ghost", Pad("Ghost", 0, ' '))
	assert.Equal(t, "Ghost", Pad("Ghost", 3, ' '))
	assert.Equal(t, "Ghost", Pad("Ghost", 5, ' '))
	assert.Equal(t, "Ghost  ", Pad("Ghost", 7, ' '))
	assert.Equal(t, "Ghost.....", Pad("Ghost", 10, '.'))
	assert.Equal(t, "\x1b[33mGhost\x1b[0  ", Pad("\x1b[33mGhost\x1b[0", 7, ' '))
	assert.Equal(t, "\x1b[33mGhost\x1b[0.....", Pad("\x1b[33mGhost\x1b[0", 10, '.'))
}

func ExampleRepeatAndTrim() {
	fmt.Printf("RepeatAndTrim(\"\", 5): %#v\n", RepeatAndTrim("", 5))
	fmt.Printf("RepeatAndTrim(\"Ghost\", 0): %#v\n", RepeatAndTrim("Ghost", 0))
	fmt.Printf("RepeatAndTrim(\"Ghost\", 3): %#v\n", RepeatAndTrim("Ghost", 3))
	fmt.Printf("RepeatAndTrim(\"Ghost\", 5): %#v\n", RepeatAndTrim("Ghost", 5))
	fmt.Printf("RepeatAndTrim(\"Ghost\", 7): %#v\n", RepeatAndTrim("Ghost", 7))
	fmt.Printf("RepeatAndTrim(\"Ghost\", 10): %#v\n", RepeatAndTrim("Ghost", 10))

	// Output: RepeatAndTrim("", 5): ""
	// RepeatAndTrim("Ghost", 0): ""
	// RepeatAndTrim("Ghost", 3): "Gho"
	// RepeatAndTrim("Ghost", 5): "Ghost"
	// RepeatAndTrim("Ghost", 7): "GhostGh"
	// RepeatAndTrim("Ghost", 10): "GhostGhost"
}

func TestRepeatAndTrim(t *testing.T) {
	assert.Equal(t, "", RepeatAndTrim("", 5))
	assert.Equal(t, "", RepeatAndTrim("Ghost", 0))
	assert.Equal(t, "Gho", RepeatAndTrim("Ghost", 3))
	assert.Equal(t, "Ghost", RepeatAndTrim("Ghost", 5))
	assert.Equal(t, "GhostGh", RepeatAndTrim("Ghost", 7))
	assert.Equal(t, "GhostGhost", RepeatAndTrim("Ghost", 10))
	assert.Equal(t, "───", RepeatAndTrim("─", 3))
}

func ExampleRuneCount() {
	fmt.Printf("RuneCount(\"\"): %d\n", RuneCount(""))
	fmt.Printf("RuneCount(\"Ghost\"): %d\n", RuneCount("Ghost"))
	fmt.Printf("RuneCount(\"Ghostツ\"): %d\n", RuneCount("Ghostツ"))
	fmt.Printf("RuneCount(\"\\x1b[33mGhost\\x1b[0m\"): %d\n", RuneCount("\x1b[33mGhost\x1b[0m"))
	fmt.Printf("RuneCount(\"\\x1b[33mGhost\\x1b[0\"): %d\n", RuneCount("\x1b[33mGhost\x1b[0"))

	// Output: RuneCount(""): 0
	// RuneCount("Ghost"): 5
	// RuneCount("Ghostツ"): 7
	// RuneCount("\x1b[33mGhost\x1b[0m"): 5
	// RuneCount("\x1b[33mGhost\x1b[0"): 5
}

func TestRuneCount(t *testing.T) {
	assert.Equal(t, 0, RuneCount(""))
	assert.Equal(t, 5, RuneCount("Ghost"))
	assert.Equal(t, 7, RuneCount("Ghostツ"))
	assert.Equal(t, 5, RuneCount("\x1b[33mGhost\x1b[0m"))
	assert.Equal(t, 5, RuneCount("\x1b[33mGhost\x1b[0"))
}

func ExampleRuneWidth() {
	fmt.Printf("RuneWidth('A'): %d\n", RuneWidth('A'))
	fmt.Printf("RuneWidth('ツ'): %d\n", RuneWidth('ツ'))
	fmt.Printf("RuneWidth('⊙'): %d\n", RuneWidth('⊙'))
	fmt.Printf("RuneWidth('︿'): %d\n", RuneWidth('︿'))
	fmt.Printf("RuneWidth(rune(27)): %d\n", RuneWidth(rune(27))) // ANSI escape sequence

	// Output: RuneWidth('A'): 1
	// RuneWidth('ツ'): 2
	// RuneWidth('⊙'): 1
	// RuneWidth('︿'): 2
	// RuneWidth(rune(27)): 0
}

func TestRuneWidth(t *testing.T) {
	assert.Equal(t, 1, RuneWidth('A'))
	assert.Equal(t, 2, RuneWidth('ツ'))
	assert.Equal(t, 1, RuneWidth('⊙'))
	assert.Equal(t, 2, RuneWidth('︿'))
	assert.Equal(t, 0, RuneWidth(rune(27))) // ANSI escape sequence
}

func ExampleSnip() {
	fmt.Printf("Snip(\"Ghost\", 0, \"~\"): %#v\n", Snip("Ghost", 0, "~"))
	fmt.Printf("Snip(\"Ghost\", 1, \"~\"): %#v\n", Snip("Ghost", 1, "~"))
	fmt.Printf("Snip(\"Ghost\", 3, \"~\"): %#v\n", Snip("Ghost", 3, "~"))
	fmt.Printf("Snip(\"Ghost\", 5, \"~\"): %#v\n", Snip("Ghost", 5, "~"))
	fmt.Printf("Snip(\"Ghost\", 7, \"~\"): %#v\n", Snip("Ghost", 7, "~"))
	fmt.Printf("Snip(\"\\x1b[33mGhost\\x1b[0m\", 7, \"~\"): %#v\n", Snip("\x1b[33mGhost\x1b[0m", 7, "~"))

	// Output: Snip("Ghost", 0, "~"): "Ghost"
	// Snip("Ghost", 1, "~"): "~"
	// Snip("Ghost", 3, "~"): "Gh~"
	// Snip("Ghost", 5, "~"): "Ghost"
	// Snip("Ghost", 7, "~"): "Ghost"
	// Snip("\x1b[33mGhost\x1b[0m", 7, "~"): "\x1b[33mGhost\x1b[0m"
}

func TestSnip(t *testing.T) {
	assert.Equal(t, "Ghost", Snip("Ghost", 0, "~"))
	assert.Equal(t, "~", Snip("Ghost", 1, "~"))
	assert.Equal(t, "Gh~", Snip("Ghost", 3, "~"))
	assert.Equal(t, "Ghost", Snip("Ghost", 5, "~"))
	assert.Equal(t, "Ghost", Snip("Ghost", 7, "~"))
	assert.Equal(t, "\x1b[33mGhost\x1b[0m", Snip("\x1b[33mGhost\x1b[0m", 7, "~"))
}

func ExampleTrim() {
	fmt.Printf("Trim(\"Ghost\", 0): %#v\n", Trim("Ghost", 0))
	fmt.Printf("Trim(\"Ghost\", 3): %#v\n", Trim("Ghost", 3))
	fmt.Printf("Trim(\"Ghost\", 6): %#v\n", Trim("Ghost", 6))
	fmt.Printf("Trim(\"\\x1b[33mGhost\\x1b[0m\", 0): %#v\n", Trim("\x1b[33mGhost\x1b[0m", 0))
	fmt.Printf("Trim(\"\\x1b[33mGhost\\x1b[0m\", 3): %#v\n", Trim("\x1b[33mGhost\x1b[0m", 3))
	fmt.Printf("Trim(\"\\x1b[33mGhost\\x1b[0m\", 6): %#v\n", Trim("\x1b[33mGhost\x1b[0m", 6))

	// Output: Trim("Ghost", 0): ""
	// Trim("Ghost", 3): "Gho"
	// Trim("Ghost", 6): "Ghost"
	// Trim("\x1b[33mGhost\x1b[0m", 0): ""
	// Trim("\x1b[33mGhost\x1b[0m", 3): "\x1b[33mGho\x1b[0m"
	// Trim("\x1b[33mGhost\x1b[0m", 6): "\x1b[33mGhost\x1b[0m"
}

func TestTrim(t *testing.T) {
	assert.Equal(t, "", Trim("Ghost", 0))
	assert.Equal(t, "Gho", Trim("Ghost", 3))
	assert.Equal(t, "Ghost", Trim("Ghost", 6))
	assert.Equal(t, "\x1b[33mGho\x1b[0m", Trim("\x1b[33mGhost\x1b[0m", 3))
	assert.Equal(t, "\x1b[33mGhost\x1b[0m", Trim("\x1b[33mGhost\x1b[0m", 6))
}
