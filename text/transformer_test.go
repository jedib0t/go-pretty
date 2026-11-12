package text

import (
	"fmt"
	"math"
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
		"nil":      nil,
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
		"nil": {
			nil: "%v",
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
			var kind reflect.Kind
			if value != nil {
				kind = reflect.TypeOf(value).Kind()
			}
			message := fmt.Sprintf("%s.%s: expected=%v, actual=%v; format=%#v",
				sign, kind, expected, actual, format)

			assert.Equal(t, expected, actual, message)
		}
	}

	// invalid input
	assert.Equal(t, "foo", NewNumberTransformer("%05d")("foo"))
}

func TestNewNumberTransformerMinInt(t *testing.T) {
	assert.Equal(t, colorsNumberNegative.Sprint(fmt.Sprint(math.MinInt)), NewNumberTransformer("%d")(int(math.MinInt)))
}

func TestNewNumberTransformerMinInt64(t *testing.T) {
	for _, tc := range []struct{ format, want string }{
		{"%d", "-9223372036854775808"},
		{"%020d", "-09223372036854775808"},
		{"%x", "-8000000000000000"},
	} {
		t.Run(tc.format, func(t *testing.T) {
			assert.Equal(t, colorsNumberNegative.Sprint(tc.want), NewNumberTransformer(tc.format)(int64(math.MinInt64)))
		})
	}
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

	// Test error case: object that cannot be marshaled (channel) - triggers error path in json.MarshalIndent
	ch := make(chan int)
	result := transformer(ch)
	assert.Contains(t, result, "chan int")

	// Test error case: function that cannot be marshaled - triggers error path in json.MarshalIndent
	fn := func() {}
	result = transformer(fn)
	assert.Contains(t, result, "func()")
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

	// Test case where string doesn't match any time layout (returns original string)
	transformer = NewTimeTransformer(time.RFC3339, nil)
	nonTimeString := "not a time string"
	assert.Equal(t, nonTimeString, transformer(nonTimeString))

	// Test with nil location
	transformer = NewTimeTransformer(time.RFC3339, nil)
	assert.Equal(t, "2010-11-12T13:14:15-07:00", transformer(inTime))

	// a time before the epoch is a value like any other
	transformer = NewTimeTransformer(time.RFC3339, time.UTC)
	assert.Equal(t, "1960-01-01T00:00:00Z", transformer(time.Date(1960, time.January, 1, 0, 0, 0, 0, time.UTC)))
	assert.Equal(t, "1969-12-31T23:59:59Z", transformer(time.Unix(-1, 0)))
	assert.Equal(t, "1955-02-24T10:30:00Z", transformer("1955-02-24T10:30:00Z"))
	assert.Equal(t, "1970-01-01T00:00:00.5Z", NewTimeTransformer(time.RFC3339Nano, time.UTC)(time.Unix(0, 500000000)))
	// the zero value and the epoch are a time that was never set
	assert.Equal(t, "", transformer(time.Time{}))
	assert.Equal(t, "", transformer(time.Unix(0, 0)))
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

	// timestamps before the epoch, in every unit
	transformer = NewUnixTimeTransformer(time.RFC3339, time.UTC)
	inUnixTime = time.Date(1955, time.February, 24, 10, 30, 0, 0, time.UTC).Unix()
	expected = "1955-02-24T10:30:00Z"
	assert.Equal(t, expected, transformer(fmt.Sprint(inUnixTime)), "seconds in string")
	assert.Equal(t, expected, transformer(inUnixTime), "seconds")
	assert.Equal(t, expected, transformer(inUnixTime*1000), "milliseconds")
	assert.Equal(t, expected, transformer(inUnixTime*1000000), "microseconds")
	assert.Equal(t, expected, transformer(inUnixTime*1000000000), "nanoseconds")
	assert.Equal(t, "1969-12-31T23:59:59Z", transformer(int64(-1)))
	assert.Equal(t, "1677-09-21T00:12:43Z", transformer(int64(math.MinInt64)))
	// the epoch is a time that was never set
	assert.Equal(t, "", transformer(int64(0)))
	assert.Equal(t, "", transformer("0"))
}

func TestNewUnixTimeTransformerSubseconds(t *testing.T) {
	instant := time.Date(2010, time.November, 12, 20, 14, 15, 123456789, time.UTC)
	transformer := NewUnixTimeTransformer(time.RFC3339Nano, time.UTC)
	for _, unit := range []time.Duration{time.Second, time.Millisecond, time.Microsecond, time.Nanosecond} {
		timestamp := instant.UnixNano() / int64(unit)
		expected := instant.Truncate(unit).Format(time.RFC3339Nano)
		t.Run(unit.String(), func(t *testing.T) {
			assert.Equal(t, expected, transformer(timestamp))
			assert.Equal(t, expected, transformer(fmt.Sprint(timestamp)))
		})
	}

	// before the epoch, the fraction counts back from the second after it
	for _, test := range []struct {
		timestamp int64
		expected  string
	}{
		{-15000000500, "1969-07-11T09:19:59.5Z"},
		{-15000000000500, "1969-07-11T09:19:59.9995Z"},
		{-15000000000000500, "1969-07-11T09:19:59.9999995Z"},
	} {
		assert.Equal(t, test.expected, transformer(test.timestamp))
		assert.Equal(t, test.expected, transformer(fmt.Sprint(test.timestamp)))
	}
}

func TestNewURLTransformer(t *testing.T) {
	url := "https://winter.is.coming"

	transformer := NewURLTransformer()
	assert.Equal(t, colorsURL.Sprint(url), transformer(url))

	transformer2 := NewURLTransformer(FgRed, BgWhite, Bold)
	assert.Equal(t, Colors{FgRed, BgWhite, Bold}.Sprint(url), transformer2(url))
	assert.Equal(t, colorsURL.Sprint(url), transformer(url))
}
