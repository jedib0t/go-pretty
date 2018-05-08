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

// RowStr defines a single row in the Table comprised of just string objects.
type RowStr []string

// Table helps print a 2-dimensional array in a human readable pretty-table.
type Table struct {
	// align describes the horizontal-align for each column
	align []text.Align
	// allowedColumnLengths contains the max allowed length for each column
	allowedColumnLengths []int
	// allowedRowLength is the max allowed length for a row (or line of output)
	allowedRowLength int
	// enable automatic indexing of the rows and columns like a spreadsheet?
	autoIndex bool
	// autoIndexVIndexMaxLength denotes the length in chars for the last rownum
	autoIndexVIndexMaxLength int
	// caption stores the text to be rendered just below the table; and doesn't
	// get used when rendered as a CSV
	caption string
	// colors contains Colorization options for the body
	colors []*color.Color
	// colorsFooter contains Colorization options for the footer
	colorsFooter []*color.Color
	// colorsHeader contains Colorization options for the header
	colorsHeader []*color.Color
	// columnIsNonNumeric stores if a column contains non-numbers in all rows
	columnIsNonNumeric []bool
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
	rows []RowStr
	// rowsFooter stores the rows that make up the footer
	rowsFooter []RowStr
	// rowsHeader stores the rows that make up the header
	rowsHeader []RowStr
	// rowSeparator is a dummy row that contains the separator columns (dashes
	// that make up the separator between header/body/footer
	rowSeparator RowStr
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

// SetAlign sets the horizontal-align for each column in all the rows.
func (t *Table) SetAlign(align []text.Align) {
	t.align = align
}

// SetAllowedColumnLengths sets the maximum allowed length for each column in
// all the rows. Columns with content longer than the allowed limit will be
// wrapped to fit the length. Length has to be a positive value to take effect.
func (t *Table) SetAllowedColumnLengths(lengths []int) {
	t.allowedColumnLengths = lengths
}

// SetAllowedRowLength sets the maximum allowed length or a row (or line of
// output) when rendered as a table. Rows that are longer than this limit will
// be "snipped" to the length. Length has to be a positive value to take effect.
func (t *Table) SetAllowedRowLength(length int) {
	t.allowedRowLength = length
}

// SetAutoIndex adds a generated header with columns such as "A", "B", "C", etc.
// and a leading column with the row number similar to what you'd see on any
// spreadsheet application. NOTE: Appending a Header will void this
// functionality.
func (t *Table) SetAutoIndex(autoIndex bool) {
	t.autoIndex = true
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

func (t *Table) analyzeAndStringify(row Row, isHeader bool, isFooter bool) RowStr {
	// update t.numColumns if this row is the longest seen till now
	if len(row) > t.numColumns {
		// init the slice for the first time; and pad it the rest of the time
		if t.numColumns == 0 {
			t.columnIsNonNumeric = make([]bool, len(row))
		} else {
			t.columnIsNonNumeric = append(t.columnIsNonNumeric, make([]bool, len(row)-t.numColumns)...)
		}
		// update t.numColumns
		t.numColumns = len(row)
	}

	// convert each column to string and figure out if it has non-numeric data
	rowOut := make(RowStr, len(row))
	for colIdx, col := range row {
		// if the column is not a number, keep track of it
		if !isHeader && !isFooter && !t.columnIsNonNumeric[colIdx] && !util.IsNumber(col) {
			t.columnIsNonNumeric[colIdx] = true
		}

		// convert to a string and store it in the row
		colStr := util.AsString(col)
		if strings.Contains(colStr, "\t") {
			colStr = strings.Replace(colStr, "\t", "    ", -1)
		}
		rowOut[colIdx] = colStr
	}
	return rowOut
}

func (t *Table) getAlign(colIdx int) text.Align {
	align := text.AlignDefault
	if colIdx < len(t.align) {
		align = t.align[colIdx]
	}
	if align == text.AlignDefault && !t.columnIsNonNumeric[colIdx] {
		align = text.AlignRight
	}
	return align
}

func (t *Table) getAllowedColumnLength(colIdx int) int {
	if colIdx < len(t.allowedColumnLengths) {
		return t.allowedColumnLengths[colIdx]
	}
	return 0
}

func (t *Table) getAutoIndexColumnIDRow() RowStr {
	row := make(RowStr, t.numColumns)
	for colIdx, maxColumnLength := range t.maxColumnLengths {
		row[colIdx] = text.AlignCenter.Apply(util.AutoIndexColumnID(colIdx), maxColumnLength)
	}
	return row
}

func (t *Table) getVAlign(colIdx int) text.VAlign {
	vAlign := text.VAlignDefault
	if colIdx < len(t.vAlign) {
		vAlign = t.vAlign[colIdx]
	}
	return vAlign
}

func (t *Table) initForRender() {
	// pick a default style
	if t.style == nil {
		t.style = &StyleDefault
	}

	// turn off auto-index if a header is found
	if t.autoIndex && len(t.rowsHeader) > 0 {
		t.autoIndex = false
	}
	t.autoIndexVIndexMaxLength = len(fmt.Sprint(len(t.rows)))

	// default to a HTML CSS Class if none-defined
	if t.htmlCSSClass == "" {
		t.htmlCSSClass = DefaultHTMLCSSClass
	}

	// find the longest continuous line in the column string
	t.initForRenderMaxColumnLength()

	// generate a separator row and calculate maximum row length
	t.initForRenderRowSeparator()
}

func (t *Table) initForRenderMaxColumnLength() {
	var findMaxColumnLengths = func(rows []RowStr) {
		for _, row := range rows {
			for colIdx, colStr := range row {
				allowedColumnLength := t.getAllowedColumnLength(colIdx)
				if allowedColumnLength > 0 {
					t.maxColumnLengths[colIdx] = allowedColumnLength
				} else {
					colLongestLineLength := util.GetLongestLineLength(colStr)
					if colLongestLineLength > t.maxColumnLengths[colIdx] {
						t.maxColumnLengths[colIdx] = colLongestLineLength
					}
				}
			}
		}
	}

	t.maxColumnLengths = make([]int, t.numColumns)
	findMaxColumnLengths(t.rowsHeader)
	findMaxColumnLengths(t.rows)
	findMaxColumnLengths(t.rowsFooter)
}

func (t *Table) initForRenderRowSeparator() {
	t.maxRowLength = (utf8.RuneCountInString(t.style.BoxMiddleSeparator) * t.numColumns) + 1
	t.rowSeparator = make(RowStr, t.numColumns)
	for colIdx, maxColumnLength := range t.maxColumnLengths {
		maxColumnLength += utf8.RuneCountInString(t.style.BoxPaddingLeft)
		maxColumnLength += utf8.RuneCountInString(t.style.BoxPaddingRight)
		// TODO: handle case where BoxMiddleHorizontal is longer than 1 rune
		horizontalSeparatorCol := strings.Repeat(t.style.BoxMiddleHorizontal, maxColumnLength)
		t.maxRowLength += maxColumnLength
		t.rowSeparator[colIdx] = horizontalSeparatorCol
	}
}

func (t *Table) render(out *strings.Builder) string {
	outStr := out.String()
	if t.outputMirror != nil && len(outStr) > 0 {
		t.outputMirror.Write([]byte(outStr))
		t.outputMirror.Write([]byte("\n"))
	}
	return outStr
}
