package text

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/jedib0t/go-pretty/util"
)

// Align denotes how text is to be aligned horizontally.
type Align int

// Align enumerations
const (
	AlignDefault Align = iota // same as AlignLeft
	AlignLeft                 // "left        "
	AlignCenter               // "   center   "
	AlignJustify              // "justify   it"
	AlignRight                // "       right"
)

// Apply aligns the text as directed. Examples:
//  * AlignDefault.Apply("Jon Snow", 12) returns "Jon Snow    "
//  * AlignLeft.Apply("Jon Snow",    12) returns "Jon Snow    "
//  * AlignCenter.Apply("Jon Snow",  12) returns "  Jon Snow  "
//  * AlignJustify.Apply("Jon Snow", 12) returns "Jon     Snow"
//  * AlignRight.Apply("Jon Snow",   12) returns "    Jon Snow"
func (a Align) Apply(text string, maxLength int) string {
	if strings.ContainsAny(text, " \t") {
		text = strings.Trim(text, " \t")
	}
	textLength := utf8.RuneCountInString(text)

	if a == AlignJustify {
		return a.justifyText(text, textLength, maxLength)
	}

	// use string formatter to get this done (%10s or %-10s, etc.)
	formatStr := "%"
	if a == AlignCenter {
		// left pad the string with half the number of spaces needed
		if textLength < maxLength {
			text += strings.Repeat(" ", int((maxLength-textLength)/2))
		}
	} else if a == AlignDefault || a == AlignLeft {
		// align left using the "-" in (say) %-10s
		formatStr += "-"
	}
	// finish with the 10s
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
	case AlignJustify:
		return "align=\"justify\""
	case AlignRight:
		return "align=\"right\""
	default:
		return ""
	}
}

func (a Align) justifyText(text string, textLength int, maxLength int) string {
	// split the text into individual words
	wordsUnfiltered := strings.Split(text, " ")
	words := util.FilterStrings(wordsUnfiltered, func(item string) bool {
		return item != ""
	})
	// empty string implies spaces for maxLength
	if len(words) == 0 {
		return strings.Repeat(" ", maxLength)
	}
	// get the number of spaces to insert into the text
	numSpacesNeeded := maxLength - textLength + strings.Count(text, " ")
	numSpacesNeededBetweenWords := 0
	if len(words) > 1 {
		numSpacesNeededBetweenWords = numSpacesNeeded / (len(words) - 1)
	}
	// create the output string word by word with spaces in between
	var outText strings.Builder
	outText.Grow(maxLength)
	for idx, word := range words {
		if idx > 0 {
			// insert spaces only after the first word
			if idx == len(words)-1 {
				// insert all the remaining space before the last word
				outText.WriteString(strings.Repeat(" ", numSpacesNeeded))
				numSpacesNeeded = 0
			} else {
				// insert the determined number of spaces between each word
				outText.WriteString(strings.Repeat(" ", numSpacesNeededBetweenWords))
				// and reduce the number of spaces needed after this
				numSpacesNeeded -= numSpacesNeededBetweenWords
			}
		}
		outText.WriteString(word)
		if idx == len(words)-1 && numSpacesNeeded > 0 {
			outText.WriteString(strings.Repeat(" ", numSpacesNeeded))
		}
	}
	return outText.String()
}
