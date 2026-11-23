package text

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEscSeqParser(t *testing.T) {
	t.Run("extract csi", func(t *testing.T) {
		es := EscSeqParser{}

		assert.Equal(t, "\x1b[1;3;4;5;7;9;91m", es.ParseString("\x1b[91m\x1b[1m\x1b[3m\x1b[4m\x1b[5m\x1b[7m\x1b[9m Spicy"))
		assert.Equal(t, "\x1b[3;4;5;7;9;91m", es.ParseString("\x1b[22m No Bold"))
		assert.Equal(t, "\x1b[4;5;7;9;91m", es.ParseString("\x1b[23m No Italic"))
		assert.Equal(t, "\x1b[5;7;9;91m", es.ParseString("\x1b[24m No Underline"))
		assert.Equal(t, "\x1b[7;9;91m", es.ParseString("\x1b[25m No Blink"))
		assert.Equal(t, "\x1b[9;91m", es.ParseString("\x1b[27m No Reverse"))
		assert.Equal(t, "\x1b[91m", es.ParseString("\x1b[29m No Crossed-Out"))
		assert.Equal(t, "", es.ParseString("\x1b[0m Resetted"))
	})

	t.Run("extract osi", func(t *testing.T) {
		es := EscSeqParser{}

		assert.Equal(t, "\x1b[1;3;4;5;7;9;91m", es.ParseString("\x1b]91\\\x1b]1\\\x1b]3\\\x1b]4\\\x1b]5\\\x1b]7\\\x1b]9\\ Spicy"))
		assert.Equal(t, "\x1b[3;4;5;7;9;91m", es.ParseString("\x1b]22\\ No Bold"))
		assert.Equal(t, "\x1b[4;5;7;9;91m", es.ParseString("\x1b]23\\ No Italic"))
		assert.Equal(t, "\x1b[5;7;9;91m", es.ParseString("\x1b]24\\ No Underline"))
		assert.Equal(t, "\x1b[7;9;91m", es.ParseString("\x1b]25\\ No Blink"))
		assert.Equal(t, "\x1b[9;91m", es.ParseString("\x1b]27\\ No Reverse"))
		assert.Equal(t, "\x1b[91m", es.ParseString("\x1b]29\\ No Crossed-Out"))
		assert.Equal(t, "", es.ParseString("\x1b[0m Resetted"))
	})

	t.Run("parse csi", func(t *testing.T) {
		es := EscSeqParser{}

		es.ParseSeq("\x1b[91m", escSeqKindCSI) // color
		es.ParseSeq("\x1b[1m", escSeqKindCSI)  // bold
		assert.Len(t, es.Codes(), 2)
		assert.True(t, es.IsOpen())
		assert.Equal(t, "\x1b[1;91m", es.Sequence())

		es.ParseSeq("\x1b[22m", escSeqKindCSI) // un-bold
		assert.Len(t, es.Codes(), 1)
		assert.True(t, es.IsOpen())
		assert.Equal(t, "\x1b[91m", es.Sequence())

		es.ParseSeq("\x1b[0m", escSeqKindCSI) // reset
		assert.Empty(t, es.Codes())
		assert.False(t, es.IsOpen())
		assert.Empty(t, es.Sequence())
	})

	t.Run("parse osi", func(t *testing.T) {
		es := EscSeqParser{}

		es.ParseSeq("\x1b]91\\", escSeqKindOSI) // color
		es.ParseSeq("\x1b]1\\", escSeqKindOSI)  // bold
		assert.Len(t, es.Codes(), 2)
		assert.True(t, es.IsOpen())
		assert.Equal(t, "\x1b[1;91m", es.Sequence())

		es.ParseSeq("\x1b]22\\", escSeqKindOSI) // un-bold
		assert.Len(t, es.Codes(), 1)
		assert.True(t, es.IsOpen())
		assert.Equal(t, "\x1b[91m", es.Sequence())

		es.ParseSeq("\x1b]0\\", escSeqKindOSI) // reset
		assert.Empty(t, es.Codes())
		assert.False(t, es.IsOpen())
		assert.Empty(t, es.Sequence())
	})

	t.Run("osi hyperlink with bel", func(t *testing.T) {
		// Test OSC 8 hyperlink termination with BEL character (\a)
		// Format: \x1b]8;;url\x07label\x1b]8;;\x07
		es := EscSeqParser{}

		// Simulate parsing an OSC 8 hyperlink that ends with BEL
		for _, char := range "\x1b]8;;url\x07" {
			es.Consume(char)
		}
		// After consuming BEL, the parser should have reset; the hyperlink
		// carries no attributes: its 8 is not the code to conceal text
		assert.False(t, es.InSequence())
		assert.False(t, es.IsOpen())
		assert.Empty(t, es.Codes())

		// Test with a full hyperlink sequence
		es = EscSeqParser{} // fresh parser
		result := es.ParseString("\x1b]8;;https://example.com\x07Click here\x1b]8;;\x07")
		assert.Equal(t, "", result)
	})

	t.Run("osi hyperlink has no codes", func(t *testing.T) {
		for _, terminator := range []string{"\x07", "\x1b\\"} {
			// neither the 8, nor the numbers in the params and in the URI
			// are codes; and the 0 in the URI is not a reset
			link := "\x1b]8;id=1;https://example.com/0;31;4" + terminator + "label\x1b]8;;" + terminator

			es := EscSeqParser{}
			assert.Equal(t, "", es.ParseString(link))
			assert.False(t, es.InSequence())

			// the colors around the hyperlink are left alone
			es = EscSeqParser{}
			assert.Equal(t, "\x1b[1;32m", es.ParseString("\x1b[1;32m"+link))
			assert.Equal(t, "\x1b[1;32;44m", es.ParseString(link+"\x1b[44m"))
		}

		// "OSC 8" without the rest of a hyperlink is still the conceal code
		es := EscSeqParser{}
		es.ParseSeq("\x1b]8\\", escSeqKindOSI)
		assert.Equal(t, []int{escCodeConceal}, es.Codes())
	})

	t.Run("consume directly", func(t *testing.T) {
		es := EscSeqParser{}

		// Test entering escape sequence
		es.Consume('\x1b')
		assert.True(t, es.InSequence())

		// Test identifying CSI sequence
		es.Consume('[')
		assert.True(t, es.InSequence())

		// Test consuming code
		es.Consume('1')
		assert.True(t, es.InSequence())

		// Test completing sequence
		es.Consume('m')
		assert.False(t, es.InSequence())
		assert.True(t, es.IsOpen())
		assert.Equal(t, "\x1b[1m", es.Sequence())

		// Test entering OSI sequence
		es.Reset()
		es.Consume('\x1b')
		es.Consume(']')
		assert.True(t, es.InSequence())

		// Test consuming code and completing
		es.Consume('9')
		es.Consume('1')
		es.Consume('\\')
		assert.False(t, es.InSequence())
		assert.True(t, es.IsOpen())
	})

	t.Run("consume invalid escape sequence", func(t *testing.T) {
		es := EscSeqParser{}

		// Test escape sequence that doesn't start with CSI or OSI
		es.Consume('\x1b')
		assert.True(t, es.InSequence())

		// Consume something that's neither '[' nor ']'
		es.Consume('X')
		assert.True(t, es.InSequence())

		// Reset manually
		es.Reset()
		assert.False(t, es.InSequence())
	})

	t.Run("parse seq with spaces and multiple codes", func(t *testing.T) {
		es := EscSeqParser{}

		// Test parsing sequence with spaces
		es.ParseSeq("\x1b[ 1 ; 3 ; 4 m", escSeqKindCSI)
		assert.Len(t, es.Codes(), 3)
		assert.True(t, es.IsOpen())

		// Test parsing sequence with invalid code (non-numeric)
		// Reset doesn't clear codes, so parse reset first
		es.ParseSeq("\x1b[0m", escSeqKindCSI)
		assert.Empty(t, es.Codes())
		es.ParseSeq("\x1b[1;abc;3m", escSeqKindCSI)
		// Should only have valid numeric codes (1 and 3, abc is skipped)
		assert.Len(t, es.Codes(), 2)
		assert.Contains(t, es.Codes(), 1)
		assert.Contains(t, es.Codes(), 3)

		// Test parsing sequence with invalid code (non-numeric)
		// Reset doesn't clear codes, so parse reset first
		es.ParseSeq("\x1b[0m", escSeqKindCSI)
		assert.Empty(t, es.Codes())
		es.ParseSeq("\x1b[abc;1;3m", escSeqKindCSI)
		// Should only have valid numeric codes (1 and 3, abc is skipped)
		assert.Len(t, es.Codes(), 2)
		assert.Contains(t, es.Codes(), 1)
		assert.Contains(t, es.Codes(), 3)
	})

	t.Run("parse seq edge cases", func(t *testing.T) {
		es := EscSeqParser{}

		// Test empty sequence after stripping
		es.ParseSeq("\x1b[m", escSeqKindCSI)
		assert.Empty(t, es.Codes())

		// Test all reset codes
		es.ParseSeq("\x1b[1m", escSeqKindCSI)
		es.ParseSeq("\x1b[22m", escSeqKindCSI) // reset intensity
		assert.Empty(t, es.Codes())

		es.ParseSeq("\x1b[3m", escSeqKindCSI)
		es.ParseSeq("\x1b[23m", escSeqKindCSI) // reset italic
		assert.Empty(t, es.Codes())

		es.ParseSeq("\x1b[4m", escSeqKindCSI)
		es.ParseSeq("\x1b[24m", escSeqKindCSI) // reset underline
		assert.Empty(t, es.Codes())

		es.ParseSeq("\x1b[5m", escSeqKindCSI)
		es.ParseSeq("\x1b[25m", escSeqKindCSI) // reset blink
		assert.Empty(t, es.Codes())

		es.ParseSeq("\x1b[6m", escSeqKindCSI)
		es.ParseSeq("\x1b[25m", escSeqKindCSI) // reset blink (also clears 6)
		assert.Empty(t, es.Codes())

		es.ParseSeq("\x1b[7m", escSeqKindCSI)
		es.ParseSeq("\x1b[27m", escSeqKindCSI) // reset reverse
		assert.Empty(t, es.Codes())

		es.ParseSeq("\x1b[9m", escSeqKindCSI)
		es.ParseSeq("\x1b[29m", escSeqKindCSI) // reset crossed out
		assert.Empty(t, es.Codes())
	})

	t.Run("256-color foreground", func(t *testing.T) {
		es := EscSeqParser{}

		// Test parsing 256-color foreground sequence
		es.ParseSeq("\x1b[38;5;123m", escSeqKindCSI)
		assert.Len(t, es.Codes(), 1)
		assert.True(t, es.IsOpen())
		// Should have encoded value 1000 + 123 = 1123
		assert.Contains(t, es.Codes(), escCode256FgBase+123)
		assert.Equal(t, "\x1b[38;5;123m", es.Sequence())

		// Test parsing from string
		es = EscSeqParser{}
		result := es.ParseString("\x1b[38;5;123m")
		assert.Equal(t, "\x1b[38;5;123m", result)

		// Test boundary values
		es = EscSeqParser{}
		es.ParseSeq("\x1b[38;5;0m", escSeqKindCSI)
		assert.Contains(t, es.Codes(), escCode256FgBase+0)
		assert.Equal(t, "\x1b[38;5;0m", es.Sequence())

		es = EscSeqParser{}
		es.ParseSeq("\x1b[38;5;255m", escSeqKindCSI)
		assert.Contains(t, es.Codes(), escCode256FgBase+255)
		assert.Equal(t, "\x1b[38;5;255m", es.Sequence())
	})

	t.Run("256-color background", func(t *testing.T) {
		es := EscSeqParser{}

		// Test parsing 256-color background sequence
		es.ParseSeq("\x1b[48;5;200m", escSeqKindCSI)
		assert.Len(t, es.Codes(), 1)
		assert.True(t, es.IsOpen())
		// Should have encoded value 2000 + 200 = 2200
		assert.Contains(t, es.Codes(), escCode256BgBase+200)
		assert.Equal(t, "\x1b[48;5;200m", es.Sequence())

		// Test parsing from string
		es = EscSeqParser{}
		result := es.ParseString("\x1b[48;5;200m")
		assert.Equal(t, "\x1b[48;5;200m", result)

		// Test boundary values
		es = EscSeqParser{}
		es.ParseSeq("\x1b[48;5;0m", escSeqKindCSI)
		assert.Contains(t, es.Codes(), escCode256BgBase+0)
		assert.Equal(t, "\x1b[48;5;0m", es.Sequence())

		es = EscSeqParser{}
		es.ParseSeq("\x1b[48;5;255m", escSeqKindCSI)
		assert.Contains(t, es.Codes(), escCode256BgBase+255)
		assert.Equal(t, "\x1b[48;5;255m", es.Sequence())
	})

	t.Run("256-color with regular codes", func(t *testing.T) {
		es := EscSeqParser{}

		// Test mixing 256-color with regular formatting codes
		es.ParseSeq("\x1b[38;5;100m", escSeqKindCSI)
		es.ParseSeq("\x1b[1m", escSeqKindCSI) // bold
		es.ParseSeq("\x1b[3m", escSeqKindCSI) // italic
		assert.Len(t, es.Codes(), 3)
		assert.Contains(t, es.Codes(), escCode256FgBase+100)
		assert.Contains(t, es.Codes(), escCodeBold)
		assert.Contains(t, es.Codes(), escCodeItalic)
		// Sequence should include all codes
		seq := es.Sequence()
		assert.Equal(t, "\x1b[1;3;38;5;100m", seq)

		// Test 256-color foreground and background together
		es = EscSeqParser{}
		es.ParseSeq("\x1b[38;5;50m", escSeqKindCSI)
		es.ParseSeq("\x1b[48;5;150m", escSeqKindCSI)
		assert.Len(t, es.Codes(), 2)
		assert.Contains(t, es.Codes(), escCode256FgBase+50)
		assert.Contains(t, es.Codes(), escCode256BgBase+150)
		seq = es.Sequence()
		assert.Contains(t, seq, "38;5;50")
		assert.Contains(t, seq, "48;5;150")
	})

	t.Run("256-color replaces standard colors", func(t *testing.T) {
		es := EscSeqParser{}

		// Set standard foreground color
		es.ParseSeq("\x1b[31m", escSeqKindCSI) // red
		assert.Contains(t, es.Codes(), 31)
		// Set 256-color foreground - should replace standard color
		es.ParseSeq("\x1b[38;5;123m", escSeqKindCSI)
		assert.NotContains(t, es.Codes(), 31)
		assert.Contains(t, es.Codes(), escCode256FgBase+123)

		// Set standard background color
		es = EscSeqParser{}
		es.ParseSeq("\x1b[41m", escSeqKindCSI) // red background
		assert.Contains(t, es.Codes(), 41)
		// Set 256-color background - should replace standard color
		es.ParseSeq("\x1b[48;5;200m", escSeqKindCSI)
		assert.NotContains(t, es.Codes(), 41)
		assert.Contains(t, es.Codes(), escCode256BgBase+200)
	})

	t.Run("256-color reset codes", func(t *testing.T) {
		es := EscSeqParser{}

		// Set 256-color foreground
		es.ParseSeq("\x1b[38;5;123m", escSeqKindCSI)
		assert.Contains(t, es.Codes(), escCode256FgBase+123)
		// Reset foreground (code 39)
		es.ParseSeq("\x1b[39m", escSeqKindCSI)
		assert.NotContains(t, es.Codes(), escCode256FgBase+123)
		assert.Empty(t, es.Codes())

		// Set 256-color background
		es = EscSeqParser{}
		es.ParseSeq("\x1b[48;5;200m", escSeqKindCSI)
		assert.Contains(t, es.Codes(), escCode256BgBase+200)
		// Reset background (code 49)
		es.ParseSeq("\x1b[49m", escSeqKindCSI)
		assert.NotContains(t, es.Codes(), escCode256BgBase+200)
		assert.Empty(t, es.Codes())

		// Test reset all (code 0) clears 256-color codes
		es = EscSeqParser{}
		es.ParseSeq("\x1b[38;5;123m", escSeqKindCSI)
		es.ParseSeq("\x1b[48;5;200m", escSeqKindCSI)
		assert.Len(t, es.Codes(), 2)
		es.ParseSeq("\x1b[0m", escSeqKindCSI)
		assert.Empty(t, es.Codes())
	})

	t.Run("256-color OSI format", func(t *testing.T) {
		es := EscSeqParser{}

		// Test parsing 256-color foreground in OSI format
		es.ParseSeq("\x1b]38;5;123\\", escSeqKindOSI)
		assert.Contains(t, es.Codes(), escCode256FgBase+123)
		assert.Equal(t, "\x1b[38;5;123m", es.Sequence())

		// Test parsing 256-color background in OSI format
		es = EscSeqParser{}
		es.ParseSeq("\x1b]48;5;200\\", escSeqKindOSI)
		assert.Contains(t, es.Codes(), escCode256BgBase+200)
		assert.Equal(t, "\x1b[48;5;200m", es.Sequence())
	})

	t.Run("256-color invalid sequences", func(t *testing.T) {
		es := EscSeqParser{}

		// Test incomplete 256-color sequence (missing color index)
		es.ParseSeq("\x1b[38;5m", escSeqKindCSI)
		// Should not have any 256-color codes
		for code := range es.Codes() {
			assert.False(t, code >= escCode256FgBase && code <= escCode256FgBase+escCode256Max)
		}

		// Test 256-color sequence with invalid color index (> 255)
		es = EscSeqParser{}
		es.ParseSeq("\x1b[38;5;256m", escSeqKindCSI)
		// Should not have any 256-color codes
		for code := range es.Codes() {
			assert.False(t, code >= escCode256FgBase && code <= escCode256FgBase+escCode256Max)
		}

		// Test 256-color sequence with wrong middle code (not 5)
		es = EscSeqParser{}
		es.ParseSeq("\x1b[38;4;123m", escSeqKindCSI)
		// Should not have any 256-color codes, but might have code 38 and 4
		for code := range es.Codes() {
			assert.False(t, code >= escCode256FgBase && code <= escCode256FgBase+escCode256Max)
		}
	})

	t.Run("256-color with spaces", func(t *testing.T) {
		es := EscSeqParser{}

		// Test parsing 256-color sequence with spaces
		es.ParseSeq("\x1b[ 38 ; 5 ; 123 m", escSeqKindCSI)
		assert.Contains(t, es.Codes(), escCode256FgBase+123)
		assert.Equal(t, "\x1b[38;5;123m", es.Sequence())
	})

	t.Run("256-color multiple replacements", func(t *testing.T) {
		es := EscSeqParser{}

		// Set first 256-color foreground
		es.ParseSeq("\x1b[38;5;100m", escSeqKindCSI)
		assert.Contains(t, es.Codes(), escCode256FgBase+100)
		// Replace with different 256-color foreground
		es.ParseSeq("\x1b[38;5;200m", escSeqKindCSI)
		assert.NotContains(t, es.Codes(), escCode256FgBase+100)
		assert.Contains(t, es.Codes(), escCode256FgBase+200)
		assert.Len(t, es.Codes(), 1)
	})
}

