package table

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTable_sortRows_WithName(t *testing.T) {
	table := Table{}
	table.AppendHeader(Row{"#", "First Name", "Last Name", "Salary"})
	table.AppendRows([]Row{
		{1, "Arya", "Stark", 3000},
		{11, "Sansa", "Stark", 3000},
		{20, "Jon", "Snow", 2000, "You know nothing, Jon Snow!"},
		{300, "Tyrion", "Lannister", 5000},
	})
	table.initForRenderRows()

	// sort by nothing
	assert.Equal(t, []int{0, 1, 2, 3}, table.getSortedRowIndices())

	// sort by "#"
	table.SortBy([]SortBy{{Name: "#", Mode: AscNumeric}})
	assert.Equal(t, []int{0, 1, 2, 3}, table.getSortedRowIndices())

	table.SortBy([]SortBy{{Name: "#", Mode: DscNumeric}})
	assert.Equal(t, []int{3, 2, 1, 0}, table.getSortedRowIndices())

	// sort by First Name, Last Name
	table.SortBy([]SortBy{{Name: "First Name", Mode: Asc}, {Name: "Last Name", Mode: Asc}})
	assert.Equal(t, []int{0, 2, 1, 3}, table.getSortedRowIndices())

	table.SortBy([]SortBy{{Name: "First Name", Mode: Asc}, {Name: "Last Name", Mode: Dsc}})
	assert.Equal(t, []int{0, 2, 1, 3}, table.getSortedRowIndices())

	table.SortBy([]SortBy{{Name: "First Name", Mode: Dsc}, {Name: "Last Name", Mode: Asc}})
	assert.Equal(t, []int{3, 1, 2, 0}, table.getSortedRowIndices())

	table.SortBy([]SortBy{{Name: "First Name", Mode: Dsc}, {Name: "Last Name", Mode: Dsc}})
	assert.Equal(t, []int{3, 1, 2, 0}, table.getSortedRowIndices())

	// sort by Last Name, First Name
	table.SortBy([]SortBy{{Name: "Last Name", Mode: Asc}, {Name: "First Name", Mode: Asc}})
	assert.Equal(t, []int{3, 2, 0, 1}, table.getSortedRowIndices())

	table.SortBy([]SortBy{{Name: "Last Name", Mode: Asc}, {Name: "First Name", Mode: Dsc}})
	assert.Equal(t, []int{3, 2, 1, 0}, table.getSortedRowIndices())

	table.SortBy([]SortBy{{Name: "Last Name", Mode: Dsc}, {Name: "First Name", Mode: Asc}})
	assert.Equal(t, []int{0, 1, 2, 3}, table.getSortedRowIndices())

	table.SortBy([]SortBy{{Name: "Last Name", Mode: Dsc}, {Name: "First Name", Mode: Dsc}})
	assert.Equal(t, []int{1, 0, 2, 3}, table.getSortedRowIndices())

	// sort by Unknown Column
	table.SortBy([]SortBy{{Name: "Last Name", Mode: Dsc}, {Name: "Foo Bar", Mode: Dsc}})
	assert.Equal(t, []int{0, 1, 2, 3}, table.getSortedRowIndices())

	// sort by Salary
	table.SortBy([]SortBy{{Name: "Salary", Mode: AscNumeric}})
	assert.Equal(t, []int{2, 0, 1, 3}, table.getSortedRowIndices())

	table.SortBy([]SortBy{{Name: "Salary", Mode: DscNumeric}})
	assert.Equal(t, []int{3, 0, 1, 2}, table.getSortedRowIndices())

	table.SortBy(nil)
	assert.Equal(t, []int{0, 1, 2, 3}, table.getSortedRowIndices())
}

func TestTable_sortRows_WithoutName(t *testing.T) {
	table := Table{}
	table.AppendRows([]Row{
		{1, "Arya", "Stark", 3000},
		{11, "Sansa", "Stark", 3000},
		{20, "Jon", "Snow", 2000, "You know nothing, Jon Snow!"},
		{300, "Tyrion", "Lannister", 5000},
	})
	table.initForRenderRows()

	// sort by nothing
	assert.Equal(t, []int{0, 1, 2, 3}, table.getSortedRowIndices())

	// sort by "#"
	table.SortBy([]SortBy{{Number: 1, Mode: AscNumeric}})
	assert.Equal(t, []int{0, 1, 2, 3}, table.getSortedRowIndices())

	table.SortBy([]SortBy{{Number: 1, Mode: DscNumeric}})
	assert.Equal(t, []int{3, 2, 1, 0}, table.getSortedRowIndices())

	// sort by First Name, Last Name
	table.SortBy([]SortBy{{Number: 2, Mode: Asc}, {Number: 3, Mode: Asc}})
	assert.Equal(t, []int{0, 2, 1, 3}, table.getSortedRowIndices())

	table.SortBy([]SortBy{{Number: 2, Mode: Asc}, {Number: 3, Mode: Dsc}})
	assert.Equal(t, []int{0, 2, 1, 3}, table.getSortedRowIndices())

	table.SortBy([]SortBy{{Number: 2, Mode: Dsc}, {Number: 3, Mode: Asc}})
	assert.Equal(t, []int{3, 1, 2, 0}, table.getSortedRowIndices())

	table.SortBy([]SortBy{{Number: 2, Mode: Dsc}, {Number: 3, Mode: Dsc}})
	assert.Equal(t, []int{3, 1, 2, 0}, table.getSortedRowIndices())

	// sort by Last Name, First Name
	table.SortBy([]SortBy{{Number: 3, Mode: Asc}, {Number: 2, Mode: Asc}})
	assert.Equal(t, []int{3, 2, 0, 1}, table.getSortedRowIndices())

	table.SortBy([]SortBy{{Number: 3, Mode: Asc}, {Number: 2, Mode: Dsc}})
	assert.Equal(t, []int{3, 2, 1, 0}, table.getSortedRowIndices())

	table.SortBy([]SortBy{{Number: 3, Mode: Dsc}, {Number: 2, Mode: Asc}})
	assert.Equal(t, []int{0, 1, 2, 3}, table.getSortedRowIndices())

	table.SortBy([]SortBy{{Number: 3, Mode: Dsc}, {Number: 2, Mode: Dsc}})
	assert.Equal(t, []int{1, 0, 2, 3}, table.getSortedRowIndices())

	// sort by Unknown Column
	table.SortBy([]SortBy{{Number: 3, Mode: Dsc}, {Number: 99, Mode: Dsc}})
	assert.Equal(t, []int{0, 1, 2, 3}, table.getSortedRowIndices())

	// sort by Salary
	table.SortBy([]SortBy{{Number: 4, Mode: AscNumeric}})
	assert.Equal(t, []int{2, 0, 1, 3}, table.getSortedRowIndices())

	table.SortBy([]SortBy{{Number: 4, Mode: DscNumeric}})
	assert.Equal(t, []int{3, 0, 1, 2}, table.getSortedRowIndices())

	table.SortBy(nil)
	assert.Equal(t, []int{0, 1, 2, 3}, table.getSortedRowIndices())
}
