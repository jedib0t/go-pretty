package table

import (
	"testing"

	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/stretchr/testify/assert"
)

func newClippedRowsTable(rows []Row) Writer {
	tw := NewWriter()
	tw.AppendHeader(Row{"ID", "Notes"})
	tw.AppendRows(rows)
	tw.SetColumnConfigs([]ColumnConfig{{Number: 2, WidthMax: 5, WidthMaxEnforcer: text.WrapSoft}})
	tw.Style().Size.WidthMax = 8
	return tw
}

func TestTable_Render_SkipEmptyRowsWhenClipped(t *testing.T) {
	tw := newClippedRowsTable([]Row{{"1", "alpha beta gamma"}, {"2", "delta"}})
	const original = `+----+ ~
| ID | ~
+----+ ~
| 1  | ~
|    | ~
|    | ~
| 2  | ~
+----+ ~`
	compareOutput(t, tw.Render(), original)

	tw.Style().Options.SkipEmptyRowsWhenClipped = true
	compareOutput(t, tw.Render(), `+----+ ~
| ID | ~
+----+ ~
| 1  | ~
| 2  | ~
+----+ ~`)

	tw.Style().Options.SkipEmptyRowsWhenClipped = false
	compareOutput(t, tw.Render(), original)
}

func TestTable_Render_SkipEmptyRowsWhenClipped_PreserveRows(t *testing.T) {
	tests := []struct {
		name     string
		rows     []Row
		merge    bool
		expected string
	}{
		{
			name:  "vertical merge",
			rows:  []Row{{"g1", "alpha beta"}, {"g1", "gamma"}, {"g1", "delta"}},
			merge: true,
			expected: `+----+ ~
| ID | ~
+----+ ~
| g1 | ~
|    | ~
|    | ~
+----+ ~`,
		},
		{
			name: "literal structural characters",
			rows: []Row{{"|", "alpha beta"}, {"~", "gamma delta"}},
			expected: `+----+ ~
| ID | ~
+----+ ~
| |  | ~
| ~  | ~
+----+ ~`,
		},
		{
			name: "structural characters in a continuation",
			rows: []Row{{"|\n~", "alpha beta gamma"}},
			expected: `+----+ ~
| ID | ~
+----+ ~
| |  | ~
| ~  | ~
+----+ ~`,
		},
		{
			name: "empty logical row",
			rows: []Row{{"1", "alpha"}, {"", ""}, {"2", "beta"}},
			expected: `+----+ ~
| ID | ~
+----+ ~
| 1  | ~
|    | ~
| 2  | ~
+----+ ~`,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tw := newClippedRowsTable(test.rows)
			if test.merge {
				tw.SetColumnConfigs([]ColumnConfig{
					{Number: 1, AutoMerge: true},
					{Number: 2, WidthMax: 5, WidthMaxEnforcer: text.WrapSoft},
				})
			}
			tw.Style().Options.SkipEmptyRowsWhenClipped = true
			compareOutput(t, tw.Render(), test.expected)
		})
	}
}

func TestTable_Render_SkipEmptyRowsWhenClipped_Pager(t *testing.T) {
	tw := newClippedRowsTable([]Row{{"1", "alpha beta gamma"}, {"2", "delta theta"}})
	tw.Style().Options.SkipEmptyRowsWhenClipped = true
	pg := tw.Pager(PageSize(2))
	compareOutput(t, pg.Render(), `+----+ ~
| ID | ~
+----+ ~
| 1  | ~
| 2  | ~
+----+ ~`)
	assert.Equal(t, pg.Render(), pg.Next())
	assert.Equal(t, 1, pg.Location())

	tw.AppendRow(Row{"3", "extra words"})
	pg = tw.Pager(PageSize(2))
	assert.Equal(t, `
+----+ ~
| ID | ~
+----+ ~
| 3  | ~
+----+ ~`, pg.GoTo(2))
	assert.Equal(t, pg.Render(), pg.Next())
	assert.Equal(t, 2, pg.Location())
}

func TestTable_Render_SkipEmptyRowsWhenClipped_UnchangedOutput(t *testing.T) {
	for _, width := range []int{0, 100} {
		for _, pageSize := range []int{0, 2} {
			for _, separateRows := range []bool{false, true} {
				tw := newClippedRowsTable([]Row{{"1", "alpha beta"}, {"2", "gamma"}, {"3", "delta"}})
				tw.AppendFooter(Row{"F", "footer"})
				tw.Style().Size.WidthMax = width
				tw.Style().Options.SeparateRows = separateRows
				tw.SetPageSize(pageSize)
				original := tw.Render()
				tw.Style().Options.SkipEmptyRowsWhenClipped = true
				assert.Equal(t, original, tw.Render(), "width=%d pageSize=%d separateRows=%t",
					width, pageSize, separateRows)
			}
		}
	}
}

