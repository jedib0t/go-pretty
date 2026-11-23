package text

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode/utf8"
)

// Constants
const (
	EscapeReset        = EscapeResetCSI
	EscapeResetCSI     = EscapeStartCSI + "0" + EscapeStopCSI
	EscapeResetOSI     = EscapeStartOSI + "0" + EscapeStopOSI
	EscapeStart        = EscapeStartCSI
	EscapeStartCSI     = "\x1b["
	EscapeStartOSI     = "\x1b]"
	EscapeStartRune    = rune(27) // \x1b
	EscapeStartRuneCSI = '['      // [
	EscapeStartRuneOSI = ']'      // ]
	EscapeStop         = EscapeStopCSI
	EscapeStopCSI      = "m"
	EscapeStopOSI      = "\\"
	EscapeStopRune     = EscapeStopRuneCSI
	EscapeStopRuneCSI  = 'm'
	EscapeStopRuneOSI  = '\\'
)

// Deprecated Constants
const (
	CSIStartRune = EscapeStartRuneCSI
	CSIStopRune  = EscapeStopRuneCSI
	OSIStartRune = EscapeStartRuneOSI
	OSIStopRune  = EscapeStopRuneOSI
)

type escSeqKind int

const (
	escSeqKindUnknown escSeqKind = iota
	escSeqKindCSI
	escSeqKindOSI
	escSeqKindString
)

// private constants
const (
	escCodeResetAll        = 0
	escCodeResetIntensity  = 22
	escCodeResetItalic     = 23
	escCodeResetUnderline  = 24
	escCodeResetBlink      = 25
	escCodeResetReverse    = 27
	escCodeResetCrossedOut = 29
	escCodeBold            = 1
	escCodeDim             = 2
	escCodeItalic          = 3
	escCodeUnderline       = 4
	escCodeBlinkSlow       = 5
	escCodeBlinkRapid      = 6
	escCodeReverse         = 7
	escCodeConceal         = 8
	escCodeCrossedOut      = 9

	// OSC 8 hyperlink sequences
	escapeStartConcealOSI = "\x1b]8;"
	escapeStopConcealOSI  = "\x1b\\"

	// the runes that follow ESC to start a sequence that carries a payload
	// until ST: Device Control String, Start Of String, Privacy Message and
	// Application Program Command
	escapeStartRuneDCS = 'P'
	escapeStartRuneSOS = 'X'
	escapeStartRunePM  = '^'
	escapeStartRuneAPC = '_'

	// escSeqMaxLength bounds the accumulated escape sequence; real terminals
	// cap sequences in the low KBs, and anything longer here is most likely
	// malformed input that would otherwise grow the buffer without limit
	escSeqMaxLength = 2048
)

// []byte forms of the OSC 8 hyperlink markers for prefix/suffix checks on the
// escape sequence buffer without allocating per call
var (
	escapeStartConcealOSIBytes = []byte(escapeStartConcealOSI)
	escapeStopConcealOSIBytes  = []byte(escapeStopConcealOSI)
)

// 256-color and 24-bit ("true") color codes: "38;5;n" and "38;2;r;g;b" set
// the foreground, and "48;5;n" and "48;2;r;g;b" set the background
const (
	escCode256FgStart = 38
	escCode256BgStart = 48
	escCode256Color   = 5
	escCodeRGBColor   = 2
	escCodeResetFg    = 39
	escCodeResetBg    = 49
	escCode256Max     = 255
)

// Internal encoding for 256-color codes uses fg256Start and bg256Start from color.go
// Private constants initialized from private constants to avoid repeated casting in hot paths
// Foreground 256-color: fg256Start + colorIndex (1000-1255)
// Background 256-color: bg256Start + colorIndex (2000-2255)
const (
	escCode256FgBase = int(fg256Start) // 1000
	escCode256BgBase = int(bg256Start) // 2000
)

// Standard color code ranges
const (
	// Standard foreground colors (30-37)
	escCodeFgStdStart = 30
	escCodeFgStdEnd   = 37
	// Bright foreground colors (90-97)
	escCodeFgBrightStart = 90
	escCodeFgBrightEnd   = 97
	// Standard background colors (40-47)
	escCodeBgStdStart = 40
	escCodeBgStdEnd   = 47
	// Bright background colors (100-107)
	escCodeBgBrightStart = 100
	escCodeBgBrightEnd   = 107
)

