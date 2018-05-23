package table

import (
	"io"

	"github.com/jedib0t/go-pretty/text"
)

// Writer declares the interfaces that can be used to setup and render a table.
type Writer interface {
	AppendFooter(row Row)
	AppendHeader(row Row)
	AppendRow(row Row)
	AppendRows(rows []Row)
	Length() int
	Render() string
	RenderCSV() string
	RenderHTML() string
	RenderMarkdown() string
	SetAlign(align []text.Align)
	SetAlignFooter(align []text.Align)
	SetAlignHeader(align []text.Align)
	SetAllowedColumnLengths(lengths []int)
	SetAllowedRowLength(length int)
	SetAutoIndex(autoIndex bool)
	SetCaption(format string, a ...interface{})
	SetColors(colors []text.Colors)
	SetColorsFooter(colors []text.Colors)
	SetColorsHeader(colors []text.Colors)
	SetHTMLCSSClass(cssClass string)
	SetIndexColumn(colNum int)
	SetOutputMirror(mirror io.Writer)
	SetPageSize(numLines int)
	SetStyle(style Style)
	SetVAlign(vAlign []text.VAlign)
	SetVAlignFooter(vAlign []text.VAlign)
	SetVAlignHeader(vAlign []text.VAlign)
	SortBy(sortBy []SortBy)
	Style() *Style
}

// NewWriter initializes and returns a Writer.
func NewWriter() Writer {
	return &Table{}
}