func TestEscSeqParser_ConsumeMalformedSequence(t *testing.T) {
	es := EscSeqParser{}

	// an unterminated sequence should get discarded once it crosses
	// escSeqMaxLength instead of accumulating without limit
	es.Consume(EscapeStartRune)
	es.Consume(EscapeStartRuneCSI)
	for idx := 0; idx < 1000000; idx++ {
		es.Consume('1')
	}
	assert.False(t, es.InSequence())
	assert.LessOrEqual(t, len(es.escapeSeq), escSeqMaxLength+4)
	assert.Empty(t, es.Codes())

	// the parser should recover and parse well-formed sequences that follow
	for _, char := range "\x1b[1m" {
		es.Consume(char)
	}
	assert.Equal(t, []int{1}, es.Codes())
}

func TestEscSeqParser_ControlSequences(t *testing.T) {
	// the Cursor type of this package emits CSI sequences that end in a byte
	// other than "m"; they must end the sequence without changing the state in
	// force, and without swallowing the text that follows them
	for _, test := range []struct {
		name  string
		input string
		want  string
	}{
		{"erase line", EraseLine.Sprint(), ""},
		{"cursor up", CursorUp.Sprint(), ""},
		{"cursor down repeated", CursorDown.Sprintn(3), ""},
		{"cursor left", CursorLeft.Sprint(), ""},
		{"cursor right repeated", CursorRight.Sprintn(2), ""},
		{"cursor position", "\x1b[2;3H", ""},
		{"hide cursor", "\x1b[?25l", ""},
		{"plain text after a cursor sequence", CursorUp.Sprint() + "Ghost", ""},
		{"colour in force after a cursor sequence", "\x1b[31m" + CursorUp.Sprint(), "\x1b[31m"},
		{"colour set after a cursor sequence", CursorUp.Sprint() + "\x1b[31m", "\x1b[31m"},
		{"cursor sequence between two colours", "\x1b[31m" + EraseLine.Sprint() + "\x1b[1m", "\x1b[1;31m"},
		{"reset after a cursor sequence", "\x1b[31m" + CursorUp.Sprint() + "\x1b[0m", ""},
	} {
		t.Run(test.name, func(t *testing.T) {
			es := EscSeqParser{}
			assert.Equal(t, test.want, es.ParseString(test.input))
			assert.False(t, es.InSequence())
		})
	}
}

