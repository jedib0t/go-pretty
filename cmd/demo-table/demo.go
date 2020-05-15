package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

func demoTableColors() {
	tw := table.NewWriter()
	tw.AppendHeader(table.Row{"#", "First Name", "Last Name", "Salary"})
	tw.AppendRows([]table.Row{
		{1, "Arya", "Stark", 3000},
		{20, "Jon", "Snow", 2000, "You know nothing, Jon Snow!"},
		{300, "Tyrion", "Lannister", 5000},
	})
	tw.AppendFooter(table.Row{"", "", "Total", 10000})
	tw.SetIndexColumn(1)
	tw.SetTitle("Game Of Thrones")

	stylePairs := [][]table.Style{
		{table.StyleColoredBright, table.StyleColoredDark},
		{table.StyleColoredBlackOnBlueWhite, table.StyleColoredBlueWhiteOnBlack},
		{table.StyleColoredBlackOnCyanWhite, table.StyleColoredCyanWhiteOnBlack},
		{table.StyleColoredBlackOnGreenWhite, table.StyleColoredGreenWhiteOnBlack},
		{table.StyleColoredBlackOnMagentaWhite, table.StyleColoredMagentaWhiteOnBlack},
		{table.StyleColoredBlackOnRedWhite, table.StyleColoredRedWhiteOnBlack},
		{table.StyleColoredBlackOnYellowWhite, table.StyleColoredYellowWhiteOnBlack},
	}

	twOuter := table.NewWriter()
	twOuter.AppendHeader(table.Row{"Bright", "Dark"})
	for _, stylePair := range stylePairs {
		row := make(table.Row, 2)
		for idx, style := range stylePair {
			tw.SetCaption(style.Name)
			tw.SetStyle(style)
			tw.Style().Title.Align = text.AlignCenter
			row[idx] = tw.Render()
		}
		twOuter.AppendRow(row)
	}
	twOuter.SetColumnConfigs([]table.ColumnConfig{
		{Name: "Bright", Align: text.AlignCenter, AlignHeader: text.AlignCenter},
		{Name: "Dark", Align: text.AlignCenter, AlignHeader: text.AlignCenter},
	})
	twOuter.SetStyle(table.StyleLight)
	twOuter.Style().Title.Align = text.AlignCenter
	twOuter.SetTitle("C O L O R S")
	twOuter.Style().Options.SeparateRows = true
	fmt.Println(twOuter.Render())
}

