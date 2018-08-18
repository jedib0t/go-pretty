package text

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

// The logic here is inspired from github.com/fatih/color; the following is
// the the bare minimum logic required to print Colored to the console.
// The differences:
// * This one does not determine if the terminal/console can display Colored
//    it is left to the user to determine using other packages/utilities
// * This one caches the escape sequences for cases with multiple colors
// * This one handles cases where the incoming already has colors in the
//   form of escape sequences; in which case, the that does not have any
//   escape sequences are colored/escaped

// Color represents a single color to render with.
type Color int

// Base colors -- attributes in reality
const (
	Reset Color = iota
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

// Foreground colors
const (
	FgBlack Color = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
)

// Foreground Hi-Intensity colors
const (
	FgHiBlack Color = iota + 90
	FgHiRed
	FgHiGreen
	FgHiYellow
	FgHiBlue
	FgHiMagenta
	FgHiCyan
	FgHiWhite
)

// Background colors
const (
	BgBlack Color = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
)

// Background Hi-Intensity colors
const (
	BgHiBlack Color = iota + 100
	BgHiRed
	BgHiGreen
	BgHiYellow
	BgHiBlue
	BgHiMagenta
	BgHiCyan
	BgHiWhite
)

// GetEscapeSeq returns the ANSI escape sequence for the color.
func (c Color) GetEscapeSeq() string {
	return EscapeStart + strconv.Itoa(int(c)) + EscapeStop
}

// Sprint colorizes and prints the given string(s).
func (c Color) Sprint(a ...interface{}) string {
	return colorize(fmt.Sprint(a...), c.GetEscapeSeq())
}

// Sprintf formats and colorizes and prints the given string(s).
func (c Color) Sprintf(format string, a ...interface{}) string {
	return colorize(fmt.Sprintf(format, a...), c.GetEscapeSeq())
}

// Colors represents an array of Color objects to render with.
// Example: Colors{FgCyan, BgBlack}
type Colors []Color

var (
	// colorsSeqMap caches the escape sequence for a set of colors
	colorsSeqMap = make(map[string]string)

	// colorsSeqMapMutex should be used to lock writes to the map
	colorsSeqMapMutex = sync.Mutex{}
)

// GetEscapeSeq returns the ANSI escape sequence for the colors set.
func (c Colors) GetEscapeSeq() string {
	if len(c) == 0 {
		return ""
	}
	colorsKey := fmt.Sprintf("%#v", c)
	escapeSeq := colorsSeqMap[colorsKey]
	if escapeSeq == "" {
		colorNums := make([]string, len(c))
		for idx, c := range c {
			colorNums[idx] = strconv.Itoa(int(c))
		}
		escapeSeq = EscapeStart + strings.Join(colorNums, ";") + EscapeStop
		colorsSeqMapMutex.Lock()
		colorsSeqMap[colorsKey] = escapeSeq
		colorsSeqMapMutex.Unlock()
	}
	return escapeSeq
}

// Sprint colorizes and prints the given string(s).
func (c Colors) Sprint(a ...interface{}) string {
	return colorize(fmt.Sprint(a...), c.GetEscapeSeq())
}

// Sprintf formats and colorizes and prints the given string(s).
func (c Colors) Sprintf(format string, a ...interface{}) string {
	return colorize(fmt.Sprintf(format, a...), c.GetEscapeSeq())
}

func colorize(s string, escapeSeq string) string {
	if !ANSISupported || escapeSeq == "" {
		return s
	}

	out := ""
	if !strings.HasPrefix(s, EscapeStart) {
		out += escapeSeq
	}
	out += strings.Replace(s, EscapeReset, EscapeReset+escapeSeq, -1)
	if !strings.HasSuffix(out, EscapeReset) {
		out += EscapeReset
	}
	if strings.Contains(out, escapeSeq+EscapeReset) {
		out = strings.Replace(out, escapeSeq+EscapeReset, "", -1)
	}
	return out
}
