# go-pretty

[![Build Status](https://travis-ci.org/jedib0t/go-pretty.svg?branch=master)](https://travis-ci.org/jedib0t/go-pretty)
[![Go Report Card](https://goreportcard.com/badge/github.com/jedib0t/go-pretty)](https://goreportcard.com/report/github.com/jedib0t/go-pretty)
[![GoDoc](https://godoc.org/github.com/jedib0t/go-pretty?status.svg)](https://godoc.org/github.com/jedib0t/go-pretty)
<!-- [![Coverage Status](https://coveralls.io/repos/github/jedib0t/go-pretty/badge.svg?branch=master)](https://coveralls.io/github/jedib0t/go-pretty?branch=master) -->

Utilities to prettify console output of tables, lists, text, etc.

_Note_: Coveralls Integration is [broken](https://github.com/mattn/goveralls/issues/114) as of now.

## Table

Pretty-print tables into ASCII/Unicode strings.

  - Supports Header and Footer
  - Supports Adding Rows one-by-one or as a group
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

A demonstration of all the capabilities can be found here: [table/demo/demo.go](table/demo/demo.go)

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

A demonstration of all the capabilities can be found here: [list/demo/demo.go](list/demo/demo.go)

## Benchmarks

Partial output of `make bench`:
```
BenchmarkList_Render-8                   1000000              1651 ns/op             608 B/op         24 allocs/op
BenchmarkTable_Render-8                    50000             26410 ns/op            7138 B/op        416 allocs/op
BenchmarkTable_RenderCSV-8                300000              5827 ns/op            2656 B/op         90 allocs/op
BenchmarkTable_RenderHTML-8               200000              7435 ns/op            4129 B/op         89 allocs/op
BenchmarkTable_RenderMarkdown-8           200000              7493 ns/op            4129 B/op         89 allocs/op
```
