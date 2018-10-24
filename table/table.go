package table

import (
	"fmt"
	"io"
	"strings"
	"unicode/utf8"

	"github.com/jedib0t/go-pretty/text"
)

// Row defines a single row in the Table.
type Row []interface{}

// rowStr defines a single row in the Table comprised of just string objects.
type rowStr []string

// Table helps print a 2-dimensional array in a human readable pretty-table.
type Table struct {
	// align describes the horizontal-align for each column
	align []text.Align
	// alignFooter describes the horizontal-align for each column in the footer
	alignFooter []text.Align
	// alignHeader describes the horizontal-align for each column in the header
	alignHeader []text.Align
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
	colors []text.Colors
	// colorsFooter contains Colorization options for the footer
	colorsFooter []text.Colors
	// colorsHeader contains Colorization options for the header
	colorsHeader []text.Colors
	// columnIsNonNumeric stores if a column contains non-numbers in all rows
	columnIsNonNumeric []bool
	// htmlCSSClass stores the HTML CSS Class to use on the <table> node
	htmlCSSClass string
	// indexColumn stores the number of the column considered as the "index"
	indexColumn int
	// maxColumnLengths stores the length of the longest line in each column
	maxColumnLengths []int
	// maxRowLength stores the length of the longest row
	maxRowLength int
	// numColumns stores the (max.) number of columns seen
	numColumns int
	// numLinesRendered keeps track of the number of lines rendered and helps in
	// paginating long tables
	numLinesRendered int
	// outputMirror stores an io.Writer where the "Render" functions would write
	outputMirror io.Writer
	// pageSize stores the maximum lines to render before rendering the header
	// again (to denote a page break) - useful when you are dealing with really
	// long tables
	pageSize int
	// rows stores the rows that make up the body
	rows []rowStr
	// rowsFooter stores the rows that make up the footer
	rowsFooter []rowStr
	// rowsHeader stores the rows that make up the header
	rowsHeader []rowStr
	// rowSeparator is a dummy row that contains the separator columns (dashes
	// that make up the separator between header/body/footer
	rowSeparator rowStr
	// sortBy stores a map of Column
	sortBy []SortBy
	// style contains all the strings used to draw the table, and more
	style *Style
	// vAlign describes the vertical-align for each column
	vAlign []text.VAlign
	// vAlign describes the vertical-align for each column in the footer
	vAlignFooter []text.VAlign
	// vAlign describes the vertical-align for each column in the header
	vAlignHeader []text.VAlign
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

// SetAlign sets the horizontal-align for each column in the (data) rows.
func (t *Table) SetAlign(align []text.Align) {
	t.align = align
}

// SetAlignFooter sets the horizontal-align for each column in the footer.
func (t *Table) SetAlignFooter(align []text.Align) {
	t.alignFooter = align
}

// SetAlignHeader sets the horizontal-align for each column in the header.
func (t *Table) SetAlignHeader(align []text.Align) {
	t.alignHeader = align
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
	t.autoIndex = autoIndex
}

// SetCaption sets the text to be rendered just below the table. This will not
// show up when the Table is rendered as a CSV.
func (t *Table) SetCaption(format string, a ...interface{}) {
	t.caption = fmt.Sprintf(format, a...)
}

// SetColors sets the colors for the rows in the Body.
func (t *Table) SetColors(colors []text.Colors) {
	t.colors = colors
}

// SetColorsFooter sets the colors for the rows in the Footer.
func (t *Table) SetColorsFooter(colors []text.Colors) {
	t.colorsFooter = colors
}

// SetColorsHeader sets the colors for the rows in the Header.
func (t *Table) SetColorsHeader(colors []text.Colors) {
	t.colorsHeader = colors
}

// SetHTMLCSSClass sets the the HTML CSS Class to use on the <table> node
// when rendering the Table in HTML format.
func (t *Table) SetHTMLCSSClass(cssClass string) {
	t.htmlCSSClass = cssClass
}

// SetIndexColumn sets the given Column # as the column that has the row
// "Number". Valid values range from 1 to N. Note that this is not 0-indexed.
func (t *Table) SetIndexColumn(colNum int) {
	t.indexColumn = colNum
}

// SetOutputMirror sets an io.Writer for all the Render functions to "Write" to
// in addition to returning a string.
func (t *Table) SetOutputMirror(mirror io.Writer) {
	t.outputMirror = mirror
}

// SetPageSize sets the maximum number of lines to render before rendering the
// header rows again. This can be useful when dealing with tables containing a
// long list of rows that can span pages. Please note that the pagination logic
// will not consider Header/Footer lines for paging.
func (t *Table) SetPageSize(numLines int) {
	t.pageSize = numLines
}

// SetStyle overrides the DefaultStyle with the provided one.
func (t *Table) SetStyle(style Style) {
	t.style = &style
}

// SetVAlign sets the vertical-align for each column in all the rows.
func (t *Table) SetVAlign(vAlign []text.VAlign) {
	t.vAlign = vAlign
}

// SetVAlignFooter sets the horizontal-align for each column in the footer.
func (t *Table) SetVAlignFooter(vAlign []text.VAlign) {
	t.vAlignFooter = vAlign
}

// SetVAlignHeader sets the horizontal-align for each column in the header.
func (t *Table) SetVAlignHeader(vAlign []text.VAlign) {
	t.vAlignHeader = vAlign
}

// SortBy sets the rules for sorting the Rows in the order specified. i.e., the
// first SortBy instruction takes precedence over the second and so on. Any
// duplicate instructions on the same column will be discarded while sorting.
func (t *Table) SortBy(sortBy []SortBy) {
	t.sortBy = sortBy
}

// Style returns the current style.
func (t *Table) Style() *Style {
	if t.style == nil {
		tempStyle := StyleDefault
		t.style = &tempStyle
	}
	return t.style
}

func (t *Table) analyzeAndStringify(row Row, isHeader bool, isFooter bool) rowStr {
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
	rowOut := make(rowStr, len(row))
	for colIdx, col := range row {
		// if the column is not a number, keep track of it
		if !isHeader && !isFooter && !t.columnIsNonNumeric[colIdx] && !isNumber(col) {
			t.columnIsNonNumeric[colIdx] = true
		}

		// convert to a string and store it in the row
		var colStr string
		if colStrVal, ok := col.(string); ok {
			colStr = colStrVal
		} else {
			colStr = fmt.Sprint(col)
		}
		if strings.Contains(colStr, "\t") {
			colStr = strings.Replace(colStr, "\t", "    ", -1)
		}
		if strings.Contains(colStr, "\r") {
			colStr = strings.Replace(colStr, "\r", "", -1)
		}
		rowOut[colIdx] = colStr
	}
	return rowOut
}

func (t *Table) getAlign(colIdx int, hint renderHint) text.Align {
	align := text.AlignDefault
	if hint.isHeaderRow {
		if colIdx < len(t.alignHeader) {
			align = t.alignHeader[colIdx]
		}
	} else if hint.isFooterRow {
		if colIdx < len(t.alignFooter) {
			align = t.alignFooter[colIdx]
		}
	} else if colIdx < len(t.align) {
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

func (t *Table) getAutoIndexColumnIDs() rowStr {
	row := make(rowStr, t.numColumns)
	for colIdx, maxColumnLength := range t.maxColumnLengths {
		row[colIdx] = text.AlignCenter.Apply(AutoIndexColumnID(colIdx), maxColumnLength)
	}
	return row
}

func (t *Table) getFormat(hint renderHint) text.Format {
	if hint.isSeparatorRow {
		return text.FormatDefault
	} else if hint.isHeaderRow {
		return t.style.Format.Header
	} else if hint.isFooterRow {
		return t.style.Format.Footer
	}
	return t.style.Format.Row
}

func (t *Table) getRowColors(hint renderHint) []text.Colors {
	if hint.isSeparatorRow {
		return nil
	} else if hint.isHeaderRow {
		return t.colorsHeader
	} else if hint.isFooterRow {
		return t.colorsFooter
	}
	return t.colors
}

func (t *Table) getRowsSorted() []rowStr {
	if t.sortBy == nil || len(t.sortBy) == 0 {
		return t.rows
	}

	sortedRowIndices := t.sortRows(t.rows)
	sortedRows := make([]rowStr, len(t.rows))
	for idx := range t.rows {
		sortedRows[idx] = t.rows[sortedRowIndices[idx]]
	}
	return sortedRows
}

func (t *Table) getBorderColors(hint renderHint) text.Colors {
	if hint.isFooterRow {
		return t.style.Color.Footer
	} else if t.autoIndex {
		return t.style.Color.IndexColumn
	}
	return t.style.Color.Header
}

func (t *Table) getSeparatorColors(hint renderHint) text.Colors {
	if hint.isHeaderRow {
		return t.style.Color.Header
	} else if hint.isFooterRow {
		return t.style.Color.Footer
	} else if hint.isAutoIndexColumn {
		return t.style.Color.IndexColumn
	} else if hint.rowNumber > 0 && hint.rowNumber%2 == 0 {
		return t.style.Color.RowAlternate
	}
	return t.style.Color.Row
}

func (t *Table) getVAlign(colIdx int, hint renderHint) text.VAlign {
	vAlign := text.VAlignDefault
	if hint.isHeaderRow {
		if colIdx < len(t.vAlignHeader) {
			vAlign = t.vAlignHeader[colIdx]
		}
	} else if hint.isFooterRow {
		if colIdx < len(t.vAlignFooter) {
			vAlign = t.vAlignFooter[colIdx]
		}
	} else if colIdx < len(t.vAlign) {
		vAlign = t.vAlign[colIdx]
	}
	return vAlign
}

func (t *Table) initForRender() {
	// pick a default style
	t.Style()

	// auto-index: calc the index column's max length
	t.autoIndexVIndexMaxLength = len(fmt.Sprint(len(t.rows)))

	// find the longest continuous line in the column string
	t.initForRenderMaxColumnLength()

	// generate a separator row and calculate maximum row length
	t.initForRenderRowSeparator()

	// reset the counter for the number of lines rendered
	t.numLinesRendered = 0
}

func (t *Table) initForRenderMaxColumnLength() {
	var findMaxColumnLengths = func(rows []rowStr) {
		for _, row := range rows {
			for colIdx, colStr := range row {
				longestLineLen := text.LongestLineLen(colStr)
				if longestLineLen > t.maxColumnLengths[colIdx] {
					t.maxColumnLengths[colIdx] = longestLineLen
				}
			}
		}
	}

	t.maxColumnLengths = make([]int, t.numColumns)
	findMaxColumnLengths(t.rowsHeader)
	findMaxColumnLengths(t.rows)
	findMaxColumnLengths(t.rowsFooter)

	// restrict the column lengths if any are overthe allowed lengths
	for colIdx := range t.maxColumnLengths {
		allowedLen := t.getAllowedColumnLength(colIdx)
		if allowedLen > 0 && t.maxColumnLengths[colIdx] > allowedLen {
			t.maxColumnLengths[colIdx] = allowedLen
		}
	}
}

func (t *Table) initForRenderRowSeparator() {
	t.maxRowLength = (utf8.RuneCountInString(t.style.Box.MiddleSeparator) * t.numColumns) + 1
	t.rowSeparator = make(rowStr, t.numColumns)
	for colIdx, maxColumnLength := range t.maxColumnLengths {
		maxColumnLength += utf8.RuneCountInString(t.style.Box.PaddingLeft)
		maxColumnLength += utf8.RuneCountInString(t.style.Box.PaddingRight)
		horizontalSeparatorCol := text.RepeatAndTrim(t.style.Box.MiddleHorizontal, maxColumnLength)
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

// renderHint has hints for the Render*() logic
type renderHint struct {
	isAutoIndexColumn bool // auto-index column?
	isBorderBottom    bool // bottom-border?
	isBorderTop       bool // top-border?
	isFirstRow        bool // first-row of header/footer/regular-rows?
	isFooterRow       bool // footer row?
	isHeaderRow       bool // header row?
	isLastLineOfRow   bool // last-line of the current row?
	isLastRow         bool // last-row of header/footer/regular-rows?
	isSeparatorRow    bool // separator row?
	rowLineNumber     int  // the line number for a multi-line row
	rowNumber         int  // the row number/index
}

func (h *renderHint) isRegularRow() bool {
	return !h.isHeaderRow && !h.isFooterRow
}

func (h *renderHint) isLastLineOfLastRow() bool {
	return h.isLastLineOfRow && h.isLastRow
}
