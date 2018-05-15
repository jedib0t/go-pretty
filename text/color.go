package text

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jedib0t/go-pretty/util"
)

// The logic here is inspired from github.com/fatih/color; the following is
// the the bare minimum logic required to print Colored text to the console.
// The differences:
// * This one does not determine if the terminal/console can display Colored
//   text; it is left to the user to determine using other packages/utilities
// * This one caches the escape sequences for cases with multiple colors
// * This one handles cases where the incoming text already has colors in the
//   form of escape sequences; in which case, the text that does not have any
//   escape sequences are colored/escaped

// Color represents a single color to render text with.
type Color int

// Base colors -- attributes in reality
const (
	Reset        Color = iota
	Bold
	Faint
	Italic
	Underline
	BlinkSlow
	BlinkRapid
	ReverseVideo
	Concealed
	CrossedOut
)

// Foreground text colors
const (
	FgBlack   Color = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
)

// Foreground Hi-Intensity text colors
const (
	FgHiBlack   Color = iota + 90
	FgHiRed
	FgHiGreen
	FgHiYellow
	FgHiBlue
	FgHiMagenta
	FgHiCyan
	FgHiWhite
)

// Background text colors
const (
	BgBlack   Color = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
)

// Background Hi-Intensity text colors
const (
	BgHiBlack   Color = iota + 100
	BgHiRed
	BgHiGreen
	BgHiYellow
	BgHiBlue
	BgHiMagenta
	BgHiCyan
	BgHiWhite
)

// GetEscapeSeq returns the ANSI escape sequence for the color.
func (ca Color) GetEscapeSeq() string {
	return util.EscapeStart + strconv.Itoa(int(ca)) + util.EscapeStop
}

// Sprint colorizes and prints the given string(s).
func (ca Color) Sprint(a ...interface{}) string {
	return ca.colorize(fmt.Sprint(a...))
}

// Sprintf formats and colorizes and prints the given string(s).
func (ca Color) Sprintf(format string, a ...interface{}) string {
	return ca.colorize(fmt.Sprintf(format, a...))
}

func (ca Color) colorize(s string) string {
	escapeSeq := ca.GetEscapeSeq()

	out := ""
	if !strings.HasPrefix(s, util.EscapeStart) {
		out += escapeSeq
	}
	out += strings.Replace(s, util.EscapeReset, util.EscapeReset+escapeSeq, -1)
	if !strings.HasSuffix(out, util.EscapeReset) {
		out += util.EscapeReset
	}
	if strings.Contains(out, escapeSeq+util.EscapeReset) {
		out = strings.Replace(out, escapeSeq+util.EscapeReset, "", -1)
	}
	return out
}

// Colors represents an array of Color objects to render text with.
// Example: Colors{FgCyan, BgBlack}
type Colors []Color

var (
	// colorsSeqMap caches the escape sequence for a set of colors
	colorsSeqMap = make(map[string]string)
)

// GetEscapeSeq returns the ANSI escape sequence for the colors set.
func (ca Colors) GetEscapeSeq() string {
	if len(ca) == 0 {
		return ""
	}
	colorsKey := fmt.Sprintf("%#v", ca)
	escapeSeq := colorsSeqMap[colorsKey]
	if escapeSeq == "" {
		escapeSeqForColors := make([]string, len(ca))
		for idx, c := range ca {
			escapeSeqForColors[idx] = strconv.Itoa(int(c))
		}
		escapeSeq = util.EscapeStart + strings.Join(escapeSeqForColors, ";") + util.EscapeStop
		colorsSeqMap[colorsKey] = escapeSeq
	}
	return escapeSeq
}

// Sprint colorizes and prints the given string(s).
func (ca Colors) Sprint(a ...interface{}) string {
	return ca.colorize(fmt.Sprint(a...))
}

// Sprintf formats and colorizes and prints the given string(s).
func (ca Colors) Sprintf(format string, a ...interface{}) string {
	return ca.colorize(fmt.Sprintf(format, a...))
}

func (ca Colors) colorize(s string) string {
	escapeSeq := ca.GetEscapeSeq()
	if escapeSeq == "" {
		return s
	}

	out := ""
	if !strings.HasPrefix(s, util.EscapeStart) {
		out += escapeSeq
	}
	out += strings.Replace(s, util.EscapeReset, util.EscapeReset+escapeSeq, -1)
	if !strings.HasSuffix(out, util.EscapeReset) {
		out += util.EscapeReset
	}
	if strings.Contains(out, escapeSeq+util.EscapeReset) {
		out = strings.Replace(out, escapeSeq+util.EscapeReset, "", -1)
	}
	return out
}
