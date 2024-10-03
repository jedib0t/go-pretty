package text

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_escSeqParser(t *testing.T) {
	t.Run("extract", func(t *testing.T) {
		es := escSeqParser{}

		assert.Equal(t, "\x1b[1;3;4;5;7;9;91m", es.Extract("\x1b[91m\x1b[1m\x1b[3m\x1b[4m\x1b[5m\x1b[7m\x1b[9m Spicy"))
		assert.Equal(t, "\x1b[3;4;5;7;9;91m", es.Extract("\x1b[22m No Bold"))
		assert.Equal(t, "\x1b[4;5;7;9;91m", es.Extract("\x1b[23m No Italic"))
		assert.Equal(t, "\x1b[5;7;9;91m", es.Extract("\x1b[24m No Underline"))
		assert.Equal(t, "\x1b[7;9;91m", es.Extract("\x1b[25m No Blink"))
		assert.Equal(t, "\x1b[9;91m", es.Extract("\x1b[27m No Reverse"))
		assert.Equal(t, "\x1b[91m", es.Extract("\x1b[29m No Crossed-Out"))
		assert.Equal(t, "", es.Extract("\x1b[0m Resetted"))
	})

	t.Run("parse", func(t *testing.T) {
		es := escSeqParser{}

		es.Parse("\x1b[91m") // color
		es.Parse("\x1b[1m")  // bold
		assert.Len(t, es.Codes(), 2)
		assert.True(t, es.IsOpen())
		assert.Equal(t, "\x1b[1;91m", es.Sequence())

		es.Parse("\x1b[22m") // un-bold
		assert.Len(t, es.Codes(), 1)
		assert.True(t, es.IsOpen())
		assert.Equal(t, "\x1b[91m", es.Sequence())

		es.Parse("\x1b[0m") // reset
		assert.Empty(t, es.Codes())
		assert.False(t, es.IsOpen())
		assert.Empty(t, es.Sequence())
	})
}
