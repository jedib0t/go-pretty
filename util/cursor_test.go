package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCursor_Sprint(t *testing.T) {
	assert.Equal(t, "\x1b[5B", CursorDown.Sprint(5))
	assert.Equal(t, "\x1b[5D", CursorLeft.Sprint(5))
	assert.Equal(t, "\x1b[5C", CursorRight.Sprint(5))
	assert.Equal(t, "\x1b[5A", CursorUp.Sprint(5))
	assert.Equal(t, "\x1b[K", EraseLine.Sprint())
}
