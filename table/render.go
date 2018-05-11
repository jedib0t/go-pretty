package table

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/text"
	"github.com/jedib0t/go-pretty/util"
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
		if t.style.Options.DrawBorder {
			t.renderRowSeparator(&out, true, false)
		}
		if len(t.rowsHeader) > 0 || t.autoIndex {
			if len(t.rowsHeader) > 0 {
				t.renderRows(&out, t.rowsHeader, t.colorsHeader, t.style.Format.Header)
			} else {
				t.renderRow(&out, 0, t.getAutoIndexColumnIDs(), t.colorsHeader, false, false, false, text.FormatUpper)
			}
			t.renderRowSeparator(&out, false, false)
		}
		t.renderRows(&out, t.rows, t.colors, t.style.Format.Row)
		if len(t.rowsFooter) > 0 {
			t.renderRowSeparator(&out, false, false)
			t.renderRows(&out, t.rowsFooter, t.colorsFooter, t.style.Format.Footer)
		}
		if t.style.Options.DrawBorder {
			t.renderRowSeparator(&out, false, true)
		}
		if t.caption != "" {
			out.WriteRune('\n')
			out.WriteString(t.caption)
		}
	}
	return t.render(&out)
}

func (t *Table) renderColumn(out *strings.Builder, rowNum int, row RowStr, colIdx int, maxColumnLength int, colors []*color.Color, isFirstRow bool, isLastRow bool, isSeparatorRow bool, format text.Format) {
	// when working on the first column, and autoIndex is true, insert a new
	// column with the row number on it.
	if colIdx == 0 && t.autoIndex {
		if rowNum < 0 {
			numChars := t.autoIndexVIndexMaxLength + utf8.RuneCountInString(t.style.Box.PaddingLeft) +
				utf8.RuneCountInString(t.style.Box.PaddingRight)
			out.WriteString(strings.Repeat(t.style.Box.MiddleHorizontal, numChars))
		} else {
			out.WriteString(t.style.Box.PaddingLeft)
			rowNumStr := fmt.Sprint(rowNum)
			if rowNum == 0 {
				rowNumStr = strings.Repeat(" ", t.autoIndexVIndexMaxLength)
			}
			out.WriteString(text.AlignRight.Apply(rowNumStr, t.autoIndexVIndexMaxLength))
			out.WriteString(t.style.Box.PaddingRight)
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
		colStr = format.Apply(row[colIdx])
	}
	colStr = t.getAlign(colIdx).Apply(colStr, maxColumnLength)

	// pad both sides of the column (when not a separator row)
	if !isSeparatorRow {
		colStr = t.style.Box.PaddingLeft + colStr + t.style.Box.PaddingRight
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
			out.WriteString(t.style.Box.TopSeparator)
		} else if isLastRow {
			out.WriteString(t.style.Box.BottomSeparator)
		} else {
			out.WriteString(t.style.Box.MiddleSeparator)
		}
	} else {
		out.WriteString(t.style.Box.MiddleVertical)
	}
}

func (t *Table) renderLine(out *strings.Builder, rowNum int, row RowStr, colors []*color.Color, isFirstRow bool, isLastRow bool, isSeparatorRow bool, format text.Format) {
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

		if t.style.Options.DrawBorder {
			t.renderMarginLeft(outLine, isFirstRow, isLastRow, isSeparatorRow)
		}
		for colIdx, maxColumnLength := range t.maxColumnLengths {
			t.renderColumn(outLine, rowNum, row, colIdx, maxColumnLength, colors, isFirstRow, isLastRow, isSeparatorRow, format)
		}
		if t.style.Options.DrawBorder {
			t.renderMarginRight(outLine, isFirstRow, isLastRow, isSeparatorRow)
		}

		if outLine != out {
			outLineStr := outLine.String()
			if util.RuneCountWithoutEscapeSeq(outLineStr) > t.allowedRowLength {
				trimLength := t.allowedRowLength - utf8.RuneCountInString(t.style.Box.UnfinishedRow)
				if trimLength > 0 {
					out.WriteString(util.TrimTextWithoutEscapeSeq(outLineStr, trimLength))
					out.WriteString(t.style.Box.UnfinishedRow)
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
		out.WriteString(t.style.Box.TopLeft)
	} else if isLastRow {
		out.WriteString(t.style.Box.BottomLeft)
	} else if isSeparatorRow {
		out.WriteString(t.style.Box.LeftSeparator)
	} else {
		out.WriteString(t.style.Box.Left)
	}
}

func (t *Table) renderMarginRight(out *strings.Builder, isFirstRow bool, isLastRow bool, isSeparatorRow bool) {
	// type of row determines the character used (top/bottom/separator/etc.)
	if isFirstRow {
		out.WriteString(t.style.Box.TopRight)
	} else if isLastRow {
		out.WriteString(t.style.Box.BottomRight)
	} else if isSeparatorRow {
		out.WriteString(t.style.Box.RightSeparator)
	} else {
		out.WriteString(t.style.Box.Right)
	}
}

func (t *Table) renderRow(out *strings.Builder, rowNum int, row RowStr, colors []*color.Color, isFirstRow bool, isLastRow bool, isSeparatorRow bool, format text.Format) {
	// fit every column into the allowedColumnLength/maxColumnLength limit and
	// in the process find the max. number of lines in any column in this row
	colMaxLines := 0
	rowWrapped := make(RowStr, len(row))
	for colIdx, colStr := range row {
		rowWrapped[colIdx] = util.WrapText(colStr, t.maxColumnLengths[colIdx])
		colNumLines := strings.Count(rowWrapped[colIdx], "\n") + 1
		if colNumLines > colMaxLines {
			colMaxLines = colNumLines
		}
	}

	// if there is just 1 line in all columns, add the row as such; else split
	// each column into individual lines and render them one-by-one
	if colMaxLines == 1 {
		t.renderLine(out, rowNum, row, colors, isFirstRow, isLastRow, isSeparatorRow, format)
	} else {
		// convert one row into N # of rows based on colMaxLines
		rowLines := make([]RowStr, len(row))
		for colIdx, colStr := range rowWrapped {
			rowLines[colIdx] = t.getVAlign(colIdx).ApplyStr(colStr, colMaxLines)
		}
		for colLineIdx := 0; colLineIdx < colMaxLines; colLineIdx++ {
			rowLine := make(RowStr, len(rowLines))
			for colIdx, colLines := range rowLines {
				rowLine[colIdx] = colLines[colLineIdx]
			}
			t.renderLine(out, rowNum, rowLine, colors, isFirstRow, isLastRow, isSeparatorRow, format)
		}
	}
}

func (t *Table) renderRows(out *strings.Builder, rows []RowStr, colors []*color.Color, format text.Format) {
	for idx, row := range rows {
		t.renderRow(out, idx+1, row, colors, false, false, false, format)
		if t.style.Options.SeparateRows && idx < len(rows)-1 {
			t.renderRowSeparator(out, false, false)
		}
	}
}

func (t *Table) renderRowSeparator(out *strings.Builder, isFirstRow bool, isLastRow bool) {
	t.renderLine(out, -1, t.rowSeparator, nil, isFirstRow, isLastRow, true, text.FormatDefault)
}
