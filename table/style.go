package table

import (
	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/text"
)

// Style declares how to render the Table and provides very fine-grained control
// on how the Table gets rendered on the Console.
type Style struct {
	Name    string
	Box     BoxStyle
	Color   ColorOptions
	Format  FormatOptions
	Options Options
}

var (
	// StyleDefault renders a Table like below:
	//  +-----+------------+-----------+--------+-----------------------------+
	//  |   # | FIRST NAME | LAST NAME | SALARY |                             |
	//  +-----+------------+-----------+--------+-----------------------------+
	//  |   1 | Arya       | Stark     |   3000 |                             |
	//  |  20 | Jon        | Snow      |   2000 | You know nothing, Jon Snow! |
	//  | 300 | Tyrion     | Lannister |   5000 |                             |
	//  +-----+------------+-----------+--------+-----------------------------+
	//  |     |            | TOTAL     |  10000 |                             |
	//  +-----+------------+-----------+--------+-----------------------------+
	StyleDefault = Style{
		Name:    "StyleDefault",
		Box:     StyleBoxDefault,
		Color:   ColorOptionsDefault,
		Format:  FormatOptionsDefault,
		Options: OptionsDefault,
	}

	// StyleBold renders a Table like below:
	//  ┏━━━━━┳━━━━━━━━━━━━┳━━━━━━━━━━━┳━━━━━━━━┳━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
	//  ┃   # ┃ FIRST NAME ┃ LAST NAME ┃ SALARY ┃                             ┃
	//  ┣━━━━━╋━━━━━━━━━━━━╋━━━━━━━━━━━╋━━━━━━━━╋━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
	//  ┃   1 ┃ Arya       ┃ Stark     ┃   3000 ┃                             ┃
	//  ┃  20 ┃ Jon        ┃ Snow      ┃   2000 ┃ You know nothing, Jon Snow! ┃
	//  ┃ 300 ┃ Tyrion     ┃ Lannister ┃   5000 ┃                             ┃
	//  ┣━━━━━╋━━━━━━━━━━━━╋━━━━━━━━━━━╋━━━━━━━━╋━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
	//  ┃     ┃            ┃ TOTAL     ┃  10000 ┃                             ┃
	//  ┗━━━━━┻━━━━━━━━━━━━┻━━━━━━━━━━━┻━━━━━━━━┻━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛
	StyleBold = Style{
		Name:    "StyleBold",
		Box:     StyleBoxBold,
		Color:   ColorOptionsDefault,
		Format:  FormatOptionsDefault,
		Options: OptionsDefault,
	}

	// StyleColoredBright renders a Table without any borders or separators, and
	// with every row colored with a Bright background.
	StyleColoredBright = Style{
		Name:    "StyleColoredBright",
		Box:     StyleBoxDefault,
		Color:   ColorOptionsBright,
		Format:  FormatOptionsDefault,
		Options: OptionsNoBordersAndSeparators,
	}

	// StyleColoredDark renders a Table without any borders or separators, and
	// with every row colored with a Dark background.
	StyleColoredDark = Style{
		Name:    "StyleColoredDark",
		Box:     StyleBoxDefault,
		Color:   ColorOptionsDark,
		Format:  FormatOptionsDefault,
		Options: OptionsNoBordersAndSeparators,
	}

	// StyleDouble renders a Table like below:
	//  ╔═════╦════════════╦═══════════╦════════╦═════════════════════════════╗
	//  ║   # ║ FIRST NAME ║ LAST NAME ║ SALARY ║                             ║
	//  ╠═════╬════════════╬═══════════╬════════╬═════════════════════════════╣
	//  ║   1 ║ Arya       ║ Stark     ║   3000 ║                             ║
	//  ║  20 ║ Jon        ║ Snow      ║   2000 ║ You know nothing, Jon Snow! ║
	//  ║ 300 ║ Tyrion     ║ Lannister ║   5000 ║                             ║
	//  ╠═════╬════════════╬═══════════╬════════╬═════════════════════════════╣
	//  ║     ║            ║ TOTAL     ║  10000 ║                             ║
	//  ╚═════╩════════════╩═══════════╩════════╩═════════════════════════════╝
	StyleDouble = Style{
		Name:    "StyleDouble",
		Box:     StyleBoxDouble,
		Color:   ColorOptionsDefault,
		Format:  FormatOptionsDefault,
		Options: OptionsDefault,
	}

	// StyleLight renders a Table like below:
	//  ┌─────┬────────────┬───────────┬────────┬─────────────────────────────┐
	//  │   # │ FIRST NAME │ LAST NAME │ SALARY │                             │
	//  ├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
	//  │   1 │ Arya       │ Stark     │   3000 │                             │
	//  │  20 │ Jon        │ Snow      │   2000 │ You know nothing, Jon Snow! │
	//  │ 300 │ Tyrion     │ Lannister │   5000 │                             │
	//  ├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
	//  │     │            │ TOTAL     │  10000 │                             │
	//  └─────┴────────────┴───────────┴────────┴─────────────────────────────┘
	StyleLight = Style{
		Name:    "StyleLight",
		Box:     StyleBoxLight,
		Color:   ColorOptionsDefault,
		Format:  FormatOptionsDefault,
		Options: OptionsDefault,
	}

	// StyleRounded renders a Table like below:
	//  ╭─────┬────────────┬───────────┬────────┬─────────────────────────────╮
	//  │   # │ FIRST NAME │ LAST NAME │ SALARY │                             │
	//  ├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
	//  │   1 │ Arya       │ Stark     │   3000 │                             │
	//  │  20 │ Jon        │ Snow      │   2000 │ You know nothing, Jon Snow! │
	//  │ 300 │ Tyrion     │ Lannister │   5000 │                             │
	//  ├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
	//  │     │            │ TOTAL     │  10000 │                             │
	//  ╰─────┴────────────┴───────────┴────────┴─────────────────────────────╯
	StyleRounded = Style{
		Name:    "StyleRounded",
		Box:     StyleBoxRounded,
		Color:   ColorOptionsDefault,
		Format:  FormatOptionsDefault,
		Options: OptionsDefault,
	}

	// styleTest renders a Table like below:
	//  (-----^------------^-----------^--------^-----------------------------)
	//  [<  #>|<FIRST NAME>|<LAST NAME>|<SALARY>|<                           >]
	//  {-----+------------+-----------+--------+-----------------------------}
	//  [<  1>|<Arya      >|<Stark    >|<  3000>|<                           >]
	//  [< 20>|<Jon       >|<Snow     >|<  2000>|<You know nothing, Jon Snow!>]
	//  [<300>|<Tyrion    >|<Lannister>|<  5000>|<                           >]
	//  {-----+------------+-----------+--------+-----------------------------}
	//  [<   >|<          >|<TOTAL    >|< 10000>|<                           >]
	//  \-----v------------v-----------v--------v-----------------------------/
	styleTest = Style{
		Name:    "styleTest",
		Box:     styleBoxTest,
		Color:   ColorOptionsDefault,
		Format:  FormatOptionsDefault,
		Options: OptionsDefault,
	}
)