func TestEscSeqParser_CodeOrder(t *testing.T) {
	// the codes of a sequence apply in the order they appear in, and a
	// terminal has only one foreground and one background color at a time
	for _, test := range []struct {
		name  string
		input string
		want  string
	}{
		{"fg replaces fg", "\x1b[31m\x1b[32m", "\x1b[32m"},
		{"fg replaces fg in one sequence", "\x1b[31;32m", "\x1b[32m"},
		{"bright fg replaces fg", "\x1b[31m\x1b[92m", "\x1b[92m"},
		{"fg replaces bright fg", "\x1b[92m\x1b[31m", "\x1b[31m"},
		{"bg replaces bg", "\x1b[41m\x1b[42m", "\x1b[42m"},
		{"bright bg replaces bg", "\x1b[41m\x1b[102m", "\x1b[102m"},
		{"fg leaves bg and attributes alone", "\x1b[1;31;44m\x1b[32m", "\x1b[1;32;44m"},
		{"bg leaves fg and attributes alone", "\x1b[1;31;44m\x1b[42m", "\x1b[1;31;42m"},
		{"fg replaces 256-color fg", "\x1b[38;5;10m\x1b[31m", "\x1b[31m"},
		{"bg replaces 256-color bg", "\x1b[48;5;10m\x1b[41m", "\x1b[41m"},
		{"256-color fg replaces fg in one sequence", "\x1b[31;38;5;10m", "\x1b[38;5;10m"},
		{"fg replaces 256-color fg in one sequence", "\x1b[38;5;10;31m", "\x1b[31m"},
		{"reset before a color", "\x1b[1;31m\x1b[0;32m", "\x1b[32m"},
		{"reset before a 256-color", "\x1b[1;31m\x1b[0;38;5;10m", "\x1b[38;5;10m"},
		{"reset after a 256-color", "\x1b[38;5;10;0m", ""},
		{"reset without a parameter", "\x1b[1;31m\x1b[m", ""},
		{"reset without a parameter before a color", "\x1b[1;31m\x1b[;32m", "\x1b[32m"},
		{"256-color with a bad index is skipped whole", "\x1b[31m\x1b[38;5;999m", "\x1b[31m"},
		{"256-color without an index is skipped whole", "\x1b[31m\x1b[38;5m", "\x1b[31m"},
		{"38 with nothing after it is not a color", "\x1b[31m\x1b[38m", "\x1b[31;38m"},
		{"48 with an unknown color type is not a color", "\x1b[41m\x1b[48;4m", "\x1b[4;41;48m"},
	} {
		t.Run(test.name, func(t *testing.T) {
			es := EscSeqParser{}
			assert.Equal(t, test.want, es.ParseString(test.input))
		})
	}
}

