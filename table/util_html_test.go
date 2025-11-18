package table

import (
	"testing"

	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/stretchr/testify/assert"
)

func Test_convertEscSequencesToSpans(t *testing.T) {
	t.Run("basic cases", func(t *testing.T) {
		tests := []struct {
			name     string
			input    string
			expected string
		}{
			{"empty string", "", ""},
			{"no escape sequences", "Hello World", "Hello World"},
			{"single foreground color", text.FgRed.Sprint("Red Text"), "<span class=\"fg-red\">Red Text</span>"},
			{"single background color", text.BgBlue.Sprint("Blue Background"), "<span class=\"bg-blue\">Blue Background</span>"},
			{"only escape sequence with reset", text.FgYellow.EscapeSeq() + text.Reset.EscapeSeq(), "<span class=\"fg-yellow\"></span>"},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := convertEscSequencesToSpans(tt.input)
				assert.Equal(t, tt.expected, result)
			})
		}
	})

	t.Run("multiple colors and attributes", func(t *testing.T) {
		tests := []struct {
			name     string
			input    string
			expected string
		}{
			{"two colors", text.Colors{text.FgRed, text.Bold}.Sprint("Bold Red"), "<span class=\"bold fg-red\">Bold Red</span>"},
			{"three attributes", text.Colors{text.Bold, text.Italic, text.Underline}.Sprint("Bold Italic Underline"), "<span class=\"bold italic underline\">Bold Italic Underline</span>"},
			{"complex combination", text.Colors{text.FgCyan, text.Bold, text.Underline}.Sprint("Styled"), "<span class=\"bold fg-cyan underline\">Styled</span>"},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := convertEscSequencesToSpans(tt.input)
				assert.Equal(t, tt.expected, result)
			})
		}
	})

	t.Run("color changes and resets", func(t *testing.T) {
		tests := []struct {
			name     string
			input    string
			expected string
		}{
			{"color reset", text.FgRed.Sprint("Red") + text.Reset.Sprint("Normal"), "<span class=\"fg-red\">Red</span>Normal"},
			{"multiple color changes", "Start" + text.FgRed.Sprint("Red") + text.FgBlue.Sprint("Blue") + "End", "Start<span class=\"fg-red\">Red</span><span class=\"fg-blue\">Blue</span>End"},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := convertEscSequencesToSpans(tt.input)
				assert.Equal(t, tt.expected, result)
			})
		}
	})

	t.Run("HTML escaping", func(t *testing.T) {
		tests := []struct {
			name     string
			input    string
			expected string
		}{
			{"all HTML special chars with color", text.FgBlue.Sprint("Test < > & \" '"), "<span class=\"fg-blue\">Test &lt; &gt; &amp; &#34; &#39;</span>"},
			{"HTML special chars without color", "Text <bold> & \"quoted\"", "Text &lt;bold&gt; &amp; &#34;quoted&#34;"},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := convertEscSequencesToSpans(tt.input)
				assert.Equal(t, tt.expected, result)
			})
		}
	})

	t.Run("edge cases", func(t *testing.T) {
		tests := []struct {
			name     string
			input    string
			expected string
		}{
			{"color with newlines", text.FgMagenta.Sprint("Line1\nLine2"), "<span class=\"fg-magenta\">Line1\nLine2</span>"},
			{"unmapped color code", "\x1b[99mText\x1b[0m", "Text"},
			{"same colors set twice", text.FgRed.EscapeSeq() + "Red1" + text.FgRed.EscapeSeq() + "Red2" + text.Reset.EscapeSeq(), "<span class=\"fg-red\">Red1Red2</span>"},
			{"open span at end", text.FgRed.EscapeSeq() + "Red Text", "<span class=\"fg-red\">Red Text</span>"},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := convertEscSequencesToSpans(tt.input)
				assert.Equal(t, tt.expected, result)
			})
		}
	})
}
