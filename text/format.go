package text

import (
	"strings"
	"unicode"
)

// Format denotes the "case" to use for text.
type Format int

// Format enumerations
const (
	FormatDefault Format = iota // default_Case
	FormatLower                 // lower
	FormatTitle                 // Title
	FormatUpper                 // UPPER
)

// Apply converts the text as directed.
func (tc Format) Apply(text string) string {
	switch tc {
	case FormatLower:
		return strings.ToLower(text)
	case FormatTitle:
		return tc.toTitle(text)
	case FormatUpper:
		return tc.toUpper(text)
	default:
		return text
	}
}

func (tc Format) toUpper(text string) string {
	inEscSeq := false
	return strings.Map(
		func(r rune) rune {
			if r == EscapeStartRune {
				inEscSeq = true
			}
			if !inEscSeq {
				r = unicode.ToUpper(r)
			}
			if inEscSeq && r == EscapeStopRune {
				inEscSeq = false
			}
			return r
		},
		text,
	)
}

func (tc Format) toTitle(text string) string {
	prev, inEscSeq := ' ', false
	return strings.Map(
		func(r rune) rune {
			if r == EscapeStartRune {
				inEscSeq = true
			}
			if !inEscSeq {
				if tc.isSeparator(prev) {
					prev = r
					r = unicode.ToUpper(r)
				} else {
					prev = r
				}
			}
			if inEscSeq && r == EscapeStopRune {
				inEscSeq = false
			}
			return r
		},
		text,
	)
}

func (tc Format) isSeparator(r rune) bool {
	// ASCII alphanumerics and underscore are not separators
	if r <= 0x7F {
		switch {
		case '0' <= r && r <= '9':
			return false
		case 'a' <= r && r <= 'z':
			return false
		case 'A' <= r && r <= 'Z':
			return false
		case r == '_':
			return false
		}
		return true
	}
	// Letters and digits are not separators
	if unicode.IsLetter(r) || unicode.IsDigit(r) {
		return false
	}
	// Otherwise, all we can do for now is treat spaces as separators.
	return unicode.IsSpace(r)
}
