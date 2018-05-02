package table

import (
	"fmt"
	"io"
	"strings"
	"unicode/utf8"

	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/text"
	"github.com/jedib0t/go-pretty/util"
)

const (
	// DefaultHTMLCSSClass stores the css-class to use when none-provided via
	// SetHTMLCSSClass(cssClass string).
	DefaultHTMLCSSClass = "go-pretty-table"
)

// Row defines a single row in the Table.
type Row []interface{}

// Table helps print a 2-dimensional array in a human readable pretty-table.
type Table struct {
	// align describes the horizontal-align for each column
	align []text.Align
	// caption stores the text to be rendered just below the table; and doesn't
	// get used when rendered as a CSV
	caption string
	// colors contains Colorization options for the body
	colors []*color.Color
	// colorsFooter contains Colorization options for the footer
	colorsFooter []*color.Color
	// colorsHeader contains Colorization options for the header
	colorsHeader []*color.Color
	// columnIsNumeric stores if a column contains numbers in all rows or not
	columnIsNumeric []bool
	// disableBorder disables drawing the border around the table
	disableBorder bool
	// enableSeparators enables drawing separators between each row
	enableSeparators bool
	// htmlCSSClass stores the HTML CSS Class to use on the <table> node
	htmlCSSClass string
	// maxColumnLengths stores the length of the longest line in each column
	maxColumnLengths []int
	// maxRowLength stores the length of the longest row
	maxRowLength int
	// numColumns stores the (max.) number of columns seen
	numColumns int
	// outputMirror stores an io.Writer where the "Render" functions would write
	outputMirror io.Writer
	// rows stores the rows that make up the body
	rows []Row
	// rowsFooter stores the rows that make up the footer
	rowsFooter []Row
	// rowsHeader stores the rows that make up the header
	rowsHeader []Row
	// rowSeparator is a dummy row that contains the separator columns (dashes
	// that make up the separator between header/body/footer
	rowSeparator Row
	// style contains all the strings used to draw the table, and more
	style *Style
	// vAlign describes the vertical-align for each column
	vAlign []text.VAlign
}

// AppendFooter appends the row to the List of footers to render.
func (t *Table) AppendFooter(row Row) {
	t.rowsFooter = append(t.rowsFooter, t.analyzeAndStringify(row, false, true))
}

// AppendHeader appends the row to the List of headers to render.
func (t *Table) AppendHeader(row Row) {
	t.rowsHeader = append(t.rowsHeader, t.analyzeAndStringify(row, true, false))
}

// AppendRow appends the row to the List of rows to render.
func (t *Table) AppendRow(row Row) {
	t.rows = append(t.rows, t.analyzeAndStringify(row, false, false))
}

// AppendRows appends the rows to the List of rows to render.
func (t *Table) AppendRows(rows []Row) {
	for _, row := range rows {
		t.AppendRow(row)
	}
}

// Length returns the number of rows to be rendered.
func (t *Table) Length() int {
	return len(t.rows)
}

// Render renders the Table in a human-readable "pretty" format. Example:
//  ┌─────┬────────────┬───────────┬────────┬─────────────────────────────┐
//  │   # │ FIRST NAME │ LAST NAME │ SALARY │                             │
//  ├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
//  │   1 │ Arya       │ Stark     │   3000 │                             │
//  │  20 │ Jon        │ Snow      │   2000 │ You know nothing, Jon Snow! │
//  │ 300 │ Tyrion     │ Lannister │   5000 │                             │
//  ├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
//  │     │            │ TOTAL     │  10000 │                             │
//  └─────┴────────────┴───────────┴────────┴─────────────────────────────┘
func (t *Table) Render() string {
	t.init()

	var out strings.Builder
	if !t.disableBorder {
		t.renderRowSeparator(&out, true, false)
	}
	if len(t.rowsHeader) > 0 {
		t.renderRows(&out, t.rowsHeader, t.colorsHeader, t.style.FormatHeader)
		t.renderRowSeparator(&out, false, false)
	}
	t.renderRows(&out, t.rows, t.colors, t.style.FormatRows)
	if len(t.rowsFooter) > 0 {
		t.renderRowSeparator(&out, false, false)
		t.renderRows(&out, t.rowsFooter, t.colorsFooter, t.style.FormatFooter)
	}
	if !t.disableBorder {
		t.renderRowSeparator(&out, false, true)
	}
	if t.caption != "" {
		out.WriteRune('\n')
		out.WriteString(t.caption)
	}
	return t.render(&out)
}