// BoxStyle defines the characters/strings to use to render the borders and
// separators for the Table.
type BoxStyle struct {
	BottomLeft       string
	BottomRight      string
	BottomSeparator  string
	Left             string
	LeftSeparator    string
	MiddleHorizontal string
	MiddleSeparator  string
	MiddleVertical   string
	PaddingLeft      string
	PaddingRight     string
	Right            string
	RightSeparator   string
	TopLeft          string
	TopRight         string
	TopSeparator     string
	UnfinishedRow    string
}

var (
	// StyleBoxDefault defines a Boxed-Table like below:
	//  +-----+------------+-----------+--------+-----------------------------+
	//  |   # | FIRST NAME | LAST NAME | SALARY |                             |
	//  +-----+------------+-----------+--------+-----------------------------+
	//  |   1 | Arya       | Stark     |   3000 |                             |
	//  |  20 | Jon        | Snow      |   2000 | You know nothing, Jon Snow! |
	//  | 300 | Tyrion     | Lannister |   5000 |                             |
	//  +-----+------------+-----------+--------+-----------------------------+
	//  |     |            | TOTAL     |  10000 |                             |
	//  +-----+------------+-----------+--------+-----------------------------+
	StyleBoxDefault = BoxStyle{
		BottomLeft:       "+",
		BottomRight:      "+",
		BottomSeparator:  "+",
		Left:             "|",
		LeftSeparator:    "+",
		MiddleHorizontal: "-",
		MiddleSeparator:  "+",
		MiddleVertical:   "|",
		PaddingLeft:      " ",
		PaddingRight:     " ",
		Right:            "|",
		RightSeparator:   "+",
		TopLeft:          "+",
		TopRight:         "+",
		TopSeparator:     "+",
		UnfinishedRow:    " ~",
	}

	// StyleBoxBold defines a Boxed-Table like below:
	//  ┏━━━━━┳━━━━━━━━━━━━┳━━━━━━━━━━━┳━━━━━━━━┳━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
	//  ┃   # ┃ FIRST NAME ┃ LAST NAME ┃ SALARY ┃                             ┃
	//  ┣━━━━━╋━━━━━━━━━━━━╋━━━━━━━━━━━╋━━━━━━━━╋━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
	//  ┃   1 ┃ Arya       ┃ Stark     ┃   3000 ┃                             ┃
	//  ┃  20 ┃ Jon        ┃ Snow      ┃   2000 ┃ You know nothing, Jon Snow! ┃
	//  ┃ 300 ┃ Tyrion     ┃ Lannister ┃   5000 ┃                             ┃
	//  ┣━━━━━╋━━━━━━━━━━━━╋━━━━━━━━━━━╋━━━━━━━━╋━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
	//  ┃     ┃            ┃ TOTAL     ┃  10000 ┃                             ┃
	//  ┗━━━━━┻━━━━━━━━━━━━┻━━━━━━━━━━━┻━━━━━━━━┻━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛
	StyleBoxBold = BoxStyle{
		BottomLeft:       text.BoxBottomLeftBold,
		BottomRight:      text.BoxBottomRightBold,
		BottomSeparator:  text.BoxBottomSeparatorBold,
		Left:             text.BoxLeftBold,
		LeftSeparator:    text.BoxLeftSeparatorBold,
		MiddleHorizontal: text.BoxHorizontalBold,
		MiddleSeparator:  text.BoxSeparatorBold,
		MiddleVertical:   text.BoxVerticalBold,
		PaddingLeft:      " ",
		PaddingRight:     " ",
		Right:            text.BoxRightBold,
		RightSeparator:   text.BoxRightSeparatorBold,
		TopLeft:          text.BoxTopLeftBold,
		TopRight:         text.BoxTopRightBold,
		TopSeparator:     text.BoxTopSeparatorBold,
		UnfinishedRow:    " " + text.BoxUnfinishedLine,
	}

	// StyleBoxDouble defines a Boxed-Table like below:
	//  ╔═════╦════════════╦═══════════╦════════╦═════════════════════════════╗
	//  ║   # ║ FIRST NAME ║ LAST NAME ║ SALARY ║                             ║
	//  ╠═════╬════════════╬═══════════╬════════╬═════════════════════════════╣
	//  ║   1 ║ Arya       ║ Stark     ║   3000 ║                             ║
	//  ║  20 ║ Jon        ║ Snow      ║   2000 ║ You know nothing, Jon Snow! ║
	//  ║ 300 ║ Tyrion     ║ Lannister ║   5000 ║                             ║
	//  ╠═════╬════════════╬═══════════╬════════╬═════════════════════════════╣
	//  ║     ║            ║ TOTAL     ║  10000 ║                             ║
	//  ╚═════╩════════════╩═══════════╩════════╩═════════════════════════════╝
	StyleBoxDouble = BoxStyle{
		BottomLeft:       text.BoxBottomLeftDouble,
		BottomRight:      text.BoxBottomRightDouble,
		BottomSeparator:  text.BoxBottomSeparatorDouble,
		Left:             text.BoxLeftDouble,
		LeftSeparator:    text.BoxLeftSeparatorDouble,
		MiddleHorizontal: text.BoxHorizontalDouble,
		MiddleSeparator:  text.BoxSeparatorDouble,
		MiddleVertical:   text.BoxVerticalDouble,
		PaddingLeft:      " ",
		PaddingRight:     " ",
		Right:            text.BoxRightDouble,
		RightSeparator:   text.BoxRightSeparatorDouble,
		TopLeft:          text.BoxTopLeftDouble,
		TopRight:         text.BoxTopRightDouble,
		TopSeparator:     text.BoxTopSeparatorDouble,
		UnfinishedRow:    " " + text.BoxUnfinishedLine,
	}

	// StyleBoxLight defines a Boxed-Table like below:
	//  ┌─────┬────────────┬───────────┬────────┬─────────────────────────────┐
	//  │   # │ FIRST NAME │ LAST NAME │ SALARY │                             │
	//  ├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
	//  │   1 │ Arya       │ Stark     │   3000 │                             │
	//  │  20 │ Jon        │ Snow      │   2000 │ You know nothing, Jon Snow! │
	//  │ 300 │ Tyrion     │ Lannister │   5000 │                             │
	//  ├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
	//  │     │            │ TOTAL     │  10000 │                             │
	//  └─────┴────────────┴───────────┴────────┴─────────────────────────────┘
	StyleBoxLight = BoxStyle{
		BottomLeft:       text.BoxBottomLeft,
		BottomRight:      text.BoxBottomRight,
		BottomSeparator:  text.BoxBottomSeparator,
		Left:             text.BoxLeft,
		LeftSeparator:    text.BoxLeftSeparator,
		MiddleHorizontal: text.BoxHorizontal,
		MiddleSeparator:  text.BoxSeparator,
		MiddleVertical:   text.BoxVertical,
		PaddingLeft:      " ",
		PaddingRight:     " ",
		Right:            text.BoxRight,
		RightSeparator:   text.BoxRightSeparator,
		TopLeft:          text.BoxTopLeft,
		TopRight:         text.BoxTopRight,
		TopSeparator:     text.BoxTopSeparator,
		UnfinishedRow:    " " + text.BoxUnfinishedLine,
	}

	// StyleBoxRounded defines a Boxed-Table like below:
	//  ╭─────┬────────────┬───────────┬────────┬─────────────────────────────╮
	//  │   # │ FIRST NAME │ LAST NAME │ SALARY │                             │
	//  ├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
	//  │   1 │ Arya       │ Stark     │   3000 │                             │
	//  │  20 │ Jon        │ Snow      │   2000 │ You know nothing, Jon Snow! │
	//  │ 300 │ Tyrion     │ Lannister │   5000 │                             │
	//  ├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
	//  │     │            │ TOTAL     │  10000 │                             │
	//  ╰─────┴────────────┴───────────┴────────┴─────────────────────────────╯
	StyleBoxRounded = BoxStyle{
		BottomLeft:       text.BoxBottomLeftRounded,
		BottomRight:      text.BoxBottomRightRounded,
		BottomSeparator:  text.BoxBottomSeparator,
		Left:             text.BoxLeft,
		LeftSeparator:    text.BoxLeftSeparator,
		MiddleHorizontal: text.BoxHorizontal,
		MiddleSeparator:  text.BoxSeparator,
		MiddleVertical:   text.BoxVertical,
		PaddingLeft:      " ",
		PaddingRight:     " ",
		Right:            text.BoxRight,
		RightSeparator:   text.BoxRightSeparator,
		TopLeft:          text.BoxTopLeftRounded,
		TopRight:         text.BoxTopRightRounded,
		TopSeparator:     text.BoxTopSeparator,
		UnfinishedRow:    " " + text.BoxUnfinishedLine,
	}

	// styleBoxTest defines a Boxed-Table like below:
	//  (-----^------------^-----------^--------^-----------------------------)
	//  [<  #>|<FIRST NAME>|<LAST NAME>|<SALARY>|<                           >]
	//  {-----+------------+-----------+--------+-----------------------------}
	//  [<  1>|<Arya      >|<Stark    >|<  3000>|<                           >]
	//  [< 20>|<Jon       >|<Snow     >|<  2000>|<You know nothing, Jon Snow!>]
	//  [<300>|<Tyrion    >|<Lannister>|<  5000>|<                           >]
	//  {-----+------------+-----------+--------+-----------------------------}
	//  [<   >|<          >|<TOTAL    >|< 10000>|<                           >]
	//  \-----v------------v-----------v--------v-----------------------------/
	styleBoxTest = BoxStyle{
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
	}
)

