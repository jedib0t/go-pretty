package table

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/jedib0t/go-pretty/text"
	"github.com/stretchr/testify/assert"
)

func TestTable_Render(t *testing.T) {
	tw := NewWriter()
	tw.AppendHeader(testHeader)
	tw.AppendRows(testRows)
	tw.AppendRow(testRowMultiLine)
	tw.AppendFooter(testFooter)
	tw.SetAlign(testAlign)
	tw.SetCaption(testCaption)
	tw.SetStyle(styleTest)

	expectedOut := `(-----^------------^-----------^--------^-----------------------------)
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
test-caption`
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
	tw.SetAlign(testAlign)
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
		"\x1b[106;30m+\x1b[0m\x1b[106;30m---\x1b[0m\x1b[106;30m+\x1b[0m\x1b[107;30m-----\x1b[0m\x1b[107;30m+\x1b[0m\x1b[107;30m------------\x1b[0m\x1b[107;30m+\x1b[0m\x1b[107;30m-----------\x1b[0m\x1b[107;30m+\x1b[0m\x1b[107;30m--------\x1b[0m\x1b[107;30m+\x1b[0m\x1b[107;30m-----------------------------\x1b[0m\x1b[106;30m+\x1b[0m",
		"\x1b[106;30m|\x1b[0m\x1b[106;30m 3 \x1b[0m\x1b[106;30m|\x1b[0m\x1b[107;30m 300 \x1b[0m\x1b[107;30m|\x1b[0m\x1b[107;30m Tyrion     \x1b[0m\x1b[107;30m|\x1b[0m\x1b[107;30m Lannister \x1b[0m\x1b[107;30m|\x1b[0m\x1b[107;30m   5000 \x1b[0m\x1b[107;30m|\x1b[0m\x1b[107;30m                             \x1b[0m\x1b[106;30m|\x1b[0m",
		"\x1b[106;30m+\x1b[0m\x1b[106;30m---\x1b[0m\x1b[106;30m+\x1b[0m\x1b[107;30m-----\x1b[0m\x1b[107;30m+\x1b[0m\x1b[107;30m------------\x1b[0m\x1b[107;30m+\x1b[0m\x1b[107;30m-----------\x1b[0m\x1b[107;30m+\x1b[0m\x1b[107;30m--------\x1b[0m\x1b[107;30m+\x1b[0m\x1b[107;30m-----------------------------\x1b[0m\x1b[106;30m+\x1b[0m",
		"\x1b[106;30m|\x1b[0m\x1b[106;30m 4 \x1b[0m\x1b[106;30m|\x1b[0m\x1b[47;30m   0 \x1b[0m\x1b[47;30m|\x1b[0m\x1b[47;30m Winter     \x1b[0m\x1b[47;30m|\x1b[0m\x1b[47;30m Is        \x1b[0m\x1b[47;30m|\x1b[0m\x1b[47;30m      0 \x1b[0m\x1b[47;30m|\x1b[0m\x1b[47;30m Coming.                     \x1b[0m\x1b[106;30m|\x1b[0m",
		"\x1b[106;30m|\x1b[0m\x1b[106;30m   \x1b[0m\x1b[106;30m|\x1b[0m\x1b[47;30m     \x1b[0m\x1b[47;30m|\x1b[0m\x1b[47;30m            \x1b[0m\x1b[47;30m|\x1b[0m\x1b[47;30m           \x1b[0m\x1b[47;30m|\x1b[0m\x1b[47;30m        \x1b[0m\x1b[47;30m|\x1b[0m\x1b[47;30m The North Remembers!        \x1b[0m\x1b[106;30m|\x1b[0m",
		"\x1b[106;30m|\x1b[0m\x1b[106;30m   \x1b[0m\x1b[106;30m|\x1b[0m\x1b[47;30m     \x1b[0m\x1b[47;30m|\x1b[0m\x1b[47;30m            \x1b[0m\x1b[47;30m|\x1b[0m\x1b[47;30m           \x1b[0m\x1b[47;30m|\x1b[0m\x1b[47;30m        \x1b[0m\x1b[47;30m|\x1b[0m\x1b[47;30m This is known.              \x1b[0m\x1b[106;30m|\x1b[0m",
		"\x1b[46;30m+\x1b[0m\x1b[46;30m---\x1b[0m\x1b[46;30m+\x1b[0m\x1b[46;30m-----\x1b[0m\x1b[46;30m+\x1b[0m\x1b[46;30m------------\x1b[0m\x1b[46;30m+\x1b[0m\x1b[46;30m-----------\x1b[0m\x1b[46;30m+\x1b[0m\x1b[46;30m--------\x1b[0m\x1b[46;30m+\x1b[0m\x1b[46;30m-----------------------------\x1b[0m\x1b[46;30m+\x1b[0m",
		"\x1b[46;30m|\x1b[0m\x1b[46;30m   \x1b[0m\x1b[46;30m|\x1b[0m\x1b[46;30m     \x1b[0m\x1b[46;30m|\x1b[0m\x1b[46;30m            \x1b[0m\x1b[46;30m|\x1b[0m\x1b[46;30m TOTAL     \x1b[0m\x1b[46;30m|\x1b[0m\x1b[46;30m  10000 \x1b[0m\x1b[46;30m|\x1b[0m\x1b[46;30m                             \x1b[0m\x1b[46;30m|\x1b[0m",
		"\x1b[46;30m+\x1b[0m\x1b[46;30m---\x1b[0m\x1b[46;30m+\x1b[0m\x1b[46;30m-----\x1b[0m\x1b[46;30m+\x1b[0m\x1b[46;30m------------\x1b[0m\x1b[46;30m+\x1b[0m\x1b[46;30m-----------\x1b[0m\x1b[46;30m+\x1b[0m\x1b[46;30m--------\x1b[0m\x1b[46;30m+\x1b[0m\x1b[46;30m-----------------------------\x1b[0m\x1b[46;30m+\x1b[0m",
	}
	assert.Equal(t, strings.Join(expectedOut, "\n"), tw.Render())
}

func TestTable_Render_ColoredCustom(t *testing.T) {
	tw := NewWriter()
	tw.AppendHeader(testHeader)
	tw.AppendRows(testRows)
	tw.AppendRow(testRowMultiLine)
	tw.AppendFooter(testFooter)
	tw.SetAlign(testAlign)
	tw.SetCaption(testCaption)
	tw.SetColors(testColors)
	tw.SetColorsFooter(testColorsFooter)
	tw.SetColorsHeader(testColorsHeader)
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
		"test-caption",
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
	tableOuter.SetAlignHeader([]text.Align{text.AlignCenter})
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

	expectedOut := strings.Join([]string{
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

func TestTable_Render_Empty(t *testing.T) {
	tw := NewWriter()
	assert.Empty(t, tw.Render())
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

func TestList_Render_Styles(t *testing.T) {
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

func TestTable_Render_TableWithinTable(t *testing.T) {
	twInner := NewWriter()
	twInner.AppendHeader(testHeader)
	twInner.AppendRows(testRows)
	twInner.AppendFooter(testFooter)
	twInner.SetAlignFooter([]text.Align{text.AlignDefault, text.AlignDefault, text.AlignLeft, text.AlignRight})
	twInner.SetStyle(StyleLight)

	twOuter := NewWriter()
	twOuter.AppendHeader(Row{"Table within a Table"})
	twOuter.AppendRow(Row{twInner.Render()})
	twOuter.SetAlignHeader([]text.Align{text.AlignCenter})
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
