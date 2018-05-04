package table

import "github.com/jedib0t/go-pretty/text"

// Style declares how to render the Table.
type Style struct {
	BoxBottomLeft       string
	BoxBottomRight      string
	BoxBottomSeparator  string
	BoxLeft             string
	BoxLeftSeparator    string
	BoxMiddleHorizontal string
	BoxMiddleSeparator  string
	BoxMiddleVertical   string
	BoxPaddingLeft      string
	BoxPaddingRight     string
	BoxRight            string
	BoxRightSeparator   string
	BoxTopLeft          string
	BoxTopRight         string
	BoxTopSeparator     string
	BoxUnfinishedRow    string
	FormatHeader        text.Format
	FormatFooter        text.Format
	FormatRows          text.Format
	Name                string
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
		BoxBottomLeft:       "+",
		BoxBottomRight:      "+",
		BoxBottomSeparator:  "+",
		BoxLeft:             "|",
		BoxLeftSeparator:    "+",
		BoxMiddleHorizontal: "-",
		BoxMiddleSeparator:  "+",
		BoxMiddleVertical:   "|",
		BoxPaddingLeft:      " ",
		BoxPaddingRight:     " ",
		BoxRight:            "|",
		BoxRightSeparator:   "+",
		BoxTopLeft:          "+",
		BoxTopRight:         "+",
		BoxTopSeparator:     "+",
		BoxUnfinishedRow:    " ~",
		FormatFooter:        text.FormatUpper,
		FormatHeader:        text.FormatUpper,
		FormatRows:          text.FormatDefault,
		Name:                "StyleDefault",
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
		BoxBottomLeft:       text.BoxBottomLeftBold,
		BoxBottomRight:      text.BoxBottomRightBold,
		BoxBottomSeparator:  text.BoxBottomSeparatorBold,
		BoxLeft:             text.BoxLeftBold,
		BoxLeftSeparator:    text.BoxLeftSeparatorBold,
		BoxMiddleHorizontal: text.BoxHorizontalBold,
		BoxMiddleSeparator:  text.BoxSeparatorBold,
		BoxMiddleVertical:   text.BoxVerticalBold,
		BoxPaddingLeft:      " ",
		BoxPaddingRight:     " ",
		BoxRight:            text.BoxRightBold,
		BoxRightSeparator:   text.BoxRightSeparatorBold,
		BoxTopLeft:          text.BoxTopLeftBold,
		BoxTopRight:         text.BoxTopRightBold,
		BoxTopSeparator:     text.BoxTopSeparatorBold,
		BoxUnfinishedRow:    " " + text.BoxUnfinishedLine,
		FormatFooter:        text.FormatUpper,
		FormatHeader:        text.FormatUpper,
		FormatRows:          text.FormatDefault,
		Name:                "StyleBold",
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
		BoxBottomLeft:       text.BoxBottomLeftDouble,
		BoxBottomRight:      text.BoxBottomRightDouble,
		BoxBottomSeparator:  text.BoxBottomSeparatorDouble,
		BoxLeft:             text.BoxLeftDouble,
		BoxLeftSeparator:    text.BoxLeftSeparatorDouble,
		BoxMiddleHorizontal: text.BoxHorizontalDouble,
		BoxMiddleSeparator:  text.BoxSeparatorDouble,
		BoxMiddleVertical:   text.BoxVerticalDouble,
		BoxPaddingLeft:      " ",
		BoxPaddingRight:     " ",
		BoxRight:            text.BoxRightDouble,
		BoxRightSeparator:   text.BoxRightSeparatorDouble,
		BoxTopLeft:          text.BoxTopLeftDouble,
		BoxTopRight:         text.BoxTopRightDouble,
		BoxTopSeparator:     text.BoxTopSeparatorDouble,
		BoxUnfinishedRow:    " " + text.BoxUnfinishedLine,
		FormatFooter:        text.FormatUpper,
		FormatHeader:        text.FormatUpper,
		FormatRows:          text.FormatDefault,
		Name:                "StyleDouble",
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
		BoxBottomLeft:       text.BoxBottomLeft,
		BoxBottomRight:      text.BoxBottomRight,
		BoxBottomSeparator:  text.BoxBottomSeparator,
		BoxLeft:             text.BoxLeft,
		BoxLeftSeparator:    text.BoxLeftSeparator,
		BoxMiddleHorizontal: text.BoxHorizontal,
		BoxMiddleSeparator:  text.BoxSeparator,
		BoxMiddleVertical:   text.BoxVertical,
		BoxPaddingLeft:      " ",
		BoxPaddingRight:     " ",
		BoxRight:            text.BoxRight,
		BoxRightSeparator:   text.BoxRightSeparator,
		BoxTopLeft:          text.BoxTopLeft,
		BoxTopRight:         text.BoxTopRight,
		BoxTopSeparator:     text.BoxTopSeparator,
		BoxUnfinishedRow:    " " + text.BoxUnfinishedLine,
		FormatFooter:        text.FormatUpper,
		FormatHeader:        text.FormatUpper,
		FormatRows:          text.FormatDefault,
		Name:                "StyleLight",
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
		BoxBottomLeft:       text.BoxBottomLeftRounded,
		BoxBottomRight:      text.BoxBottomRightRounded,
		BoxBottomSeparator:  text.BoxBottomSeparator,
		BoxLeft:             text.BoxLeft,
		BoxLeftSeparator:    text.BoxLeftSeparator,
		BoxMiddleHorizontal: text.BoxHorizontal,
		BoxMiddleSeparator:  text.BoxSeparator,
		BoxMiddleVertical:   text.BoxVertical,
		BoxPaddingLeft:      " ",
		BoxPaddingRight:     " ",
		BoxRight:            text.BoxRight,
		BoxRightSeparator:   text.BoxRightSeparator,
		BoxTopLeft:          text.BoxTopLeftRounded,
		BoxTopRight:         text.BoxTopRightRounded,
		BoxTopSeparator:     text.BoxTopSeparator,
		BoxUnfinishedRow:    " " + text.BoxUnfinishedLine,
		FormatFooter:        text.FormatUpper,
		FormatHeader:        text.FormatUpper,
		FormatRows:          text.FormatDefault,
		Name:                "StyleRounded",
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
		BoxBottomLeft:       "\\",
		BoxBottomRight:      "/",
		BoxBottomSeparator:  "v",
		BoxLeft:             "[",
		BoxLeftSeparator:    "{",
		BoxMiddleHorizontal: "-",
		BoxMiddleSeparator:  "+",
		BoxMiddleVertical:   "|",
		BoxPaddingLeft:      "<",
		BoxPaddingRight:     ">",
		BoxRight:            "]",
		BoxRightSeparator:   "}",
		BoxTopLeft:          "(",
		BoxTopRight:         ")",
		BoxTopSeparator:     "^",
		BoxUnfinishedRow:    " ~~~",
		FormatFooter:        text.FormatUpper,
		FormatHeader:        text.FormatUpper,
		FormatRows:          text.FormatDefault,
		Name:                "styleTest",
	}
)
