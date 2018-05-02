package table

import (
	"strings"
	"testing"

	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/text"
	"github.com/stretchr/testify/assert"
)

var (
	testAlign        = []text.Align{text.AlignDefault, text.AlignLeft, text.AlignLeft, text.AlignRight}
	testCaption      = "test-caption"
	testCSSClass     = "test-css-class"
	testColor        = text.Colors{color.FgGreen}
	testColorHeader  = text.Colors{color.FgHiRed, color.Bold}
	testColorFooter  = text.Colors{color.FgHiBlue, color.Bold}
	testColors       = []text.Colors{testColor, testColor, testColor, testColor, {color.FgCyan}}
	testColorsFooter = []text.Colors{{}, {}, testColorFooter, testColorFooter}
	testColorsHeader = []text.Colors{testColorHeader, testColorHeader, testColorHeader, testColorHeader}
	testFooter       = Row{"", "", "Total", 10000}
	testHeader       = Row{"#", "First Name", "Last Name", "Salary"}
	testRows         = []Row{
		{1, "Arya", "Stark", 3000},
		{20, "Jon", "Snow", 2000, "You know nothing, Jon Snow!"},
		{300, "Tyrion", "Lannister", 5000},
	}
	testRowMultiLine = Row{0, "Winter", "Is", 0, "Coming.\nThe North Remembers!"}
	testRowTabs      = Row{0, "Valar", "Morghulis", 0, "\t"}
	testColors1      = text.Colors{color.FgWhite, color.BgBlack}
	testColors2      = text.Colors{color.FgBlack, color.BgWhite}
)

type myMockOutputMirror struct{
	mirroredOutput string
}

func (t *myMockOutputMirror) Write(p []byte) (n int, err error) {
	t.mirroredOutput = string(p)
	return len(p), nil
}

func BenchmarkTable_Render(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tw := NewWriter()
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
		tw := NewWriter()
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
		tw := NewWriter()
		tw.AppendHeader(testHeader)
		tw.AppendRows(testRows)
		tw.AppendFooter(testFooter)
		tw.SetAlign(testAlign)
		tw.SetCaption(testCaption)
		tw.RenderHTML()
	}
}

func TestNewWriter(t *testing.T) {
	tw := NewWriter()
	assert.Nil(t, tw.Style())

	tw.SetStyle(StyleBold)
	assert.NotNil(t, tw.Style())
	assert.Equal(t, StyleBold, *tw.Style())
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

	table.AppendRows([]Row{{}})
	assert.Equal(t, 1, table.Length())
	assert.Equal(t, 0, len(table.rowsFooter))
	assert.Equal(t, 0, len(table.rowsHeader))

	table.AppendRows([]Row{{}})
	assert.Equal(t, 2, table.Length())
	assert.Equal(t, 0, len(table.rowsFooter))
	assert.Equal(t, 0, len(table.rowsHeader))
}

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

