package text

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleFormat_Apply() {
	fmt.Printf("FormatDefault: %#v\n", FormatDefault.Apply("jon Snow"))
	fmt.Printf("FormatLower  : %#v\n", FormatLower.Apply("jon Snow"))
	fmt.Printf("FormatTitle  : %#v\n", FormatTitle.Apply("jon Snow"))
	fmt.Printf("FormatUpper  : %#v\n", FormatUpper.Apply("jon Snow"))
	fmt.Println()
	fmt.Printf("FormatDefault (w/EscSeq): %#v\n", FormatDefault.Apply(Bold.Sprint("jon Snow")))
	fmt.Printf("FormatLower   (w/EscSeq): %#v\n", FormatLower.Apply(Bold.Sprint("jon Snow")))
	fmt.Printf("FormatTitle   (w/EscSeq): %#v\n", FormatTitle.Apply(Bold.Sprint("jon Snow")))
	fmt.Printf("FormatUpper   (w/EscSeq): %#v\n", FormatUpper.Apply(Bold.Sprint("jon Snow")))

	// Output: FormatDefault: "jon Snow"
	// FormatLower  : "jon snow"
	// FormatTitle  : "Jon Snow"
	// FormatUpper  : "JON SNOW"
	//
	// FormatDefault (w/EscSeq): "\x1b[1mjon Snow\x1b[0m"
	// FormatLower   (w/EscSeq): "\x1b[1mjon snow\x1b[0m"
	// FormatTitle   (w/EscSeq): "\x1b[1mJon Snow\x1b[0m"
	// FormatUpper   (w/EscSeq): "\x1b[1mJON SNOW\x1b[0m"
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
