package text

import (
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
		// After consuming BEL, the parser should have reset
		// Note: code 8 (conceal) will be parsed before reset
		assert.False(t, es.InSequence())
		assert.True(t, es.IsOpen())
		assert.Contains(t, es.Codes(), 8)

		// Test with a full hyperlink sequence
		es = EscSeqParser{} // fresh parser
		result := es.ParseString("\x1b]8;;https://example.com\x07Click here\x1b]8;;\x07")
		// Code 8 (conceal) will be parsed
		assert.Equal(t, "\x1b[8m", result)
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
