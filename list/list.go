package list

import (
	"fmt"
	"io"
	"strings"
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
	items []*listItem
	// outputMirror stores an io.Writer where the "Render" functions would write
	outputMirror io.Writer
	// style contains all the strings used to draw the List, and more
	style *Style
}

// AppendItem appends the item to the List of items to render.
func (l *List) AppendItem(item interface{}) {
	l.items = append(l.items, l.analyzeAndStringify(item))
}

// AppendItems appends the items to the List of items to render.
func (l *List) AppendItems(items []interface{}) {
	for _, item := range items {
		l.AppendItem(item)
	}
}

// Indent indents the following items to appear right-shifted.
func (l *List) Indent() {
	if len(l.items) == 0 {
		// should not indent when there is no item in the current level
	} else if l.level > l.items[len(l.items)-1].Level {
		// already indented compared to previous item; do not indent more
	} else {
		l.level++
	}
}

// Length returns the number of items to be rendered.
func (l *List) Length() int {
	return len(l.items)
}

// Reset sets the List to its initial state.
func (l *List) Reset() {
	l.approxSize = 0
	l.items = make([]*listItem, 0)
	l.level = 0
	l.style = nil
}

// SetOutputMirror sets an io.Writer for all the Render functions to "Write" to
// in addition to returning a string.
func (l *List) SetOutputMirror(mirror io.Writer) {
	l.outputMirror = mirror
}

// SetStyle overrides the DefaultStyle with the provided one.
func (l *List) SetStyle(style Style) {
	l.style = &style
}

// Style returns the current style.
func (l *List) Style() *Style {
	return l.style
}

func (l *List) analyzeAndStringify(item interface{}) *listItem {
	listEntry := &listItem{
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

	return listEntry
}

// UnIndent un-indents the following items to appear left-shifted.
func (l *List) UnIndent() {
	if l.level > 0 {
		l.level--
	}
}

func (l *List) initForRender() {
	if l.style == nil {
		l.style = &StyleDefault
	}
}

func (l *List) hasMoreItemsInLevel(levelIdx int, fromItemIdx int) bool {
	for idx := fromItemIdx + 1; idx >= 0 && idx < len(l.items); idx++ {
		if l.items[idx].Level < levelIdx {
			return false
		} else if l.items[idx].Level == levelIdx {
			return true
		}
	}
	return false
}

func (l *List) render(out *strings.Builder) string {
	outStr := out.String()
	if l.outputMirror != nil {
		l.outputMirror.Write([]byte(outStr))
		l.outputMirror.Write([]byte("\n"))
	}
	return outStr
}
