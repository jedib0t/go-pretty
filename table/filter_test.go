package table

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTable_FilterRows(t *testing.T) {
	t.Run("BasicOperators", func(t *testing.T) {
		t.Run("Equal", func(t *testing.T) {
			table := Table{}
			table.AppendHeader(Row{"#", "First Name", "Last Name", "Salary"})
			table.AppendRows([]Row{
				{1, "Arya", "Stark", 3000},
				{20, "Jon", "Snow", 2000},
				{300, "Tyrion", "Lannister", 5000},
			})
			table.SetStyle(StyleDefault)

			table.FilterBy([]FilterBy{{Number: 2, Operator: Equal, Value: "Jon"}})
			table.initForRenderRows()

			assert.Equal(t, 1, len(table.rowsRawFiltered))
			assert.Equal(t, "Jon", table.rowsRawFiltered[0][1])
		})

		t.Run("NotEqual", func(t *testing.T) {
			table := Table{}
			table.AppendHeader(Row{"#", "First Name", "Last Name", "Salary"})
			table.AppendRows([]Row{
				{1, "Arya", "Stark", 3000},
				{20, "Jon", "Snow", 2000},
				{300, "Tyrion", "Lannister", 5000},
			})
			table.SetStyle(StyleDefault)

			table.FilterBy([]FilterBy{{Number: 2, Operator: NotEqual, Value: "Jon"}})
			table.initForRenderRows()
			assert.Equal(t, 2, len(table.rowsRawFiltered))
		})

		t.Run("GreaterThan", func(t *testing.T) {
			table := Table{}
			table.AppendHeader(Row{"#", "First Name", "Last Name", "Salary"})
			table.AppendRows([]Row{
				{1, "Arya", "Stark", 3000},
				{20, "Jon", "Snow", 2000},
				{300, "Tyrion", "Lannister", 5000},
			})
			table.SetStyle(StyleDefault)

			table.FilterBy([]FilterBy{{Number: 4, Operator: GreaterThan, Value: 2000}})
			table.initForRenderRows()
			assert.Equal(t, 2, len(table.rowsRawFiltered))
		})

		t.Run("LessThan", func(t *testing.T) {
			table := Table{}
			table.AppendHeader(Row{"#", "First Name", "Last Name", "Salary"})
			table.AppendRows([]Row{
				{1, "Arya", "Stark", 3000},
				{20, "Jon", "Snow", 2000},
				{300, "Tyrion", "Lannister", 5000},
			})
			table.SetStyle(StyleDefault)

			table.FilterBy([]FilterBy{{Number: 4, Operator: LessThan, Value: 4000}})
			table.initForRenderRows()
			assert.Equal(t, 2, len(table.rowsRawFiltered))
		})

		t.Run("GreaterThanOrEqual", func(t *testing.T) {
			table := Table{}
			table.AppendHeader(Row{"#", "First Name", "Last Name", "Salary"})
			table.AppendRows([]Row{
				{1, "Arya", "Stark", 3000},
				{20, "Jon", "Snow", 2000},
				{300, "Tyrion", "Lannister", 5000},
			})
			table.SetStyle(StyleDefault)

			table.FilterBy([]FilterBy{{Number: 4, Operator: GreaterThanOrEqual, Value: 3000}})
			table.initForRenderRows()
			assert.Equal(t, 2, len(table.rowsRawFiltered))
		})

		t.Run("LessThanOrEqual", func(t *testing.T) {
			table := Table{}
			table.AppendHeader(Row{"#", "First Name", "Last Name", "Salary"})
			table.AppendRows([]Row{
				{1, "Arya", "Stark", 3000},
				{20, "Jon", "Snow", 2000},
				{300, "Tyrion", "Lannister", 5000},
			})
			table.SetStyle(StyleDefault)

			table.FilterBy([]FilterBy{{Number: 4, Operator: LessThanOrEqual, Value: 3000}})
			table.initForRenderRows()
			assert.Equal(t, 2, len(table.rowsRawFiltered))
		})
	})

	t.Run("StringOperations", func(t *testing.T) {
		t.Run("Contains", func(t *testing.T) {
			table := Table{}
			table.AppendHeader(Row{"#", "First Name", "Last Name", "Salary"})
			table.AppendRows([]Row{
				{1, "Arya", "Stark", 3000},
				{20, "Jon", "Snow", 2000},
				{300, "Tyrion", "Lannister", 5000},
			})
			table.SetStyle(StyleDefault)

			table.FilterBy([]FilterBy{{Number: 2, Operator: Contains, Value: "on"}})
			table.initForRenderRows()
			assert.Equal(t, 2, len(table.rowsRawFiltered))
		})

		t.Run("NotContains", func(t *testing.T) {
			table := Table{}
			table.AppendHeader(Row{"#", "First Name", "Last Name", "Salary"})
			table.AppendRows([]Row{
				{1, "Arya", "Stark", 3000},
				{20, "Jon", "Snow", 2000},
				{300, "Tyrion", "Lannister", 5000},
			})
			table.SetStyle(StyleDefault)

			table.FilterBy([]FilterBy{{Number: 2, Operator: NotContains, Value: "on"}})
			table.initForRenderRows()
			assert.Equal(t, 1, len(table.rowsRawFiltered))
		})

		t.Run("StartsWith", func(t *testing.T) {
			table := Table{}
			table.AppendHeader(Row{"#", "First Name", "Last Name", "Salary"})
			table.AppendRows([]Row{
				{1, "Arya", "Stark", 3000},
				{20, "Jon", "Snow", 2000},
				{300, "Tyrion", "Lannister", 5000},
			})
			table.SetStyle(StyleDefault)

			table.FilterBy([]FilterBy{{Number: 2, Operator: StartsWith, Value: "T"}})
			table.initForRenderRows()
			assert.Equal(t, 1, len(table.rowsRawFiltered))
		})

		t.Run("EndsWith", func(t *testing.T) {
			table := Table{}
			table.AppendHeader(Row{"#", "First Name", "Last Name", "Salary"})
			table.AppendRows([]Row{
				{1, "Arya", "Stark", 3000},
				{20, "Jon", "Snow", 2000},
				{300, "Tyrion", "Lannister", 5000},
			})
			table.SetStyle(StyleDefault)

			table.FilterBy([]FilterBy{{Number: 2, Operator: EndsWith, Value: "a"}})
			table.initForRenderRows()
			assert.Equal(t, 1, len(table.rowsRawFiltered))
		})
	})

	t.Run("RegexOperations", func(t *testing.T) {
		t.Run("RegexMatch", func(t *testing.T) {
			table := Table{}
			table.AppendHeader(Row{"#", "First Name", "Last Name", "Salary"})
			table.AppendRows([]Row{
				{1, "Arya", "Stark", 3000},
				{20, "Jon", "Snow", 2000},
				{300, "Tyrion", "Lannister", 5000},
			})
			table.SetStyle(StyleDefault)

			table.FilterBy([]FilterBy{{Number: 2, Operator: RegexMatch, Value: "^[JT]"}})
			table.initForRenderRows()
			assert.Equal(t, 2, len(table.rowsRawFiltered))
		})

		t.Run("RegexNotMatch", func(t *testing.T) {
			table := Table{}
			table.AppendHeader(Row{"#", "First Name", "Last Name", "Salary"})
			table.AppendRows([]Row{
				{1, "Arya", "Stark", 3000},
				{20, "Jon", "Snow", 2000},
				{300, "Tyrion", "Lannister", 5000},
			})
			table.SetStyle(StyleDefault)

			table.FilterBy([]FilterBy{{Number: 2, Operator: RegexNotMatch, Value: "^[JT]"}})
			table.initForRenderRows()
			assert.Equal(t, 1, len(table.rowsRawFiltered))
		})

		t.Run("InvalidRegex", func(t *testing.T) {
			table := Table{}
			table.AppendHeader(Row{"#", "Name"})
			table.AppendRows([]Row{
				{1, "Arya"},
				{2, "Jon"},
				{3, "Tyrion"},
			})
			table.SetStyle(StyleDefault)

			table.FilterBy([]FilterBy{{Number: 2, Operator: RegexMatch, Value: "[", IgnoreCase: false}})
			table.initForRenderRows()
			assert.Empty(t, table.rowsRawFiltered)
		})
	})

	t.Run("CaseSensitivity", func(t *testing.T) {
		t.Run("Equal", func(t *testing.T) {
			table := Table{}
			table.AppendHeader(Row{"#", "First Name", "Last Name", "Salary"})
			table.AppendRows([]Row{
				{1, "Arya", "Stark", 3000},
				{20, "Jon", "Snow", 2000},
				{300, "Tyrion", "Lannister", 5000},
			})
			table.SetStyle(StyleDefault)

			table.FilterBy([]FilterBy{{Number: 2, Operator: Equal, Value: "JON", IgnoreCase: true}})
			table.initForRenderRows()
			assert.Equal(t, 1, len(table.rowsRawFiltered))
		})

		t.Run("Contains", func(t *testing.T) {
			table := Table{}
			table.AppendHeader(Row{"#", "Name"})
			table.AppendRows([]Row{
				{1, "Arya"},
				{2, "JON"},
				{3, "Tyrion"},
			})
			table.SetStyle(StyleDefault)

			table.FilterBy([]FilterBy{{Number: 2, Operator: Contains, Value: "jon", IgnoreCase: true}})
			table.initForRenderRows()
			assert.Equal(t, 1, len(table.rowsRawFiltered))
		})

		t.Run("StartsWith", func(t *testing.T) {
			table := Table{}
			table.AppendHeader(Row{"#", "Name"})
			table.AppendRows([]Row{
				{1, "Arya"},
				{2, "JON"},
				{3, "Tyrion"},
			})
			table.SetStyle(StyleDefault)

			table.FilterBy([]FilterBy{{Number: 2, Operator: StartsWith, Value: "j", IgnoreCase: true}})
			table.initForRenderRows()
			assert.Equal(t, 1, len(table.rowsRawFiltered))
		})

		t.Run("EndsWith", func(t *testing.T) {
			table := Table{}
			table.AppendHeader(Row{"#", "Name"})
			table.AppendRows([]Row{
				{1, "Arya"},
				{2, "JON"},
				{3, "Tyrion"},
			})
			table.SetStyle(StyleDefault)

			table.FilterBy([]FilterBy{{Number: 2, Operator: EndsWith, Value: "n", IgnoreCase: true}})
			table.initForRenderRows()
			assert.Equal(t, 2, len(table.rowsRawFiltered))
		})

		t.Run("RegexMatch", func(t *testing.T) {
			table := Table{}
			table.AppendHeader(Row{"#", "Name"})
			table.AppendRows([]Row{
				{1, "Arya"},
				{2, "JON"},
				{3, "Tyrion"},
			})
			table.SetStyle(StyleDefault)

			table.FilterBy([]FilterBy{{Number: 2, Operator: RegexMatch, Value: "^j", IgnoreCase: true}})
			table.initForRenderRows()
			assert.Equal(t, 1, len(table.rowsRawFiltered))
		})
	})

	t.Run("NumericTypes", func(t *testing.T) {
		t.Run("Int", func(t *testing.T) {
			table := Table{}
			table.AppendHeader(Row{"#", "Value"})
			table.AppendRows([]Row{
				{1, "100"},
				{2, "200"},
				{3, "300"},
			})
			table.SetStyle(StyleDefault)

			table.FilterBy([]FilterBy{{Number: 2, Operator: GreaterThan, Value: int(150)}})
			table.initForRenderRows()
			assert.Equal(t, 2, len(table.rowsRawFiltered))
		})

		t.Run("Int64", func(t *testing.T) {
			table := Table{}
			table.AppendHeader(Row{"#", "Value"})
			table.AppendRows([]Row{
				{1, "100"},
				{2, "200"},
				{3, "300"},
			})
			table.SetStyle(StyleDefault)

			table.FilterBy([]FilterBy{{Number: 2, Operator: GreaterThan, Value: int64(150)}})
			table.initForRenderRows()
			assert.Equal(t, 2, len(table.rowsRawFiltered))
		})

		t.Run("Float32", func(t *testing.T) {
			table := Table{}
			table.AppendHeader(Row{"#", "Value"})
			table.AppendRows([]Row{
				{1, "100"},
				{2, "200"},
				{3, "300"},
			})
			table.SetStyle(StyleDefault)

			table.FilterBy([]FilterBy{{Number: 2, Operator: LessThan, Value: float32(250.5)}})
			table.initForRenderRows()
			assert.Equal(t, 2, len(table.rowsRawFiltered))
		})

		t.Run("Float64", func(t *testing.T) {
			table := Table{}
			table.AppendHeader(Row{"#", "Value"})
			table.AppendRows([]Row{
				{1, "100.5"},
				{2, "200.5"},
				{3, "300.5"},
			})
			table.SetStyle(StyleDefault)

			table.FilterBy([]FilterBy{{Number: 2, Operator: LessThan, Value: float64(250.0)}})
			table.initForRenderRows()
			assert.Equal(t, 2, len(table.rowsRawFiltered))
		})

		t.Run("StringNumeric", func(t *testing.T) {
			table := Table{}
			table.AppendHeader(Row{"#", "First Name", "Last Name", "Salary"})
			table.AppendRows([]Row{
				{1, "Arya", "Stark", 3000},
				{20, "Jon", "Snow", 2000},
				{300, "Tyrion", "Lannister", 5000},
			})
			table.SetStyle(StyleDefault)

			table.FilterBy([]FilterBy{{Number: 4, Operator: GreaterThan, Value: "2500"}})
			table.initForRenderRows()
			assert.Equal(t, 2, len(table.rowsRawFiltered))
		})

		t.Run("NonNumericCellValue", func(t *testing.T) {
			table := Table{}
			table.AppendHeader(Row{"#", "Name", "Value"})
			table.AppendRows([]Row{
				{1, "Arya", "abc"},
				{2, "Jon", "200"},
			})
			table.SetStyle(StyleDefault)

			table.FilterBy([]FilterBy{{Number: 3, Operator: GreaterThan, Value: 100}})
			table.initForRenderRows()
			assert.Equal(t, 1, len(table.rowsRawFiltered))
		})

		t.Run("NonNumericFilterValue", func(t *testing.T) {
			table := Table{}
			table.AppendHeader(Row{"#", "Value"})
			table.AppendRows([]Row{
				{1, "100"},
				{2, "200"},
			})
			table.SetStyle(StyleDefault)

			type customType struct{ val int }
			table.FilterBy([]FilterBy{{Number: 2, Operator: GreaterThan, Value: customType{val: 150}}})
			table.initForRenderRows()
			assert.Empty(t, table.rowsRawFiltered)
		})

		t.Run("StringFilterValueInvalid", func(t *testing.T) {
			table := Table{}
			table.AppendHeader(Row{"#", "Value"})
			table.AppendRows([]Row{
				{1, "100"},
				{2, "200"},
			})
			table.SetStyle(StyleDefault)

			table.FilterBy([]FilterBy{{Number: 2, Operator: GreaterThan, Value: "not-a-number"}})
			table.initForRenderRows()
			assert.Empty(t, table.rowsRawFiltered)
		})

		t.Run("DefaultCaseParsable", func(t *testing.T) {
			table := Table{}
			table.AppendHeader(Row{"#", "Value"})
			table.AppendRows([]Row{
				{1, "100"},
				{2, "200"},
			})
			table.SetStyle(StyleDefault)

			table.FilterBy([]FilterBy{{Number: 2, Operator: GreaterThan, Value: uint8(150)}})
			table.initForRenderRows()
			assert.Equal(t, 1, len(table.rowsRawFiltered))
		})

		t.Run("DefaultCaseUnparsable", func(t *testing.T) {
			table := Table{}
			table.AppendHeader(Row{"#", "Value"})
			table.AppendRows([]Row{
				{1, "100"},
				{2, "200"},
			})
			table.SetStyle(StyleDefault)

			type unparsableType struct {
				Field string
			}
			table.FilterBy([]FilterBy{{Number: 2, Operator: GreaterThan, Value: unparsableType{Field: "150"}}})
			table.initForRenderRows()
			assert.Empty(t, table.rowsRawFiltered)
		})
	})

	t.Run("EdgeCases", func(t *testing.T) {
		t.Run("InvalidColumn", func(t *testing.T) {
			table := Table{}
			table.AppendHeader(Row{"#", "First Name", "Last Name", "Salary"})
			table.AppendRows([]Row{
				{1, "Arya", "Stark", 3000},
				{20, "Jon", "Snow", 2000},
			})
			table.SetStyle(StyleDefault)

			table.FilterBy([]FilterBy{{Number: 99, Operator: Equal, Value: "test"}})
			table.initForRenderRows()
			assert.Equal(t, 2, len(table.rowsRawFiltered))
		})

		t.Run("InvalidColumnIndexNegative", func(t *testing.T) {
			table := Table{}
			table.AppendHeader(Row{"#", "Name"})
			table.AppendRows([]Row{
				{1, "Arya"},
				{2, "Jon"},
			})
			table.SetStyle(StyleDefault)

			table.FilterBy([]FilterBy{{Number: 0, Operator: Equal, Value: "test"}})
			table.initForRenderRows()
			assert.Equal(t, 2, len(table.rowsRawFiltered))
		})

		t.Run("ColumnIndexOutOfBounds", func(t *testing.T) {
			table := Table{}
			table.AppendHeader(Row{"#", "Name"})
			table.AppendRows([]Row{
				{1, "Arya"},
				{2, "Jon"},
			})
			table.SetStyle(StyleDefault)

			table.FilterBy([]FilterBy{{Number: 10, Operator: Equal, Value: "test"}})
			table.initForRenderRows()
			assert.Equal(t, 2, len(table.rowsRawFiltered))
		})

		t.Run("RowWithFewerColumns", func(t *testing.T) {
			table := Table{}
			table.AppendHeader(Row{"#", "Name", "Salary"})
			table.AppendRows([]Row{
				{1, "Arya", 3000},
				{2, "Jon"},
				{300, "Tyrion", 5000},
			})
			table.SetStyle(StyleDefault)

			table.FilterBy([]FilterBy{{Number: 3, Operator: GreaterThan, Value: 2000}})
			table.initForRenderRows()
			assert.Equal(t, 2, len(table.rowsRawFiltered))
			assert.Equal(t, "Arya", table.rowsRawFiltered[0][1])
			assert.Equal(t, "Tyrion", table.rowsRawFiltered[1][1])
		})

		t.Run("InvalidOperator", func(t *testing.T) {
			table := Table{}
			table.AppendHeader(Row{"#", "Name"})
			table.AppendRows([]Row{
				{1, "Arya"},
				{2, "Jon"},
			})
			table.SetStyle(StyleDefault)

			invalidFilter := FilterBy{Number: 2, Operator: FilterOperator(999), Value: "Arya"}
			table.FilterBy([]FilterBy{invalidFilter})
			table.initForRenderRows()
			assert.Empty(t, table.rowsRawFiltered)
		})

		t.Run("NoFilters", func(t *testing.T) {
			table := Table{}
			table.AppendHeader(Row{"#", "First Name", "Last Name", "Salary"})
			table.AppendRows([]Row{
				{1, "Arya", "Stark", 3000},
				{20, "Jon", "Snow", 2000},
				{300, "Tyrion", "Lannister", 5000},
			})
			table.SetStyle(StyleDefault)

			table.initForRenderRows()
			assert.Equal(t, 3, len(table.rowsRawFiltered))
		})
	})

	t.Run("SpecialFeatures", func(t *testing.T) {
		t.Run("MultipleFilters", func(t *testing.T) {
			table := Table{}
			table.AppendHeader(Row{"#", "First Name", "Last Name", "Salary"})
			table.AppendRows([]Row{
				{1, "Arya", "Stark", 3000},
				{20, "Jon", "Snow", 2000},
				{300, "Tyrion", "Lannister", 5000},
				{400, "Sansa", "Stark", 3000},
			})
			table.SetStyle(StyleDefault)

			table.FilterBy([]FilterBy{
				{Number: 4, Operator: GreaterThanOrEqual, Value: 3000},
				{Number: 3, Operator: Contains, Value: "Stark"},
			})
			table.initForRenderRows()
			assert.Equal(t, 2, len(table.rowsRawFiltered))
		})

		t.Run("CustomFilter", func(t *testing.T) {
			table := Table{}
			table.AppendHeader(Row{"#", "First Name", "Last Name", "Salary"})
			table.AppendRows([]Row{
				{1, "Arya", "Stark", 3000},
				{20, "Jon", "Snow", 2000},
				{300, "Tyrion", "Lannister", 5000},
			})
			table.SetStyle(StyleDefault)

			table.FilterBy([]FilterBy{{
				Number: 2,
				CustomFilter: func(cellValue string) bool {
					return len(cellValue) > 3
				},
			}})
			table.initForRenderRows()
			assert.Equal(t, 2, len(table.rowsRawFiltered))
		})

		t.Run("ByName", func(t *testing.T) {
			table := Table{}
			table.AppendHeader(Row{"#", "First Name", "Last Name", "Salary"})
			table.AppendRows([]Row{
				{1, "Arya", "Stark", 3000},
				{20, "Jon", "Snow", 2000},
				{300, "Tyrion", "Lannister", 5000},
			})
			table.SetStyle(StyleDefault)

			table.FilterBy([]FilterBy{{Name: "First Name", Operator: Equal, Value: "Jon"}})
			table.initForRenderRows()
			assert.Equal(t, 1, len(table.rowsRawFiltered))
		})

		t.Run("WithSorting", func(t *testing.T) {
			table := Table{}
			table.AppendHeader(Row{"#", "First Name", "Last Name", "Salary"})
			table.AppendRows([]Row{
				{1, "Arya", "Stark", 3000},
				{20, "Jon", "Snow", 2000},
				{300, "Tyrion", "Lannister", 5000},
				{400, "Sansa", "Stark", 2500},
			})
			table.SetStyle(StyleDefault)

			table.FilterBy([]FilterBy{{Number: 4, Operator: GreaterThan, Value: 2000}})
			table.SortBy([]SortBy{{Number: 4, Mode: AscNumeric}})
			table.initForRenderRows()

			assert.Equal(t, 3, len(table.rows))
			assert.Equal(t, "2500", table.rows[0][3])
			assert.Equal(t, "3000", table.rows[1][3])
			assert.Equal(t, "5000", table.rows[2][3])
		})

		t.Run("WithSeparators", func(t *testing.T) {
			table := Table{}
			table.AppendHeader(Row{"#", "Name", "Salary"})
			table.AppendRow(Row{1, "Arya", 3000})
			table.AppendSeparator()
			table.AppendRow(Row{20, "Jon", 2000})
			table.AppendRow(Row{300, "Tyrion", 5000})
			table.AppendSeparator()
			table.AppendRow(Row{400, "Sansa", 2500})
			table.SetStyle(StyleDefault)

			table.FilterBy([]FilterBy{{Number: 3, Operator: GreaterThan, Value: 2000}})
			table.initForRenderRows()

			assert.Equal(t, 3, len(table.rowsRawFiltered))
			assert.True(t, table.separators[0])
			assert.True(t, table.separators[1])
			assert.False(t, table.separators[2])
		})
	})
}

func TestTable_calculateNumColumnsFromRaw(t *testing.T) {
	t.Run("WithFooter", func(t *testing.T) {
		table := Table{}
		table.AppendHeader(Row{"#", "Name"})
		table.AppendRows([]Row{
			{1, "Arya"},
			{2, "Jon"},
		})
		table.AppendFooter(Row{"Total", "Count", "Extra", "Column"})
		table.SetStyle(StyleDefault)
		table.initForRenderRows()

		table.calculateNumColumnsFromRaw()
		assert.Equal(t, 4, table.numColumns)
	})

	t.Run("WithMultipleHeaders", func(t *testing.T) {
		table := Table{}
		table.AppendHeader(Row{"#", "Name"})
		table.AppendHeader(Row{"ID", "First", "Last", "Extra"})
		table.AppendRows([]Row{
			{1, "Arya"},
			{2, "Jon"},
		})
		table.SetStyle(StyleDefault)
		table.initForRenderRows()

		table.calculateNumColumnsFromRaw()
		assert.Equal(t, 4, table.numColumns)
	})
}
