package main

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/table"
	"github.com/jedib0t/go-pretty/text"
)

func main() {
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
	//Simple Table with 3 Rows.
	//==========================================================================

	//==========================================================================
	// A table needs to have a Header & Footer (for this demo at least!)
	//==========================================================================
	// start with a header
	t.AppendHeader(table.Row{"#", "First Name", "Last Name", "Salary"})
	// time to take a peek
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
	t.SetAlign([]text.Align{text.AlignDefault, text.AlignRight, text.AlignDefault, text.AlignDefault, text.AlignJustify})
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
	// time to VAlign the columns... and ignore the last column in the process
	t.SetVAlign([]text.VAlign{text.VAlignDefault, text.VAlignMiddle, text.VAlignBottom, text.VAlignMiddle})
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
	t.SetAlign([]text.Align{text.AlignDefault, text.AlignRight, text.AlignDefault, text.AlignDefault, text.AlignCenter})
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
	// Time to begin anew. Too much on the screen for a demo!
	//==========================================================================
	t = table.NewWriter()
	t.AppendHeader(table.Row{"#", "First Name", "Last Name", "Salary"})
	t.AppendRow(table.Row{1, "Arya", "Stark", 3000})
	t.AppendRow(table.Row{20, "Jon", "Snow", 2000, "You know nothing, Jon Snow!"})
	t.AppendRow([]interface{}{300, "Tyrion", "Lannister", 5000})
	t.AppendFooter(table.Row{"", "", "Total", 10000})
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
	funkyStyle := table.Style{
		CharBottomLeft:       "\\",
		CharBottomRight:      "/",
		CharBottomSeparator:  "v",
		CharLeft:             "[",
		CharLeftSeparator:    "{",
		CharMiddleHorizontal: "-",
		CharMiddleSeparator:  "+",
		CharMiddleVertical:   "|",
		CharPaddingLeft:      "<",
		CharPaddingRight:     ">",
		CharRight:            "]",
		CharRightSeparator:   "}",
		CharTopLeft:          "(",
		CharTopRight:         ")",
		CharTopSeparator:     "^",
		FormatFooter:         text.FormatLower,
		FormatHeader:         text.FormatLower,
		FormatRows:           text.FormatUpper,
		Name:                 "funkyStyle",
	}
	t.SetStyle(funkyStyle)
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
	colorBOnW := text.Colors{color.BgWhite, color.FgBlack}
	t.SetColorsHeader([]text.Colors{colorBOnW, colorBOnW, colorBOnW, colorBOnW, colorBOnW})
	t.SetColors([]text.Colors{{color.FgYellow}, {color.FgHiRed}, {color.FgHiRed}, {color.FgGreen}, {color.FgCyan}})
	t.SetColorsFooter([]text.Colors{{}, {}, colorBOnW, colorBOnW})
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
	// lets print the same table line-by-line using "%#v" to see the escape
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
	t.SetColorsHeader([]text.Colors{}) // disable colors on the header
	t.SetColors([]text.Colors{})       // disable colors on the body
	t.SetColorsFooter([]text.Colors{}) // disable colors on the footer
	//==========================================================================

	//==========================================================================
	// I don't like borders!
	//==========================================================================
	t.ShowBorder(false)
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
	t.ShowBorder(true)
	t.ShowSeparators(true)
	t.SetCaption("Table with Borders Everywhere!\n")
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
	// I want CSV.
	//==========================================================================
	for _, line := range strings.Split(t.RenderCSV(), "\n") {
		fmt.Printf("CSV | %s\n", line)
	}
	fmt.Println()
	//CSV | #,First Name,Last Name,Salary,
	//CSV | 1,Arya,Stark,3000,
	//CSV | 20,Jon,Snow,2000,"You know nothing\, Jon Snow!"
	//CSV | 300,Tyrion,Lannister,5000,
	//CSV | ,,Total,10000,
	//==========================================================================

	//==========================================================================
	// Nope. I want a HTML Table.
	//==========================================================================
	for _, line := range strings.Split(t.RenderHTML(), "\n") {
		fmt.Printf("HTML | %s\n", line)
	}
	fmt.Println()
	//HTML | <table class="go-pretty-table">
	//HTML |   <thead>
	//HTML |   <tr>
	//HTML |     <th align="right">#</th>
	//HTML |     <th>First Name</th>
	//HTML |     <th>Last Name</th>
	//HTML |     <th align="right">Salary</th>
	//HTML |     <th>&nbsp;</th>
	//HTML |   </tr>
	//HTML |   </thead>
	//HTML |   <tbody>
	//HTML |   <tr>
	//HTML |     <td align="right">1</td>
	//HTML |     <td>Arya</td>
	//HTML |     <td>Stark</td>
	//HTML |     <td align="right">3000</td>
	//HTML |     <td>&nbsp;</td>
	//HTML |   </tr>
	//HTML |   <tr>
	//HTML |     <td align="right">20</td>
	//HTML |     <td>Jon</td>
	//HTML |     <td>Snow</td>
	//HTML |     <td align="right">2000</td>
	//HTML |     <td>You know nothing, Jon Snow!</td>
	//HTML |   </tr>
	//HTML |   <tr>
	//HTML |     <td align="right">300</td>
	//HTML |     <td>Tyrion</td>
	//HTML |     <td>Lannister</td>
	//HTML |     <td align="right">5000</td>
	//HTML |     <td>&nbsp;</td>
	//HTML |   </tr>
	//HTML |   </tbody>
	//HTML |   <tfoot>
	//HTML |   <tr>
	//HTML |     <td align="right">&nbsp;</td>
	//HTML |     <td>&nbsp;</td>
	//HTML |     <td>Total</td>
	//HTML |     <td align="right">10000</td>
	//HTML |     <td>&nbsp;</td>
	//HTML |   </tr>
	//HTML |   </tfoot>
	//HTML | </table>
	//==========================================================================

	//==========================================================================
	// That's it for today! New features will always find a place in this demo!
	//==========================================================================
}