func TestTable_Render_SkipEmptyRowsWhenClipped_PreserveEmptyContinuation(t *testing.T) {
	tw := newClippedRowsTable([]Row{{"1\n\n2", "alpha\n\nbeta"}})
	tw.SetColumnConfigs([]ColumnConfig{{
		Number: 2, WidthMax: 5,
		WidthMaxEnforcer: func(value string, _ int) string { return value },
	}})
	original := tw.Render()
	tw.Style().Options.SkipEmptyRowsWhenClipped = true
	compareOutput(t, tw.Render(), `+----+ ~
| ID | ~
+----+ ~
| 1  | ~
|    | ~
| 2  | ~
+----+ ~`)
	assert.Equal(t, original, tw.Render())
}

func TestTable_Render_SkipEmptyRowsWhenClipped_MergedBlock(t *testing.T) {
	tw := newClippedRowsTable([]Row{
		{"g1", "alpha beta gamma delta"},
		{"g1", "alpha beta gamma delta"},
		{"g1", "alpha beta gamma delta"},
	})
	tw.SetColumnConfigs([]ColumnConfig{
		{Number: 1, AutoMerge: true},
		{Number: 2, AutoMerge: true, WidthMax: 5, WidthMaxEnforcer: text.WrapSoft},
	})
	tw.Style().Options.SkipEmptyRowsWhenClipped = true
	compareOutput(t, tw.Render(), `+----+ ~
| ID | ~
+----+ ~
| g1 | ~
|    | ~
|    | ~
+----+ ~`)
}

func TestTable_Render_SkipEmptyRowsWhenClipped_FiveColumns(t *testing.T) {
	tw := NewWriter()
	tw.AppendHeader(Row{"G", "A", "B", "C", "D"})
	tw.AppendRows([]Row{
		{"g", "1", "ab\ncd", "aaa bbb ccc", "ddd eee fff"},
		{"g", "2", "x", "ggg hhh", "iii jjj"},
	})
	tw.AppendFooter(Row{"", "", "F", "", ""})
	tw.SetColumnConfigs([]ColumnConfig{
		{Number: 1, AutoMerge: true},
		{Number: 4, WidthMax: 3, WidthMaxEnforcer: text.WrapSoft},
		{Number: 5, WidthMax: 3, WidthMaxEnforcer: text.WrapSoft},
	})
	tw.Style().Size.WidthMax = 16
	tw.Style().Options.SkipEmptyRowsWhenClipped = true
	compareOutput(t, tw.Render(), `+---+---+----+ ~
| G | A | B  | ~
+---+---+----+ ~
| g | 1 | ab | ~
|   |   | cd | ~
|   | 2 | x  | ~
+---+---+----+ ~
|   |   | F  | ~
+---+---+----+ ~`)
}

func TestTable_Render_SkipEmptyRowsWhenClipped_HeaderAndFooter(t *testing.T) {
	tw := newClippedRowsTable([]Row{{"1", "alpha"}})
	tw.ResetHeaders()
	tw.AppendHeader(Row{"ID", "header words"})
	tw.AppendFooter(Row{"F", "footer words"})
	original := tw.Render()
	tw.Style().Options.SkipEmptyRowsWhenClipped = true
	assert.Equal(t, original, tw.Render())
}

func TestTable_Render_SkipEmptyRowsWhenClipped_Colors(t *testing.T) {
	tw := newClippedRowsTable([]Row{{
		"\x1b[31m|\x1b[0m\n\x1b[32m~\x1b[0m",
		"\x1b[34malpha beta gamma\x1b[0m",
	}})
	tw.Style().Options.SkipEmptyRowsWhenClipped = true
	out := tw.Render()
	assert.Contains(t, out, "\x1b[31m")
	compareOutput(t, text.StripEscape(out), `+----+ ~
| ID | ~
+----+ ~
| |  | ~
| ~  | ~
+----+ ~`)
}

func TestTable_Render_SkipEmptyRowsWhenClipped_WideCharacters(t *testing.T) {
	tw := newClippedRowsTable([]Row{{"1", "界界界"}})
	tw.ResetHeaders()
	tw.AppendHeader(Row{"ID", "N"})
	tw.SetColumnConfigs([]ColumnConfig{{Number: 2, WidthMax: 2, WidthMaxEnforcer: text.WrapHard}})
	tw.Style().Size.WidthMax = 10
	tw.Style().Options.SkipEmptyRowsWhenClipped = true
	compareOutput(t, tw.Render(), `+----+-- ~
| ID | N ~
+----+-- ~
| 1  |  ~
+----+-- ~`)
}

func TestTable_Render_SkipEmptyRowsWhenClipped_PagerWithMerge(t *testing.T) {
	tw := newClippedRowsTable([]Row{
		{"g1", "alpha beta"}, {"g1", "gamma delta"}, {"g1", "theta omega"},
	})
	tw.SetColumnConfigs([]ColumnConfig{
		{Number: 1, AutoMerge: true},
		{Number: 2, WidthMax: 5, WidthMaxEnforcer: text.WrapSoft},
	})
	tw.AppendFooter(Row{"F", "total"})
	tw.Style().Options.SkipEmptyRowsWhenClipped = true
	pg := tw.Pager(PageSize(2))
	assert.Equal(t, `
+----+ ~
| ID | ~
+----+ ~
| g1 | ~
+----+ ~
| F  | ~
+----+ ~`, pg.GoTo(2))
	assert.Equal(t, pg.Render(), pg.Next())
	assert.Equal(t, 2, pg.Location())
}