// RenderCSV renders the Table in CSV format. Example:
//  #,First Name,Last Name,Salary,
//  1,Arya,Stark,3000,
//  20,Jon,Snow,2000,"You know nothing\, Jon Snow!"
//  300,Tyrion,Lannister,5000,
//  ,,Total,10000,
func (t *Table) RenderCSV() string {
	t.init()

	var out strings.Builder
	t.renderRowsCSV(&out, t.rowsHeader)
	t.renderRowsCSV(&out, t.rows)
	t.renderRowsCSV(&out, t.rowsFooter)
	return t.render(&out)
}

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
	t.init()

	var out strings.Builder
	out.WriteString("<table class=\"")
	out.WriteString(t.htmlCSSClass)
	out.WriteString("\">\n")
	t.renderRowsHTML(&out, t.rowsHeader, true, false)
	t.renderRowsHTML(&out, t.rows, false, false)
	t.renderRowsHTML(&out, t.rowsFooter, false, true)
	out.WriteString("</table>")
	return t.render(&out)
}

// RenderMarkdown renders the Table in Markdown format. Example:
//
func (t *Table) RenderMarkdown() string {
	t.init()

	var out strings.Builder
	t.renderRowsMarkdown(&out, t.rowsHeader, true, false)
	t.renderRowsMarkdown(&out, t.rows, false, false)
	t.renderRowsMarkdown(&out, t.rowsFooter, false, true)
	if t.caption != "" {
		out.WriteRune('_')
		out.WriteString(t.caption)
		out.WriteRune('_')
	}
	return t.render(&out)
}

// SetAlign sets the horizontal-align for each column in all the rows.
func (t *Table) SetAlign(align []text.Align) {
	t.align = align
}

// SetCaption sets the text to be rendered just below the table. This will not
// show up when the Table is rendered as a CSV.
func (t *Table) SetCaption(format string, a ...interface{}) {
	t.caption = fmt.Sprintf(format, a...)
}

// SetColors sets the colors for the rows in the Body.
func (t *Table) SetColors(textColors []text.Colors) {
	t.colors = make([]*color.Color, len(textColors))
	for idx, textColor := range textColors {
		t.colors[idx] = textColor.GetColorizer()
	}
}

// SetColorsFooter sets the colors for the rows in the Footer.
func (t *Table) SetColorsFooter(textColors []text.Colors) {
	t.colorsFooter = make([]*color.Color, len(textColors))
	for idx, textColor := range textColors {
		t.colorsFooter[idx] = textColor.GetColorizer()
	}
}

// SetColorsHeader sets the colors for the rows in the Header.
func (t *Table) SetColorsHeader(textColors []text.Colors) {
	t.colorsHeader = make([]*color.Color, len(textColors))
	for idx, textColor := range textColors {
		t.colorsHeader[idx] = textColor.GetColorizer()
	}
}

// SetHTMLCSSClass sets the the HTML CSS Class to use on the <table> node
// when rendering the Table in HTML format.
func (t *Table) SetHTMLCSSClass(cssClass string) {
	t.htmlCSSClass = cssClass
}

// SetOutputMirror sets an io.Writer for all the Render functions to "Write" to
// in addition to returning a string.
func (t *Table) SetOutputMirror(mirror io.Writer) {
	t.outputMirror = mirror
}

// SetStyle overrides the DefaultStyle with the provided one.
func (t *Table) SetStyle(style Style) {
	t.style = &style
}

// SetVAlign sets the vertical-align for each column in all the rows.
func (t *Table) SetVAlign(vAlign []text.VAlign) {
	t.vAlign = vAlign
}

// ShowBorder enables or disables drawing the border around the Table. Example
// of a table where it is disabled (enabled by default):
//     # │ FIRST NAME │ LAST NAME │ SALARY │
//  ─────┼────────────┼───────────┼────────┼─────────────────────────────
//     1 │ Arya       │ Stark     │   3000 │
//    20 │ Jon        │ Snow      │   2000 │ You know nothing, Jon Snow!
//   300 │ Tyrion     │ Lannister │   5000 │
//  ─────┼────────────┼───────────┼────────┼─────────────────────────────
//       │            │ TOTAL     │  10000 │
func (t *Table) ShowBorder(show bool) {
	t.disableBorder = !show
}

