package gopretty

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/fatih/color"
)

//go:generate mockgen -source table.go -destination mocks/table_mocks.go

const (
	// DefaultHTMLCSSClass stores the css-class to use when none-provided via
	// SetHTMLCSSClass(cssClass string).
	DefaultHTMLCSSClass = "go-pretty-table"
)

// TableRow defines a single row in the Table.
type TableRow []interface{}

// Table helps print a 2-dimensional array in a human readable pretty-table.
type Table struct {
	// alignment describes the horizontal-alignment for each column
	alignment []Alignment
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
	// rows stores the rows that make up the body
	rows []TableRow
	// rowsFooter stores the rows that make up the footer
	rowsFooter []TableRow
	// rowsHeader stores the rows that make up the header
	rowsHeader []TableRow
	// rowSeparator is a dummy row that contains the separator columns (dashes
	// that make up the separator between header/body/footer
	rowSeparator TableRow
	// style contains all the strings used to draw the table, and more
	style *TableStyle
	// vAlignment describes the vertical-alignment for each column
	vAlignment []VAlignment
}

// TableWriter declares the interfaces implemented by Table.
type TableWriter interface {
	AppendFooter(row TableRow)
	AppendHeader(row TableRow)
	AppendRow(row TableRow)
	AppendRows(rows []TableRow)
	DisableBorder()
	EnableSeparators()
	Length() int
	Render() string
	RenderCSV() string
	RenderHTML() string
	SetAlignment(alignment []Alignment)
	SetCaption(caption string)
	SetColors(colors []TextColor)
	SetColorsFooter(colors []TextColor)
	SetColorsHeader(colors []TextColor)
	SetHTMLCSSClass(cssClass string)
	SetStyle(style TableStyle)
	SetVAlignment(vAlignment []VAlignment)
	Style() *TableStyle
}

// NewTableWriter initializes and returns a TableWriter.
func NewTableWriter() TableWriter {
	return &Table{}
}

// AppendFooter appends the row to the List of footers to render.
func (t *Table) AppendFooter(row TableRow) {
	t.rowsFooter = append(t.rowsFooter, t.analyzeRowAndStringify(row, false, true))
}

// AppendHeader appends the row to the List of headers to render.
func (t *Table) AppendHeader(row TableRow) {
	t.rowsHeader = append(t.rowsHeader, t.analyzeRowAndStringify(row, true, false))
}

// AppendRow appends the row to the List of rows to render.
func (t *Table) AppendRow(row TableRow) {
	t.rows = append(t.rows, t.analyzeRowAndStringify(row, false, false))
}

// AppendRows appends the rows to the List of rows to render.
func (t *Table) AppendRows(rows []TableRow) {
	for _, row := range rows {
		t.AppendRow(row)
	}
}

// DisableBorder disables drawing the border around the Table. Example:
//     # │ FIRST NAME │ LAST NAME │ SALARY │
//  ─────┼────────────┼───────────┼────────┼─────────────────────────────
//     1 │ Arya       │ Stark     │   3000 │
//    20 │ Jon        │ Snow      │   2000 │ You know nothing, Jon Snow!
//   300 │ Tyrion     │ Lannister │   5000 │
//  ─────┼────────────┼───────────┼────────┼─────────────────────────────
//       │            │ TOTAL     │  10000 │
func (t *Table) DisableBorder() {
	t.disableBorder = true
}

// EnableSeparators enables drawing separators between each row. Example:
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
func (t *Table) EnableSeparators() {
	t.enableSeparators = true
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

	// render rows one-by-one from header to footer
	var out strings.Builder
	t.renderRowSeparator(&out, true, false)
	if len(t.rowsHeader) > 0 {
		t.renderRows(&out, t.rowsHeader, t.colorsHeader, true, t.style.CaseHeader)
		t.renderRowSeparator(&out, false, false)
	}
	t.renderRows(&out, t.rows, t.colors, false, t.style.CaseRows)
	if len(t.rowsFooter) > 0 {
		t.renderRowSeparator(&out, false, false)
		t.renderRows(&out, t.rowsFooter, t.colorsFooter, true, t.style.CaseFooter)
	}
	t.renderRowSeparator(&out, false, true)

	// if a caption was set, append that to the output
	if t.caption != "" {
		out.WriteRune('\n')
		out.WriteString(t.caption)
	}

	return out.String()
}

