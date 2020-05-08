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

func TestTable_RenderMarkdown_HiddenColumns(t *testing.T) {
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
		assert.Equal(t, expectedOut, tw.RenderMarkdown())
	})

	t.Run("first column hidden", func(t *testing.T) {
		tw.SetColumnConfigs(generateColumnConfigsWithHiddenColumns([]int{0}))

		expectedOut := `| First Name | Last Name | Salary |  |
| --- | --- | ---:| --- |
| >>Tyrion | Lannister<< | 5013 |  |
| >>Arya | Stark<< | 3013 |  |
| >>Jon | Snow<< | 2013 | ~You know nothing, Jon Snow!~ |
|  | Total | 10000 |  |`
		assert.Equal(t, expectedOut, tw.RenderMarkdown())
	})

	t.Run("column hidden in the middle", func(t *testing.T) {
		tw.SetColumnConfigs(generateColumnConfigsWithHiddenColumns([]int{1}))

		expectedOut := `| # | Last Name | Salary |  |
| ---:| --- | ---:| --- |
| 307 | Lannister<< | 5013 |  |
| 8 | Stark<< | 3013 |  |
| 27 | Snow<< | 2013 | ~You know nothing, Jon Snow!~ |
|  | Total | 10000 |  |`
		assert.Equal(t, expectedOut, tw.RenderMarkdown())
	})

	t.Run("last column hidden", func(t *testing.T) {
		tw.SetColumnConfigs(generateColumnConfigsWithHiddenColumns([]int{4}))

		expectedOut := `| # | First Name | Last Name | Salary |
| ---:| --- | --- | ---:|
| 307 | >>Tyrion | Lannister<< | 5013 |
| 8 | >>Arya | Stark<< | 3013 |
| 27 | >>Jon | Snow<< | 2013 |
|  |  | Total | 10000 |`
		assert.Equal(t, expectedOut, tw.RenderMarkdown())
	})
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
