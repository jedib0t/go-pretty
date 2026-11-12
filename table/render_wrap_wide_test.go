package table

import (
	"strings"
	"testing"

	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/stretchr/testify/assert"
)

// TestTable_Render_WidthMax_GraphemeClusters ensures that a grapheme cluster
// is not split between two lines of a cell when it fits in one.
func TestTable_Render_WidthMax_GraphemeClusters(t *testing.T) {
	thumbsUp := "\U0001F44D\U0001F3FD" // 4 columns
	for _, enforcer := range []struct {
		name string
		fn   WidthEnforcer
	}{
		{"default", nil},
		{"WrapHard", text.WrapHard},
		{"WrapSoft", text.WrapSoft},
		{"WrapText", text.WrapText},
	} {
		t.Run(enforcer.name, func(t *testing.T) {
			tw := NewWriter()
			tw.AppendRow(Row{"ab" + thumbsUp + "cd", "x"})
			tw.SetColumnConfigs([]ColumnConfig{{Number: 1, WidthMax: 4, WidthMaxEnforcer: enforcer.fn}})

			out := tw.Render()
			compareOutput(t, out, `
+------+---+
| ab   | x |
| `+thumbsUp+` |   |
| cd   |   |
+------+---+`)
			for _, line := range strings.Split(out, "\n") {
				assert.Equal(t, 12, text.StringWidthWithoutEscSequences(line), "line=%q", line)
			}
		})
	}
}

// TestTable_Render_WidthMax_WideRunes ensures that wrapping wide (CJK) content
// to a WidthMax that is not a multiple of the rune width never produces rows
// wider than the table borders.
func TestTable_Render_WidthMax_WideRunes(t *testing.T) {
	for _, enforcer := range []struct {
		name string
		fn   WidthEnforcer
	}{
		{"default", nil},
		{"WrapHard", text.WrapHard},
		{"WrapSoft", text.WrapSoft},
		{"WrapText", text.WrapText},
	} {
		t.Run(enforcer.name, func(t *testing.T) {
			tw := NewWriter()
			tw.AppendRow(Row{"生命生命生命", "x"})
			tw.SetColumnConfigs([]ColumnConfig{{Number: 1, WidthMax: 3, WidthMaxEnforcer: enforcer.fn}})

			out := tw.Render()
			compareOutput(t, out, `
+-----+---+
| 生  | x |
| 命  |   |
| 生  |   |
| 命  |   |
| 生  |   |
| 命  |   |
+-----+---+`)
			for _, line := range strings.Split(out, "\n") {
				assert.Equal(t, 11, text.StringWidthWithoutEscSequences(line), "line=%q", line)
			}
		})
	}
}

// TestTable_Render_WidthMax_MultiLine ensures that a column is wide enough for
// content with many lines, after the enforcer has had a go at it: WrapHard and
// WrapSoft join the lines before wrapping, and can end up with longer lines
// than the content had to begin with.
func TestTable_Render_WidthMax_MultiLine(t *testing.T) {
	for _, enforcer := range []struct {
		name     string
		fn       WidthEnforcer
		expected string
	}{
		{"default", nil, `
+----------+---+
| line one | x |
| line two |   |
+----------+---+`},
		{"WrapHard", text.WrapHard, `
+------------+---+
| line one l | x |
| ine two    |   |
+------------+---+`},
		{"WrapSoft", text.WrapSoft, `
+------------+---+
| line one   | x |
| line two   |   |
+------------+---+`},
		{"WrapText", text.WrapText, `
+----------+---+
| line one | x |
| line two |   |
+----------+---+`},
	} {
		t.Run(enforcer.name, func(t *testing.T) {
			tw := NewWriter()
			tw.AppendRow(Row{"line one\nline two", "x"})
			tw.SetColumnConfigs([]ColumnConfig{{Number: 1, WidthMax: 10, WidthMaxEnforcer: enforcer.fn}})

			compareOutput(t, tw.Render(), enforcer.expected)
		})
	}
}

// TestTable_Render_WidthMax_WideRunes_Hyperlink ensures the same for wide
// content that is the label of a hyperlink.
func TestTable_Render_WidthMax_WideRunes_Hyperlink(t *testing.T) {
	for _, enforcer := range []struct {
		name string
		fn   WidthEnforcer
	}{
		{"default", nil},
		{"WrapHard", text.WrapHard},
		{"WrapSoft", text.WrapSoft},
		{"WrapText", text.WrapText},
	} {
		t.Run(enforcer.name, func(t *testing.T) {
			tw := NewWriter()
			tw.AppendRow(Row{text.Hyperlink("https://example.com/x", "生命生命生命"), "x"})
			tw.SetColumnConfigs([]ColumnConfig{{Number: 1, WidthMax: 3, WidthMaxEnforcer: enforcer.fn}})

			out := tw.Render()
			compareOutput(t, text.StripEscape(out), `
+-----+---+
| 生  | x |
| 命  |   |
| 生  |   |
| 命  |   |
| 生  |   |
| 命  |   |
+-----+---+`)
			assert.Equal(t, 6, strings.Count(out, "https://example.com/x"))
			for _, line := range strings.Split(out, "\n") {
				assert.Equal(t, 11, text.StringWidthWithoutEscSequences(line), "line=%q", line)
			}
		})
	}
}

// TestTable_Render_WidthMax_ControlSequences ensures that content with
// sequences that move the cursor or erase text is wrapped to WidthMax too.
func TestTable_Render_WidthMax_ControlSequences(t *testing.T) {
	for _, enforcer := range []struct {
		name string
		fn   WidthEnforcer
	}{
		{"default", nil},
		{"WrapHard", text.WrapHard},
		{"WrapSoft", text.WrapSoft},
		{"WrapText", text.WrapText},
	} {
		t.Run(enforcer.name, func(t *testing.T) {
			tw := NewWriter()
			tw.AppendRow(Row{text.EraseLine.Sprint() + "abcdefgh", "x"})
			tw.SetColumnConfigs([]ColumnConfig{{Number: 1, WidthMax: 4, WidthMaxEnforcer: enforcer.fn}})

			compareOutput(t, text.StripEscape(tw.Render()), `
+------+---+
| abcd | x |
| efgh |   |
+------+---+`)
		})
	}
}