// RenderCSV renders the Table in a CSV format. Example:
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
	return out.String()
}

// RenderHTML renders the Table in a HTML format. Example:
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

	return out.String()
}

// SetAlignment sets the horizontal-alignment for each column in all the rows.
func (t *Table) SetAlignment(alignment []Alignment) {
	t.alignment = alignment
}

// SetCaption sets the text to be rendered just below the table. This will not
// show up when the Table is rendered as a CSV.
func (t *Table) SetCaption(caption string) {
	t.caption = caption
}

// SetColors sets the colors for the rows in the Body.
func (t *Table) SetColors(textColors []TextColor) {
	t.colors = make([]*color.Color, len(textColors))
	for idx, textColor := range textColors {
		t.colors[idx] = textColor.getColorizer()
	}
}

// SetColorsFooter sets the colors for the rows in the Footer.
func (t *Table) SetColorsFooter(textColors []TextColor) {
	t.colorsFooter = make([]*color.Color, len(textColors))
	for idx, textColor := range textColors {
		t.colorsFooter[idx] = textColor.getColorizer()
	}
}

// SetColorsHeader sets the colors for the rows in the Header.
func (t *Table) SetColorsHeader(textColors []TextColor) {
	t.colorsHeader = make([]*color.Color, len(textColors))
	for idx, textColor := range textColors {
		t.colorsHeader[idx] = textColor.getColorizer()
	}
}

// SetHTMLCSSClass sets the the HTML CSS Class to use on the <table> node
// when rendering the Table in HTML format.
func (t *Table) SetHTMLCSSClass(cssClass string) {
	t.htmlCSSClass = cssClass
}

// SetStyle overrides the DefaultStyle with the provided one.
func (t *Table) SetStyle(style TableStyle) {
	t.style = &style
}

// SetVAlignment sets the vertical-alignment for each column in all the rows.
func (t *Table) SetVAlignment(vAlignment []VAlignment) {
	t.vAlignment = vAlignment
}

// Style returns the current style.
func (t *Table) Style() *TableStyle {
	return t.style
}