func TestEscSeqParser_LoneEscapeSequences(t *testing.T) {
	// an ESC that is not followed by "[" or "]" starts a two-byte escape
	// sequence like "ESC 7" (save cursor) or "ESC c" (reset terminal); it must
	// end at its final byte (ECMA-48) without swallowing the text after it
	for _, test := range []struct {
		name  string
		input string
	}{
		{"save cursor", "\x1b7"},
		{"reset terminal", "\x1bc"},
		{"charset selection", "\x1b(B"},
		{"text after save cursor", "\x1b7Ghost"},
		{"text after an unknown sequence", "a\x1bQbc"},
		{"DCS payload until ST", "\x1bP61;1|1;2\x1b\\"},
		{"DCS parameters are not SGR codes", "\x1bP1;2;3\x1b\\"},
		{"string starter after an intermediate byte", "\x1b(X"},
		{"CSI starter after an intermediate byte", "\x1b#["},
		{"OSI starter after an intermediate byte", "\x1b ]"},
	} {
		t.Run(test.name, func(t *testing.T) {
			es := EscSeqParser{}
			assert.Equal(t, "", es.ParseString(test.input))
			assert.False(t, es.InSequence())
			assert.Empty(t, es.Codes())
			assert.False(t, es.IsOpen())
		})
	}
}

