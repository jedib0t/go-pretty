package text

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewNumberTransformer(t *testing.T) {
	signColorsMap := map[string]Colors{
		"negative": colorsNumberNegative,
		"positive": colorsNumberPositive,
		"zero":     nil,
	}
	colorValuesMap := map[string]map[interface{}]string{
		"negative": {
			int(-5):           "%05d",
			int8(-5):          "%05d",
			int16(-5):         "%05d",
			int32(-5):         "%05d",
			int64(-5):         "%05d",
			float32(-5.55555): "%08.2f",
			float64(-5.55555): "%08.2f",
		},
		"positive": {
			int(5):           "%05d",
			int8(5):          "%05d",
			int16(5):         "%05d",
			int32(5):         "%05d",
			int64(5):         "%05d",
			uint(5):          "%05d",
			uint8(5):         "%05d",
			uint16(5):        "%05d",
			uint32(5):        "%05d",
			uint64(5):        "%05d",
			float32(5.55555): "%08.2f",
			float64(5.55555): "%08.2f",
		},
		"zero": {
			int(0):           "%05d",
			int8(0):          "%05d",
			int16(0):         "%05d",
			int32(0):         "%05d",
			int64(0):         "%05d",
			uint(0):          "%05d",
			uint8(0):         "%05d",
			uint16(0):        "%05d",
			uint32(0):        "%05d",
			uint64(0):        "%05d",
			float32(0.00000): "%08.2f",
			float64(0.00000): "%08.2f",
		},
	}

	for sign, valuesFormatMap := range colorValuesMap {
		for value, format := range valuesFormatMap {
			transformer := NewNumberTransformer(format)
			expected := signColorsMap[sign].Sprintf(format, value)
			if sign == "negative" {
				expected = strings.Replace(expected, "-0", "-00", 1)
			}
			actual := transformer(value)
			message := fmt.Sprintf("%s.%s: expected=%v, actual=%v; format=%#v",
				sign, reflect.TypeOf(value).Kind(), expected, actual, format)

			assert.Equal(t, expected, actual, message)
		}
	}

	// invalid input
	assert.Equal(t, "foo", NewNumberTransformer("%05d")("foo"))
}

type jsonTest struct {
	Foo string       `json:"foo"`
	Bar int32        `json:"bar"`
	Baz float64      `json:"baz"`
	Nan jsonNestTest `json:"nan"`
}

type jsonNestTest struct {
	A string
	B int32
	C float64
}

func TestNewJSONTransformer(t *testing.T) {
	transformer := NewJSONTransformer("", "    ")

	// instance of a struct
	inputObj := jsonTest{
		Foo: "fooooooo",
		Bar: 13,
		Baz: 3.14,
		Nan: jsonNestTest{
			A: "a",
			B: 2,
			C: 3.0,
		},
	}
	expectedOutput := `{
    "foo": "fooooooo",
    "bar": 13,
    "baz": 3.14,
    "nan": {
        "A": "a",
        "B": 2,
        "C": 3
    }
}`
	assert.Equal(t, expectedOutput, transformer(inputObj))

	// numbers
	assert.Equal(t, "1", transformer(int(1)))
	assert.Equal(t, "1.2345", transformer(float32(1.2345)))

	// slices
	assert.Equal(t, "[\n    1,\n    2,\n    3\n]", transformer([]uint{1, 2, 3}))

	// strings
	assert.Equal(t, "\"foo\"", transformer("foo"))
	assert.Equal(t, "\"{foo...\"", transformer("{foo...")) // malformed JSON

	// strings with valid JSON
	input := "{\"foo\":\"bar\",\"baz\":[1,2,3]}"
	expectedOutput = `{
    "foo": "bar",
    "baz": [
        1,
        2,
        3
    ]
}`
	assert.Equal(t, expectedOutput, transformer(input))
}

