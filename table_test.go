package gopretty

import (
	"strings"
	"testing"

	"github.com/fatih/color"
	"github.com/stretchr/testify/assert"
)

var (
	testAlign        = []Align{AlignDefault, AlignLeft, AlignLeft, AlignRight}
	testCaption      = "test-caption"
	testCSSClass     = "test-css-class"
	testColor        = TextColor{color.FgGreen}
	testColorHeader  = TextColor{color.FgHiRed, color.Bold}
	testColorFooter  = TextColor{color.FgHiBlue, color.Bold}
	testColors       = []TextColor{testColor, testColor, testColor, testColor, {color.FgCyan}}
	testColorsFooter = []TextColor{{}, {}, testColorFooter, testColorFooter}
	testColorsHeader = []TextColor{testColorHeader, testColorHeader, testColorHeader, testColorHeader}
	testFooter       = TableRow{"", "", "Total", 10000}
	testHeader       = TableRow{"#", "First Name", "Last Name", "Salary"}
	testRows         = []TableRow{
		{1, "Arya", "Stark", 3000},
		{20, "Jon", "Snow", 2000, "You know nothing, Jon Snow!"},
		{300, "Tyrion", "Lannister", 5000},
	}
	testRowMultiLine = TableRow{0, "Winter", "Is", 0, "Coming.\nThe North Remembers!"}
	testTextColor1   = TextColor{color.FgWhite, color.BgBlack}
	testTextColor2   = TextColor{color.FgBlack, color.BgWhite}
)

func BenchmarkTable_Render(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tw := NewTableWriter()
		tw.AppendHeader(testHeader)
		tw.AppendRows(testRows)
		tw.AppendFooter(testFooter)
		tw.SetAlign(testAlign)
		tw.SetCaption(testCaption)
		tw.Render()
	}
}

func BenchmarkTable_RenderCSV(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tw := NewTableWriter()
		tw.AppendHeader(testHeader)
		tw.AppendRows(testRows)
		tw.AppendFooter(testFooter)
		tw.SetAlign(testAlign)
		tw.SetCaption(testCaption)
		tw.RenderCSV()
	}
}

func BenchmarkTable_RenderHTML(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tw := NewTableWriter()
		tw.AppendHeader(testHeader)
		tw.AppendRows(testRows)
		tw.AppendFooter(testFooter)
		tw.SetAlign(testAlign)
		tw.SetCaption(testCaption)
		tw.RenderHTML()
	}
}

func TestNewTableWriter(t *testing.T) {
	tw := NewTableWriter()
	assert.Nil(t, tw.Style())

	tw.SetStyle(TableStyleBold)
	assert.NotNil(t, tw.Style())
	assert.Equal(t, TableStyleBold, *tw.Style())
}

func TestTable_AppendFooter(t *testing.T) {
	table := Table{}
	assert.Equal(t, 0, len(table.rowsFooter))

	table.AppendFooter([]interface{}{})
	assert.Equal(t, 0, table.Length())
	assert.Equal(t, 1, len(table.rowsFooter))
	assert.Equal(t, 0, len(table.rowsHeader))

	table.AppendFooter([]interface{}{})
	assert.Equal(t, 0, table.Length())
	assert.Equal(t, 2, len(table.rowsFooter))
	assert.Equal(t, 0, len(table.rowsHeader))
}

func TestTable_AppendHeader(t *testing.T) {
	table := Table{}
	assert.Equal(t, 0, len(table.rowsHeader))

	table.AppendHeader([]interface{}{})
	assert.Equal(t, 0, table.Length())
	assert.Equal(t, 0, len(table.rowsFooter))
	assert.Equal(t, 1, len(table.rowsHeader))

	table.AppendHeader([]interface{}{})
	assert.Equal(t, 0, table.Length())
	assert.Equal(t, 0, len(table.rowsFooter))
	assert.Equal(t, 2, len(table.rowsHeader))
}

