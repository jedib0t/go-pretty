package text

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_escSeqParser(t *testing.T) {
	t.Run("extract csi", func(t *testing.T) {
		es := escSeqParser{}

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
		es := escSeqParser{}

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
		es := escSeqParser{}

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
		es := escSeqParser{}

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
		es := escSeqParser{}

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
		es = escSeqParser{} // fresh parser
		result := es.ParseString("\x1b]8;;https://example.com\x07Click here\x1b]8;;\x07")
		// Code 8 (conceal) will be parsed
		assert.Equal(t, "\x1b[8m", result)
	})

	t.Run("consume directly", func(t *testing.T) {
		es := escSeqParser{}

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
		es := escSeqParser{}

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
		es := escSeqParser{}

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
	})

	t.Run("parse seq edge cases", func(t *testing.T) {
		es := escSeqParser{}

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
}
