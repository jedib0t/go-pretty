package table

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

func TestIsNumber(t *testing.T) {
	assert.True(t, IsNumber(int(1)))
	assert.True(t, IsNumber(int8(1)))
	assert.True(t, IsNumber(int16(1)))
	assert.True(t, IsNumber(int32(1)))
	assert.True(t, IsNumber(int64(1)))
	assert.True(t, IsNumber(uint(1)))
	assert.True(t, IsNumber(uint8(1)))
	assert.True(t, IsNumber(uint16(1)))
	assert.True(t, IsNumber(uint32(1)))
	assert.True(t, IsNumber(uint64(1)))
	assert.True(t, IsNumber(float32(1)))
	assert.True(t, IsNumber(float64(1)))
	assert.False(t, IsNumber("1"))
}
