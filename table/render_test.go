package table

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/stretchr/testify/assert"
)

func generateColumnConfigsWithHiddenColumns(colsToHide []int) []ColumnConfig {
	cc := []ColumnConfig{
		{
			Name: "#",
			Transformer: func(val interface{}) string {
				return fmt.Sprint(val.(int) + 7)
			},
		}, {
			Name: "First Name",
			Transformer: func(val interface{}) string {
				return fmt.Sprintf(">>%s", val)
			},
		}, {
			Name: "Last Name",
			Transformer: func(val interface{}) string {
				return fmt.Sprintf("%s<<", val)
			},
		}, {
			Name: "Salary",
			Transformer: func(val interface{}) string {
				return fmt.Sprint(val.(int) + 13)
			},
		}, {
			Number: 5,
			Transformer: func(val interface{}) string {
				return fmt.Sprintf("~%s~", val)
			},
		},
	}
	for _, colToHide := range colsToHide {
		cc[colToHide].Hidden = true
	}
	return cc
}

func TestTable_Render(t *testing.T) {
	tw := NewWriter()
	tw.AppendHeader(testHeader)
	tw.AppendRows(testRows)
	tw.AppendRow(testRowMultiLine)
	tw.AppendFooter(testFooter)
	tw.SetCaption(testCaption)
	tw.SetStyle(styleTest)
	tw.SetTitle(testTitle2)

	expectedOut := `(---------------------------------------------------------------------)
[<When you play the Game of Thrones, you win or you die. There is no >]
[<middle ground.                                                     >]
{-----^------------^-----------^--------^-----------------------------}
[<  #>|<FIRST NAME>|<LAST NAME>|<SALARY>|<                           >]
{-----+------------+-----------+--------+-----------------------------}
[<  1>|<Arya      >|<Stark    >|<  3000>|<                           >]
[< 20>|<Jon       >|<Snow     >|<  2000>|<You know nothing, Jon Snow!>]
[<300>|<Tyrion    >|<Lannister>|<  5000>|<                           >]
[<  0>|<Winter    >|<Is       >|<     0>|<Coming.                    >]
[<   >|<          >|<         >|<      >|<The North Remembers!       >]
[<   >|<          >|<         >|<      >|<This is known.             >]
{-----+------------+-----------+--------+-----------------------------}
[<   >|<          >|<TOTAL    >|< 10000>|<                           >]
\-----v------------v-----------v--------v-----------------------------/
A Song of Ice and Fire`
	assert.Equal(t, expectedOut, tw.Render())
}

func TestTable_Render_AutoIndex(t *testing.T) {
	tw := NewWriter()
	for rowIdx := 0; rowIdx < 10; rowIdx++ {
		row := make(Row, 10)
		for colIdx := 0; colIdx < 10; colIdx++ {
			row[colIdx] = fmt.Sprintf("%s%d", AutoIndexColumnID(colIdx), rowIdx+1)
		}
		tw.AppendRow(row)
	}
	tw.SetAutoIndex(true)
	tw.SetStyle(StyleLight)

	expectedOut := `┌────┬─────┬─────┬─────┬─────┬─────┬─────┬─────┬─────┬─────┬─────┐
│    │  A  │  B  │  C  │  D  │  E  │  F  │  G  │  H  │  I  │  J  │
├────┼─────┼─────┼─────┼─────┼─────┼─────┼─────┼─────┼─────┼─────┤
│  1 │ A1  │ B1  │ C1  │ D1  │ E1  │ F1  │ G1  │ H1  │ I1  │ J1  │
│  2 │ A2  │ B2  │ C2  │ D2  │ E2  │ F2  │ G2  │ H2  │ I2  │ J2  │
│  3 │ A3  │ B3  │ C3  │ D3  │ E3  │ F3  │ G3  │ H3  │ I3  │ J3  │
│  4 │ A4  │ B4  │ C4  │ D4  │ E4  │ F4  │ G4  │ H4  │ I4  │ J4  │
│  5 │ A5  │ B5  │ C5  │ D5  │ E5  │ F5  │ G5  │ H5  │ I5  │ J5  │
│  6 │ A6  │ B6  │ C6  │ D6  │ E6  │ F6  │ G6  │ H6  │ I6  │ J6  │
│  7 │ A7  │ B7  │ C7  │ D7  │ E7  │ F7  │ G7  │ H7  │ I7  │ J7  │
│  8 │ A8  │ B8  │ C8  │ D8  │ E8  │ F8  │ G8  │ H8  │ I8  │ J8  │
│  9 │ A9  │ B9  │ C9  │ D9  │ E9  │ F9  │ G9  │ H9  │ I9  │ J9  │
│ 10 │ A10 │ B10 │ C10 │ D10 │ E10 │ F10 │ G10 │ H10 │ I10 │ J10 │
└────┴─────┴─────┴─────┴─────┴─────┴─────┴─────┴─────┴─────┴─────┘`
	assert.Equal(t, expectedOut, tw.Render())
}

func TestTable_Render_AutoMerge(t *testing.T) {
	rowConfigAutoMerge := RowConfig{AutoMerge: true}

	tw := NewWriter()
	tw.AppendHeader(Row{"Node IP", "Pods", "Namespace", "Container", "RCE", "RCE"}, rowConfigAutoMerge)
	tw.AppendHeader(Row{"", "", "", "", "EXE", "RUN"})
	tw.AppendRow(Row{"1.1.1.1", "Pod 1A", "NS 1A", "C 1", "Y", "Y"}, rowConfigAutoMerge)
	tw.AppendRow(Row{"1.1.1.1", "Pod 1A", "NS 1A", "C 2", "Y", "N"}, rowConfigAutoMerge)
	tw.AppendRow(Row{"1.1.1.1", "Pod 1A", "NS 1B", "C 3", "N", "N"}, rowConfigAutoMerge)
	tw.AppendRow(Row{"1.1.1.1", "Pod 1B", "NS 2", "C 4", "N", "N"}, rowConfigAutoMerge)
	tw.AppendRow(Row{"1.1.1.1", "Pod 1B", "NS 2", "C 5", "Y", "N"}, rowConfigAutoMerge)
	tw.AppendRow(Row{"2.2.2.2", "Pod 2", "NS 3", "C 6", "Y", "Y"}, rowConfigAutoMerge)
	tw.AppendRow(Row{"2.2.2.2", "Pod 2", "NS 3", "C 7", "Y", "Y"}, rowConfigAutoMerge)
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

	expectedOut := `┌───┬─────────┬────────┬───────────┬───────────┬───────────┐
│   │ NODE IP │ PODS   │ NAMESPACE │ CONTAINER │    RCE    │
│   │         │        │           │           ├─────┬─────┤
│   │         │        │           │           │ EXE │ RUN │
├───┼─────────┼────────┼───────────┼───────────┼─────┴─────┤
│ 1 │ 1.1.1.1 │ Pod 1A │ NS 1A     │ C 1       │     Y     │
├───┤         │        │           ├───────────┼─────┬─────┤
│ 2 │         │        │           │ C 2       │  Y  │  N  │
├───┤         │        ├───────────┼───────────┼─────┴─────┤
│ 3 │         │        │ NS 1B     │ C 3       │     N     │
├───┤         ├────────┼───────────┼───────────┼───────────┤
│ 4 │         │ Pod 1B │ NS 2      │ C 4       │     N     │
├───┤         │        │           ├───────────┼─────┬─────┤
│ 5 │         │        │           │ C 5       │  Y  │  N  │
├───┼─────────┼────────┼───────────┼───────────┼─────┴─────┤
│ 6 │ 2.2.2.2 │ Pod 2  │ NS 3      │ C 6       │     Y     │
├───┤         │        │           ├───────────┼───────────┤
│ 7 │         │        │           │ C 7       │     Y     │
├───┼─────────┼────────┼───────────┼───────────┼─────┬─────┤
│   │         │        │           │ 7         │  5  │  3  │
└───┴─────────┴────────┴───────────┴───────────┴─────┴─────┘`
	assert.Equal(t, expectedOut, tw.Render())
}

func TestTable_Render_AutoMerge_Complex(t *testing.T) {
	tw := NewWriter()
	tw.AppendHeader(Row{"Node IP", "Pods", "Namespace", "Container", "RCE", "RCE", "ID"}, RowConfig{AutoMerge: true})
	tw.AppendHeader(Row{"", "", "", "", "EXE", "RUN", ""})
	tw.AppendRow(Row{"1.1.1.1", "Pod 1A", "NS 1A", "C 1", "Y", "Y", 123}, RowConfig{AutoMerge: true})
	tw.AppendRow(Row{"1.1.1.1", "Pod 1A", "NS 1A", "C 2", "Y", "N", 234})
	tw.AppendRow(Row{"1.1.1.1", "Pod 1A", "NS 1B", "C 3", "N", "N", 345})
	tw.AppendRow(Row{"1.1.1.1", "Pod 1B", "NS 2", "C 4", "N", "N", 456}, RowConfig{AutoMerge: true})
	tw.AppendRow(Row{"1.1.1.1", "Pod 1B", "NS 2", "C 5", "Y", "N", 567})
	tw.AppendRow(Row{"2.2.2.2", "Pod 2", "NS 3", "C 6", "Y", "Y", 678}, RowConfig{AutoMerge: true})
	tw.AppendRow(Row{"2.2.2.2", "Pod 2", "NS 3", "C 7", "Y", "Y", 789}, RowConfig{AutoMerge: true})
	tw.AppendFooter(Row{"", "", "", 7, 5, 5}, RowConfig{AutoMerge: true})
	tw.AppendFooter(Row{"", "", "", 7, 5, 3}, RowConfig{AutoMerge: true})
	tw.AppendFooter(Row{"", "", "", 7, 5, 5}, RowConfig{AutoMerge: true})
	tw.AppendFooter(Row{"", "", "", 7, 5, 3}, RowConfig{AutoMerge: true})
	tw.AppendFooter(Row{"", "", "", 7, 5, 5}, RowConfig{AutoMerge: true})
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

	expectedOut := `┌───┬─────────┬────────┬───────────┬───────────┬───────────┬─────┐
│   │ NODE IP │ PODS   │ NAMESPACE │ CONTAINER │    RCE    │  ID │
│   │         │        │           │           ├─────┬─────┼─────┤
│   │         │        │           │           │ EXE │ RUN │     │
├───┼─────────┼────────┼───────────┼───────────┼─────┴─────┼─────┤
│ 1 │ 1.1.1.1 │ Pod 1A │ NS 1A     │ C 1       │     Y     │ 123 │
├───┤         │        │           ├───────────┼─────┬─────┼─────┤
│ 2 │         │        │           │ C 2       │  Y  │  N  │ 234 │
├───┤         │        ├───────────┼───────────┼─────┼─────┼─────┤
│ 3 │         │        │ NS 1B     │ C 3       │  N  │  N  │ 345 │
├───┤         ├────────┼───────────┼───────────┼─────┴─────┼─────┤
│ 4 │         │ Pod 1B │ NS 2      │ C 4       │     N     │ 456 │
├───┤         │        │           ├───────────┼─────┬─────┼─────┤
│ 5 │         │        │           │ C 5       │  Y  │  N  │ 567 │
├───┼─────────┼────────┼───────────┼───────────┼─────┴─────┼─────┤
│ 6 │ 2.2.2.2 │ Pod 2  │ NS 3      │ C 6       │     Y     │ 678 │
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
└───┴──────────────────────────────┴───────────┴───────────┴─────┘`
	assert.Equal(t, expectedOut, tw.Render())
}

