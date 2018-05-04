package table

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/text"
)

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
	t.initForRender()

	var out strings.Builder
	if t.numColumns > 0 {
		if !t.disableBorder {
			t.renderRowSeparator(&out, true, false)
		}
		if len(t.rowsHeader) > 0 || t.autoIndex {
			if len(t.rowsHeader) > 0 {
				t.renderRows(&out, t.rowsHeader, t.colorsHeader, t.style.FormatHeader)
			} else {
				t.renderRow(&out, 0, t.getAutoIndexColumnIDRow(), t.colorsHeader, false, false, false, text.FormatUpper)
			}
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
	}
	return t.render(&out)
}

func (t *Table) renderColumn(out *strings.Builder, rowNum int, row Row, colIdx int, maxColumnLength int, colors []*color.Color, isFirstRow bool, isLastRow bool, isSeparatorRow bool, format text.Format) {
	// when working on the first column, and autoIndex is true, insert a new
	// column with the row number on it.
	if colIdx == 0 && t.autoIndex {
		if rowNum < 0 {
			numChars := t.autoIndexVIndexMaxLength + utf8.RuneCountInString(t.style.BoxPaddingLeft) +
				utf8.RuneCountInString(t.style.BoxPaddingRight)
			out.WriteString(strings.Repeat(t.style.BoxMiddleHorizontal, numChars))
		} else {
			out.WriteString(t.style.BoxPaddingLeft)
			rowNumStr := fmt.Sprint(rowNum)
			if rowNum == 0 {
				rowNumStr = strings.Repeat(" ", t.autoIndexVIndexMaxLength)
			}
			out.WriteString(text.AlignRight.Apply(rowNumStr, t.autoIndexVIndexMaxLength))
			out.WriteString(t.style.BoxPaddingRight)
		}
		t.renderColumnSeparator(out, isFirstRow, isLastRow, rowNum < 0)
	}

	// when working on column number 2 or more, render the column separator
	if colIdx > 0 {
		t.renderColumnSeparator(out, isFirstRow, isLastRow, isSeparatorRow)
	}

	// extract the text, convert-case if not-empty and align horizontally
	var colStr string
	if colIdx < len(row) {
		colStr = format.Apply(row[colIdx].(string))
	}
	colStr = t.getAlign(colIdx).Apply(colStr, maxColumnLength)

	// pad both sides of the column (when not a separator row)
	if !isSeparatorRow {
		colStr = t.style.BoxPaddingLeft + colStr + t.style.BoxPaddingRight
	}

	// colorize and then render the column content if a color has been set
	if colIdx < len(colors) && colors[colIdx] != nil {
		out.WriteString(colors[colIdx].Sprint(colStr))
	} else {
		out.WriteString(colStr)
	}
}

func (t *Table) renderColumnSeparator(out *strings.Builder, isFirstRow bool, isLastRow bool, isSeparatorRow bool) {
	// type of row determines the character used (top/bottom/separator)
	if isSeparatorRow {
		if isFirstRow {
			out.WriteString(t.style.BoxTopSeparator)
		} else if isLastRow {
			out.WriteString(t.style.BoxBottomSeparator)
		} else {
			out.WriteString(t.style.BoxMiddleSeparator)
		}
	} else {
		out.WriteString(t.style.BoxMiddleVertical)
	}
}

func (t *Table) renderLine(out *strings.Builder, rowNum int, row Row, colors []*color.Color, isFirstRow bool, isLastRow bool, isSeparatorRow bool, format text.Format) {
	if len(row) > 0 {
		// if the output has content, it means that this call is working on line
		// number 2 or more; separate them with a newline
		if out.Len() > 0 {
			out.WriteRune('\n')
		}

		// use a brand new strings.Builder if a row length limit has been set
		var outLine *strings.Builder
		if t.allowedRowLength > 0 {
			outLine = &strings.Builder{}
		} else {
			outLine = out
		}
		// grow the strings.Builder to the maximum possible row length
		outLine.Grow(t.maxRowLength)

		if !t.disableBorder {
			t.renderMarginLeft(outLine, isFirstRow, isLastRow, isSeparatorRow)
		}
		for colIdx, maxColumnLength := range t.maxColumnLengths {
			t.renderColumn(outLine, rowNum, row, colIdx, maxColumnLength, colors, isFirstRow, isLastRow, isSeparatorRow, format)
		}
		if !t.disableBorder {
			t.renderMarginRight(outLine, isFirstRow, isLastRow, isSeparatorRow)
		}

		if outLine != out {
			outLineStr := outLine.String()
			if utf8.RuneCountInString(outLineStr) >= t.allowedRowLength {
				finalCharIdx := t.allowedRowLength - utf8.RuneCountInString(t.style.BoxUnfinishedRow)
				if finalCharIdx > 0 {
					outLineStr = string([]rune(outLineStr)[0:finalCharIdx])
					out.WriteString(outLineStr)
					out.WriteString(t.style.BoxUnfinishedRow)
				}
			} else {
				out.WriteString(outLineStr)
			}
		}
	}
}

func (t *Table) renderMarginLeft(out *strings.Builder, isFirstRow bool, isLastRow bool, isSeparatorRow bool) {
	// type of row determines the character used (top/bottom/separator/etc.)
	if isFirstRow {
		out.WriteString(t.style.BoxTopLeft)
	} else if isLastRow {
		out.WriteString(t.style.BoxBottomLeft)
	} else if isSeparatorRow {
		out.WriteString(t.style.BoxLeftSeparator)
	} else {
		out.WriteString(t.style.BoxLeft)
	}
}

func (t *Table) renderMarginRight(out *strings.Builder, isFirstRow bool, isLastRow bool, isSeparatorRow bool) {
	// type of row determines the character used (top/bottom/separator/etc.)
	if isFirstRow {
		out.WriteString(t.style.BoxTopRight)
	} else if isLastRow {
		out.WriteString(t.style.BoxBottomRight)
	} else if isSeparatorRow {
		out.WriteString(t.style.BoxRightSeparator)
	} else {
		out.WriteString(t.style.BoxRight)
	}
}

func (t *Table) renderRow(out *strings.Builder, rowNum int, row Row, colors []*color.Color, isFirstRow bool, isLastRow bool, isSeparatorRow bool, format text.Format) {
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
		t.renderLine(out, rowNum, row, colors, isFirstRow, isLastRow, isSeparatorRow, format)
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
			t.renderLine(out, rowNum, rowLine, colors, isFirstRow, isLastRow, isSeparatorRow, format)
		}
	}
}

func (t *Table) renderRows(out *strings.Builder, rows []Row, colors []*color.Color, format text.Format) {
	for idx, row := range rows {
		t.renderRow(out, idx+1, row, colors, false, false, false, format)
		if t.enableSeparators && idx < len(rows)-1 {
			t.renderRowSeparator(out, false, false)
		}
	}
}

func (t *Table) renderRowSeparator(out *strings.Builder, isFirstRow bool, isLastRow bool) {
	t.renderLine(out, -1, t.rowSeparator, nil, isFirstRow, isLastRow, true, text.FormatDefault)
}
