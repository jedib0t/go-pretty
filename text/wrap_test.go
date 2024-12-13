package text

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	textTable = "+---+------+-------+------+\n| 1 | Arya | Stark | 3000 |\n+---+------+-------+------+"

	// colored text with nested escape codes: "{red}{bold}...{un-bold}...{reset}"
	textUnBold     = "\x1b[91m\x1b[1mBold Title\x1b[22m Regular Red Text\x1b[0m"
	expectedUnBold = "\x1b[91m\x1b[1mBold Title\x1b[22m Regular Red \x1b[0m\n\x1b[91mText\x1b[0m"

	// text with wide characters
	textWide     = "abcd甲乙丙丁abcd"
	expectedWide = "abcd甲乙丙\n丁abcd"

	// colored text with wide characters
	textWideColored     = "\x1b[22mabcd甲乙丙丁abcd\x1b[0m"
	expectedWideColored = "\x1b[22mabcd甲乙丙\x1b[0m\n\u001B[22m丁abcd\u001B[0m"
)

func ExampleWrapHard() {
	str := `The quick brown fox jumped over the lazy dog.

A big crocodile died empty-fanged, gulping horribly in jerking kicking little
motions. Nonchalant old Peter Quinn ruthlessly shot the under-water vermin with
Xavier yelling Zap!`
	strWrapped := WrapHard(str, 30)
	for idx, line := range strings.Split(strWrapped, "\n") {
		fmt.Printf("Line #%02d: '%s'\n", idx+1, line)
	}

	// Output: Line #01: 'The quick brown fox jumped ove'
	// Line #02: 'r the lazy dog.'
	// Line #03: ''
	// Line #04: 'A big crocodile died empty-fan'
	// Line #05: 'ged, gulping horribly in jerki'
	// Line #06: 'ng kicking little motions. Non'
	// Line #07: 'chalant old Peter Quinn ruthle'
	// Line #08: 'ssly shot the under-water verm'
	// Line #09: 'in with Xavier yelling Zap!'
}

func TestWrapHard(t *testing.T) {
	assert.Equal(t, "", WrapHard("Ghost", 0))
	assert.Equal(t, "G\nh\no\ns\nt", WrapHard("Ghost", 1))
	assert.Equal(t, "Gh\nos\nt", WrapHard("Ghost", 2))
	assert.Equal(t, "Gho\nst", WrapHard("Ghost", 3))
	assert.Equal(t, "Ghos\nt", WrapHard("Ghost", 4))
	assert.Equal(t, "Ghost", WrapHard("Ghost", 5))
	assert.Equal(t, "Ghost", WrapHard("Ghost", 6))
	assert.Equal(t, "Jo\nn \nSn\now", WrapHard("Jon\nSnow", 2))
	assert.Equal(t, "Jo\nn \nSn\now", WrapHard("Jon\nSnow\n", 2))
	assert.Equal(t, "Jon\nSno\nw", WrapHard("Jon\nSnow\n", 3))
	assert.Equal(t, "Jon i\ns a S\nnow", WrapHard("Jon is a Snow", 5))
	assert.Equal(t, "\x1b[33mJon\x1b[0m\nSno\nw", WrapHard("\x1b[33mJon\x1b[0m\nSnow", 3))
	assert.Equal(t, "\x1b[33mJon\x1b[0m\nSno\nw", WrapHard("\x1b[33mJon\x1b[0m\nSnow\n", 3))
	assert.Equal(t, "\x1b[33mJon\x1b[0m\n\x1b[33mSno\x1b[0m\n\x1b[33mw\x1b[0m", WrapHard("\x1b[33mJon Snow\x1b[0m", 3))
	assert.Equal(t, "\x1b[33mJon\x1b[0m\n\x1b[33mSno\x1b[0m\n\x1b[33mw\x1b[0m", WrapHard("\x1b[33mJon Snow\n", 3))
	assert.Equal(t, "\x1b[33mJon\x1b[0m\n\x1b[33mSno\x1b[0m\n\x1b[33mw \x1b[0m", WrapHard("\x1b[33mJon Snow\n\x1b[0m", 3))

	assert.Equal(t, textTable, WrapHard(textTable, 27))
	assert.Equal(t, expectedUnBold, WrapHard(textUnBold, 23))
	assert.Equal(t, expectedWide, WrapHard(textWide, 10))
	assert.Equal(t, expectedWideColored, WrapHard(textWideColored, 10))
}

func TestFoo(t *testing.T) {
	assert.Equal(t, "\x1b[33mJon\x1b[0m\n\x1b[33mSno\x1b[0m\n\x1b[33mw\x1b[0m", WrapHard("\x1b[33mJon Snow\x1b[0m", 3))
}

func ExampleWrapSoft() {
	str := `The quick brown fox jumped over the lazy dog.

A big crocodile died empty-fanged, gulping horribly in jerking kicking little
motions. Nonchalant old Peter Quinn ruthlessly shot the under-water vermin with
Xavier yelling Zap!`
	strWrapped := WrapSoft(str, 30)
	for idx, line := range strings.Split(strWrapped, "\n") {
		fmt.Printf("Line #%02d: '%s'\n", idx+1, line)
	}

	// Output: Line #01: 'The quick brown fox jumped    '
	// Line #02: 'over the lazy dog.'
	// Line #03: ''
	// Line #04: 'A big crocodile died          '
	// Line #05: 'empty-fanged, gulping horribly'
	// Line #06: 'in jerking kicking little     '
	// Line #07: 'motions. Nonchalant old Peter '
	// Line #08: 'Quinn ruthlessly shot the     '
	// Line #09: 'under-water vermin with Xavier'
	// Line #10: 'yelling Zap!'
}