// Special characters
const (
	escRuneBEL = '\a' // BEL character (ASCII 7)
)

// EscSeqParser parses ANSI escape sequences from text and tracks active formatting codes.
// It supports both CSI (Control Sequence Introducer) and OSI (Operating System Command)
// escape sequence formats.
type EscSeqParser struct {
	// codes tracks active escape sequence codes (e.g., 1 for bold, 3 for italic).
	codes map[int]bool
	// fgRGB and bgRGB track the active 24-bit foreground and background
	// colors; these have no code of their own, and are not part of Codes():
	// use FgRGB() and BgRGB() for these.
	fgRGB, bgRGB escRGBColor

	// inEscSeq indicates whether the parser is currently inside an escape sequence.
	inEscSeq bool
	// escSeqKind identifies the type of escape sequence being parsed (CSI or OSI).
	escSeqKind escSeqKind
	// escapeSeq accumulates the current escape sequence being parsed; it is
	// bounded by escSeqMaxLength and reused across sequences.
	escapeSeq []byte

	// skipParse makes Consume track only where escape sequences begin and
	// end, without parsing the codes in them; meant for the callers that
	// measure or cut text and never ask for Codes, IsOpen or Sequence.
	skipParse bool
}

// escRGBColor is a 24-bit color set by a "38;2;r;g;b" or "48;2;r;g;b" sequence.
type escRGBColor struct {
	set     bool
	r, g, b int
}

// BgRGB returns the active 24-bit background color, as set by a sequence like
// "\x1b[48;2;r;g;b m"; ok is false if no such color is active.
func (s *EscSeqParser) BgRGB() (r, g, b int, ok bool) {
	return s.bgRGB.r, s.bgRGB.g, s.bgRGB.b, s.bgRGB.set
}

// Codes returns the active codes in ascending order. 256-colors are returned
// in the encoding that Fg256Color and Bg256Color use. 24-bit colors have no
// Color equivalent and are left out: FgRGB() and BgRGB() return those, and
// Sequence() includes them.
func (s *EscSeqParser) Codes() []int {
	codes := make([]int, 0)
	for code, val := range s.codes {
		if val {
			codes = append(codes, code)
		}
	}
	sort.Ints(codes)
	return codes
}

// escSeqKindFor returns the kind of sequence introduced by char after ESC.
func escSeqKindFor(char rune) escSeqKind {
	switch char {
	case EscapeStartRuneCSI:
		return escSeqKindCSI
	case EscapeStartRuneOSI:
		return escSeqKindOSI
	case escapeStartRuneDCS, escapeStartRuneSOS, escapeStartRunePM, escapeStartRuneAPC:
		// DCS, SOS, PM and APC carry a payload; the sequence runs until
		// ST ("ESC \") instead of ending at a final byte (ECMA-48)
		return escSeqKindString
	}
	return escSeqKindUnknown
}

func (s *EscSeqParser) Consume(char rune) {
	if !s.inEscSeq && char == EscapeStartRune {
		s.inEscSeq = true
		s.escSeqKind = escSeqKindUnknown
		s.escapeSeq = s.escapeSeq[:0]
	} else if s.inEscSeq && s.escSeqKind == escSeqKindUnknown {
		s.escSeqKind = escSeqKindFor(char)
	}

	if s.inEscSeq {
		s.escapeSeq = utf8.AppendRune(s.escapeSeq, char)

		// abort on absurdly long (most likely malformed/unterminated)
		// sequences instead of accumulating them without limit
		if len(s.escapeSeq) > escSeqMaxLength {
			s.Reset()
			return
		}

		// --- FIX for OSC 8 hyperlinks (e.g. \x1b]8;;url\x07label\x1b]8;;\x07)
		if s.escSeqKind == escSeqKindOSI &&
			bytes.HasPrefix(s.escapeSeq, escapeStartConcealOSIBytes) &&
			char == escRuneBEL { // BEL

			s.parseEscapeSeq()
			s.Reset()
			return
		}

		if s.isEscapeStopRune(char) {
			s.parseEscapeSeq()
			s.Reset()
		}
	}
}

// FgRGB returns the active 24-bit foreground color, as set by a sequence like
// "\x1b[38;2;r;g;b m"; ok is false if no such color is active.
func (s *EscSeqParser) FgRGB() (r, g, b int, ok bool) {
	return s.fgRGB.r, s.fgRGB.g, s.fgRGB.b, s.fgRGB.set
}

