package table

import (
	"fmt"
	"os"
	"testing"

	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/stretchr/testify/assert"
)

func TestTable_RenderHTML(t *testing.T) {
	tw := NewWriter()
	tw.AppendHeader(testHeader)
	tw.AppendRows(testRows)
	tw.AppendRow(testRowMultiLine)
	tw.AppendFooter(testFooter)
	tw.SetColumnConfigs([]ColumnConfig{
		{Name: "Salary", VAlign: text.VAlignBottom},
		{Number: 5, VAlign: text.VAlignBottom},
	})
	tw.SetTitle(testTitle1)
	tw.SetCaption(testCaption)
	tw.Style().Title = TitleOptions{
		Align:  text.AlignLeft,
		Colors: text.Colors{text.BgBlack, text.Bold, text.FgHiBlue},
		Format: text.FormatTitle,
	}

	expectedOut := `<table class="go-pretty-table">
  <caption class="title" align="left" class="bg-black bold fg-hi-blue">Game Of Thrones</caption>
  <thead>
  <tr>
    <th align="right">#</th>
    <th>First Name</th>
    <th>Last Name</th>
    <th align="right">Salary</th>
    <th>&nbsp;</th>
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
    <td valign="bottom">Coming.<br/>The North Remembers!<br/>This is known.</td>
  </tr>
  </tbody>
  <tfoot>
  <tr>
    <td align="right">&nbsp;</td>
    <td>&nbsp;</td>
    <td>Total</td>
    <td align="right">10000</td>
    <td>&nbsp;</td>
  </tr>
  </tfoot>
  <caption class="caption" style="caption-side: bottom;">A Song of Ice and Fire</caption>
</table>`

	assert.Equal(t, expectedOut, tw.RenderHTML())
}

func TestTable_RenderHTML_AutoIndex(t *testing.T) {
	tw := NewWriter()
	for rowIdx := 0; rowIdx < 3; rowIdx++ {
		row := make(Row, 3)
		for colIdx := 0; colIdx < 3; colIdx++ {
			row[colIdx] = fmt.Sprintf("%s%d", AutoIndexColumnID(colIdx), rowIdx+1)
		}
		tw.AppendRow(row)
	}
	for rowIdx := 0; rowIdx < 1; rowIdx++ {
		row := make(Row, 3)
		for colIdx := 0; colIdx < 3; colIdx++ {
			row[colIdx] = AutoIndexColumnID(colIdx) + "F"
		}
		tw.AppendFooter(row)
	}
	tw.SetOutputMirror(os.Stdout)
	tw.SetAutoIndex(true)
	tw.SetStyle(StyleLight)

	expectedOut := `<table class="go-pretty-table">
  <thead>
  <tr>
    <th>&nbsp;</th>
    <th align="center">A</th>
    <th align="center">B</th>
    <th align="center">C</th>
  </tr>
  </thead>
  <tbody>
  <tr>
    <td align="right">1</td>
    <td>A1</td>
    <td>B1</td>
    <td>C1</td>
  </tr>
  <tr>
    <td align="right">2</td>
    <td>A2</td>
    <td>B2</td>
    <td>C2</td>
  </tr>
  <tr>
    <td align="right">3</td>
    <td>A3</td>
    <td>B3</td>
    <td>C3</td>
  </tr>
  </tbody>
  <tfoot>
  <tr>
    <td>&nbsp;</td>
    <td>AF</td>
    <td>BF</td>
    <td>CF</td>
  </tr>
  </tfoot>
</table>`
	assert.Equal(t, expectedOut, tw.RenderHTML())
}