// ShowSeparators enables or disable drawing separators between each row.
// Example of a table where it is enabled (disabled by default):
//  ┌─────┬────────────┬───────────┬────────┬─────────────────────────────┐
//  │   # │ FIRST NAME │ LAST NAME │ SALARY │                             │
//  ├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
//  │   1 │ Arya       │ Stark     │   3000 │                             │
//  ├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
//  │  20 │ Jon        │ Snow      │   2000 │ You know nothing, Jon Snow! │
//  ├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
//  │ 300 │ Tyrion     │ Lannister │   5000 │                             │
//  ├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
//  │     │            │ TOTAL     │  10000 │                             │
//  └─────┴────────────┴───────────┴────────┴─────────────────────────────┘
func (t *Table) ShowSeparators(show bool) {
	t.enableSeparators = show
}

// Style returns the current style.
func (t *Table) Style() *Style {
	return t.style
}

func (t *Table) analyzeAndStringify(row Row, isHeader bool, isFooter bool) Row {
	// update t.numColumns if this row is the longest seen till now
	if len(row) > t.numColumns {
		// pad t.columnIsNumeric with extra "true" values
		columnIsNumeric := make([]bool, len(row)-t.numColumns)
		for idx := range columnIsNumeric {
			columnIsNumeric[idx] = true
		}
		t.columnIsNumeric = append(t.columnIsNumeric, columnIsNumeric...)

		// pad t.maxColumnLengths with extra "0" values
		maxColumnLengths := make([]int, len(row)-t.numColumns)
		t.maxColumnLengths = append(t.maxColumnLengths, maxColumnLengths...)

		// update t.numColumns
		t.numColumns = len(row)
	}

	return t.stringify(row, isHeader, isFooter)
}

func (t *Table) csvFixCommas(str string) string {
	return strings.Replace(str, ",", "\\,", -1)
}

func (t *Table) csvFixDoubleQuotes(str string) string {
	return strings.Replace(str, "\"", "\\\"", -1)
}

func (t *Table) getAlign(colIdx int) text.Align {
	align := text.AlignDefault
	if colIdx < len(t.align) {
		align = t.align[colIdx]
	}
	if align == text.AlignDefault && t.columnIsNumeric[colIdx] {
		align = text.AlignRight
	}
	return align
}

func (t *Table) getVAlign(colIdx int) text.VAlign {
	vAlign := text.VAlignDefault
	if colIdx < len(t.vAlign) {
		vAlign = t.vAlign[colIdx]
	}
	return vAlign
}

func (t *Table) init() {
	// pick the default style
	if t.style == nil {
		t.style = &StyleDefault
	}

	// default to a HTML CSS Class if none-defined
	if t.htmlCSSClass == "" {
		t.htmlCSSClass = DefaultHTMLCSSClass
	}

	// generate a separator row and calculate maximum row length
	t.maxRowLength = (utf8.RuneCountInString(t.style.CharMiddleSeparator) * t.numColumns) + 1
	t.rowSeparator = make([]interface{}, t.numColumns)
	for colIdx, maxColumnLength := range t.maxColumnLengths {
		horizontalSeparatorCol := strings.Repeat(t.style.CharMiddleHorizontal,
			len(t.style.CharPaddingLeft)+maxColumnLength+len(t.style.CharPaddingRight))
		t.maxRowLength += utf8.RuneCountInString(horizontalSeparatorCol)
		t.rowSeparator[colIdx] = horizontalSeparatorCol
	}
}

func (t *Table) render(out *strings.Builder) string {
	outStr := out.String()
	if t.outputMirror != nil {
		t.outputMirror.Write([]byte(outStr))
	}
	return outStr
}

func (t *Table) renderColumn(out *strings.Builder, row Row, colIdx int, maxColumnLength int, colors []*color.Color, isFirstRow bool, isLastRow bool, isSeparatorRow bool, format text.Format) {
	// when working on column number 2 or more, render the column separator
	if colIdx > 0 {
		// type of row determines the character used (top/bottom/separator)
		if isSeparatorRow {
			if isFirstRow {
				out.WriteString(t.style.CharTopSeparator)
			} else if isLastRow {
				out.WriteString(t.style.CharBottomSeparator)
			} else {
				out.WriteString(t.style.CharMiddleSeparator)
			}
		} else {
			out.WriteString(t.style.CharMiddleVertical)
		}
	}

	// extract the text, convert-case if not-empty and align horizontally
	var colStr string
	if colIdx < len(row) {
		colStr = format.Apply(row[colIdx].(string))
	}
	colStr = t.getAlign(colIdx).Apply(colStr, maxColumnLength)

	// pad both sides of the column (when not a separator row)
	if !isSeparatorRow {
		colStr = t.style.CharPaddingLeft + colStr + t.style.CharPaddingRight
	}

	// colorize and then render the column content if a color has been set
	if colIdx < len(colors) && colors[colIdx] != nil {
		out.WriteString(colors[colIdx].Sprint(colStr))
	} else {
		out.WriteString(colStr)
	}
}

