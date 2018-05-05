package table

import "strings"

// RenderHTML renders the Table in HTML format. Example:
//  <table class="go-pretty-table">
//    <thead>
//    <tr>
//      <th align="right">#</th>
//      <th>First Name</th>
//      <th>Last Name</th>
//      <th align="right">Salary</th>
//      <th>&nbsp;</th>
//    </tr>
//    </thead>
//    <tbody>
//    <tr>
//      <td align="right">1</td>
//      <td>Arya</td>
//      <td>Stark</td>
//      <td align="right">3000</td>
//      <td>&nbsp;</td>
//    </tr>
//    <tr>
//      <td align="right">20</td>
//      <td>Jon</td>
//      <td>Snow</td>
//      <td align="right">2000</td>
//      <td>You know nothing, Jon Snow!</td>
//    </tr>
//    <tr>
//      <td align="right">300</td>
//      <td>Tyrion</td>
//      <td>Lannister</td>
//      <td align="right">5000</td>
//      <td>&nbsp;</td>
//    </tr>
//    </tbody>
//    <tfoot>
//    <tr>
//      <td align="right">&nbsp;</td>
//      <td>&nbsp;</td>
//      <td>Total</td>
//      <td align="right">10000</td>
//      <td>&nbsp;</td>
//    </tr>
//    </tfoot>
//  </table>
func (t *Table) RenderHTML() string {
	t.initForRender()

	var out strings.Builder
	if t.numColumns > 0 {
		out.WriteString("<table class=\"")
		out.WriteString(t.htmlCSSClass)
		out.WriteString("\">\n")
		t.htmlRenderRows(&out, t.rowsHeader, true, false)
		t.htmlRenderRows(&out, t.rows, false, false)
		t.htmlRenderRows(&out, t.rowsFooter, false, true)
		out.WriteString("</table>")
	}
	return t.render(&out)
}

func (t *Table) htmlRenderRow(out *strings.Builder, row RowStr, isHeader bool, isFooter bool) {
	out.WriteString("  <tr>\n")
	for colIdx := 0; colIdx < t.numColumns; colIdx++ {
		var colStr string
		if colIdx < len(row) {
			colStr = row[colIdx]
		}

		// header uses "th" instead of "td"
		colTagName := "td"
		if isHeader {
			colTagName = "th"
		}

		// determine the HTML "align"/"valign" property values
		align := t.getAlign(colIdx).HTMLProperty()
		vAlign := t.getVAlign(colIdx).HTMLProperty()

		// write the row
		out.WriteString("    <")
		out.WriteString(colTagName)
		if align != "" {
			out.WriteRune(' ')
			out.WriteString(align)
		}
		if vAlign != "" {
			out.WriteRune(' ')
			out.WriteString(vAlign)
		}
		out.WriteString(">")
		if len(colStr) > 0 {
			out.WriteString(strings.Replace(colStr, "\n", "<br/>", -1))
		} else {
			out.WriteString("&nbsp;")
		}
		out.WriteString("</")
		out.WriteString(colTagName)
		out.WriteString(">\n")
	}
	out.WriteString("  </tr>\n")
}

func (t *Table) htmlRenderRows(out *strings.Builder, rows []RowStr, isHeader bool, isFooter bool) {
	if len(rows) > 0 {
		// determine that tag to use based on the type of the row
		rowsTag := "tbody"
		if isHeader {
			rowsTag = "thead"
		} else if isFooter {
			rowsTag = "tfoot"
		}

		var renderedTagOpen, shouldRenderTagClose bool
		for _, row := range rows {
			if len(row) > 0 {
				if !renderedTagOpen {
					out.WriteString("  <")
					out.WriteString(rowsTag)
					out.WriteString(">\n")
					renderedTagOpen = true
				}
				t.htmlRenderRow(out, row, isHeader, isFooter)
				shouldRenderTagClose = true
			}
		}
		if shouldRenderTagClose {
			out.WriteString("  </")
			out.WriteString(rowsTag)
			out.WriteString(">\n")
		}
	}
}
