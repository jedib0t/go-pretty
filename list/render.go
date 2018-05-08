package list

import (
	"strings"
)

// Render renders the List in a human-readable "pretty" format. Example:
//  * Game Of Thrones
//    * Winter
//    * Is
//    * Coming
//      * This
//      * Is
//      * Known
//  * The Dark Tower
//    * The Gunslinger
func (l *List) Render() string {
	l.initForRender()

	var out strings.Builder
	out.Grow(l.approxSize)
	for idx, item := range l.items {
		l.renderItem(&out, idx, item)
	}
	return l.render(&out)
}

func (l *List) renderItem(out *strings.Builder, idx int, item *listItem) {
	// when working on item number 2 or more, render a newline first
	if idx > 0 {
		out.WriteRune('\n')
	}

	var renderItemLine = func(lineIdx int, lineStr string) {
		// render the prefix or the leading text before the actual item
		l.renderItemBulletPrefix(out, idx, item.Level, lineIdx)
		l.renderItemBullet(out, idx, item.Level, lineIdx)

		// render the actual item
		out.WriteString(lineStr)
	}

	itemStr := l.style.Format.Apply(item.Text)
	if strings.Contains(itemStr, "\n") {
		for lineIdx, lineStr := range strings.Split(itemStr, "\n") {
			if lineIdx > 0 {
				out.WriteRune('\n')
			}
			renderItemLine(lineIdx, lineStr)
		}
	} else {
		renderItemLine(0, itemStr)
	}
}

func (l *List) renderItemBullet(out *strings.Builder, itemIdx int, itemLevel int, lineIdx int) {
	isFirstItem := bool(itemIdx == 0)
	isLastItem := bool(itemIdx == (len(l.items) - 1))
	isLastItemInSubList := bool(itemIdx < (len(l.items)-1) && itemLevel > l.items[itemIdx+1].Level)
	isNewLevel := bool(itemIdx > 0 && l.items[itemIdx].Level > l.items[itemIdx-1].Level)

	if lineIdx > 0 {
		out.WriteString(l.style.CharVertical)
		out.WriteString("  ")
	} else {
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
	}
}

func (l *List) renderItemBulletPrefix(out *strings.Builder, itemIdx int, itemLevel int, lineIdx int) {
	if l.style.LinePrefix != "" {
		out.WriteString(l.style.LinePrefix)
	}

	isFirstItem := bool(itemIdx == 0)
	isFirstLine := bool(lineIdx == 0)
	isFirstLineOfNonFirstItem := bool(isFirstLine && !isFirstItem)
	isIndentedFromPreviousItem := bool(!isFirstItem && itemLevel > l.items[itemIdx-1].Level)

	// render spaces and connectors until the item's position
	for levelIdx := 0; levelIdx < itemLevel; levelIdx++ {
		if l.hasMoreItemsInLevel(levelIdx, itemIdx) {
			if isFirstLineOfNonFirstItem && isIndentedFromPreviousItem && levelIdx == itemLevel-1 {
				out.WriteString(l.style.CharVerticalConnect)
				out.WriteString(l.style.CharHorizontal)
			} else {
				out.WriteString(l.style.CharVertical)
				out.WriteRune(' ')
			}
		} else if isFirstLineOfNonFirstItem && levelIdx == l.items[itemIdx-1].Level {
			out.WriteString(l.style.CharConnectBottom)
			out.WriteString(l.style.CharHorizontal)
		} else {
			out.WriteString("  ")
		}
	}
}