func (t *Table) renderLine(out *strings.Builder, row Row, colors []*color.Color, isFirstRow bool, isLastRow bool, isSeparatorRow bool, format text.Format) {
	// grow the strings.Builder by using the horizontal-row-separator length
	// and by the number of columns to account for the column-separator
	out.Grow(t.maxRowLength + t.numColumns + 1)

	// if the output has content, it means that this call is working on line
	// number 2 or more; separate them with a newline
	if out.Len() > 0 {
		out.WriteRune('\n')
	}

	if !t.disableBorder {
		t.renderMarginLeft(out, isFirstRow, isLastRow, isSeparatorRow)
	}
	for colIdx, maxColumnLength := range t.maxColumnLengths {
		t.renderColumn(out, row, colIdx, maxColumnLength, colors, isFirstRow, isLastRow, isSeparatorRow, format)
	}
	if !t.disableBorder {
		t.renderMarginRight(out, isFirstRow, isLastRow, isSeparatorRow)
	}
}

func (t *Table) renderMarginLeft(out *strings.Builder, isFirstRow bool, isLastRow bool, isSeparatorRow bool) {
	// type of row determines the character used (top/bottom/separator/etc.)
	if isFirstRow {
		out.WriteString(t.style.CharTopLeft)
	} else if isLastRow {
		out.WriteString(t.style.CharBottomLeft)
	} else if isSeparatorRow {
		out.WriteString(t.style.CharLeftSeparator)
	} else {
		out.WriteString(t.style.CharLeft)
	}
}

func (t *Table) renderMarginRight(out *strings.Builder, isFirstRow bool, isLastRow bool, isSeparatorRow bool) {
	// type of row determines the character used (top/bottom/separator/etc.)
	if isFirstRow {
		out.WriteString(t.style.CharTopRight)
	} else if isLastRow {
		out.WriteString(t.style.CharBottomRight)
	} else if isSeparatorRow {
		out.WriteString(t.style.CharRightSeparator)
	} else {
		out.WriteString(t.style.CharRight)
	}
}

func (t *Table) renderRow(out *strings.Builder, row Row, colors []*color.Color, isFirstRow bool, isLastRow bool, isSeparatorRow bool, format text.Format) {
	// find the max. # of lines found in all the columns and split each column
	// into a list of strings
	maxColLines := 0
	for _, col := range row {
		numLines := strings.Count(col.(string), "\n") + 1
		if numLines > maxColLines {
			maxColLines = numLines
		}
	}

	// if there is just 1 line in all columns, add the row as such; else split
	// each column into individual lines and render them one-by-one
	if maxColLines == 1 {
		t.renderLine(out, row, colors, isFirstRow, isLastRow, isSeparatorRow, format)
	} else {
		// convert one row into N # of rows based on maxColLines
		rowLines := make([][]string, len(row))
		for colIdx, col := range row {
			rowLines[colIdx] = t.getVAlign(colIdx).ApplyStr(col.(string), maxColLines)
		}
		for colLineIdx := 0; colLineIdx < maxColLines; colLineIdx++ {
			rowLine := make(Row, len(rowLines))
			for colIdx, colLines := range rowLines {
				rowLine[colIdx] = colLines[colLineIdx]
			}
			t.renderLine(out, rowLine, colors, isFirstRow, isLastRow, isSeparatorRow, format)
		}
	}
}

func (t *Table) renderRows(out *strings.Builder, rows []Row, colors []*color.Color, format text.Format) {
	for idx, row := range rows {
		t.renderRow(out, row, colors, false, false, false, format)
		if t.enableSeparators && idx < len(rows)-1 {
			t.renderRowSeparator(out, false, false)
		}
	}
}

func (t *Table) renderRowsCSV(out *strings.Builder, rows []Row) {
	for _, row := range rows {
		t.renderRowCSV(out, row)
	}
}

