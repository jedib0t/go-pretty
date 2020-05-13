package text

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	EnableColors()
}

func TestColor_EnableAndDisable(t *testing.T) {
	defer EnableColors()

	EnableColors()
	assert.Equal(t, "\x1b[31mtest\x1b[0m", FgRed.Sprint("test"))

	DisableColors()
	assert.Equal(t, "test", FgRed.Sprint("test"))

	EnableColors()
	assert.Equal(t, "\x1b[31mtest\x1b[0m", FgRed.Sprint("test"))
}

func ExampleColor_EscapeSeq() {
	fmt.Printf("Black Background: %#v\n", BgBlack.EscapeSeq())
	fmt.Printf("Black Foreground: %#v\n", FgBlack.EscapeSeq())

	// Output: Black Background: "\x1b[40m"
	// Black Foreground: "\x1b[30m"
}

func TestColor_EscapeSeq(t *testing.T) {
	assert.Equal(t, "\x1b[40m", BgBlack.EscapeSeq())
}

func ExampleColor_HTMLProperty() {
	fmt.Printf("Bold: %#v\n", Bold.HTMLProperty())
	fmt.Printf("Black Background: %#v\n", BgBlack.HTMLProperty())
	fmt.Printf("Black Foreground: %#v\n", FgBlack.HTMLProperty())

	// Output: Bold: "class=\"bold\""
	// Black Background: "class=\"bg-black\""
	// Black Foreground: "class=\"fg-black\""
}

func TestColor_HTMLProperty(t *testing.T) {
	assert.Equal(t, "class=\"bold\"", Bold.HTMLProperty())
	assert.Equal(t, "class=\"bg-black\"", BgBlack.HTMLProperty())
	assert.Equal(t, "class=\"fg-black\"", FgBlack.HTMLProperty())
}

func ExampleColor_Sprint() {
	fmt.Printf("%#v\n", BgBlack.Sprint("Black Background"))
	fmt.Printf("%#v\n", FgBlack.Sprint("Black Foreground"))

	// Output: "\x1b[40mBlack Background\x1b[0m"
	// "\x1b[30mBlack Foreground\x1b[0m"
}

func TestColor_Sprint(t *testing.T) {
	assert.Equal(t, "\x1b[31mtest true\x1b[0m", FgRed.Sprint("test ", true))

	assert.Equal(t, "\x1b[32mtest\x1b[0m\x1b[31mtrue\x1b[0m", FgRed.Sprint("\x1b[32mtest\x1b[0m", true))
	assert.Equal(t, "\x1b[32mtest true\x1b[0m", FgRed.Sprint("\x1b[32mtest ", true))
	assert.Equal(t, "\x1b[32mtest\x1b[0m\x1b[31m \x1b[0m", FgRed.Sprint("\x1b[32mtest\x1b[0m "))
	assert.Equal(t, "\x1b[32mtest\x1b[0m", FgRed.Sprint("\x1b[32mtest\x1b[0m"))
}

func ExampleColor_Sprintf() {
	fmt.Printf("%#v\n", BgBlack.Sprintf("%s %s", "Black", "Background"))
	fmt.Printf("%#v\n", FgBlack.Sprintf("%s %s", "Black", "Foreground"))

	// Output: "\x1b[40mBlack Background\x1b[0m"
	// "\x1b[30mBlack Foreground\x1b[0m"
}

func TestColor_Sprintf(t *testing.T) {
	assert.Equal(t, "\x1b[31mtest true\x1b[0m", FgRed.Sprintf("test %s", "true"))
}

func ExampleColors_EscapeSeq() {
	fmt.Printf("Black Background: %#v\n", Colors{BgBlack}.EscapeSeq())
	fmt.Printf("Black Foreground: %#v\n", Colors{FgBlack}.EscapeSeq())
	fmt.Printf("Black Background, White Foreground: %#v\n", Colors{BgBlack, FgWhite}.EscapeSeq())
	fmt.Printf("Black Foreground, White Background: %#v\n", Colors{FgBlack, BgWhite}.EscapeSeq())

	// Output: Black Background: "\x1b[40m"
	// Black Foreground: "\x1b[30m"
	// Black Background, White Foreground: "\x1b[40;37m"
	// Black Foreground, White Background: "\x1b[30;47m"
}

func TestColors_EscapeSeq(t *testing.T) {
	assert.Equal(t, "", Colors{}.EscapeSeq())
	assert.Equal(t, "\x1b[40;37m", Colors{BgBlack, FgWhite}.EscapeSeq())
}

