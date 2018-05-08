package list

import (
	"html"
	"strconv"
	"strings"
)

// RenderHTML renders the List in the HTML format. Example:
//
func (l *List) RenderHTML() string {
	l.initForRender()

	var out strings.Builder
	if len(l.items) > 0 {
		l.htmlRenderRecursively(&out, 0, l.items[0])
	}
	return l.render(&out)
}

func (l *List) htmlRenderRecursively(out *strings.Builder, idx int, item *listItem) int {
	linePrefix := strings.Repeat("  ", item.Level)

	out.WriteString(linePrefix)
	out.WriteString("<ul class=\"")
	out.WriteString(l.htmlCSSClass)
	if item.Level > 0 {
		out.WriteRune('-')
		out.WriteString(strconv.Itoa(item.Level))
	}
	out.WriteString("\">\n")
	var numItemsRendered int
	for itemIdx := idx; itemIdx < len(l.items); itemIdx++ {
		if l.items[itemIdx].Level == item.Level {
			out.WriteString(linePrefix)
			out.WriteString("  <li>")
			out.WriteString(html.EscapeString(l.items[itemIdx].Text))
			out.WriteString("</li>\n")
			numItemsRendered++
		} else if l.items[itemIdx].Level > item.Level { // indent
			numItemsRenderedRecursively := l.htmlRenderRecursively(out, itemIdx, l.items[itemIdx])
			numItemsRendered += numItemsRenderedRecursively
			itemIdx += numItemsRenderedRecursively - 1
			if numItemsRendered > 0 {
				out.WriteRune('\n')
			}
		} else { // un-indent
			break
		}
	}
	out.WriteString(linePrefix)
	out.WriteString("</ul>")
	return numItemsRendered
}
