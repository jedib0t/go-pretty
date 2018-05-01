package table

import "github.com/jedib0t/go-pretty/text"

// Style declares how to render the Table.
type Style struct {
	CharBottomLeft       string
	CharBottomRight      string
	CharBottomSeparator  string
	CharLeft             string
	CharLeftSeparator    string
	CharMiddleHorizontal string
	CharMiddleSeparator  string
	CharMiddleVertical   string
	CharPaddingLeft      string
	CharPaddingRight     string
	CharRight            string
	CharRightSeparator   string
	CharTopLeft          string
	CharTopRight         string
	CharTopSeparator     string
	FormatHeader         text.Format
	FormatFooter         text.Format
	FormatRows           text.Format
	Name                 string
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
		CharBottomLeft:       "+",
		CharBottomRight:      "+",
		CharBottomSeparator:  "+",
		CharLeft:             "|",
		CharLeftSeparator:    "+",
		CharMiddleHorizontal: "-",
		CharMiddleSeparator:  "+",
		CharMiddleVertical:   "|",
		CharPaddingLeft:      " ",
		CharPaddingRight:     " ",
		CharRight:            "|",
		CharRightSeparator:   "+",
		CharTopLeft:          "+",
		CharTopRight:         "+",
		CharTopSeparator:     "+",
		FormatFooter:         text.FormatUpper,
		FormatHeader:         text.FormatUpper,
		FormatRows:           text.FormatDefault,
		Name:                 "StyleDefault",
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
		CharBottomLeft:       text.BoxBottomLeftBold,
		CharBottomRight:      text.BoxBottomRightBold,
		CharBottomSeparator:  text.BoxBottomSeparatorBold,
		CharLeft:             text.BoxLeftBold,
		CharLeftSeparator:    text.BoxLeftSeparatorBold,
		CharMiddleHorizontal: text.BoxHorizontalBold,
		CharMiddleSeparator:  text.BoxSeparatorBold,
		CharMiddleVertical:   text.BoxVerticalBold,
		CharPaddingLeft:      " ",
		CharPaddingRight:     " ",
		CharRight:            text.BoxRightBold,
		CharRightSeparator:   text.BoxRightSeparatorBold,
		CharTopLeft:          text.BoxTopLeftBold,
		CharTopRight:         text.BoxTopRightBold,
		CharTopSeparator:     text.BoxTopSeparatorBold,
		FormatFooter:         text.FormatUpper,
		FormatHeader:         text.FormatUpper,
		FormatRows:           text.FormatDefault,
		Name:                 "StyleBold",
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
		CharBottomLeft:       text.BoxBottomLeftDouble,
		CharBottomRight:      text.BoxBottomRightDouble,
		CharBottomSeparator:  text.BoxBottomSeparatorDouble,
		CharLeft:             text.BoxLeftDouble,
		CharLeftSeparator:    text.BoxLeftSeparatorDouble,
		CharMiddleHorizontal: text.BoxHorizontalDouble,
		CharMiddleSeparator:  text.BoxSeparatorDouble,
		CharMiddleVertical:   text.BoxVerticalDouble,
		CharPaddingLeft:      " ",
		CharPaddingRight:     " ",
		CharRight:            text.BoxRightDouble,
		CharRightSeparator:   text.BoxRightSeparatorDouble,
		CharTopLeft:          text.BoxTopLeftDouble,
		CharTopRight:         text.BoxTopRightDouble,
		CharTopSeparator:     text.BoxTopSeparatorDouble,
		FormatFooter:         text.FormatUpper,
		FormatHeader:         text.FormatUpper,
		FormatRows:           text.FormatDefault,
		Name:                 "StyleDouble",
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
		CharBottomLeft:       text.BoxBottomLeft,
		CharBottomRight:      text.BoxBottomRight,
		CharBottomSeparator:  text.BoxBottomSeparator,
		CharLeft:             text.BoxLeft,
		CharLeftSeparator:    text.BoxLeftSeparator,
		CharMiddleHorizontal: text.BoxHorizontal,
		CharMiddleSeparator:  text.BoxSeparator,
		CharMiddleVertical:   text.BoxVertical,
		CharPaddingLeft:      " ",
		CharPaddingRight:     " ",
		CharRight:            text.BoxRight,
		CharRightSeparator:   text.BoxRightSeparator,
		CharTopLeft:          text.BoxTopLeft,
		CharTopRight:         text.BoxTopRight,
		CharTopSeparator:     text.BoxTopSeparator,
		FormatFooter:         text.FormatUpper,
		FormatHeader:         text.FormatUpper,
		FormatRows:           text.FormatDefault,
		Name:                 "StyleLight",
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
		CharBottomLeft:       text.BoxBottomLeftRounded,
		CharBottomRight:      text.BoxBottomRightRounded,
		CharBottomSeparator:  text.BoxBottomSeparator,
		CharLeft:             text.BoxLeft,
		CharLeftSeparator:    text.BoxLeftSeparator,
		CharMiddleHorizontal: text.BoxHorizontal,
		CharMiddleSeparator:  text.BoxSeparator,
		CharMiddleVertical:   text.BoxVertical,
		CharPaddingLeft:      " ",
		CharPaddingRight:     " ",
		CharRight:            text.BoxRight,
		CharRightSeparator:   text.BoxRightSeparator,
		CharTopLeft:          text.BoxTopLeftRounded,
		CharTopRight:         text.BoxTopRightRounded,
		CharTopSeparator:     text.BoxTopSeparator,
		FormatFooter:         text.FormatUpper,
		FormatHeader:         text.FormatUpper,
		FormatRows:           text.FormatDefault,
		Name:                 "StyleRounded",
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
		FormatFooter:         text.FormatUpper,
		FormatHeader:         text.FormatUpper,
		FormatRows:           text.FormatDefault,
		Name:                 "styleTest",
	}
)
