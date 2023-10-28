package table

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// RenderCSV renders the Table in CSV format. Example:
//
//	#,First Name,Last Name,Salary,
//	1,Arya,Stark,3000,
//	20,Jon,Snow,2000,"You know nothing\, Jon Snow!"
//	300,Tyrion,Lannister,5000,
//	,,Total,10000,
func (t *Table) RenderCSV() string {
	t.initForRender()

	out := newOutputWriter(t.debugWriter)
	if t.numColumns > 0 {
		if t.title != "" {
			_, _ = out.WriteString(t.title)
		}
		if t.autoIndex && len(t.rowsHeader) == 0 {
			t.csvRenderRow(out, t.getAutoIndexColumnIDs(), renderHint{isAutoIndexRow: true, isHeaderRow: true})
		}
		t.csvRenderRows(out, t.rowsHeader, renderHint{isHeaderRow: true})
		t.csvRenderRows(out, t.rows, renderHint{})
		t.csvRenderRows(out, t.rowsFooter, renderHint{isFooterRow: true})
		if t.caption != "" {
			_, _ = out.WriteRune('\n')
			_, _ = out.WriteString(t.caption)
		}
	}
	return t.render(out)
}

func (t *Table) csvFixCommas(str string) string {
	return strings.Replace(str, ",", "\\,", -1)
}

func (t *Table) csvFixDoubleQuotes(str string) string {
	return strings.Replace(str, "\"", "\\\"", -1)
}

func (t *Table) csvRenderRow(out outputWriter, row rowStr, hint renderHint) {
	// when working on line number 2 or more, insert a newline first
	if out.Len() > 0 {
		_, _ = out.WriteRune('\n')
	}

	// generate the columns to render in CSV format and append to "out"
	for colIdx, colStr := range row {
		// auto-index column
		if colIdx == 0 && t.autoIndex {
			if hint.isRegularRow() {
				_, _ = out.WriteString(fmt.Sprint(hint.rowNumber))
			}
			_, _ = out.WriteRune(',')
		}
		if colIdx > 0 {
			_, _ = out.WriteRune(',')
		}
		if strings.ContainsAny(colStr, "\",\n") {
			_, _ = out.WriteRune('"')
			_, _ = out.WriteString(t.csvFixCommas(t.csvFixDoubleQuotes(colStr)))
			_, _ = out.WriteRune('"')
		} else if utf8.RuneCountInString(colStr) > 0 {
			_, _ = out.WriteString(colStr)
		}
	}
	for colIdx := len(row); colIdx < t.numColumns; colIdx++ {
		_, _ = out.WriteRune(',')
	}
}

func (t *Table) csvRenderRows(out outputWriter, rows []rowStr, hint renderHint) {
	for rowIdx, row := range rows {
		hint.rowNumber = rowIdx + 1
		t.csvRenderRow(out, row, hint)
	}
}
