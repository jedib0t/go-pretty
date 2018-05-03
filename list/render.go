package list

import "strings"

// Render renders the List in a human-readable "pretty" format. Example:
//  ┌─ Game Of Thrones
//  └─┬─ Winter
//    ├─ Is
//    ├─ Coming
//    └─┬─ This
//      ├─ Is
//      └─ Known
func (l *List) Render() string {
	l.init()

	// init a new strings.Builder to build the output efficiently and grow it by
	// the pre-calculated "approxSize"
	var out strings.Builder
	out.Grow(l.approxSize)

	// render the List item-by-item and use the "Level" property in the item to
	// determine the prefix/padding
	for idx, item := range l.items {
		// when working on item number 2 or more, render a newline first
		if idx > 0 {
			out.WriteRune('\n')
		}

		// if there is a change in level, render the connector; else just pad
		// the output with spaces based on the current level
		levelChanged := bool(idx > 0 && l.items[idx].Level > l.items[idx-1].Level)
		if levelChanged {
			if item.Level > 1 {
				out.WriteString(strings.Repeat(" ", (item.Level-1)*2))
			}
			out.WriteString(l.style.CharConnect + l.style.CharPaddingLeft)
		} else {
			out.WriteString(strings.Repeat(" ", item.Level*2))
		}

		// render the "bullet"
		if idx == 0 {
			out.WriteString(l.style.CharItemTop)
		} else if levelChanged {
			out.WriteString(l.style.CharItemFirst)
		} else if idx == len(l.items)-1 {
			out.WriteString(l.style.CharItemBottom)
		} else {
			out.WriteString(l.style.CharItem)
		}
		// pad as directed before rendering the item text
		out.WriteString(l.style.CharPaddingRight)
		out.WriteRune(' ')
		out.WriteString(l.style.Format.Apply(item.Text))
	}

	return out.String()
}
