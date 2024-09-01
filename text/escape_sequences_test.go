package text

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_escSeqParser(t *testing.T) {
	t.Run("extract", func(t *testing.T) {
		es := escSeqParser{}

		assert.Equal(t, "\x1b[1;91m", es.Extract("\x1b[91m\x1b[1m Bold text"))
		assert.Equal(t, "\x1b[91m", es.Extract("\x1b[22m Regular text"))
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
