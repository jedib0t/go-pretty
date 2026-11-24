package table

import (
	"fmt"
	"strings"
)

func (t *Table) RenderTSV() string {
	t.initForRender(renderModeTSV)

	var out strings.Builder

	if t.numColumns > 0 {
		out.Grow(t.estimatedRenderLength())
		if t.title != "" {
			out.WriteString(t.title)
		}

		if t.autoIndex && len(t.rowsHeader) == 0 {
			t.tsvRenderRow(&out, t.getAutoIndexColumnIDs(), renderHint{isAutoIndexRow: true, isHeaderRow: true})
		}

		t.tsvRenderRows(&out, t.rowsHeader, renderHint{isHeaderRow: true})
		t.tsvRenderRows(&out, t.rows, renderHint{})
		t.tsvRenderRows(&out, t.rowsFooter, renderHint{isFooterRow: true})

		if t.caption != "" {
			out.WriteRune('\n')
			out.WriteString(t.caption)
		}
	}

	return t.render(&out)
}

func (t *Table) tsvRenderRow(out *strings.Builder, row rowStr, hint renderHint) {
	// when working on line number 2 or more, insert a newline first; the
	// lines are counted as a lone empty field renders as an empty line
	if out.Len() > 0 || t.numLinesRendered > 0 {
		out.WriteRune('\n')
	}
	t.numLinesRendered++

	// auto-index column
	if t.autoIndex {
		if hint.isRegularRow() {
			fmt.Fprint(out, hint.rowNumber)
		}
		out.WriteRune('\t')
	}

	// rows with fewer columns than the table are padded with empty fields
	for colIdx := 0; colIdx < t.numColumns; colIdx++ {
		if colIdx > 0 {
			out.WriteRune('\t')
		}
		var col string
		if colIdx < len(row) {
			col = row[colIdx]
		}
		if strings.ContainsAny(col, "\t\n\"") || strings.Contains(col, "    ") {
			col = strings.ReplaceAll(col, "\"", "\"\"") // fix double-quotes
			fmt.Fprintf(out, "\"%s\"", col)
		} else {
			out.WriteString(col)
		}
	}
}

func (t *Table) tsvRenderRows(out *strings.Builder, rows []rowStr, hint renderHint) {
	for idx, row := range rows {
		hint.rowNumber = idx + 1
		t.tsvRenderRow(out, row, hint)
	}
}