func TestTable_Render_AutoMerge_ColumnsOnly(t *testing.T) {
	tw := NewWriter()
	tw.AppendHeader(Row{"Node IP", "Pods", "Namespace", "Container", "RCE\nEXE", "RCE\nRUN"})
	tw.AppendRow(Row{"1.1.1.1", "Pod 1A", "NS 1A", "C 1", "Y", "Y"})
	tw.AppendRow(Row{"1.1.1.1", "Pod 1A", "NS 1A", "C 2", "Y", "N"})
	tw.AppendRow(Row{"1.1.1.1", "Pod 1A", "NS 1B", "C 3", "N", "N"})
	tw.AppendRow(Row{"1.1.1.1", "Pod 1B", "NS 2", "C 4", "N", "N"})
	tw.AppendRow(Row{"1.1.1.1", "Pod 1B", "NS 2", "C 5", "Y", "N"})
	tw.AppendRow(Row{"2.2.2.2", "Pod 2", "NS 3", "C 6", "Y", "Y"})
	tw.AppendRow(Row{"2.2.2.2", "Pod 2", "NS 3", "C 7", "Y", "Y"})
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

	expectedOut := `┌───┬─────────┬────────┬───────────┬───────────┬─────┬─────┐
│   │ NODE IP │ PODS   │ NAMESPACE │ CONTAINER │ RCE │ RCE │
│   │         │        │           │           │ EXE │ RUN │
├───┼─────────┼────────┼───────────┼───────────┼─────┼─────┤
│ 1 │ 1.1.1.1 │ Pod 1A │ NS 1A     │ C 1       │  Y  │  Y  │
├───┤         │        │           ├───────────┼─────┼─────┤
│ 2 │         │        │           │ C 2       │  Y  │  N  │
├───┤         │        ├───────────┼───────────┼─────┼─────┤
│ 3 │         │        │ NS 1B     │ C 3       │  N  │  N  │
├───┤         ├────────┼───────────┼───────────┼─────┼─────┤
│ 4 │         │ Pod 1B │ NS 2      │ C 4       │  N  │  N  │
├───┤         │        │           ├───────────┼─────┼─────┤
│ 5 │         │        │           │ C 5       │  Y  │  N  │
├───┼─────────┼────────┼───────────┼───────────┼─────┼─────┤
│ 6 │ 2.2.2.2 │ Pod 2  │ NS 3      │ C 6       │  Y  │  Y  │
├───┤         │        │           ├───────────┼─────┼─────┤
│ 7 │         │        │           │ C 7       │  Y  │  Y  │
├───┼─────────┼────────┼───────────┼───────────┼─────┼─────┤
│   │         │        │           │ 7         │  5  │  3  │
└───┴─────────┴────────┴───────────┴───────────┴─────┴─────┘`
	assert.Equal(t, expectedOut, tw.Render())
}

func TestTable_Render_AutoMerge_RowsOnly(t *testing.T) {
	tw := NewWriter()
	tw.AppendHeader(Row{"Node IP", "Pods", "Namespace", "Container", "RCE", "RCE"}, RowConfig{AutoMerge: true})
	tw.AppendHeader(Row{"", "", "", "", "EXE", "RUN"})
	tw.AppendRow(Row{"1.1.1.1", "Pod 1A", "NS 1A", "C 1", "Y", "Y"}, RowConfig{AutoMerge: true})
	tw.AppendRow(Row{"1.1.1.1", "Pod 1A", "NS 1A", "C 2", "Y", "N"})
	tw.AppendRow(Row{"1.1.1.1", "Pod 1A", "NS 1B", "C 3", "N", "N"})
	tw.AppendRow(Row{"1.1.1.1", "Pod 1B", "NS 2", "C 4", "N", "N"}, RowConfig{AutoMerge: true})
	tw.AppendRow(Row{"1.1.1.1", "Pod 1B", "NS 2", "C 5", "Y", "N"})
	tw.AppendRow(Row{"2.2.2.2", "Pod 2", "NS 3", "C 6", "Y", "Y"}, RowConfig{AutoMerge: true})
	tw.AppendRow(Row{"2.2.2.2", "Pod 2", "NS 3", "C 7", "Y", "Y"}, RowConfig{AutoMerge: true})
	tw.AppendFooter(Row{"", "", "", 7, 5, 3})
	tw.SetAutoIndex(true)
	tw.SetColumnConfigs([]ColumnConfig{
		{Number: 5, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter},
		{Number: 6, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter},
	})
	tw.SetStyle(StyleLight)
	tw.Style().Options.SeparateRows = true

	expectedOut := `┌───┬─────────┬────────┬───────────┬───────────┬───────────┐
│   │ NODE IP │ PODS   │ NAMESPACE │ CONTAINER │    RCE    │
│   ├─────────┼────────┼───────────┼───────────┼─────┬─────┤
│   │         │        │           │           │ EXE │ RUN │
├───┼─────────┼────────┼───────────┼───────────┼─────┴─────┤
│ 1 │ 1.1.1.1 │ Pod 1A │ NS 1A     │ C 1       │     Y     │
├───┼─────────┼────────┼───────────┼───────────┼─────┬─────┤
│ 2 │ 1.1.1.1 │ Pod 1A │ NS 1A     │ C 2       │  Y  │  N  │
├───┼─────────┼────────┼───────────┼───────────┼─────┼─────┤
│ 3 │ 1.1.1.1 │ Pod 1A │ NS 1B     │ C 3       │  N  │  N  │
├───┼─────────┼────────┼───────────┼───────────┼─────┴─────┤
│ 4 │ 1.1.1.1 │ Pod 1B │ NS 2      │ C 4       │     N     │
├───┼─────────┼────────┼───────────┼───────────┼─────┬─────┤
│ 5 │ 1.1.1.1 │ Pod 1B │ NS 2      │ C 5       │  Y  │  N  │
├───┼─────────┼────────┼───────────┼───────────┼─────┴─────┤
│ 6 │ 2.2.2.2 │ Pod 2  │ NS 3      │ C 6       │     Y     │
├───┼─────────┼────────┼───────────┼───────────┼───────────┤
│ 7 │ 2.2.2.2 │ Pod 2  │ NS 3      │ C 7       │     Y     │
├───┼─────────┼────────┼───────────┼───────────┼─────┬─────┤
│   │         │        │           │ 7         │  5  │  3  │
└───┴─────────┴────────┴───────────┴───────────┴─────┴─────┘`
	assert.Equal(t, expectedOut, tw.Render())
}

func TestTable_Render_AutoMerge_WithHiddenRows(t *testing.T) {
	tw := NewWriter()
	tw.AppendHeader(Row{"Node IP", "Pods", "Namespace", "Container", "RCE\nEXE", "RCE\nRUN"})
	tw.AppendRow(Row{"1.1.1.1", "Pod 1A", "NS 1A", "C 1", "Y", "Y"})
	tw.AppendRow(Row{"1.1.1.1", "Pod 1A", "NS 1A", "C 2", "Y", "N"})
	tw.AppendRow(Row{"1.1.1.1", "Pod 1A", "NS 1B", "C 3", "N", "N"})
	tw.AppendRow(Row{"1.1.1.1", "Pod 1B", "NS 2", "C 4", "Y", "Y"})
	tw.AppendRow(Row{"1.1.1.1", "Pod 1B", "NS 2", "C 5", "Y", "N"})
	tw.AppendRow(Row{"2.2.2.2", "Pod 2", "NS 3", "C 6", "Y", "Y"})
	tw.AppendRow(Row{"2.2.2.2", "Pod 2", "NS 3", "C 7", "Y", "N"})
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

	expectedOut := `┌─────────┬────────┬───────────┐
│ NODE IP │ PODS   │ NAMESPACE │
├─────────┼────────┼───────────┤
│ 1.1.1.1 │ Pod 1A │ NS 1A     │
│         │        │           │
│         │        │           │
│         │        ├───────────┤
│         │        │ NS 1B     │
│         ├────────┼───────────┤
│         │ Pod 1B │ NS 2      │
│         │        │           │
│         │        │           │
├─────────┼────────┼───────────┤
│ 2.2.2.2 │ Pod 2  │ NS 3      │
│         │        │           │
│         │        │           │
├─────────┼────────┼───────────┤
│         │        │           │
└─────────┴────────┴───────────┘`
	assert.Equal(t, expectedOut, tw.Render())
}

func TestTable_Render_BorderAndSeparators(t *testing.T) {
	table := Table{}
	table.AppendHeader(testHeader)
	table.AppendRows(testRows)
	table.AppendFooter(testFooter)
	expectedOut := `+-----+------------+-----------+--------+-----------------------------+
|   # | FIRST NAME | LAST NAME | SALARY |                             |
+-----+------------+-----------+--------+-----------------------------+
|   1 | Arya       | Stark     |   3000 |                             |
|  20 | Jon        | Snow      |   2000 | You know nothing, Jon Snow! |
| 300 | Tyrion     | Lannister |   5000 |                             |
+-----+------------+-----------+--------+-----------------------------+
|     |            | TOTAL     |  10000 |                             |
+-----+------------+-----------+--------+-----------------------------+`
	assert.Equal(t, expectedOut, table.Render())

	table.Style().Options = OptionsNoBorders
	expectedOut = `   # | FIRST NAME | LAST NAME | SALARY |                             
-----+------------+-----------+--------+-----------------------------
   1 | Arya       | Stark     |   3000 |                             
  20 | Jon        | Snow      |   2000 | You know nothing, Jon Snow! 
 300 | Tyrion     | Lannister |   5000 |                             
-----+------------+-----------+--------+-----------------------------
     |            | TOTAL     |  10000 |                             `
	assert.Equal(t, expectedOut, table.Render())

	table.Style().Options.SeparateColumns = false
	expectedOut = `   #  FIRST NAME  LAST NAME  SALARY                              
-----------------------------------------------------------------
   1  Arya        Stark        3000                              
  20  Jon         Snow         2000  You know nothing, Jon Snow! 
 300  Tyrion      Lannister    5000                              
-----------------------------------------------------------------
                  TOTAL       10000                              `
	assert.Equal(t, expectedOut, table.Render())

	table.Style().Options.SeparateFooter = false
	expectedOut = `   #  FIRST NAME  LAST NAME  SALARY                              
-----------------------------------------------------------------
   1  Arya        Stark        3000                              
  20  Jon         Snow         2000  You know nothing, Jon Snow! 
 300  Tyrion      Lannister    5000                              
                  TOTAL       10000                              `
	assert.Equal(t, expectedOut, table.Render())

	table.Style().Options = OptionsNoBordersAndSeparators
	expectedOut = `   #  FIRST NAME  LAST NAME  SALARY                              
   1  Arya        Stark        3000                              
  20  Jon         Snow         2000  You know nothing, Jon Snow! 
 300  Tyrion      Lannister    5000                              
                  TOTAL       10000                              `
	assert.Equal(t, expectedOut, table.Render())

	table.Style().Options.DrawBorder = true
	expectedOut = `+-----------------------------------------------------------------+
|   #  FIRST NAME  LAST NAME  SALARY                              |
|   1  Arya        Stark        3000                              |
|  20  Jon         Snow         2000  You know nothing, Jon Snow! |
| 300  Tyrion      Lannister    5000                              |
|                  TOTAL       10000                              |
+-----------------------------------------------------------------+`
	assert.Equal(t, expectedOut, table.Render())

	table.Style().Options.SeparateFooter = true
	expectedOut = `+-----------------------------------------------------------------+
|   #  FIRST NAME  LAST NAME  SALARY                              |
|   1  Arya        Stark        3000                              |
|  20  Jon         Snow         2000  You know nothing, Jon Snow! |
| 300  Tyrion      Lannister    5000                              |
+-----------------------------------------------------------------+
|                  TOTAL       10000                              |
+-----------------------------------------------------------------+`
	assert.Equal(t, expectedOut, table.Render())

	table.Style().Options.SeparateHeader = true
	expectedOut = `+-----------------------------------------------------------------+
|   #  FIRST NAME  LAST NAME  SALARY                              |
+-----------------------------------------------------------------+
|   1  Arya        Stark        3000                              |
|  20  Jon         Snow         2000  You know nothing, Jon Snow! |
| 300  Tyrion      Lannister    5000                              |
+-----------------------------------------------------------------+
|                  TOTAL       10000                              |
+-----------------------------------------------------------------+`
	assert.Equal(t, expectedOut, table.Render())

	table.Style().Options.SeparateRows = true
	expectedOut = `+-----------------------------------------------------------------+
|   #  FIRST NAME  LAST NAME  SALARY                              |
+-----------------------------------------------------------------+
|   1  Arya        Stark        3000                              |
+-----------------------------------------------------------------+
|  20  Jon         Snow         2000  You know nothing, Jon Snow! |
+-----------------------------------------------------------------+
| 300  Tyrion      Lannister    5000                              |
+-----------------------------------------------------------------+
|                  TOTAL       10000                              |
+-----------------------------------------------------------------+`
	assert.Equal(t, expectedOut, table.Render())

	table.Style().Options.SeparateColumns = true
	expectedOut = `+-----+------------+-----------+--------+-----------------------------+
|   # | FIRST NAME | LAST NAME | SALARY |                             |
+-----+------------+-----------+--------+-----------------------------+
|   1 | Arya       | Stark     |   3000 |                             |
+-----+------------+-----------+--------+-----------------------------+
|  20 | Jon        | Snow      |   2000 | You know nothing, Jon Snow! |
+-----+------------+-----------+--------+-----------------------------+
| 300 | Tyrion     | Lannister |   5000 |                             |
+-----+------------+-----------+--------+-----------------------------+
|     |            | TOTAL     |  10000 |                             |
+-----+------------+-----------+--------+-----------------------------+`
	assert.Equal(t, expectedOut, table.Render())
}