func TestTable_AppendRow(t *testing.T) {
	table := Table{}
	assert.Equal(t, 0, table.Length())

	table.AppendRow([]interface{}{})
	assert.Equal(t, 1, table.Length())
	assert.Equal(t, 0, len(table.rowsFooter))
	assert.Equal(t, 0, len(table.rowsHeader))

	table.AppendRow([]interface{}{})
	assert.Equal(t, 2, table.Length())
	assert.Equal(t, 0, len(table.rowsFooter))
	assert.Equal(t, 0, len(table.rowsHeader))
}

func TestTable_AppendRows(t *testing.T) {
	table := Table{}
	assert.Equal(t, 0, table.Length())

	table.AppendRows([]TableRow{{}})
	assert.Equal(t, 1, table.Length())
	assert.Equal(t, 0, len(table.rowsFooter))
	assert.Equal(t, 0, len(table.rowsHeader))

	table.AppendRows([]TableRow{{}})
	assert.Equal(t, 2, table.Length())
	assert.Equal(t, 0, len(table.rowsFooter))
	assert.Equal(t, 0, len(table.rowsHeader))
}

func TestTable_DisableBorder(t *testing.T) {
	table := Table{}
	assert.False(t, table.disableBorder)

	table.DisableBorder()
	assert.True(t, table.disableBorder)

	table.AppendRow(testRows[0])
	out := table.Render()

	assert.NotEmpty(t, out)
	assert.Equal(t, 0, strings.Count(out, "\n"))
	assert.Equal(t, " 1 | Arya | Stark | 3000 ", out)
}

func TestTable_EnableSeparators(t *testing.T) {
	table := Table{}
	assert.False(t, table.enableSeparators)

	table.EnableSeparators()
	assert.True(t, table.enableSeparators)

	table.AppendRows(testRows)

	expectedOut := `+-----+--------+-----------+------+-----------------------------+
|   1 | Arya   | Stark     | 3000 |                             |
+-----+--------+-----------+------+-----------------------------+
|  20 | Jon    | Snow      | 2000 | You know nothing, Jon Snow! |
+-----+--------+-----------+------+-----------------------------+
| 300 | Tyrion | Lannister | 5000 |                             |
+-----+--------+-----------+------+-----------------------------+`

	assert.Equal(t, expectedOut, table.Render())
}

