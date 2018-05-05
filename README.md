# go-pretty

[![Build Status](https://travis-ci.org/jedib0t/go-pretty.svg?branch=master)](https://travis-ci.org/jedib0t/go-pretty)
[![Coverage Status](https://coveralls.io/repos/github/jedib0t/go-pretty/badge.svg?branch=master)](https://coveralls.io/github/jedib0t/go-pretty?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/jedib0t/go-pretty)](https://goreportcard.com/report/github.com/jedib0t/go-pretty)
[![GoDoc](https://godoc.org/github.com/jedib0t/go-pretty?status.svg)](https://godoc.org/github.com/jedib0t/go-pretty)


Utilities to prettify console output of tables, lists, text, etc.

## Table

Pretty-print tables into ASCII/Unicode strings.

  - Add Rows one-by-one or as a group
  - Add Header(s) and Footer(s)
  - Auto Index Rows (1, 2, 3 ...) and Columns (A, B, C, ...)
  - Mirror output to an io.Writer object (like os.StdOut)
  - Limit the length of the Rows; limit the length of individual Columns
  - Alignment - Horizontal & Vertical
    - Auto (horizontal) Align (numeric columns are aligned Right)
    - Custom (horizontal) Align per column
    - Custom (vertical) VAlign per column (and multi-line column support)
  - Completely customizable styles
    - Many ready-to-use styles: [table/style.go](table/style.go)
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

A demonstration of all the capabilities can be found here: [cmd/demo-table](cmd/demo-table)

## List

Pretty-print lists with multiple levels/indents into ASCII/Unicode strings.

  - Completely customizable styles
    - Many ready-to-use styles: [list/style.go](list/style.go)

```
- Game Of Thrones
--- Winter
  - Is
  - Coming
  --- This
    - Is
    - Known
```

A demonstration of all the capabilities can be found here: [cmd/demo-list](cmd/demo-list)

## Text

The following features are all used by the other packages in this project.
Specifically, `table` and `list` use these extensively:

   - Align text horizontally
     - [text/align.go](text/align.go)
   - Align text vertically
     - [text/valign.go](text/valign.go)
   - Colorize text using a simpler interface to the awesome package [github.com/fatih/color](github.com/fatih.color)
     - [text/color.go](text/color.go)
   - Format text (convert case for now)
     - [text/format.go](text/format.go)

The unit-tests for each of the above show how these are to be used.

## Benchmarks

Partial output of `make bench`:
```
BenchmarkList_Render-8                   1000000              1644 ns/op             608 B/op         24 allocs/op
BenchmarkTable_Render-8                   100000             18828 ns/op            5009 B/op        190 allocs/op
BenchmarkTable_RenderCSV-8                300000              4066 ns/op            1920 B/op         44 allocs/op
BenchmarkTable_RenderHTML-8               300000              5621 ns/op            3377 B/op         43 allocs/op
BenchmarkTable_RenderMarkdown-8           300000              4221 ns/op            1856 B/op         42 allocs/op
```
