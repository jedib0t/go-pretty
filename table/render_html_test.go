package table

import (
	"testing"

	"github.com/jedib0t/go-pretty/text"
	"github.com/stretchr/testify/assert"
)

func TestTable_RenderHTML(t *testing.T) {
	tw := NewWriter()
	tw.AppendHeader(testHeader)
	tw.AppendRows(testRows)
	tw.AppendRow(testRowMultiLine)
	tw.AppendFooter(testFooter)
	tw.SetVAlign([]text.VAlign{
		text.VAlignDefault, text.VAlignDefault, text.VAlignDefault, text.VAlignBottom, text.VAlignBottom,
	})

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

func TestTable_RenderHTML_Empty(t *testing.T) {
	tw := NewWriter()
	assert.Empty(t, tw.RenderHTML())
}
