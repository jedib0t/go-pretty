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
			{"color change without a reset", "\x1b[31mred\x1b[32mgreen\x1b[0m", "<span class=\"fg-red\">red</span><span class=\"fg-green\">green</span>"},
			{"background change without a reset", "\x1b[41mred\x1b[42mgreen\x1b[0m", "<span class=\"bg-red\">red</span><span class=\"bg-green\">green</span>"},
			{"256-color change without a reset", "\x1b[38;5;9mred\x1b[32mgreen\x1b[0m", "<span class=\"fg-256-255-0-0\">red</span><span class=\"fg-green\">green</span>"},
			{"reset and 256-color in one sequence", "\x1b[31mred\x1b[0;38;5;10mgreen\x1b[0m", "<span class=\"fg-red\">red</span><span class=\"fg-256-0-255-0\">green</span>"},
			{"reset without a parameter", "\x1b[31mred\x1b[m plain", "<span class=\"fg-red\">red</span> plain"},
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

	t.Run("24-bit colors", func(t *testing.T) {
		tests := []struct {
			name     string
			input    string
			expected string
		}{
			{"foreground", "\x1b[38;2;255;0;10mText\x1b[0m", "<span style=\"color: #ff000a;\">Text</span>"},
			{"background", "\x1b[48;2;0;128;255mText\x1b[0m", "<span style=\"background-color: #0080ff;\">Text</span>"},
			{"foreground and background", "\x1b[38;2;255;0;0;48;2;0;255;0mText\x1b[0m", "<span style=\"color: #ff0000; background-color: #00ff00;\">Text</span>"},
			{"with attributes", "\x1b[1;4;38;2;1;0;3mText\x1b[0m", "<span class=\"bold underline\" style=\"color: #010003;\">Text</span>"},
			{"with a standard background", "\x1b[44;38;2;255;0;0mText\x1b[0m", "<span class=\"bg-blue\" style=\"color: #ff0000;\">Text</span>"},
			{"color changes", "\x1b[38;2;255;0;0ma\x1b[38;2;0;255;0mb\x1b[0mc", "<span style=\"color: #ff0000;\">a</span><span style=\"color: #00ff00;\">b</span>c"},
			{"same color set twice", "\x1b[38;2;255;0;0ma\x1b[38;2;255;0;0mb\x1b[0m", "<span style=\"color: #ff0000;\">ab</span>"},
			{"replaces a standard color", "\x1b[1;31ma\x1b[38;2;1;2;3mb\x1b[0m", "<span class=\"bold fg-red\">a</span><span class=\"bold\" style=\"color: #010203;\">b</span>"},
			{"replaced by a standard color", "\x1b[38;2;1;2;3ma\x1b[31mb\x1b[0m", "<span style=\"color: #010203;\">a</span><span class=\"fg-red\">b</span>"},
			{"foreground reset", "\x1b[38;2;1;2;3;48;2;4;5;6ma\x1b[39mb\x1b[0m", "<span style=\"color: #010203; background-color: #040506;\">a</span><span style=\"background-color: #040506;\">b</span>"},
			{"open span at end", "\x1b[38;2;1;2;3mText", "<span style=\"color: #010203;\">Text</span>"},
			{"malformed color is ignored", "\x1b[38;2;256;0;0mText\x1b[0m", "Text"},
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
			{"same classes from different codes", "\x1b[31;21ma\x1b[0;31;53mb\x1b[0m", "<span class=\"fg-red\">ab</span>"},
			{"hyperlink is not concealed text", "\x1b]8;;http://example.com\x1b\\link\x1b]8;;\x1b\\ after", "link after"},
			{"window title that ends in a BEL", "\x1b[31mred \x1b]0;my title\aafter\x1b[0m", "<span class=\"fg-red\">red after</span>"},
			{"window title is not a reset", "\x1b[31mred \x1b]0;my title\x1b\\after\x1b[0m", "<span class=\"fg-red\">red after</span>"},
			{"working directory is not reverse video", "\x1b]7;file://host/home\x1b\\text", "text"},
			{"malformed sequence leaves the text alone", "\x1b[31mred\x1b\nmore\x1b[3世界\x1b[0m", "<span class=\"fg-red\">red\nmore世界</span>"},
			{"hyperlink within a color", "\x1b[31mred \x1b]8;;http://example.com/0\x07link\x1b]8;;\x07 after\x1b[0m", "<span class=\"fg-red\">red link after</span>"},
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