func TestTable_Render_Colored(t *testing.T) {
	tw := NewWriter()
	tw.AppendHeader(testHeader)
	tw.AppendRows(testRows)
	tw.AppendRow(testRowMultiLine)
	tw.AppendFooter(testFooter)
	tw.SetAutoIndex(true)
	tw.SetStyle(StyleColoredBright)
	tw.Style().Options.DrawBorder = true
	tw.Style().Options.SeparateColumns = true
	tw.Style().Options.SeparateFooter = true
	tw.Style().Options.SeparateHeader = true
	tw.Style().Options.SeparateRows = true

	expectedOut := []string{
		"\x1b[106;30m+\x1b[0m\x1b[106;30m---\x1b[0m\x1b[106;30m+\x1b[0m\x1b[106;30m-----\x1b[0m\x1b[106;30m+\x1b[0m\x1b[106;30m------------\x1b[0m\x1b[106;30m+\x1b[0m\x1b[106;30m-----------\x1b[0m\x1b[106;30m+\x1b[0m\x1b[106;30m--------\x1b[0m\x1b[106;30m+\x1b[0m\x1b[106;30m-----------------------------\x1b[0m\x1b[106;30m+\x1b[0m",
		"\x1b[106;30m|\x1b[0m\x1b[106;30m   \x1b[0m\x1b[106;30m|\x1b[0m\x1b[106;30m   # \x1b[0m\x1b[106;30m|\x1b[0m\x1b[106;30m FIRST NAME \x1b[0m\x1b[106;30m|\x1b[0m\x1b[106;30m LAST NAME \x1b[0m\x1b[106;30m|\x1b[0m\x1b[106;30m SALARY \x1b[0m\x1b[106;30m|\x1b[0m\x1b[106;30m                             \x1b[0m\x1b[106;30m|\x1b[0m",
		"\x1b[106;30m+\x1b[0m\x1b[106;30m---\x1b[0m\x1b[106;30m+\x1b[0m\x1b[106;30m-----\x1b[0m\x1b[106;30m+\x1b[0m\x1b[106;30m------------\x1b[0m\x1b[106;30m+\x1b[0m\x1b[106;30m-----------\x1b[0m\x1b[106;30m+\x1b[0m\x1b[106;30m--------\x1b[0m\x1b[106;30m+\x1b[0m\x1b[106;30m-----------------------------\x1b[0m\x1b[106;30m+\x1b[0m",
		"\x1b[106;30m|\x1b[0m\x1b[106;30m 1 \x1b[0m\x1b[106;30m|\x1b[0m\x1b[107;30m   1 \x1b[0m\x1b[107;30m|\x1b[0m\x1b[107;30m Arya       \x1b[0m\x1b[107;30m|\x1b[0m\x1b[107;30m Stark     \x1b[0m\x1b[107;30m|\x1b[0m\x1b[107;30m   3000 \x1b[0m\x1b[107;30m|\x1b[0m\x1b[107;30m                             \x1b[0m\x1b[106;30m|\x1b[0m",
		"\x1b[106;30m+\x1b[0m\x1b[106;30m---\x1b[0m\x1b[106;30m+\x1b[0m\x1b[107;30m-----\x1b[0m\x1b[107;30m+\x1b[0m\x1b[107;30m------------\x1b[0m\x1b[107;30m+\x1b[0m\x1b[107;30m-----------\x1b[0m\x1b[107;30m+\x1b[0m\x1b[107;30m--------\x1b[0m\x1b[107;30m+\x1b[0m\x1b[107;30m-----------------------------\x1b[0m\x1b[106;30m+\x1b[0m",
		"\x1b[106;30m|\x1b[0m\x1b[106;30m 2 \x1b[0m\x1b[106;30m|\x1b[0m\x1b[47;30m  20 \x1b[0m\x1b[47;30m|\x1b[0m\x1b[47;30m Jon        \x1b[0m\x1b[47;30m|\x1b[0m\x1b[47;30m Snow      \x1b[0m\x1b[47;30m|\x1b[0m\x1b[47;30m   2000 \x1b[0m\x1b[47;30m|\x1b[0m\x1b[47;30m You know nothing, Jon Snow! \x1b[0m\x1b[106;30m|\x1b[0m",
		"\x1b[106;30m+\x1b[0m\x1b[106;30m---\x1b[0m\x1b[106;30m+\x1b[0m\x1b[47;30m-----\x1b[0m\x1b[47;30m+\x1b[0m\x1b[47;30m------------\x1b[0m\x1b[47;30m+\x1b[0m\x1b[47;30m-----------\x1b[0m\x1b[47;30m+\x1b[0m\x1b[47;30m--------\x1b[0m\x1b[47;30m+\x1b[0m\x1b[47;30m-----------------------------\x1b[0m\x1b[106;30m+\x1b[0m",
		"\x1b[106;30m|\x1b[0m\x1b[106;30m 3 \x1b[0m\x1b[106;30m|\x1b[0m\x1b[107;30m 300 \x1b[0m\x1b[107;30m|\x1b[0m\x1b[107;30m Tyrion     \x1b[0m\x1b[107;30m|\x1b[0m\x1b[107;30m Lannister \x1b[0m\x1b[107;30m|\x1b[0m\x1b[107;30m   5000 \x1b[0m\x1b[107;30m|\x1b[0m\x1b[107;30m                             \x1b[0m\x1b[106;30m|\x1b[0m",
		"\x1b[106;30m+\x1b[0m\x1b[106;30m---\x1b[0m\x1b[106;30m+\x1b[0m\x1b[107;30m-----\x1b[0m\x1b[107;30m+\x1b[0m\x1b[107;30m------------\x1b[0m\x1b[107;30m+\x1b[0m\x1b[107;30m-----------\x1b[0m\x1b[107;30m+\x1b[0m\x1b[107;30m--------\x1b[0m\x1b[107;30m+\x1b[0m\x1b[107;30m-----------------------------\x1b[0m\x1b[106;30m+\x1b[0m",
		"\x1b[106;30m|\x1b[0m\x1b[106;30m 4 \x1b[0m\x1b[106;30m|\x1b[0m\x1b[47;30m   0 \x1b[0m\x1b[47;30m|\x1b[0m\x1b[47;30m Winter     \x1b[0m\x1b[47;30m|\x1b[0m\x1b[47;30m Is        \x1b[0m\x1b[47;30m|\x1b[0m\x1b[47;30m      0 \x1b[0m\x1b[47;30m|\x1b[0m\x1b[47;30m Coming.                     \x1b[0m\x1b[106;30m|\x1b[0m",
		"\x1b[106;30m|\x1b[0m\x1b[106;30m   \x1b[0m\x1b[106;30m|\x1b[0m\x1b[47;30m     \x1b[0m\x1b[47;30m|\x1b[0m\x1b[47;30m            \x1b[0m\x1b[47;30m|\x1b[0m\x1b[47;30m           \x1b[0m\x1b[47;30m|\x1b[0m\x1b[47;30m        \x1b[0m\x1b[47;30m|\x1b[0m\x1b[47;30m The North Remembers!        \x1b[0m\x1b[106;30m|\x1b[0m",
		"\x1b[106;30m|\x1b[0m\x1b[106;30m   \x1b[0m\x1b[106;30m|\x1b[0m\x1b[47;30m     \x1b[0m\x1b[47;30m|\x1b[0m\x1b[47;30m            \x1b[0m\x1b[47;30m|\x1b[0m\x1b[47;30m           \x1b[0m\x1b[47;30m|\x1b[0m\x1b[47;30m        \x1b[0m\x1b[47;30m|\x1b[0m\x1b[47;30m This is known.              \x1b[0m\x1b[106;30m|\x1b[0m",
		"\x1b[46;30m+\x1b[0m\x1b[46;30m---\x1b[0m\x1b[46;30m+\x1b[0m\x1b[46;30m-----\x1b[0m\x1b[46;30m+\x1b[0m\x1b[46;30m------------\x1b[0m\x1b[46;30m+\x1b[0m\x1b[46;30m-----------\x1b[0m\x1b[46;30m+\x1b[0m\x1b[46;30m--------\x1b[0m\x1b[46;30m+\x1b[0m\x1b[46;30m-----------------------------\x1b[0m\x1b[46;30m+\x1b[0m",
		"\x1b[46;30m|\x1b[0m\x1b[46;30m   \x1b[0m\x1b[46;30m|\x1b[0m\x1b[46;30m     \x1b[0m\x1b[46;30m|\x1b[0m\x1b[46;30m            \x1b[0m\x1b[46;30m|\x1b[0m\x1b[46;30m TOTAL     \x1b[0m\x1b[46;30m|\x1b[0m\x1b[46;30m  10000 \x1b[0m\x1b[46;30m|\x1b[0m\x1b[46;30m                             \x1b[0m\x1b[46;30m|\x1b[0m",
		"\x1b[46;30m+\x1b[0m\x1b[46;30m---\x1b[0m\x1b[46;30m+\x1b[0m\x1b[46;30m-----\x1b[0m\x1b[46;30m+\x1b[0m\x1b[46;30m------------\x1b[0m\x1b[46;30m+\x1b[0m\x1b[46;30m-----------\x1b[0m\x1b[46;30m+\x1b[0m\x1b[46;30m--------\x1b[0m\x1b[46;30m+\x1b[0m\x1b[46;30m-----------------------------\x1b[0m\x1b[46;30m+\x1b[0m",
	}
	out := tw.Render()
	assert.Equal(t, strings.Join(expectedOut, "\n"), out)
	if strings.Join(expectedOut, "\n") != out {
		for _, line := range strings.Split(out, "\n") {
			fmt.Printf("%#v,\n", line)
		}
	}
}

func TestTable_Render_ColoredCustom(t *testing.T) {
	tw := NewWriter()
	tw.AppendHeader(testHeader)
	tw.AppendRows(testRows)
	tw.AppendRow(testRowMultiLine)
	tw.AppendFooter(testFooter)
	tw.SetCaption(testCaption)
	tw.SetColumnConfigs([]ColumnConfig{
		{
			Name:         "#",
			Colors:       testColor,
			ColorsHeader: testColorHiRedBold,
		}, {
			Name:         "First Name",
			Colors:       testColor,
			ColorsHeader: testColorHiRedBold,
		}, {
			Name:         "Last Name",
			Colors:       testColor,
			ColorsHeader: testColorHiRedBold,
			ColorsFooter: testColorHiBlueBold,
		}, {
			Name:         "Salary",
			Colors:       testColor,
			ColorsHeader: testColorHiRedBold,
			ColorsFooter: testColorHiBlueBold,
		}, {
			Number: 5,
			Colors: text.Colors{text.FgCyan},
		},
	})
	tw.SetStyle(StyleRounded)

	expectedOut := []string{
		"╭─────┬────────────┬───────────┬────────┬─────────────────────────────╮",
		"│\x1b[91;1m   # \x1b[0m│\x1b[91;1m FIRST NAME \x1b[0m│\x1b[91;1m LAST NAME \x1b[0m│\x1b[91;1m SALARY \x1b[0m│                             │",
		"├─────┼────────────┼───────────┼────────┼─────────────────────────────┤",
		"│\x1b[32m   1 \x1b[0m│\x1b[32m Arya       \x1b[0m│\x1b[32m Stark     \x1b[0m│\x1b[32m   3000 \x1b[0m│\x1b[36m                             \x1b[0m│",
		"│\x1b[32m  20 \x1b[0m│\x1b[32m Jon        \x1b[0m│\x1b[32m Snow      \x1b[0m│\x1b[32m   2000 \x1b[0m│\x1b[36m You know nothing, Jon Snow! \x1b[0m│",
		"│\x1b[32m 300 \x1b[0m│\x1b[32m Tyrion     \x1b[0m│\x1b[32m Lannister \x1b[0m│\x1b[32m   5000 \x1b[0m│\x1b[36m                             \x1b[0m│",
		"│\x1b[32m   0 \x1b[0m│\x1b[32m Winter     \x1b[0m│\x1b[32m Is        \x1b[0m│\x1b[32m      0 \x1b[0m│\x1b[36m Coming.                     \x1b[0m│",
		"│\x1b[32m     \x1b[0m│\x1b[32m            \x1b[0m│\x1b[32m           \x1b[0m│\x1b[32m        \x1b[0m│\x1b[36m The North Remembers!        \x1b[0m│",
		"│\x1b[32m     \x1b[0m│\x1b[32m            \x1b[0m│\x1b[32m           \x1b[0m│\x1b[32m        \x1b[0m│\x1b[36m This is known.              \x1b[0m│",
		"├─────┼────────────┼───────────┼────────┼─────────────────────────────┤",
		"│     │            │\x1b[94;1m TOTAL     \x1b[0m│\x1b[94;1m  10000 \x1b[0m│                             │",
		"╰─────┴────────────┴───────────┴────────┴─────────────────────────────╯",
		"A Song of Ice and Fire",
	}
	assert.Equal(t, strings.Join(expectedOut, "\n"), tw.Render())
}

