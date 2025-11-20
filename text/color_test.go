package text

import (
	"fmt"
	"os"
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

func TestColor_areColorsOnInTheEnv(t *testing.T) {
	// Default: TERM matters when FORCE_COLOR and NO_COLOR are not set
	os.Setenv("FORCE_COLOR", "0")
	os.Setenv("NO_COLOR", "0")
	os.Setenv("TERM", "xterm")
	assert.True(t, areColorsOnInTheEnv())

	os.Setenv("TERM", "dumb")
	assert.False(t, areColorsOnInTheEnv())

	// NO_COLOR disables colors (overrides TERM)
	os.Setenv("FORCE_COLOR", "0")
	os.Setenv("NO_COLOR", "1")
	os.Setenv("TERM", "xterm")
	assert.False(t, areColorsOnInTheEnv())

	// FORCE_COLOR enables colors (overrides NO_COLOR and TERM)
	os.Setenv("FORCE_COLOR", "1")
	os.Setenv("NO_COLOR", "1")
	os.Setenv("TERM", "dumb")
	assert.True(t, areColorsOnInTheEnv())

	// FORCE_COLOR alternative values
	os.Setenv("FORCE_COLOR", "true")
	assert.True(t, areColorsOnInTheEnv())

	// FORCE_COLOR=false is treated as falsy
	os.Setenv("FORCE_COLOR", "false")
	os.Setenv("NO_COLOR", "0")
	os.Setenv("TERM", "xterm")
	assert.True(t, areColorsOnInTheEnv())
}

func ExampleColor_CSSClasses() {
	fmt.Printf("Bold: %#v\n", Bold.CSSClasses())
	fmt.Printf("Black Background: %#v\n", BgBlack.CSSClasses())
	fmt.Printf("Black Foreground: %#v\n", FgBlack.CSSClasses())

	// Output: Bold: "bold"
	// Black Background: "bg-black"
	// Black Foreground: "fg-black"
}

func TestColor_CSSClasses(t *testing.T) {
	assert.Equal(t, "bold", Bold.CSSClasses())
	assert.Equal(t, "bg-black", BgBlack.CSSClasses())
	assert.Equal(t, "fg-black", FgBlack.CSSClasses())
	assert.Equal(t, "", Reset.CSSClasses()) // Reset has no CSS class
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
	assert.Equal(t, "", Reset.HTMLProperty()) // Reset has no CSS class mapping
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

func ExampleColors_CSSClasses() {
	fmt.Printf("Black Background: %#v\n", Colors{BgBlack}.CSSClasses())
	fmt.Printf("Black Foreground: %#v\n", Colors{FgBlack}.CSSClasses())
	fmt.Printf("Black Background, White Foreground: %#v\n", Colors{BgBlack, FgWhite}.CSSClasses())
	fmt.Printf("Black Foreground, White Background: %#v\n", Colors{FgBlack, BgWhite}.CSSClasses())
	fmt.Printf("Bold Italic Underline Red Text: %#v\n", Colors{Bold, Italic, Underline, FgRed}.CSSClasses())

	// Output: Black Background: "bg-black"
	// Black Foreground: "fg-black"
	// Black Background, White Foreground: "bg-black fg-white"
	// Black Foreground, White Background: "bg-white fg-black"
	// Bold Italic Underline Red Text: "bold fg-red italic underline"
}

func TestColors_CSSClasses(t *testing.T) {
	assert.Equal(t, "", Colors{}.CSSClasses())
	assert.Equal(t, "bg-black fg-white", Colors{BgBlack, FgWhite}.CSSClasses())
	assert.Equal(t, "bold fg-red", Colors{Bold, FgRed}.CSSClasses())
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

// 256-color tests

func TestFg256Color(t *testing.T) {
	// Valid indices
	assert.Equal(t, fg256Start+0, Fg256Color(0))
	assert.Equal(t, fg256Start+123, Fg256Color(123))
	assert.Equal(t, fg256Start+255, Fg256Color(255))

	// Invalid indices should return Reset
	assert.Equal(t, Reset, Fg256Color(-1))
	assert.Equal(t, Reset, Fg256Color(256))
	assert.Equal(t, Reset, Fg256Color(1000))
}

func TestBg256Color(t *testing.T) {
	// Valid indices
	assert.Equal(t, bg256Start+0, Bg256Color(0))
	assert.Equal(t, bg256Start+200, Bg256Color(200))
	assert.Equal(t, bg256Start+255, Bg256Color(255))

	// Invalid indices should return Reset
	assert.Equal(t, Reset, Bg256Color(-1))
	assert.Equal(t, Reset, Bg256Color(256))
	assert.Equal(t, Reset, Bg256Color(1000))
}

func TestFg256RGB(t *testing.T) {
	// Valid RGB values (0-5)
	assert.Equal(t, fg256Start+16, Fg256RGB(0, 0, 0))  // First color in cube
	assert.Equal(t, fg256Start+51, Fg256RGB(0, 5, 5))  // r=0, g=5, b=5
	assert.Equal(t, fg256Start+231, Fg256RGB(5, 5, 5)) // Last color in cube

	// Invalid RGB values should return Reset
	assert.Equal(t, Reset, Fg256RGB(-1, 0, 0))
	assert.Equal(t, Reset, Fg256RGB(0, -1, 0))
	assert.Equal(t, Reset, Fg256RGB(0, 0, -1))
	assert.Equal(t, Reset, Fg256RGB(6, 0, 0))
	assert.Equal(t, Reset, Fg256RGB(0, 6, 0))
	assert.Equal(t, Reset, Fg256RGB(0, 0, 6))
}

func TestBg256RGB(t *testing.T) {
	// Valid RGB values (0-5)
	assert.Equal(t, bg256Start+16, Bg256RGB(0, 0, 0))  // First color in cube
	assert.Equal(t, bg256Start+51, Bg256RGB(0, 5, 5))  // r=0, g=5, b=5
	assert.Equal(t, bg256Start+231, Bg256RGB(5, 5, 5)) // Last color in cube

	// Invalid RGB values should return Reset
	assert.Equal(t, Reset, Bg256RGB(-1, 0, 0))
	assert.Equal(t, Reset, Bg256RGB(0, -1, 0))
	assert.Equal(t, Reset, Bg256RGB(0, 0, -1))
	assert.Equal(t, Reset, Bg256RGB(6, 0, 0))
	assert.Equal(t, Reset, Bg256RGB(0, 6, 0))
	assert.Equal(t, Reset, Bg256RGB(0, 0, 6))
}

func TestColor_EscapeSeq_256Color(t *testing.T) {
	// 256-color foreground
	assert.Equal(t, "\x1b[38;5;0m", Fg256Color(0).EscapeSeq())
	assert.Equal(t, "\x1b[38;5;123m", Fg256Color(123).EscapeSeq())
	assert.Equal(t, "\x1b[38;5;255m", Fg256Color(255).EscapeSeq())

	// 256-color background
	assert.Equal(t, "\x1b[48;5;0m", Bg256Color(0).EscapeSeq())
	assert.Equal(t, "\x1b[48;5;200m", Bg256Color(200).EscapeSeq())
	assert.Equal(t, "\x1b[48;5;255m", Bg256Color(255).EscapeSeq())

	// Regular colors still work
	assert.Equal(t, "\x1b[31m", FgRed.EscapeSeq())
	assert.Equal(t, "\x1b[40m", BgBlack.EscapeSeq())
}

func TestColors_EscapeSeq_256Color(t *testing.T) {
	// Single 256-color
	assert.Equal(t, "\x1b[38;5;123m", Colors{Fg256Color(123)}.EscapeSeq())
	assert.Equal(t, "\x1b[48;5;200m", Colors{Bg256Color(200)}.EscapeSeq())

	// Multiple 256-colors
	assert.Equal(t, "\x1b[38;5;100;48;5;200m", Colors{Fg256Color(100), Bg256Color(200)}.EscapeSeq())

	// Mixed regular and 256-colors
	assert.Equal(t, "\x1b[1;38;5;123m", Colors{Bold, Fg256Color(123)}.EscapeSeq())
	assert.Equal(t, "\x1b[31;48;5;200m", Colors{FgRed, Bg256Color(200)}.EscapeSeq())
	assert.Equal(t, "\x1b[1;3;38;5;100;48;5;200m", Colors{Bold, Italic, Fg256Color(100), Bg256Color(200)}.EscapeSeq())
}

func TestColor_CSSClasses_256Color(t *testing.T) {
	// 256-color foreground
	assert.Equal(t, "fg-256-0-0-0", Fg256Color(0).CSSClasses())   // Black
	assert.Equal(t, "fg-256-255-0-0", Fg256Color(9).CSSClasses()) // Bright red
	assert.Contains(t, Fg256Color(123).CSSClasses(), "fg-256-")

	// 256-color background
	assert.Equal(t, "bg-256-0-0-0", Bg256Color(0).CSSClasses())   // Black
	assert.Equal(t, "bg-256-255-0-0", Bg256Color(9).CSSClasses()) // Bright red
	assert.Contains(t, Bg256Color(200).CSSClasses(), "bg-256-")

	// Regular colors still work
	assert.Equal(t, "fg-red", FgRed.CSSClasses())
	assert.Equal(t, "bg-black", BgBlack.CSSClasses())
}

func TestColors_CSSClasses_256Color(t *testing.T) {
	// Single 256-color
	assert.Contains(t, Colors{Fg256Color(123)}.CSSClasses(), "fg-256-")
	assert.Contains(t, Colors{Bg256Color(200)}.CSSClasses(), "bg-256-")

	// Mixed regular and 256-colors
	classes := Colors{Bold, Fg256Color(100)}.CSSClasses()
	assert.Contains(t, classes, "bold")
	assert.Contains(t, classes, "fg-256-")

	classes = Colors{FgRed, Bg256Color(200)}.CSSClasses()
	assert.Contains(t, classes, "fg-red")
	assert.Contains(t, classes, "bg-256-")
}

func TestColor_Sprint_256Color(t *testing.T) {
	EnableColors()
	defer EnableColors()

	assert.Equal(t, "\x1b[38;5;123mtest\x1b[0m", Fg256Color(123).Sprint("test"))
	assert.Equal(t, "\x1b[48;5;200mtest\x1b[0m", Bg256Color(200).Sprint("test"))

	DisableColors()
	assert.Equal(t, "test", Fg256Color(123).Sprint("test"))
	EnableColors()
}

func TestColor_Sprintf_256Color(t *testing.T) {
	EnableColors()
	defer EnableColors()

	assert.Equal(t, "\x1b[38;5;123mtest 123\x1b[0m", Fg256Color(123).Sprintf("test %d", 123))
	assert.Equal(t, "\x1b[48;5;200mtest 200\x1b[0m", Bg256Color(200).Sprintf("test %d", 200))
}

func TestColors_Sprint_256Color(t *testing.T) {
	EnableColors()
	defer EnableColors()

	assert.Equal(t, "\x1b[38;5;100;48;5;200mtest\x1b[0m", Colors{Fg256Color(100), Bg256Color(200)}.Sprint("test"))
	assert.Equal(t, "\x1b[1;38;5;123mtest\x1b[0m", Colors{Bold, Fg256Color(123)}.Sprint("test"))
}

func TestColors_Sprintf_256Color(t *testing.T) {
	EnableColors()
	defer EnableColors()

	assert.Equal(t, "\x1b[38;5;100;48;5;200mtest 100 200\x1b[0m", Colors{Fg256Color(100), Bg256Color(200)}.Sprintf("test %d %d", 100, 200))
}

func TestColor256ToRGB_Indirect(t *testing.T) {
	// Test RGB conversion indirectly through CSS classes
	// Standard 16 colors (0-15)
	assert.Equal(t, "fg-256-0-0-0", Fg256Color(0).CSSClasses())   // Black
	assert.Equal(t, "fg-256-255-0-0", Fg256Color(9).CSSClasses()) // Bright red

	// RGB cube (16-231) - test first and last
	css16 := Fg256Color(16).CSSClasses()
	assert.Contains(t, css16, "fg-256-")
	css231 := Fg256Color(231).CSSClasses()
	assert.Contains(t, css231, "fg-256-")
	assert.Contains(t, css231, "255") // Should be white (255,255,255)

	// Grayscale (232-255) - test first and last
	css232 := Fg256Color(232).CSSClasses()
	assert.Contains(t, css232, "fg-256-")
	css255 := Fg256Color(255).CSSClasses()
	assert.Contains(t, css255, "fg-256-")
}