func TestTable_RenderHTML_Colored(t *testing.T) {
	tw := NewWriter()
	tw.AppendHeader(testHeader)
	tw.AppendRows(testRows)
	tw.AppendRow(testRowMultiLine)
	tw.AppendFooter(testFooter)
	tw.SetCaption(testCaption)
	tw.SetTitle(testTitle1)
	tw.Style().HTML.CSSClass = "go-pretty-table-colored"
	colorsBlackOnWhite := text.Colors{text.BgWhite, text.FgBlack}
	tw.SetColumnConfigs([]ColumnConfig{
		{
			Name:         "#",
			Colors:       text.Colors{text.Bold},
			ColorsHeader: colorsBlackOnWhite,
		}, {
			Name:         "First Name",
			Colors:       text.Colors{text.FgCyan},
			ColorsHeader: colorsBlackOnWhite,
		}, {
			Name:         "Last Name",
			Colors:       text.Colors{text.FgMagenta},
			ColorsHeader: colorsBlackOnWhite,
			ColorsFooter: colorsBlackOnWhite,
		}, {
			Name:         "Salary",
			Colors:       text.Colors{text.FgYellow},
			ColorsHeader: colorsBlackOnWhite,
			ColorsFooter: colorsBlackOnWhite,
			VAlign:       text.VAlignBottom,
		}, {
			Number:       5,
			Colors:       text.Colors{text.FgBlack},
			ColorsHeader: colorsBlackOnWhite,
			VAlign:       text.VAlignBottom,
		},
	})

	expectedOut := `<table class="go-pretty-table-colored">
  <caption class="title">Game of Thrones</caption>
  <thead>
  <tr>
    <th align="right" class="bg-white fg-black">#</th>
    <th class="bg-white fg-black">First Name</th>
    <th class="bg-white fg-black">Last Name</th>
    <th align="right" class="bg-white fg-black">Salary</th>
    <th class="bg-white fg-black">&nbsp;</th>
  </tr>
  </thead>
  <tbody>
  <tr>
    <td align="right" class="bold">1</td>
    <td class="fg-cyan">Arya</td>
    <td class="fg-magenta">Stark</td>
    <td align="right" class="fg-yellow" valign="bottom">3000</td>
    <td class="fg-black" valign="bottom">&nbsp;</td>
  </tr>
  <tr>
    <td align="right" class="bold">20</td>
    <td class="fg-cyan">Jon</td>
    <td class="fg-magenta">Snow</td>
    <td align="right" class="fg-yellow" valign="bottom">2000</td>
    <td class="fg-black" valign="bottom">You know nothing, Jon Snow!</td>
  </tr>
  <tr>
    <td align="right" class="bold">300</td>
    <td class="fg-cyan">Tyrion</td>
    <td class="fg-magenta">Lannister</td>
    <td align="right" class="fg-yellow" valign="bottom">5000</td>
    <td class="fg-black" valign="bottom">&nbsp;</td>
  </tr>
  <tr>
    <td align="right" class="bold">0</td>
    <td class="fg-cyan">Winter</td>
    <td class="fg-magenta">Is</td>
    <td align="right" class="fg-yellow" valign="bottom">0</td>
    <td class="fg-black" valign="bottom">Coming.<br/>The North Remembers!<br/>This is known.</td>
  </tr>
  </tbody>
  <tfoot>
  <tr>
    <td align="right">&nbsp;</td>
    <td>&nbsp;</td>
    <td class="bg-white fg-black">Total</td>
    <td align="right" class="bg-white fg-black">10000</td>
    <td>&nbsp;</td>
  </tr>
  </tfoot>
  <caption class="caption" style="caption-side: bottom;">A Song of Ice and Fire</caption>
</table>`
	assert.Equal(t, expectedOut, tw.RenderHTML())
}