func TestTable_Render_ColoredTableWithinATable(t *testing.T) {
	table := Table{}
	table.AppendHeader(testHeader)
	table.AppendRows(testRows)
	table.AppendFooter(testFooter)
	table.SetStyle(StyleColoredBright)
	table.SetIndexColumn(1)

	// colored is simple; render the colored table into another table
	tableOuter := Table{}
	tableOuter.AppendRow(Row{table.Render()})
	tableOuter.SetStyle(StyleRounded)

	expectedOut := strings.Join([]string{
		"╭───────────────────────────────────────────────────────────────────╮",
		"│ \x1b[106;30m   # \x1b[0m\x1b[106;30m FIRST NAME \x1b[0m\x1b[106;30m LAST NAME \x1b[0m\x1b[106;30m SALARY \x1b[0m\x1b[106;30m                             \x1b[0m │",
		"│ \x1b[106;30m   1 \x1b[0m\x1b[107;30m Arya       \x1b[0m\x1b[107;30m Stark     \x1b[0m\x1b[107;30m   3000 \x1b[0m\x1b[107;30m                             \x1b[0m │",
		"│ \x1b[106;30m  20 \x1b[0m\x1b[47;30m Jon        \x1b[0m\x1b[47;30m Snow      \x1b[0m\x1b[47;30m   2000 \x1b[0m\x1b[47;30m You know nothing, Jon Snow! \x1b[0m │",
		"│ \x1b[106;30m 300 \x1b[0m\x1b[107;30m Tyrion     \x1b[0m\x1b[107;30m Lannister \x1b[0m\x1b[107;30m   5000 \x1b[0m\x1b[107;30m                             \x1b[0m │",
		"│ \x1b[46;30m     \x1b[0m\x1b[46;30m            \x1b[0m\x1b[46;30m TOTAL     \x1b[0m\x1b[46;30m  10000 \x1b[0m\x1b[46;30m                             \x1b[0m │",
		"╰───────────────────────────────────────────────────────────────────╯",
	}, "\n")
	out := tableOuter.Render()
	assert.Equal(t, expectedOut, out)

	// dump it out in a easy way to update the test if things are meant to
	// change due to some other feature
	if expectedOut != out {
		for _, line := range strings.Split(out, "\n") {
			fmt.Printf("%#v,\n", line)
		}
		fmt.Println()
	}
}

func TestTable_Render_ColoredTableWithinAColoredTable(t *testing.T) {
	table := Table{}
	table.AppendHeader(testHeader)
	table.AppendRows(testRows)
	table.AppendFooter(testFooter)
	table.SetStyle(StyleColoredBright)
	table.SetIndexColumn(1)

	// colored is simple; render the colored table into another colored table
	tableOuter := Table{}
	tableOuter.AppendHeader(Row{"Colored Table within a Colored Table"})
	tableOuter.AppendRow(Row{"\n" + table.Render() + "\n"})
	tableOuter.SetColumnConfigs([]ColumnConfig{{Number: 1, AlignHeader: text.AlignCenter}})
	tableOuter.SetStyle(StyleColoredBright)

	expectedOut := strings.Join([]string{
		"\x1b[106;30m                COLORED TABLE WITHIN A COLORED TABLE               \x1b[0m",
		"\x1b[107;30m                                                                   \x1b[0m",
		"\x1b[107;30m \x1b[106;30m   # \x1b[0m\x1b[107;30m\x1b[106;30m FIRST NAME \x1b[0m\x1b[107;30m\x1b[106;30m LAST NAME \x1b[0m\x1b[107;30m\x1b[106;30m SALARY \x1b[0m\x1b[107;30m\x1b[106;30m                             \x1b[0m\x1b[107;30m \x1b[0m",
		"\x1b[107;30m \x1b[106;30m   1 \x1b[0m\x1b[107;30m\x1b[107;30m Arya       \x1b[0m\x1b[107;30m\x1b[107;30m Stark     \x1b[0m\x1b[107;30m\x1b[107;30m   3000 \x1b[0m\x1b[107;30m\x1b[107;30m                             \x1b[0m\x1b[107;30m \x1b[0m",
		"\x1b[107;30m \x1b[106;30m  20 \x1b[0m\x1b[107;30m\x1b[47;30m Jon        \x1b[0m\x1b[107;30m\x1b[47;30m Snow      \x1b[0m\x1b[107;30m\x1b[47;30m   2000 \x1b[0m\x1b[107;30m\x1b[47;30m You know nothing, Jon Snow! \x1b[0m\x1b[107;30m \x1b[0m",
		"\x1b[107;30m \x1b[106;30m 300 \x1b[0m\x1b[107;30m\x1b[107;30m Tyrion     \x1b[0m\x1b[107;30m\x1b[107;30m Lannister \x1b[0m\x1b[107;30m\x1b[107;30m   5000 \x1b[0m\x1b[107;30m\x1b[107;30m                             \x1b[0m\x1b[107;30m \x1b[0m",
		"\x1b[107;30m \x1b[46;30m     \x1b[0m\x1b[107;30m\x1b[46;30m            \x1b[0m\x1b[107;30m\x1b[46;30m TOTAL     \x1b[0m\x1b[107;30m\x1b[46;30m  10000 \x1b[0m\x1b[107;30m\x1b[46;30m                             \x1b[0m\x1b[107;30m \x1b[0m",
		"\x1b[107;30m                                                                   \x1b[0m",
	}, "\n")
	out := tableOuter.Render()
	assert.Equal(t, expectedOut, out)

	// dump it out in a easy way to update the test if things are meant to
	// change due to some other feature
	if expectedOut != out {
		for _, line := range strings.Split(out, "\n") {
			fmt.Printf("%#v,\n", line)
		}
		fmt.Println()
	}
}

func TestTable_Render_ColoredStyleAutoIndex(t *testing.T) {
	table := Table{}
	table.AppendHeader(testHeader)
	table.AppendRows(testRows)
	table.AppendFooter(testFooter)
	table.SetAutoIndex(true)
	table.SetStyle(StyleColoredDark)
	table.SetTitle(testTitle2)

	expectedOut := strings.Join([]string{
		"\x1b[106;30;1m When you play the Game of Thrones, you win or you die. There is no \x1b[0m",
		"\x1b[106;30;1m middle ground.                                                     \x1b[0m",
		"\x1b[96;100m   \x1b[0m\x1b[96;100m   # \x1b[0m\x1b[96;100m FIRST NAME \x1b[0m\x1b[96;100m LAST NAME \x1b[0m\x1b[96;100m SALARY \x1b[0m\x1b[96;100m                             \x1b[0m",
		"\x1b[96;100m 1 \x1b[0m\x1b[97;40m   1 \x1b[0m\x1b[97;40m Arya       \x1b[0m\x1b[97;40m Stark     \x1b[0m\x1b[97;40m   3000 \x1b[0m\x1b[97;40m                             \x1b[0m",
		"\x1b[96;100m 2 \x1b[0m\x1b[37;40m  20 \x1b[0m\x1b[37;40m Jon        \x1b[0m\x1b[37;40m Snow      \x1b[0m\x1b[37;40m   2000 \x1b[0m\x1b[37;40m You know nothing, Jon Snow! \x1b[0m",
		"\x1b[96;100m 3 \x1b[0m\x1b[97;40m 300 \x1b[0m\x1b[97;40m Tyrion     \x1b[0m\x1b[97;40m Lannister \x1b[0m\x1b[97;40m   5000 \x1b[0m\x1b[97;40m                             \x1b[0m",
		"\x1b[36;100m   \x1b[0m\x1b[36;100m     \x1b[0m\x1b[36;100m            \x1b[0m\x1b[36;100m TOTAL     \x1b[0m\x1b[36;100m  10000 \x1b[0m\x1b[36;100m                             \x1b[0m",
	}, "\n")
	out := table.Render()
	assert.Equal(t, expectedOut, out)

	// dump it out in a easy way to update the test if things are meant to
	// change due to some other feature
	if expectedOut != out {
		for _, line := range strings.Split(out, "\n") {
			fmt.Printf("%#v,\n", line)
		}
		fmt.Println()
	}
}

func TestTable_Render_ColumnConfigs(t *testing.T) {
	generatePrefixTransformer := func(prefix string) text.Transformer {
		return func(val interface{}) string {
			return fmt.Sprintf("%s%v", prefix, val)
		}
	}
	generateSuffixTransformer := func(suffix string) text.Transformer {
		return func(val interface{}) string {
			return fmt.Sprintf("%v%s", val, suffix)
		}
	}
	salaryTransformer := text.Transformer(func(val interface{}) string {
		if valInt, ok := val.(int); ok {
			return fmt.Sprintf("$ %.2f", float64(valInt)+0.03)
		}
		return strings.Replace(fmt.Sprint(val), "ry", "riii", -1)
	})

	tw := NewWriter()
	tw.AppendHeader(testHeaderMultiLine)
	tw.AppendRows(testRows)
	tw.AppendRow(testRowMultiLine)
	tw.AppendFooter(testFooterMultiLine)
	tw.SetAutoIndex(true)
	tw.SetColumnConfigs([]ColumnConfig{
		{
			Name:              fmt.Sprint(testHeaderMultiLine[1]), // First Name
			Align:             text.AlignRight,
			AlignFooter:       text.AlignRight,
			AlignHeader:       text.AlignRight,
			Colors:            text.Colors{text.BgBlack, text.FgRed},
			ColorsHeader:      text.Colors{text.BgRed, text.FgBlack, text.Bold},
			ColorsFooter:      text.Colors{text.BgRed, text.FgBlack},
			Transformer:       generatePrefixTransformer("(r_"),
			TransformerFooter: generatePrefixTransformer("(f_"),
			TransformerHeader: generatePrefixTransformer("(h_"),
			VAlign:            text.VAlignTop,
			VAlignFooter:      text.VAlignTop,
			VAlignHeader:      text.VAlignTop,
			WidthMax:          10,
		}, {
			Name:              fmt.Sprint(testHeaderMultiLine[2]), // Last Name
			Align:             text.AlignLeft,
			AlignFooter:       text.AlignLeft,
			AlignHeader:       text.AlignLeft,
			Colors:            text.Colors{text.BgBlack, text.FgGreen},
			ColorsHeader:      text.Colors{text.BgGreen, text.FgBlack, text.Bold},
			ColorsFooter:      text.Colors{text.BgGreen, text.FgBlack},
			Transformer:       generateSuffixTransformer("_r)"),
			TransformerFooter: generateSuffixTransformer("_f)"),
			TransformerHeader: generateSuffixTransformer("_h)"),
			VAlign:            text.VAlignMiddle,
			VAlignFooter:      text.VAlignMiddle,
			VAlignHeader:      text.VAlignMiddle,
			WidthMax:          10,
		}, {
			Number:            4, // Salary
			Colors:            text.Colors{text.BgBlack, text.FgBlue},
			ColorsHeader:      text.Colors{text.BgBlue, text.FgBlack, text.Bold},
			ColorsFooter:      text.Colors{text.BgBlue, text.FgBlack},
			Transformer:       salaryTransformer,
			TransformerFooter: salaryTransformer,
			TransformerHeader: salaryTransformer,
			VAlign:            text.VAlignBottom,
			VAlignFooter:      text.VAlignBottom,
			VAlignHeader:      text.VAlignBottom,
			WidthMin:          16,
		}, {
			Name:   "Non-existent Column",
			Colors: text.Colors{text.BgYellow, text.FgHiRed},
		},
	})
	tw.SetStyle(styleTest)

	expectedOutLines := []string{
		"(---^-----^-----------^------------^------------------^-----------------------------)",
		"[< >|<  #>|\x1b[41;30;1m< (H_FIRST>\x1b[0m|\x1b[42;30;1m<LAST      >\x1b[0m|\x1b[44;30;1m<                >\x1b[0m|<                           >]",
		"[< >|<   >|\x1b[41;30;1m<     NAME>\x1b[0m|\x1b[42;30;1m<NAME_H)   >\x1b[0m|\x1b[44;30;1m<        SALARIII>\x1b[0m|<                           >]",
		"{---+-----+-----------+------------+------------------+-----------------------------}",
		"[<1>|<  1>|\x1b[40;31m<  (r_Arya>\x1b[0m|\x1b[40;32m<Stark_r)  >\x1b[0m|\x1b[40;34m<       $ 3000.03>\x1b[0m|<                           >]",
		"[<2>|< 20>|\x1b[40;31m<   (r_Jon>\x1b[0m|\x1b[40;32m<Snow_r)   >\x1b[0m|\x1b[40;34m<       $ 2000.03>\x1b[0m|<You know nothing, Jon Snow!>]",
		"[<3>|<300>|\x1b[40;31m<(r_Tyrion>\x1b[0m|\x1b[40;32m<Lannister_>\x1b[0m|\x1b[40;34m<                >\x1b[0m|<                           >]",
		"[< >|<   >|\x1b[40;31m<         >\x1b[0m|\x1b[40;32m<r)        >\x1b[0m|\x1b[40;34m<       $ 5000.03>\x1b[0m|<                           >]",
		"[<4>|<  0>|\x1b[40;31m<(r_Winter>\x1b[0m|\x1b[40;32m<          >\x1b[0m|\x1b[40;34m<                >\x1b[0m|<Coming.                    >]",
		"[< >|<   >|\x1b[40;31m<         >\x1b[0m|\x1b[40;32m<Is_r)     >\x1b[0m|\x1b[40;34m<                >\x1b[0m|<The North Remembers!       >]",
		"[< >|<   >|\x1b[40;31m<         >\x1b[0m|\x1b[40;32m<          >\x1b[0m|\x1b[40;34m<          $ 0.03>\x1b[0m|<This is known.             >]",
		"{---+-----+-----------+------------+------------------+-----------------------------}",
		"[< >|<   >|\x1b[41;30m<      (F_>\x1b[0m|\x1b[42;30m<TOTAL     >\x1b[0m|\x1b[44;30m<                >\x1b[0m|<                           >]",
		"[< >|<   >|\x1b[41;30m<         >\x1b[0m|\x1b[42;30m<SALARY_F) >\x1b[0m|\x1b[44;30m<      $ 10000.03>\x1b[0m|<                           >]",
		"\\---v-----v-----------v------------v------------------v-----------------------------/",
	}
	expectedOut := strings.Join(expectedOutLines, "\n")
	assert.Equal(t, expectedOut, tw.Render())
}

