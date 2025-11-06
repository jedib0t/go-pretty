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

func Test_convertValueToString(t *testing.T) {
	t.Run("int types", func(t *testing.T) {
		assert.Equal(t, "0", convertValueToString(int(0)))
		assert.Equal(t, "42", convertValueToString(int(42)))
		assert.Equal(t, "-42", convertValueToString(int(-42)))
		assert.Equal(t, "-128", convertValueToString(int8(-128)))
		assert.Equal(t, "127", convertValueToString(int8(127)))
		assert.Equal(t, "-32768", convertValueToString(int16(-32768)))
		assert.Equal(t, "32767", convertValueToString(int16(32767)))
		assert.Equal(t, "-2147483648", convertValueToString(int32(-2147483648)))
		assert.Equal(t, "2147483647", convertValueToString(int32(2147483647)))
		assert.Equal(t, "-9223372036854775808", convertValueToString(int64(-9223372036854775808)))
		assert.Equal(t, "9223372036854775807", convertValueToString(int64(9223372036854775807)))
	})

	t.Run("uint types", func(t *testing.T) {
		assert.Equal(t, "0", convertValueToString(uint(0)))
		assert.Equal(t, "42", convertValueToString(uint(42)))
		assert.Equal(t, "255", convertValueToString(uint8(255)))
		assert.Equal(t, "0", convertValueToString(uint8(0)))
		assert.Equal(t, "65535", convertValueToString(uint16(65535)))
		assert.Equal(t, "4294967295", convertValueToString(uint32(4294967295)))
		assert.Equal(t, "18446744073709551615", convertValueToString(uint64(18446744073709551615)))
	})

	t.Run("float types", func(t *testing.T) {
		assert.Equal(t, "0", convertValueToString(float32(0)))
		assert.Equal(t, "3.14", convertValueToString(float32(3.14)))
		assert.Equal(t, "-3.14", convertValueToString(float32(-3.14)))
		assert.Equal(t, "0", convertValueToString(float64(0)))
		assert.Equal(t, "3.141592653589793", convertValueToString(float64(3.141592653589793)))
		assert.Equal(t, "-3.141592653589793", convertValueToString(float64(-3.141592653589793)))
		// Test scientific notation
		assert.Contains(t, convertValueToString(float64(1e10)), "1e+10")
		assert.Contains(t, convertValueToString(float32(1e10)), "1e+10")
	})

	t.Run("bool", func(t *testing.T) {
		assert.Equal(t, "true", convertValueToString(bool(true)))
		assert.Equal(t, "false", convertValueToString(bool(false)))
	})

	t.Run("default case - string", func(t *testing.T) {
		assert.Equal(t, "hello", convertValueToString("hello"))
		assert.Equal(t, "world", convertValueToString("world"))
		assert.Equal(t, "", convertValueToString(""))
	})

	t.Run("default case - nil", func(t *testing.T) {
		assert.Equal(t, "<nil>", convertValueToString(nil))
	})

	t.Run("default case - slice", func(t *testing.T) {
		assert.Equal(t, "[1 2 3]", convertValueToString([]int{1, 2, 3}))
		assert.Equal(t, "[]", convertValueToString([]string{}))
	})

	t.Run("default case - map", func(t *testing.T) {
		assert.Equal(t, "map[a:1]", convertValueToString(map[string]int{"a": 1}))
		assert.Equal(t, "map[]", convertValueToString(map[string]int{}))
	})

	t.Run("default case - struct", func(t *testing.T) {
		type testStruct struct {
			Field1 string
			Field2 int
		}
		result := convertValueToString(testStruct{Field1: "test", Field2: 42})
		assert.Contains(t, result, "test")
		assert.Contains(t, result, "42")
	})

	t.Run("default case - pointer", func(t *testing.T) {
		x := 42
		result := convertValueToString(&x)
		// fmt.Sprint of a pointer shows the address, which includes "0x"
		assert.Contains(t, result, "0x")
		assert.NotEmpty(t, result)
	})

	t.Run("edge cases", func(t *testing.T) {
		// Zero values
		assert.Equal(t, "0", convertValueToString(int(0)))
		assert.Equal(t, "0", convertValueToString(uint(0)))
		assert.Equal(t, "0", convertValueToString(float32(0)))
		assert.Equal(t, "0", convertValueToString(float64(0)))
		assert.Equal(t, "false", convertValueToString(bool(false)))

		// Max values
		assert.Equal(t, "127", convertValueToString(int8(127)))
		assert.Equal(t, "255", convertValueToString(uint8(255)))
		assert.Equal(t, "32767", convertValueToString(int16(32767)))
		assert.Equal(t, "65535", convertValueToString(uint16(65535)))
	})
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
