package gopretty

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/fatih/color"
)

// Alignment denotes how text is to be aligned horizontally.
type Alignment int

// Alignment enumerations
const (
	AlignmentDefault Alignment = iota // same as AlignmentLeft
	AlignmentLeft                     // "left        "
	AlignmentCenter                   // "   center   "
	AlignmentRight                    // "       right"
)

// Apply aligns the text as directed. Examples:
//  * AlignmentLeft.Apply("Ghost",   7) returns "Ghost  "
//  * AlignmentCenter.Apply("Ghost", 7) returns " Ghost "
//  * AlignmentRight.Apply("Ghost",  7) returns "  Ghost"
func (a Alignment) Apply(text string, maxLength int) string {
	textLength := utf8.RuneCountInString(text)

	formatStr := "%"
	if a == AlignmentCenter {
		if textLength < maxLength {
			text += strings.Repeat(" ", int((maxLength-textLength)/2))
		}
	} else if a == AlignmentDefault || a == AlignmentLeft {
		formatStr += "-"
	}
	formatStr += fmt.Sprint(maxLength)
	formatStr += "s"

	return fmt.Sprintf(formatStr, text)
}

// HTMLProperty returns the equivalent HTML horizontal-alignment tag property.
func (a Alignment) HTMLProperty() string {
	switch a {
	case AlignmentLeft:
		return "align=\"left\""
	case AlignmentCenter:
		return "align=\"center\""
	case AlignmentRight:
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

// VAlignment denotes how text is to be aligned vertically.
type VAlignment int

// VAlignment enumerations
const (
	VAlignmentDefault VAlignment = iota // same as VAlignmentTop
	VAlignmentTop                       // "left"
	VAlignmentMiddle                    // "\nmiddle\n"
	VAlignmentBottom                    // "\n\nright"
)

// Apply aligns the lines vertically. Examples:
//  * VAlignmentTop.Apply({"Game", "Of", "Thrones"},    5)
// 	    returns {"Game", "Of", "Thrones", "", ""}
//  * VAlignmentMiddle.Apply({"Game", "Of", "Thrones"}, 5)
// 	    returns {"", "Game", "Of", "Thrones", ""}
//  * VAlignmentBottom.Apply({"Game", "Of", "Thrones"}, 5)
// 	    returns {"", "", "Game", "Of", "Thrones"}
func (va VAlignment) Apply(lines []string, maxLines int) []string {
	if len(lines) == maxLines {
		return lines
	}

	insertIdx := 0
	if va == VAlignmentMiddle {
		insertIdx = int(maxLines-len(lines)) / 2
	} else if va == VAlignmentBottom {
		insertIdx = maxLines - len(lines)
	}

	linesOut := strings.Split(strings.Repeat("\n", maxLines-1), "\n")
	for idx, line := range lines {
		linesOut[idx+insertIdx] = line
	}
	return linesOut
}

// ApplyStr aligns the string (of 1 or more lines) vertically. Examples:
//  * VAlignmentTop.ApplyStr("Game\nOf\nThrones",    5)
// 	    returns {"Game", "Of", "Thrones", "", ""}
//  * VAlignmentMiddle.ApplyStr("Game\nOf\nThrones", 5)
// 	    returns {"", "Game", "Of", "Thrones", ""}
//  * VAlignmentBottom.ApplyStr("Game\nOf\nThrones", 5)
// 	    returns {"", "", "Game", "Of", "Thrones"}
func (va VAlignment) ApplyStr(text string, maxLines int) []string {
	return va.Apply(strings.Split(text, "\n"), maxLines)
}

// HTMLProperty returns the equivalent HTML vertical-alignment tag property.
func (va VAlignment) HTMLProperty() string {
	switch va {
	case VAlignmentTop:
		return "valign=\"top\""
	case VAlignmentMiddle:
		return "valign=\"middle\""
	case VAlignmentBottom:
		return "valign=\"bottom\""
	default:
		return ""
	}
}
