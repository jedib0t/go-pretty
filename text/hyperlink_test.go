package text

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHyperlink(t *testing.T) {
	assert.Equal(t, "Ghost", Hyperlink("", "Ghost"))
	assert.Equal(t, "https://example.com", Hyperlink("https://example.com", ""))
	assert.Equal(t, "\x1b]8;;https://example.com\x1b\\Ghost\x1b]8;;\x1b\\", Hyperlink("https://example.com", "Ghost"))

	// control characters in the URL should not be able to terminate the OSC 8
	// sequence early and inject escape sequences into the output
	assert.Equal(t, "\x1b]8;;https://example.com/\\[31mred\x1b\\Ghost\x1b]8;;\x1b\\",
		Hyperlink("https://example.com/\x1b\\\x1b[31mred", "Ghost"))
	assert.Equal(t, "\x1b]8;;https://example.com/ab\x1b\\Ghost\x1b]8;;\x1b\\",
		Hyperlink("https://example.com/a\x07b", "Ghost"))
	assert.Equal(t, "Ghost", Hyperlink("\x1b\x07\n", "Ghost"))
}