func TestTable_Render_Empty(t *testing.T) {
	tw := NewWriter()
	assert.Empty(t, tw.Render())
}

func TestTable_Render_HiddenColumns(t *testing.T) {
	tw := NewWriter()
	tw.AppendHeader(testHeader)
	tw.AppendRows(testRows)
	tw.AppendFooter(testFooter)

	// ensure sorting is done before hiding the columns
	tw.SortBy([]SortBy{
		{Name: "Salary", Mode: DscNumeric},
	})

	t.Run("no columns hidden", func(t *testing.T) {
		tw.SetColumnConfigs(generateColumnConfigsWithHiddenColumns(nil))

		expectedOut := `+-----+------------+-------------+--------+-------------------------------+
|   # | FIRST NAME | LAST NAME   | SALARY |                               |
+-----+------------+-------------+--------+-------------------------------+
| 307 | >>Tyrion   | Lannister<< |   5013 |                               |
|   8 | >>Arya     | Stark<<     |   3013 |                               |
|  27 | >>Jon      | Snow<<      |   2013 | ~You know nothing, Jon Snow!~ |
+-----+------------+-------------+--------+-------------------------------+
|     |            | TOTAL       |  10000 |                               |
+-----+------------+-------------+--------+-------------------------------+`
		assert.Equal(t, expectedOut, tw.Render())
	})

	t.Run("every column hidden", func(t *testing.T) {
		tw.SetColumnConfigs(generateColumnConfigsWithHiddenColumns([]int{0, 1, 2, 3, 4}))

		expectedOut := ``
		assert.Equal(t, expectedOut, tw.Render())
	})

	t.Run("some columns hidden (1)", func(t *testing.T) {
		tw.SetColumnConfigs(generateColumnConfigsWithHiddenColumns([]int{1, 2, 3, 4}))

		expectedOut := `+-----+
|   # |
+-----+
| 307 |
|   8 |
|  27 |
+-----+
|     |
+-----+`
		assert.Equal(t, expectedOut, tw.Render())
	})

	t.Run("some columns hidden (2)", func(t *testing.T) {
		tw.SetColumnConfigs(generateColumnConfigsWithHiddenColumns([]int{1, 2, 3}))

		expectedOut := `+-----+-------------------------------+
|   # |                               |
+-----+-------------------------------+
| 307 |                               |
|   8 |                               |
|  27 | ~You know nothing, Jon Snow!~ |
+-----+-------------------------------+
|     |                               |
+-----+-------------------------------+`
		assert.Equal(t, expectedOut, tw.Render())
	})

	t.Run("some columns hidden (3)", func(t *testing.T) {
		tw.SetColumnConfigs(generateColumnConfigsWithHiddenColumns([]int{0, 4}))

		expectedOut := `+------------+-------------+--------+
| FIRST NAME | LAST NAME   | SALARY |
+------------+-------------+--------+
| >>Tyrion   | Lannister<< |   5013 |
| >>Arya     | Stark<<     |   3013 |
| >>Jon      | Snow<<      |   2013 |
+------------+-------------+--------+
|            | TOTAL       |  10000 |
+------------+-------------+--------+`
		assert.Equal(t, expectedOut, tw.Render())
	})

	t.Run("first column hidden", func(t *testing.T) {
		tw.SetColumnConfigs(generateColumnConfigsWithHiddenColumns([]int{0}))

		expectedOut := `+------------+-------------+--------+-------------------------------+
| FIRST NAME | LAST NAME   | SALARY |                               |
+------------+-------------+--------+-------------------------------+
| >>Tyrion   | Lannister<< |   5013 |                               |
| >>Arya     | Stark<<     |   3013 |                               |
| >>Jon      | Snow<<      |   2013 | ~You know nothing, Jon Snow!~ |
+------------+-------------+--------+-------------------------------+
|            | TOTAL       |  10000 |                               |
+------------+-------------+--------+-------------------------------+`
		assert.Equal(t, expectedOut, tw.Render())
	})

	t.Run("column hidden in the middle", func(t *testing.T) {
		tw.SetColumnConfigs(generateColumnConfigsWithHiddenColumns([]int{1}))

		expectedOut := `+-----+-------------+--------+-------------------------------+
|   # | LAST NAME   | SALARY |                               |
+-----+-------------+--------+-------------------------------+
| 307 | Lannister<< |   5013 |                               |
|   8 | Stark<<     |   3013 |                               |
|  27 | Snow<<      |   2013 | ~You know nothing, Jon Snow!~ |
+-----+-------------+--------+-------------------------------+
|     | TOTAL       |  10000 |                               |
+-----+-------------+--------+-------------------------------+`
		assert.Equal(t, expectedOut, tw.Render())
	})

	t.Run("last column hidden", func(t *testing.T) {
		tw.SetColumnConfigs(generateColumnConfigsWithHiddenColumns([]int{4}))

		expectedOut := `+-----+------------+-------------+--------+
|   # | FIRST NAME | LAST NAME   | SALARY |
+-----+------------+-------------+--------+
| 307 | >>Tyrion   | Lannister<< |   5013 |
|   8 | >>Arya     | Stark<<     |   3013 |
|  27 | >>Jon      | Snow<<      |   2013 |
+-----+------------+-------------+--------+
|     |            | TOTAL       |  10000 |
+-----+------------+-------------+--------+`
		assert.Equal(t, expectedOut, tw.Render())
	})
}

func TestTable_Render_Paged(t *testing.T) {
	tw := NewWriter()
	tw.AppendHeader(testHeader)
	tw.AppendRows(testRows)
	tw.AppendRow(testRowMultiLine)
	tw.AppendFooter(Row{"", "", "Total", 10000})
	tw.SetPageSize(1)

	expectedOut := `+-----+------------+-----------+--------+-----------------------------+
|   # | FIRST NAME | LAST NAME | SALARY |                             |
+-----+------------+-----------+--------+-----------------------------+
|   1 | Arya       | Stark     |   3000 |                             |
+-----+------------+-----------+--------+-----------------------------+
|     |            | TOTAL     |  10000 |                             |
+-----+------------+-----------+--------+-----------------------------+

+-----+------------+-----------+--------+-----------------------------+
|   # | FIRST NAME | LAST NAME | SALARY |                             |
+-----+------------+-----------+--------+-----------------------------+
|  20 | Jon        | Snow      |   2000 | You know nothing, Jon Snow! |
+-----+------------+-----------+--------+-----------------------------+
|     |            | TOTAL     |  10000 |                             |
+-----+------------+-----------+--------+-----------------------------+

+-----+------------+-----------+--------+-----------------------------+
|   # | FIRST NAME | LAST NAME | SALARY |                             |
+-----+------------+-----------+--------+-----------------------------+
| 300 | Tyrion     | Lannister |   5000 |                             |
+-----+------------+-----------+--------+-----------------------------+
|     |            | TOTAL     |  10000 |                             |
+-----+------------+-----------+--------+-----------------------------+

+-----+------------+-----------+--------+-----------------------------+
|   # | FIRST NAME | LAST NAME | SALARY |                             |
+-----+------------+-----------+--------+-----------------------------+
|   0 | Winter     | Is        |      0 | Coming.                     |
+-----+------------+-----------+--------+-----------------------------+
|     |            | TOTAL     |  10000 |                             |
+-----+------------+-----------+--------+-----------------------------+

+-----+------------+-----------+--------+-----------------------------+
|   # | FIRST NAME | LAST NAME | SALARY |                             |
+-----+------------+-----------+--------+-----------------------------+
|     |            |           |        | The North Remembers!        |
+-----+------------+-----------+--------+-----------------------------+
|     |            | TOTAL     |  10000 |                             |
+-----+------------+-----------+--------+-----------------------------+

+-----+------------+-----------+--------+-----------------------------+
|   # | FIRST NAME | LAST NAME | SALARY |                             |
+-----+------------+-----------+--------+-----------------------------+
|     |            |           |        | This is known.              |
+-----+------------+-----------+--------+-----------------------------+
|     |            | TOTAL     |  10000 |                             |
+-----+------------+-----------+--------+-----------------------------+`
	assert.Equal(t, expectedOut, tw.Render())
}

func TestTable_Render_Reset(t *testing.T) {
	tw := NewWriter()
	tw.AppendHeader(testHeader)
	tw.AppendRows(testRows)
	tw.AppendFooter(testFooter)
	tw.SetStyle(StyleLight)

	expectedOut := `┌─────┬────────────┬───────────┬────────┬─────────────────────────────┐
│   # │ FIRST NAME │ LAST NAME │ SALARY │                             │
├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
│   1 │ Arya       │ Stark     │   3000 │                             │
│  20 │ Jon        │ Snow      │   2000 │ You know nothing, Jon Snow! │
│ 300 │ Tyrion     │ Lannister │   5000 │                             │
├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
│     │            │ TOTAL     │  10000 │                             │
└─────┴────────────┴───────────┴────────┴─────────────────────────────┘`
	assert.Equal(t, expectedOut, tw.Render())

	tw.ResetFooters()
	expectedOut = `┌─────┬────────────┬───────────┬────────┬─────────────────────────────┐
│   # │ FIRST NAME │ LAST NAME │ SALARY │                             │
├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
│   1 │ Arya       │ Stark     │   3000 │                             │
│  20 │ Jon        │ Snow      │   2000 │ You know nothing, Jon Snow! │
│ 300 │ Tyrion     │ Lannister │   5000 │                             │
└─────┴────────────┴───────────┴────────┴─────────────────────────────┘`
	assert.Equal(t, expectedOut, tw.Render())

	tw.ResetHeaders()
	expectedOut = `┌─────┬────────┬───────────┬──────┬─────────────────────────────┐
│   1 │ Arya   │ Stark     │ 3000 │                             │
│  20 │ Jon    │ Snow      │ 2000 │ You know nothing, Jon Snow! │
│ 300 │ Tyrion │ Lannister │ 5000 │                             │
└─────┴────────┴───────────┴──────┴─────────────────────────────┘`
	assert.Equal(t, expectedOut, tw.Render())

	tw.ResetRows()
	assert.Empty(t, tw.Render())
}

