package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestIsString(t *testing.T) {
	assert.False(t, IsString(int(1)))
	assert.False(t, IsString(int8(1)))
	assert.False(t, IsString(int16(1)))
	assert.False(t, IsString(int32(1)))
	assert.False(t, IsString(int64(1)))
	assert.False(t, IsString(uint(1)))
	assert.False(t, IsString(uint8(1)))
	assert.False(t, IsString(uint16(1)))
	assert.False(t, IsString(uint32(1)))
	assert.False(t, IsString(uint64(1)))
	assert.False(t, IsString(float32(1)))
	assert.False(t, IsString(float64(1)))
	assert.True(t, IsString("1"))
}
