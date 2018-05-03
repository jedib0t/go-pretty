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
