package table

import (
	"testing"

	"github.com/jedib0t/go-pretty/v6/text"
)

func TestTable_Render_AutoMerge(t *testing.T) {
	rcAutoMerge := RowConfig{AutoMerge: true}

	t.Run("columns only", func(t *testing.T) {
		tw := NewWriter()
		tw.AppendHeader(Row{"Node IP", "Pods", "Namespace", "Container", "RCE\nEXE", "RCE\nRUN"})
		tw.AppendRow(Row{"a.a.a.a", "Pod 1A", "NS 1A", "C 1", "Y", "Y"})
		tw.AppendRow(Row{"a.a.a.a", "Pod 1A", "NS 1A", "C 2", "Y", "N"})
		tw.AppendRow(Row{"a.a.a.a", "Pod 1A", "NS 1B", "C 3", "N", "N"})
		tw.AppendRow(Row{"a.a.a.a", "Pod 1B", "NS 2", "C 4", "N", "N"})
		tw.AppendRow(Row{"a.a.a.a", "Pod 1B", "NS 2", "C 5", "Y", "N"})
		tw.AppendRow(Row{"b.b.b.b", "Pod 2", "NS 3", "C 6", "Y", "Y"})
		tw.AppendRow(Row{"b.b.b.b", "Pod 2", "NS 3", "C 7", "Y", "Y"})
		tw.AppendFooter(Row{"", "", "", 7, 5, 3})
		tw.SetAutoIndex(true)
		tw.SetColumnConfigs([]ColumnConfig{
			{Number: 1, AutoMerge: true},
			{Number: 2, AutoMerge: true},
			{Number: 3, AutoMerge: true},
			{Number: 4, AutoMerge: true},
			{Number: 5, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter},
			{Number: 6, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter},
		})
		tw.SetStyle(StyleLight)
		tw.Style().Options.SeparateRows = true

		compareOutput(t, tw.Render(), `
┌───┬─────────┬────────┬───────────┬───────────┬─────┬─────┐
│   │ NODE IP │ PODS   │ NAMESPACE │ CONTAINER │ RCE │ RCE │
│   │         │        │           │           │ EXE │ RUN │
├───┼─────────┼────────┼───────────┼───────────┼─────┼─────┤
│ 1 │ a.a.a.a │ Pod 1A │ NS 1A     │ C 1       │  Y  │  Y  │
├───┤         │        │           ├───────────┼─────┼─────┤
│ 2 │         │        │           │ C 2       │  Y  │  N  │
├───┤         │        ├───────────┼───────────┼─────┼─────┤
│ 3 │         │        │ NS 1B     │ C 3       │  N  │  N  │
├───┤         ├────────┼───────────┼───────────┼─────┼─────┤
│ 4 │         │ Pod 1B │ NS 2      │ C 4       │  N  │  N  │
├───┤         │        │           ├───────────┼─────┼─────┤
│ 5 │         │        │           │ C 5       │  Y  │  N  │
├───┼─────────┼────────┼───────────┼───────────┼─────┼─────┤
│ 6 │ b.b.b.b │ Pod 2  │ NS 3      │ C 6       │  Y  │  Y  │
├───┤         │        │           ├───────────┼─────┼─────┤
│ 7 │         │        │           │ C 7       │  Y  │  Y  │
├───┼─────────┼────────┼───────────┼───────────┼─────┼─────┤
│   │         │        │           │ 7         │  5  │  3  │
└───┴─────────┴────────┴───────────┴───────────┴─────┴─────┘`)
	})

	t.Run("columns only with hidden columns", func(t *testing.T) {
		tw := NewWriter()
		tw.AppendHeader(Row{"Node IP", "Pods", "Namespace", "Container", "RCE\nEXE", "RCE\nRUN"})
		tw.AppendRow(Row{"a.a.a.a", "Pod 1A", "NS 1A", "C 1", "Y", "Y"})
		tw.AppendRow(Row{"a.a.a.a", "Pod 1A", "NS 1A", "C 2", "Y", "N"})
		tw.AppendRow(Row{"a.a.a.a", "Pod 1A", "NS 1B", "C 3", "N", "N"})
		tw.AppendRow(Row{"a.a.a.a", "Pod 1B", "NS 2", "C 4", "Y", "Y"})
		tw.AppendRow(Row{"a.a.a.a", "Pod 1B", "NS 2", "C 5", "Y", "N"})
		tw.AppendRow(Row{"b.b.b.b", "Pod 2", "NS 3", "C 6", "Y", "Y"})
		tw.AppendRow(Row{"b.b.b.b", "Pod 2", "NS 3", "C 7", "Y", "N"})
		tw.AppendFooter(Row{"", "", "", 7, 5, 3})
		tw.SetColumnConfigs([]ColumnConfig{
			{Number: 1, AutoMerge: true},
			{Number: 2, AutoMerge: true},
			{Number: 3, AutoMerge: true},
			{Number: 4, Hidden: true},
			{Number: 5, Hidden: true, Align: text.AlignCenter},
			{Number: 6, Hidden: true, Align: text.AlignCenter},
		})
		tw.SetStyle(StyleLight)
		tw.Style().Options.SeparateRows = true

		compareOutput(t, tw.Render(), `
┌─────────┬────────┬───────────┐
│ NODE IP │ PODS   │ NAMESPACE │
├─────────┼────────┼───────────┤
│ a.a.a.a │ Pod 1A │ NS 1A     │
│         │        │           │
│         │        │           │
│         │        ├───────────┤
│         │        │ NS 1B     │
│         ├────────┼───────────┤
│         │ Pod 1B │ NS 2      │
│         │        │           │
│         │        │           │
├─────────┼────────┼───────────┤
│ b.b.b.b │ Pod 2  │ NS 3      │
│         │        │           │
│         │        │           │
├─────────┼────────┼───────────┤
│         │        │           │
└─────────┴────────┴───────────┘`)
	})

	t.Run("rows only", func(t *testing.T) {
		tw := NewWriter()
		tw.AppendHeader(Row{"Node IP", "Pods", "Namespace", "Container", "RCE", "RCE"}, rcAutoMerge)
		tw.AppendHeader(Row{"", "", "", "", "EXE", "RUN"})
		tw.AppendRow(Row{"a.a.a.a", "Pod 1A", "NS 1A", "C 1", "Y", "Y"}, rcAutoMerge)
		tw.AppendRow(Row{"a.a.a.a", "Pod 1A", "NS 1A", "C 2", "Y", "N"})
		tw.AppendRow(Row{"a.a.a.a", "Pod 1A", "NS 1B", "C 3", "N", "N"})
		tw.AppendRow(Row{"a.a.a.a", "Pod 1B", "NS 2", "C 4", "N", "N"}, rcAutoMerge)
		tw.AppendRow(Row{"a.a.a.a", "Pod 1B", "NS 2", "C 5", "Y", "N"})
		tw.AppendRow(Row{"b.b.b.b", "Pod 2", "NS 3", "C 6", "Y", "Y"}, rcAutoMerge)
		tw.AppendRow(Row{"b.b.b.b", "Pod 2", "NS 3", "C 7", "Y", "Y"}, RowConfig{AutoMerge: true, AutoMergeAlign: text.AlignRight})
		tw.AppendFooter(Row{"", "", "", 7, 5, 3})
		tw.SetAutoIndex(true)
		tw.SetColumnConfigs([]ColumnConfig{
			{Number: 5, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter},
			{Number: 6, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter},
		})
		tw.SetStyle(StyleLight)
		tw.Style().Options.SeparateRows = true

		compareOutput(t, tw.Render(), `
┌───┬─────────┬────────┬───────────┬───────────┬───────────┐
│   │ NODE IP │ PODS   │ NAMESPACE │ CONTAINER │    RCE    │
│   ├─────────┼────────┼───────────┼───────────┼─────┬─────┤
│   │         │        │           │           │ EXE │ RUN │
├───┼─────────┼────────┼───────────┼───────────┼─────┴─────┤
│ 1 │ a.a.a.a │ Pod 1A │ NS 1A     │ C 1       │     Y     │
├───┼─────────┼────────┼───────────┼───────────┼─────┬─────┤
│ 2 │ a.a.a.a │ Pod 1A │ NS 1A     │ C 2       │  Y  │  N  │
├───┼─────────┼────────┼───────────┼───────────┼─────┼─────┤
│ 3 │ a.a.a.a │ Pod 1A │ NS 1B     │ C 3       │  N  │  N  │
├───┼─────────┼────────┼───────────┼───────────┼─────┴─────┤
│ 4 │ a.a.a.a │ Pod 1B │ NS 2      │ C 4       │     N     │
├───┼─────────┼────────┼───────────┼───────────┼─────┬─────┤
│ 5 │ a.a.a.a │ Pod 1B │ NS 2      │ C 5       │  Y  │  N  │
├───┼─────────┼────────┼───────────┼───────────┼─────┴─────┤
│ 6 │ b.b.b.b │ Pod 2  │ NS 3      │ C 6       │     Y     │
├───┼─────────┼────────┼───────────┼───────────┼───────────┤
│ 7 │ b.b.b.b │ Pod 2  │ NS 3      │ C 7       │         Y │
├───┼─────────┼────────┼───────────┼───────────┼─────┬─────┤
│   │         │        │           │ 7         │  5  │  3  │
└───┴─────────┴────────┴───────────┴───────────┴─────┴─────┘`)
	})

	t.Run("rows and columns", func(t *testing.T) {
		tw := NewWriter()
		tw.AppendHeader(Row{"Node IP", "Pods", "Namespace", "Container", "RCE", "RCE"}, rcAutoMerge)
		tw.AppendHeader(Row{"", "", "", "", "EXE", "RUN"})
		tw.AppendRow(Row{"a.a.a.a", "Pod 1A", "NS 1A", "C 1", "Y", "Y"}, rcAutoMerge)
		tw.AppendRow(Row{"a.a.a.a", "Pod 1A", "NS 1A", "C 2", "Y", "N"}, rcAutoMerge)
		tw.AppendRow(Row{"a.a.a.a", "Pod 1A", "NS 1B", "C 3", "N", "N"}, rcAutoMerge)
		tw.AppendRow(Row{"a.a.a.a", "Pod 1B", "NS 2", "C 4", "N", "N"}, rcAutoMerge)
		tw.AppendRow(Row{"a.a.a.a", "Pod 1B", "NS 2", "C 5", "Y", "N"}, rcAutoMerge)
		tw.AppendRow(Row{"b.b.b.b", "Pod 2", "NS 3", "C 6", "Y", "Y"}, rcAutoMerge)
		tw.AppendRow(Row{"b.b.b.b", "Pod 2", "NS 3", "C 7", "Y", "Y"}, rcAutoMerge)
		tw.AppendFooter(Row{"", "", "", 7, 5, 3})
		tw.SetAutoIndex(true)
		tw.SetColumnConfigs([]ColumnConfig{
			{Number: 1, AutoMerge: true},
			{Number: 2, AutoMerge: true},
			{Number: 3, AutoMerge: true},
			{Number: 4, AutoMerge: true},
			{Number: 5, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter},
			{Number: 6, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter},
		})
		tw.SetStyle(StyleLight)
		tw.Style().Options.SeparateRows = true

		compareOutput(t, tw.Render(), `
┌───┬─────────┬────────┬───────────┬───────────┬───────────┐
│   │ NODE IP │ PODS   │ NAMESPACE │ CONTAINER │    RCE    │
│   ├─────────┼────────┼───────────┼───────────┼─────┬─────┤
│   │         │        │           │           │ EXE │ RUN │
├───┼─────────┼────────┼───────────┼───────────┼─────┴─────┤
│ 1 │ a.a.a.a │ Pod 1A │ NS 1A     │ C 1       │     Y     │
├───┤         │        │           ├───────────┼─────┬─────┤
│ 2 │         │        │           │ C 2       │  Y  │  N  │
├───┤         │        ├───────────┼───────────┼─────┴─────┤
│ 3 │         │        │ NS 1B     │ C 3       │     N     │
├───┤         ├────────┼───────────┼───────────┼───────────┤
│ 4 │         │ Pod 1B │ NS 2      │ C 4       │     N     │
├───┤         │        │           ├───────────┼─────┬─────┤
│ 5 │         │        │           │ C 5       │  Y  │  N  │
├───┼─────────┼────────┼───────────┼───────────┼─────┴─────┤
│ 6 │ b.b.b.b │ Pod 2  │ NS 3      │ C 6       │     Y     │
├───┤         │        │           ├───────────┼───────────┤
│ 7 │         │        │           │ C 7       │     Y     │
├───┼─────────┼────────┼───────────┼───────────┼─────┬─────┤
│   │         │        │           │ 7         │  5  │  3  │
└───┴─────────┴────────┴───────────┴───────────┴─────┴─────┘`)
	})

	t.Run("rows and columns no headers or footers", func(t *testing.T) {
		tw := NewWriter()
		tw.AppendRow(Row{"a.a.a.a", "Pod 1A", "NS 1A", "C 1", "Y", "Y"}, rcAutoMerge)
		tw.AppendRow(Row{"a.a.a.a", "Pod 1A", "NS 1A", "C 2", "Y", "N"})
		tw.AppendRow(Row{"a.a.a.a", "Pod 1A", "NS 1B", "C 3", "N", "N"})
		tw.AppendRow(Row{"a.a.a.a", "Pod 1B", "NS 2", "C 4", "N", "N"}, rcAutoMerge)
		tw.AppendRow(Row{"a.a.a.a", "Pod 1B", "NS 2", "C 5", "Y", "N"})
		tw.AppendRow(Row{"b.b.b.b", "Pod 2", "NS 3", "C 6", "Y", "Y"}, rcAutoMerge)
		tw.AppendRow(Row{"b.b.b.b", "Pod 2", "NS 3", "C 7", "Y", "Y"}, RowConfig{AutoMerge: true, AutoMergeAlign: text.AlignRight})
		tw.SetColumnConfigs([]ColumnConfig{
			{Number: 5, Align: text.AlignCenter, AlignHeader: text.AlignCenter},
			{Number: 6, Align: text.AlignCenter, AlignHeader: text.AlignCenter},
		})
		tw.SetStyle(StyleLight)
		tw.Style().Options.SeparateRows = true

		compareOutput(t, tw.Render(), `
┌─────────┬────────┬───────┬─────┬───────┐
│ a.a.a.a │ Pod 1A │ NS 1A │ C 1 │   Y   │
├─────────┼────────┼───────┼─────┼───┬───┤
│ a.a.a.a │ Pod 1A │ NS 1A │ C 2 │ Y │ N │
├─────────┼────────┼───────┼─────┼───┼───┤
│ a.a.a.a │ Pod 1A │ NS 1B │ C 3 │ N │ N │
├─────────┼────────┼───────┼─────┼───┴───┤
│ a.a.a.a │ Pod 1B │ NS 2  │ C 4 │   N   │
├─────────┼────────┼───────┼─────┼───┬───┤
│ a.a.a.a │ Pod 1B │ NS 2  │ C 5 │ Y │ N │
├─────────┼────────┼───────┼─────┼───┴───┤
│ b.b.b.b │ Pod 2  │ NS 3  │ C 6 │   Y   │
├─────────┼────────┼───────┼─────┼───────┤
│ b.b.b.b │ Pod 2  │ NS 3  │ C 7 │     Y │
└─────────┴────────┴───────┴─────┴───────┘`)
	})

	t.Run("rows and columns no headers or footers with auto-index", func(t *testing.T) {
		tw := NewWriter()
		tw.AppendRow(Row{"a.a.a.a", "Pod 1A", "NS 1A", "C 1", "Y", "Y"}, rcAutoMerge)
		tw.AppendRow(Row{"a.a.a.a", "Pod 1A", "NS 1A", "C 2", "Y", "N"})
		tw.AppendRow(Row{"a.a.a.a", "Pod 1A", "NS 1B", "C 3", "N", "N"})
		tw.AppendRow(Row{"a.a.a.a", "Pod 1B", "NS 2", "C 4", "N", "N"}, rcAutoMerge)
		tw.AppendRow(Row{"a.a.a.a", "Pod 1B", "NS 2", "C 5", "Y", "N"})
		tw.AppendRow(Row{"b.b.b.b", "Pod 2", "NS 3", "C 6", "Y", "Y"}, rcAutoMerge)
		tw.AppendRow(Row{"b.b.b.b", "Pod 2", "NS 3", "C 7", "Y", "Y"}, RowConfig{AutoMerge: true, AutoMergeAlign: text.AlignRight})
		tw.SetAutoIndex(true)
		tw.SetColumnConfigs([]ColumnConfig{
			{Number: 5, Align: text.AlignCenter, AlignHeader: text.AlignCenter},
			{Number: 6, Align: text.AlignCenter, AlignHeader: text.AlignCenter},
		})
		tw.SetStyle(StyleLight)
		tw.Style().Options.SeparateRows = true

		compareOutput(t, tw.Render(), `
┌───┬─────────┬────────┬───────┬─────┬───┬───┐
│   │    A    │    B   │   C   │  D  │ E │ F │
├───┼─────────┼────────┼───────┼─────┼───┴───┤
│ 1 │ a.a.a.a │ Pod 1A │ NS 1A │ C 1 │   Y   │
├───┼─────────┼────────┼───────┼─────┼───┬───┤
│ 2 │ a.a.a.a │ Pod 1A │ NS 1A │ C 2 │ Y │ N │
├───┼─────────┼────────┼───────┼─────┼───┼───┤
│ 3 │ a.a.a.a │ Pod 1A │ NS 1B │ C 3 │ N │ N │
├───┼─────────┼────────┼───────┼─────┼───┴───┤
│ 4 │ a.a.a.a │ Pod 1B │ NS 2  │ C 4 │   N   │
├───┼─────────┼────────┼───────┼─────┼───┬───┤
│ 5 │ a.a.a.a │ Pod 1B │ NS 2  │ C 5 │ Y │ N │
├───┼─────────┼────────┼───────┼─────┼───┴───┤
│ 6 │ b.b.b.b │ Pod 2  │ NS 3  │ C 6 │   Y   │
├───┼─────────┼────────┼───────┼─────┼───────┤
│ 7 │ b.b.b.b │ Pod 2  │ NS 3  │ C 7 │     Y │
└───┴─────────┴────────┴───────┴─────┴───────┘`)
	})

	t.Run("rows and columns and headers and footers", func(t *testing.T) {
		tw := NewWriter()
		tw.AppendHeader(Row{"Node IP", "Pods", "Namespace", "Container", "RCE", "RCE", "ID"}, rcAutoMerge)
		tw.AppendHeader(Row{"Node IP", "Pods", "Namespace", "Container", "EXE", "RUN", ""})
		tw.AppendRow(Row{"a.a.a.a", "Pod 1A", "NS 1A", "C 1", "Y", "Y", 123}, rcAutoMerge)
		tw.AppendRow(Row{"a.a.a.a", "Pod 1A", "NS 1A", "C 2", "Y", "N", 234})
		tw.AppendRow(Row{"a.a.a.a", "Pod 1A", "NS 1B", "C 3", "N", "N", 345})
		tw.AppendRow(Row{"a.a.a.a", "Pod 1B", "NS 2", "C 4", "N", "N", 456}, rcAutoMerge)
		tw.AppendRow(Row{"a.a.a.a", "Pod 1B", "NS 2", "C 5", "Y", "N", 567})
		tw.AppendRow(Row{"b.b.b.b", "Pod 2", "NS 3", "C 6", "Y", "Y", 678}, rcAutoMerge)
		tw.AppendRow(Row{"b.b.b.b", "Pod 2", "NS 3", "C 7", "Y", "Y", 789}, rcAutoMerge)
		tw.AppendFooter(Row{"", "", "", 7, 5, 5}, rcAutoMerge)
		tw.AppendFooter(Row{"", "", "", 7, 5, 3}, rcAutoMerge)
		tw.AppendFooter(Row{"", "", "", 7, 5, 5}, rcAutoMerge)
		tw.AppendFooter(Row{"", "", "", 7, 5, 3}, rcAutoMerge)
		tw.AppendFooter(Row{"", "", "", 7, 5, 5}, rcAutoMerge)
		tw.SetAutoIndex(true)
		tw.SetColumnConfigs([]ColumnConfig{
			{Number: 1, AutoMerge: true},
			{Number: 2, AutoMerge: true},
			{Number: 3, AutoMerge: true},
			{Number: 4, AutoMerge: true},
			{Number: 5, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter},
			{Number: 6, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter},
		})
		tw.SetStyle(StyleLight)
		tw.Style().Options.SeparateRows = true

		compareOutput(t, tw.Render(), `
┌───┬─────────┬────────┬───────────┬───────────┬───────────┬─────┐
│   │ NODE IP │ PODS   │ NAMESPACE │ CONTAINER │    RCE    │  ID │
│   │         │        │           │           ├─────┬─────┼─────┤
│   │         │        │           │           │ EXE │ RUN │     │
├───┼─────────┼────────┼───────────┼───────────┼─────┴─────┼─────┤
│ 1 │ a.a.a.a │ Pod 1A │ NS 1A     │ C 1       │     Y     │ 123 │
├───┤         │        │           ├───────────┼─────┬─────┼─────┤
│ 2 │         │        │           │ C 2       │  Y  │  N  │ 234 │
├───┤         │        ├───────────┼───────────┼─────┼─────┼─────┤
│ 3 │         │        │ NS 1B     │ C 3       │  N  │  N  │ 345 │
├───┤         ├────────┼───────────┼───────────┼─────┴─────┼─────┤
│ 4 │         │ Pod 1B │ NS 2      │ C 4       │     N     │ 456 │
├───┤         │        │           ├───────────┼─────┬─────┼─────┤
│ 5 │         │        │           │ C 5       │  Y  │  N  │ 567 │
├───┼─────────┼────────┼───────────┼───────────┼─────┴─────┼─────┤
│ 6 │ b.b.b.b │ Pod 2  │ NS 3      │ C 6       │     Y     │ 678 │
├───┤         │        │           ├───────────┼───────────┼─────┤
│ 7 │         │        │           │ C 7       │     Y     │ 789 │
├───┼─────────┴────────┴───────────┼───────────┼───────────┼─────┤
│   │                              │ 7         │     5     │     │
│   │                              │           ├─────┬─────┼─────┤
│   │                              │           │  5  │  3  │     │
│   │                              │           ├─────┴─────┼─────┤
│   │                              │           │     5     │     │
│   │                              │           ├─────┬─────┼─────┤
│   │                              │           │  5  │  3  │     │
│   │                              │           ├─────┴─────┼─────┤
│   │                              │           │     5     │     │
└───┴──────────────────────────────┴───────────┴───────────┴─────┘`)
	})

	t.Run("samurai sudoku", func(t *testing.T) {
		tw := NewWriter()
		tw.AppendRow(Row{"1.1\n1.1", "1.2\n1.2", "1.3\n1.3", " ", "2.1\n2.1", "2.2\n2.2", "2.3\n2.3"})
		tw.AppendRow(Row{"1.4\n1.4", "1.5\n1.5", "1.6\n1.6", " ", "2.4\n2.4", "2.5\n2.5", "2.6\n2.6"})
		tw.AppendRow(Row{"1.7\n1.7", "1.8\n1.8", "1.9\n0.1", "0.2\n0.2", "2.7\n0.3", "2.8\n2.8", "2.9\n2.9"})
		tw.AppendRow(Row{" ", " ", "0.4\n0.4", "0.5\n0.5", "0.6\n0.6", " ", " "}, rcAutoMerge)
		tw.AppendRow(Row{"3.1\n3.1", "3.2\n3.2", "3.3\n0.7", "0.8\n0.8", "4.1\n0.9", "4.2\n4.2", "4.3\n4.3"})
		tw.AppendRow(Row{"3.4\n3.4", "3.5\n3.5", "3.6\n3.6", " ", "4.4\n4.4", "4.5\n4.5", "4.6\n4.6"})
		tw.AppendRow(Row{"3.7\n3.7", "3.8\n3.8", "3.9\n3.9", " ", "4.7\n4.7", "4.8\n4.8", "4.9\n4.9"})
		tw.SetColumnConfigs([]ColumnConfig{
			{Number: 4, AutoMerge: true},
		})
		tw.SetStyle(StyleLight)
		tw.Style().Box.PaddingLeft = ""
		tw.Style().Box.PaddingRight = ""
		tw.Style().Options.DrawBorder = true
		tw.Style().Options.SeparateRows = true
		tw.Style().Options.SeparateColumns = true

		compareOutput(t, tw.Render(), `
┌───┬───┬───┬───┬───┬───┬───┐
│1.1│1.2│1.3│   │2.1│2.2│2.3│
│1.1│1.2│1.3│   │2.1│2.2│2.3│
├───┼───┼───┤   ├───┼───┼───┤
│1.4│1.5│1.6│   │2.4│2.5│2.6│
│1.4│1.5│1.6│   │2.4│2.5│2.6│
├───┼───┼───┼───┼───┼───┼───┤
│1.7│1.8│1.9│0.2│2.7│2.8│2.9│
│1.7│1.8│0.1│0.2│0.3│2.8│2.9│
├───┴───┼───┼───┼───┼───┴───┤
│       │0.4│0.5│0.6│       │
│       │0.4│0.5│0.6│       │
├───┬───┼───┼───┼───┼───┬───┤
│3.1│3.2│3.3│0.8│4.1│4.2│4.3│
│3.1│3.2│0.7│0.8│0.9│4.2│4.3│
├───┼───┼───┼───┼───┼───┼───┤
│3.4│3.5│3.6│   │4.4│4.5│4.6│
│3.4│3.5│3.6│   │4.4│4.5│4.6│
├───┼───┼───┤   ├───┼───┼───┤
│3.7│3.8│3.9│   │4.7│4.8│4.9│
│3.7│3.8│3.9│   │4.7│4.8│4.9│
└───┴───┴───┴───┴───┴───┴───┘`)
	})

	t.Run("long column no merge", func(t *testing.T) {
		tw := NewWriter()
		tw.AppendHeader(Row{"Column 1", "Column 2", "Column 3", "Column 4", "Column 5", "Column 6", "Column 7", "Column 8"}, rcAutoMerge)
		tw.AppendRow(Row{"a.a.a.a", "Pod 1A", "NS 1A", "C 1", "4F8F5CB531E3D49A61CF417CD133792CCFA501FD8DA53EE368FED20E5FE0248C3A0B64F98A6533CEE1DA614C3A8DDEC791FF05FEE6D971D57C1348320F4EB42DR", "4F8F5CB531E3D49A61CF417CD133792CCFA501FD8DA53EE368FED20E5FE0248C3A0B64F98A6533CEE1DA614C3A8DDEC791FF05FEE6D971D57C1348320F4EB42DRW", "4F8F5CB531E3D49A61CF417CD133792CCFA501FD8DA53EE368FED20E5FE0248C3A0B64F98A6533CEE1DA614C3A8DDEC791FF05FEE6D971D57C1348320F4EB42DRH", "4F8F5CB531E3D49A61CF417CD133792CCFA501FD8DA53EE368FED20E5FE0248C3A0B64F98A6533CEE1DA614C3A8DDEC791FF05FEE6D971D57C1348320F4EB42DRY"}, rcAutoMerge)
		tw.AppendRow(Row{"a.a.a.a", "Pod 1A", "NS 1A", "C 2", "Y", "Y", "Y", "Y"}, rcAutoMerge)
		tw.AppendRow(Row{"a.a.a.a", "Pod 1A", "NS 1A", "C 2", "Y", "Y", "Y", "Y"}, rcAutoMerge)
		tw.SetAutoIndex(true)
		tw.SetColumnConfigs([]ColumnConfig{
			{Number: 5, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter, WidthMax: 24, WidthMaxEnforcer: text.WrapHard},
			{Number: 6, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter, WidthMax: 24, WidthMaxEnforcer: text.WrapHard},
			{Number: 7, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter, WidthMax: 24, WidthMaxEnforcer: text.WrapHard},
			{Number: 8, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter, WidthMax: 24, WidthMaxEnforcer: text.WrapHard},
		})
		tw.SetStyle(StyleLight)
		tw.Style().Options.SeparateRows = true

		compareOutput(t, tw.Render(), `
┌───┬──────────┬──────────┬──────────┬──────────┬──────────────────────────┬──────────────────────────┬──────────────────────────┬──────────────────────────┐
│   │ COLUMN 1 │ COLUMN 2 │ COLUMN 3 │ COLUMN 4 │         COLUMN 5         │         COLUMN 6         │         COLUMN 7         │         COLUMN 8         │
├───┼──────────┼──────────┼──────────┼──────────┼──────────────────────────┼──────────────────────────┼──────────────────────────┼──────────────────────────┤
│ 1 │ a.a.a.a  │ Pod 1A   │ NS 1A    │ C 1      │ 4F8F5CB531E3D49A61CF417C │ 4F8F5CB531E3D49A61CF417C │ 4F8F5CB531E3D49A61CF417C │ 4F8F5CB531E3D49A61CF417C │
│   │          │          │          │          │ D133792CCFA501FD8DA53EE3 │ D133792CCFA501FD8DA53EE3 │ D133792CCFA501FD8DA53EE3 │ D133792CCFA501FD8DA53EE3 │
│   │          │          │          │          │ 68FED20E5FE0248C3A0B64F9 │ 68FED20E5FE0248C3A0B64F9 │ 68FED20E5FE0248C3A0B64F9 │ 68FED20E5FE0248C3A0B64F9 │
│   │          │          │          │          │ 8A6533CEE1DA614C3A8DDEC7 │ 8A6533CEE1DA614C3A8DDEC7 │ 8A6533CEE1DA614C3A8DDEC7 │ 8A6533CEE1DA614C3A8DDEC7 │
│   │          │          │          │          │ 91FF05FEE6D971D57C134832 │ 91FF05FEE6D971D57C134832 │ 91FF05FEE6D971D57C134832 │ 91FF05FEE6D971D57C134832 │
│   │          │          │          │          │         0F4EB42DR        │        0F4EB42DRW        │        0F4EB42DRH        │        0F4EB42DRY        │
├───┼──────────┼──────────┼──────────┼──────────┼──────────────────────────┴──────────────────────────┴──────────────────────────┴──────────────────────────┤
│ 2 │ a.a.a.a  │ Pod 1A   │ NS 1A    │ C 2      │                                                     Y                                                     │
├───┼──────────┼──────────┼──────────┼──────────┼───────────────────────────────────────────────────────────────────────────────────────────────────────────┤
│ 3 │ a.a.a.a  │ Pod 1A   │ NS 1A    │ C 2      │                                                     Y                                                     │
└───┴──────────┴──────────┴──────────┴──────────┴───────────────────────────────────────────────────────────────────────────────────────────────────────────┘`)
	})

	t.Run("long column partially merged #1", func(t *testing.T) {
		tw := NewWriter()
		tw.AppendHeader(Row{"Column 1", "Column 2", "Column 3", "Column 4", "Column 5", "Column 6", "Column 7", "Column 8"}, rcAutoMerge)
		tw.AppendRow(Row{"a.a.a.a", "Pod 1A", "NS 1A", "C 1", "4F8F5CB531E3D49A61CF417CD133792CCFA501FD8DA53EE368FED20E5FE0248C3A0B64F98A6533CEE1DA614C3A8DDEC791FF05FEE6D971D57C1348320F4EB42DR", "4F8F5CB531E3D49A61CF417CD133792CCFA501FD8DA53EE368FED20E5FE0248C3A0B64F98A6533CEE1DA614C3A8DDEC791FF05FEE6D971D57C1348320F4EB42DR", "4F8F5CB531E3D49A61CF417CD133792CCFA501FD8DA53EE368FED20E5FE0248C3A0B64F98A6533CEE1DA614C3A8DDEC791FF05FEE6D971D57C1348320F4EB42DRR", "4F8F5CB531E3D49A61CF417CD133792CCFA501FD8DA53EE368FED20E5FE0248C3A0B64F98A6533CEE1DA614C3A8DDEC791FF05FEE6D971D57C1348320F4EB42DRR"}, rcAutoMerge)
		tw.AppendRow(Row{"a.a.a.a", "Pod 1A", "NS 1A", "C 2", "Y", "Y", "Y", "Y"}, rcAutoMerge)
		tw.AppendRow(Row{"a.a.a.a", "Pod 1A", "NS 1A", "C 2", "Y", "Y", "Y", "Y"}, rcAutoMerge)
		tw.SetAutoIndex(true)
		tw.SetColumnConfigs([]ColumnConfig{
			{Number: 5, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter, WidthMax: 24, WidthMaxEnforcer: text.WrapHard},
			{Number: 6, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter, WidthMax: 24, WidthMaxEnforcer: text.WrapHard},
			{Number: 7, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter, WidthMax: 24, WidthMaxEnforcer: text.WrapHard},
			{Number: 8, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter, WidthMax: 24, WidthMaxEnforcer: text.WrapHard},
		})
		tw.SetStyle(StyleLight)
		tw.Style().Options.SeparateRows = true

		compareOutput(t, tw.Render(), `
┌───┬──────────┬──────────┬──────────┬──────────┬─────────────┬─────────────┬─────────────┬─────────────┐
│   │ COLUMN 1 │ COLUMN 2 │ COLUMN 3 │ COLUMN 4 │   COLUMN 5  │   COLUMN 6  │   COLUMN 7  │   COLUMN 8  │
├───┼──────────┼──────────┼──────────┼──────────┼─────────────┴─────────────┼─────────────┴─────────────┤
│ 1 │ a.a.a.a  │ Pod 1A   │ NS 1A    │ C 1      │  4F8F5CB531E3D49A61CF417C │  4F8F5CB531E3D49A61CF417C │
│   │          │          │          │          │  D133792CCFA501FD8DA53EE3 │  D133792CCFA501FD8DA53EE3 │
│   │          │          │          │          │  68FED20E5FE0248C3A0B64F9 │  68FED20E5FE0248C3A0B64F9 │
│   │          │          │          │          │  8A6533CEE1DA614C3A8DDEC7 │  8A6533CEE1DA614C3A8DDEC7 │
│   │          │          │          │          │  91FF05FEE6D971D57C134832 │  91FF05FEE6D971D57C134832 │
│   │          │          │          │          │         0F4EB42DR         │         0F4EB42DRR        │
├───┼──────────┼──────────┼──────────┼──────────┼───────────────────────────┴───────────────────────────┤
│ 2 │ a.a.a.a  │ Pod 1A   │ NS 1A    │ C 2      │                           Y                           │
├───┼──────────┼──────────┼──────────┼──────────┼───────────────────────────────────────────────────────┤
│ 3 │ a.a.a.a  │ Pod 1A   │ NS 1A    │ C 2      │                           Y                           │
└───┴──────────┴──────────┴──────────┴──────────┴───────────────────────────────────────────────────────┘`)
	})

	t.Run("long column partially merged #2", func(t *testing.T) {
		tw := NewWriter()
		tw.AppendHeader(Row{"Column 1", "Column 2", "Column 3", "Column 4", "Column 5", "Column 6", "Column 7", "Column 8"}, rcAutoMerge)
		tw.AppendRow(Row{"a.a.a.a", "Pod 1A", "NS 1A", "C 1", "4F8F5CB531E3D49A61CF417CD133792CCFA501FD8DA53EE368FED20E5FE0248C3A0B64F98A6533CEE1DA614C3A8DDEC791FF05FEE6D971D57C1348320F4EB42DR", "4F8F5CB531E3D49A61CF417CD133792CCFA501FD8DA53EE368FED20E5FE0248C3A0B64F98A6533CEE1DA614C3A8DDEC791FF05FEE6D971D57C1348320F4EB42DR", "4F8F5CB531E3D49A61CF417CD133792CCFA501FD8DA53EE368FED20E5FE0248C3A0B64F98A6533CEE1DA614C3A8DDEC791FF05FEE6D971D57C1348320F4EB42DR", "4F8F5CB531E3D49A61CF417CD133792CCFA501FD8DA53EE368FED20E5FE0248C3A0B64F98A6533CEE1DA614C3A8DDEC791FF05FEE6D971D57C1348320F4EB42DRE"}, rcAutoMerge)
		tw.AppendRow(Row{"a.a.a.a", "Pod 1A", "NS 1A", "C 2", "Y", "Y", "Y", "Y"}, rcAutoMerge)
		tw.AppendRow(Row{"a.a.a.a", "Pod 1A", "NS 1A", "C 2", "Y", "Y", "Y", "Y"}, rcAutoMerge)
		tw.SetAutoIndex(true)
		tw.SetColumnConfigs([]ColumnConfig{
			{Number: 5, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter, WidthMax: 24, WidthMaxEnforcer: text.WrapHard},
			{Number: 6, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter, WidthMax: 24, WidthMaxEnforcer: text.WrapHard},
			{Number: 7, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter, WidthMax: 24, WidthMaxEnforcer: text.WrapHard},
			{Number: 8, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter, WidthMax: 24, WidthMaxEnforcer: text.WrapHard},
		})
		tw.SetStyle(StyleLight)
		tw.Style().Options.SeparateRows = true

		compareOutput(t, tw.Render(), `
┌───┬──────────┬──────────┬──────────┬──────────┬──────────┬──────────┬──────────┬──────────────────────────┐
│   │ COLUMN 1 │ COLUMN 2 │ COLUMN 3 │ COLUMN 4 │ COLUMN 5 │ COLUMN 6 │ COLUMN 7 │         COLUMN 8         │
├───┼──────────┼──────────┼──────────┼──────────┼──────────┴──────────┴──────────┼──────────────────────────┤
│ 1 │ a.a.a.a  │ Pod 1A   │ NS 1A    │ C 1      │    4F8F5CB531E3D49A61CF417C    │ 4F8F5CB531E3D49A61CF417C │
│   │          │          │          │          │    D133792CCFA501FD8DA53EE3    │ D133792CCFA501FD8DA53EE3 │
│   │          │          │          │          │    68FED20E5FE0248C3A0B64F9    │ 68FED20E5FE0248C3A0B64F9 │
│   │          │          │          │          │    8A6533CEE1DA614C3A8DDEC7    │ 8A6533CEE1DA614C3A8DDEC7 │
│   │          │          │          │          │    91FF05FEE6D971D57C134832    │ 91FF05FEE6D971D57C134832 │
│   │          │          │          │          │            0F4EB42DR           │        0F4EB42DRE        │
├───┼──────────┼──────────┼──────────┼──────────┼────────────────────────────────┴──────────────────────────┤
│ 2 │ a.a.a.a  │ Pod 1A   │ NS 1A    │ C 2      │                             Y                             │
├───┼──────────┼──────────┼──────────┼──────────┼───────────────────────────────────────────────────────────┤
│ 3 │ a.a.a.a  │ Pod 1A   │ NS 1A    │ C 2      │                             Y                             │
└───┴──────────┴──────────┴──────────┴──────────┴───────────────────────────────────────────────────────────┘`)
	})

	t.Run("long column fully merged", func(t *testing.T) {
		tw := NewWriter()
		tw.AppendHeader(Row{"Column 1", "Column 2", "Column 3", "Column 4", "Column 5", "Column 6", "Column 7", "Column 8"}, rcAutoMerge)
		tw.AppendRow(Row{"a.a.a.a", "Pod 1A", "NS 1A", "C 1", "4F8F5CB531E3D49A61CF417CD133792CCFA501FD8DA53EE368FED20E5FE0248C3A0B64F98A6533CEE1DA614C3A8DDEC791FF05FEE6D971D57C1348320F4EB42DR", "4F8F5CB531E3D49A61CF417CD133792CCFA501FD8DA53EE368FED20E5FE0248C3A0B64F98A6533CEE1DA614C3A8DDEC791FF05FEE6D971D57C1348320F4EB42DR", "4F8F5CB531E3D49A61CF417CD133792CCFA501FD8DA53EE368FED20E5FE0248C3A0B64F98A6533CEE1DA614C3A8DDEC791FF05FEE6D971D57C1348320F4EB42DR", "4F8F5CB531E3D49A61CF417CD133792CCFA501FD8DA53EE368FED20E5FE0248C3A0B64F98A6533CEE1DA614C3A8DDEC791FF05FEE6D971D57C1348320F4EB42DR"}, rcAutoMerge)
		tw.AppendRow(Row{"a.a.a.a", "Pod 1A", "NS 1A", "C 2", "Y", "Y", "Y", "Y"}, rcAutoMerge)
		tw.AppendRow(Row{"a.a.a.a", "Pod 1A", "NS 1A", "C 2", "Y", "Y", "Y", "Y"}, rcAutoMerge)
		tw.SetAutoIndex(true)
		tw.SetColumnConfigs([]ColumnConfig{
			{Number: 5, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter, WidthMax: 24, WidthMaxEnforcer: text.WrapHard},
			{Number: 6, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter, WidthMax: 24, WidthMaxEnforcer: text.WrapHard},
			{Number: 7, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter, WidthMax: 24, WidthMaxEnforcer: text.WrapHard},
			{Number: 8, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter, WidthMax: 24, WidthMaxEnforcer: text.WrapHard},
		})
		tw.SetStyle(StyleLight)
		tw.Style().Options.SeparateRows = true

		compareOutput(t, tw.Render(), `
┌───┬──────────┬──────────┬──────────┬──────────┬──────────┬──────────┬──────────┬──────────┐
│   │ COLUMN 1 │ COLUMN 2 │ COLUMN 3 │ COLUMN 4 │ COLUMN 5 │ COLUMN 6 │ COLUMN 7 │ COLUMN 8 │
├───┼──────────┼──────────┼──────────┼──────────┼──────────┴──────────┴──────────┴──────────┤
│ 1 │ a.a.a.a  │ Pod 1A   │ NS 1A    │ C 1      │          4F8F5CB531E3D49A61CF417C         │
│   │          │          │          │          │          D133792CCFA501FD8DA53EE3         │
│   │          │          │          │          │          68FED20E5FE0248C3A0B64F9         │
│   │          │          │          │          │          8A6533CEE1DA614C3A8DDEC7         │
│   │          │          │          │          │          91FF05FEE6D971D57C134832         │
│   │          │          │          │          │                 0F4EB42DR                 │
├───┼──────────┼──────────┼──────────┼──────────┼───────────────────────────────────────────┤
│ 2 │ a.a.a.a  │ Pod 1A   │ NS 1A    │ C 2      │                     Y                     │
├───┼──────────┼──────────┼──────────┼──────────┼───────────────────────────────────────────┤
│ 3 │ a.a.a.a  │ Pod 1A   │ NS 1A    │ C 2      │                     Y                     │
└───┴──────────┴──────────┴──────────┴──────────┴───────────────────────────────────────────┘`)
	})

	t.Run("headers and footers", func(t *testing.T) {
		tw := NewWriter()
		tw.AppendHeader(Row{"Node IP", "Pods", "Namespace", "Container", "RCE1", "RCE2"}, rcAutoMerge)
		tw.AppendHeader(Row{"", "", "", "", "EXE EXE EXE", "EXE EXE EXE"}, rcAutoMerge)
		tw.AppendRow(Row{"a.a.a.a", "Pod 1A", "NS 1A", "C 1", "Y", "Y"}, rcAutoMerge)
		tw.AppendRow(Row{"a.a.a.a", "Pod 1A", "NS 1A", "C 2", "Y", "Y"}, rcAutoMerge)
		tw.AppendRow(Row{"a.a.a.a", "Pod 1A", "NS 1B", "C 3", "N", "N"}, rcAutoMerge)
		tw.AppendRow(Row{"a.a.a.a", "Pod 1B", "NS 2", "C 4", "N", "N"}, rcAutoMerge)
		tw.AppendRow(Row{"a.a.a.a", "Pod 1B", "NS 2", "C 5", "Y", "Y"}, rcAutoMerge)
		tw.AppendRow(Row{"b.b.b.b", "Pod 2", "NS 3", "C 6", "Y", "Y"}, rcAutoMerge)
		tw.AppendRow(Row{"b.b.b.b", "Pod 2", "NS 3", "C 7", "Y", "Y"}, rcAutoMerge)
		tw.AppendFooter(Row{"", "", "", 7, 5, 5}, rcAutoMerge)
		tw.AppendFooter(Row{"", "", "", 6, 4, 4}, rcAutoMerge)
		tw.SetAutoIndex(true)
		tw.SetColumnConfigs([]ColumnConfig{
			{Number: 5, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter, WidthMax: 7, WidthMaxEnforcer: text.WrapHard},
			{Number: 6, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter, WidthMax: 7, WidthMaxEnforcer: text.WrapHard},
		})
		tw.SetStyle(StyleLight)
		tw.Style().Options.SeparateRows = true

		compareOutput(t, tw.Render(), `
┌───┬─────────┬────────┬───────────┬───────────┬──────┬──────┐
│   │ NODE IP │ PODS   │ NAMESPACE │ CONTAINER │ RCE1 │ RCE2 │
│   ├─────────┴────────┴───────────┴───────────┼──────┴──────┤
│   │                                          │   EXE EXE   │
│   │                                          │     EXE     │
├───┼─────────┬────────┬───────────┬───────────┼─────────────┤
│ 1 │ a.a.a.a │ Pod 1A │ NS 1A     │ C 1       │      Y      │
├───┼─────────┼────────┼───────────┼───────────┼─────────────┤
│ 2 │ a.a.a.a │ Pod 1A │ NS 1A     │ C 2       │      Y      │
├───┼─────────┼────────┼───────────┼───────────┼─────────────┤
│ 3 │ a.a.a.a │ Pod 1A │ NS 1B     │ C 3       │      N      │
├───┼─────────┼────────┼───────────┼───────────┼─────────────┤
│ 4 │ a.a.a.a │ Pod 1B │ NS 2      │ C 4       │      N      │
├───┼─────────┼────────┼───────────┼───────────┼─────────────┤
│ 5 │ a.a.a.a │ Pod 1B │ NS 2      │ C 5       │      Y      │
├───┼─────────┼────────┼───────────┼───────────┼─────────────┤
│ 6 │ b.b.b.b │ Pod 2  │ NS 3      │ C 6       │      Y      │
├───┼─────────┼────────┼───────────┼───────────┼─────────────┤
│ 7 │ b.b.b.b │ Pod 2  │ NS 3      │ C 7       │      Y      │
├───┼─────────┴────────┴───────────┼───────────┼─────────────┤
│   │                              │ 7         │      5      │
│   ├──────────────────────────────┼───────────┼─────────────┤
│   │                              │ 6         │      4      │
└───┴──────────────────────────────┴───────────┴─────────────┘`)
	})

	t.Run("long header column", func(t *testing.T) {
		tw := NewWriter()
		tw.AppendHeader(Row{"Node IP", "Pods", "Namespace", "Container", "RCE1", "RCE2", "RCE3"}, rcAutoMerge)
		tw.AppendHeader(Row{"", "", "", "", "EXE EXE EXE", "EXE EXE EXE", "EXE EXE EXE"}, rcAutoMerge)
		tw.AppendRow(Row{"a.a.a.a", "Pod 1A", "NS 1A", "C 1", "Y", "Y", "Y"}, rcAutoMerge)
		tw.AppendRow(Row{"a.a.a.a", "Pod 1A", "NS 1A", "C 2", "Y", "Y", "Y"}, rcAutoMerge)
		tw.AppendRow(Row{"a.a.a.a", "Pod 1A", "NS 1B", "C 3", "N", "N", "N"}, rcAutoMerge)
		tw.AppendRow(Row{"a.a.a.a", "Pod 1B", "NS 2", "C 4", "N", "N", "N"}, rcAutoMerge)
		tw.AppendRow(Row{"a.a.a.a", "Pod 1B", "NS 2", "C 5", "Y", "Y", "Y"}, rcAutoMerge)
		tw.AppendRow(Row{"b.b.b.b", "Pod 2", "NS 3", "C 6", "Y", "Y", "Y"}, rcAutoMerge)
		tw.AppendRow(Row{"b.b.b.b", "Pod 2", "NS 3", "C 7", "Y", "Y", "Y"}, rcAutoMerge)
		tw.AppendFooter(Row{"", "", "", 7, 5, 5, 5}, rcAutoMerge)
		tw.AppendFooter(Row{"", "", "", 6, 4, 4, 3}, rcAutoMerge)
		tw.SetAutoIndex(true)
		tw.SetColumnConfigs([]ColumnConfig{
			{Number: 5, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter, WidthMax: 7, WidthMaxEnforcer: text.WrapHard},
			{Number: 6, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter, WidthMax: 7, WidthMaxEnforcer: text.WrapHard},
			{Number: 7, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter, WidthMax: 7, WidthMaxEnforcer: text.WrapHard},
		})
		tw.SetStyle(StyleLight)
		tw.Style().Options.SeparateRows = true

		compareOutput(t, tw.Render(), `
┌───┬─────────┬────────┬───────────┬───────────┬──────┬──────┬──────┐
│   │ NODE IP │ PODS   │ NAMESPACE │ CONTAINER │ RCE1 │ RCE2 │ RCE3 │
│   ├─────────┴────────┴───────────┴───────────┼──────┴──────┴──────┤
│   │                                          │       EXE EXE      │
│   │                                          │         EXE        │
├───┼─────────┬────────┬───────────┬───────────┼────────────────────┤
│ 1 │ a.a.a.a │ Pod 1A │ NS 1A     │ C 1       │          Y         │
├───┼─────────┼────────┼───────────┼───────────┼────────────────────┤
│ 2 │ a.a.a.a │ Pod 1A │ NS 1A     │ C 2       │          Y         │
├───┼─────────┼────────┼───────────┼───────────┼────────────────────┤
│ 3 │ a.a.a.a │ Pod 1A │ NS 1B     │ C 3       │          N         │
├───┼─────────┼────────┼───────────┼───────────┼────────────────────┤
│ 4 │ a.a.a.a │ Pod 1B │ NS 2      │ C 4       │          N         │
├───┼─────────┼────────┼───────────┼───────────┼────────────────────┤
│ 5 │ a.a.a.a │ Pod 1B │ NS 2      │ C 5       │          Y         │
├───┼─────────┼────────┼───────────┼───────────┼────────────────────┤
│ 6 │ b.b.b.b │ Pod 2  │ NS 3      │ C 6       │          Y         │
├───┼─────────┼────────┼───────────┼───────────┼────────────────────┤
│ 7 │ b.b.b.b │ Pod 2  │ NS 3      │ C 7       │          Y         │
├───┼─────────┴────────┴───────────┼───────────┼────────────────────┤
│   │                              │ 7         │          5         │
│   ├──────────────────────────────┼───────────┼─────────────┬──────┤
│   │                              │ 6         │      4      │   3  │
└───┴──────────────────────────────┴───────────┴─────────────┴──────┘`)
	})

	t.Run("empty cells", func(t *testing.T) {
		tw := NewWriter()
		rowConfigAutoMerge := RowConfig{AutoMerge: true}
		tw.AppendRow(Row{"Product", "Standalone", "foo bar", "a.a.a.a", ""}, rowConfigAutoMerge)
		tw.AppendRow(Row{"Test", "Standalone", "bar baz", "b.b.b.b", ""}, rowConfigAutoMerge)
		tw.AppendRow(Row{"Product", "RedisCluster", "foo baz", "", "Cluster #1"}, rowConfigAutoMerge)
		tw.AppendRow(Row{"Product", "RedisCluster", "bar baz", "", "Cluster #2"}, rowConfigAutoMerge)
		tw.SetAutoIndex(true)
		tw.SetColumnConfigs([]ColumnConfig{
			{Number: 1, AutoMerge: true},
			{Number: 2, AutoMerge: true},
			{Number: 3, AutoMerge: true},
			{Number: 4, AutoMerge: true},
			{Number: 5, AutoMerge: true},
		})
		tw.SetStyle(StyleLight)
		tw.Style().Options.SeparateRows = true

		compareOutput(t, tw.Render(), `┌───┬─────────┬──────────────┬─────────┬─────────┬────────────┐
│   │    A    │       B      │    C    │    D    │      E     │
├───┼─────────┼──────────────┼─────────┼─────────┼────────────┤
│ 1 │ Product │ Standalone   │ foo bar │ a.a.a.a │            │
├───┼─────────┤              ├─────────┼─────────┤            │
│ 2 │ Test    │              │ bar baz │ b.b.b.b │            │
├───┼─────────┼──────────────┼─────────┼─────────┼────────────┤
│ 3 │ Product │ RedisCluster │ foo baz │         │ Cluster #1 │
├───┤         │              ├─────────┤         ├────────────┤
│ 4 │         │              │ bar baz │         │ Cluster #2 │
└───┴─────────┴──────────────┴─────────┴─────────┴────────────┘`)
	})

	t.Run("everything", func(t *testing.T) {
		tw := NewWriter()
		tw.AppendHeader(Row{"COLUMNS", "COLUMNS", "COLUMNS", "COLUMNS", "COLUMNS", "COLUMNS", "COLUMNS"}, rcAutoMerge)
		tw.AppendRow(Row{"a.a.a.a", "Pod 1A", "NS 1A", "C 1", "Y", "Y", "Y"}, rcAutoMerge)
		tw.AppendRow(Row{"a.a.a.a", "Pod 1A", "NS 1A", "C 2", "Y", "Y", "N"}, rcAutoMerge)
		tw.AppendRow(Row{"a.a.a.a", "Pod 1A", "NS 1B", "C 3", "N", "N", "N"}, rcAutoMerge)
		tw.AppendRow(Row{"a.a.a.a", "Pod 1B", "NS 2", "C 4", "N", "Y", "N"}, rcAutoMerge)
		tw.AppendRow(Row{"a.a.a.a", "Pod 1B", "NS 2", "C 5", "Y", "Y", "Y"}, rcAutoMerge)
		tw.AppendRow(Row{"b.b.b.b", "Pod 2", "NS 3", "C 6", "N", "Y", "Y"}, rcAutoMerge)
		tw.AppendRow(Row{"b.b.b.b", "Pod 2", "NS 3", "C 7", "Y", "Y", "Y"}, rcAutoMerge)
		tw.AppendFooter(Row{"foo", "foo", "foo", "foo", "bar", "bar", "bar"}, rcAutoMerge)
		tw.AppendFooter(Row{7, 7, 7, 7, 7, 7, 7}, rcAutoMerge)
		tw.SetAutoIndex(true)
		tw.SetColumnConfigs([]ColumnConfig{
			{Number: 1, AutoMerge: true},
			{Number: 2, AutoMerge: true},
			{Number: 3, AutoMerge: true},
			{Number: 5, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter, WidthMax: 7, WidthMaxEnforcer: text.WrapHard},
			{Number: 6, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter, WidthMax: 7, WidthMaxEnforcer: text.WrapHard},
			{Number: 7, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter, WidthMax: 7, WidthMaxEnforcer: text.WrapHard},
		})
		tw.SetStyle(StyleLight)
		tw.Style().Options.SeparateRows = true

		compareOutput(t, tw.Render(), `
┌───┬───────────────────────────────────────────────────┐
│   │                      COLUMNS                      │
├───┼─────────┬─────────┬─────────┬─────────┬───────────┤
│ 1 │ a.a.a.a │ Pod 1A  │ NS 1A   │ C 1     │     Y     │
├───┤         │         │         ├─────────┼───────┬───┤
│ 2 │         │         │         │ C 2     │   Y   │ N │
├───┤         │         ├─────────┼─────────┼───────┴───┤
│ 3 │         │         │ NS 1B   │ C 3     │     N     │
├───┤         ├─────────┼─────────┼─────────┼───┬───┬───┤
│ 4 │         │ Pod 1B  │ NS 2    │ C 4     │ N │ Y │ N │
├───┤         │         │         ├─────────┼───┴───┴───┤
│ 5 │         │         │         │ C 5     │     Y     │
├───┼─────────┼─────────┼─────────┼─────────┼───┬───────┤
│ 6 │ b.b.b.b │ Pod 2   │ NS 3    │ C 6     │ N │   Y   │
├───┤         │         │         ├─────────┼───┴───────┤
│ 7 │         │         │         │ C 7     │     Y     │
├───┼─────────┴─────────┴─────────┴─────────┼───────────┤
│   │                  FOO                  │    BAR    │
│   ├───────────────────────────────────────┴───────────┤
│   │                         7                         │
└───┴───────────────────────────────────────────────────┘`)
	})
}
