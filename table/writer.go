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
	SetAutoIndex(autoIndex bool)
	SetCaption(format string, a ...interface{})
	SetColors(colors []text.Colors)
	SetColorsFooter(colors []text.Colors)
	SetColorsHeader(colors []text.Colors)
	SetHTMLCSSClass(cssClass string)
	SetOutputMirror(mirror io.Writer)
	SetStyle(style Style)
	SetVAlign(vAlign []text.VAlign)
	ShowBorder(show bool)
	ShowSeparators(show bool)
	Style() *Style
}

// NewWriter initializes and returns a Writer.
func NewWriter() Writer {
	return &Table{}
}
