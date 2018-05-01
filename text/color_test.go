package text

import (
	"testing"

	"github.com/fatih/color"
	"github.com/stretchr/testify/assert"
)

func TestColors_GetColorizer(t *testing.T) {
	assert.Nil(t, Colors{}.GetColorizer())
	assert.NotNil(t, Colors{color.BgBlack, color.FgWhite}.GetColorizer())
}

func TestColors_Sprint(t *testing.T) {
	assert.Equal(t, "test true", Colors{}.Sprint("test ", true))
	assert.Equal(t, "\x1b[31mtest true\x1b[0m", Colors{color.FgRed}.Sprint("test ", true))
}

func TestColors_Sprintf(t *testing.T) {
	assert.Equal(t, "test true", Colors{}.Sprintf("test %s", "true"))
	assert.Equal(t, "\x1b[31mtest true\x1b[0m", Colors{color.FgRed}.Sprintf("test %s", "true"))
}