func (t *Table) renderRowsHTML(out *strings.Builder, rows []Row, isHeader bool, isFooter bool) {
	if len(rows) > 0 {
		// determine that tag to use based on the type of the row
		rowsTag := "tbody"
		if isHeader {
			rowsTag = "thead"
		} else if isFooter {
			rowsTag = "tfoot"
		}

		// render all the rows enclosed by the "rowsTag"
		out.WriteString("  <")
		out.WriteString(rowsTag)
		out.WriteString(">\n")
		for _, row := range rows {
			t.renderRowHTML(out, row, isHeader, isFooter)
		}
		out.WriteString("  </")
		out.WriteString(rowsTag)
		out.WriteString(">\n")
	}
}

func (t *Table) renderRowsMarkdown(out *strings.Builder, rows []Row, isHeader bool, isFooter bool) {
	if len(rows) > 0 {
		for idx, row := range rows {
			t.renderRowMarkdown(out, row, false)
			if idx == len(rows)-1 && isHeader {
				t.renderRowMarkdown(out, t.rowSeparator, true)
			}
		}
	}
}

func (t *Table) renderRowMarkdown(out *strings.Builder, row Row, isSeparator bool) {
	if len(row) > 0 {
		out.WriteRune('|')
		for colIdx := range t.maxColumnLengths {
			if !isSeparator {
				var colStr string
				if colIdx < len(row) {
					colStr = row[colIdx].(string)
				}
				out.WriteRune(' ')
				if strings.Contains(colStr, "|") {
					colStr = strings.Replace(colStr, "|", "\\|", -1)
				}
				if strings.Contains(colStr, "\n") {
					colStr = strings.Replace(colStr, "\n", "<br>", -1)
				}
				out.WriteString(colStr)
				out.WriteRune(' ')
			} else {
				out.WriteString(t.getAlign(colIdx).MarkdownProperty())
			}
			out.WriteRune('|')
		}
		out.WriteRune('\n')
	}
}

func (t *Table) renderRowCSV(out *strings.Builder, row Row) {
	// when working on line number 2 or more, insert a newline first
	if out.Len() > 0 {
		out.WriteRune('\n')
	}

	// generate the columns to render in CSV format and append to "out"
	for idx, col := range row {
		if idx > 0 {
			out.WriteRune(',')
		}
		colStr := col.(string)
		if strings.ContainsAny(colStr, "\",\n") {
			out.WriteRune('"')
			out.WriteString(t.csvFixCommas(t.csvFixDoubleQuotes(colStr)))
			out.WriteRune('"')
		} else if utf8.RuneCountInString(colStr) > 0 {
			out.WriteString(colStr)
		}
	}
	for idx := len(row); idx < t.numColumns; idx++ {
		out.WriteRune(',')
	}
}

func (t *Table) renderRowHTML(out *strings.Builder, row Row, isHeader bool, isFooter bool) {
	out.WriteString("  <tr>\n")
	for idx := range t.maxColumnLengths {
		colStr := ""
		if idx < len(row) {
			colStr = row[idx].(string)
		}

		// header uses "th" instead of "td"
		colTagName := "td"
		if isHeader {
			colTagName = "th"
		}

		// determine the HTML "align"/"valign" property values
		align := t.getAlign(idx).HTMLProperty()
		vAlign := t.getVAlign(idx).HTMLProperty()

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

func (t *Table) renderRowSeparator(out *strings.Builder, isFirstRow bool, isLastRow bool) {
	t.renderLine(out, t.rowSeparator, nil, isFirstRow, isLastRow, true, text.FormatDefault)
}

func (t *Table) stringify(row Row, isHeader bool, isFooter bool) Row {
	// convert each column to string and figure out the longest column in all
	// available rows
	rowOut := make(Row, len(row))
	for colIdx, col := range row {
		// if the column is not a number, keep track of it
		if !isHeader && !isFooter && t.columnIsNumeric[colIdx] && !util.IsNumber(col) {
			t.columnIsNumeric[colIdx] = false
		}

		var colStr string
		if util.IsString(col) {
			colStr = col.(string)
		} else {
			colStr = fmt.Sprint(col)
		}
		if strings.Contains(colStr, "\t") {
			colStr = strings.Replace(colStr, "\t", "    ", -1)
		}
		rowOut[colIdx] = colStr

		// split the string into multiple string based on newlines, and find
		// the longest "line" in it
		for _, colLine := range strings.Split(colStr, "\n") {
			colLineLength := utf8.RuneCountInString(colLine)
			if colLineLength > t.maxColumnLengths[colIdx] {
				t.maxColumnLengths[colIdx] = colLineLength
			}
		}
	}
	return rowOut
}