func TestEscSeqParser_StringSequences(t *testing.T) {
	// a sequence that carries a payload (OSC, DCS, SOS, PM, APC) ends at the
	// ST ("ESC \"), and a backslash on its own is a part of the payload
	for _, test := range []struct {
		name  string
		input string
		want  string // the sequence in force after the input
	}{
		{"DCS with a backslash", "\x1bPa\\b\x1b\\", ""},
		{"APC with a backslash", "\x1b_a\\b\x1b\\", ""},
		{"OSC with a backslash", "\x1b]0;C:\\dir\x1b\\", ""},
		{"hyperlink with a backslash", "\x1b]8;;http://example.com/a\\b\x1b\\", ""},
		{"empty DCS", "\x1bP\x1b\\", ""},
		// an ESC that is not followed by a backslash cuts the string short
		// and starts a new sequence
		{"OSC cut short by a CSI sequence", "\x1b]0;title\x1b[31m", "\x1b[31m"},
		{"DCS cut short by a plain sequence", "\x1bPabc\x1b7", ""},
		{"SOS cut short by two ESCs", "\x1bXabc\x1b\x1b[31m", "\x1b[31m"},
		{"PM cut short by another string", "\x1b^abc\x1b]0;title\x1b\\", ""},
		{"hyperlink cut short by a CSI sequence", "\x1b]8;;http://example.com\x1b[31m", "\x1b[31m"},
	} {
		t.Run(test.name, func(t *testing.T) {
			es := EscSeqParser{}
			assert.Equal(t, test.want, es.ParseString(test.input))
			assert.False(t, es.InSequence())
			assert.Equal(t, 5, StringWidthWithoutEscSequences(test.input+"Ghost"))
			assert.Equal(t, "Ghost", StripEscape(test.input+"Ghost"))
		})
	}

	// without the ST, the sequence is not over
	for _, input := range []string{"\x1bPa\\b", "\x1b]0;C:\\dir", "\x1b]8;;\\", "\x1bXabc\x1b"} {
		es := EscSeqParser{}
		es.ParseString(input)
		assert.True(t, es.InSequence(), "input=%q", input)
	}

	// the OSI form of a "CSI ... m" sequence is a list of codes that ends at
	// a backslash without an ESC before it
	for _, test := range []struct{ input, want string }{
		{"\x1b]31\\", "\x1b[31m"},
		{"\x1b]1;31\\", "\x1b[1;31m"},
		{"\x1b] 1 ; 31 \\", "\x1b[1;31m"},
		{"\x1b]31\\\x1b]0\\", ""},
		{"\x1b]\\", ""},
	} {
		es := EscSeqParser{}
		assert.Equal(t, test.want, es.ParseString(test.input), "input=%q", test.input)
		assert.False(t, es.InSequence(), "input=%q", test.input)
	}
}