func TestTable_Render_RowPainter(t *testing.T) {
	tw := NewWriter()
	tw.AppendHeader(testHeader)
	tw.AppendRows(testRows)
	tw.AppendRow(testRowMultiLine)
	tw.AppendFooter(testFooter)
	tw.SetIndexColumn(1)
	tw.SetRowPainter(RowPainter(func(row Row) text.Colors {
		if salary, ok := row[3].(int); ok {
			if salary > 3000 {
				return text.Colors{text.BgYellow, text.FgBlack}
			} else if salary < 2000 {
				return text.Colors{text.BgRed, text.FgBlack}
			}
		}
		return nil
	}))
	tw.SetStyle(StyleLight)
	tw.SortBy([]SortBy{{Name: "Salary", Mode: AscNumeric}})

	expectedOutLines := []string{
		"┌─────┬────────────┬───────────┬────────┬─────────────────────────────┐",
		"│   # │ FIRST NAME │ LAST NAME │ SALARY │                             │",
		"├─────┼────────────┼───────────┼────────┼─────────────────────────────┤",
		"│   0 │\x1b[41;30m Winter     \x1b[0m│\x1b[41;30m Is        \x1b[0m│\x1b[41;30m      0 \x1b[0m│\x1b[41;30m Coming.                     \x1b[0m│",
		"│     │\x1b[41;30m            \x1b[0m│\x1b[41;30m           \x1b[0m│\x1b[41;30m        \x1b[0m│\x1b[41;30m The North Remembers!        \x1b[0m│",
		"│     │\x1b[41;30m            \x1b[0m│\x1b[41;30m           \x1b[0m│\x1b[41;30m        \x1b[0m│\x1b[41;30m This is known.              \x1b[0m│",
		"│  20 │ Jon        │ Snow      │   2000 │ You know nothing, Jon Snow! │",
		"│   1 │ Arya       │ Stark     │   3000 │                             │",
		"│ 300 │\x1b[43;30m Tyrion     \x1b[0m│\x1b[43;30m Lannister \x1b[0m│\x1b[43;30m   5000 \x1b[0m│\x1b[43;30m                             \x1b[0m│",
		"├─────┼────────────┼───────────┼────────┼─────────────────────────────┤",
		"│     │            │ TOTAL     │  10000 │                             │",
		"└─────┴────────────┴───────────┴────────┴─────────────────────────────┘",
	}
	expectedOut := strings.Join(expectedOutLines, "\n")
	assert.Equal(t, expectedOut, tw.Render())

	tw.SetStyle(StyleColoredBright)
	tw.Style().Color.RowAlternate = tw.Style().Color.Row
	expectedOutLines = []string{
		"\x1b[106;30m   # \x1b[0m\x1b[106;30m FIRST NAME \x1b[0m\x1b[106;30m LAST NAME \x1b[0m\x1b[106;30m SALARY \x1b[0m\x1b[106;30m                             \x1b[0m",
		"\x1b[106;30m   0 \x1b[0m\x1b[41;30m Winter     \x1b[0m\x1b[41;30m Is        \x1b[0m\x1b[41;30m      0 \x1b[0m\x1b[41;30m Coming.                     \x1b[0m",
		"\x1b[106;30m     \x1b[0m\x1b[41;30m            \x1b[0m\x1b[41;30m           \x1b[0m\x1b[41;30m        \x1b[0m\x1b[41;30m The North Remembers!        \x1b[0m",
		"\x1b[106;30m     \x1b[0m\x1b[41;30m            \x1b[0m\x1b[41;30m           \x1b[0m\x1b[41;30m        \x1b[0m\x1b[41;30m This is known.              \x1b[0m",
		"\x1b[106;30m  20 \x1b[0m\x1b[107;30m Jon        \x1b[0m\x1b[107;30m Snow      \x1b[0m\x1b[107;30m   2000 \x1b[0m\x1b[107;30m You know nothing, Jon Snow! \x1b[0m",
		"\x1b[106;30m   1 \x1b[0m\x1b[107;30m Arya       \x1b[0m\x1b[107;30m Stark     \x1b[0m\x1b[107;30m   3000 \x1b[0m\x1b[107;30m                             \x1b[0m",
		"\x1b[106;30m 300 \x1b[0m\x1b[43;30m Tyrion     \x1b[0m\x1b[43;30m Lannister \x1b[0m\x1b[43;30m   5000 \x1b[0m\x1b[43;30m                             \x1b[0m",
		"\x1b[46;30m     \x1b[0m\x1b[46;30m            \x1b[0m\x1b[46;30m TOTAL     \x1b[0m\x1b[46;30m  10000 \x1b[0m\x1b[46;30m                             \x1b[0m",
	}
	expectedOut = strings.Join(expectedOutLines, "\n")
	assert.Equal(t, expectedOut, tw.Render())
}

func TestTable_Render_Sorted(t *testing.T) {
	tw := NewWriter()
	tw.AppendHeader(testHeader)
	tw.AppendRows(testRows)
	tw.AppendRow(Row{11, "Sansa", "Stark", 6000})
	tw.AppendFooter(testFooter)
	tw.SetStyle(StyleLight)
	tw.SortBy([]SortBy{{Name: "Last Name", Mode: Asc}, {Name: "First Name", Mode: Asc}})

	expectedOut := `┌─────┬────────────┬───────────┬────────┬─────────────────────────────┐
│   # │ FIRST NAME │ LAST NAME │ SALARY │                             │
├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
│ 300 │ Tyrion     │ Lannister │   5000 │                             │
│  20 │ Jon        │ Snow      │   2000 │ You know nothing, Jon Snow! │
│   1 │ Arya       │ Stark     │   3000 │                             │
│  11 │ Sansa      │ Stark     │   6000 │                             │
├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
│     │            │ TOTAL     │  10000 │                             │
└─────┴────────────┴───────────┴────────┴─────────────────────────────┘`
	assert.Equal(t, expectedOut, tw.Render())
}

func TestTable_Render_Separator(t *testing.T) {
	tw := NewWriter()
	tw.AppendHeader(testHeader)
	tw.AppendSeparator() // doesn't make any difference
	tw.AppendRows(testRows)
	tw.AppendSeparator()
	tw.AppendSeparator() // doesn't make any difference
	tw.AppendRow(testRowMultiLine)
	tw.AppendSeparator()
	tw.AppendSeparator() // doesn't make any difference
	tw.AppendSeparator() // doesn't make any difference
	tw.AppendRow(Row{11, "Sansa", "Stark", 6000})
	tw.AppendSeparator() // doesn't make any difference
	tw.AppendSeparator() // doesn't make any difference
	tw.AppendSeparator() // doesn't make any difference
	tw.AppendSeparator() // doesn't make any difference
	tw.AppendFooter(testFooter)
	tw.SetStyle(StyleLight)

	expectedOut := `┌─────┬────────────┬───────────┬────────┬─────────────────────────────┐
│   # │ FIRST NAME │ LAST NAME │ SALARY │                             │
├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
│   1 │ Arya       │ Stark     │   3000 │                             │
│  20 │ Jon        │ Snow      │   2000 │ You know nothing, Jon Snow! │
│ 300 │ Tyrion     │ Lannister │   5000 │                             │
├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
│   0 │ Winter     │ Is        │      0 │ Coming.                     │
│     │            │           │        │ The North Remembers!        │
│     │            │           │        │ This is known.              │
├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
│  11 │ Sansa      │ Stark     │   6000 │                             │
├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
│     │            │ TOTAL     │  10000 │                             │
└─────┴────────────┴───────────┴────────┴─────────────────────────────┘`
	assert.Equal(t, expectedOut, tw.Render())
}

