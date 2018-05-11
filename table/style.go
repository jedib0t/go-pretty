package table

import (
	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/text"
)

// Style declares how to render the Table and provides very fine-grained control
// on how the Table gets rendered on the Console.
type Style struct {
	Name    string
	Box     StyleBox
	Color   StyleColor
	Format  StyleFormat
	Options StyleOptions
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
		Color:   StyleColorDefault,
		Format:  StyleFormatDefault,
		Options: StyleOptionsDefault,
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
		Color:   StyleColorDefault,
		Format:  StyleFormatDefault,
		Options: StyleOptionsDefault,
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
		Color:   StyleColorDefault,
		Format:  StyleFormatDefault,
		Options: StyleOptionsDefault,
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
		Color:   StyleColorDefault,
		Format:  StyleFormatDefault,
		Options: StyleOptionsDefault,
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
		Color:   StyleColorDefault,
		Format:  StyleFormatDefault,
		Options: StyleOptionsDefault,
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
		Color:   StyleColorDefault,
		Format:  StyleFormatDefault,
		Options: StyleOptionsDefault,
	}
)

// StyleBox defines the characters/strings to use to render the borders and
// separators for the Table.
type StyleBox struct {
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
	StyleBoxDefault = StyleBox{
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
	StyleBoxBold = StyleBox{
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
	StyleBoxDouble = StyleBox{
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
	StyleBoxLight = StyleBox{
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
	StyleBoxRounded = StyleBox{
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
	styleBoxTest = StyleBox{
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

// StyleColor defines the ANSI colors to use for parts of the Table.
type StyleColor struct {
	AutoIndexColumn text.Colors
	FirstColumn     text.Colors
	Footer          text.Colors
	Header          text.Colors
	Row             text.Colors
	RowAlternate    text.Colors
}

var (
	// StyleColorDefault defines sensible ANSI color options - basically NONE.
	StyleColorDefault = StyleColor{
		AutoIndexColumn: nil,
		FirstColumn:     nil,
		Footer:          nil,
		Header:          nil,
		Row:             nil,
		RowAlternate:    nil,
	}

	// StyleColorBright defines ANSI color options to render dark text on bright
	// background.
	StyleColorBright = StyleColor{
		AutoIndexColumn: text.Colors{color.BgHiCyan, color.FgBlack},
		FirstColumn:     nil,
		Footer:          text.Colors{color.BgCyan, color.FgBlack},
		Header:          text.Colors{color.BgHiCyan, color.FgBlack},
		Row:             text.Colors{color.BgWhite, color.FgBlack},
		RowAlternate:    text.Colors{color.BgHiWhite, color.FgBlack},
	}

	// StyleColorDark defines ANSI color options to render bright text on dark
	// background.
	StyleColorDark = StyleColor{
		AutoIndexColumn: text.Colors{color.FgHiCyan, color.BgBlack},
		FirstColumn:     nil,
		Footer:          text.Colors{color.FgCyan, color.BgBlack},
		Header:          text.Colors{color.FgHiCyan, color.BgBlack},
		Row:             text.Colors{color.FgWhite, color.BgBlack},
		RowAlternate:    text.Colors{color.FgHiWhite, color.BgBlack},
	}
)

// StyleFormat defines the text-formatting to perform on parts of the Table.
type StyleFormat struct {
	FirstColumn text.Format
	Footer      text.Format
	Header      text.Format
	Row         text.Format
}

var (
	// StyleFormatDefault defines sensible formatting options.
	StyleFormatDefault = StyleFormat{
		FirstColumn: text.FormatDefault,
		Footer:      text.FormatUpper,
		Header:      text.FormatUpper,
		Row:         text.FormatDefault,
	}
)

// StyleOptions defines the global options that determine how the Table is
// rendered.
type StyleOptions struct {
	DrawBorder      bool
	SeparateColumns bool
	SeparateRows    bool
}

var (
	// StyleOptionsDefault defines sensible global options.
	StyleOptionsDefault = StyleOptions{
		DrawBorder:      true,
		SeparateColumns: true,
		SeparateRows:    false,
	}
)