func TestTable_Render(t *testing.T) {
	tw := NewTableWriter()
	tw.AppendHeader(testHeader)
	tw.AppendRows(testRows)
	tw.AppendRow(testRowMultiLine)
	tw.AppendFooter(testFooter)
	tw.SetAlign(testAlign)
	tw.SetCaption(testCaption)
	tw.SetStyle(tableStyleTest)

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

func TestTable_RenderColored(t *testing.T) {
	tw := NewTableWriter()
	tw.AppendHeader(testHeader)
	tw.AppendRows(testRows)
	tw.AppendRow(testRowMultiLine)
	tw.AppendFooter(testFooter)
	tw.SetAlign(testAlign)
	tw.SetCaption(testCaption)
	tw.SetColors(testColors)
	tw.SetColorsFooter(testColorsFooter)
	tw.SetColorsHeader(testColorsHeader)
	tw.SetStyle(TableStyleRounded)

	expectedOut := []string{
		"╭─────┬────────────┬───────────┬────────┬─────────────────────────────╮",
		"│ \x1b[91;1m  #\x1b[0m │ \x1b[91;1mFIRST NAME\x1b[0m │ \x1b[91;1mLAST NAME\x1b[0m │ \x1b[91;1mSALARY\x1b[0m │                             │",
		"├─────┼────────────┼───────────┼────────┼─────────────────────────────┤",
		"│ \x1b[32m  1\x1b[0m │ \x1b[32mArya      \x1b[0m │ \x1b[32mStark    \x1b[0m │ \x1b[32m  3000\x1b[0m │ \x1b[36m                           \x1b[0m │",
		"│ \x1b[32m 20\x1b[0m │ \x1b[32mJon       \x1b[0m │ \x1b[32mSnow     \x1b[0m │ \x1b[32m  2000\x1b[0m │ \x1b[36mYou know nothing, Jon Snow!\x1b[0m │",
		"│ \x1b[32m300\x1b[0m │ \x1b[32mTyrion    \x1b[0m │ \x1b[32mLannister\x1b[0m │ \x1b[32m  5000\x1b[0m │ \x1b[36m                           \x1b[0m │",
		"│ \x1b[32m  0\x1b[0m │ \x1b[32mWinter    \x1b[0m │ \x1b[32mIs       \x1b[0m │ \x1b[32m     0\x1b[0m │ \x1b[36mComing.                    \x1b[0m │",
		"│ \x1b[32m   \x1b[0m │ \x1b[32m          \x1b[0m │ \x1b[32m         \x1b[0m │ \x1b[32m      \x1b[0m │ \x1b[36mThe North Remembers!       \x1b[0m │",
		"├─────┼────────────┼───────────┼────────┼─────────────────────────────┤",
		"│     │            │ \x1b[94;1mTOTAL    \x1b[0m │ \x1b[94;1m 10000\x1b[0m │                             │",
		"╰─────┴────────────┴───────────┴────────┴─────────────────────────────╯",
		"test-caption",
	}

	assert.Equal(t, strings.Join(expectedOut, "\n"), tw.Render())
}

func TestTable_RenderCSV(t *testing.T) {
	tw := NewTableWriter()
	tw.AppendHeader(testHeader)
	tw.AppendRows(testRows)
	tw.AppendRow(testRowMultiLine)
	tw.AppendFooter(testFooter)

	expectedOut := `#,First Name,Last Name,Salary,
1,Arya,Stark,3000,
20,Jon,Snow,2000,"You know nothing\, Jon Snow!"
300,Tyrion,Lannister,5000,
0,Winter,Is,0,"Coming.
The North Remembers!"
,,Total,10000,`

	assert.Equal(t, expectedOut, tw.RenderCSV())
}

func TestTable_RenderHTML(t *testing.T) {
	tw := NewTableWriter()
	tw.AppendHeader(testHeader)
	tw.AppendRows(testRows)
	tw.AppendRow(testRowMultiLine)
	tw.AppendFooter(testFooter)
	tw.SetHTMLCSSClass(testCSSClass)
	tw.SetVAlign([]VAlign{VAlignDefault, VAlignDefault, VAlignDefault, VAlignBottom, VAlignBottom})

	expectedOut := `<table class="test-css-class">
  <thead>
  <tr>
    <th align="right">#</th>
    <th>First Name</th>
    <th>Last Name</th>
    <th align="right" valign="bottom">Salary</th>
    <th valign="bottom">&nbsp;</th>
  </tr>
  </thead>
  <tbody>
  <tr>
    <td align="right">1</td>
    <td>Arya</td>
    <td>Stark</td>
    <td align="right" valign="bottom">3000</td>
    <td valign="bottom">&nbsp;</td>
  </tr>
  <tr>
    <td align="right">20</td>
    <td>Jon</td>
    <td>Snow</td>
    <td align="right" valign="bottom">2000</td>
    <td valign="bottom">You know nothing, Jon Snow!</td>
  </tr>
  <tr>
    <td align="right">300</td>
    <td>Tyrion</td>
    <td>Lannister</td>
    <td align="right" valign="bottom">5000</td>
    <td valign="bottom">&nbsp;</td>
  </tr>
  <tr>
    <td align="right">0</td>
    <td>Winter</td>
    <td>Is</td>
    <td align="right" valign="bottom">0</td>
    <td valign="bottom">Coming.<br/>The North Remembers!</td>
  </tr>
  </tbody>
  <tfoot>
  <tr>
    <td align="right">&nbsp;</td>
    <td>&nbsp;</td>
    <td>Total</td>
    <td align="right" valign="bottom">10000</td>
    <td valign="bottom">&nbsp;</td>
  </tr>
  </tfoot>
</table>`

	assert.Equal(t, expectedOut, tw.RenderHTML())
}

func TestTable_SetAlign(t *testing.T) {
	table := Table{}
	assert.Nil(t, table.align)

	table.SetAlign([]Align{})
	assert.NotNil(t, table.align)

	table.AppendRows(testRows)
	table.AppendRow(testRowMultiLine)
	table.SetAlign([]Align{AlignDefault, AlignLeft, AlignLeft, AlignRight, AlignRight})

	expectedOut := `+-----+--------+-----------+------+-----------------------------+
|   1 | Arya   | Stark     | 3000 |                             |
|  20 | Jon    | Snow      | 2000 | You know nothing, Jon Snow! |
| 300 | Tyrion | Lannister | 5000 |                             |
|   0 | Winter | Is        |    0 |                     Coming. |
|     |        |           |      |        The North Remembers! |
+-----+--------+-----------+------+-----------------------------+`

	assert.Equal(t, expectedOut, table.Render())
}

func TestTable_SetCaption(t *testing.T) {
	table := Table{}
	assert.Empty(t, table.caption)

	table.SetCaption(testCaption)
	assert.NotEmpty(t, table.caption)
	assert.Equal(t, testCaption, table.caption)
}

func TestTable_SetColors(t *testing.T) {
	table := Table{}
	assert.Empty(t, table.colors)
	assert.Empty(t, table.colorsFooter)
	assert.Empty(t, table.colorsHeader)

	table.SetColors([]TextColor{testTextColor1, testTextColor2})
	assert.NotEmpty(t, table.colors)
	assert.Empty(t, table.colorsFooter)
	assert.Empty(t, table.colorsHeader)
	assert.Equal(t, 2, len(table.colors))
}

func TestTable_SetColorsFooter(t *testing.T) {
	table := Table{}
	assert.Empty(t, table.colors)
	assert.Empty(t, table.colorsFooter)
	assert.Empty(t, table.colorsHeader)

	table.SetColorsFooter([]TextColor{testTextColor1, testTextColor2})
	assert.Empty(t, table.colors)
	assert.NotEmpty(t, table.colorsFooter)
	assert.Empty(t, table.colorsHeader)
	assert.Equal(t, 2, len(table.colorsFooter))
}

func TestTable_SetColorsHeader(t *testing.T) {
	table := Table{}
	assert.Empty(t, table.colors)
	assert.Empty(t, table.colorsFooter)
	assert.Empty(t, table.colorsHeader)

	table.SetColorsHeader([]TextColor{testTextColor1, testTextColor2})
	assert.Empty(t, table.colors)
	assert.Empty(t, table.colorsFooter)
	assert.NotEmpty(t, table.colorsHeader)
	assert.Equal(t, 2, len(table.colorsHeader))
}

func TestTable_SetVAlign(t *testing.T) {
	table := Table{}
	assert.Nil(t, table.vAlign)

	table.SetVAlign([]VAlign{})
	assert.NotNil(t, table.vAlign)

	table.AppendRow(testRowMultiLine)
	table.SetVAlign([]VAlign{VAlignTop, VAlignMiddle, VAlignBottom, VAlignDefault})

	expectedOut := `+---+--------+----+---+----------------------+
| 0 | Winter |    | 0 | Coming.              |
|   |        | Is |   | The North Remembers! |
+---+--------+----+---+----------------------+`

	assert.Equal(t, expectedOut, table.Render())
}

func TestTable_SetStyle(t *testing.T) {
	table := Table{}
	assert.Nil(t, table.Style())

	table.SetStyle(TableStyleDefault)
	assert.NotNil(t, table.Style())
	assert.Equal(t, &TableStyleDefault, table.Style())
}
