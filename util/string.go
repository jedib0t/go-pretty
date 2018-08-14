package util

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// Constants
const (
	EscapeReset     = EscapeStart + "0" + EscapeStop
	EscapeStart     = "\x1b["
	EscapeStartRune = rune(27) // \x1b
	EscapeStop      = "m"
	EscapeStopRune  = 'm'
)

// FixedLengthString returns the given string with a fixed length. For ex.:
//  FixedLengthString("Ghost", 0, "~") == "Ghost"
//  FixedLengthString("Ghost", 1, "~") == "~"
//  FixedLengthString("Ghost", 3, "~") == "Gh~"
//  FixedLengthString("Ghost", 5, "~") == "Ghost"
//  FixedLengthString("Ghost", 7, "~") == "Ghost  "
func FixedLengthString(s string, length int, snipIndicator string) string {
	if length > 0 {
		lenStr := RuneCountWithoutEscapeSeq(s)
		if lenStr < length {
			return fmt.Sprintf("%-"+fmt.Sprint(length)+"s", s)
		} else if lenStr > length {
			lenStrFinal := length - RuneCountWithoutEscapeSeq(snipIndicator)
			return TrimTextWithoutEscapeSeq(s, lenStrFinal) + snipIndicator
		}
	}
	return s
}

// GetLongestLineLength returns the length of the longest "line" within the
// argument string. For ex.:
//  GetLongestLineLength("Ghost!\nCome back here!\nRight now!") == 15
func GetLongestLineLength(s string) int {
	maxLength, currLength, isEscSeq := 0, 0, false
	for _, c := range s {
		if c == EscapeStartRune {
			isEscSeq = true
		} else if isEscSeq && c == EscapeStopRune {
			isEscSeq = false
			continue
		}

		if c == '\n' {
			if currLength > maxLength {
				maxLength = currLength
			}
			currLength = 0
		} else if !isEscSeq {
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
func InsertRuneEveryN(s string, r rune, n int) string {
	sLen := RuneCountWithoutEscapeSeq(s)
	var out strings.Builder
	out.Grow(sLen + (sLen / n))

	outLen, isEscSeq := 0, false
	for idx, c := range s {
		if c == EscapeStartRune {
			isEscSeq = true
		}

		if !isEscSeq && outLen > 0 && (outLen%n) == 0 && idx != sLen {
			out.WriteRune(r)
		}
		out.WriteRune(c)
		if !isEscSeq {
			outLen++
		}

		if isEscSeq && c == EscapeStopRune {
			isEscSeq = false
		}
	}
	return out.String()
}

// RepeatAndTrim repeats the given string until it is as long as maxRunes.
// For ex.:
//  RepeatAndTrim("Ghost", 0) == ""
//  RepeatAndTrim("Ghost", 5) == "Ghost"
//  RepeatAndTrim("Ghost", 7) == "GhostGh"
//  RepeatAndTrim("Ghost", 10) == "GhostGhost"
func RepeatAndTrim(s string, maxRunes int) string {
	if maxRunes == 0 {
		return ""
	} else if maxRunes == len(s) {
		return s
	}
	return TrimTextWithoutEscapeSeq(strings.Repeat(s, int(maxRunes/utf8.RuneCountInString(s))+1), maxRunes)
}

// RuneCountWithoutEscapeSeq is similar to utf8.RuneCountInString, except for
// the fact that it ignores escape sequences in the counting. For ex.:
//  RuneCountWithoutEscapeSeq("") == 0
//  RuneCountWithoutEscapeSeq("Ghost") == 5
//  RuneCountWithoutEscapeSeq("\x1b[33mGhost\x1b[0m") == 5
//  RuneCountWithoutEscapeSeq("\x1b[33mGhost\x1b[0") == 5
func RuneCountWithoutEscapeSeq(s string) int {
	count, isEscSeq := 0, false
	for _, c := range s {
		if c == EscapeStartRune {
			isEscSeq = true
		} else if isEscSeq {
			if c == EscapeStopRune {
				isEscSeq = false
			}
		} else {
			count++
		}
	}
	return count
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
		if sChr == EscapeStartRune {
			isEscSeq = true
			lastEscSeq.Reset()
			lastEscSeq.WriteRune(sChr)
		} else if isEscSeq {
			lastEscSeq.WriteRune(sChr)
			if sChr == EscapeStopRune {
				isEscSeq = false
			}
		} else {
			outLen++
			if outLen == n {
				break
			}
		}
	}
	if lastEscSeq.Len() > 0 && lastEscSeq.String() != EscapeReset {
		out.WriteString(EscapeReset)
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
	lineIdx, isEscSeq := 0, false
	for idx, c := range s {
		if c == EscapeStartRune {
			isEscSeq = true
		}

		if !isEscSeq && lineIdx == n && c != '\n' {
			out.WriteRune('\n')
			lineIdx = 0
		}
		out.WriteRune(c)
		if c == '\n' {
			lineIdx = 0
		} else if !isEscSeq && idx < sLen {
			lineIdx++
		}

		if isEscSeq && c == EscapeStopRune {
			isEscSeq = false
		}
	}
	return out.String()
}
