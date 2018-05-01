package text

import (
	"fmt"

	"github.com/fatih/color"
)

// Colors represents an array of color.Attributes to define the color to
// render text with. Example: Colors{color.FgCyan, color.BgBlack}
type Colors []color.Attribute

// GetColorizer returns a *color.Color object based on the color attributes set.
func (tc Colors) GetColorizer() *color.Color {
	var colorizer *color.Color
	if len(tc) > 0 {
		colorizer = color.New(tc...)
		colorizer.EnableColor()
	}
	return colorizer
}

// Sprint colorizes and prints the given string(s).
func (tc Colors) Sprint(a ...interface{}) string {
	colorizer := tc.GetColorizer()
	if colorizer != nil {
		return colorizer.Sprint(a...)
	}
	return fmt.Sprint(a...)
}

// Sprintf formats and colorizes and prints the given string(s).
func (tc Colors) Sprintf(format string, a ...interface{}) string {
	colorizer := tc.GetColorizer()
	if colorizer != nil {
		return colorizer.Sprintf(format, a...)
	}
	return fmt.Sprintf(format, a...)
}
