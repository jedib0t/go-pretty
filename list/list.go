package list

import (
	"fmt"
	"unicode/utf8"
)

// listItem represents one line in the List
type listItem struct {
	Level int
	Text  string
}

// List helps print a 2-dimensional array in a human readable pretty-List.
type List struct {
	// approxSize stores the approximate output length/size
	approxSize int
	// level stores the current indentation level
	level int
	// items contains the list of items to render
	items []listItem
	// style contains all the strings used to draw the List, and more
	style *Style
}

// AppendItem appends the item to the List of items to render.
func (l *List) AppendItem(item interface{}) {
	listEntry := listItem{
		Level: l.level,
		Text:  fmt.Sprint(item),
	}

	// account for the following when incrementing approxSize: 1. length of
	// text, 2. left-padding, 3. list-prefix, 4. right-padding, 5. newline
	l.approxSize += utf8.RuneCountInString(listEntry.Text) + (l.level * 2) + 2 + 1 + 1
	// 6. connector in case of level change
	if len(l.items) > 0 && listEntry.Level > l.items[len(l.items)-1].Level {
		l.approxSize++
	}

	l.items = append(l.items, listEntry)
}

// AppendItems appends the items to the List of items to render.
func (l *List) AppendItems(items []interface{}) {
	for _, item := range items {
		l.AppendItem(item)
	}
}

// Indent indents the following items to appear right-shifted.
func (l *List) Indent() {
	l.level++
}

// Length returns the number of items to be rendered.
func (l *List) Length() int {
	return len(l.items)
}

// SetStyle overrides the DefaultStyle with the provided one.
func (l *List) SetStyle(style Style) {
	l.style = &style
}

// Style returns the current style.
func (l *List) Style() *Style {
	return l.style
}

func (l *List) init() {
	if l.style == nil {
		l.style = &StyleDefault
	}
}
