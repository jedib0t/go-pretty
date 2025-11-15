package table

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTable_sortRows_MissingCells(t *testing.T) {
	table := Table{}
	table.AppendRows([]Row{
		{1, "Arya", "Stark", 3000, 9},
		{11, "Sansa", "Stark", 3000},
		{20, "Jon", "Snow", 2000, "You know nothing, Jon Snow!"},
		{300, "Tyrion", "Lannister", 5000, 7},
	})
	table.SetStyle(StyleDefault)
	table.initForRenderRows()

	// sort by "First Name"
	table.SortBy([]SortBy{{Number: 5, Mode: Asc}})
	assert.Equal(t, []int{1, 3, 0, 2}, table.getSortedRowIndices())
}

func TestTable_sortRows_InvalidMode(t *testing.T) {
	table := Table{}
	table.AppendRows([]Row{
		{1, "Arya", "Stark", 3000},
		{11, "Sansa", "Stark", 3000},
		{20, "Jon", "Snow", 2000, "You know nothing, Jon Snow!"},
		{300, "Tyrion", "Lannister", 5000},
	})
	table.SetStyle(StyleDefault)
	table.initForRenderRows()

	// sort by "First Name"
	table.SortBy([]SortBy{{Number: 2, Mode: AscNumeric}})
	assert.Equal(t, []int{0, 1, 2, 3}, table.getSortedRowIndices())
}

func TestTable_sortRows_MixedMode(t *testing.T) {
	table := Table{}
	table.AppendHeader(Row{"#", "First Name", "Last Name", "Salary"})
	table.AppendRows([]Row{
		/* 0 */ {1, "Arya", "Stark", 3000, 4},
		/* 1 */ {11, "Sansa", "Stark", 3000},
		/* 2 */ {20, "Jon", "Snow", 2000, "You know nothing, Jon Snow!"},
		/* 3 */ {300, "Tyrion", "Lannister", 5000, -7.54},
		/* 4 */ {400, "Jamie", "Lannister", 5000, nil},
		/* 5 */ {500, "Tywin", "Lannister", 5000, "-7.540"},
	})
	table.SetStyle(StyleDefault)
	table.initForRenderRows()

	// sort by nothing
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5}, table.getSortedRowIndices())

	// sort column #5 in Ascending order alphabetically and then numerically
	table.SortBy([]SortBy{{Number: 5, Mode: AscAlphaNumeric}, {Number: 1, Mode: AscNumeric}})
	assert.Equal(t, []int{1, 4, 2, 3, 5, 0}, table.getSortedRowIndices())

	// sort column #5 in Ascending order numerically and then alphabetically
	table.SortBy([]SortBy{{Number: 5, Mode: AscNumericAlpha}, {Number: 1, Mode: AscNumeric}})
	assert.Equal(t, []int{3, 5, 0, 1, 4, 2}, table.getSortedRowIndices())

	// sort column #5 in Descending order alphabetically and then numerically
	table.SortBy([]SortBy{{Number: 5, Mode: DscAlphaNumeric}, {Number: 1, Mode: AscNumeric}})
	assert.Equal(t, []int{2, 4, 1, 0, 3, 5}, table.getSortedRowIndices())

	// sort column #5 in Descending order numerically and then alphabetically
	table.SortBy([]SortBy{{Number: 5, Mode: DscNumericAlpha}, {Number: 1, Mode: AscNumeric}})
	assert.Equal(t, []int{0, 3, 5, 2, 4, 1}, table.getSortedRowIndices())
}

