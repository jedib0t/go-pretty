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
  - Page results by a specified number of Lines
  - Alignment - Horizontal & Vertical
    - Auto (horizontal) Align (numeric columns are aligned Right)
    - Custom (horizontal) Align per column
    - Custom (vertical) VAlign per column (and multi-line column support)
  - Mirror output to an io.Writer object (like os.StdOut)
  - Sort by any of the Columns (by Column Name or Number)
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

A demonstration of all the capabilities can be found here:
[cmd/demo-table](cmd/demo-table)

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

A demonstration of all the capabilities can be found here:
[cmd/demo-list](cmd/demo-list)

# Progress

Track the Progress of one or more Tasks (like downloading multiple files in
parallel).

  - Track one or more Tasks at the same time
  - Dynamically add one or more Task Trackers while `Render()` is in progress
  - Choose to have the Writer auto-stop the Render when no more Trackers are
    in queue, or manually stop using `Stop()`
  - Redirect output to an io.Writer object (like os.StdOut)
  - Completely customizable styles
    - Many ready-to-use styles: [progress/style.go](progress/style.go)
    - Colorize various parts of the Tracker using `StyleColors`
    - Customize how Trackers get rendered using `StyleOptions`

Sample Progress Tracking:
```
Calculating Total   #  1 ... done! [3.25K in 100ms]
Calculating Total   #  2 ... done! [6.50K in 100ms]
Downloading File    #  3 ... done! [9.75KB in 100ms]
Transferring Amount #  4 ... done! [$26.00K in 200ms]
Transferring Amount #  5 ... done! [£32.50K in 201ms]
Downloading File    #  6 ... done! [58.50KB in 300ms]
Calculating Total   #  7 ... done! [91.00K in 400ms]
Transferring Amount #  8 ... 60.9% (●●●●●●●●●●●●●●◌◌◌◌◌◌◌◌◌) [$78.00K in 399.071ms]
Downloading File    #  9 ... 32.1% (●●●●●●●○◌◌◌◌◌◌◌◌◌◌◌◌◌◌◌) [58.50KB in 298.947ms]
Transferring Amount # 10 ... 13.0% (●●○◌◌◌◌◌◌◌◌◌◌◌◌◌◌◌◌◌◌◌◌) [£32.50K in 198.84ms]
```

A demonstration of all the capabilities can be found here:
[cmd/demo-progress](cmd/demo-progress)

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
BenchmarkList_Render-8                    500000              2182 ns/op             760 B/op         40 allocs/op
BenchmarkProgress_Render-8                     2         800863000 ns/op            7200 B/op        209 allocs/op
BenchmarkTable_Render-8                   100000             20839 ns/op            5538 B/op        188 allocs/op
BenchmarkTable_RenderCSV-8                300000              4479 ns/op            2464 B/op         45 allocs/op
BenchmarkTable_RenderHTML-8               200000              6422 ns/op            3921 B/op         44 allocs/op
BenchmarkTable_RenderMarkdown-8           300000              4755 ns/op            2400 B/op         43 allocs/op
```
