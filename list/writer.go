package list

// Writer declares the interfaces that can be used to setup and render a list.
type Writer interface {
	AppendItem(item interface{})
	AppendItems(items []interface{})
	Indent()
	Length() int
	Render() string
	SetStyle(style Style)
	Style() *Style
	UnIndent()
}

// NewWriter initializes and returns a Writer.
func NewWriter() Writer {
	return &List{}
}
