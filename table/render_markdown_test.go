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

	expectedOut := `| # | First Name | Last Name | Salary |  |
| ---:| --- | --- | ---:| --- |
| 1 | Arya | Stark | 3000 |  |
| 20 | Jon | Snow | 2000 | You know nothing, Jon Snow! |
| 300 | Tyrion | Lannister | 5000 |  |
| 0 | Valar | Morghulis | 0 | Faceless<br/>Men |
| 0 | Valar | Morghulis | 0 | Faceless\|Men |
|  |  | Total | 10000 |  |
_test-caption_`

	assert.Equal(t, expectedOut, tw.RenderMarkdown())
}

func TestTable_RenderMarkdown_Empty(t *testing.T) {
	tw := NewWriter()
	assert.Empty(t, tw.RenderMarkdown())
}
