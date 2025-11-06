package progress

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/stretchr/testify/assert"
)

func TestIndeterminateIndicatorDominoes(t *testing.T) {
	maxLen := 10
	expectedTexts := []string{
		`\\\\\\\\\\`,
		`/\\\\\\\\\`,
		`//\\\\\\\\`,
		`///\\\\\\\`,
		`////\\\\\\`,
		`/////\\\\\`,
		`//////\\\\`,
		`///////\\\`,
		`////////\\`,
		`/////////\`,
		`//////////`,
		`/////////\`,
		`////////\\`,
		`///////\\\`,
		`//////\\\\`,
		`/////\\\\\`,
		`////\\\\\\`,
		`///\\\\\\\`,
		`//\\\\\\\\`,
		`/\\\\\\\\\`,
		`\\\\\\\\\\`,
		`/\\\\\\\\\`,
		`//\\\\\\\\`,
		`///\\\\\\\`,
		`////\\\\\\`,
		`/////\\\\\`,
		`//////\\\\`,
		`///////\\\`,
		`////////\\`,
		`/////////\`,
	}

	out := strings.Builder{}
	f := IndeterminateIndicatorDominoes(time.Microsecond * 10)
	for idx, expectedText := range expectedTexts {
		actual := f(maxLen)
		assert.Equal(t, 0, actual.Position, fmt.Sprintf("expectedTexts[%d]", idx))
		assert.Equal(t, expectedText, actual.Text, fmt.Sprintf("expectedTexts[%d]", idx))
		out.WriteString(fmt.Sprintf("`%v`,\n", actual.Text))
		time.Sleep(time.Microsecond * 10)
	}
	if t.Failed() {
		fmt.Println(out.String())
	}
}

func TestIndeterminateIndicatorColoredDominoes(t *testing.T) {
	maxLen := 10
	colorize := func(s string) string {
		s = strings.ReplaceAll(s, "/", text.FgHiGreen.Sprint("/"))
		s = strings.ReplaceAll(s, "\\", text.FgHiBlack.Sprint("\\"))
		return s
	}

	expectedTexts := []string{
		colorize(`\\\\\\\\\\`),
		colorize(`/\\\\\\\\\`),
		colorize(`//\\\\\\\\`),
		colorize(`///\\\\\\\`),
		colorize(`////\\\\\\`),
		colorize(`/////\\\\\`),
		colorize(`//////\\\\`),
		colorize(`///////\\\`),
		colorize(`////////\\`),
		colorize(`/////////\`),
		colorize(`//////////`),
		colorize(`/////////\`),
		colorize(`////////\\`),
		colorize(`///////\\\`),
		colorize(`//////\\\\`),
		colorize(`/////\\\\\`),
		colorize(`////\\\\\\`),
		colorize(`///\\\\\\\`),
		colorize(`//\\\\\\\\`),
		colorize(`/\\\\\\\\\`),
		colorize(`\\\\\\\\\\`),
		colorize(`/\\\\\\\\\`),
		colorize(`//\\\\\\\\`),
		colorize(`///\\\\\\\`),
		colorize(`////\\\\\\`),
		colorize(`/////\\\\\`),
		colorize(`//////\\\\`),
		colorize(`///////\\\`),
		colorize(`////////\\`),
		colorize(`/////////\`),
	}

	out := strings.Builder{}
	f := IndeterminateIndicatorColoredDominoes(time.Microsecond*10, text.FgHiGreen, text.FgHiBlack)
	for idx, expectedText := range expectedTexts {
		actual := f(maxLen)
		assert.Equal(t, 0, actual.Position, fmt.Sprintf("expectedTexts[%d]", idx))
		assert.Equal(t, expectedText, actual.Text, fmt.Sprintf("expectedTexts[%d]", idx))
		out.WriteString(fmt.Sprintf("`%v`,\n", actual.Text))
		time.Sleep(time.Microsecond * 10)
	}
	if t.Failed() {
		fmt.Println(out.String())
	}
}

func TestIndeterminateIndicatorMovingBackAndForth(t *testing.T) {
	maxLen := 10
	indicator := "<=>"
	expectedPositions := []int{
		0, 1, 2, 3, 4, 5, 6, 7, 6, 5, 4, 3, 2, 1,
		0, 1, 2, 3, 4, 5, 6, 7, 6, 5, 4, 3, 2, 1,
	}

	f := IndeterminateIndicatorMovingBackAndForth(indicator, time.Microsecond*10)
	for idx, expectedPosition := range expectedPositions {
		actual := f(maxLen)
		assert.Equal(t, expectedPosition, actual.Position, fmt.Sprintf("expectedPositions[%d]", idx))
		time.Sleep(time.Microsecond * 10)
	}
}

func Test_indeterminateIndicatorMovingBackAndForth1(t *testing.T) {
	maxLen := 10
	indicator := "?"
	expectedPositions := []int{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2, 1,
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2, 1,
	}

	f := indeterminateIndicatorMovingBackAndForth(indicator)
	for idx, expectedPosition := range expectedPositions {
		actual := f(maxLen)
		assert.Equal(t, expectedPosition, actual.Position, fmt.Sprintf("expectedPositions[%d]", idx))
	}
}

func Test_indeterminateIndicatorMovingBackAndForth2(t *testing.T) {
	maxLen := 10
	indicator := "<>"
	expectedPositions := []int{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 7, 6, 5, 4, 3, 2, 1,
		0, 1, 2, 3, 4, 5, 6, 7, 8, 7, 6, 5, 4, 3, 2, 1,
	}

	f := indeterminateIndicatorMovingBackAndForth(indicator)
	for idx, expectedPosition := range expectedPositions {
		actual := f(maxLen)
		assert.Equal(t, expectedPosition, actual.Position, fmt.Sprintf("expectedPositions[%d]", idx))
	}
}

func Test_indeterminateIndicatorMovingBackAndForth3(t *testing.T) {
	maxLen := 10
	indicator := "<=>"
	expectedPositions := []int{
		0, 1, 2, 3, 4, 5, 6, 7, 6, 5, 4, 3, 2, 1,
		0, 1, 2, 3, 4, 5, 6, 7, 6, 5, 4, 3, 2, 1,
	}

	f := indeterminateIndicatorMovingBackAndForth(indicator)
	for idx, expectedPosition := range expectedPositions {
		actual := f(maxLen)
		assert.Equal(t, expectedPosition, actual.Position, fmt.Sprintf("expectedPositions[%d]", idx))
	}
}

func TestIndeterminateIndicatorMovingLeftToRight(t *testing.T) {
	maxLen := 10
	indicator := "?"
	expectedPositions := []int{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
	}

	f := IndeterminateIndicatorMovingLeftToRight(indicator, time.Microsecond*10)
	for idx, expectedPosition := range expectedPositions {
		actual := f(maxLen)
		assert.Equal(t, expectedPosition, actual.Position, fmt.Sprintf("expectedPositions[%d]", idx))
		time.Sleep(time.Microsecond * 10)
	}
}

func Test_indeterminateIndicatorMovingLeftToRight1(t *testing.T) {
	maxLen := 10
	indicator := "?"
	expectedPositions := []int{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
	}

	f := indeterminateIndicatorMovingLeftToRight(indicator)
	for idx, expectedPosition := range expectedPositions {
		actual := f(maxLen)
		assert.Equal(t, expectedPosition, actual.Position, fmt.Sprintf("expectedPositions[%d]", idx))
	}
}

func Test_indeterminateIndicatorMovingLeftToRight2(t *testing.T) {
	maxLen := 10
	indicator := "<>"
	expectedPositions := []int{
		0, 1, 2, 3, 4, 5, 6, 7, 8,
		0, 1, 2, 3, 4, 5, 6, 7, 8,
	}

	f := indeterminateIndicatorMovingLeftToRight(indicator)
	for idx, expectedPosition := range expectedPositions {
		actual := f(maxLen)
		assert.Equal(t, expectedPosition, actual.Position, fmt.Sprintf("expectedPositions[%d]", idx))
	}
}

func Test_indeterminateIndicatorMovingLeftToRight3(t *testing.T) {
	maxLen := 10
	indicator := "<=>"
	expectedPositions := []int{
		0, 1, 2, 3, 4, 5, 6, 7,
		0, 1, 2, 3, 4, 5, 6, 7,
	}

	f := indeterminateIndicatorMovingLeftToRight(indicator)
	for idx, expectedPosition := range expectedPositions {
		actual := f(maxLen)
		assert.Equal(t, expectedPosition, actual.Position, fmt.Sprintf("expectedPositions[%d]", idx))
	}
}

func TestIndeterminateIndicatorMovingRightToLeft(t *testing.T) {
	maxLen := 10
	indicator := "?"
	expectedPositions := []int{
		9, 8, 7, 6, 5, 4, 3, 2, 1, 0,
		9, 8, 7, 6, 5, 4, 3, 2, 1, 0,
	}

	f := IndeterminateIndicatorMovingRightToLeft(indicator, time.Microsecond*10)
	for idx, expectedPosition := range expectedPositions {
		actual := f(maxLen)
		assert.Equal(t, expectedPosition, actual.Position, fmt.Sprintf("expectedPositions[%d]", idx))
		time.Sleep(time.Microsecond * 10)
	}
}

func Test_indeterminateIndicatorMovingRightToLeft1(t *testing.T) {
	maxLen := 10
	indicator := "?"
	expectedPositions := []int{
		9, 8, 7, 6, 5, 4, 3, 2, 1, 0,
		9, 8, 7, 6, 5, 4, 3, 2, 1, 0,
	}

	f := indeterminateIndicatorMovingRightToLeft(indicator)
	for idx, expectedPosition := range expectedPositions {
		actual := f(maxLen)
		assert.Equal(t, expectedPosition, actual.Position, fmt.Sprintf("expectedPositions[%d]", idx))
	}
}

func Test_indeterminateIndicatorMovingRightToLeft2(t *testing.T) {
	maxLen := 10
	indicator := "<>"
	expectedPositions := []int{
		8, 7, 6, 5, 4, 3, 2, 1, 0,
		8, 7, 6, 5, 4, 3, 2, 1, 0,
	}

	f := indeterminateIndicatorMovingRightToLeft(indicator)
	for idx, expectedPosition := range expectedPositions {
		actual := f(maxLen)
		assert.Equal(t, expectedPosition, actual.Position, fmt.Sprintf("expectedPositions[%d]", idx))
	}
}

func Test_indeterminateIndicatorMovingRightToLeft3(t *testing.T) {
	maxLen := 10
	indicator := "<=>"
	expectedPositions := []int{
		7, 6, 5, 4, 3, 2, 1, 0,
		7, 6, 5, 4, 3, 2, 1, 0,
	}

	f := indeterminateIndicatorMovingRightToLeft(indicator)
	for idx, expectedPosition := range expectedPositions {
		actual := f(maxLen)
		assert.Equal(t, expectedPosition, actual.Position, fmt.Sprintf("expectedPositions[%d]", idx))
	}
}

func TestIndeterminateIndicatorPacMan(t *testing.T) {
	maxLen := 10
	expectedTexts := []string{
		"ᗧ         ",
		" ᗧ        ",
		"  ᗧ       ",
		"   ᗧ      ",
		"    ᗧ     ",
		"     ᗧ    ",
		"      ᗧ   ",
		"       ᗧ  ",
		"        ᗧ ",
		"         ᗧ",
		"        ᗤ ",
		"       ᗤ  ",
		"      ᗤ   ",
		"     ᗤ    ",
		"    ᗤ     ",
		"   ᗤ      ",
		"  ᗤ       ",
		" ᗤ        ",
		"ᗤ         ",
		" ᗧ        ",
		"  ᗧ       ",
		"   ᗧ      ",
		"    ᗧ     ",
		"     ᗧ    ",
		"      ᗧ   ",
		"       ᗧ  ",
		"        ᗧ ",
		"         ᗧ",
	}

	out := strings.Builder{}
	f := IndeterminateIndicatorPacMan(time.Microsecond * 10)
	for idx, expectedText := range expectedTexts {
		actual := f(maxLen)
		assert.Equal(t, expectedText, actual.Text, fmt.Sprintf("expectedTexts[%d]", idx))
		out.WriteString(fmt.Sprintf("%#v,\n", actual.Text))
		time.Sleep(time.Microsecond * 10)
	}
	if t.Failed() {
		fmt.Println(out.String())
	}
}

func TestIndeterminateIndicatorPacManChomp(t *testing.T) {
	maxLen := 10
	colorize := func(s string) string {
		s = strings.ReplaceAll(s, "·", text.FgWhite.Sprint("·"))
		s = strings.ReplaceAll(s, "●", text.FgHiYellow.Sprint("●"))
		s = strings.ReplaceAll(s, "ɔ", text.FgHiYellow.Sprint("ɔ"))
		s = strings.ReplaceAll(s, "c", text.FgHiYellow.Sprint("c"))
		return s
	}

	expectedTexts := []string{
		colorize(" c········"),
		colorize("  c·······"),
		colorize("   ●······"),
		colorize("    ●·····"),
		colorize("     ●····"),
		colorize("      c···"),
		colorize("       c··"),
		colorize("        c·"),
		colorize("·········●"),
		colorize("········● "),
		colorize("·······●  "),
		colorize("······ɔ   "),
		colorize("·····ɔ    "),
		colorize("····ɔ     "),
		colorize("···●      "),
		colorize("··●       "),
		colorize("·●        "),
		colorize("c·········"),
		colorize(" c········"),
		colorize("  c·······"),
		colorize("   ●······"),
		colorize("    ●·····"),
		colorize("     ●····"),
		colorize("      c···"),
		colorize("       c··"),
		colorize("        c·"),
		colorize("·········●"),
		colorize("········● "),
	}

	out := strings.Builder{}
	f := IndeterminateIndicatorPacManChomp(time.Microsecond * 10)
	for idx, expectedText := range expectedTexts {
		actual := f(maxLen)
		assert.Equal(t, expectedText, actual.Text, fmt.Sprintf("expectedTexts[%d]", idx))
		out.WriteString(fmt.Sprintf("%#v,\n", actual.Text))
		time.Sleep(time.Microsecond * 10)
	}
	if t.Failed() {
		fmt.Println(out.String())
	}
}
