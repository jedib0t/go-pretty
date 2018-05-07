## List

Pretty-print lists with multiple levels/indents into ASCII/Unicode strings.

  - Append Items one-by-one or as a group
  - Indent/UnIndent as you like
  - Mirror output to an io.Writer object (like os.StdOut)
  - Completely customizable styles
    - Many ready-to-use styles: [style.go](style.go)
  - Render as:
    - (ASCII/Unicode) List
    - Markdown List

```
 ■ Game Of Thrones
   ■ Winter
   ■ Is
   ■ Coming
     ■ This
     ■ Is
     ■ Known
   ■ The Dark Tower
     ■ The Gunslinger
```

A demonstration of all the capabilities can be found here: [../cmd/demo-list](../cmd/demo-list)

Find documentation here: [GoDoc](https://godoc.org/github.com/jedib0t/go-pretty/list)

### TODO

  - Multi-line items
  - Line-width restrictions
  - Render as HTML