func TestEscSeqParser_OperatingSystemCommands(t *testing.T) {
	// the number that an Operating System Command begins with says which
	// command it is; neither that, nor the numbers in the rest of it are codes
	for _, test := range []struct {
		name  string
		input string
	}{
		{"window title", "\x1b]0;my title\x1b\\"},
		{"window title with numbers", "\x1b]2;1;4;31\x1b\\"},
		{"working directory", "\x1b]7;file://host/home\x1b\\"},
		{"palette color", "\x1b]4;1;rgb:ff/00/00\x1b\\"},
		{"palette reset", "\x1b]104\x1b\\"},
		{"shell integration", "\x1b]133;A\x1b\\"},
		{"hyperlink", "\x1b]8;;http://example.com\x1b\\"},
	} {
		t.Run(test.name, func(t *testing.T) {
			es := EscSeqParser{}
			assert.Equal(t, "", es.ParseString(test.input))
			assert.False(t, es.InSequence())

			// the colors in force are left alone
			assert.Equal(t, "\x1b[1;31m", es.ParseString("\x1b[1;31m"+test.input))
			assert.Equal(t, "\x1b[1;31;44m", es.ParseString(test.input+"\x1b[44m"))
			assert.Equal(t, 5, StringWidthWithoutEscSequences(test.input+"Ghost"))
		})
	}

	// ParseSeq reads nothing but the OSI form of a "CSI ... m" sequence
	es := EscSeqParser{}
	es.ParseSeq("\x1b]7;file://host/home\x1b\\", escSeqKindOSI)
	es.ParseSeq("\x1b]104\x1b\\", escSeqKindOSI)
	es.ParseSeq("\x1b]104\a", escSeqKindOSI)
	es.ParseSeq("1;31\\", escSeqKindOSI)
	assert.Empty(t, es.Codes())
	es.ParseSeq("\x1b]1;31\\", escSeqKindOSI)
	es.ParseSeq("\x1b]4", escSeqKindOSI)
	assert.Equal(t, []int{1, 4, 31}, es.Codes())
}