func (t *Table) analyzeRowAndStringify(row []interface{}, isHeader bool, isFooter bool) []interface{} {
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
	// convert each column to string and figure out the longest column in all
	// available rows
	rowOut := make([]interface{}, len(row))
	for colIdx, col := range row {
		// if the column is a string, mark as so; else, convert to a string
		var colStr string
		if isString(col) {
			colStr = col.(string)
			if !isHeader && !isFooter {
				t.columnIsNumeric[colIdx] = false
			}
		} else {
			colStr = fmt.Sprint(col)
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

func (t *Table) csvFixCommas(str string) string {
	return strings.Replace(str, ",", "\\,", -1)
}

func (t *Table) csvFixDoubleQuotes(str string) string {
	return strings.Replace(str, "\"", "\\\"", -1)
}

func (t *Table) getAlignment(colIdx int) Alignment {
	alignment := AlignmentDefault
	if colIdx < len(t.alignment) {
		alignment = t.alignment[colIdx]
	}
	if alignment == AlignmentDefault && t.columnIsNumeric[colIdx] {
		return AlignmentRight
	}
	return alignment
}

func (t *Table) getVAlignment(colIdx int) VAlignment {
	if colIdx < len(t.vAlignment) {
		return t.vAlignment[colIdx]
	}
	return VAlignmentDefault
}

func (t *Table) init() {
	// pick the default style
	if t.style == nil {
		t.style = &TableStyleDefault
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

func (t *Table) renderColumn(out *strings.Builder, row TableRow, colIdx int, maxColumnLength int, colors []*color.Color, isFirstRow bool, isLastRow bool, isSeparatorRow bool, textCase TextCase) {
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
	// pad the column on the left if not a separator row
	if !isSeparatorRow {
		out.WriteString(t.style.CharPaddingLeft)
	}

	// determine the horizontal-alignment, color.Color and text to use
	alignment := t.getAlignment(colIdx)
	var colorizer *color.Color
	if colIdx < len(colors) {
		colorizer = colors[colIdx]
	}
	var colStr string
	if colIdx < len(row) {
		colStr = row[colIdx].(string)
	}

	// convert-case and then align horizontally
	colStr = alignment.Apply(textCase.Convert(colStr), maxColumnLength)
	// colorize and then render the column content
	if colorizer != nil {
		out.WriteString(colorizer.Sprint(colStr))
	} else {
		out.WriteString(colStr)
	}

	// pad the column on the right if not a separator row
	if !isSeparatorRow {
		out.WriteString(t.style.CharPaddingRight)
	}
}

func (t *Table) renderLine(out *strings.Builder, row TableRow, colors []*color.Color, isFirstRow bool, isLastRow bool, isSeparatorRow bool, textCase TextCase) {
	// do not render the first and last rows (margin rows) if disableBorder isset to true
	if t.disableBorder && (isFirstRow || isLastRow) {
		return
	}

	// grow the strings.Builder by using the horizontal-row-separator length
	// and by the number of columns to account for the column-separator
	out.Grow(t.maxRowLength + t.numColumns + 1)

	// if the output has content, it means that this call is working on line
	// number 2 or more; separate them with a newline
	if out.Len() > 0 {
		out.WriteRune('\n')
	}

	// if disable border is not set; render the first character of the row
	if !t.disableBorder {
		t.renderMarginLeft(out, isFirstRow, isLastRow, isSeparatorRow)
	}

	// render each column with enough padding one-by-one
	for colIdx, maxColumnLength := range t.maxColumnLengths {
		t.renderColumn(out, row, colIdx, maxColumnLength, colors, isFirstRow, isLastRow, isSeparatorRow, textCase)
	}

	// if disable border is not set; render the last character of the row
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

func (t *Table) renderRow(out *strings.Builder, row TableRow, colors []*color.Color, isFirstRow bool, isLastRow bool, isSeparatorRow bool, textCase TextCase) {
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
		t.renderLine(out, row, colors, isFirstRow, isLastRow, isSeparatorRow, textCase)
	} else {
		// convert one row into N # of rows based on maxColLines
		rowLines := make([][]string, len(row))
		for colIdx, col := range row {
			rowLines[colIdx] = t.getVAlignment(colIdx).ApplyStr(col.(string), maxColLines)
		}
		for colLineIdx := 0; colLineIdx < maxColLines; colLineIdx++ {
			rowLine := make(TableRow, len(rowLines))
			for colIdx, colLines := range rowLines {
				rowLine[colIdx] = colLines[colLineIdx]
			}
			t.renderLine(out, rowLine, colors, isFirstRow, isLastRow, isSeparatorRow, textCase)
		}
	}
}

func (t *Table) renderRows(out *strings.Builder, rows []TableRow, colors []*color.Color, isHeaderOrFooter bool, textCase TextCase) {
	for idx, row := range rows {
		t.renderRow(out, row, colors, false, false, false, textCase)
		if t.enableSeparators && !isHeaderOrFooter && idx < len(rows)-1 {
			t.renderRowSeparator(out, false, false)
		}
	}
}

func (t *Table) renderRowsCSV(out *strings.Builder, rows []TableRow) {
	for _, row := range rows {
		t.renderRowCSV(out, row)
	}
}

func (t *Table) renderRowsHTML(out *strings.Builder, rows []TableRow, isHeader bool, isFooter bool) {
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

func (t *Table) renderRowCSV(out *strings.Builder, row TableRow) {
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

func (t *Table) renderRowHTML(out *strings.Builder, row TableRow, isHeader bool, isFooter bool) {
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
		alignment := t.getAlignment(idx).HTMLProperty()
		vAlignment := t.getVAlignment(idx).HTMLProperty()

		// write the row
		out.WriteString("    <")
		out.WriteString(colTagName)
		if alignment != "" {
			out.WriteRune(' ')
			out.WriteString(alignment)
		}
		if vAlignment != "" {
			out.WriteRune(' ')
			out.WriteString(vAlignment)
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
	t.renderLine(out, t.rowSeparator, nil, isFirstRow, isLastRow, true, TextCaseDefault)
}