func TestTable_RenderHTML_CustomStyle(t *testing.T) {
	tw := NewWriter()
	tw.AppendHeader(testHeader)
	tw.AppendRow(Row{1, "Arya", "Stark", 3000, "<a href=\"https://duckduckgo.com/?q=arya+stark+not+today\">Not today.</a>"})
	tw.AppendRow(Row{1, "Jon", "Snow", 2000, "You know\nnothing,\nJon Snow!"})
	tw.AppendRow(Row{300, "Tyrion", "Lannister", 5000})
	tw.AppendFooter(testFooter)
	tw.SetAutoIndex(true)
	tw.Style().HTML = HTMLOptions{
		CSSClass:    "game-of-thrones",
		EmptyColumn: "<!-- test -->&nbsp;",
		EscapeText:  false,
		Newline:     "<!-- newline -->",
	}
	tw.SetOutputMirror(os.Stdout)

	expectedOut := `<table class="game-of-thrones">
  <thead>
  <tr>
    <th><!-- test -->&nbsp;</th>
    <th align="right">#</th>
    <th>First Name</th>
    <th>Last Name</th>
    <th align="right">Salary</th>
    <th><!-- test -->&nbsp;</th>
  </tr>
  </thead>
  <tbody>
  <tr>
    <td align="right">1</td>
    <td align="right">1</td>
    <td>Arya</td>
    <td>Stark</td>
    <td align="right">3000</td>
    <td><a href="https://duckduckgo.com/?q=arya+stark+not+today">Not today.</a></td>
  </tr>
  <tr>
    <td align="right">2</td>
    <td align="right">1</td>
    <td>Jon</td>
    <td>Snow</td>
    <td align="right">2000</td>
    <td>You know<!-- newline -->nothing,<!-- newline -->Jon Snow!</td>
  </tr>
  <tr>
    <td align="right">3</td>
    <td align="right">300</td>
    <td>Tyrion</td>
    <td>Lannister</td>
    <td align="right">5000</td>
    <td><!-- test -->&nbsp;</td>
  </tr>
  </tbody>
  <tfoot>
  <tr>
    <td><!-- test -->&nbsp;</td>
    <td align="right"><!-- test -->&nbsp;</td>
    <td><!-- test -->&nbsp;</td>
    <td>Total</td>
    <td align="right">10000</td>
    <td><!-- test -->&nbsp;</td>
  </tr>
  </tfoot>
</table>`
	assert.Equal(t, expectedOut, tw.RenderHTML())
}

func TestTable_RenderHTML_Empty(t *testing.T) {
	tw := NewWriter()
	assert.Empty(t, tw.RenderHTML())
}

func TestTable_RenderHTML_HiddenColumns(t *testing.T) {
	tw := NewWriter()
	tw.AppendHeader(testHeader)
	tw.AppendRows(testRows)
	tw.AppendFooter(testFooter)

	// ensure sorting is done before hiding the columns
	tw.SortBy([]SortBy{
		{Name: "Salary", Mode: DscNumeric},
	})

	t.Run("every column hidden", func(t *testing.T) {
		tw.SetColumnConfigs(generateColumnConfigsWithHiddenColumns([]int{0, 1, 2, 3, 4}))

		expectedOut := ``
		assert.Equal(t, expectedOut, tw.RenderHTML())
	})

	t.Run("first column hidden", func(t *testing.T) {
		tw.SetColumnConfigs(generateColumnConfigsWithHiddenColumns([]int{0}))

		expectedOut := `<table class="go-pretty-table">
  <thead>
  <tr>
    <th>First Name</th>
    <th>Last Name</th>
    <th align="right">Salary</th>
    <th>&nbsp;</th>
  </tr>
  </thead>
  <tbody>
  <tr>
    <td>&gt;&gt;Tyrion</td>
    <td>Lannister&lt;&lt;</td>
    <td align="right">5013</td>
    <td>&nbsp;</td>
  </tr>
  <tr>
    <td>&gt;&gt;Arya</td>
    <td>Stark&lt;&lt;</td>
    <td align="right">3013</td>
    <td>&nbsp;</td>
  </tr>
  <tr>
    <td>&gt;&gt;Jon</td>
    <td>Snow&lt;&lt;</td>
    <td align="right">2013</td>
    <td>~You know nothing, Jon Snow!~</td>
  </tr>
  </tbody>
  <tfoot>
  <tr>
    <td>&nbsp;</td>
    <td>Total</td>
    <td align="right">10000</td>
    <td>&nbsp;</td>
  </tr>
  </tfoot>
</table>`
		assert.Equal(t, expectedOut, tw.RenderHTML())
	})

	t.Run("column hidden in the middle", func(t *testing.T) {
		tw.SetColumnConfigs(generateColumnConfigsWithHiddenColumns([]int{1}))

		expectedOut := `<table class="go-pretty-table">
  <thead>
  <tr>
    <th align="right">#</th>
    <th>Last Name</th>
    <th align="right">Salary</th>
    <th>&nbsp;</th>
  </tr>
  </thead>
  <tbody>
  <tr>
    <td align="right">307</td>
    <td>Lannister&lt;&lt;</td>
    <td align="right">5013</td>
    <td>&nbsp;</td>
  </tr>
  <tr>
    <td align="right">8</td>
    <td>Stark&lt;&lt;</td>
    <td align="right">3013</td>
    <td>&nbsp;</td>
  </tr>
  <tr>
    <td align="right">27</td>
    <td>Snow&lt;&lt;</td>
    <td align="right">2013</td>
    <td>~You know nothing, Jon Snow!~</td>
  </tr>
  </tbody>
  <tfoot>
  <tr>
    <td align="right">&nbsp;</td>
    <td>Total</td>
    <td align="right">10000</td>
    <td>&nbsp;</td>
  </tr>
  </tfoot>
</table>`
		assert.Equal(t, expectedOut, tw.RenderHTML())
	})

	t.Run("last column hidden", func(t *testing.T) {
		tw.SetColumnConfigs(generateColumnConfigsWithHiddenColumns([]int{4}))

		expectedOut := `<table class="go-pretty-table">
  <thead>
  <tr>
    <th align="right">#</th>
    <th>First Name</th>
    <th>Last Name</th>
    <th align="right">Salary</th>
  </tr>
  </thead>
  <tbody>
  <tr>
    <td align="right">307</td>
    <td>&gt;&gt;Tyrion</td>
    <td>Lannister&lt;&lt;</td>
    <td align="right">5013</td>
  </tr>
  <tr>
    <td align="right">8</td>
    <td>&gt;&gt;Arya</td>
    <td>Stark&lt;&lt;</td>
    <td align="right">3013</td>
  </tr>
  <tr>
    <td align="right">27</td>
    <td>&gt;&gt;Jon</td>
    <td>Snow&lt;&lt;</td>
    <td align="right">2013</td>
  </tr>
  </tbody>
  <tfoot>
  <tr>
    <td align="right">&nbsp;</td>
    <td>&nbsp;</td>
    <td>Total</td>
    <td align="right">10000</td>
  </tr>
  </tfoot>
</table>`
		assert.Equal(t, expectedOut, tw.RenderHTML())
	})
}

