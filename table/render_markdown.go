package table

import "strings"

// RenderMarkdown renders the Table in Markdown format. Example:
//  | # | First Name | Last Name | Salary |  |
//  | ---:| --- | --- | ---:| --- |
//  | 1 | Arya | Stark | 3000 |  |
//  | 20 | Jon | Snow | 2000 | You know nothing, Jon Snow! |
//  | 300 | Tyrion | Lannister | 5000 |  |
//  |  |  | Total | 10000 |  |
func (t *Table) RenderMarkdown() string {
	t.initForRender()

	var out strings.Builder
	if t.numColumns > 0 {
		if t.title != "" {
			out.WriteString("# ")
			out.WriteString(t.title)
		}
		t.markdownRenderRows(&out, t.rowsHeader, renderHint{isHeaderRow: true})
		t.markdownRenderRows(&out, t.rows, renderHint{})
		t.markdownRenderRows(&out, t.rowsFooter, renderHint{isFooterRow: true})
		if t.caption != "" {
			out.WriteRune('\n')
			out.WriteRune('_')
			out.WriteString(t.caption)
			out.WriteRune('_')
		}
	}
	return t.render(&out)
}

func (t *Table) markdownRenderRow(out *strings.Builder, row rowStr, hint renderHint) {
	if len(row) > 0 {
		// when working on line number 2 or more, insert a newline first
		if out.Len() > 0 {
			out.WriteRune('\n')
		}

		// render each column up to the max. columns seen in all the rows
		out.WriteRune('|')
		for colIdx := 0; colIdx < t.numColumns; colIdx++ {
			if hint.isSeparatorRow {
				out.WriteString(t.getAlign(colIdx, hint).MarkdownProperty())
			} else {
				var colStr string
				if colIdx < len(row) {
					colStr = row[colIdx]
				}
				out.WriteRune(' ')
				if strings.Contains(colStr, "|") {
					colStr = strings.Replace(colStr, "|", "\\|", -1)
				}
				if strings.Contains(colStr, "\n") {
					colStr = strings.Replace(colStr, "\n", "<br/>", -1)
				}
				out.WriteString(colStr)
				out.WriteRune(' ')
			}
			out.WriteRune('|')
		}
	}
}

func (t *Table) markdownRenderRows(out *strings.Builder, rows []rowStr, hint renderHint) {
	if len(rows) > 0 {
		for idx, row := range rows {
			t.markdownRenderRow(out, row, renderHint{})
			if idx == len(rows)-1 && hint.isHeaderRow {
				t.markdownRenderRow(out, t.rowSeparator, renderHint{isSeparatorRow: true})
			}
		}
	}
}