func ExampleColors_HTMLProperty() {
	fmt.Printf("Black Background: %#v\n", Colors{BgBlack}.HTMLProperty())
	fmt.Printf("Black Foreground: %#v\n", Colors{FgBlack}.HTMLProperty())
	fmt.Printf("Black Background, White Foreground: %#v\n", Colors{BgBlack, FgWhite}.HTMLProperty())
	fmt.Printf("Black Foreground, White Background: %#v\n", Colors{FgBlack, BgWhite}.HTMLProperty())
	fmt.Printf("Bold Italic Underline Red Text: %#v\n", Colors{Bold, Italic, Underline, FgRed}.HTMLProperty())

	// Output: Black Background: "class=\"bg-black\""
	// Black Foreground: "class=\"fg-black\""
	// Black Background, White Foreground: "class=\"bg-black fg-white\""
	// Black Foreground, White Background: "class=\"bg-white fg-black\""
	// Bold Italic Underline Red Text: "class=\"bold fg-red italic underline\""
}

func TestColors_HTMLProperty(t *testing.T) {
	assert.Equal(t, "", Colors{}.HTMLProperty())
	assert.Equal(t, "class=\"bg-black fg-white\"", Colors{BgBlack, FgWhite}.HTMLProperty())
	assert.Equal(t, "class=\"bold fg-red\"", Colors{Bold, FgRed}.HTMLProperty())
}

func ExampleColors_Sprint() {
	fmt.Printf("%#v\n", Colors{BgBlack}.Sprint("Black Background"))
	fmt.Printf("%#v\n", Colors{BgBlack, FgWhite}.Sprint("Black Background, White Foreground"))
	fmt.Printf("%#v\n", Colors{FgBlack}.Sprint("Black Foreground"))
	fmt.Printf("%#v\n", Colors{FgBlack, BgWhite}.Sprint("Black Foreground, White Background"))

	// Output: "\x1b[40mBlack Background\x1b[0m"
	// "\x1b[40;37mBlack Background, White Foreground\x1b[0m"
	// "\x1b[30mBlack Foreground\x1b[0m"
	// "\x1b[30;47mBlack Foreground, White Background\x1b[0m"
}

func TestColors_Sprint(t *testing.T) {
	assert.Equal(t, "test true", Colors{}.Sprint("test ", true))
	assert.Equal(t, "\x1b[31mtest true\x1b[0m", Colors{FgRed}.Sprint("test ", true))

	assert.Equal(t, "\x1b[32mtest\x1b[0m\x1b[31mtrue\x1b[0m", Colors{FgRed}.Sprint("\x1b[32mtest\x1b[0m", true))
	assert.Equal(t, "\x1b[32mtest true\x1b[0m", Colors{FgRed}.Sprint("\x1b[32mtest ", true))
	assert.Equal(t, "\x1b[32mtest\x1b[0m\x1b[31m \x1b[0m", Colors{FgRed}.Sprint("\x1b[32mtest\x1b[0m "))
	assert.Equal(t, "\x1b[32mtest\x1b[0m", Colors{FgRed}.Sprint("\x1b[32mtest\x1b[0m"))
}

func ExampleColors_Sprintf() {
	fmt.Printf("%#v\n", Colors{BgBlack}.Sprintf("%s %s", "Black", "Background"))
	fmt.Printf("%#v\n", Colors{BgBlack, FgWhite}.Sprintf("%s, %s", "Black Background", "White Foreground"))
	fmt.Printf("%#v\n", Colors{FgBlack}.Sprintf("%s %s", "Black", "Foreground"))
	fmt.Printf("%#v\n", Colors{FgBlack, BgWhite}.Sprintf("%s, %s", "Black Foreground", "White Background"))

	// Output: "\x1b[40mBlack Background\x1b[0m"
	// "\x1b[40;37mBlack Background, White Foreground\x1b[0m"
	// "\x1b[30mBlack Foreground\x1b[0m"
	// "\x1b[30;47mBlack Foreground, White Background\x1b[0m"
}

func TestColors_Sprintf(t *testing.T) {
	assert.Equal(t, "test true", Colors{}.Sprintf("test %s", "true"))
	assert.Equal(t, "\x1b[31mtest true\x1b[0m", Colors{FgRed}.Sprintf("test %s", "true"))
}
