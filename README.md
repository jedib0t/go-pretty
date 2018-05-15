# go-pretty

[![Build Status](https://travis-ci.com/jedib0t/go-pretty.svg?branch=master)](https://travis-ci.com/jedib0t/go-pretty)
[![Coverage Status](https://coveralls.io/repos/github/jedib0t/go-pretty/badge.svg?branch=master)](https://coveralls.io/github/jedib0t/go-pretty?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/jedib0t/go-pretty)](https://goreportcard.com/report/github.com/jedib0t/go-pretty)
[![GoDoc](https://godoc.org/github.com/jedib0t/go-pretty?status.svg)](https://godoc.org/github.com/jedib0t/go-pretty)


Utilities to prettify console output of tables, lists, text, etc.

## Table

Pretty-print tables into ASCII/Unicode strings.

  - Add Rows one-by-one or as a group
  - Add Header(s) and Footer(s)
  - Auto Index Rows (1, 2, 3 ...) and Columns (A, B, C, ...)
  - Limit the length of the Rows; limit the length of individual Columns
  - Alignment - Horizontal & Vertical
    - Auto (horizontal) Align (numeric columns are aligned Right)
    - Custom (horizontal) Align per column
    - Custom (vertical) VAlign per column (and multi-line column support)
  - Mirror output to an io.Writer object (like os.StdOut)
  - Completely customizable styles
    - Many ready-to-use styles: [table/style.go](table/style.go)
    - Colorize Headers/Body/Footers using [text/color](text/color.go)
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

  - Append Items one-by-one or as a group
  - Indent/UnIndent as you like
  - Limit the length of the Lines/Items
  - Support Items with Multiple-lines
  - Mirror output to an io.Writer object (like os.StdOut)
  - Completely customizable styles
    - Many ready-to-use styles: [list/style.go](list/style.go)
  - Render as:
    - (ASCII/Unicode) List
    - HTML List (with custom CSS Class)
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

A demonstration of all the capabilities can be found here: [cmd/demo-list](cmd/demo-list)

## Text

The following features are all used by the other packages in this project.
Specifically, `table` and `list` use these extensively:

   - Align text horizontally
     - [text/align.go](text/align.go)
   - Align text vertically
     - [text/valign.go](text/valign.go)
   - Colorize text
     - [text/color.go](text/color.go)
   - Format text (convert case for now)
     - [text/format.go](text/format.go)

The unit-tests for each of the above show how these are to be used.

## Benchmarks

Partial output of `make bench`:
```
BenchmarkList_Render-8                   1000000              1836 ns/op             808 B/op         22 allocs/op
BenchmarkTable_Render-8                   100000             20736 ns/op            5426 B/op        191 allocs/op
BenchmarkTable_RenderCSV-8                300000              4394 ns/op            2336 B/op         45 allocs/op
BenchmarkTable_RenderHTML-8               200000              6563 ns/op            3793 B/op         44 allocs/op
BenchmarkTable_RenderMarkdown-8           300000              4666 ns/op            2272 B/op         43 allocs/op
```