// ColorOptions defines the ANSI colors to use for parts of the Table.
type ColorOptions struct {
	AutoIndexColumn *color.Color
	FirstColumn     *color.Color
	Footer          *color.Color
	Header          *color.Color
	Row             *color.Color
	RowAlternate    *color.Color
}

var (
	// ColorOptionsDefault defines sensible ANSI color options - basically NONE.
	ColorOptionsDefault = ColorOptions{
		AutoIndexColumn: nil,
		FirstColumn:     nil,
		Footer:          nil,
		Header:          nil,
		Row:             nil,
		RowAlternate:    nil,
	}

	// ColorOptionsBright defines ANSI color options to render dark text on
	// bright background.
	ColorOptionsBright = ColorOptions{
		AutoIndexColumn: text.Colors{color.BgHiCyan, color.FgBlack}.GetColorizer(),
		FirstColumn:     nil,
		Footer:          text.Colors{color.BgCyan, color.FgBlack}.GetColorizer(),
		Header:          text.Colors{color.BgHiCyan, color.FgBlack}.GetColorizer(),
		Row:             text.Colors{color.BgHiWhite, color.FgBlack}.GetColorizer(),
		RowAlternate:    text.Colors{color.BgWhite, color.FgBlack}.GetColorizer(),
	}

	// ColorOptionsDark defines ANSI color options to render bright text on dark
	// background.
	ColorOptionsDark = ColorOptions{
		AutoIndexColumn: text.Colors{color.FgHiCyan, color.BgBlack}.GetColorizer(),
		FirstColumn:     nil,
		Footer:          text.Colors{color.FgCyan, color.BgBlack}.GetColorizer(),
		Header:          text.Colors{color.FgHiCyan, color.BgBlack}.GetColorizer(),
		Row:             text.Colors{color.FgHiWhite, color.BgBlack}.GetColorizer(),
		RowAlternate:    text.Colors{color.FgWhite, color.BgBlack}.GetColorizer(),
	}
)

