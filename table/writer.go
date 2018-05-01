package table

import "github.com/jedib0t/go-pretty/text"

// Writer declares the interfaces implemented by Table.
type Writer interface {
	AppendFooter(row Row)
	AppendHeader(row Row)
	AppendRow(row Row)
	AppendRows(rows []Row)
	Length() int
	Render() string
	RenderCSV() string
	RenderHTML() string
	SetAlign(align []text.Align)
	SetCaption(format string, a ...interface{})
	SetColors(colors []text.Colors)
	SetColorsFooter(colors []text.Colors)
	SetColorsHeader(colors []text.Colors)
	SetHTMLCSSClass(cssClass string)
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
