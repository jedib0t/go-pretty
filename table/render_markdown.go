package table

import (
	"fmt"
	"strings"
)

// RenderMarkdown renders the Table in Markdown format. Example:
//
//	| # | First Name | Last Name | Salary |  |
//	| ---:| --- | --- | ---:| --- |
//	| 1 | Arya | Stark | 3000 |  |
//	| 20 | Jon | Snow | 2000 | You know nothing, Jon Snow! |
//	| 300 | Tyrion | Lannister | 5000 |  |
//	|  |  | Total | 10000 |  |
func (t *Table) RenderMarkdown() string {
	t.initForRender()

	out := newOutputWriter(t.debugWriter)
	if t.numColumns > 0 {
		t.markdownRenderTitle(out)
		t.markdownRenderRowsHeader(out)
		t.markdownRenderRows(out, t.rows, renderHint{})
		t.markdownRenderRowsFooter(out)
		t.markdownRenderCaption(out)
	}
	return t.render(out)
}

func (t *Table) markdownRenderCaption(out outputWriter) {
	if t.caption != "" {
		_, _ = out.WriteRune('\n')
		_, _ = out.WriteRune('_')
		_, _ = out.WriteString(t.caption)
		_, _ = out.WriteRune('_')
	}
}

func (t *Table) markdownRenderRow(out outputWriter, row rowStr, hint renderHint) {
	// when working on line number 2 or more, insert a newline first
	if out.Len() > 0 {
		_, _ = out.WriteRune('\n')
	}

	// render each column up to the max. columns seen in all the rows
	_, _ = out.WriteRune('|')
	for colIdx := 0; colIdx < t.numColumns; colIdx++ {
		t.markdownRenderRowAutoIndex(out, colIdx, hint)

		if hint.isSeparatorRow {
			_, _ = out.WriteString(t.getAlign(colIdx, hint).MarkdownProperty())
		} else {
			var colStr string
			if colIdx < len(row) {
				colStr = row[colIdx]
			}
			_, _ = out.WriteRune(' ')
			colStr = strings.ReplaceAll(colStr, "|", "\\|")
			colStr = strings.ReplaceAll(colStr, "\n", "<br/>")
			_, _ = out.WriteString(colStr)
			_, _ = out.WriteRune(' ')
		}
		_, _ = out.WriteRune('|')
	}
}

func (t *Table) markdownRenderRowAutoIndex(out outputWriter, colIdx int, hint renderHint) {
	if colIdx == 0 && t.autoIndex {
		_, _ = out.WriteRune(' ')
		if hint.isSeparatorRow {
			_, _ = out.WriteString("---:")
		} else if hint.isRegularRow() {
			_, _ = out.WriteString(fmt.Sprintf("%d ", hint.rowNumber))
		}
		_, _ = out.WriteRune('|')
	}
}

func (t *Table) markdownRenderRows(out outputWriter, rows []rowStr, hint renderHint) {
	if len(rows) > 0 {
		for idx, row := range rows {
			hint.rowNumber = idx + 1
			t.markdownRenderRow(out, row, hint)

			if idx == len(rows)-1 && hint.isHeaderRow {
				t.markdownRenderRow(out, t.rowSeparator, renderHint{isSeparatorRow: true})
			}
		}
	}
}

func (t *Table) markdownRenderRowsFooter(out outputWriter) {
	t.markdownRenderRows(out, t.rowsFooter, renderHint{isFooterRow: true})
}

func (t *Table) markdownRenderRowsHeader(out outputWriter) {
	if len(t.rowsHeader) > 0 {
		t.markdownRenderRows(out, t.rowsHeader, renderHint{isHeaderRow: true})
	} else if t.autoIndex {
		t.markdownRenderRows(out, []rowStr{t.getAutoIndexColumnIDs()}, renderHint{isAutoIndexRow: true, isHeaderRow: true})
	}
}

func (t *Table) markdownRenderTitle(out outputWriter) {
	if t.title != "" {
		_, _ = out.WriteString("# ")
		_, _ = out.WriteString(t.title)
	}
}
