package gopretty

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isNumber(t *testing.T) {
	assert.True(t, isNumber(int(1)))
	assert.True(t, isNumber(int8(1)))
	assert.True(t, isNumber(int16(1)))
	assert.True(t, isNumber(int32(1)))
	assert.True(t, isNumber(int64(1)))
	assert.True(t, isNumber(uint(1)))
	assert.True(t, isNumber(uint8(1)))
	assert.True(t, isNumber(uint16(1)))
	assert.True(t, isNumber(uint32(1)))
	assert.True(t, isNumber(uint64(1)))
	assert.True(t, isNumber(float32(1)))
	assert.True(t, isNumber(float64(1)))
	assert.False(t, isNumber("1"))
}

func Test_isString(t *testing.T) {
	assert.False(t, isString(int(1)))
	assert.False(t, isString(int8(1)))
	assert.False(t, isString(int16(1)))
	assert.False(t, isString(int32(1)))
	assert.False(t, isString(int64(1)))
	assert.False(t, isString(uint(1)))
	assert.False(t, isString(uint8(1)))
	assert.False(t, isString(uint16(1)))
	assert.False(t, isString(uint32(1)))
	assert.False(t, isString(uint64(1)))
	assert.False(t, isString(float32(1)))
	assert.False(t, isString(float64(1)))
	assert.True(t, isString("1"))
}
