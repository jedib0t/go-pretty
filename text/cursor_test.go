package text

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCursor_Sprint(t *testing.T) {
	assert.Equal(t, "\x1b[B", CursorDown.Sprint())
	assert.Equal(t, "\x1b[D", CursorLeft.Sprint())
	assert.Equal(t, "\x1b[C", CursorRight.Sprint())
	assert.Equal(t, "\x1b[A", CursorUp.Sprint())
	assert.Equal(t, "\x1b[K", EraseLine.Sprint())
}

func TestCursor_Sprintn(t *testing.T) {
	assert.Equal(t, "\x1b[5B", CursorDown.Sprintn(5))
	assert.Equal(t, "\x1b[5D", CursorLeft.Sprintn(5))
	assert.Equal(t, "\x1b[5C", CursorRight.Sprintn(5))
	assert.Equal(t, "\x1b[5A", CursorUp.Sprintn(5))
	assert.Equal(t, "\x1b[K", EraseLine.Sprintn(5))
}
