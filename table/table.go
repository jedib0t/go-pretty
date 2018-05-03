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
