package text

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDirection_Modifier(t *testing.T) {
	assert.Equal(t, "", Default.Modifier())
	assert.Equal(t, string(RuneL2R), LeftToRight.Modifier())
	assert.Equal(t, string(RuneR2L), RightToLeft.Modifier())
}
