# go-pretty

[![Go Reference](https://pkg.go.dev/badge/github.com/jedib0t/go-pretty/v6.svg)](https://pkg.go.dev/github.com/jedib0t/go-pretty/v6)
[![Build Status](https://github.com/jedib0t/go-pretty/workflows/CI/badge.svg?branch=main)](https://github.com/jedib0t/go-pretty/actions?query=workflow%3ACI+event%3Apush+branch%3Amain)
[![Coverage Status](https://coveralls.io/repos/github/jedib0t/go-pretty/badge.svg?branch=main)](https://coveralls.io/github/jedib0t/go-pretty?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/jedib0t/go-pretty)](https://goreportcard.com/report/github.com/jedib0t/go-pretty)

Utilities to prettify console output of tables, lists, progress-bars, text, etc.

## Table

Pretty-print tables into ASCII/Unicode strings.

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

<img src="table/images/table-StyleColoredBright.png" width="640px"/>

More details can be found here: [table/](table)

## List

Pretty-print lists with multiple levels/indents into ASCII/Unicode strings.

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

More details can be found here: [list/](list)

# Progress

Track the Progress of one or more Tasks (like downloading multiple files in
parallel).

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

More details can be found here: [progress/](progress)

## Text

Utility functions to manipulate text with or without ANSI escape sequences. Most
of the functions available are used in one or more of the other packages here.

   - Align text horizontally or vertically
     - [text/align.go](text/align.go) and [text/valign.go](text/valign.go)
   - Colorize text
     - [text/color.go](text/color.go)
   - Cursor Movement
     - [text/cursor.go](text/cursor.go)
   - Format text (convert case)
     - [text/format.go](text/format.go)
   - String Manipulation (Pad, RepeatAndTrim, RuneCount, Trim, etc.)
     - [text/string.go](text/string.go)
   - Transform text (UnixTime to human-readable-time, pretty-JSON, etc.)
     - [text/transformer.go](text/transformer.go)
   - Wrap text
     - [text/wrap.go](text/wrap.go)

The unit-tests for each of the above show how these can be used. There GoDoc
should also have examples for all the available functions.

## Benchmarks

Partial output of `make bench` on CI:
```
BenchmarkList_Render-2            	  372352	      3179 ns/op	     856 B/op	      38 allocs/op
BenchmarkProgress_Render-2        	       4	 300318682 ns/op	    3438 B/op	      87 allocs/op
BenchmarkTable_Render-2           	   27208	     44154 ns/op	    5616 B/op	     179 allocs/op
BenchmarkTable_RenderCSV-2        	  108732	     11059 ns/op	    2624 B/op	      46 allocs/op
BenchmarkTable_RenderHTML-2       	   88633	     13425 ns/op	    4080 B/op	      45 allocs/op
BenchmarkTable_RenderMarkdown-2   	  107420	     10991 ns/op	    2560 B/op	      44 allocs/op
```

## v6.0.0++

If you are using a version of this library older than `v6.0.0` and want to move
to a newer version of this library, you'd have to modify the import paths from
something like:
```golang
    "github.com/jedib0t/go-pretty/list"
    "github.com/jedib0t/go-pretty/progress"
    "github.com/jedib0t/go-pretty/table"
    "github.com/jedib0t/go-pretty/text"
```
to:
```golang
    "github.com/jedib0t/go-pretty/v6/list"
    "github.com/jedib0t/go-pretty/v6/progress"
    "github.com/jedib0t/go-pretty/v6/table"
    "github.com/jedib0t/go-pretty/v6/text"
```

I'd recommend you fire up your favorite IDE and do a mass search and replace for
all occurrences of `jedib0t/go-pretty/` to `jedib0t/go-pretty/v6/`. If you are
on a system with access to `find`, `grep`, `xargs` and `sed`, you could just run
the following from within your code folder to do the same:
```
find . -type f -name "*.go" | grep -v vendor | xargs sed -i 's/jedib0t\/go-pretty\//jedib0t\/go-pretty\/v6\//'g
```
