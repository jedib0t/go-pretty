package gopretty

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/fatih/color"
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

// TextCase denotes the "case" to use for text.
type TextCase int

// TextCase enumerations
const (
	TextCaseDefault TextCase = iota // default_Case
	TextCaseLower                   // lower
	TextCaseTitle                   // Title
	TextCaseUpper                   // UPPER
)

// Convert converts the text as directed.
func (tc TextCase) Convert(text string) string {
	switch tc {
	case TextCaseLower:
		return strings.ToLower(text)
	case TextCaseTitle:
		return strings.Title(text)
	case TextCaseUpper:
		return strings.ToUpper(text)
	default:
		return text
	}
}

// TextColor represents an array of color.Attributes to define the color to
// render text with. Example: TextColor{color.FgCyan, color.BgBlack}
type TextColor []color.Attribute

// getColorizer returns a *color.Color object based on the color attributes set.
func (tc TextColor) getColorizer() *color.Color {
	var colorizer *color.Color
	if len(tc) > 0 {
		colorizer = color.New(tc...)
		colorizer.EnableColor()
	}
	return colorizer
}

// Sprint colorizes and prints the given string(s).
func (tc TextColor) Sprint(a ...interface{}) string {
	colorizer := tc.getColorizer()
	if colorizer != nil {
		return colorizer.Sprint(a...)
	}
	return fmt.Sprint(a...)
}

// Sprintf formats and colorizes and prints the given string(s).
func (tc TextColor) Sprintf(format string, a ...interface{}) string {
	colorizer := tc.getColorizer()
	if colorizer != nil {
		return colorizer.Sprintf(format, a...)
	}
	return fmt.Sprintf(format, a...)
}

// VAlign denotes how text is to be aligned vertically.
type VAlign int

// VAlign enumerations
const (
	VAlignDefault VAlign = iota // same as VAlignTop
	VAlignTop                   // "top\n\n"
	VAlignMiddle                // "\nmiddle\n"
	VAlignBottom                // "\n\nbottom"
)

// Apply aligns the lines vertically. Examples:
//  * VAlignTop.Apply({"Game", "Of", "Thrones"},    5)
// 	    returns {"Game", "Of", "Thrones", "", ""}
//  * VAlignMiddle.Apply({"Game", "Of", "Thrones"}, 5)
// 	    returns {"", "Game", "Of", "Thrones", ""}
//  * VAlignBottom.Apply({"Game", "Of", "Thrones"}, 5)
// 	    returns {"", "", "Game", "Of", "Thrones"}
func (va VAlign) Apply(lines []string, maxLines int) []string {
	if len(lines) == maxLines {
		return lines
	}

	insertIdx := 0
	if va == VAlignMiddle {
		insertIdx = int(maxLines-len(lines)) / 2
	} else if va == VAlignBottom {
		insertIdx = maxLines - len(lines)
	}

	linesOut := strings.Split(strings.Repeat("\n", maxLines-1), "\n")
	for idx, line := range lines {
		linesOut[idx+insertIdx] = line
	}
	return linesOut
}

// ApplyStr aligns the string (of 1 or more lines) vertically. Examples:
//  * VAlignTop.ApplyStr("Game\nOf\nThrones",    5)
// 	    returns {"Game", "Of", "Thrones", "", ""}
//  * VAlignMiddle.ApplyStr("Game\nOf\nThrones", 5)
// 	    returns {"", "Game", "Of", "Thrones", ""}
//  * VAlignBottom.ApplyStr("Game\nOf\nThrones", 5)
// 	    returns {"", "", "Game", "Of", "Thrones"}
func (va VAlign) ApplyStr(text string, maxLines int) []string {
	return va.Apply(strings.Split(text, "\n"), maxLines)
}

// HTMLProperty returns the equivalent HTML vertical-align tag property.
func (va VAlign) HTMLProperty() string {
	switch va {
	case VAlignTop:
		return "valign=\"top\""
	case VAlignMiddle:
		return "valign=\"middle\""
	case VAlignBottom:
		return "valign=\"bottom\""
	default:
		return ""
	}
}
