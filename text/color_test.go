package text

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColor_GetEscapeSeq(t *testing.T) {
	assert.Equal(t, "\x1b[40m", BgBlack.GetEscapeSeq())
}

func TestColor_Sprint(t *testing.T) {
	assert.Equal(t, "\x1b[31mtest true\x1b[0m", FgRed.Sprint("test ", true))

	assert.Equal(t, "\x1b[32mtest\x1b[0m\x1b[31mtrue\x1b[0m", FgRed.Sprint("\x1b[32mtest\x1b[0m", true))
	assert.Equal(t, "\x1b[32mtest true\x1b[0m", FgRed.Sprint("\x1b[32mtest ", true))
	assert.Equal(t, "\x1b[32mtest\x1b[0m\x1b[31m \x1b[0m", FgRed.Sprint("\x1b[32mtest\x1b[0m "))
	assert.Equal(t, "\x1b[32mtest\x1b[0m", FgRed.Sprint("\x1b[32mtest\x1b[0m"))
}

func TestColor_Sprintf(t *testing.T) {
	assert.Equal(t, "\x1b[31mtest true\x1b[0m", FgRed.Sprintf("test %s", "true"))
}

func TestColors_GetEscapeSeq(t *testing.T) {
	assert.Equal(t, "", Colors{}.GetEscapeSeq())
	assert.Equal(t, "\x1b[40;37m", Colors{BgBlack, FgWhite}.GetEscapeSeq())
}

func TestColors_Sprint(t *testing.T) {
	assert.Equal(t, "test true", Colors{}.Sprint("test ", true))
	assert.Equal(t, "\x1b[31mtest true\x1b[0m", Colors{FgRed}.Sprint("test ", true))

	assert.Equal(t, "\x1b[32mtest\x1b[0m\x1b[31mtrue\x1b[0m", Colors{FgRed}.Sprint("\x1b[32mtest\x1b[0m", true))
	assert.Equal(t, "\x1b[32mtest true\x1b[0m", Colors{FgRed}.Sprint("\x1b[32mtest ", true))
	assert.Equal(t, "\x1b[32mtest\x1b[0m\x1b[31m \x1b[0m", Colors{FgRed}.Sprint("\x1b[32mtest\x1b[0m "))
	assert.Equal(t, "\x1b[32mtest\x1b[0m", Colors{FgRed}.Sprint("\x1b[32mtest\x1b[0m"))
}

func TestColors_Sprintf(t *testing.T) {
	assert.Equal(t, "test true", Colors{}.Sprintf("test %s", "true"))
	assert.Equal(t, "\x1b[31mtest true\x1b[0m", Colors{FgRed}.Sprintf("test %s", "true"))
}
