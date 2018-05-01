package list

// Writer declares the interfaces implemented by List.
type Writer interface {
	AppendItem(item interface{})
	AppendItems(items []interface{})
	Indent()
	Length() int
	Render() string
	SetStyle(style Style)
	Style() *Style
}

// NewWriter initializes and returns a Writer.
func NewWriter() Writer {
	return &List{}
}
