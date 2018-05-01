## Table

Pretty-print tables into ASCII/Unicode strings.

  - Supports Header and Footer
  - Supports Adding Rows one-by-one or as a group
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

A demonstration of all the capabilities can be found here: [demo/demo.go](demo/demo.go)

Documentation: [GoDoc](https://godoc.org/github.com/jedib0t/go-pretty/table)
