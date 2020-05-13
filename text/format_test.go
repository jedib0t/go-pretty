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
	text := "A big crocodile, Died. Empty."
	assert.Equal(t, text, FormatDefault.Apply(text))
	assert.Equal(t, "a big crocodile, died. empty.", FormatLower.Apply(text))
	assert.Equal(t, "A Big Crocodile, Died. Empty.", FormatTitle.Apply(text))
	assert.Equal(t, "A BIG CROCODILE, DIED. EMPTY.", FormatUpper.Apply(text))

	// test with escape sequences
	text = Colors{Bold}.Sprint(text)
	assert.Equal(t, "\x1b[1mA big crocodile, Died. Empty.\x1b[0m", FormatDefault.Apply(text))
	assert.Equal(t, "\x1b[1ma big crocodile, died. empty.\x1b[0m", FormatLower.Apply(text))
	assert.Equal(t, "\x1b[1mA Big Crocodile, Died. Empty.\x1b[0m", FormatTitle.Apply(text))
	assert.Equal(t, "\x1b[1mA BIG CROCODILE, DIED. EMPTY.\x1b[0m", FormatUpper.Apply(text))
}