func TestNewTimeTransformer(t *testing.T) {
	inStr := "2010-11-12T13:14:15-07:00"
	inTime, err := time.Parse(time.RFC3339, inStr)
	assert.Nil(t, err)

	location, err := time.LoadLocation("America/Los_Angeles")
	assert.Nil(t, err)
	transformer := NewTimeTransformer(time.RFC3339, location)
	expected := "2010-11-12T12:14:15-08:00"
	assert.Equal(t, expected, transformer(inStr))
	assert.Equal(t, expected, transformer(inTime))
	for _, possibleTimeLayout := range possibleTimeLayouts {
		assert.Equal(t, expected, transformer(inTime.Format(possibleTimeLayout)), possibleTimeLayout)
	}

	location, err = time.LoadLocation("Asia/Singapore")
	assert.Nil(t, err)
	transformer = NewTimeTransformer(time.UnixDate, location)
	expected = "Sat Nov 13 04:14:15 +08 2010"
	assert.Equal(t, expected, transformer(inStr))
	assert.Equal(t, expected, transformer(inTime))
	for _, possibleTimeLayout := range possibleTimeLayouts {
		assert.Equal(t, expected, transformer(inTime.Format(possibleTimeLayout)), possibleTimeLayout)
	}

	location, err = time.LoadLocation("Europe/London")
	assert.Nil(t, err)
	transformer = NewTimeTransformer(time.RFC3339, location)
	expected = "2010-11-12T20:14:15Z"
	assert.Equal(t, expected, transformer(inStr))
	assert.Equal(t, expected, transformer(inTime))
	for _, possibleTimeLayout := range possibleTimeLayouts {
		assert.Equal(t, expected, transformer(inTime.Format(possibleTimeLayout)), possibleTimeLayout)
	}
}

func TestNewUnixTimeTransformer(t *testing.T) {
	inStr := "2010-11-12T13:14:15-07:00"
	inTime, err := time.Parse(time.RFC3339, inStr)
	assert.Nil(t, err)
	inUnixTime := inTime.Unix()

	location, err := time.LoadLocation("America/Los_Angeles")
	assert.Nil(t, err)
	transformer := NewUnixTimeTransformer(time.RFC3339, location)
	expected := "2010-11-12T12:14:15-08:00"
	assert.Equal(t, expected, transformer(fmt.Sprint(inUnixTime)), "seconds in string")
	assert.Equal(t, expected, transformer(inUnixTime), "seconds")
	assert.Equal(t, expected, transformer(inUnixTime*1000), "milliseconds")
	assert.Equal(t, expected, transformer(inUnixTime*1000000), "microseconds")
	assert.Equal(t, expected, transformer(inUnixTime*1000000000), "nanoseconds")

	location, err = time.LoadLocation("Asia/Singapore")
	assert.Nil(t, err)
	transformer = NewUnixTimeTransformer(time.UnixDate, location)
	expected = "Sat Nov 13 04:14:15 +08 2010"
	assert.Equal(t, expected, transformer(fmt.Sprint(inUnixTime)), "seconds in string")
	assert.Equal(t, expected, transformer(inUnixTime), "seconds")
	assert.Equal(t, expected, transformer(inUnixTime*1000), "milliseconds")
	assert.Equal(t, expected, transformer(inUnixTime*1000000), "microseconds")
	assert.Equal(t, expected, transformer(inUnixTime*1000000000), "nanoseconds")

	location, err = time.LoadLocation("Europe/London")
	assert.Nil(t, err)
	transformer = NewUnixTimeTransformer(time.RFC3339, location)
	expected = "2010-11-12T20:14:15Z"
	assert.Equal(t, expected, transformer(fmt.Sprint(inUnixTime)), "seconds in string")
	assert.Equal(t, expected, transformer(inUnixTime), "seconds")
	assert.Equal(t, expected, transformer(inUnixTime*1000), "milliseconds")
	assert.Equal(t, expected, transformer(inUnixTime*1000000), "microseconds")
	assert.Equal(t, expected, transformer(inUnixTime*1000000000), "nanoseconds")

	assert.Equal(t, "0.123456", transformer(float32(0.123456)))
}

func TestNewURLTransformer(t *testing.T) {
	url := "https://winter.is.coming"
	transformer := NewURLTransformer()

	assert.Equal(t, colorsURL.Sprint(url), transformer(url))
}
