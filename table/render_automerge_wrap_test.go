package table

import (
	"testing"

	"github.com/jedib0t/go-pretty/v6/text"
)

func TestTable_Render_AutoMerge_WrapWithVerticalMerge(t *testing.T) {
	t.Run("basic wrap with vertical merge", func(t *testing.T) {
		tw := NewWriter()

		tw.SetColumnConfigs([]ColumnConfig{
			{
				Name:             "C1",
				AutoMerge:        true,
				WidthMax:         10,
				WidthMaxEnforcer: text.WrapSoft,
			},
			{
				Name: "C2",
			},
		})

		tw.AppendHeader(Row{"C1", "C2"})
		tw.AppendRow(Row{"very long row value that wraps", "x"})
		tw.AppendRow(Row{"very long row value that wraps", "y"})

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

	t.Run("multiple rows merge", func(t *testing.T) {
		tw := NewWriter()

		tw.SetColumnConfigs([]ColumnConfig{
			{
				Name:             "C1",
				AutoMerge:        true,
				WidthMax:         12,
				WidthMaxEnforcer: text.WrapSoft,
			},
			{
				Name: "C2",
			},
		})

		tw.AppendHeader(Row{"C1", "C2"})
		tw.AppendRow(Row{"very long row value that wraps", "x"})
		tw.AppendRow(Row{"very long row value that wraps", "y"})
		tw.AppendRow(Row{"very long row value that wraps", "z"})

		compareOutput(t, tw.Render(), `
+--------------+----+
| C1           | C2 |
+--------------+----+
| very long    | x  |
| row value    |    |
| that wraps   |    |
|              | y  |
|              | z  |
+--------------+----+`)
	})

	t.Run("three columns", func(t *testing.T) {
		tw := NewWriter()

		tw.SetColumnConfigs([]ColumnConfig{
			{
				Name:             "C1",
				AutoMerge:        true,
				WidthMax:         10,
				WidthMaxEnforcer: text.WrapSoft,
			},
			{
				Name: "C2",
			},
			{
				Name: "C3",
			},
		})

		tw.AppendHeader(Row{"C1", "C2", "C3"})
		tw.AppendRow(Row{"very long row value that wraps", "x", "A"})
		tw.AppendRow(Row{"very long row value that wraps", "y", "B"})

		compareOutput(t, tw.Render(), `
+------------+----+----+
| C1         | C2 | C3 |
+------------+----+----+
| very long  | x  | A  |
| row value  |    |    |
| that wraps |    |    |
|            | y  | B  |
+------------+----+----+`)
	})

	t.Run("non-merged column wraps", func(t *testing.T) {
		tw := NewWriter()

		tw.SetColumnConfigs([]ColumnConfig{
			{
				Name:             "C1",
				AutoMerge:        true,
				WidthMax:         10,
				WidthMaxEnforcer: text.WrapSoft,
			},
			{
				Name:             "C2",
				WidthMax:         8,
				WidthMaxEnforcer: text.WrapSoft,
			},
		})

		tw.AppendHeader(Row{"C1", "C2"})
		tw.AppendRow(Row{"very long row value that wraps", "short value"})
		tw.AppendRow(Row{"very long row value that wraps", "another short value"})

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

	t.Run("different values break merge", func(t *testing.T) {
		tw := NewWriter()

		tw.SetColumnConfigs([]ColumnConfig{
			{
				Name:             "C1",
				AutoMerge:        true,
				WidthMax:         10,
				WidthMaxEnforcer: text.WrapSoft,
			},
			{
				Name: "C2",
			},
		})

		tw.AppendHeader(Row{"C1", "C2"})
		tw.AppendRow(Row{"very long row value that wraps", "x"})
		tw.AppendRow(Row{"different value that also wraps", "y"})

		compareOutput(t, tw.Render(), `
+------------+----+
| C1         | C2 |
+------------+----+
| very long  | x  |
| row value  |    |
| that wraps |    |
| different  | y  |
| value that |    |
| also wraps |    |
+------------+----+`)
	})

	t.Run("previous row has fewer lines", func(t *testing.T) {
		tw := NewWriter()

		tw.SetColumnConfigs([]ColumnConfig{
			{
				Name:             "C1",
				AutoMerge:        true,
				WidthMax:         10,
				WidthMaxEnforcer: text.WrapSoft,
			},
			{
				Name: "C2",
			},
		})

		tw.AppendHeader(Row{"C1", "C2"})
		tw.AppendRow(Row{"short", "x"})
		tw.AppendRow(Row{"very long row value that wraps", "y"})

		compareOutput(t, tw.Render(), `
+------------+----+
| C1         | C2 |
+------------+----+
| short      | x  |
| very long  | y  |
| row value  |    |
| that wraps |    |
+------------+----+`)
	})

	t.Run("current row has more lines than previous", func(t *testing.T) {
		tw := NewWriter()

		tw.SetColumnConfigs([]ColumnConfig{
			{
				Name:             "C1",
				AutoMerge:        true,
				WidthMax:         10,
				WidthMaxEnforcer: text.WrapSoft,
			},
			{
				Name: "C2",
			},
		})

		tw.AppendHeader(Row{"C1", "C2"})
		tw.AppendRow(Row{"very long row value that wraps", "x"})
		tw.AppendRow(Row{"very long row value that wraps and extends further", "y"})

		compareOutput(t, tw.Render(), `
+------------+----+
| C1         | C2 |
+------------+----+
| very long  | x  |
| row value  |    |
| that wraps |    |
| very long  | y  |
| row value  |    |
| that wraps |    |
| and        |    |
| extends    |    |
| further    |    |
+------------+----+`)
	})

	t.Run("no width max uses column length", func(t *testing.T) {
		tw := NewWriter()

		tw.SetColumnConfigs([]ColumnConfig{
			{
				Name:             "C1",
				AutoMerge:        true,
				WidthMaxEnforcer: text.WrapSoft,
			},
			{
				Name: "C2",
			},
		})

		tw.AppendHeader(Row{"C1", "C2"})
		tw.AppendRow(Row{"very long row value that wraps", "x"})
		tw.AppendRow(Row{"very long row value that wraps", "y"})

		compareOutput(t, tw.Render(), `
+--------------------------------+----+
| C1                             | C2 |
+--------------------------------+----+
| very long row value that wraps | x  |
|                                | y  |
+--------------------------------+----+`)
	})

	t.Run("non-merged column determines max height", func(t *testing.T) {
		tw := NewWriter()

		tw.SetColumnConfigs([]ColumnConfig{
			{
				Name:             "C1",
				AutoMerge:        true,
				WidthMax:         10,
				WidthMaxEnforcer: text.WrapSoft,
			},
			{
				Name:             "C2",
				WidthMax:         5,
				WidthMaxEnforcer: text.WrapSoft,
			},
		})

		tw.AppendHeader(Row{"C1", "C2"})
		tw.AppendRow(Row{"short", "very long value that wraps"})
		tw.AppendRow(Row{"short", "very long value that wraps"})

		compareOutput(t, tw.Render(), `
+-------+-------+
| C1    | C2    |
+-------+-------+
| short | very  |
|       | long  |
|       | value |
|       | that  |
|       | wraps |
|       | very  |
|       | long  |
|       | value |
|       | that  |
|       | wraps |
+-------+-------+`)
	})

	t.Run("separator row style", func(t *testing.T) {
		tw := NewWriter()
		tw.SetStyle(StyleLight)
		tw.Style().Options.SeparateRows = true

		tw.SetColumnConfigs([]ColumnConfig{
			{
				Name:             "C1",
				AutoMerge:        true,
				WidthMax:         10,
				WidthMaxEnforcer: text.WrapSoft,
			},
			{
				Name: "C2",
			},
		})

		tw.AppendHeader(Row{"C1", "C2"})
		tw.AppendRow(Row{"very long row value that wraps", "x"})
		tw.AppendRow(Row{"very long row value that wraps", "y"})

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

	t.Run("short values merge across multiple rows", func(t *testing.T) {
		tw := NewWriter()

		tw.SetColumnConfigs([]ColumnConfig{
			{
				Name:             "C1",
				AutoMerge:        true,
				WidthMax:         10,
				WidthMaxEnforcer: text.WrapSoft,
			},
			{
				Name: "C2",
			},
		})

		tw.AppendHeader(Row{"C1", "C2"})
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

	t.Run("non-merged column wraps more than merged", func(t *testing.T) {
		tw := NewWriter()

		tw.SetColumnConfigs([]ColumnConfig{
			{
				Name:             "C1",
				AutoMerge:        true,
				WidthMax:         10,
				WidthMaxEnforcer: text.WrapSoft,
			},
			{
				Name:             "C2",
				WidthMax:         15,
				WidthMaxEnforcer: text.WrapSoft,
			},
		})

		tw.AppendHeader(Row{"C1", "C2"})
		tw.AppendRow(Row{"short", "very long value that wraps"})
		tw.AppendRow(Row{"short", "very long value that wraps"})

		compareOutput(t, tw.Render(), `
+-------+-----------------+
| C1    | C2              |
+-------+-----------------+
| short | very long value |
|       | that wraps      |
|       | very long value |
|       | that wraps      |
+-------+-----------------+`)
	})
}
