package text

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type escSeqParser struct {
	openSeq map[int]bool
}

func (s *escSeqParser) Codes() []int {
	codes := make([]int, 0)
	for code, val := range s.openSeq {
		if val {
			codes = append(codes, code)
		}
	}
	sort.Ints(codes)
	return codes
}

func (s *escSeqParser) Extract(str string) string {
	escapeSeq, inEscSeq := "", false
	for _, char := range str {
		if char == EscapeStartRune {
			inEscSeq = true
			escapeSeq = ""
		}
		if inEscSeq {
			escapeSeq += string(char)
		}
		if char == EscapeStopRune {
			inEscSeq = false
			s.Parse(escapeSeq)
		}
	}
	return s.Sequence()
}

func (s *escSeqParser) IsOpen() bool {
	return len(s.openSeq) > 0
}

func (s *escSeqParser) Sequence() string {
	out := strings.Builder{}
	if s.IsOpen() {
		out.WriteString(EscapeStart)
		for idx, code := range s.Codes() {
			if idx > 0 {
				out.WriteRune(';')
			}
			out.WriteString(fmt.Sprint(code))
		}
		out.WriteString(EscapeStop)
	}

	return out.String()
}

func (s *escSeqParser) Parse(seq string) {
	if s.openSeq == nil {
		s.openSeq = make(map[int]bool)
	}

	seq = strings.Replace(seq, EscapeStart, "", 1)
	seq = strings.Replace(seq, EscapeStop, "", 1)
	codes := strings.Split(seq, ";")
	for _, code := range codes {
		code = strings.TrimSpace(code)
		if codeNum, err := strconv.Atoi(code); err == nil {
			switch codeNum {
			case 0: // reset
				s.openSeq = make(map[int]bool) // clear everything
			case 22: // reset intensity
				delete(s.openSeq, 1) // remove bold
				delete(s.openSeq, 2) // remove faint
			case 23: // not italic
				delete(s.openSeq, 3) // remove italic
			case 24: // not underlined
				delete(s.openSeq, 4) // remove underline
			case 25: // not blinking
				delete(s.openSeq, 5) // remove slow blink
				delete(s.openSeq, 6) // remove rapid blink
			case 27: // not reversed
				delete(s.openSeq, 7) // remove reverse
			case 29: // not crossed-out
				delete(s.openSeq, 9) // remove crossed-out
			default:
				s.openSeq[codeNum] = true
			}
		}
	}
}