func TestTable_RenderHTML_Sorted(t *testing.T) {
	tw := NewWriter()
	tw.AppendHeader(testHeader)
	tw.AppendRows(testRows)
	tw.AppendRow(Row{11, "Sansa", "Stark", 6000})
	tw.AppendFooter(testFooter)
	tw.SortBy([]SortBy{{Name: "Last Name", Mode: Asc}, {Name: "First Name", Mode: Asc}})

	expectedOut := `<table class="go-pretty-table">
  <thead>
  <tr>
    <th align="right">#</th>
    <th>First Name</th>
    <th>Last Name</th>
    <th align="right">Salary</th>
    <th>&nbsp;</th>
  </tr>
  </thead>
  <tbody>
  <tr>
    <td align="right">300</td>
    <td>Tyrion</td>
    <td>Lannister</td>
    <td align="right">5000</td>
    <td>&nbsp;</td>
  </tr>
  <tr>
    <td align="right">20</td>
    <td>Jon</td>
    <td>Snow</td>
    <td align="right">2000</td>
    <td>You know nothing, Jon Snow!</td>
  </tr>
  <tr>
    <td align="right">1</td>
    <td>Arya</td>
    <td>Stark</td>
    <td align="right">3000</td>
    <td>&nbsp;</td>
  </tr>
  <tr>
    <td align="right">11</td>
    <td>Sansa</td>
    <td>Stark</td>
    <td align="right">6000</td>
    <td>&nbsp;</td>
  </tr>
  </tbody>
  <tfoot>
  <tr>
    <td align="right">&nbsp;</td>
    <td>&nbsp;</td>
    <td>Total</td>
    <td align="right">10000</td>
    <td>&nbsp;</td>
  </tr>
  </tfoot>
</table>`
	assert.Equal(t, expectedOut, tw.RenderHTML())
}
