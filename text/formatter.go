package text

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/go-openapi/strfmt"
)

// Formatter related constants
const (
	unixTimeMinMilliseconds = int64(10000000000)
	unixTimeMinMicroseconds = unixTimeMinMilliseconds * 1000
	unixTimeMinNanoSeconds  = unixTimeMinMicroseconds * 1000
)

// Formatter related variables
var (
	colorsNumberPositive = Colors{FgHiGreen}
	colorsNumberNegative = Colors{FgHiRed}
	colorsNumberZero     = Colors{}
	colorsURL            = Colors{Underline, FgBlue}
)

// Formatter helps format the contents of a Column to the user's liking.
type Formatter func(val interface{}) string

// NewNumberFormatter returns a number Formatter that:
//   * formats the number as directed by 'format' (ex.: %.2f)
//   * colors negative values Red
//   * colors positive values Green
func NewNumberFormatter(format string) Formatter {
	return func(val interface{}) string {
		if number, ok := val.(int); ok {
			return formatInt(format, int64(number))
		}
		if number, ok := val.(int8); ok {
			return formatInt(format, int64(number))
		}
		if number, ok := val.(int16); ok {
			return formatInt(format, int64(number))
		}
		if number, ok := val.(int32); ok {
			return formatInt(format, int64(number))
		}
		if number, ok := val.(int64); ok {
			return formatInt(format, int64(number))
		}
		if number, ok := val.(uint); ok {
			return formatUint(format, uint64(number))
		}
		if number, ok := val.(uint8); ok {
			return formatUint(format, uint64(number))
		}
		if number, ok := val.(uint16); ok {
			return formatUint(format, uint64(number))
		}
		if number, ok := val.(uint32); ok {
			return formatUint(format, uint64(number))
		}
		if number, ok := val.(uint64); ok {
			return formatUint(format, uint64(number))
		}
		if number, ok := val.(float32); ok {
			return formatFloat(format, float64(number))
		}
		if number, ok := val.(float64); ok {
			return formatFloat(format, float64(number))
		}
		return fmt.Sprint(val)
	}
}

func formatInt(format string, val int64) string {
	if val < 0 {
		return colorsNumberNegative.Sprintf("-"+format, -val)
	}
	if val > 0 {
		return colorsNumberPositive.Sprintf(format, val)
	}
	return colorsNumberZero.Sprintf(format, val)
}

func formatUint(format string, val uint64) string {
	if val > 0 {
		return colorsNumberPositive.Sprintf(format, val)
	}
	return colorsNumberZero.Sprintf(format, val)
}

func formatFloat(format string, val float64) string {
	if val < 0 {
		return colorsNumberNegative.Sprintf("-"+format, -val)
	}
	if val > 0 {
		return colorsNumberPositive.Sprintf(format, val)
	}
	return colorsNumberZero.Sprintf(format, val)
}

// NewJSONFormatter returns a Formatter that can format a JSON string or an
// object into pretty-indented JSON-strings.
func NewJSONFormatter(prefix string, indent string) Formatter {
	return func(val interface{}) string {
		if valStr, ok := val.(string); ok {
			var b bytes.Buffer
			if err := json.Indent(&b, []byte(strings.TrimSpace(valStr)), prefix, indent); err == nil {
				return string(b.Bytes())
			}
		} else if b, err := json.MarshalIndent(val, prefix, indent); err == nil {
			return string(b)
		}
		return fmt.Sprintf("%#v", val)
	}
}

// NewTimeFormatter returns a Formatter that can format a timestamp (a time.Time
// or strfmt.DateTime object) into a well-defined time format defined using
// the provided layout (ex.: time.RFC3339).
//
// If a non-nil location value is provided, the time will be localized to that
// location (use time.Local to get localized timestamps).
func NewTimeFormatter(layout string, location *time.Location) Formatter {
	return func(val interface{}) string {
		formatTime := func(t time.Time) string {
			rsp := ""
			if t.Unix() > 0 {
				if location != nil {
					t = t.In(location)
				}
				rsp = t.Format(layout)
			}
			return rsp
		}

		rsp := fmt.Sprint(val)
		if valDate, ok := val.(strfmt.DateTime); ok {
			rsp = formatTime(time.Time(valDate))
		} else if valTime, ok := val.(time.Time); ok {
			rsp = formatTime(valTime)
		} else if valStr, ok := val.(string); ok {
			if valTime, err := time.Parse(time.RFC3339, valStr); err == nil {
				rsp = formatTime(valTime)
			}
		}
		return rsp
	}
}

// NewUnixTimeFormatter returns a Formatter that can format a unix-timestamp
// into a well-defined time format as defined by 'layout'. This can handle
// unix-time in Seconds, MilliSeconds, Microseconds and Nanoseconds.
//
// If a non-nil location value is provided, the time will be localized to that
// location (use time.Local to get localized timestamps).
func NewUnixTimeFormatter(layout string, location *time.Location) Formatter {
	timeFormatter := NewTimeFormatter(layout, location)
	formatUnixTime := func(unixTime int64) string {
		if unixTime >= unixTimeMinNanoSeconds {
			unixTime = unixTime / time.Second.Nanoseconds()
		} else if unixTime >= unixTimeMinMicroseconds {
			unixTime = unixTime / (time.Second.Nanoseconds() / 1000)
		} else if unixTime >= unixTimeMinMilliseconds {
			unixTime = unixTime / (time.Second.Nanoseconds() / 1000000)
		}
		return timeFormatter(time.Unix(unixTime, 0))
	}

	return func(val interface{}) string {
		if unixTime, ok := val.(int64); ok {
			return formatUnixTime(unixTime)
		} else if unixTimeStr, ok := val.(string); ok {
			if unixTime, err := strconv.ParseInt(unixTimeStr, 10, 64); err == nil {
				return formatUnixTime(unixTime)
			}
		}
		return fmt.Sprint(val)
	}
}

// NewURLFormatter returns a Formatter that can format and pretty print a string
// that contains an URL (the text is underlined and colored Blue).
func NewURLFormatter() Formatter {
	return func(val interface{}) string {
		return colorsURL.Sprint(val)
	}
}