func TestTable_Render_Styles(t *testing.T) {
	tw := NewWriter()
	tw.AppendHeader(testHeader)
	tw.AppendRows(testRows)
	tw.AppendFooter(testFooter)
	tw.SetStyle(StyleLight)

	styles := map[*Style]string{
		&StyleDefault:                    "+-----+------------+-----------+--------+-----------------------------+\n|   # | FIRST NAME | LAST NAME | SALARY |                             |\n+-----+------------+-----------+--------+-----------------------------+\n|   1 | Arya       | Stark     |   3000 |                             |\n|  20 | Jon        | Snow      |   2000 | You know nothing, Jon Snow! |\n| 300 | Tyrion     | Lannister |   5000 |                             |\n+-----+------------+-----------+--------+-----------------------------+\n|     |            | TOTAL     |  10000 |                             |\n+-----+------------+-----------+--------+-----------------------------+",
		&StyleBold:                       "┏━━━━━┳━━━━━━━━━━━━┳━━━━━━━━━━━┳━━━━━━━━┳━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓\n┃   # ┃ FIRST NAME ┃ LAST NAME ┃ SALARY ┃                             ┃\n┣━━━━━╋━━━━━━━━━━━━╋━━━━━━━━━━━╋━━━━━━━━╋━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫\n┃   1 ┃ Arya       ┃ Stark     ┃   3000 ┃                             ┃\n┃  20 ┃ Jon        ┃ Snow      ┃   2000 ┃ You know nothing, Jon Snow! ┃\n┃ 300 ┃ Tyrion     ┃ Lannister ┃   5000 ┃                             ┃\n┣━━━━━╋━━━━━━━━━━━━╋━━━━━━━━━━━╋━━━━━━━━╋━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫\n┃     ┃            ┃ TOTAL     ┃  10000 ┃                             ┃\n┗━━━━━┻━━━━━━━━━━━━┻━━━━━━━━━━━┻━━━━━━━━┻━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛",
		&StyleColoredBlackOnBlueWhite:    "\x1b[104;30m   # \x1b[0m\x1b[104;30m FIRST NAME \x1b[0m\x1b[104;30m LAST NAME \x1b[0m\x1b[104;30m SALARY \x1b[0m\x1b[104;30m                             \x1b[0m\n\x1b[107;30m   1 \x1b[0m\x1b[107;30m Arya       \x1b[0m\x1b[107;30m Stark     \x1b[0m\x1b[107;30m   3000 \x1b[0m\x1b[107;30m                             \x1b[0m\n\x1b[47;30m  20 \x1b[0m\x1b[47;30m Jon        \x1b[0m\x1b[47;30m Snow      \x1b[0m\x1b[47;30m   2000 \x1b[0m\x1b[47;30m You know nothing, Jon Snow! \x1b[0m\n\x1b[107;30m 300 \x1b[0m\x1b[107;30m Tyrion     \x1b[0m\x1b[107;30m Lannister \x1b[0m\x1b[107;30m   5000 \x1b[0m\x1b[107;30m                             \x1b[0m\n\x1b[44;30m     \x1b[0m\x1b[44;30m            \x1b[0m\x1b[44;30m TOTAL     \x1b[0m\x1b[44;30m  10000 \x1b[0m\x1b[44;30m                             \x1b[0m",
		&StyleColoredBlackOnCyanWhite:    "\x1b[106;30m   # \x1b[0m\x1b[106;30m FIRST NAME \x1b[0m\x1b[106;30m LAST NAME \x1b[0m\x1b[106;30m SALARY \x1b[0m\x1b[106;30m                             \x1b[0m\n\x1b[107;30m   1 \x1b[0m\x1b[107;30m Arya       \x1b[0m\x1b[107;30m Stark     \x1b[0m\x1b[107;30m   3000 \x1b[0m\x1b[107;30m                             \x1b[0m\n\x1b[47;30m  20 \x1b[0m\x1b[47;30m Jon        \x1b[0m\x1b[47;30m Snow      \x1b[0m\x1b[47;30m   2000 \x1b[0m\x1b[47;30m You know nothing, Jon Snow! \x1b[0m\n\x1b[107;30m 300 \x1b[0m\x1b[107;30m Tyrion     \x1b[0m\x1b[107;30m Lannister \x1b[0m\x1b[107;30m   5000 \x1b[0m\x1b[107;30m                             \x1b[0m\n\x1b[46;30m     \x1b[0m\x1b[46;30m            \x1b[0m\x1b[46;30m TOTAL     \x1b[0m\x1b[46;30m  10000 \x1b[0m\x1b[46;30m                             \x1b[0m",
		&StyleColoredBlackOnGreenWhite:   "\x1b[102;30m   # \x1b[0m\x1b[102;30m FIRST NAME \x1b[0m\x1b[102;30m LAST NAME \x1b[0m\x1b[102;30m SALARY \x1b[0m\x1b[102;30m                             \x1b[0m\n\x1b[107;30m   1 \x1b[0m\x1b[107;30m Arya       \x1b[0m\x1b[107;30m Stark     \x1b[0m\x1b[107;30m   3000 \x1b[0m\x1b[107;30m                             \x1b[0m\n\x1b[47;30m  20 \x1b[0m\x1b[47;30m Jon        \x1b[0m\x1b[47;30m Snow      \x1b[0m\x1b[47;30m   2000 \x1b[0m\x1b[47;30m You know nothing, Jon Snow! \x1b[0m\n\x1b[107;30m 300 \x1b[0m\x1b[107;30m Tyrion     \x1b[0m\x1b[107;30m Lannister \x1b[0m\x1b[107;30m   5000 \x1b[0m\x1b[107;30m                             \x1b[0m\n\x1b[42;30m     \x1b[0m\x1b[42;30m            \x1b[0m\x1b[42;30m TOTAL     \x1b[0m\x1b[42;30m  10000 \x1b[0m\x1b[42;30m                             \x1b[0m",
		&StyleColoredBlackOnMagentaWhite: "\x1b[105;30m   # \x1b[0m\x1b[105;30m FIRST NAME \x1b[0m\x1b[105;30m LAST NAME \x1b[0m\x1b[105;30m SALARY \x1b[0m\x1b[105;30m                             \x1b[0m\n\x1b[107;30m   1 \x1b[0m\x1b[107;30m Arya       \x1b[0m\x1b[107;30m Stark     \x1b[0m\x1b[107;30m   3000 \x1b[0m\x1b[107;30m                             \x1b[0m\n\x1b[47;30m  20 \x1b[0m\x1b[47;30m Jon        \x1b[0m\x1b[47;30m Snow      \x1b[0m\x1b[47;30m   2000 \x1b[0m\x1b[47;30m You know nothing, Jon Snow! \x1b[0m\n\x1b[107;30m 300 \x1b[0m\x1b[107;30m Tyrion     \x1b[0m\x1b[107;30m Lannister \x1b[0m\x1b[107;30m   5000 \x1b[0m\x1b[107;30m                             \x1b[0m\n\x1b[45;30m     \x1b[0m\x1b[45;30m            \x1b[0m\x1b[45;30m TOTAL     \x1b[0m\x1b[45;30m  10000 \x1b[0m\x1b[45;30m                             \x1b[0m",
		&StyleColoredBlackOnRedWhite:     "\x1b[101;30m   # \x1b[0m\x1b[101;30m FIRST NAME \x1b[0m\x1b[101;30m LAST NAME \x1b[0m\x1b[101;30m SALARY \x1b[0m\x1b[101;30m                             \x1b[0m\n\x1b[107;30m   1 \x1b[0m\x1b[107;30m Arya       \x1b[0m\x1b[107;30m Stark     \x1b[0m\x1b[107;30m   3000 \x1b[0m\x1b[107;30m                             \x1b[0m\n\x1b[47;30m  20 \x1b[0m\x1b[47;30m Jon        \x1b[0m\x1b[47;30m Snow      \x1b[0m\x1b[47;30m   2000 \x1b[0m\x1b[47;30m You know nothing, Jon Snow! \x1b[0m\n\x1b[107;30m 300 \x1b[0m\x1b[107;30m Tyrion     \x1b[0m\x1b[107;30m Lannister \x1b[0m\x1b[107;30m   5000 \x1b[0m\x1b[107;30m                             \x1b[0m\n\x1b[41;30m     \x1b[0m\x1b[41;30m            \x1b[0m\x1b[41;30m TOTAL     \x1b[0m\x1b[41;30m  10000 \x1b[0m\x1b[41;30m                             \x1b[0m",
		&StyleColoredBlackOnYellowWhite:  "\x1b[103;30m   # \x1b[0m\x1b[103;30m FIRST NAME \x1b[0m\x1b[103;30m LAST NAME \x1b[0m\x1b[103;30m SALARY \x1b[0m\x1b[103;30m                             \x1b[0m\n\x1b[107;30m   1 \x1b[0m\x1b[107;30m Arya       \x1b[0m\x1b[107;30m Stark     \x1b[0m\x1b[107;30m   3000 \x1b[0m\x1b[107;30m                             \x1b[0m\n\x1b[47;30m  20 \x1b[0m\x1b[47;30m Jon        \x1b[0m\x1b[47;30m Snow      \x1b[0m\x1b[47;30m   2000 \x1b[0m\x1b[47;30m You know nothing, Jon Snow! \x1b[0m\n\x1b[107;30m 300 \x1b[0m\x1b[107;30m Tyrion     \x1b[0m\x1b[107;30m Lannister \x1b[0m\x1b[107;30m   5000 \x1b[0m\x1b[107;30m                             \x1b[0m\n\x1b[43;30m     \x1b[0m\x1b[43;30m            \x1b[0m\x1b[43;30m TOTAL     \x1b[0m\x1b[43;30m  10000 \x1b[0m\x1b[43;30m                             \x1b[0m",
		&StyleColoredBlueWhiteOnBlack:    "\x1b[94;100m   # \x1b[0m\x1b[94;100m FIRST NAME \x1b[0m\x1b[94;100m LAST NAME \x1b[0m\x1b[94;100m SALARY \x1b[0m\x1b[94;100m                             \x1b[0m\n\x1b[97;40m   1 \x1b[0m\x1b[97;40m Arya       \x1b[0m\x1b[97;40m Stark     \x1b[0m\x1b[97;40m   3000 \x1b[0m\x1b[97;40m                             \x1b[0m\n\x1b[37;40m  20 \x1b[0m\x1b[37;40m Jon        \x1b[0m\x1b[37;40m Snow      \x1b[0m\x1b[37;40m   2000 \x1b[0m\x1b[37;40m You know nothing, Jon Snow! \x1b[0m\n\x1b[97;40m 300 \x1b[0m\x1b[97;40m Tyrion     \x1b[0m\x1b[97;40m Lannister \x1b[0m\x1b[97;40m   5000 \x1b[0m\x1b[97;40m                             \x1b[0m\n\x1b[34;100m     \x1b[0m\x1b[34;100m            \x1b[0m\x1b[34;100m TOTAL     \x1b[0m\x1b[34;100m  10000 \x1b[0m\x1b[34;100m                             \x1b[0m",
		&StyleColoredBright:              "\x1b[106;30m   # \x1b[0m\x1b[106;30m FIRST NAME \x1b[0m\x1b[106;30m LAST NAME \x1b[0m\x1b[106;30m SALARY \x1b[0m\x1b[106;30m                             \x1b[0m\n\x1b[107;30m   1 \x1b[0m\x1b[107;30m Arya       \x1b[0m\x1b[107;30m Stark     \x1b[0m\x1b[107;30m   3000 \x1b[0m\x1b[107;30m                             \x1b[0m\n\x1b[47;30m  20 \x1b[0m\x1b[47;30m Jon        \x1b[0m\x1b[47;30m Snow      \x1b[0m\x1b[47;30m   2000 \x1b[0m\x1b[47;30m You know nothing, Jon Snow! \x1b[0m\n\x1b[107;30m 300 \x1b[0m\x1b[107;30m Tyrion     \x1b[0m\x1b[107;30m Lannister \x1b[0m\x1b[107;30m   5000 \x1b[0m\x1b[107;30m                             \x1b[0m\n\x1b[46;30m     \x1b[0m\x1b[46;30m            \x1b[0m\x1b[46;30m TOTAL     \x1b[0m\x1b[46;30m  10000 \x1b[0m\x1b[46;30m                             \x1b[0m",
		&StyleColoredCyanWhiteOnBlack:    "\x1b[96;100m   # \x1b[0m\x1b[96;100m FIRST NAME \x1b[0m\x1b[96;100m LAST NAME \x1b[0m\x1b[96;100m SALARY \x1b[0m\x1b[96;100m                             \x1b[0m\n\x1b[97;40m   1 \x1b[0m\x1b[97;40m Arya       \x1b[0m\x1b[97;40m Stark     \x1b[0m\x1b[97;40m   3000 \x1b[0m\x1b[97;40m                             \x1b[0m\n\x1b[37;40m  20 \x1b[0m\x1b[37;40m Jon        \x1b[0m\x1b[37;40m Snow      \x1b[0m\x1b[37;40m   2000 \x1b[0m\x1b[37;40m You know nothing, Jon Snow! \x1b[0m\n\x1b[97;40m 300 \x1b[0m\x1b[97;40m Tyrion     \x1b[0m\x1b[97;40m Lannister \x1b[0m\x1b[97;40m   5000 \x1b[0m\x1b[97;40m                             \x1b[0m\n\x1b[36;100m     \x1b[0m\x1b[36;100m            \x1b[0m\x1b[36;100m TOTAL     \x1b[0m\x1b[36;100m  10000 \x1b[0m\x1b[36;100m                             \x1b[0m",
		&StyleColoredDark:                "\x1b[96;100m   # \x1b[0m\x1b[96;100m FIRST NAME \x1b[0m\x1b[96;100m LAST NAME \x1b[0m\x1b[96;100m SALARY \x1b[0m\x1b[96;100m                             \x1b[0m\n\x1b[97;40m   1 \x1b[0m\x1b[97;40m Arya       \x1b[0m\x1b[97;40m Stark     \x1b[0m\x1b[97;40m   3000 \x1b[0m\x1b[97;40m                             \x1b[0m\n\x1b[37;40m  20 \x1b[0m\x1b[37;40m Jon        \x1b[0m\x1b[37;40m Snow      \x1b[0m\x1b[37;40m   2000 \x1b[0m\x1b[37;40m You know nothing, Jon Snow! \x1b[0m\n\x1b[97;40m 300 \x1b[0m\x1b[97;40m Tyrion     \x1b[0m\x1b[97;40m Lannister \x1b[0m\x1b[97;40m   5000 \x1b[0m\x1b[97;40m                             \x1b[0m\n\x1b[36;100m     \x1b[0m\x1b[36;100m            \x1b[0m\x1b[36;100m TOTAL     \x1b[0m\x1b[36;100m  10000 \x1b[0m\x1b[36;100m                             \x1b[0m",
		&StyleColoredGreenWhiteOnBlack:   "\x1b[92;100m   # \x1b[0m\x1b[92;100m FIRST NAME \x1b[0m\x1b[92;100m LAST NAME \x1b[0m\x1b[92;100m SALARY \x1b[0m\x1b[92;100m                             \x1b[0m\n\x1b[97;40m   1 \x1b[0m\x1b[97;40m Arya       \x1b[0m\x1b[97;40m Stark     \x1b[0m\x1b[97;40m   3000 \x1b[0m\x1b[97;40m                             \x1b[0m\n\x1b[37;40m  20 \x1b[0m\x1b[37;40m Jon        \x1b[0m\x1b[37;40m Snow      \x1b[0m\x1b[37;40m   2000 \x1b[0m\x1b[37;40m You know nothing, Jon Snow! \x1b[0m\n\x1b[97;40m 300 \x1b[0m\x1b[97;40m Tyrion     \x1b[0m\x1b[97;40m Lannister \x1b[0m\x1b[97;40m   5000 \x1b[0m\x1b[97;40m                             \x1b[0m\n\x1b[32;100m     \x1b[0m\x1b[32;100m            \x1b[0m\x1b[32;100m TOTAL     \x1b[0m\x1b[32;100m  10000 \x1b[0m\x1b[32;100m                             \x1b[0m",
		&StyleColoredMagentaWhiteOnBlack: "\x1b[95;100m   # \x1b[0m\x1b[95;100m FIRST NAME \x1b[0m\x1b[95;100m LAST NAME \x1b[0m\x1b[95;100m SALARY \x1b[0m\x1b[95;100m                             \x1b[0m\n\x1b[97;40m   1 \x1b[0m\x1b[97;40m Arya       \x1b[0m\x1b[97;40m Stark     \x1b[0m\x1b[97;40m   3000 \x1b[0m\x1b[97;40m                             \x1b[0m\n\x1b[37;40m  20 \x1b[0m\x1b[37;40m Jon        \x1b[0m\x1b[37;40m Snow      \x1b[0m\x1b[37;40m   2000 \x1b[0m\x1b[37;40m You know nothing, Jon Snow! \x1b[0m\n\x1b[97;40m 300 \x1b[0m\x1b[97;40m Tyrion     \x1b[0m\x1b[97;40m Lannister \x1b[0m\x1b[97;40m   5000 \x1b[0m\x1b[97;40m                             \x1b[0m\n\x1b[35;100m     \x1b[0m\x1b[35;100m            \x1b[0m\x1b[35;100m TOTAL     \x1b[0m\x1b[35;100m  10000 \x1b[0m\x1b[35;100m                             \x1b[0m",
		&StyleColoredRedWhiteOnBlack:     "\x1b[91;100m   # \x1b[0m\x1b[91;100m FIRST NAME \x1b[0m\x1b[91;100m LAST NAME \x1b[0m\x1b[91;100m SALARY \x1b[0m\x1b[91;100m                             \x1b[0m\n\x1b[97;40m   1 \x1b[0m\x1b[97;40m Arya       \x1b[0m\x1b[97;40m Stark     \x1b[0m\x1b[97;40m   3000 \x1b[0m\x1b[97;40m                             \x1b[0m\n\x1b[37;40m  20 \x1b[0m\x1b[37;40m Jon        \x1b[0m\x1b[37;40m Snow      \x1b[0m\x1b[37;40m   2000 \x1b[0m\x1b[37;40m You know nothing, Jon Snow! \x1b[0m\n\x1b[97;40m 300 \x1b[0m\x1b[97;40m Tyrion     \x1b[0m\x1b[97;40m Lannister \x1b[0m\x1b[97;40m   5000 \x1b[0m\x1b[97;40m                             \x1b[0m\n\x1b[31;100m     \x1b[0m\x1b[31;100m            \x1b[0m\x1b[31;100m TOTAL     \x1b[0m\x1b[31;100m  10000 \x1b[0m\x1b[31;100m                             \x1b[0m",
		&StyleColoredYellowWhiteOnBlack:  "\x1b[93;100m   # \x1b[0m\x1b[93;100m FIRST NAME \x1b[0m\x1b[93;100m LAST NAME \x1b[0m\x1b[93;100m SALARY \x1b[0m\x1b[93;100m                             \x1b[0m\n\x1b[97;40m   1 \x1b[0m\x1b[97;40m Arya       \x1b[0m\x1b[97;40m Stark     \x1b[0m\x1b[97;40m   3000 \x1b[0m\x1b[97;40m                             \x1b[0m\n\x1b[37;40m  20 \x1b[0m\x1b[37;40m Jon        \x1b[0m\x1b[37;40m Snow      \x1b[0m\x1b[37;40m   2000 \x1b[0m\x1b[37;40m You know nothing, Jon Snow! \x1b[0m\n\x1b[97;40m 300 \x1b[0m\x1b[97;40m Tyrion     \x1b[0m\x1b[97;40m Lannister \x1b[0m\x1b[97;40m   5000 \x1b[0m\x1b[97;40m                             \x1b[0m\n\x1b[33;100m     \x1b[0m\x1b[33;100m            \x1b[0m\x1b[33;100m TOTAL     \x1b[0m\x1b[33;100m  10000 \x1b[0m\x1b[33;100m                             \x1b[0m",
		&StyleDouble:                     "╔═════╦════════════╦═══════════╦════════╦═════════════════════════════╗\n║   # ║ FIRST NAME ║ LAST NAME ║ SALARY ║                             ║\n╠═════╬════════════╬═══════════╬════════╬═════════════════════════════╣\n║   1 ║ Arya       ║ Stark     ║   3000 ║                             ║\n║  20 ║ Jon        ║ Snow      ║   2000 ║ You know nothing, Jon Snow! ║\n║ 300 ║ Tyrion     ║ Lannister ║   5000 ║                             ║\n╠═════╬════════════╬═══════════╬════════╬═════════════════════════════╣\n║     ║            ║ TOTAL     ║  10000 ║                             ║\n╚═════╩════════════╩═══════════╩════════╩═════════════════════════════╝",
		&StyleLight:                      "┌─────┬────────────┬───────────┬────────┬─────────────────────────────┐\n│   # │ FIRST NAME │ LAST NAME │ SALARY │                             │\n├─────┼────────────┼───────────┼────────┼─────────────────────────────┤\n│   1 │ Arya       │ Stark     │   3000 │                             │\n│  20 │ Jon        │ Snow      │   2000 │ You know nothing, Jon Snow! │\n│ 300 │ Tyrion     │ Lannister │   5000 │                             │\n├─────┼────────────┼───────────┼────────┼─────────────────────────────┤\n│     │            │ TOTAL     │  10000 │                             │\n└─────┴────────────┴───────────┴────────┴─────────────────────────────┘",
		&StyleRounded:                    "╭─────┬────────────┬───────────┬────────┬─────────────────────────────╮\n│   # │ FIRST NAME │ LAST NAME │ SALARY │                             │\n├─────┼────────────┼───────────┼────────┼─────────────────────────────┤\n│   1 │ Arya       │ Stark     │   3000 │                             │\n│  20 │ Jon        │ Snow      │   2000 │ You know nothing, Jon Snow! │\n│ 300 │ Tyrion     │ Lannister │   5000 │                             │\n├─────┼────────────┼───────────┼────────┼─────────────────────────────┤\n│     │            │ TOTAL     │  10000 │                             │\n╰─────┴────────────┴───────────┴────────┴─────────────────────────────╯",
		&styleTest:                       "(-----^------------^-----------^--------^-----------------------------)\n[<  #>|<FIRST NAME>|<LAST NAME>|<SALARY>|<                           >]\n{-----+------------+-----------+--------+-----------------------------}\n[<  1>|<Arya      >|<Stark    >|<  3000>|<                           >]\n[< 20>|<Jon       >|<Snow     >|<  2000>|<You know nothing, Jon Snow!>]\n[<300>|<Tyrion    >|<Lannister>|<  5000>|<                           >]\n{-----+------------+-----------+--------+-----------------------------}\n[<   >|<          >|<TOTAL    >|< 10000>|<                           >]\n\\-----v------------v-----------v--------v-----------------------------/",
	}
	var mismatches []string
	for style, expectedOut := range styles {
		tw.SetStyle(*style)
		out := tw.Render()
		assert.Equal(t, expectedOut, out)
		if expectedOut != out {
			mismatches = append(mismatches, fmt.Sprintf("&%s: %#v,", style.Name, out))
			fmt.Printf("// %s renders a Table like below:\n", style.Name)
			for _, line := range strings.Split(out, "\n") {
				fmt.Printf("//  %s\n", line)
			}
			fmt.Println()
		}
	}
	sort.Strings(mismatches)
	for _, mismatch := range mismatches {
		fmt.Println(mismatch)
	}
}