func TestEscSeqParser_Restart(t *testing.T) {
	// an ESC inside an escape sequence cuts that sequence short, and starts a
	// new one: nothing of the first one is left to garble the second
	for _, test := range []struct {
		name  string
		input string
		want  string
	}{
		{"ESC after ESC", "\x1b\x1b[31m", "\x1b[31m"},
		{"ESC inside a CSI sequence", "\x1b[1;3\x1b[31m", "\x1b[31m"},
		{"ESC inside a plain sequence", "\x1b(\x1b[31m", "\x1b[31m"},
		{"many ESCs", "\x1b\x1b\x1b\x1b[31m", "\x1b[31m"},
		{"the sequence cut short is dropped", "\x1b[32m\x1b[1\x1b7", "\x1b[32m"},
	} {
		t.Run(test.name, func(t *testing.T) {
			es := EscSeqParser{}
			assert.Equal(t, test.want, es.ParseString(test.input))
			assert.False(t, es.InSequence())
			assert.Equal(t, 5, StringWidthWithoutEscSequences(test.input+"Ghost"))
			assert.Equal(t, "Ghost", StripEscape(test.input+"Ghost"))
		})
	}

	// an ESC inside a sequence with a payload is the first half of its ST
	es := EscSeqParser{}
	assert.Equal(t, "", es.ParseString("\x1b]8;;http://example.com\x1b\\"))
	assert.False(t, es.InSequence())
	assert.Equal(t, "", es.ParseString("\x1bP1$r0m\x1b\\"))
	assert.False(t, es.InSequence())
}

