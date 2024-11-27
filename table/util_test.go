package table

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleAutoIndexColumnID() {
	fmt.Printf("AutoIndexColumnID(    0): \"%s\"\n", AutoIndexColumnID(0))
	fmt.Printf("AutoIndexColumnID(    1): \"%s\"\n", AutoIndexColumnID(1))
	fmt.Printf("AutoIndexColumnID(    2): \"%s\"\n", AutoIndexColumnID(2))
	fmt.Printf("AutoIndexColumnID(   25): \"%s\"\n", AutoIndexColumnID(25))
	fmt.Printf("AutoIndexColumnID(   26): \"%s\"\n", AutoIndexColumnID(26))
	fmt.Printf("AutoIndexColumnID(  702): \"%s\"\n", AutoIndexColumnID(702))
	fmt.Printf("AutoIndexColumnID(18278): \"%s\"\n", AutoIndexColumnID(18278))

	// Output: AutoIndexColumnID(    0): "A"
	// AutoIndexColumnID(    1): "B"
	// AutoIndexColumnID(    2): "C"
	// AutoIndexColumnID(   25): "Z"
	// AutoIndexColumnID(   26): "AA"
	// AutoIndexColumnID(  702): "AAA"
	// AutoIndexColumnID(18278): "AAAA"
}

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
	assert.False(t, isNumber(nil))
}

func Test_objAsSlice(t *testing.T) {
	a, b, c := 1, 2, 3
	assert.Equal(t, "[1 2 3]", fmt.Sprint(objAsSlice([]int{a, b, c})))
	assert.Equal(t, "[1 2 3]", fmt.Sprint(objAsSlice(&[]int{a, b, c})))
	assert.Equal(t, "[1 2 3]", fmt.Sprint(objAsSlice(&[]*int{&a, &b, &c})))
	assert.Equal(t, "[1 2]", fmt.Sprint(objAsSlice(&[]*int{&a, &b, nil})))
	assert.Equal(t, "[1]", fmt.Sprint(objAsSlice(&[]*int{&a, nil, nil})))
	assert.Equal(t, "[]", fmt.Sprint(objAsSlice(&[]*int{nil, nil, nil})))
	assert.Equal(t, "[<nil> 2]", fmt.Sprint(objAsSlice(&[]*int{nil, &b, nil})))
}

func Test_objIsSlice(t *testing.T) {
	assert.True(t, objIsSlice([]int{}))
	assert.True(t, objIsSlice([]*int{}))
	assert.False(t, objIsSlice(&[]int{}))
	assert.False(t, objIsSlice(&[]*int{}))
	assert.False(t, objIsSlice(Table{}))
	assert.False(t, objIsSlice(&Table{}))
	assert.False(t, objIsSlice(nil))
}
