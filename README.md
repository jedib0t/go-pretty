# go-pretty

[![Go Reference](https://pkg.go.dev/badge/github.com/jedib0t/go-pretty/v6.svg)](https://pkg.go.dev/github.com/jedib0t/go-pretty/v6)
[![Build Status](https://github.com/jedib0t/go-pretty/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/jedib0t/go-pretty/actions?query=workflow%3ACI+event%3Apush+branch%3Amain)
[![Coverage Status](https://coveralls.io/repos/github/jedib0t/go-pretty/badge.svg?branch=main)](https://coveralls.io/github/jedib0t/go-pretty?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/jedib0t/go-pretty/v6)](https://goreportcard.com/report/github.com/jedib0t/go-pretty/v6)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=jedib0t_go-pretty&metric=alert_status)](https://sonarcloud.io/dashboard?id=jedib0t_go-pretty)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Utilities to prettify console output of tables, lists, progress bars, text, and more
with a heavy emphasis on customization and flexibility.

## Quick Start

```bash
go get github.com/jedib0t/go-pretty/v6
```

Import the packages you need:
```go
import (
    "github.com/jedib0t/go-pretty/v6/table"
    "github.com/jedib0t/go-pretty/v6/list"
    "github.com/jedib0t/go-pretty/v6/progress"
    "github.com/jedib0t/go-pretty/v6/text"
)
```

**Note**: Current major version is **v6**. See [Go modules versioning](https://go.dev/doc/modules/version-numbers#major-version) for details.

## Packages

### Table

Pretty-print tables with colors, auto-merge, sorting, paging, and multiple output formats (ASCII, HTML, Markdown, CSV, TSV).

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

Try the nested colored tables demo:
```bash
go run github.com/jedib0t/go-pretty/v6/cmd/demo-table@latest colors
```

<img src="cmd/demo-table/demo-colors.png" alt="Tables with Colors within a Table in a Terminal" width="640px"/>

📖 [Full documentation →](table/)

### Progress

Track progress of one or more tasks with ETA, speed calculation, indeterminate indicators, and customizable styles.

<img src="progress/images/demo.gif" alt="Progress Demo in a Terminal" width="640px"/>

📖 [Full documentation →](progress/)

### List

Pretty-print hierarchical lists with multiple levels, indentation, and multiple output formats (ASCII, HTML, Markdown).

```
╭─ Game Of Thrones
│  ├─ Winter
│  ├─ Is
│  ╰─ Coming
│     ├─ This
│     ├─ Is
│     ╰─ Known
╰─ The Dark Tower
   ╰─ The Gunslinger
```

📖 [Full documentation →](list/)

### Text

Utility functions to manipulate strings/text with full ANSI escape sequence support. Used extensively by other packages in this repo.

**Features**: Alignment (horizontal/vertical), colors & formatting, cursor control, text transformation (case, JSON, time, URLs), string manipulation (pad, trim, wrap), and more.

📖 [Full documentation →](text/)
