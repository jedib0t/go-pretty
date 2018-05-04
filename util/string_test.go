package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetLongestLineLength(t *testing.T) {
	assert.Equal(t, 0, GetLongestLineLength(""))
	assert.Equal(t, 0, GetLongestLineLength("\n\n"))
	assert.Equal(t, 5, GetLongestLineLength("Ghost"))
	assert.Equal(t, 6, GetLongestLineLength("Winter\nIs\nComing"))
	assert.Equal(t, 7, GetLongestLineLength("Mother\nOf\nDragons"))
}