func (s *EscSeqParser) InSequence() bool {
	return s.inEscSeq
}

func (s *EscSeqParser) IsOpen() bool {
	return len(s.codes) > 0 || s.fgRGB.set || s.bgRGB.set
}

func (s *EscSeqParser) ParseSeq(seq string, seqKind escSeqKind) {
	// only CSI and OSI sequences carry character attributes. A bare escape
	// ("ESC 7"), a DCS/SOS/PM/APC string or any other kind must not be parsed
	// as SGR: their parameter bytes are payload, not attribute codes.
	if seqKind != escSeqKindCSI && seqKind != escSeqKindOSI {
		return
	}
	if s.codes == nil {
		s.codes = make(map[int]bool)
	}

	// only SGR sequences ("CSI ... m") carry character attributes; the other
	// CSI final bytes (cursor movement, line erasing, ...) carry none, and
	// their parameter bytes are not attribute codes
	if seqKind == escSeqKindCSI && !strings.HasSuffix(seq, EscapeStopCSI) {
		return
	}
	// an OSC 8 hyperlink ("OSC 8 ; params ; URI ST") carries a link and no
	// attributes: its 8 is not the code to conceal text, and the numbers in
	// its params and URI are not codes either
	if seqKind == escSeqKindOSI && strings.HasPrefix(seq, escapeStartConcealOSI) {
		return
	}

	seq = s.stripEscapeSequence(seq, seqKind)
	codes := s.splitAndTrimCodes(seq)
	if seqKind == escSeqKindCSI {
		// a missing SGR parameter stands for 0: "CSI m" is a reset
		for i, code := range codes {
			if code == "" {
				codes[i] = "0"
			}
		}
	}
	s.processCodes(codes)
}

func (s *EscSeqParser) ParseString(str string) string {
	s.escapeSeq, s.inEscSeq, s.escSeqKind = s.escapeSeq[:0], false, escSeqKindUnknown
	for _, char := range str {
		s.Consume(char)
	}
	return s.Sequence()
}

func (s *EscSeqParser) Reset() {
	s.inEscSeq = false
	s.escSeqKind = escSeqKindUnknown
	s.escapeSeq = s.escapeSeq[:0]
}

func (s *EscSeqParser) Sequence() string {
	if !s.IsOpen() {
		return ""
	}

	parts := make([]string, 0, len(s.codes)+2)
	for _, code := range s.Codes() {
		if code >= escCode256FgBase && code <= escCode256FgBase+escCode256Max {
			// 256-color foreground code (1000-1255)
			parts = append(parts, fmt.Sprintf("%d;%d;%d", escCode256FgStart, escCode256Color, code-escCode256FgBase))
		} else if code >= escCode256BgBase && code <= escCode256BgBase+escCode256Max {
			// 256-color background code (2000-2255)
			parts = append(parts, fmt.Sprintf("%d;%d;%d", escCode256BgStart, escCode256Color, code-escCode256BgBase))
		} else {
			// Regular code
			parts = append(parts, strconv.Itoa(code))
		}
	}
	if c := s.fgRGB; c.set {
		parts = append(parts, fmt.Sprintf("%d;%d;%d;%d;%d", escCode256FgStart, escCodeRGBColor, c.r, c.g, c.b))
	}
	if c := s.bgRGB; c.set {
		parts = append(parts, fmt.Sprintf("%d;%d;%d;%d;%d", escCode256BgStart, escCodeRGBColor, c.r, c.g, c.b))
	}
	return EscapeStart + strings.Join(parts, ";") + EscapeStop
}

// clearAllBackgroundColors clears the active background color, whichever
// form (standard, bright, 256-color or 24-bit) it was set in.
func (s *EscSeqParser) clearAllBackgroundColors() {
	for code := range s.codes {
		if isBackgroundColorCode(code) {
			delete(s.codes, code)
		}
	}
	s.bgRGB = escRGBColor{}
}

// clearAllForegroundColors clears the active foreground color, whichever
// form (standard, bright, 256-color or 24-bit) it was set in.
func (s *EscSeqParser) clearAllForegroundColors() {
	for code := range s.codes {
		if isForegroundColorCode(code) {
			delete(s.codes, code)
		}
	}
	s.fgRGB = escRGBColor{}
}

