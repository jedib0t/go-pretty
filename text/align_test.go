package text

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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

	// Align Right
	assert.Equal(t, "    Jon Snow", AlignRight.Apply("Jon Snow", 12))
	assert.Equal(t, "   Jon Snow ", AlignRight.Apply("Jon Snow ", 12))
	assert.Equal(t, "   Jon Snow ", AlignRight.Apply("  Jon Snow ", 12))
	assert.Equal(t, "            ", AlignRight.Apply("", 12))

	// AlignJustify
	assert.Equal(t, "Jon     Snow", AlignJustify.Apply("Jon Snow", 12))
	assert.Equal(t, "JS  vs.   DT", AlignJustify.Apply("JS vs. DT", 12))
	assert.Equal(t, "JS   is   AT", AlignJustify.Apply("JS is AT", 12))
	assert.Equal(t, "JS   is   AT", AlignJustify.Apply("JS is  AT", 12))
	assert.Equal(t, "JonSnow     ", AlignJustify.Apply("JonSnow", 12))
	assert.Equal(t, "JonSnow     ", AlignJustify.Apply("  JonSnow", 12))
	assert.Equal(t, "            ", AlignJustify.Apply("", 12))
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

	// Align Right
	assert.Equal(t, "    \x1b[33mJon Snow\x1b[0m", AlignRight.Apply("\x1b[33mJon Snow\x1b[0m", 12))
	assert.Equal(t, "   \x1b[33mJon Snow \x1b[0m", AlignRight.Apply("\x1b[33mJon Snow \x1b[0m", 12))
	assert.Equal(t, " \x1b[33m  Jon Snow \x1b[0m", AlignRight.Apply("\x1b[33m  Jon Snow \x1b[0m", 12))
	assert.Equal(t, "            \x1b[33m\x1b[0m", AlignRight.Apply("\x1b[33m\x1b[0m", 12))

	// AlignJustify
	assert.Equal(t, "\x1b[33mJon     Snow\x1b[0m", AlignJustify.Apply("\x1b[33mJon Snow\x1b[0m", 12))
	assert.Equal(t, "\x1b[33mJS  vs.   DT\x1b[0m", AlignJustify.Apply("\x1b[33mJS vs. DT\x1b[0m", 12))
	assert.Equal(t, "\x1b[33mJS   is   AT\x1b[0m", AlignJustify.Apply("\x1b[33mJS is AT\x1b[0m", 12))
	assert.Equal(t, "\x1b[33mJS   is   AT\x1b[0m", AlignJustify.Apply("\x1b[33mJS is  AT\x1b[0m", 12))
	assert.Equal(t, "\x1b[33mJonSnow\x1b[0m     ", AlignJustify.Apply("\x1b[33mJonSnow\x1b[0m", 12))
	assert.Equal(t, "\x1b[33m     JonSnow\x1b[0m", AlignJustify.Apply("\x1b[33m  JonSnow\x1b[0m", 12))
	assert.Equal(t, "\x1b[33m\x1b[0m            ", AlignJustify.Apply("\x1b[33m\x1b[0m", 12))
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