func TestTable_RenderColored(t *testing.T) {
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

func TestTable_RenderCSV(t *testing.T) {
	tw := NewWriter()
	tw.AppendHeader(testHeader)
	tw.AppendRows(testRows)
	tw.AppendRow(testRowMultiLine)
	tw.AppendRow(testRowTabs)
	tw.AppendFooter(testFooter)

	expectedOut := `#,First Name,Last Name,Salary,
1,Arya,Stark,3000,
20,Jon,Snow,2000,"You know nothing\, Jon Snow!"
300,Tyrion,Lannister,5000,
0,Winter,Is,0,"Coming.
The North Remembers!"
0,Valar,Morghulis,0,    
,,Total,10000,`

	assert.Equal(t, expectedOut, tw.RenderCSV())
}

func TestTable_RenderHTML(t *testing.T) {
	tw := NewWriter()
	tw.AppendHeader(testHeader)
	tw.AppendRows(testRows)
	tw.AppendRow(testRowMultiLine)
	tw.AppendFooter(testFooter)
	tw.SetVAlign([]text.VAlign{
		text.VAlignDefault, text.VAlignDefault, text.VAlignDefault, text.VAlignBottom, text.VAlignBottom})

	expectedOut := `<table class="go-pretty-table">
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

	table.SetAlign([]text.Align{})
	assert.NotNil(t, table.align)

	table.AppendRows(testRows)
	table.AppendRow(testRowMultiLine)
	table.SetAlign([]text.Align{text.AlignDefault, text.AlignLeft, text.AlignLeft, text.AlignRight, text.AlignRight})

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

	table.SetColors([]text.Colors{testColors1, testColors2})
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

	table.SetColorsFooter([]text.Colors{testColors1, testColors2})
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

	table.SetColorsHeader([]text.Colors{testColors1, testColors2})
	assert.Empty(t, table.colors)
	assert.Empty(t, table.colorsFooter)
	assert.NotEmpty(t, table.colorsHeader)
	assert.Equal(t, 2, len(table.colorsHeader))
}

func TestTable_SetHTMLCSSClass(t *testing.T) {
	table := Table{}
	table.AppendRow(testRows[0])
	expectedHTML := `<table class="` + DefaultHTMLCSSClass + `">
  <tbody>
  <tr>
    <td align="right">1</td>
    <td>Arya</td>
    <td>Stark</td>
    <td align="right">3000</td>
  </tr>
  </tbody>
</table>`
	assert.Equal(t, "", table.htmlCSSClass)
	assert.Equal(t, expectedHTML, table.RenderHTML())

	table.SetHTMLCSSClass(testCSSClass)
	assert.Equal(t, testCSSClass, table.htmlCSSClass)
	assert.Equal(t, strings.Replace(expectedHTML, DefaultHTMLCSSClass, testCSSClass, -1), table.RenderHTML())
}

func TestTable_SetOutputMirror(t *testing.T) {
	table := Table{}
	table.AppendRow(testRows[0])
	expectedOut := `+---+------+-------+------+
| 1 | Arya | Stark | 3000 |
+---+------+-------+------+`
	assert.Equal(t, nil, table.outputMirror)
	assert.Equal(t, expectedOut, table.Render())

	mockOutputMirror := &myMockOutputMirror{}
	table.SetOutputMirror(mockOutputMirror)
	assert.Equal(t, mockOutputMirror, table.outputMirror)
	assert.Equal(t, expectedOut, table.Render())
	assert.Equal(t, expectedOut, mockOutputMirror.mirroredOutput)
}

func TestTable_SetVAlign(t *testing.T) {
	table := Table{}
	assert.Nil(t, table.vAlign)

	table.SetVAlign([]text.VAlign{})
	assert.NotNil(t, table.vAlign)

	table.AppendRow(testRowMultiLine)
	table.SetVAlign([]text.VAlign{text.VAlignTop, text.VAlignMiddle, text.VAlignBottom, text.VAlignDefault})

	expectedOut := `+---+--------+----+---+----------------------+
| 0 | Winter |    | 0 | Coming.              |
|   |        | Is |   | The North Remembers! |
+---+--------+----+---+----------------------+`

	assert.Equal(t, expectedOut, table.Render())
}

func TestTable_SetStyle(t *testing.T) {
	table := Table{}
	assert.Nil(t, table.Style())

	table.SetStyle(StyleDefault)
	assert.NotNil(t, table.Style())
	assert.Equal(t, &StyleDefault, table.Style())
}

func TestTable_ShowBorder(t *testing.T) {
	table := Table{}
	assert.False(t, table.disableBorder)

	table.ShowBorder(false)
	assert.True(t, table.disableBorder)

	table.AppendRow(testRows[0])
	out := table.Render()

	assert.NotEmpty(t, out)
	assert.Equal(t, 0, strings.Count(out, "\n"))
	assert.Equal(t, " 1 | Arya | Stark | 3000 ", out)
}

func TestTable_ShowSeparators(t *testing.T) {
	table := Table{}
	assert.False(t, table.enableSeparators)

	table.ShowSeparators(true)
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