func demoTableFeatures() {
	//==========================================================================
	// Initialization
	//==========================================================================
	t := table.NewWriter()
	// you can also instantiate the object directly
	tTemp := table.Table{}
	tTemp.Render() // just to avoid the compile error of not using the object
	//==========================================================================

	//==========================================================================
	// Append a few rows and render to console
	//==========================================================================
	// a row need not be just strings
	t.AppendRow(table.Row{1, "Arya", "Stark", 3000})
	// all rows need not have the same number of columns
	t.AppendRow(table.Row{20, "Jon", "Snow", 2000, "You know nothing, Jon Snow!"})
	// table.Row is just a shorthand for []interface{}
	t.AppendRow([]interface{}{300, "Tyrion", "Lannister", 5000})
	// time to take a peek
	t.SetCaption("Simple Table with 3 Rows.\n")
	fmt.Println(t.Render())
	//+-----+--------+-----------+------+-----------------------------+
	//|   1 | Arya   | Stark     | 3000 |                             |
	//|  20 | Jon    | Snow      | 2000 | You know nothing, Jon Snow! |
	//| 300 | Tyrion | Lannister | 5000 |                             |
	//+-----+--------+-----------+------+-----------------------------+
	//Simple Table with 3 Rows and a separator.
	//==========================================================================

	//==========================================================================
	// Can you index the columns?
	//==========================================================================
	t.SetAutoIndex(true)
	t.SetCaption("Table with Auto-Indexing.\n")
	fmt.Println(t.Render())
	//+---+-----+--------+-----------+------+-----------------------------+
	//|   |  A  |    B   |     C     |   D  |              E              |
	//+---+-----+--------+-----------+------+-----------------------------+
	//| 1 |   1 | Arya   | Stark     | 3000 |                             |
	//| 2 |  20 | Jon    | Snow      | 2000 | You know nothing, Jon Snow! |
	//| 3 | 300 | Tyrion | Lannister | 5000 |                             |
	//+---+-----+--------+-----------+------+-----------------------------+
	//Table with Auto-Indexing.
	//
	t.AppendHeader(table.Row{"#", "First Name", "Last Name", "Salary"})
	t.SetCaption("Table with Auto-Indexing (columns-only).\n")
	fmt.Println(t.Render())
	//+---+-----+------------+-----------+--------+-----------------------------+
	//|   |   # | FIRST NAME | LAST NAME | SALARY |                             |
	//+---+-----+------------+-----------+--------+-----------------------------+
	//| 1 |   1 | Arya       | Stark     |   3000 |                             |
	//| 2 |  20 | Jon        | Snow      |   2000 | You know nothing, Jon Snow! |
	//| 3 | 300 | Tyrion     | Lannister |   5000 |                             |
	//+---+-----+------------+-----------+--------+-----------------------------+
	//==========================================================================

	//==========================================================================
	// A table needs to have a Header & Footer (for this demo at least!)
	//==========================================================================
	t.SetAutoIndex(false)
	t.SetCaption("Table with 3 Rows & and a Header.\n")
	fmt.Println(t.Render())
	//+-----+------------+-----------+--------+-----------------------------+
	//|   # | FIRST NAME | LAST NAME | SALARY |                             |
	//+-----+------------+-----------+--------+-----------------------------+
	//|   1 | Arya       | Stark     |   3000 |                             |
	//|  20 | Jon        | Snow      |   2000 | You know nothing, Jon Snow! |
	//| 300 | Tyrion     | Lannister |   5000 |                             |
	//+-----+------------+-----------+--------+-----------------------------+
	//Table with 3 Rows & and a Header.
	//
	// and then add a footer
	t.AppendFooter(table.Row{"", "", "Total", 10000})
	// time to take a peek
	t.SetCaption("Table with 3 Rows, a Header & a Footer.\n")
	fmt.Println(t.Render())
	//+-----+------------+-----------+--------+-----------------------------+
	//|   # | FIRST NAME | LAST NAME | SALARY |                             |
	//+-----+------------+-----------+--------+-----------------------------+
	//|   1 | Arya       | Stark     |   3000 |                             |
	//|  20 | Jon        | Snow      |   2000 | You know nothing, Jon Snow! |
	//| 300 | Tyrion     | Lannister |   5000 |                             |
	//+-----+------------+-----------+--------+-----------------------------+
	//|     |            | TOTAL     |  10000 |                             |
	//+-----+------------+-----------+--------+-----------------------------+
	//Table with 3 Rows, a Header & a Footer.
	//==========================================================================

	//==========================================================================
	// Alignment?
	//==========================================================================
	// did you notice that the numeric columns were auto-aligned? when you don't
	// specify alignment, all the columns default to text.AlignDefault - numbers
	// go right and everything else left. but what if you want the first name to
	// go right too? and the last column to be "justified"?
	t.SetColumnConfigs([]table.ColumnConfig{
		{Name: "First Name", Align: text.AlignRight},
		// the 5th column does not have a title, so use the column number as the
		// identifier for the column
		{Number: 5, Align: text.AlignJustify},
	})
	// to show AlignJustify in action, lets add one more row
	t.AppendRow(table.Row{4, "Faceless", "Man", 0, "Needs a\tname."})
	// time to take a peek:
	t.SetCaption("Table with Custom Alignment for 2 columns.\n")
	fmt.Println(t.Render())
	//+-----+------------+-----------+--------+-----------------------------+
	//|   # | FIRST NAME | LAST NAME | SALARY |                             |
	//+-----+------------+-----------+--------+-----------------------------+
	//|   1 |       Arya | Stark     |   3000 |                             |
	//|  20 |        Jon | Snow      |   2000 | You know nothing, Jon Snow! |
	//| 300 |     Tyrion | Lannister |   5000 |                             |
	//|   4 |   Faceless | Man       |      0 | Needs        a        name. |
	//+-----+------------+-----------+--------+-----------------------------+
	//|     |            | TOTAL     |  10000 |                             |
	//+-----+------------+-----------+--------+-----------------------------+
	//Table with Custom Alignment for 2 columns.
	//==========================================================================

	//==========================================================================
	// Vertical Alignment?
	//==========================================================================
	// horizontal alignment is fine... what about vertical? lets add a row with
	// a column having multiple lines; and then play with VAlign
	t.AppendRow(table.Row{13, "Winter\nIs\nComing", "Valar\nMorghulis", 0, "You\n know\n  nothing,\n   Jon\n    Snow!"})
	// first without any custom VAlign
	t.SetCaption("Table with a Multi-line Row.\n")
	fmt.Println(t.Render())
	//+-----+------------+-----------+--------+-----------------------------+
	//|   # | FIRST NAME | LAST NAME | SALARY |                             |
	//+-----+------------+-----------+--------+-----------------------------+
	//|   1 |       Arya | Stark     |   3000 |                             |
	//|  20 |        Jon | Snow      |   2000 | You know nothing, Jon Snow! |
	//| 300 |     Tyrion | Lannister |   5000 |                             |
	//|   4 |   Faceless | Man       |      0 | Needs        a        name. |
	//|  13 |     Winter | Valar     |      0 | You                         |
	//|     |         Is | Morghulis |        | know                        |
	//|     |     Coming |           |        | nothing,                    |
	//|     |            |           |        | Jon                         |
	//|     |            |           |        | Snow!                       |
	//+-----+------------+-----------+--------+-----------------------------+
	//|     |            | TOTAL     |  10000 |                             |
	//+-----+------------+-----------+--------+-----------------------------+
	//Table with a Multi-line Row.
	//
	// time to Align/VAlign the columns...
	t.SetColumnConfigs([]table.ColumnConfig{
		{Name: "First Name", Align: text.AlignRight, VAlign: text.VAlignMiddle},
		{Name: "Last Name", VAlign: text.VAlignBottom},
		{Name: "Salary", Align: text.AlignRight, VAlign: text.VAlignMiddle},
		// the 5th column does not have a title, so use the column number
		{Number: 5, Align: text.AlignJustify},
	})
	t.SetCaption("Table with a Multi-line Row with VAlign.\n")
	fmt.Println(t.Render())
	//+-----+------------+-----------+--------+-----------------------------+
	//|   # | FIRST NAME | LAST NAME | SALARY |                             |
	//+-----+------------+-----------+--------+-----------------------------+
	//|   1 |       Arya | Stark     |   3000 |                             |
	//|  20 |        Jon | Snow      |   2000 | You know nothing, Jon Snow! |
	//| 300 |     Tyrion | Lannister |   5000 |                             |
	//|   4 |   Faceless | Man       |      0 | Needs        a        name. |
	//|  13 |            |           |        | You                         |
	//|     |     Winter |           |        | know                        |
	//|     |         Is |           |      0 | nothing,                    |
	//|     |     Coming | Valar     |        | Jon                         |
	//|     |            | Morghulis |        | Snow!                       |
	//+-----+------------+-----------+--------+-----------------------------+
	//|     |            | TOTAL     |  10000 |                             |
	//+-----+------------+-----------+--------+-----------------------------+
	//Table with a Multi-line Row with VAlign.
	//
	// changed your mind about AlignJustify?
	t.SetColumnConfigs([]table.ColumnConfig{
		{Name: "First Name", Align: text.AlignRight, VAlign: text.VAlignMiddle},
		{Name: "Last Name", VAlign: text.VAlignBottom},
		{Name: "Salary", Align: text.AlignRight, VAlign: text.VAlignMiddle},
		{Number: 5, Align: text.AlignCenter},
	})
	t.SetCaption("Table with a Multi-line Row with VAlign and changed Align.\n")
	fmt.Println(t.Render())
	//+-----+------------+-----------+--------+-----------------------------+
	//|   # | FIRST NAME | LAST NAME | SALARY |                             |
	//+-----+------------+-----------+--------+-----------------------------+
	//|   1 |       Arya | Stark     |   3000 |                             |
	//|  20 |        Jon | Snow      |   2000 | You know nothing, Jon Snow! |
	//| 300 |     Tyrion | Lannister |   5000 |                             |
	//|   4 |   Faceless | Man       |      0 |       Needs a    name.      |
	//|  13 |            |           |        |             You             |
	//|     |     Winter |           |        |             know            |
	//|     |         Is |           |      0 |           nothing,          |
	//|     |     Coming | Valar     |        |             Jon             |
	//|     |            | Morghulis |        |            Snow!            |
	//+-----+------------+-----------+--------+-----------------------------+
	//|     |            | TOTAL     |  10000 |                             |
	//+-----+------------+-----------+--------+-----------------------------+
	//Table with a Multi-line Row with VAlign and changed Align.
	//==========================================================================

	//==========================================================================
	// Time to begin anew. Too much on the screen for a demo! How about some
	// custom separators?
	//==========================================================================
	t.ResetRows()
	t.AppendRow(table.Row{1, "Arya", "Stark", 3000})
	t.AppendRow(table.Row{20, "Jon", "Snow", 2000, "You know nothing, Jon Snow!"})
	t.AppendSeparator()
	t.AppendRow([]interface{}{300, "Tyrion", "Lannister", 5000})
	t.SetCaption("Simple Table with 3 Rows and a Separator in-between.\n")
	fmt.Println(t.Render())
	//+-----+--------+-----------+------+-----------------------------+
	//|   1 | Arya   | Stark     | 3000 |                             |
	//|  20 | Jon    | Snow      | 2000 | You know nothing, Jon Snow! |
	//+-----+--------+-----------+------+-----------------------------+
	//| 300 | Tyrion | Lannister | 5000 |                             |
	//+-----+--------+-----------+------+-----------------------------+
	//Simple Table with 3 Rows and a Separator in-between.
	//==========================================================================

	//==========================================================================
	// Never-mind, lets start over yet again!
	//==========================================================================
	t.ResetRows()
	t.SetColumnConfigs(nil)
	t.AppendRow(table.Row{1, "Arya", "Stark", 3000})
	t.AppendRow(table.Row{20, "Jon", "Snow", 2000, "You know nothing, Jon Snow!"})
	t.AppendRow([]interface{}{300, "Tyrion", "Lannister", 5000})
	t.SetCaption("Starting afresh with a Simple Table again.\n")
	fmt.Println(t.Render())
	//+-----+------------+-----------+--------+-----------------------------+
	//|   # | FIRST NAME | LAST NAME | SALARY |                             |
	//+-----+------------+-----------+--------+-----------------------------+
	//|   1 | Arya       | Stark     |   3000 |                             |
	//|  20 | Jon        | Snow      |   2000 | You know nothing, Jon Snow! |
	//| 300 | Tyrion     | Lannister |   5000 |                             |
	//+-----+------------+-----------+--------+-----------------------------+
	//|     |            | TOTAL     |  10000 |                             |
	//+-----+------------+-----------+--------+-----------------------------+
	//Starting afresh with a Simple Table again.
	//==========================================================================

	//==========================================================================
	// Does it support paging?
	//==========================================================================
	t.SetPageSize(1)
	t.Style().Box.PageSeparator = "\n... page break ..."
	t.SetCaption("Table with a PageSize of 1.\n")
	fmt.Println(t.Render())
	//+-----+------------+-----------+--------+-----------------------------+
	//|   # | FIRST NAME | LAST NAME | SALARY |                             |
	//+-----+------------+-----------+--------+-----------------------------+
	//|   1 | Arya       | Stark     |   3000 |                             |
	//+-----+------------+-----------+--------+-----------------------------+
	//|     |            | TOTAL     |  10000 |                             |
	//+-----+------------+-----------+--------+-----------------------------+
	//... page break ...
	//+-----+------------+-----------+--------+-----------------------------+
	//|   # | FIRST NAME | LAST NAME | SALARY |                             |
	//+-----+------------+-----------+--------+-----------------------------+
	//|  20 | Jon        | Snow      |   2000 | You know nothing, Jon Snow! |
	//+-----+------------+-----------+--------+-----------------------------+
	//|     |            | TOTAL     |  10000 |                             |
	//+-----+------------+-----------+--------+-----------------------------+
	//... page break ...
	//+-----+------------+-----------+--------+-----------------------------+
	//|   # | FIRST NAME | LAST NAME | SALARY |                             |
	//+-----+------------+-----------+--------+-----------------------------+
	//| 300 | Tyrion     | Lannister |   5000 |                             |
	//+-----+------------+-----------+--------+-----------------------------+
	//|     |            | TOTAL     |  10000 |                             |
	//+-----+------------+-----------+--------+-----------------------------+
	//Table with a PageSize of 1.
	t.SetPageSize(0) // disables paging
	//==========================================================================

	//==========================================================================
	// How about limiting the length of every Row?
	//==========================================================================
	t.SetAllowedRowLength(50)
	t.SetCaption("Table with an Allowed Row Length of 50.\n")
	fmt.Println(t.Render())
	//+-----+------------+-----------+--------+------- ~
	//|   # | FIRST NAME | LAST NAME | SALARY |        ~
	//+-----+------------+-----------+--------+------- ~
	//|   1 | Arya       | Stark     |   3000 |        ~
	//|  20 | Jon        | Snow      |   2000 | You kn ~
	//| 300 | Tyrion     | Lannister |   5000 |        ~
	//+-----+------------+-----------+--------+------- ~
	//|     |            | TOTAL     |  10000 |        ~
	//+-----+------------+-----------+--------+------- ~
	t.SetStyle(table.StyleDouble)
	t.SetCaption("Table with an Allowed Row Length of 50 in 'StyleDouble'.\n")
	fmt.Println(t.Render())
	//╔═════╦════════════╦═══════════╦════════╦═══════ ≈
	//║   # ║ FIRST NAME ║ LAST NAME ║ SALARY ║        ≈
	//╠═════╬════════════╬═══════════╬════════╬═══════ ≈
	//║   1 ║ Arya       ║ Stark     ║   3000 ║        ≈
	//║  20 ║ Jon        ║ Snow      ║   2000 ║ You kn ≈
	//║ 300 ║ Tyrion     ║ Lannister ║   5000 ║        ≈
	//╠═════╬════════════╬═══════════╬════════╬═══════ ≈
	//║     ║            ║ TOTAL     ║  10000 ║        ≈
	//╚═════╩════════════╩═══════════╩════════╩═══════ ≈
	//Table with an Allowed Row Length of 50 in 'StyleDouble'.
	//==========================================================================

	//==========================================================================
	// But I want to see all the data!
	//==========================================================================
	t.SetColumnConfigs([]table.ColumnConfig{
		{Name: "First Name", WidthMax: 6},
		{Name: "Last Name", WidthMax: 9},
		{Name: "Salary", WidthMax: 6},
		{Number: 5, WidthMax: 10},
	})
	t.SetCaption("Table on a diet.\n")
	t.SetStyle(table.StyleRounded)
	fmt.Println(t.Render())
	//╭─────┬────────┬───────────┬────────┬────────────╮
	//│   # │ FIRST  │ LAST NAME │ SALARY │            │
	//│     │ NAME   │           │        │            │
	//├─────┼────────┼───────────┼────────┼────────────┤
	//│   1 │ Arya   │ Stark     │   3000 │            │
	//│  20 │ Jon    │ Snow      │   2000 │ You know n │
	//│     │        │           │        │ othing, Jo │
	//│     │        │           │        │ n Snow!    │
	//│ 300 │ Tyrion │ Lannister │   5000 │            │
	//├─────┼────────┼───────────┼────────┼────────────┤
	//│     │        │ TOTAL     │  10000 │            │
	//╰─────┴────────┴───────────┴────────┴────────────╯
	//Table on a diet.
	t.SetAllowedRowLength(0)
	// remove the width restrictions
	t.SetColumnConfigs([]table.ColumnConfig{})
	//==========================================================================

	//==========================================================================
	// ASCII is too simple for me.
	//==========================================================================
	t.SetStyle(table.StyleLight)
	t.SetCaption("Table using the style 'StyleLight'.\n")
	fmt.Println(t.Render())
	//┌─────┬────────────┬───────────┬────────┬─────────────────────────────┐
	//│   # │ FIRST NAME │ LAST NAME │ SALARY │                             │
	//├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
	//│   1 │ Arya       │ Stark     │   3000 │                             │
	//│  20 │ Jon        │ Snow      │   2000 │ You know nothing, Jon Snow! │
	//│ 300 │ Tyrion     │ Lannister │   5000 │                             │
	//├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
	//│     │            │ TOTAL     │  10000 │                             │
	//└─────┴────────────┴───────────┴────────┴─────────────────────────────┘
	//Table using the style 'StyleLight'.
	t.SetStyle(table.StyleDouble)
	t.SetCaption("Table using the style '%s'.\n", t.Style().Name)
	fmt.Println(t.Render())
	//╔═════╦════════════╦═══════════╦════════╦═════════════════════════════╗
	//║   # ║ FIRST NAME ║ LAST NAME ║ SALARY ║                             ║
	//╠═════╬════════════╬═══════════╬════════╬═════════════════════════════╣
	//║   1 ║ Arya       ║ Stark     ║   3000 ║                             ║
	//║  20 ║ Jon        ║ Snow      ║   2000 ║ You know nothing, Jon Snow! ║
	//║ 300 ║ Tyrion     ║ Lannister ║   5000 ║                             ║
	//╠═════╬════════════╬═══════════╬════════╬═════════════════════════════╣
	//║     ║            ║ TOTAL     ║  10000 ║                             ║
	//╚═════╩════════════╩═══════════╩════════╩═════════════════════════════╝
	//Table using the style 'StyleDouble'.
	//==========================================================================

	//==========================================================================
	// I don't like any of the ready-made styles.
	//==========================================================================
	t.SetStyle(table.Style{
		Name: "funkyStyle",
		Box: table.BoxStyle{
			BottomLeft:       "\\",
			BottomRight:      "/",
			BottomSeparator:  "v",
			Left:             "[",
			LeftSeparator:    "{",
			MiddleHorizontal: "-",
			MiddleSeparator:  "+",
			MiddleVertical:   "|",
			PaddingLeft:      "<",
			PaddingRight:     ">",
			Right:            "]",
			RightSeparator:   "}",
			TopLeft:          "(",
			TopRight:         ")",
			TopSeparator:     "^",
			UnfinishedRow:    " ~~~",
		},
	})
	t.Style().Format = table.FormatOptions{
		Footer: text.FormatLower,
		Header: text.FormatLower,
		Row:    text.FormatUpper,
	}
	t.Style().Options.DrawBorder = true
	t.Style().Options.SeparateColumns = true
	t.Style().Options.SeparateFooter = true
	t.Style().Options.SeparateHeader = true
	t.SetCaption("Table using the style 'funkyStyle'.\n")
	fmt.Println(t.Render())
	//(-----^------------^-----------^--------^-----------------------------)
	//[<  #>|<first name>|<last name>|<salary>|<                           >]
	//{-----+------------+-----------+--------+-----------------------------}
	//[<  1>|<ARYA      >|<STARK    >|<  3000>|<                           >]
	//[< 20>|<JON       >|<SNOW     >|<  2000>|<YOU KNOW NOTHING, JON SNOW!>]
	//[<300>|<TYRION    >|<LANNISTER>|<  5000>|<                           >]
	//{-----+------------+-----------+--------+-----------------------------}
	//[<   >|<          >|<total    >|< 10000>|<                           >]
	//\-----v------------v-----------v--------v-----------------------------/
	//Table using the style 'funkyStyle'.
	//==========================================================================

	//==========================================================================
	// I need some color in my life!
	//==========================================================================
	t.SetStyle(table.StyleBold)
	colorBOnW := text.Colors{text.BgWhite, text.FgBlack}
	// set colors using Colors/ColorsHeader/ColorsFooter
	t.SetColumnConfigs([]table.ColumnConfig{
		{Name: "#", Colors: text.Colors{text.FgYellow}, ColorsHeader: colorBOnW},
		{Name: "First Name", Colors: text.Colors{text.FgHiRed}, ColorsHeader: colorBOnW},
		{Name: "Last Name", Colors: text.Colors{text.FgHiRed}, ColorsHeader: colorBOnW, ColorsFooter: colorBOnW},
		{Name: "Salary", Colors: text.Colors{text.FgGreen}, ColorsHeader: colorBOnW, ColorsFooter: colorBOnW},
		{Number: 5, Colors: text.Colors{text.FgCyan}, ColorsHeader: colorBOnW},
	})
	t.SetCaption("Table with Colors.\n")
	fmt.Println(t.Render())
	//┏━━━━━┳━━━━━━━━━━━━┳━━━━━━━━━━━┳━━━━━━━━┳━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
	//┃   # ┃ FIRST NAME ┃ LAST NAME ┃ SALARY ┃                             ┃
	//┣━━━━━╋━━━━━━━━━━━━╋━━━━━━━━━━━╋━━━━━━━━╋━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
	//┃   1 ┃ Arya       ┃ Stark     ┃   3000 ┃                             ┃
	//┃  20 ┃ Jon        ┃ Snow      ┃   2000 ┃ You know nothing, Jon Snow! ┃
	//┃ 300 ┃ Tyrion     ┃ Lannister ┃   5000 ┃                             ┃
	//┣━━━━━╋━━━━━━━━━━━━╋━━━━━━━━━━━╋━━━━━━━━╋━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
	//┃     ┃            ┃ TOTAL     ┃  10000 ┃                             ┃
	//┗━━━━━┻━━━━━━━━━━━━┻━━━━━━━━━━━┻━━━━━━━━┻━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛
	//Table with Colors.
	//
	// "Table with Colors"??? where? i don't see any! well, you have to trust me
	// on this... the colors show on a terminal that supports it. to prove it,
	// lets print the same table line-by-line using "%#v" to see the control
	// sequences ...
	t.SetCaption("Table with Colors in Raw Mode.\n")
	for _, line := range strings.Split(t.Render(), "\n") {
		if line != "" {
			fmt.Printf("%#v\n", line)
		}
	}
	fmt.Println()
	//"┏━━━━━┳━━━━━━━━━━━━┳━━━━━━━━━━━┳━━━━━━━━┳━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓"
	//"┃\x1b[47;30m   # \x1b[0m┃\x1b[47;30m FIRST NAME \x1b[0m┃\x1b[47;30m LAST NAME \x1b[0m┃\x1b[47;30m SALARY \x1b[0m┃\x1b[47;30m                             \x1b[0m┃"
	//"┣━━━━━╋━━━━━━━━━━━━╋━━━━━━━━━━━╋━━━━━━━━╋━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫"
	//"┃\x1b[33m   1 \x1b[0m┃\x1b[91m Arya       \x1b[0m┃\x1b[91m Stark     \x1b[0m┃\x1b[32m   3000 \x1b[0m┃\x1b[36m                             \x1b[0m┃"
	//"┃\x1b[33m  20 \x1b[0m┃\x1b[91m Jon        \x1b[0m┃\x1b[91m Snow      \x1b[0m┃\x1b[32m   2000 \x1b[0m┃\x1b[36m You know nothing, Jon Snow! \x1b[0m┃"
	//"┃\x1b[33m 300 \x1b[0m┃\x1b[91m Tyrion     \x1b[0m┃\x1b[91m Lannister \x1b[0m┃\x1b[32m   5000 \x1b[0m┃\x1b[36m                             \x1b[0m┃"
	//"┣━━━━━╋━━━━━━━━━━━━╋━━━━━━━━━━━╋━━━━━━━━╋━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫"
	//"┃     ┃            ┃\x1b[47;30m TOTAL     \x1b[0m┃\x1b[47;30m  10000 \x1b[0m┃                             ┃"
	//"┗━━━━━┻━━━━━━━━━━━━┻━━━━━━━━━━━┻━━━━━━━━┻━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛"
	//"Table with Colors in Raw Mode."
	//""
	// disable colors and revert to previous version of the column configs
	t.SetColumnConfigs([]table.ColumnConfig{})
	//==========================================================================

	//==========================================================================
	// How about not asking me to set colors in such a verbose way? And I don't
	// like wasting my terminal space with borders and separators.
	//==========================================================================
	t.SetStyle(table.StyleColoredBright)
	t.SetCaption("Table with style 'StyleColoredBright'.\n")
	fmt.Println(t.Render())
	//   #  FIRST NAME  LAST NAME  SALARY
	//   1  Arya        Stark        3000
	//  20  Jon         Snow         2000  You know nothing, Jon Snow!
	// 300  Tyrion      Lannister    5000
	//                  TOTAL       10000
	//Table with style 'StyleColoredBright'.
	t.SetStyle(table.StyleBold)
	//==========================================================================

	//==========================================================================
	// I don't like borders!
	//==========================================================================
	t.Style().Options.DrawBorder = false
	t.SetCaption("Table without Borders.\n")
	fmt.Println(t.Render())
	//   # ┃ FIRST NAME ┃ LAST NAME ┃ SALARY ┃
	//━━━━━╋━━━━━━━━━━━━╋━━━━━━━━━━━╋━━━━━━━━╋━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
	//   1 ┃ Arya       ┃ Stark     ┃   3000 ┃
	//  20 ┃ Jon        ┃ Snow      ┃   2000 ┃ You know nothing, Jon Snow!
	// 300 ┃ Tyrion     ┃ Lannister ┃   5000 ┃
	//━━━━━╋━━━━━━━━━━━━╋━━━━━━━━━━━╋━━━━━━━━╋━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
	//     ┃            ┃ TOTAL     ┃  10000 ┃
	//Table without Borders.
	//==========================================================================

	//==========================================================================
	// I like walls and borders everywhere!
	//==========================================================================
	t.Style().Options.DrawBorder = true
	t.Style().Options.SeparateRows = true
	t.SetCaption("Table with Borders Everywhere!\n")
	t.SetTitle("Divide!")
	fmt.Println(t.Render())
	//┏━━━━━┳━━━━━━━━━━━━┳━━━━━━━━━━━┳━━━━━━━━┳━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
	//┃   # ┃ FIRST NAME ┃ LAST NAME ┃ SALARY ┃                             ┃
	//┣━━━━━╋━━━━━━━━━━━━╋━━━━━━━━━━━╋━━━━━━━━╋━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
	//┃   1 ┃ Arya       ┃ Stark     ┃   3000 ┃                             ┃
	//┣━━━━━╋━━━━━━━━━━━━╋━━━━━━━━━━━╋━━━━━━━━╋━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
	//┃  20 ┃ Jon        ┃ Snow      ┃   2000 ┃ You know nothing, Jon Snow! ┃
	//┣━━━━━╋━━━━━━━━━━━━╋━━━━━━━━━━━╋━━━━━━━━╋━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
	//┃ 300 ┃ Tyrion     ┃ Lannister ┃   5000 ┃                             ┃
	//┣━━━━━╋━━━━━━━━━━━━╋━━━━━━━━━━━╋━━━━━━━━╋━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
	//┃     ┃            ┃ TOTAL     ┃  10000 ┃                             ┃
	//┗━━━━━┻━━━━━━━━━━━━┻━━━━━━━━━━━┻━━━━━━━━┻━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛
	//Table with Borders Everywhere!
	//==========================================================================

	//==========================================================================
	// There is strength in Unity.
	//==========================================================================
	t.Style().Options.DrawBorder = false
	t.Style().Options.SeparateColumns = false
	t.Style().Options.SeparateFooter = false
	t.Style().Options.SeparateHeader = false
	t.Style().Options.SeparateRows = false
	t.SetCaption("(c) No one!")
	t.SetTitle("Unite!")
	fmt.Println(t.Render())
	fmt.Println()
	//   #  FIRST NAME  LAST NAME  SALARY
	//   1  Arya        Stark        3000
	//  20  Jon         Snow         2000  You know nothing, Jon Snow!
	// 300  Tyrion      Lannister    5000
	//                  TOTAL       10000
	//Table without Any Borders or Separators!
	//==========================================================================

	//==========================================================================
	// I want CSV.
	//==========================================================================
	for _, line := range strings.Split(t.RenderCSV(), "\n") {
		fmt.Printf("[CSV] %s\n", line)
	}
	fmt.Println()
	//[CSV] #,First Name,Last Name,Salary,
	//[CSV] 1,Arya,Stark,3000,
	//[CSV] 20,Jon,Snow,2000,"You know nothing\, Jon Snow!"
	//[CSV] 300,Tyrion,Lannister,5000,
	//[CSV] ,,Total,10000,
	//==========================================================================

	//==========================================================================
	// Nope. I want a HTML Table.
	//==========================================================================
	for _, line := range strings.Split(t.RenderHTML(), "\n") {
		fmt.Printf("[HTML] %s\n", line)
	}
	fmt.Println()
	//[HTML] <table class="go-pretty-table">
	//[HTML]   <thead>
	//[HTML]   <tr>
	//[HTML]     <th align="right">#</th>
	//[HTML]     <th>First Name</th>
	//[HTML]     <th>Last Name</th>
	//[HTML]     <th align="right">Salary</th>
	//[HTML]     <th>&nbsp;</th>
	//[HTML]   </tr>
	//[HTML]   </thead>
	//[HTML]   <tbody>
	//[HTML]   <tr>
	//[HTML]     <td align="right">1</td>
	//[HTML]     <td>Arya</td>
	//[HTML]     <td>Stark</td>
	//[HTML]     <td align="right">3000</td>
	//[HTML]     <td>&nbsp;</td>
	//[HTML]   </tr>
	//[HTML]   <tr>
	//[HTML]     <td align="right">20</td>
	//[HTML]     <td>Jon</td>
	//[HTML]     <td>Snow</td>
	//[HTML]     <td align="right">2000</td>
	//[HTML]     <td>You know nothing, Jon Snow!</td>
	//[HTML]   </tr>
	//[HTML]   <tr>
	//[HTML]     <td align="right">300</td>
	//[HTML]     <td>Tyrion</td>
	//[HTML]     <td>Lannister</td>
	//[HTML]     <td align="right">5000</td>
	//[HTML]     <td>&nbsp;</td>
	//[HTML]   </tr>
	//[HTML]   </tbody>
	//[HTML]   <tfoot>
	//[HTML]   <tr>
	//[HTML]     <td align="right">&nbsp;</td>
	//[HTML]     <td>&nbsp;</td>
	//[HTML]     <td>Total</td>
	//[HTML]     <td align="right">10000</td>
	//[HTML]     <td>&nbsp;</td>
	//[HTML]   </tr>
	//[HTML]   </tfoot>
	//[HTML] </table>
	//==========================================================================

	//==========================================================================
	// Nope. I want a Markdown Table now.
	//==========================================================================
	for _, line := range strings.Split(t.RenderMarkdown(), "\n") {
		fmt.Printf("[Markdown] %s\n", line)
	}
	fmt.Println()
	//[Markdown] | # | First Name | Last Name | Salary |  |
	//[Markdown] | ---:| --- | --- | ---:| --- |
	//[Markdown] | 1 | Arya | Stark | 3000 |  |
	//[Markdown] | 20 | Jon | Snow | 2000 | You know nothing, Jon Snow! |
	//[Markdown] | 300 | Tyrion | Lannister | 5000 |  |
	//[Markdown] |  |  | Total | 10000 |  |
	//==========================================================================

	//==========================================================================
	// That's it for today! New features will always find a place in this demo!
	//==========================================================================
}

func demoTableEmoji() {
	styles := []table.Style{
		table.StyleDefault,
		table.StyleLight,
		table.StyleColoredBright,
	}
	for _, style := range styles {
		tw := table.NewWriter()
		tw.AppendHeader(table.Row{"Key", "Value"})
		tw.AppendRows([]table.Row{
			{"Emoji 1 🥰", 1000},
			{"Emoji 2 ⚔️", 2000},
			{"Emoji 3 🎁", 3000},
			{"Emoji 4 ツ", 4000},
		})
		tw.AppendFooter(table.Row{"Total", 10000})
		tw.SetAutoIndex(true)
		tw.SetStyle(style)

		fmt.Println(tw.Render())
		fmt.Println()
	}
}

func main() {
	demoWhat := "features"
	if len(os.Args) > 1 {
		demoWhat = os.Args[1]
	}

	switch strings.ToLower(demoWhat) {
	case "colors":
		demoTableColors()
	case "emoji":
		demoTableEmoji()
	default:
		demoTableFeatures()
	}
}
