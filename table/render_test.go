package table

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"fmt"
	"github.com/jedib0t/go-pretty/util"
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
			row[colIdx] = fmt.Sprintf("%s%d", util.AutoIndexColumnID(colIdx), rowIdx+1)
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

func TestTable_Render_Colored(t *testing.T) {
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
		"├─────┼────────────┼───────────┼────────┼─────────────────────────────┤",
		"│     │            │\x1b[94;1m TOTAL     \x1b[0m│\x1b[94;1m  10000 \x1b[0m│                             │",
		"╰─────┴────────────┴───────────┴────────┴─────────────────────────────╯",
		"test-caption",
	}

	assert.Equal(t, strings.Join(expectedOut, "\n"), tw.Render())
}

func TestTable_Render_Empty(t *testing.T) {
	tw := NewWriter()
	assert.Empty(t, tw.Render())
}
