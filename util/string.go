package util

import (
	"strings"
	"unicode/utf8"
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

// WrapText wraps a text to the given length using a newline. For ex.:
//  WrapText("Ghost", 1) == "G\nh\no\ns\nt"
//  WrapText("Ghost", 2) == "Gh\nos\nt"
//  WrapText("Ghost", 3) == "Gho\nst"
//  WrapText("Ghost", 4) == "Ghos\nt"
//  WrapText("Ghost", 5) == "Ghost"
//  WrapText("Ghost", 6) == "Ghost"
//  WrapText("Jon\nSnow", 2) == "Jo\nn\nSn\now"
//  WrapText("Jon\nSnow\n", 2) == "Jo\nn\nSn\now\n"
func WrapText(s string, n int) string {
	sLen := utf8.RuneCountInString(s)

	var out strings.Builder
	out.Grow(sLen + (sLen / n))
	sLineIdx := 0
	for sIdx, sChr := range s {
		if sLineIdx == n {
			if sIdx == (sLen-1) && sChr == '\n' {
				// last letter and it is a newline; don't add one more
			} else {
				out.WriteRune('\n')
				sLineIdx = 0
			}
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
