package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAutoIndexColumnID(t *testing.T) {
	assert.Equal(t, "A", AutoIndexColumnID(0))
	assert.Equal(t, "Z", AutoIndexColumnID(25))
	assert.Equal(t, "AA", AutoIndexColumnID(26))
	assert.Equal(t, "ZZ", AutoIndexColumnID(701))
	assert.Equal(t, "AAA", AutoIndexColumnID(702))
	assert.Equal(t, "ZZZ", AutoIndexColumnID(18277))
	assert.Equal(t, "AAAA", AutoIndexColumnID(18278))
}
