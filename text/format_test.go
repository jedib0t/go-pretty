package text

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleFormat_Apply() {
	text := "jon Snow"
	fmt.Printf("FormatDefault: '%s'\n", FormatDefault.Apply(text))
	fmt.Printf("FormatLower  : '%s'\n", FormatLower.Apply(text))
	fmt.Printf("FormatTitle  : '%s'\n", FormatTitle.Apply(text))
	fmt.Printf("FormatUpper  : '%s'\n", FormatUpper.Apply(text))

	// Output: FormatDefault: 'jon Snow'
	// FormatLower  : 'jon snow'
	// FormatTitle  : 'Jon Snow'
	// FormatUpper  : 'JON SNOW'
}

func TestFormat_Apply(t *testing.T) {
	text := "A big croc0dile; Died - Empty_fanged ツ \u2008."
	assert.Equal(t, text, FormatDefault.Apply(text))
	assert.Equal(t, "a big croc0dile; died - empty_fanged ツ \u2008.", FormatLower.Apply(text))
	assert.Equal(t, "A Big Croc0dile; Died - Empty_fanged ツ \u2008.", FormatTitle.Apply(text))
	assert.Equal(t, "A BIG CROC0DILE; DIED - EMPTY_FANGED ツ \u2008.", FormatUpper.Apply(text))

	// test with escape sequences
	text = Colors{Bold}.Sprint(text)
	assert.Equal(t, "\x1b[1mA big croc0dile; Died - Empty_fanged ツ \u2008.\x1b[0m", FormatDefault.Apply(text))
	assert.Equal(t, "\x1b[1ma big croc0dile; died - empty_fanged ツ \u2008.\x1b[0m", FormatLower.Apply(text))
	assert.Equal(t, "\x1b[1mA Big Croc0dile; Died - Empty_fanged ツ \u2008.\x1b[0m", FormatTitle.Apply(text))
	assert.Equal(t, "\x1b[1mA BIG CROC0DILE; DIED - EMPTY_FANGED ツ \u2008.\x1b[0m", FormatUpper.Apply(text))
}
