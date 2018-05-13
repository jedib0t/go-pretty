package text

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/util"
	"strconv"
	"strings"
)

// Colors represents an array of color.Attributes to define the color to
// render text with. Example: Colors{color.FgCyan, color.BgBlack}
type Colors []color.Attribute

var (
	// escapeSeqMap caches the escape sequence for a set of colors
	escapeSeqMap = make(map[string]string)
)

// GetEscapeSeq returns the ANSI escape sequence for the color attributes set.
func (tc Colors) GetEscapeSeq() string {
	if len(tc) == 0 {
		return ""
	}
	colorsKey := fmt.Sprintf("%#v", tc)
	escapeSeq := escapeSeqMap[colorsKey]
	if escapeSeq == "" {
		escapeSeqForColors := make([]string, len(tc))
		for idx, c := range tc {
			escapeSeqForColors[idx] = strconv.Itoa(int(c))
		}
		escapeSeq = util.EscapeStart + strings.Join(escapeSeqForColors, ";") + util.EscapeStop
		escapeSeqMap[colorsKey] = escapeSeq
	}
	return escapeSeq
}

// Sprint colorizes and prints the given string(s).
func (tc Colors) Sprint(a ...interface{}) string {
	return tc.wrap(fmt.Sprint(a...))
}

// Sprintf formats and colorizes and prints the given string(s).
func (tc Colors) Sprintf(format string, a ...interface{}) string {
	return tc.wrap(fmt.Sprintf(format, a...))
}

func (tc Colors) wrap(s string) string {
	escapeSeq := tc.GetEscapeSeq()
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
	return out
}
