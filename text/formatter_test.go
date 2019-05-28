package text

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
)

func TestNewNumberFormatter(t *testing.T) {
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
			formatter := NewNumberFormatter(format)
			expected := signColorsMap[sign].Sprintf(format, value)
			if sign == "negative" {
				expected = strings.Replace(expected, "-0", "-00", 1)
			}
			actual := formatter(value)
			message := fmt.Sprintf("%s.%s: expected=%v, actual=%v; format=%#v",
				sign, reflect.TypeOf(value).Kind(), expected, actual, format)

			assert.Equal(t, expected, actual, message)
		}
	}

	// invalid input
	assert.Equal(t, "foo", NewNumberFormatter("%05d")("foo"))
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

func TestNewJSONFormatter(t *testing.T) {
	formatter := NewJSONFormatter("", "    ")

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
	assert.Equal(t, expectedOutput, formatter(inputObj))

	// numbers
	assert.Equal(t, "1", formatter(int(1)))
	assert.Equal(t, "1.2345", formatter(float32(1.2345)))

	// slices
	assert.Equal(t, "[\n    1,\n    2,\n    3\n]", formatter([]uint{1, 2, 3}))

	// strings
	assert.Equal(t, "\"foo\"", formatter("foo"))
	assert.Equal(t, "\"{foo...\"", formatter("{foo...")) // malformed JSON

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
	assert.Equal(t, expectedOutput, formatter(input))
}

func TestNewTimeFormatter(t *testing.T) {
	inStr := "2010-11-12T13:14:15-07:00"
	inTime, err := time.Parse(time.RFC3339, inStr)
	assert.Nil(t, err)
	inDateTime := strfmt.DateTime(inTime)

	location, err := time.LoadLocation("America/Los_Angeles")
	assert.Nil(t, err)
	formatter := NewTimeFormatter(time.RFC3339, location)
	expected := "2010-11-12T12:14:15-08:00"
	assert.Equal(t, expected, formatter(inStr))
	assert.Equal(t, expected, formatter(inTime))
	assert.Equal(t, expected, formatter(inDateTime))

	location, err = time.LoadLocation("Asia/Singapore")
	assert.Nil(t, err)
	formatter = NewTimeFormatter(time.UnixDate, location)
	expected = "Sat Nov 13 04:14:15 +08 2010"
	assert.Equal(t, expected, formatter(inStr))
	assert.Equal(t, expected, formatter(inTime))
	assert.Equal(t, expected, formatter(inDateTime))

	location, err = time.LoadLocation("Europe/London")
	assert.Nil(t, err)
	formatter = NewTimeFormatter(time.RFC3339, location)
	expected = "2010-11-12T20:14:15Z"
	assert.Equal(t, expected, formatter(inStr))
	assert.Equal(t, expected, formatter(inTime))
	assert.Equal(t, expected, formatter(inDateTime))
}

func TestNewUnixTimeFormatter(t *testing.T) {
	inStr := "2010-11-12T13:14:15-07:00"
	inTime, err := time.Parse(time.RFC3339, inStr)
	assert.Nil(t, err)
	inUnixTime := inTime.Unix()

	location, err := time.LoadLocation("America/Los_Angeles")
	assert.Nil(t, err)
	formatter := NewUnixTimeFormatter(time.RFC3339, location)
	expected := "2010-11-12T12:14:15-08:00"
	assert.Equal(t, expected, formatter(fmt.Sprint(inUnixTime)), "seconds in string")
	assert.Equal(t, expected, formatter(inUnixTime), "seconds")
	assert.Equal(t, expected, formatter(inUnixTime*1000), "milliseconds")
	assert.Equal(t, expected, formatter(inUnixTime*1000000), "microseconds")
	assert.Equal(t, expected, formatter(inUnixTime*1000000000), "nanoseconds")

	location, err = time.LoadLocation("Asia/Singapore")
	assert.Nil(t, err)
	formatter = NewUnixTimeFormatter(time.UnixDate, location)
	expected = "Sat Nov 13 04:14:15 +08 2010"
	assert.Equal(t, expected, formatter(fmt.Sprint(inUnixTime)), "seconds in string")
	assert.Equal(t, expected, formatter(inUnixTime), "seconds")
	assert.Equal(t, expected, formatter(inUnixTime*1000), "milliseconds")
	assert.Equal(t, expected, formatter(inUnixTime*1000000), "microseconds")
	assert.Equal(t, expected, formatter(inUnixTime*1000000000), "nanoseconds")

	location, err = time.LoadLocation("Europe/London")
	assert.Nil(t, err)
	formatter = NewUnixTimeFormatter(time.RFC3339, location)
	expected = "2010-11-12T20:14:15Z"
	assert.Equal(t, expected, formatter(fmt.Sprint(inUnixTime)), "seconds in string")
	assert.Equal(t, expected, formatter(inUnixTime), "seconds")
	assert.Equal(t, expected, formatter(inUnixTime*1000), "milliseconds")
	assert.Equal(t, expected, formatter(inUnixTime*1000000), "microseconds")
	assert.Equal(t, expected, formatter(inUnixTime*1000000000), "nanoseconds")

	assert.Equal(t, "0.123456", formatter(float32(0.123456)))
}

func TestNewURLFormatter(t *testing.T) {
	url := "https://winter.is.coming"
	formatter := NewURLFormatter()

	assert.Equal(t, colorsURL.Sprint(url), formatter(url))
}
