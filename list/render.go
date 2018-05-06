package list

import (
	"strings"
)

// Render renders the List in a human-readable "pretty" format. Example:
//  ┌─ Game Of Thrones
//  └─┬─ Winter
//    ├─ Is
//    ├─ Coming
//    └─┬─ This
//      ├─ Is
//      └─ Known
func (l *List) Render() string {
	l.initForRender()

	// initForRender a new strings.Builder to build the output efficiently and
	// grow it by the pre-calculated "approxSize"
	var out strings.Builder
	out.Grow(l.approxSize)

	// render the List item-by-item
	for idx, item := range l.items {
		l.renderItem(&out, idx, item)
	}

	return l.render(&out)
}

func (l *List) renderItem(out *strings.Builder, idx int, item *listItem) {
	isFirstItem := bool(idx == 0)
	isLastItem := bool(idx == (len(l.items) - 1))
	isLastItemInSubList := bool(idx < (len(l.items)-1) && item.Level > l.items[idx+1].Level)

	// when working on item number 2 or more, render a newline first
	if idx > 0 {
		out.WriteRune('\n')
	}

	// render the prefix or the leading text before the actual item
	l.renderItemPrefix(out, idx, item)

	// render the "bullet"
	isNewLevel := bool(idx > 0 && l.items[idx].Level > l.items[idx-1].Level)
	if isFirstItem {
		out.WriteString(l.style.CharItemTop)
	} else if isNewLevel {
		if isLastItem {
			out.WriteString(l.style.CharItemSingle)
		} else {
			out.WriteString(l.style.CharItemFirst)
		}
	} else if isLastItem {
		out.WriteString(l.style.CharItemBottom)
	} else if isLastItemInSubList {
		out.WriteString(l.style.CharItemBottom)
	} else {
		out.WriteString(l.style.CharItem)
	}

	// pad as directed before rendering the item text
	out.WriteString(l.style.CharPaddingRight)
	out.WriteRune(' ')
	out.WriteString(l.style.Format.Apply(item.Text))
}

func (l *List) renderItemPrefix(out *strings.Builder, idx int, item *listItem) {
	// render spaces and connectors until the item's position
	for levelIdx := 0; levelIdx < item.Level; levelIdx++ {
		if l.hasMoreItemsInLevel(levelIdx, idx) {
			if idx > 0 && item.Level > l.items[idx-1].Level {
				if levelIdx == item.Level-1 {
					out.WriteString(l.style.CharVerticalConnect)
					out.WriteString(l.style.CharHorizontal)
				} else {
					out.WriteString(l.style.CharVertical)
					out.WriteString(" ")
				}
			} else {
				out.WriteString(l.style.CharVertical)
				out.WriteString(" ")
			}
		} else if idx > 0 && levelIdx == l.items[idx-1].Level {
			out.WriteString(l.style.CharConnectBottom)
			out.WriteString(l.style.CharHorizontal)
		} else {
			out.WriteString("  ")
		}
	}
}
