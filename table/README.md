## Table

Pretty-print tables into ASCII/Unicode strings.

  - Add Rows one-by-one or as a group
  - Add Header(s) and Footer(s)
  - Auto Index Rows (1, 2, 3 ...) and Columns (A, B, C, ...)
  - Set output to be mirrored to an io.Writer object like os.StdOut
  - Alignment/ - Horizontal & Vertical
    - Auto (horizontal) Align (numeric columns are aligned Right)
    - Custom (horizontal) Align per column
    - Custom (vertical) VAlign per column (and multi-line column support)
  - Completely customizable styles
    - Many ready-to-use styles: [style.go](style.go)
    - Colorize Headers/Body/Footers using [github.com/fatih/color](https://github.com/fatih/color)
    - Custom text-case for Headers/Body/Footers
    - Enable separators between each row
    - Render table without a Border
  - Render as:
    - (ASCII/Unicode) Table
    - CSV
    - HTML Table (with custom CSS Style)
    - Markdown Table


```
+-----+------------+-----------+--------+-----------------------------+
|   # | FIRST NAME | LAST NAME | SALARY |                             |
+-----+------------+-----------+--------+-----------------------------+
|   1 | Arya       | Stark     |   3000 |                             |
|  20 | Jon        | Snow      |   2000 | You know nothing, Jon Snow! |
| 300 | Tyrion     | Lannister |   5000 |                             |
+-----+------------+-----------+--------+-----------------------------+
|     |            | TOTAL     |  10000 |                             |
+-----+------------+-----------+--------+-----------------------------+
```

A demonstration of all the capabilities can be found here: [demo](demo)

Documentation: [GoDoc](https://godoc.org/github.com/jedib0t/go-pretty/table)

### TODO

  - Performance Optimizations (Memory Usage)
  - Row and Cell Width Restrictions
  - Generic Cell Content Transformers (with some ready-made ones)
    - Base64 Decoder
    - Currency Formatter
    - UnixTime to Date & Time
    - Status Formatter (color "FAILED" in RED, etc.)
