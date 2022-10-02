package text

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDirection_Modifier(t *testing.T) {
	assert.Equal(t, "", Default.Modifier())
	assert.Equal(t, "\u202a", LeftToRight.Modifier())
	assert.Equal(t, "\u202b", RightToLeft.Modifier())
}