func TestTable_sortRows_WithName(t *testing.T) {
	table := Table{}
	table.AppendHeader(Row{"#", "First Name", "Last Name", "Salary"})
	table.AppendRows([]Row{
		{1, "Arya", "Stark", 3000},
		{11, "Sansa", "Stark", 3000},
		{20, "Jon", "Snow", 2000, "You know nothing, Jon Snow!"},
		{300, "Tyrion", "Lannister", 5000},
	})
	table.SetStyle(StyleDefault)
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
	table.SetStyle(StyleDefault)
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

//gocyclo:ignore
func TestTable_sortRows_CustomLess(t *testing.T) {
	t.Run("BasicAscending", func(t *testing.T) {
		table := Table{}
		table.AppendRows([]Row{
			{"zebra", "apple"},
			{"apple", "banana"},
			{"banana", "cherry"},
		})
		table.SetStyle(StyleDefault)
		table.initForRenderRows()

		// CustomLess that sorts alphabetically in ascending order
		table.SortBy([]SortBy{{
			Number: 1,
			CustomLess: func(iVal string, jVal string) int {
				if iVal < jVal {
					return -1
				}
				if iVal > jVal {
					return 1
				}
				return 0
			},
		}})
		assert.Equal(t, []int{1, 2, 0}, table.getSortedRowIndices())
	})

	t.Run("BasicDescending", func(t *testing.T) {
		table := Table{}
		table.AppendRows([]Row{
			{"apple", "banana"},
			{"zebra", "cherry"},
			{"banana", "apple"},
		})
		table.SetStyle(StyleDefault)
		table.initForRenderRows()

		// CustomLess that sorts alphabetically in descending order
		table.SortBy([]SortBy{{
			Number: 1,
			CustomLess: func(iVal string, jVal string) int {
				if iVal > jVal {
					return -1
				}
				if iVal < jVal {
					return 1
				}
				return 0
			},
		}})
		// For descending: when iVal > jVal, return -1 means i comes before j
		// So zebra > apple means CustomLess("zebra", "apple") = -1, zebra comes first
		// Expected order: zebra, banana, apple -> indices {1, 2, 0}
		assert.Equal(t, []int{1, 2, 0}, table.getSortedRowIndices())
	})

	t.Run("EqualValuesContinueToNextColumn", func(t *testing.T) {
		table := Table{}
		table.AppendRows([]Row{
			{"same", "zebra"},
			{"same", "apple"},
			{"same", "banana"},
		})
		table.SetStyle(StyleDefault)
		table.initForRenderRows()

		// First column: all equal (returns 0), second column: alphabetical ascending
		table.SortBy([]SortBy{
			{
				Number: 1,
				CustomLess: func(iVal string, jVal string) int {
					// All values are "same", so always return 0
					return 0
				},
			},
			{
				Number: 2,
				Mode:   Asc,
			},
		})
		assert.Equal(t, []int{1, 2, 0}, table.getSortedRowIndices())
	})

	t.Run("NumericSorting", func(t *testing.T) {
		table := Table{}
		table.AppendRows([]Row{
			{"100", "item"},
			{"2", "item"},
			{"10", "item"},
			{"1", "item"},
		})
		table.SetStyle(StyleDefault)
		table.initForRenderRows()

		// CustomLess that sorts numerically
		table.SortBy([]SortBy{{
			Number: 1,
			CustomLess: func(iVal string, jVal string) int {
				iNum, iErr := strconv.Atoi(iVal)
				jNum, jErr := strconv.Atoi(jVal)
				if iErr != nil || jErr != nil {
					// Fallback to string comparison if not numeric
					if iVal < jVal {
						return -1
					}
					if iVal > jVal {
						return 1
					}
					return 0
				}
				if iNum < jNum {
					return -1
				}
				if iNum > jNum {
					return 1
				}
				return 0
			},
		}})
		// Numeric sort: 1, 2, 10, 100 -> indices {3, 1, 2, 0}
		assert.Equal(t, []int{3, 1, 2, 0}, table.getSortedRowIndices())
	})

	t.Run("EmptyStrings", func(t *testing.T) {
		table := Table{}
		table.AppendRows([]Row{
			{"", "item"},
			{"zebra", "item"},
			{"apple", "item"},
		})
		table.SetStyle(StyleDefault)
		table.initForRenderRows()

		// CustomLess that treats empty strings as less than non-empty
		table.SortBy([]SortBy{{
			Number: 1,
			CustomLess: func(iVal string, jVal string) int {
				if iVal == "" && jVal != "" {
					return -1
				}
				if iVal != "" && jVal == "" {
					return 1
				}
				if iVal < jVal {
					return -1
				}
				if iVal > jVal {
					return 1
				}
				return 0
			},
		}})
		assert.Equal(t, []int{0, 2, 1}, table.getSortedRowIndices())
	})

	t.Run("MissingCells", func(t *testing.T) {
		table := Table{}
		table.AppendRows([]Row{
			{"zebra"},
			{"apple", "extra"},
			{"banana"},
		})
		table.SetStyle(StyleDefault)
		table.initForRenderRows()

		// CustomLess that handles missing cells (empty string)
		table.SortBy([]SortBy{{
			Number: 2,
			CustomLess: func(iVal string, jVal string) int {
				// Treat empty as less than non-empty
				if iVal == "" && jVal != "" {
					return -1
				}
				if iVal != "" && jVal == "" {
					return 1
				}
				if iVal < jVal {
					return -1
				}
				if iVal > jVal {
					return 1
				}
				return 0
			},
		}})
		assert.Equal(t, []int{0, 2, 1}, table.getSortedRowIndices())
	})

	t.Run("WithColumnName", func(t *testing.T) {
		table := Table{}
		table.AppendHeader(Row{"Value", "Other"})
		table.AppendRows([]Row{
			{"zebra", "item"},
			{"apple", "item"},
			{"banana", "item"},
		})
		table.SetStyle(StyleDefault)
		table.initForRenderRows()

		// CustomLess with column name
		table.SortBy([]SortBy{{
			Name: "Value",
			CustomLess: func(iVal string, jVal string) int {
				if iVal < jVal {
					return -1
				}
				if iVal > jVal {
					return 1
				}
				return 0
			},
		}})
		// Should sort: apple, banana, zebra -> indices {1, 2, 0}
		assert.Equal(t, []int{1, 2, 0}, table.getSortedRowIndices())
	})

	t.Run("MultiColumnComplex", func(t *testing.T) {
		table := Table{}
		table.AppendRows([]Row{
			{"same", "zebra", "100"},
			{"same", "apple", "50"},
			{"same", "apple", "200"},
			{"different", "banana", "75"},
		})
		table.SetStyle(StyleDefault)
		table.initForRenderRows()

		// First column: custom logic (group by prefix)
		// Second column: alphabetical
		// Third column: numeric
		table.SortBy([]SortBy{
			{
				Number: 1,
				CustomLess: func(iVal string, jVal string) int {
					// "same" values come first
					if iVal == "same" && jVal != "same" {
						return -1
					}
					if iVal != "same" && jVal == "same" {
						return 1
					}
					return 0
				},
			},
			{
				Number: 2,
				Mode:   Asc,
			},
			{
				Number: 3,
				Mode:   AscNumeric,
			},
		})
		// Expected: "same" rows first, then sorted by col2, then col3
		// Rows: 0={"same","zebra","100"}, 1={"same","apple","50"}, 2={"same","apple","200"}, 3={"different","banana","75"}
		// Expected order: 1, 2, 0 (all "same" sorted by col2 then col3), then 3
		assert.Equal(t, []int{1, 2, 0, 3}, table.getSortedRowIndices())
	})

	t.Run("AllReturnValues", func(t *testing.T) {
		table := Table{}
		table.AppendRows([]Row{
			{"c", "item"}, // will be greater
			{"a", "item"}, // will be less
			{"b", "item"}, // will be equal to itself, but in between
		})
		table.SetStyle(StyleDefault)
		table.initForRenderRows()

		// CustomLess that explicitly returns -1, 0, 1
		table.SortBy([]SortBy{{
			Number: 1,
			CustomLess: func(iVal string, jVal string) int {
				if iVal == "a" {
					return -1 // a is always less
				}
				if iVal == "c" {
					return 1 // c is always greater
				}
				if jVal == "a" {
					return 1 // b > a
				}
				if jVal == "c" {
					return -1 // b < c
				}
				return 0 // b == b
			},
		}})
		assert.Equal(t, []int{1, 2, 0}, table.getSortedRowIndices())
	})

	t.Run("CaseInsensitiveCustom", func(t *testing.T) {
		table := Table{}
		table.AppendRows([]Row{
			{"Zebra", "item"},
			{"apple", "item"},
			{"Banana", "item"},
		})
		table.SetStyle(StyleDefault)
		table.initForRenderRows()

		// CustomLess with case-insensitive comparison
		table.SortBy([]SortBy{{
			Number: 1,
			CustomLess: func(iVal string, jVal string) int {
				iLower := strings.ToLower(iVal)
				jLower := strings.ToLower(jVal)
				if iLower < jLower {
					return -1
				}
				if iLower > jLower {
					return 1
				}
				// If case-insensitive equal, compare case-sensitive
				if iVal < jVal {
					return -1
				}
				if iVal > jVal {
					return 1
				}
				return 0
			},
		}})
		// Rows: 0={"Zebra"}, 1={"apple"}, 2={"Banana"}
		// Case-insensitive: apple < Banana < Zebra
		// Expected: {1, 2, 0}
		assert.Equal(t, []int{1, 2, 0}, table.getSortedRowIndices())
	})
}
