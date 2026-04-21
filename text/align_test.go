package text

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleAlign_Apply() {
	fmt.Printf("AlignDefault: '%s'\n", AlignDefault.Apply("Jon Snow", 12))
	fmt.Printf("AlignLeft   : '%s'\n", AlignLeft.Apply("Jon Snow", 12))
	fmt.Printf("AlignCenter : '%s'\n", AlignCenter.Apply("Jon Snow", 12))
	fmt.Printf("AlignJustify: '%s'\n", AlignJustify.Apply("Jon Snow", 12))
	fmt.Printf("AlignRight  : '%s'\n", AlignRight.Apply("Jon Snow", 12))
	fmt.Printf("AlignAuto   : '%s'\n", AlignAuto.Apply("Jon Snow", 12))
	fmt.Printf("AlignAuto   : '%s'\n", AlignAuto.Apply("-5.43", 12))

	// Output: AlignDefault: 'Jon Snow    '
	// AlignLeft   : 'Jon Snow    '
	// AlignCenter : '  Jon Snow  '
	// AlignJustify: 'Jon     Snow'
	// AlignRight  : '    Jon Snow'
	// AlignAuto   : 'Jon Snow    '
	// AlignAuto   : '       -5.43'
}

func TestAlign_Apply(t *testing.T) {
	// AlignDefault & AlignLeft are the same
	assert.Equal(t, "Jon Snow    ", AlignDefault.Apply("Jon Snow", 12))
	assert.Equal(t, " Jon Snow   ", AlignDefault.Apply(" Jon Snow", 12))
	assert.Equal(t, "            ", AlignDefault.Apply("", 12))
	assert.Equal(t, "Jon Snow    ", AlignLeft.Apply("Jon Snow   ", 12))
	assert.Equal(t, " Jon Snow   ", AlignLeft.Apply(" Jon Snow  ", 12))
	assert.Equal(t, "            ", AlignLeft.Apply("", 12))

	// AlignCenter
	assert.Equal(t, "  Jon Snow  ", AlignCenter.Apply("Jon Snow ", 12))
	assert.Equal(t, "  Jon Snow  ", AlignCenter.Apply(" Jon Snow", 12))
	assert.Equal(t, "  Jon Snow  ", AlignCenter.Apply(" Jon Snow  ", 12))
	assert.Equal(t, "            ", AlignCenter.Apply("", 12))

	// AlignJustify
	assert.Equal(t, "Jon     Snow", AlignJustify.Apply("Jon Snow", 12))
	assert.Equal(t, "JS  vs.   DT", AlignJustify.Apply("JS vs. DT", 12))
	assert.Equal(t, "JS   is   AT", AlignJustify.Apply("JS is AT", 12))
	assert.Equal(t, "JS   is   AT", AlignJustify.Apply("JS is  AT", 12))
	assert.Equal(t, "JonSnow     ", AlignJustify.Apply("JonSnow", 12))
	assert.Equal(t, "JonSnow     ", AlignJustify.Apply("  JonSnow", 12))
	assert.Equal(t, "            ", AlignJustify.Apply("", 12))

	// Align Right
	assert.Equal(t, "    Jon Snow", AlignRight.Apply("Jon Snow", 12))
	assert.Equal(t, "   Jon Snow ", AlignRight.Apply("Jon Snow ", 12))
	assert.Equal(t, "   Jon Snow ", AlignRight.Apply("  Jon Snow ", 12))
	assert.Equal(t, "            ", AlignRight.Apply("", 12))

	// Align Auto
	assert.Equal(t, "Jon Snow    ", AlignAuto.Apply("Jon Snow", 12))
	assert.Equal(t, "Jon Snow    ", AlignAuto.Apply("Jon Snow ", 12))
	assert.Equal(t, "  Jon Snow  ", AlignAuto.Apply("  Jon Snow ", 12))
	assert.Equal(t, "            ", AlignAuto.Apply("", 12))
	assert.Equal(t, "          13", AlignAuto.Apply("13", 12))
	assert.Equal(t, "       -5.43", AlignAuto.Apply("-5.43", 12))
	assert.Equal(t, "        +.43", AlignAuto.Apply("+.43", 12))
	assert.Equal(t, "       +5.43", AlignAuto.Apply("+5.43", 12))
	assert.Equal(t, "+5.43x      ", AlignAuto.Apply("+5.43x", 12))
}