// isBackgroundColorCode returns true if the code is a background color, in
// the encoding used by the codes map.
func isBackgroundColorCode(code int) bool {
	return (code >= escCodeBgStdStart && code <= escCodeBgStdEnd) ||
		(code >= escCodeBgBrightStart && code <= escCodeBgBrightEnd) ||
		(code >= escCode256BgBase && code <= escCode256BgBase+escCode256Max)
}

func (s *EscSeqParser) isEscapeStopRune(char rune) bool {
	if bytes.HasPrefix(s.escapeSeq, escapeStartConcealOSIBytes) {
		if bytes.HasSuffix(s.escapeSeq, escapeStopConcealOSIBytes) {
			return true
		}
	} else if (s.escSeqKind == escSeqKindCSI && len(s.escapeSeq) > len(EscapeStartCSI) && isCSIFinalRune(char)) ||
		(s.escSeqKind == escSeqKindOSI && char == EscapeStopRuneOSI) {
		return true
	} else if s.escSeqKind == escSeqKindString && char == EscapeStopRuneOSI {
		// ST is "ESC \\", but this stops at the backslash itself, just like
		// the OSI branch above; a backslash inside a DCS/SOS/PM/APC payload
		// therefore also ends the sequence early. Tightening both is a
		// separate change.
		return true
	} else if s.escSeqKind == escSeqKindUnknown && isEscapeFinalRune(char) {
		// an ESC that starts neither a CSI, an OSI, nor a string sequence
		// forms an escape sequence of its own: "ESC 7" (save cursor), "ESC c"
		// (reset terminal), or "ESC ( B" (select charset); such a sequence
		// ends at its final byte (ECMA-48), and everything after it is
		// visible text again
		return true
	}
	return false
}

// isCSIFinalRune returns true if the given rune ends a CSI sequence. A CSI
// sequence is "ESC [", followed by parameter bytes (0x30-0x3F) and
// intermediate bytes (0x20-0x2F), and is terminated by the first byte in the
// range 0x40-0x7E (ECMA-48, section 5.4). "m" (SGR) is only one of those final
// bytes; the Cursor type of this package emits others like "A" and "K".
func isCSIFinalRune(char rune) bool {
	return char >= 0x40 && char <= 0x7E
}

// isForegroundColorCode returns true if the code is a foreground color, in
// the encoding used by the codes map.
func isForegroundColorCode(code int) bool {
	return (code >= escCodeFgStdStart && code <= escCodeFgStdEnd) ||
		(code >= escCodeFgBrightStart && code <= escCodeFgBrightEnd) ||
		(code >= escCode256FgBase && code <= escCode256FgBase+escCode256Max)
}

// isEscapeFinalRune returns true if the given rune can end an escape sequence
// that is not a CSI or an OSI sequence: "ESC", followed by intermediate bytes
// (0x20-0x2F) and a final byte in the range 0x30-0x7E (ECMA-48, section 8.3.143).
func isEscapeFinalRune(char rune) bool {
	return char >= 0x30 && char <= 0x7E
}

// isRegularCode checks if a code is a regular code (not a 256-color encoded value).
func (s *EscSeqParser) isRegularCode(codeNum int) bool {
	return codeNum < escCode256FgBase || codeNum > escCode256BgBase+escCode256Max
}

// parseColorParams parses the first n of the given parameters as the numbers
// of an extended color: the index of a 256-color, or the R, G and B of a 24-bit
// color. It fails if any of them is missing or is not in the range 0-255.
func parseColorParams(params []string, n int) ([3]int, bool) {
	var values [3]int
	if len(params) < n {
		return values, false
	}
	for i := 0; i < n; i++ {
		value, err := strconv.Atoi(params[i])
		if err != nil || value < 0 || value > escCode256Max {
			return values, false
		}
		values[i] = value
	}
	return values, true
}

// parseEscapeSeq parses the escape sequence that was just consumed, unless the
// parser has been asked to skip that.
func (s *EscSeqParser) parseEscapeSeq() {
	if !s.skipParse {
		s.ParseSeq(string(s.escapeSeq), s.escSeqKind)
	}
}

