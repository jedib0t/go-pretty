package text

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHyperlink(t *testing.T) {
	assert.Equal(t, "Ghost", Hyperlink("", "Ghost"))
	assert.Equal(t, "https://example.com", Hyperlink("https://example.com", ""))
	assert.Equal(t, "\x1b]8;;https://example.com\x1b\\Ghost\x1b]8;;\x1b\\", Hyperlink("https://example.com", "Ghost"))
}