func TestAlign_Apply_JustifyCJKOverflow(t *testing.T) {
	// U+4222 (䈢) has display width 2; the string below has display width 42,
	// which exceeds maxLength=40. AlignJustify.Apply must not panic.
	cell := "0000000000000000000000000000000000 0000䈢0"
	assert.NotPanics(t, func() {
		assert.Equal(t, cell, AlignJustify.Apply(cell, 40))
	})

	// Further sanity checks: any CJK-only cell whose display width exceeds
	// maxLength must also be returned unchanged without panicking.
	assert.NotPanics(t, func() {
		assert.Equal(t, "中文", AlignJustify.Apply("中文", 3))
	})
}

func TestAlign_Apply_ColoredText(t *testing.T) {
	// AlignDefault & AlignLeft are the same
	assert.Equal(t, "\x1b[33mJon Snow\x1b[0m    ", AlignDefault.Apply("\x1b[33mJon Snow\x1b[0m", 12))
	assert.Equal(t, "\x1b[33m Jon Snow\x1b[0m   ", AlignDefault.Apply("\x1b[33m Jon Snow\x1b[0m", 12))
	assert.Equal(t, "\x1b[33m\x1b[0m            ", AlignDefault.Apply("\x1b[33m\x1b[0m", 12))
	assert.Equal(t, "\x1b[33mJon Snow   \x1b[0m ", AlignLeft.Apply("\x1b[33mJon Snow   \x1b[0m", 12))
	assert.Equal(t, "\x1b[33m Jon Snow  \x1b[0m ", AlignLeft.Apply("\x1b[33m Jon Snow  \x1b[0m", 12))
	assert.Equal(t, "\x1b[33m\x1b[0m            ", AlignLeft.Apply("\x1b[33m\x1b[0m", 12))

	// AlignCenter
	assert.Equal(t, "  \x1b[33mJon Snow \x1b[0m ", AlignCenter.Apply("\x1b[33mJon Snow \x1b[0m", 12))
	assert.Equal(t, "  \x1b[33m Jon Snow\x1b[0m ", AlignCenter.Apply("\x1b[33m Jon Snow\x1b[0m", 12))
	assert.Equal(t, " \x1b[33m Jon Snow  \x1b[0m", AlignCenter.Apply("\x1b[33m Jon Snow  \x1b[0m", 12))
	assert.Equal(t, "      \x1b[33m\x1b[0m      ", AlignCenter.Apply("\x1b[33m\x1b[0m", 12))

	// AlignJustify
	assert.Equal(t, "\x1b[33mJon     Snow\x1b[0m", AlignJustify.Apply("\x1b[33mJon Snow\x1b[0m", 12))
	assert.Equal(t, "\x1b[33mJS  vs.   DT\x1b[0m", AlignJustify.Apply("\x1b[33mJS vs. DT\x1b[0m", 12))
	assert.Equal(t, "\x1b[33mJS   is   AT\x1b[0m", AlignJustify.Apply("\x1b[33mJS is AT\x1b[0m", 12))
	assert.Equal(t, "\x1b[33mJS   is   AT\x1b[0m", AlignJustify.Apply("\x1b[33mJS is  AT\x1b[0m", 12))
	assert.Equal(t, "\x1b[33mJonSnow\x1b[0m     ", AlignJustify.Apply("\x1b[33mJonSnow\x1b[0m", 12))
	assert.Equal(t, "\x1b[33m     JonSnow\x1b[0m", AlignJustify.Apply("\x1b[33m  JonSnow\x1b[0m", 12))
	assert.Equal(t, "\x1b[33m\x1b[0m            ", AlignJustify.Apply("\x1b[33m\x1b[0m", 12))

	// Align Right
	assert.Equal(t, "    \x1b[33mJon Snow\x1b[0m", AlignRight.Apply("\x1b[33mJon Snow\x1b[0m", 12))
	assert.Equal(t, "   \x1b[33mJon Snow \x1b[0m", AlignRight.Apply("\x1b[33mJon Snow \x1b[0m", 12))
	assert.Equal(t, " \x1b[33m  Jon Snow \x1b[0m", AlignRight.Apply("\x1b[33m  Jon Snow \x1b[0m", 12))
	assert.Equal(t, "            \x1b[33m\x1b[0m", AlignRight.Apply("\x1b[33m\x1b[0m", 12))
}