func TestTable_Render_SuppressEmptyColumns(t *testing.T) {
	tw := NewWriter()
	tw.AppendHeader(testHeader)
	tw.AppendRows([]Row{
		{1, "Arya", "", 3000},
		{20, "Jon", "", 2000, "You know nothing, Jon Snow!"},
		{300, "Tyrion", "", 5000},
	})
	tw.AppendRow(Row{11, "Sansa", "", 6000})
	tw.AppendFooter(Row{"", "", "TOTAL", 10000})
	tw.SetStyle(StyleLight)

	expectedOut := `┌─────┬────────────┬───────────┬────────┬─────────────────────────────┐
│   # │ FIRST NAME │ LAST NAME │ SALARY │                             │
├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
│   1 │ Arya       │           │   3000 │                             │
│  20 │ Jon        │           │   2000 │ You know nothing, Jon Snow! │
│ 300 │ Tyrion     │           │   5000 │                             │
│  11 │ Sansa      │           │   6000 │                             │
├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
│     │            │ TOTAL     │  10000 │                             │
└─────┴────────────┴───────────┴────────┴─────────────────────────────┘`
	assert.Equal(t, expectedOut, tw.Render())

	tw.SuppressEmptyColumns()
	expectedOut = `┌─────┬────────────┬────────┬─────────────────────────────┐
│   # │ FIRST NAME │ SALARY │                             │
├─────┼────────────┼────────┼─────────────────────────────┤
│   1 │ Arya       │   3000 │                             │
│  20 │ Jon        │   2000 │ You know nothing, Jon Snow! │
│ 300 │ Tyrion     │   5000 │                             │
│  11 │ Sansa      │   6000 │                             │
├─────┼────────────┼────────┼─────────────────────────────┤
│     │            │  10000 │                             │
└─────┴────────────┴────────┴─────────────────────────────┘`
	assert.Equal(t, expectedOut, tw.Render())
}

func TestTable_Render_TableWithinTable(t *testing.T) {
	twInner := NewWriter()
	twInner.AppendHeader(testHeader)
	twInner.AppendRows(testRows)
	twInner.AppendFooter(testFooter)
	twInner.SetStyle(StyleLight)

	twOuter := NewWriter()
	twOuter.AppendHeader(Row{"Table within a Table"})
	twOuter.AppendRow(Row{twInner.Render()})
	twOuter.SetColumnConfigs([]ColumnConfig{{Number: 1, AlignHeader: text.AlignCenter}})
	twOuter.SetStyle(StyleDouble)

	expectedOut := `╔═════════════════════════════════════════════════════════════════════════╗
║                           TABLE WITHIN A TABLE                          ║
╠═════════════════════════════════════════════════════════════════════════╣
║ ┌─────┬────────────┬───────────┬────────┬─────────────────────────────┐ ║
║ │   # │ FIRST NAME │ LAST NAME │ SALARY │                             │ ║
║ ├─────┼────────────┼───────────┼────────┼─────────────────────────────┤ ║
║ │   1 │ Arya       │ Stark     │   3000 │                             │ ║
║ │  20 │ Jon        │ Snow      │   2000 │ You know nothing, Jon Snow! │ ║
║ │ 300 │ Tyrion     │ Lannister │   5000 │                             │ ║
║ ├─────┼────────────┼───────────┼────────┼─────────────────────────────┤ ║
║ │     │            │ TOTAL     │  10000 │                             │ ║
║ └─────┴────────────┴───────────┴────────┴─────────────────────────────┘ ║
╚═════════════════════════════════════════════════════════════════════════╝`
	assert.Equal(t, expectedOut, twOuter.Render())
}

func TestTable_Render_TableWithTransformers(t *testing.T) {
	bolden := func(val interface{}) string {
		return text.Bold.Sprint(val)
	}
	tw := NewWriter()
	tw.AppendHeader(testHeader)
	tw.AppendRows(testRows)
	tw.AppendFooter(testFooter)
	tw.SetColumnConfigs([]ColumnConfig{{
		Name:              "Salary",
		Transformer:       bolden,
		TransformerFooter: bolden,
		TransformerHeader: bolden,
	}})
	tw.SetStyle(StyleLight)

	expectedOut := []string{
		"┌─────┬────────────┬───────────┬────────┬─────────────────────────────┐",
		"│   # │ FIRST NAME │ LAST NAME │ \x1b[1mSALARY\x1b[0m │                             │",
		"├─────┼────────────┼───────────┼────────┼─────────────────────────────┤",
		"│   1 │ Arya       │ Stark     │   \x1b[1m3000\x1b[0m │                             │",
		"│  20 │ Jon        │ Snow      │   \x1b[1m2000\x1b[0m │ You know nothing, Jon Snow! │",
		"│ 300 │ Tyrion     │ Lannister │   \x1b[1m5000\x1b[0m │                             │",
		"├─────┼────────────┼───────────┼────────┼─────────────────────────────┤",
		"│     │            │ TOTAL     │  \x1b[1m10000\x1b[0m │                             │",
		"└─────┴────────────┴───────────┴────────┴─────────────────────────────┘",
	}
	out := tw.Render()
	assert.Equal(t, strings.Join(expectedOut, "\n"), out)
	if strings.Join(expectedOut, "\n") != out {
		for _, line := range strings.Split(out, "\n") {
			fmt.Printf("%#v,\n", line)
		}
	}
}

func TestTable_Render_SetWidth_Title(t *testing.T) {
	tw := NewWriter()
	tw.AppendHeader(testHeader)
	tw.AppendRows(testRows)
	tw.AppendFooter(testFooter)
	tw.SetTitle("Game Of Thrones")

	t.Run("length 20", func(t *testing.T) {
		tw.SetAllowedRowLength(20)

		expectedOut := []string{
			"+------------------+",
			"| Game Of Thrones  |",
			"+-----+----------- ~",
			"|   # | FIRST NAME ~",
			"+-----+----------- ~",
			"|   1 | Arya       ~",
			"|  20 | Jon        ~",
			"| 300 | Tyrion     ~",
			"+-----+----------- ~",
			"|     |            ~",
			"+-----+----------- ~",
		}

		assert.Equal(t, strings.Join(expectedOut, "\n"), tw.Render())
	})

	t.Run("length 30", func(t *testing.T) {
		tw.SetAllowedRowLength(30)

		expectedOut := []string{
			"+----------------------------+",
			"| Game Of Thrones            |",
			"+-----+------------+-------- ~",
			"|   # | FIRST NAME | LAST NA ~",
			"+-----+------------+-------- ~",
			"|   1 | Arya       | Stark   ~",
			"|  20 | Jon        | Snow    ~",
			"| 300 | Tyrion     | Lannist ~",
			"+-----+------------+-------- ~",
			"|     |            | TOTAL   ~",
			"+-----+------------+-------- ~",
		}

		assert.Equal(t, strings.Join(expectedOut, "\n"), tw.Render())
	})

}
