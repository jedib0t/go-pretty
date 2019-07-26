package table

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTable_RenderMarkdown(t *testing.T) {
	tw := NewWriter()
	tw.AppendHeader(testHeader)
	tw.AppendRows(testRows)
	tw.AppendRow(testRowNewLines)
	tw.AppendRow(testRowPipes)
	tw.AppendFooter(testFooter)
	tw.SetCaption(testCaption)
	tw.SetTitle(testTitle1)

	expectedOut := `# Game of Thrones
| # | First Name | Last Name | Salary |  |
| ---:| --- | --- | ---:| --- |
| 1 | Arya | Stark | 3000 |  |
| 20 | Jon | Snow | 2000 | You know nothing, Jon Snow! |
| 300 | Tyrion | Lannister | 5000 |  |
| 0 | Valar | Morghulis | 0 | Faceless<br/>Men |
| 0 | Valar | Morghulis | 0 | Faceless\|Men |
|  |  | Total | 10000 |  |
_A Song of Ice and Fire_`

	assert.Equal(t, expectedOut, tw.RenderMarkdown())
}

func TestTable_RenderMarkdown_Empty(t *testing.T) {
	tw := NewWriter()
	assert.Empty(t, tw.RenderMarkdown())
}

func TestTable_RendeMarkdown_Sorted(t *testing.T) {
	tw := NewWriter()
	tw.AppendHeader(testHeader)
	tw.AppendRows(testRows)
	tw.AppendRow(Row{11, "Sansa", "Stark", 6000})
	tw.AppendFooter(testFooter)
	tw.SortBy([]SortBy{{Name: "Last Name", Mode: Asc}, {Name: "First Name", Mode: Asc}})

	expectedOut := `| # | First Name | Last Name | Salary |  |
| ---:| --- | --- | ---:| --- |
| 300 | Tyrion | Lannister | 5000 |  |
| 20 | Jon | Snow | 2000 | You know nothing, Jon Snow! |
| 1 | Arya | Stark | 3000 |  |
| 11 | Sansa | Stark | 6000 |  |
|  |  | Total | 10000 |  |`
	assert.Equal(t, expectedOut, tw.RenderMarkdown())
}