func ExampleAlign_HTMLProperty() {
	fmt.Printf("AlignDefault: '%s'\n", AlignDefault.HTMLProperty())
	fmt.Printf("AlignLeft   : '%s'\n", AlignLeft.HTMLProperty())
	fmt.Printf("AlignCenter : '%s'\n", AlignCenter.HTMLProperty())
	fmt.Printf("AlignJustify: '%s'\n", AlignJustify.HTMLProperty())
	fmt.Printf("AlignRight  : '%s'\n", AlignRight.HTMLProperty())

	// Output: AlignDefault: ''
	// AlignLeft   : 'align="left"'
	// AlignCenter : 'align="center"'
	// AlignJustify: 'align="justify"'
	// AlignRight  : 'align="right"'
}

func TestAlign_HTMLProperty(t *testing.T) {
	aligns := map[Align]string{
		AlignDefault: "",
		AlignLeft:    "left",
		AlignCenter:  "center",
		AlignJustify: "justify",
		AlignRight:   "right",
	}
	for align, htmlStyle := range aligns {
		assert.Contains(t, align.HTMLProperty(), htmlStyle)
	}
}

func ExampleAlign_MarkdownProperty() {
	fmt.Printf("AlignDefault: '%s'\n", AlignDefault.MarkdownProperty())
	fmt.Printf("AlignLeft   : '%s'\n", AlignLeft.MarkdownProperty())
	fmt.Printf("AlignCenter : '%s'\n", AlignCenter.MarkdownProperty())
	fmt.Printf("AlignJustify: '%s'\n", AlignJustify.MarkdownProperty())
	fmt.Printf("AlignRight  : '%s'\n", AlignRight.MarkdownProperty())

	// Output: AlignDefault: ' --- '
	// AlignLeft   : ':--- '
	// AlignCenter : ':---:'
	// AlignJustify: ' --- '
	// AlignRight  : ' ---:'
}

func TestAlign_MarkdownProperty(t *testing.T) {
	aligns := map[Align]string{
		AlignDefault: " --- ",
		AlignLeft:    ":--- ",
		AlignCenter:  ":---:",
		AlignJustify: " --- ",
		AlignRight:   " ---:",
	}
	for align, markdownSeparator := range aligns {
		assert.Contains(t, align.MarkdownProperty(), markdownSeparator)
	}
}

func TestAlign_MarkdownProperty_WithMinLength(t *testing.T) {
	assert.Equal(t, " ---------- ", AlignDefault.MarkdownProperty(10))
	assert.Equal(t, ":---------- ", AlignLeft.MarkdownProperty(10))
	assert.Equal(t, ":----------:", AlignCenter.MarkdownProperty(10))
	assert.Equal(t, " ---------- ", AlignJustify.MarkdownProperty(10))
	assert.Equal(t, " ----------:", AlignRight.MarkdownProperty(10))

	// minimum width of 3
	assert.Equal(t, " --- ", AlignDefault.MarkdownProperty(1))
	assert.Equal(t, " --- ", AlignDefault.MarkdownProperty(3))
	assert.Equal(t, " ---- ", AlignDefault.MarkdownProperty(4))
}

// FuzzAlign_Apply exercises Align.Apply across all alignment modes with
// arbitrary UTF-8 input (including wide Unicode characters) and arbitrary
// maxLength values. It guards against panics such as the "strings: negative
// Repeat count" crash in justifyText when the display width of the cell
// exceeds maxLength.
func FuzzAlign_Apply(f *testing.F) {
	f.Add("Jon Snow", 12)
	f.Add("0000000000000000000000000000000000 0000䈢0", 40)
	f.Add("中文 字符", 3)
	f.Add("", 5)
	f.Add("a b c", 0)
	f.Add("\x1b[33mJon Snow\x1b[0m", 12)

	aligns := []Align{AlignDefault, AlignLeft, AlignCenter, AlignJustify, AlignRight, AlignAuto}
	f.Fuzz(func(t *testing.T, s string, maxLength int) {
		// keep maxLength in a sane range; negative/huge values are not the
		// concern of this fuzz target.
		if maxLength < 0 || maxLength > 1024 {
			t.Skip()
		}
		for _, a := range aligns {
			a.Apply(s, maxLength)
		}
	})
}
