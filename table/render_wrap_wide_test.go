package table

import (
	"strings"
	"testing"

	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/stretchr/testify/assert"
)

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
