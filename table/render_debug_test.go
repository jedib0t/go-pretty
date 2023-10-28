package table

import (
	"strings"
	"testing"
)

func TestTable_Render_Debug(t *testing.T) {
	debugData := strings.Builder{}
	tw := Table{}
	tw.AppendHeader(testHeader)
	tw.AppendRows(testRows)
	tw.AppendFooter(testFooter)
	tw.SetDebugWriter(&debugData)

	compareOutput(t, tw.Render(), `
+-----+------------+-----------+--------+-----------------------------+
|   # | FIRST NAME | LAST NAME | SALARY |                             |
+-----+------------+-----------+--------+-----------------------------+
|   1 | Arya       | Stark     |   3000 |                             |
|  20 | Jon        | Snow      |   2000 | You know nothing, Jon Snow! |
| 300 | Tyrion     | Lannister |   5000 |                             |
+-----+------------+-----------+--------+-----------------------------+
|     |            | TOTAL     |  10000 |                             |
+-----+------------+-----------+--------+-----------------------------+`)

	compareOutput(t, debugData.String(), `(table.Table) {
    allowedRowLength: (int) 0,
    autoIndex: (bool) false,
    autoIndexVIndexMaxLength: (int) 1,
    caption: (string) "",
    columnIsNonNumeric: ([]bool) (len=5) {
        (bool) false,
        (bool) true,
        (bool) true,
        (bool) false,
        (bool) true
    },
    columnConfigs: ([]table.ColumnConfig) <nil>,
    columnConfigMap: (map[int]table.ColumnConfig) {
    },
    debugWriter: (*strings.Builder)(),
    htmlCSSClass: (string) "",
    indexColumn: (int) 0,
    maxColumnLengths: ([]int) (len=5) {
        (int) 3,
        (int) 10,
        (int) 9,
        (int) 6,
        (int) 27
    },
    maxRowLength: (int) 71,
    numColumns: (int) 5,
    numLinesRendered: (int) 0,
    outputMirror: (io.Writer) <nil>,
    pageSize: (int) 0,
    rows: ([]table.rowStr) (len=3) {
        (table.rowStr) (len=4) {
            (string) (len=1) "1",
            (string) (len=4) "Arya",
            (string) (len=5) "Stark",
            (string) (len=4) "3000"
        },
        (table.rowStr) (len=5) {
            (string) (len=2) "20",
            (string) (len=3) "Jon",
            (string) (len=4) "Snow",
            (string) (len=4) "2000",
            (string) (len=27) "You know nothing, Jon Snow!"
        },
        (table.rowStr) (len=4) {
            (string) (len=3) "300",
            (string) (len=6) "Tyrion",
            (string) (len=9) "Lannister",
            (string) (len=4) "5000"
        }
    },
    rowsColors: ([]text.Colors) <nil>,
    rowsConfigMap: (map[int]table.RowConfig) <nil>,
    rowsRaw: ([]table.Row) (len=3) {
        (table.Row) (len=4) {
            (int) 1,
            (string) (len=4) "Arya",
            (string) (len=5) "Stark",
            (int) 3000
        },
        (table.Row) (len=5) {
            (int) 20,
            (string) (len=3) "Jon",
            (string) (len=4) "Snow",
            (int) 2000,
            (string) (len=27) "You know nothing, Jon Snow!"
        },
        (table.Row) (len=4) {
            (int) 300,
            (string) (len=6) "Tyrion",
            (string) (len=9) "Lannister",
            (int) 5000
        }
    },
    rowsFooter: ([]table.rowStr) (len=1) {
        (table.rowStr) (len=4) {
            (string) "",
            (string) "",
            (string) (len=5) "Total",
            (string) (len=5) "10000"
        }
    },
    rowsFooterConfigMap: (map[int]table.RowConfig) <nil>,
    rowsFooterRaw: ([]table.Row) (len=1) {
        (table.Row) (len=4) {
            (string) "",
            (string) "",
            (string) (len=5) "Total",
            (int) 10000
        }
    },
    rowsHeader: ([]table.rowStr) (len=1) {
        (table.rowStr) (len=4) {
            (string) (len=1) "#",
            (string) (len=10) "First Name",
            (string) (len=9) "Last Name",
            (string) (len=6) "Salary"
        }
    },
    rowsHeaderConfigMap: (map[int]table.RowConfig) <nil>,
    rowsHeaderRaw: ([]table.Row) (len=1) {
        (table.Row) (len=4) {
            (string) (len=1) "#",
            (string) (len=10) "First Name",
            (string) (len=9) "Last Name",
            (string) (len=6) "Salary"
        }
    },
    rowPainter: (table.RowPainter) <nil>,
    rowSeparator: (table.rowStr) (len=5) {
        (string) (len=5) "-----",
        (string) (len=12) "------------",
        (string) (len=11) "-----------",
        (string) (len=8) "--------",
        (string) (len=29) "-----------------------------"
    },
    separators: (map[int]bool) <nil>,
    sortBy: ([]table.SortBy) <nil>,
    style: (*table.Style)({
        Name: (string) (len=12) "StyleDefault",
        Box: (table.BoxStyle) {
            BottomLeft: (string) (len=1) "+",
            BottomRight: (string) (len=1) "+",
            BottomSeparator: (string) (len=1) "+",
            EmptySeparator: (string) (len=1) " ",
            Left: (string) (len=1) "|",
            LeftSeparator: (string) (len=1) "+",
            MiddleHorizontal: (string) (len=1) "-",
            MiddleSeparator: (string) (len=1) "+",
            MiddleVertical: (string) (len=1) "|",
            PaddingLeft: (string) (len=1) " ",
            PaddingRight: (string) (len=1) " ",
            PageSeparator: (string) (len=1) "\n",
            Right: (string) (len=1) "|",
            RightSeparator: (string) (len=1) "+",
            TopLeft: (string) (len=1) "+",
            TopRight: (string) (len=1) "+",
            TopSeparator: (string) (len=1) "+",
            UnfinishedRow: (string) (len=2) " ~"
        },
        Color: (table.ColorOptions) {
            Border: (text.Colors) <nil>,
            Footer: (text.Colors) <nil>,
            Header: (text.Colors) <nil>,
            IndexColumn: (text.Colors) <nil>,
            Row: (text.Colors) <nil>,
            RowAlternate: (text.Colors) <nil>,
            Separator: (text.Colors) <nil>
        },
        Format: (table.FormatOptions) {
            Direction: (text.Direction) 0,
            Footer: (text.Format) 3,
            Header: (text.Format) 3,
            Row: (text.Format) 0
        },
        HTML: (table.HTMLOptions) {
            CSSClass: (string) (len=15) "go-pretty-table",
            EmptyColumn: (string) (len=6) "&nbsp;",
            EscapeText: (bool) true,
            Newline: (string) (len=5) "<br/>"
        },
        Options: (table.Options) {
            DoNotColorBordersAndSeparators: (bool) false,
            DrawBorder: (bool) true,
            SeparateColumns: (bool) true,
            SeparateFooter: (bool) true,
            SeparateHeader: (bool) true,
            SeparateRows: (bool) false
        },
        Title: (table.TitleOptions) {
            Align: (text.Align) 0,
            Colors: (text.Colors) <nil>,
            Format: (text.Format) 0
        }
    }),
    suppressEmptyColumns: (bool) false,
    title: (string) ""
}
>> grow buffer by 71 bytes
++ [00] ""
++ [01] "+"
++ [05] "-----"
++ [01] "+"
++ [12] "------------"
++ [01] "+"
++ [11] "-----------"
++ [01] "+"
++ [08] "--------"
++ [01] "+"
++ [29] "-----------------------------"
++ [01] "+"
++ [01] 10
>> grow buffer by 71 bytes
++ [00] ""
++ [01] "|"
++ [05] "   # "
++ [01] "|"
++ [12] " FIRST NAME "
++ [01] "|"
++ [11] " LAST NAME "
++ [01] "|"
++ [08] " SALARY "
++ [01] "|"
++ [29] "                             "
++ [01] "|"
++ [01] 10
>> grow buffer by 71 bytes
++ [00] ""
++ [01] "+"
++ [05] "-----"
++ [01] "+"
++ [12] "------------"
++ [01] "+"
++ [11] "-----------"
++ [01] "+"
++ [08] "--------"
++ [01] "+"
++ [29] "-----------------------------"
++ [01] "+"
++ [01] 10
>> grow buffer by 71 bytes
++ [00] ""
++ [01] "|"
++ [05] "   1 "
++ [01] "|"
++ [12] " Arya       "
++ [01] "|"
++ [11] " Stark     "
++ [01] "|"
++ [08] "   3000 "
++ [01] "|"
++ [29] "                             "
++ [01] "|"
++ [01] 10
>> grow buffer by 71 bytes
++ [00] ""
++ [01] "|"
++ [05] "  20 "
++ [01] "|"
++ [12] " Jon        "
++ [01] "|"
++ [11] " Snow      "
++ [01] "|"
++ [08] "   2000 "
++ [01] "|"
++ [29] " You know nothing, Jon Snow! "
++ [01] "|"
++ [01] 10
>> grow buffer by 71 bytes
++ [00] ""
++ [01] "|"
++ [05] " 300 "
++ [01] "|"
++ [12] " Tyrion     "
++ [01] "|"
++ [11] " Lannister "
++ [01] "|"
++ [08] "   5000 "
++ [01] "|"
++ [29] "                             "
++ [01] "|"
++ [01] 10
>> grow buffer by 71 bytes
++ [00] ""
++ [01] "+"
++ [05] "-----"
++ [01] "+"
++ [12] "------------"
++ [01] "+"
++ [11] "-----------"
++ [01] "+"
++ [08] "--------"
++ [01] "+"
++ [29] "-----------------------------"
++ [01] "+"
++ [01] 10
>> grow buffer by 71 bytes
++ [00] ""
++ [01] "|"
++ [05] "     "
++ [01] "|"
++ [12] "            "
++ [01] "|"
++ [11] " TOTAL     "
++ [01] "|"
++ [08] "  10000 "
++ [01] "|"
++ [29] "                             "
++ [01] "|"
++ [01] 10
>> grow buffer by 71 bytes
++ [00] ""
++ [01] "+"
++ [05] "-----"
++ [01] "+"
++ [12] "------------"
++ [01] "+"
++ [11] "-----------"
++ [01] "+"
++ [08] "--------"
++ [01] "+"
++ [29] "-----------------------------"
++ [01] "+"
[00] "+-----+------------+-----------+--------+-----------------------------+"
[01] "|   # | FIRST NAME | LAST NAME | SALARY |                             |"
[02] "+-----+------------+-----------+--------+-----------------------------+"
[03] "|   1 | Arya       | Stark     |   3000 |                             |"
[04] "|  20 | Jon        | Snow      |   2000 | You know nothing, Jon Snow! |"
[05] "| 300 | Tyrion     | Lannister |   5000 |                             |"
[06] "+-----+------------+-----------+--------+-----------------------------+"
[07] "|     |            | TOTAL     |  10000 |                             |"
[08] "+-----+------------+-----------+--------+-----------------------------+"
`)
}
