package util

import (
	"strings"
	"unicode/utf8"
)

// Constants
const (
	colorReset  = "\x1b[0m"
	escapeStart = rune(27) // \x1b
	escapeStop  = 'm'
)

// GetLongestLineLength returns the length of the longest "line" within the
// argument string. For ex.:
//  GetLongestLineLength("Ghost!\nCome back here!\nRight now!") == 15
func GetLongestLineLength(s string) int {
	maxLength, currLength := 0, 0
	for _, c := range s {
		if c == '\n' {
			if currLength > maxLength {
				maxLength = currLength
			}
			currLength = 0
		} else {
			currLength++
		}
	}
	if currLength > maxLength {
		return currLength
	}
	return maxLength
}

// InsertRuneEveryN inserts the rune every N characters in the string. For ex.:
//  InsertRuneEveryN("Ghost", '-', 1) == "G-h-o-s-t"
//  InsertRuneEveryN("Ghost", '-', 2) == "Gh-os-t"
//  InsertRuneEveryN("Ghost", '-', 3) == "Gho-st"
//  InsertRuneEveryN("Ghost", '-', 4) == "Ghos-t"
//  InsertRuneEveryN("Ghost", '-', 5) == "Ghost"
func InsertRuneEveryN(s string, c rune, n int) string {
	sLen := utf8.RuneCountInString(s)
	var out strings.Builder
	out.Grow(sLen + (sLen / n))

	for scIdx, sc := range s {
		out.WriteRune(sc)
		if ((scIdx+1)%n) == 0 && scIdx != (sLen-1) {
			out.WriteRune(c)
		}
	}
	return out.String()
}

// RuneCountWithoutEscapeSeq is similar to utf8.RuneCountInString, except for
// the fact that it ignores escape sequences in the counting. For ex.:
//  RuneCountWithoutEscapeSeq("") == 0
//  RuneCountWithoutEscapeSeq("Ghost") == 5
//  RuneCountWithoutEscapeSeq("\x1b[33mGhost\x1b[0m") == 5
//  RuneCountWithoutEscapeSeq("\x1b[33mGhost\x1b[0") == 5
func RuneCountWithoutEscapeSeq(s string) int {
	out, isEscSeq := 0, false
	for _, sChr := range s {
		if sChr == escapeStart {
			isEscSeq = true
		} else if isEscSeq {
			if sChr == escapeStop {
				isEscSeq = false
			}
		} else {
			out++
		}
	}
	return out
}

// TrimTextWithoutEscapeSeq trims a string to the given length accounting for
// escape sequences For ex.:
//  TrimTextWithoutEscapeSeq("Ghost", 3) == "Gho"
//  TrimTextWithoutEscapeSeq("Ghost", 6) == "Ghost"
//  TrimTextWithoutEscapeSeq("\x1b[33mGhost\x1b[0m", 3) == "\x1b[33mGho\x1b[0m"
//  TrimTextWithoutEscapeSeq("\x1b[33mGhost\x1b[0m", 6) == "\x1b[33mGhost\x1b[0m"
func TrimTextWithoutEscapeSeq(s string, n int) string {
	if n <= 0 {
		return ""
	}

	var out strings.Builder
	out.Grow(n)

	outLen, isEscSeq, lastEscSeq := 0, false, strings.Builder{}
	for _, sChr := range s {
		out.WriteRune(sChr)
		if sChr == escapeStart {
			isEscSeq = true
			lastEscSeq.Reset()
			lastEscSeq.WriteRune(sChr)
		} else if isEscSeq {
			lastEscSeq.WriteRune(sChr)
			if sChr == escapeStop {
				isEscSeq = false
			}
		} else {
			outLen++
			if outLen == n {
				break
			}
		}
	}
	if lastEscSeq.Len() > 0 && lastEscSeq.String() != colorReset {
		out.WriteString(colorReset)
	}
	return out.String()
}

// WrapText wraps a string to the given length using a newline. For ex.:
//  WrapText("Ghost", 1) == "G\nh\no\ns\nt"
//  WrapText("Ghost", 2) == "Gh\nos\nt"
//  WrapText("Ghost", 3) == "Gho\nst"
//  WrapText("Ghost", 4) == "Ghos\nt"
//  WrapText("Ghost", 5) == "Ghost"
//  WrapText("Ghost", 6) == "Ghost"
//  WrapText("Jon\nSnow", 2) == "Jo\nn\nSn\now"
//  WrapText("Jon\nSnow\n", 2) == "Jo\nn\nSn\now\n"
func WrapText(s string, n int) string {
	if n <= 0 {
		return ""
	}

	var out strings.Builder
	sLen := utf8.RuneCountInString(s)
	out.Grow(sLen + (sLen / n))
	sLineIdx := 0
	for sIdx, sChr := range s {
		if sLineIdx == n && sChr != '\n' {
			out.WriteRune('\n')
			sLineIdx = 0
		}
		out.WriteRune(sChr)
		if sChr == '\n' {
			sLineIdx = 0
		} else if sIdx < (sLen - 1) {
			sLineIdx++
		}
	}
	return out.String()
}
