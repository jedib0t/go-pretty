package table

import (
	"testing"

	"github.com/jedib0t/go-pretty/v6/text"
)

// TestTable_Render_AutoMerge_Wrap covers issue #261: when a vertically merged
// (ColumnConfig.AutoMerge) column is wrapped by WidthMax, the rows merged into
// it should stack their differing columns into the merged cell's wrapped height
// instead of leaving blank lines below it.
func TestTable_Render_AutoMerge_Wrap(t *testing.T) {
	newMergedTable := func(widthMax int) Writer {
		tw := NewWriter()
		tw.SetColumnConfigs([]ColumnConfig{
			{Name: "C1", AutoMerge: true, WidthMax: widthMax, WidthMaxEnforcer: text.WrapSoft},
			{Name: "C2"},
		})
		tw.AppendHeader(Row{"C1", "C2"})
		return tw
	}

	t.Run("issue 261 example", func(t *testing.T) {
		tw := newMergedTable(10)
		tw.AppendRow(Row{"very long row value that wraps", "x"})
		tw.AppendRow(Row{"very long row value that wraps", "y"})

		compareOutput(t, tw.Render(), `
+------------+----+
| C1         | C2 |
+------------+----+
| very long  | x  |
| row value  | y  |
| that wraps |    |
+------------+----+`)
	})

	t.Run("group fully fills wrapped height", func(t *testing.T) {
		tw := newMergedTable(12)
		tw.AppendRow(Row{"very long row value that wraps", "x"})
		tw.AppendRow(Row{"very long row value that wraps", "y"})
		tw.AppendRow(Row{"very long row value that wraps", "z"})

		compareOutput(t, tw.Render(), `
+--------------+----+
| C1           | C2 |
+--------------+----+
| very long    | x  |
| row value    | y  |
| that wraps   | z  |
+--------------+----+`)
	})

	t.Run("more members than wrapped lines", func(t *testing.T) {
		tw := newMergedTable(10)
		tw.AppendRow(Row{"very long row value that wraps", "a"})
		tw.AppendRow(Row{"very long row value that wraps", "b"})
		tw.AppendRow(Row{"very long row value that wraps", "c"})
		tw.AppendRow(Row{"very long row value that wraps", "d"})

		compareOutput(t, tw.Render(), `
+------------+----+
| C1         | C2 |
+------------+----+
| very long  | a  |
| row value  | b  |
| that wraps | c  |
|            | d  |
+------------+----+`)
	})

	t.Run("three columns", func(t *testing.T) {
		tw := NewWriter()
		tw.SetColumnConfigs([]ColumnConfig{
			{Name: "C1", AutoMerge: true, WidthMax: 10, WidthMaxEnforcer: text.WrapSoft},
			{Name: "C2"},
			{Name: "C3"},
		})
		tw.AppendHeader(Row{"C1", "C2", "C3"})
		tw.AppendRow(Row{"very long row value that wraps", "x", "A"})
		tw.AppendRow(Row{"very long row value that wraps", "y", "B"})

		compareOutput(t, tw.Render(), `
+------------+----+----+
| C1         | C2 | C3 |
+------------+----+----+
| very long  | x  | A  |
| row value  | y  | B  |
| that wraps |    |    |
+------------+----+----+`)
	})

	t.Run("different values start a new group", func(t *testing.T) {
		tw := newMergedTable(10)
		tw.AppendRow(Row{"very long row value that wraps", "x"})
		tw.AppendRow(Row{"very long row value that wraps", "y"})
		tw.AppendRow(Row{"different value that also wraps", "z"})

		compareOutput(t, tw.Render(), `
+------------+----+
| C1         | C2 |
+------------+----+
| very long  | x  |
| row value  | y  |
| that wraps |    |
| different  | z  |
| value that |    |
| also wraps |    |
+------------+----+`)
	})

	t.Run("two merged columns with different boundaries", func(t *testing.T) {
		tw := NewWriter()
		tw.SetColumnConfigs([]ColumnConfig{
			{Name: "C1", AutoMerge: true, WidthMax: 8, WidthMaxEnforcer: text.WrapSoft},
			{Name: "C2", AutoMerge: true},
			{Name: "C3"},
		})
		tw.AppendHeader(Row{"C1", "C2", "C3"})
		tw.AppendRow(Row{"alpha beta gamma", "K1", "v1"})
		tw.AppendRow(Row{"alpha beta gamma", "K1", "v2"})
		tw.AppendRow(Row{"alpha beta gamma", "K2", "v3"})

		// C1 spans all three rows, so the block is its wrapped height; C2 does
		// its own vertical merge within that block (K1 over the first two rows,
		// K2 on the third).
		compareOutput(t, tw.Render(), `
+----------+----+----+
| C1       | C2 | C3 |
+----------+----+----+
| alpha    | K1 | v1 |
| beta     |    | v2 |
| gamma    | K2 | v3 |
+----------+----+----+`)
	})

	t.Run("inner merged column re-merges within the block", func(t *testing.T) {
		tw := NewWriter()
		tw.SetColumnConfigs([]ColumnConfig{
			{Name: "C1", AutoMerge: true, WidthMax: 8, WidthMaxEnforcer: text.WrapSoft},
			{Name: "C2", AutoMerge: true},
			{Name: "C3"},
		})
		tw.AppendHeader(Row{"C1", "C2", "C3"})
		tw.AppendRow(Row{"alpha beta gamma", "K1", "v1"})
		tw.AppendRow(Row{"alpha beta gamma", "K2", "v2"})
		tw.AppendRow(Row{"alpha beta gamma", "K2", "v3"})

		compareOutput(t, tw.Render(), `
+----------+----+----+
| C1       | C2 | C3 |
+----------+----+----+
| alpha    | K1 | v1 |
| beta     | K2 | v2 |
| gamma    |    | v3 |
+----------+----+----+`)
	})

	t.Run("merged cell that wraps then changes is not stacked", func(t *testing.T) {
		tw := NewWriter()
		tw.SetColumnConfigs([]ColumnConfig{
			{Name: "C1", AutoMerge: true, WidthMax: 4, WidthMaxEnforcer: text.WrapSoft},
			{Name: "C2"},
		})
		tw.AppendHeader(Row{"C1", "C2"})
		tw.AppendRow(Row{"aa bb", "x"})
		tw.AppendRow(Row{"z", "y"})

		// "aa bb" wraps to two lines, so the next row cannot slot into it
		// without misaligning; each row keeps its own line.
		compareOutput(t, tw.Render(), `
+------+----+
| C1   | C2 |
+------+----+
| aa   | x  |
| bb   |    |
| z    | y  |
+------+----+`)
	})

	t.Run("short values are unaffected", func(t *testing.T) {
		tw := newMergedTable(10)
		tw.AppendRow(Row{"short", "x"})
		tw.AppendRow(Row{"short", "y"})
		tw.AppendRow(Row{"short", "z"})

		compareOutput(t, tw.Render(), `
+-------+----+
| C1    | C2 |
+-------+----+
| short | x  |
|       | y  |
|       | z  |
+-------+----+`)
	})

	t.Run("fallback when a non-merged column also wraps", func(t *testing.T) {
		tw := NewWriter()
		tw.SetColumnConfigs([]ColumnConfig{
			{Name: "C1", AutoMerge: true, WidthMax: 10, WidthMaxEnforcer: text.WrapSoft},
			{Name: "C2", WidthMax: 8, WidthMaxEnforcer: text.WrapSoft},
		})
		tw.AppendHeader(Row{"C1", "C2"})
		tw.AppendRow(Row{"very long row value that wraps", "short value"})
		tw.AppendRow(Row{"very long row value that wraps", "another short value"})

		// the second column wraps too, so the rows cannot be stacked without
		// misaligning columns; fall back to trimming the trailing blank lines.
		compareOutput(t, tw.Render(), `
+------------+----------+
| C1         | C2       |
+------------+----------+
| very long  | short    |
| row value  | value    |
| that wraps |          |
|            | another  |
|            | short    |
|            | value    |
+------------+----------+`)
	})

	t.Run("differing row config breaks the group", func(t *testing.T) {
		tw := newMergedTable(10)
		tw.AppendRow(Row{"very long row value that wraps", "x"})
		tw.AppendRow(Row{"very long row value that wraps", "y"}, RowConfig{AutoMerge: true})

		// the second row carries a different row config, so it is not stacked
		// into the first; only the trailing blank lines are trimmed.
		compareOutput(t, tw.Render(), `
+------------+----+
| C1         | C2 |
+------------+----+
| very long  | x  |
| row value  |    |
| that wraps |    |
|            | y  |
+------------+----+`)
	})

	t.Run("later non-stackable row breaks the group partway", func(t *testing.T) {
		tw := NewWriter()
		tw.SetColumnConfigs([]ColumnConfig{
			{Name: "C1", AutoMerge: true, WidthMax: 10, WidthMaxEnforcer: text.WrapSoft},
			{Name: "C2", WidthMax: 8, WidthMaxEnforcer: text.WrapSoft},
		})
		tw.AppendHeader(Row{"C1", "C2"})
		tw.AppendRow(Row{"very long row value that wraps", "x"})
		tw.AppendRow(Row{"very long row value that wraps", "y"})
		tw.AppendRow(Row{"very long row value that wraps", "multi word value"})

		// the first two rows stack; the third wraps in C2 and cannot join, so
		// it renders on its own below the merged block.
		compareOutput(t, tw.Render(), `
+------------+----------+
| C1         | C2       |
+------------+----------+
| very long  | x        |
| row value  | y        |
| that wraps |          |
|            | multi    |
|            | word     |
|            | value    |
+------------+----------+`)
	})

	t.Run("gated off when rows are separated", func(t *testing.T) {
		tw := newMergedTable(10)
		tw.SetStyle(StyleLight)
		tw.Style().Options.SeparateRows = true
		tw.AppendRow(Row{"very long row value that wraps", "x"})
		tw.AppendRow(Row{"very long row value that wraps", "y"})

		// with separators the rows stay visually distinct; only the trailing
		// blank lines are trimmed.
		compareOutput(t, tw.Render(), `
┌────────────┬────┐
│ C1         │ C2 │
├────────────┼────┤
│ very long  │ x  │
│ row value  │    │
│ that wraps │    │
│            ├────┤
│            │ y  │
└────────────┴────┘`)
	})
}
