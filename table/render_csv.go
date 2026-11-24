package table

import (
	"fmt"
	"strings"
)

// RenderCSV renders the Table in CSV format. Example:
//
//	#,First Name,Last Name,Salary,
//	1,Arya,Stark,3000,
//	20,Jon,Snow,2000,"You know nothing, Jon Snow!"
//	300,Tyrion,Lannister,5000,
//	,,Total,10000,
func (t *Table) RenderCSV() string {
	t.initForRender(renderModeCSV)

	var out strings.Builder
	if t.numColumns > 0 {
		out.Grow(t.estimatedRenderLength())
		if t.title != "" {
			out.WriteString(t.title)
		}
		if t.autoIndex && len(t.rowsHeader) == 0 {
			t.csvRenderRow(&out, t.getAutoIndexColumnIDs(), renderHint{isAutoIndexRow: true, isHeaderRow: true})
		}
		t.csvRenderRows(&out, t.rowsHeader, renderHint{isHeaderRow: true})
		t.csvRenderRows(&out, t.rows, renderHint{})
		t.csvRenderRows(&out, t.rowsFooter, renderHint{isFooterRow: true})
		if t.caption != "" {
			out.WriteRune('\n')
			out.WriteString(t.caption)
		}
	}
	return t.render(&out)
}

// csvEscapeDoubleQuotes escapes double-quotes inside a quoted field by
// doubling them, as mandated by RFC 4180.
func (t *Table) csvEscapeDoubleQuotes(str string) string {
	return strings.ReplaceAll(str, "\"", "\"\"")
}

// csvNeedsQuotes returns true if the field has to be enclosed in double-quotes
// to be read back as-is. A lone empty field is quoted as CSV readers would
// otherwise skip its record as a blank line.
func (t *Table) csvNeedsQuotes(str string) bool {
	return strings.ContainsAny(str, "\",\n") ||
		(str == "" && t.numColumns == 1 && !t.autoIndex)
}

// csvProtectField neutralizes fields that spreadsheet applications would
// interpret as formulas (=, +, -, @, tab or CR prefix) by prefixing them with
// a single-quote.
func (t *Table) csvProtectField(str string) string {
	switch str[0] {
	case '=', '+', '-', '@', '\t', '\r':
		return "'" + str
	}
	return str
}

func (t *Table) csvRenderRow(out *strings.Builder, row rowStr, hint renderHint) {
	// when working on line number 2 or more, insert a newline first
	if out.Len() > 0 {
		out.WriteRune('\n')
	}

	// auto-index column
	if t.autoIndex {
		if hint.isRegularRow() {
			fmt.Fprint(out, hint.rowNumber)
		}
		out.WriteRune(',')
	}

	// generate the columns to render in CSV format and append to "out"; rows
	// with fewer columns than the table are padded with empty fields
	for colIdx := 0; colIdx < t.numColumns; colIdx++ {
		if colIdx > 0 {
			out.WriteRune(',')
		}
		var colStr string
		if colIdx < len(row) {
			colStr = row[colIdx]
		}
		if t.style.CSV.FieldProtection && colStr != "" {
			colStr = t.csvProtectField(colStr)
		}
		if t.csvNeedsQuotes(colStr) {
			out.WriteRune('"')
			out.WriteString(t.csvEscapeDoubleQuotes(colStr))
			out.WriteRune('"')
		} else {
			out.WriteString(colStr)
		}
	}
}

func (t *Table) csvRenderRows(out *strings.Builder, rows []rowStr, hint renderHint) {
	for rowIdx, row := range rows {
		hint.rowNumber = rowIdx + 1
		t.csvRenderRow(out, row, hint)
	}
}
