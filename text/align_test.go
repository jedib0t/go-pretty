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

	// Output: AlignDefault: 'Jon Snow    '
	// AlignLeft   : 'Jon Snow    '
	// AlignCenter : '  Jon Snow  '
	// AlignJustify: 'Jon     Snow'
	// AlignRight  : '    Jon Snow'
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
