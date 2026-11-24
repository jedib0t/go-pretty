package list

import (
	"html"
	"strconv"
	"strings"
)

// RenderHTML renders the List in the HTML format. Example:
//
//	<ul class="go-pretty-table">
//	  <li>Game Of Thrones
//	    <ul class="go-pretty-table-1">
//	      <li>Winter</li>
//	      <li>Is</li>
//	      <li>Coming
//	        <ul class="go-pretty-table-2">
//	          <li>This</li>
//	          <li>Is</li>
//	          <li>Known</li>
//	        </ul>
//	      </li>
//	    </ul>
//	  </li>
//	  <li>The Dark Tower
//	    <ul class="go-pretty-table-1">
//	      <li>The Gunslinger</li>
//	    </ul>
//	  </li>
//	</ul>
func (l *List) RenderHTML() string {
	l.initForRender()

	var out strings.Builder
	if len(l.items) > 0 {
		out.Grow(l.approxSize)
		l.htmlRenderRecursively(&out, 0)
	}
	return l.render(&out)
}

// htmlRenderRecursively renders one list and returns the index after its last item.
func (l *List) htmlRenderRecursively(out *strings.Builder, idx int) int {
	level := l.items[idx].Level
	linePrefix := strings.Repeat("    ", level)

	out.WriteString(linePrefix)
	out.WriteString("<ul class=\"")
	out.WriteString(html.EscapeString(l.htmlCSSClass))
	if level > 0 {
		out.WriteRune('-')
		out.WriteString(strconv.Itoa(level))
	}
	out.WriteString("\">\n")

	for idx < len(l.items) && l.items[idx].Level == level {
		out.WriteString(linePrefix)
		out.WriteString("  <li>")
		out.WriteString(strings.ReplaceAll(html.EscapeString(l.items[idx].Text), "\n", "<br/>"))
		idx++
		if idx < len(l.items) && l.items[idx].Level > level {
			out.WriteRune('\n')
			idx = l.htmlRenderRecursively(out, idx)
			out.WriteRune('\n')
			out.WriteString(linePrefix)
			out.WriteString("  ")
		}
		out.WriteString("</li>\n")
	}
	out.WriteString(linePrefix)
	out.WriteString("</ul>")
	return idx
}
