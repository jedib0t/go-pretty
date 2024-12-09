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
}