// FormatOptions defines the text-formatting to perform on parts of the Table.
type FormatOptions struct {
	Footer text.Format
	Header text.Format
	Row    text.Format
}

var (
	// FormatOptionsDefault defines sensible formatting options.
	FormatOptionsDefault = FormatOptions{
		Footer: text.FormatUpper,
		Header: text.FormatUpper,
		Row:    text.FormatDefault,
	}
)

// Options defines the global options that determine how the Table is
// rendered.
type Options struct {
	// DrawBorder enables or disables drawing the border around the Table.
	// Example of a table where it is disabled:
	//     # │ FIRST NAME │ LAST NAME │ SALARY │
	//  ─────┼────────────┼───────────┼────────┼─────────────────────────────
	//     1 │ Arya       │ Stark     │   3000 │
	//    20 │ Jon        │ Snow      │   2000 │ You know nothing, Jon Snow!
	//   300 │ Tyrion     │ Lannister │   5000 │
	//  ─────┼────────────┼───────────┼────────┼─────────────────────────────
	//       │            │ TOTAL     │  10000 │
	DrawBorder bool

	// SeparateColumns enables or disable drawing border between columns.
	// Example of a table where it is disabled:
	//  ┌─────────────────────────────────────────────────────────────────┐
	//  │   #  FIRST NAME  LAST NAME  SALARY                              │
	//  ├─────────────────────────────────────────────────────────────────┤
	//  │   1  Arya        Stark        3000                              │
	//  │  20  Jon         Snow         2000  You know nothing, Jon Snow! │
	//  │ 300  Tyrion      Lannister    5000                              │
	//  │                  TOTAL       10000                              │
	//  └─────────────────────────────────────────────────────────────────┘
	SeparateColumns bool

	// SeparateFooter enables or disable drawing border between the footer and
	// the rows. Example of a table where it is disabled:
	//  ┌─────┬────────────┬───────────┬────────┬─────────────────────────────┐
	//  │   # │ FIRST NAME │ LAST NAME │ SALARY │                             │
	//  ├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
	//  │   1 │ Arya       │ Stark     │   3000 │                             │
	//  │  20 │ Jon        │ Snow      │   2000 │ You know nothing, Jon Snow! │
	//  │ 300 │ Tyrion     │ Lannister │   5000 │                             │
	//  │     │            │ TOTAL     │  10000 │                             │
	//  └─────┴────────────┴───────────┴────────┴─────────────────────────────┘
	SeparateFooter bool

	// SeparateHeader enables or disable drawing border between the header and
	// the rows. Example of a table where it is disabled:
	//  ┌─────┬────────────┬───────────┬────────┬─────────────────────────────┐
	//  │   # │ FIRST NAME │ LAST NAME │ SALARY │                             │
	//  │   1 │ Arya       │ Stark     │   3000 │                             │
	//  │  20 │ Jon        │ Snow      │   2000 │ You know nothing, Jon Snow! │
	//  │ 300 │ Tyrion     │ Lannister │   5000 │                             │
	//  ├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
	//  │     │            │ TOTAL     │  10000 │                             │
	//  └─────┴────────────┴───────────┴────────┴─────────────────────────────┘
	SeparateHeader bool

	// SeparateRows enables or disables drawing separators between each row.
	// Example of a table where it is enabled:
	//  ┌─────┬────────────┬───────────┬────────┬─────────────────────────────┐
	//  │   # │ FIRST NAME │ LAST NAME │ SALARY │                             │
	//  ├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
	//  │   1 │ Arya       │ Stark     │   3000 │                             │
	//  ├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
	//  │  20 │ Jon        │ Snow      │   2000 │ You know nothing, Jon Snow! │
	//  ├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
	//  │ 300 │ Tyrion     │ Lannister │   5000 │                             │
	//  ├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
	//  │     │            │ TOTAL     │  10000 │                             │
	//  └─────┴────────────┴───────────┴────────┴─────────────────────────────┘
	SeparateRows bool
}

var (
	// OptionsDefault defines sensible global options.
	OptionsDefault = Options{
		DrawBorder:      true,
		SeparateColumns: true,
		SeparateFooter:  true,
		SeparateHeader:  true,
		SeparateRows:    false,
	}

	// OptionsNoBorders sets up a table without any borders.
	OptionsNoBorders = Options{
		DrawBorder:      false,
		SeparateColumns: true,
		SeparateFooter:  true,
		SeparateHeader:  true,
		SeparateRows:    false,
	}

	// OptionsNoBordersAndSeparators sets up a table without any borders or
	// separators.
	OptionsNoBordersAndSeparators = Options{
		DrawBorder:      false,
		SeparateColumns: false,
		SeparateFooter:  false,
		SeparateHeader:  false,
		SeparateRows:    false,
	}
)