func TestWrapSoft(t *testing.T) {
	assert.Equal(t, "", WrapSoft("Ghost", 0))
	assert.Equal(t, "G\nh\no\ns\nt", WrapSoft("Ghost", 1))
	assert.Equal(t, "Gh\nos\nt", WrapSoft("Ghost", 2))
	assert.Equal(t, "Gho\nst", WrapSoft("Ghost", 3))
	assert.Equal(t, "Ghos\nt", WrapSoft("Ghost", 4))
	assert.Equal(t, "Ghost", WrapSoft("Ghost", 5))
	assert.Equal(t, "Ghost", WrapSoft("Ghost", 6))
	assert.Equal(t, "Jo\nn \nSn\now", WrapSoft("Jon\nSnow", 2))
	assert.Equal(t, "Jo\nn \nSn\now", WrapSoft("Jon\nSnow\n", 2))
	assert.Equal(t, "Jon\nSno\nw", WrapSoft("Jon\nSnow\n", 3))
	assert.Equal(t, "Jon \nSnow", WrapSoft("Jon\nSnow", 4))
	assert.Equal(t, "Jon \nSnow", WrapSoft("Jon\nSnow\n", 4))
	assert.Equal(t, "Jon  \nis a \nSnow", WrapSoft("Jon is a Snow", 5))
	assert.Equal(t, "\x1b[33mJon\x1b[0m\nSno\nw", WrapSoft("\x1b[33mJon\x1b[0m\nSnow", 3))
	assert.Equal(t, "\x1b[33mJon\x1b[0m\nSno\nw", WrapSoft("\x1b[33mJon\x1b[0m\nSnow\n", 3))
	assert.Equal(t, "\x1b[33mJon\x1b[0m\n\x1b[33mSno\x1b[0m\n\x1b[33mw\x1b[0m", WrapSoft("\x1b[33mJon Snow\x1b[0m", 3))
	assert.Equal(t, "\x1b[33mJon\x1b[0m\n\x1b[33mSno\x1b[0m\n\x1b[33mw\x1b[0m", WrapSoft("\x1b[33mJon Snow\n", 3))
	assert.Equal(t, "\x1b[33mJon\x1b[0m\n\x1b[33mSno\x1b[0m\n\x1b[33mw \x1b[0m", WrapSoft("\x1b[33mJon Snow\n\x1b[0m", 3))
	assert.Equal(t, "\x1b[33mJon \x1b[0m\n\x1b[33mSnow\x1b[0m", WrapSoft("\x1b[33mJon Snow\x1b[0m", 4))
	assert.Equal(t, "\x1b[33mJon \x1b[0m\n\x1b[33mSnow\x1b[0m\n\x1b[33m???\x1b[0m", WrapSoft("\x1b[33mJon Snow???\x1b[0m", 4))

	assert.Equal(t, textTable, WrapSoft(textTable, 27))
	assert.Equal(t, expectedUnBold, WrapSoft(textUnBold, 23))
	assert.Equal(t, expectedWide, WrapHard(textWide, 10))
	assert.Equal(t, expectedWideColored, WrapHard(textWideColored, 10))
}

func ExampleWrapText() {
	str := `The quick brown fox jumped over the lazy dog.

A big crocodile died empty-fanged, gulping horribly in jerking kicking little
motions. Nonchalant old Peter Quinn ruthlessly shot the under-water vermin with
Xavier yelling Zap!`
	strWrapped := WrapText(str, 30)
	for idx, line := range strings.Split(strWrapped, "\n") {
		fmt.Printf("Line #%02d: '%s'\n", idx+1, line)
	}

	// Output: Line #01: 'The quick brown fox jumped ove'
	// Line #02: 'r the lazy dog.'
	// Line #03: ''
	// Line #04: 'A big crocodile died empty-fan'
	// Line #05: 'ged, gulping horribly in jerki'
	// Line #06: 'ng kicking little'
	// Line #07: 'motions. Nonchalant old Peter '
	// Line #08: 'Quinn ruthlessly shot the unde'
	// Line #09: 'r-water vermin with'
	// Line #10: 'Xavier yelling Zap!'
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
	assert.Equal(t, "\x1b[33mJon\x1b[0m\n\x1b[33mSno\x1b[0m\n\x1b[33mw\x1b[0m", WrapText("\x1b[33mJon Snow\x1b[0m", 3))
	assert.Equal(t, "\x1b[33mJon\x1b[0m\n\x1b[33mSno\x1b[0m\n\x1b[33mw\x1b[0m\n", WrapText("\x1b[33mJon Snow\n", 3))
	assert.Equal(t, "\x1b[33mJon\x1b[0m\n\x1b[33mSno\x1b[0m\n\x1b[33mw\x1b[0m\n\x1b[0m", WrapText("\x1b[33mJon Snow\n\x1b[0m", 3))

	assert.Equal(t, textTable, WrapText(textTable, 27))
	assert.Equal(t, expectedUnBold, WrapText(textUnBold, 23))
	assert.Equal(t, expectedWide, WrapHard(textWide, 10))
	assert.Equal(t, expectedWideColored, WrapHard(textWideColored, 10))
}