func TestEscSeqParser_RGBColors(t *testing.T) {
	// the numbers in "38;2;r;g;b" and "48;2;r;g;b" are color components, and
	// not codes: a 0 in there is no reset, and a 1 does not turn bold on
	for _, test := range []struct {
		name  string
		input string
		want  string
	}{
		{"fg", "\x1b[38;2;255;0;0m", "\x1b[38;2;255;0;0m"},
		{"bg", "\x1b[48;2;0;255;0m", "\x1b[48;2;0;255;0m"},
		{"fg and bg", "\x1b[38;2;1;2;3m\x1b[48;2;4;5;6m", "\x1b[38;2;1;2;3;48;2;4;5;6m"},
		{"fg and bg in one sequence", "\x1b[38;2;1;2;3;48;2;4;5;6m", "\x1b[38;2;1;2;3;48;2;4;5;6m"},
		{"with attributes", "\x1b[1;38;2;1;2;3;4m", "\x1b[1;4;38;2;1;2;3m"},
		{"with spaces", "\x1b[ 38 ; 2 ; 1 ; 2 ; 3 m", "\x1b[38;2;1;2;3m"},
		{"fg replaces fg", "\x1b[38;2;1;2;3m\x1b[38;2;4;5;6m", "\x1b[38;2;4;5;6m"},
		{"fg replaces standard fg", "\x1b[31m\x1b[38;2;1;2;3m", "\x1b[38;2;1;2;3m"},
		{"fg replaces 256-color fg", "\x1b[38;5;10m\x1b[38;2;1;2;3m", "\x1b[38;2;1;2;3m"},
		{"standard fg replaces fg", "\x1b[38;2;1;2;3m\x1b[31m", "\x1b[31m"},
		{"256-color fg replaces fg", "\x1b[38;2;1;2;3m\x1b[38;5;10m", "\x1b[38;5;10m"},
		{"bg replaces standard bg", "\x1b[41m\x1b[48;2;1;2;3m", "\x1b[48;2;1;2;3m"},
		{"standard bg replaces bg", "\x1b[48;2;1;2;3m\x1b[41m", "\x1b[41m"},
		{"fg leaves bg alone", "\x1b[48;2;4;5;6m\x1b[31m", "\x1b[31;48;2;4;5;6m"},
		{"reset", "\x1b[38;2;1;2;3;48;2;4;5;6m\x1b[0m", ""},
		{"reset fg", "\x1b[38;2;1;2;3;48;2;4;5;6m\x1b[39m", "\x1b[48;2;4;5;6m"},
		{"reset bg", "\x1b[38;2;1;2;3;48;2;4;5;6m\x1b[49m", "\x1b[38;2;1;2;3m"},
		{"reset before a color", "\x1b[1m\x1b[0;38;2;1;2;3m", "\x1b[38;2;1;2;3m"},
		{"OSI format", "\x1b]38;2;1;2;3\\", "\x1b[38;2;1;2;3m"},
		{"component out of range is skipped whole", "\x1b[1m\x1b[38;2;256;0;0m", "\x1b[1m"},
		{"missing components are skipped whole", "\x1b[1m\x1b[38;2;0;0m", "\x1b[1m"},
		{"codes after a bad color still apply", "\x1b[38;2;256;0;0;4m", "\x1b[4m"},
	} {
		t.Run(test.name, func(t *testing.T) {
			es := EscSeqParser{}
			assert.Equal(t, test.want, es.ParseString(test.input))
			assert.Equal(t, test.want != "", es.IsOpen())
		})
	}

	// 24-bit colors have no Color to stand for them, and are not in Codes()
	es := EscSeqParser{}
	_, _, _, ok := es.FgRGB()
	assert.False(t, ok)
	_, _, _, ok = es.BgRGB()
	assert.False(t, ok)

	es.ParseString("\x1b[1;38;2;1;2;3;48;2;4;5;6m")
	assert.Equal(t, []int{1}, es.Codes())
	r, g, b, ok := es.FgRGB()
	assert.Equal(t, []int{1, 2, 3}, []int{r, g, b})
	assert.True(t, ok)
	r, g, b, ok = es.BgRGB()
	assert.Equal(t, []int{4, 5, 6}, []int{r, g, b})
	assert.True(t, ok)

	es.ParseString("\x1b[39m")
	_, _, _, ok = es.FgRGB()
	assert.False(t, ok)
	_, _, _, ok = es.BgRGB()
	assert.True(t, ok)
}

func TestEscSeqParser_SkipParse(t *testing.T) {
	// a parser that skips parsing must agree with one that does not on where
	// every escape sequence begins and ends
	for _, input := range []string{
		"",
		"plain text",
		"\x1b[33mGhost\x1b[0m",
		"\x1b[38;5;208mOrange\x1b[0m and \x1b[48;5;21mBlue\x1b[0m",
		"\x1b[KGhost\x1b[3A\x1b[?25l",
		"\x1b]8;;http://example.com\x07label\x1b]8;;\x07",
		"\x1b]8;;http://example.com\x1b\\label\x1b]8;;\x1b\\",
		"\x1b[33mUnterminated\x1b[0",
		"\x1b[" + strings.Repeat("1;", escSeqMaxLength) + "m overflow",
		"\x1b\x1b[1m\u751f\u547d\x1b[22m",
		"\x1b(B\x1b[1mGhost",
	} {
		full, skip := EscSeqParser{}, EscSeqParser{skipParse: true}
		for idx, char := range input {
			full.Consume(char)
			skip.Consume(char)
			assert.Equal(t, full.InSequence(), skip.InSequence(), "input %q, byte %d", input, idx)
		}
		// nothing gets parsed, so nothing is ever seen as being open
		assert.False(t, skip.IsOpen(), "input %q", input)
		assert.Empty(t, skip.Codes(), "input %q", input)
	}
}