// processCode handles a single escape code.
func (s *EscSeqParser) processCode(codeNum int) {
	switch codeNum {
	case escCodeResetAll:
		s.codes = make(map[int]bool)
		s.fgRGB, s.bgRGB = escRGBColor{}, escRGBColor{}
	case escCodeResetIntensity:
		delete(s.codes, escCodeBold)
		delete(s.codes, escCodeDim)
	case escCodeResetItalic:
		delete(s.codes, escCodeItalic)
	case escCodeResetUnderline:
		delete(s.codes, escCodeUnderline)
	case escCodeResetBlink:
		delete(s.codes, escCodeBlinkSlow)
		delete(s.codes, escCodeBlinkRapid)
	case escCodeResetReverse:
		delete(s.codes, escCodeReverse)
	case escCodeResetCrossedOut:
		delete(s.codes, escCodeCrossedOut)
	case escCodeResetFg:
		s.clearAllForegroundColors()
	case escCodeResetBg:
		s.clearAllBackgroundColors()
	default:
		if s.isRegularCode(codeNum) {
			s.setCode(codeNum)
		}
	}
}

// processCodes applies the codes of a sequence in the order they appear in,
// like a terminal does: the color in "0;31" survives the reset before it.
func (s *EscSeqParser) processCodes(codes []string) {
	for i := 0; i < len(codes); i++ {
		codeNum, err := strconv.Atoi(codes[i])
		if err != nil {
			continue
		}

		if codeNum == escCode256FgStart || codeNum == escCode256BgStart {
			if consumed := s.processExtendedColor(codeNum == escCode256FgStart, codes[i+1:]); consumed > 0 {
				i += consumed
				continue
			}
		}
		s.processCode(codeNum)
	}
}

// processExtendedColor handles the parameters that follow a 38 (foreground)
// or a 48 (background): "5;n" for a 256-color, or "2;r;g;b" for a 24-bit
// color. It returns the number of parameters that belong to the color, which
// the caller should skip; 0 means that the parameters do not describe a color.
//
// The parameters of a color are never read as codes, even if the color is
// malformed and gets ignored: the zeroes in "38;2;255;0;0" and "38;2;256;0;0"
// do not reset anything, and the 5 in "38;5;999" does not start blinking.
func (s *EscSeqParser) processExtendedColor(isForeground bool, params []string) int {
	if len(params) == 0 {
		return 0
	}

	numValues := 0
	switch colorType, _ := strconv.Atoi(params[0]); colorType {
	case escCode256Color:
		numValues = 1
	case escCodeRGBColor:
		numValues = 3
	default:
		return 0
	}

	values, ok := parseColorParams(params[1:], numValues)
	if !ok {
		if len(params) < 1+numValues { // cut short; nothing else follows
			return len(params)
		}
		return 1 + numValues
	}

	if numValues == 1 {
		if isForeground {
			s.setCode(escCode256FgBase + values[0])
		} else {
			s.setCode(escCode256BgBase + values[0])
		}
	} else {
		color := escRGBColor{set: true, r: values[0], g: values[1], b: values[2]}
		if isForeground {
			s.clearAllForegroundColors()
			s.fgRGB = color
		} else {
			s.clearAllBackgroundColors()
			s.bgRGB = color
		}
	}
	return 1 + numValues
}

// setCode turns on the given code. A terminal has just one foreground and one
// background color at a time, so a color replaces the one that was active.
func (s *EscSeqParser) setCode(code int) {
	if isForegroundColorCode(code) {
		s.clearAllForegroundColors()
	} else if isBackgroundColorCode(code) {
		s.clearAllBackgroundColors()
	}
	s.codes[code] = true
}

// splitAndTrimCodes splits the sequence by semicolons and trims whitespace.
func (s *EscSeqParser) splitAndTrimCodes(seq string) []string {
	codes := strings.Split(seq, ";")
	for i := range codes {
		codes[i] = strings.TrimSpace(codes[i])
	}
	return codes
}

// stripEscapeSequence removes escape sequence markers from the input string.
func (s *EscSeqParser) stripEscapeSequence(seq string, seqKind escSeqKind) string {
	if seqKind == escSeqKindOSI {
		seq = strings.Replace(seq, EscapeStartOSI, "", 1)
		seq = strings.Replace(seq, EscapeStopOSI, "", 1)
	} else {
		seq = strings.Replace(seq, EscapeStartCSI, "", 1)
		seq = strings.Replace(seq, EscapeStopCSI, "", 1)
	}
	return seq
}
