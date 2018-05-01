package text

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// Align denotes how text is to be aligned horizontally.
type Align int

// Align enumerations
const (
	AlignDefault Align = iota // same as AlignLeft
	AlignLeft                 // "left        "
	AlignCenter               // "   center   "
	AlignRight                // "       right"
)

// Apply aligns the text as directed. Examples:
//  * AlignLeft.Apply("Ghost",   7) returns "Ghost  "
//  * AlignCenter.Apply("Ghost", 7) returns " Ghost "
//  * AlignRight.Apply("Ghost",  7) returns "  Ghost"
func (a Align) Apply(text string, maxLength int) string {
	textLength := utf8.RuneCountInString(text)

	formatStr := "%"
	if a == AlignCenter {
		if textLength < maxLength {
			text += strings.Repeat(" ", int((maxLength-textLength)/2))
		}
	} else if a == AlignDefault || a == AlignLeft {
		formatStr += "-"
	}
	formatStr += fmt.Sprint(maxLength)
	formatStr += "s"

	return fmt.Sprintf(formatStr, text)
}

// HTMLProperty returns the equivalent HTML horizontal-align tag property.
func (a Align) HTMLProperty() string {
	switch a {
	case AlignLeft:
		return "align=\"left\""
	case AlignCenter:
		return "align=\"center\""
	case AlignRight:
		return "align=\"right\""
	default:
		return ""
	}
}
