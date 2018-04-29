package gopretty

var (
	// TableStyleDefault renders a Table like below:
	//  +-----+------------+-----------+--------+-----------------------------+
	//  |   # | FIRST NAME | LAST NAME | SALARY |                             |
	//  +-----+------------+-----------+--------+-----------------------------+
	//  |   1 | Arya       | Stark     |   3000 |                             |
	//  |  20 | Jon        | Snow      |   2000 | You know nothing, Jon Snow! |
	//  | 300 | Tyrion     | Lannister |   5000 |                             |
	//  +-----+------------+-----------+--------+-----------------------------+
	//  |     |            | TOTAL     |  10000 |                             |
	//  +-----+------------+-----------+--------+-----------------------------+
	TableStyleDefault = TableStyle{
		CaseFooter:           TextCaseUpper,
		CaseHeader:           TextCaseUpper,
		CaseRows:             TextCaseDefault,
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
		Name:                 "TableStyleDefault",
	}

	// TableStyleBold renders a Table like below:
	//  ┏━━━━━┳━━━━━━━━━━━━┳━━━━━━━━━━━┳━━━━━━━━┳━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
	//  ┃   # ┃ FIRST NAME ┃ LAST NAME ┃ SALARY ┃                             ┃
	//  ┣━━━━━╋━━━━━━━━━━━━╋━━━━━━━━━━━╋━━━━━━━━╋━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
	//  ┃   1 ┃ Arya       ┃ Stark     ┃   3000 ┃                             ┃
	//  ┃  20 ┃ Jon        ┃ Snow      ┃   2000 ┃ You know nothing, Jon Snow! ┃
	//  ┃ 300 ┃ Tyrion     ┃ Lannister ┃   5000 ┃                             ┃
	//  ┣━━━━━╋━━━━━━━━━━━━╋━━━━━━━━━━━╋━━━━━━━━╋━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
	//  ┃     ┃            ┃ TOTAL     ┃  10000 ┃                             ┃
	//  ┗━━━━━┻━━━━━━━━━━━━┻━━━━━━━━━━━┻━━━━━━━━┻━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛
	TableStyleBold = TableStyle{
		CaseFooter:           TextCaseUpper,
		CaseHeader:           TextCaseUpper,
		CaseRows:             TextCaseDefault,
		CharBottomLeft:       BorderBottomLeftBold,
		CharBottomRight:      BorderBottomRightBold,
		CharBottomSeparator:  BorderBottomSeparatorBold,
		CharLeft:             BorderLeftBold,
		CharLeftSeparator:    BorderLeftSeparatorBold,
		CharMiddleHorizontal: BorderHorizontalBold,
		CharMiddleSeparator:  BorderSeparatorBold,
		CharMiddleVertical:   BorderVerticalBold,
		CharPaddingLeft:      " ",
		CharPaddingRight:     " ",
		CharRight:            BorderRightBold,
		CharRightSeparator:   BorderRightSeparatorBold,
		CharTopLeft:          BorderTopLeftBold,
		CharTopRight:         BorderTopRightBold,
		CharTopSeparator:     BorderTopSeparatorBold,
		Name:                 "TableStyleBold",
	}

	// TableStyleDouble renders a Table like below:
	//  ╔═════╦════════════╦═══════════╦════════╦═════════════════════════════╗
	//  ║   # ║ FIRST NAME ║ LAST NAME ║ SALARY ║                             ║
	//  ╠═════╬════════════╬═══════════╬════════╬═════════════════════════════╣
	//  ║   1 ║ Arya       ║ Stark     ║   3000 ║                             ║
	//  ║  20 ║ Jon        ║ Snow      ║   2000 ║ You know nothing, Jon Snow! ║
	//  ║ 300 ║ Tyrion     ║ Lannister ║   5000 ║                             ║
	//  ╠═════╬════════════╬═══════════╬════════╬═════════════════════════════╣
	//  ║     ║            ║ TOTAL     ║  10000 ║                             ║
	//  ╚═════╩════════════╩═══════════╩════════╩═════════════════════════════╝
	TableStyleDouble = TableStyle{
		CaseFooter:           TextCaseUpper,
		CaseHeader:           TextCaseUpper,
		CaseRows:             TextCaseDefault,
		CharBottomLeft:       BorderBottomLeftDouble,
		CharBottomRight:      BorderBottomRightDouble,
		CharBottomSeparator:  BorderBottomSeparatorDouble,
		CharLeft:             BorderLeftDouble,
		CharLeftSeparator:    BorderLeftSeparatorDouble,
		CharMiddleHorizontal: BorderHorizontalDouble,
		CharMiddleSeparator:  BorderSeparatorDouble,
		CharMiddleVertical:   BorderVerticalDouble,
		CharPaddingLeft:      " ",
		CharPaddingRight:     " ",
		CharRight:            BorderRightDouble,
		CharRightSeparator:   BorderRightSeparatorDouble,
		CharTopLeft:          BorderTopLeftDouble,
		CharTopRight:         BorderTopRightDouble,
		CharTopSeparator:     BorderTopSeparatorDouble,
		Name:                 "TableStyleDouble",
	}

	// TableStyleLight renders a Table like below:
	//  ┌─────┬────────────┬───────────┬────────┬─────────────────────────────┐
	//  │   # │ FIRST NAME │ LAST NAME │ SALARY │                             │
	//  ├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
	//  │   1 │ Arya       │ Stark     │   3000 │                             │
	//  │  20 │ Jon        │ Snow      │   2000 │ You know nothing, Jon Snow! │
	//  │ 300 │ Tyrion     │ Lannister │   5000 │                             │
	//  ├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
	//  │     │            │ TOTAL     │  10000 │                             │
	//  └─────┴────────────┴───────────┴────────┴─────────────────────────────┘
	TableStyleLight = TableStyle{
		CaseFooter:           TextCaseUpper,
		CaseHeader:           TextCaseUpper,
		CaseRows:             TextCaseDefault,
		CharBottomLeft:       BorderBottomLeft,
		CharBottomRight:      BorderBottomRight,
		CharBottomSeparator:  BorderBottomSeparator,
		CharLeft:             BorderLeft,
		CharLeftSeparator:    BorderLeftSeparator,
		CharMiddleHorizontal: BorderHorizontal,
		CharMiddleSeparator:  BorderSeparator,
		CharMiddleVertical:   BorderVertical,
		CharPaddingLeft:      " ",
		CharPaddingRight:     " ",
		CharRight:            BorderRight,
		CharRightSeparator:   BorderRightSeparator,
		CharTopLeft:          BorderTopLeft,
		CharTopRight:         BorderTopRight,
		CharTopSeparator:     BorderTopSeparator,
		Name:                 "TableStyleLight",
	}

	// TableStyleRounded renders a Table like below:
	//  ╭─────┬────────────┬───────────┬────────┬─────────────────────────────╮
	//  │   # │ FIRST NAME │ LAST NAME │ SALARY │                             │
	//  ├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
	//  │   1 │ Arya       │ Stark     │   3000 │                             │
	//  │  20 │ Jon        │ Snow      │   2000 │ You know nothing, Jon Snow! │
	//  │ 300 │ Tyrion     │ Lannister │   5000 │                             │
	//  ├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
	//  │     │            │ TOTAL     │  10000 │                             │
	//  ╰─────┴────────────┴───────────┴────────┴─────────────────────────────╯
	TableStyleRounded = TableStyle{
		CaseFooter:           TextCaseUpper,
		CaseHeader:           TextCaseUpper,
		CaseRows:             TextCaseDefault,
		CharBottomLeft:       BorderBottomLeftRounded,
		CharBottomRight:      BorderBottomRightRounded,
		CharBottomSeparator:  BorderBottomSeparator,
		CharLeft:             BorderLeft,
		CharLeftSeparator:    BorderLeftSeparator,
		CharMiddleHorizontal: BorderHorizontal,
		CharMiddleSeparator:  BorderSeparator,
		CharMiddleVertical:   BorderVertical,
		CharPaddingLeft:      " ",
		CharPaddingRight:     " ",
		CharRight:            BorderRight,
		CharRightSeparator:   BorderRightSeparator,
		CharTopLeft:          BorderTopLeftRounded,
		CharTopRight:         BorderTopRightRounded,
		CharTopSeparator:     BorderTopSeparator,
		Name:                 "TableStyleRounded",
	}

	// tableStyleTest renders a Table like below:
	//  (-----^------------^-----------^--------^-----------------------------)
	//  [<  #>|<FIRST NAME>|<LAST NAME>|<SALARY>|<                           >]
	//  {-----+------------+-----------+--------+-----------------------------}
	//  [<  1>|<Arya      >|<Stark    >|<  3000>|<                           >]
	//  [< 20>|<Jon       >|<Snow     >|<  2000>|<You know nothing, Jon Snow!>]
	//  [<300>|<Tyrion    >|<Lannister>|<  5000>|<                           >]
	//  {-----+------------+-----------+--------+-----------------------------}
	//  [<   >|<          >|<TOTAL    >|< 10000>|<                           >]
	//  \-----v------------v-----------v--------v-----------------------------/
	tableStyleTest = TableStyle{
		CaseFooter:           TextCaseUpper,
		CaseHeader:           TextCaseUpper,
		CaseRows:             TextCaseDefault,
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
		Name:                 "tableStyleTest",
	}
)

// TableStyle declares how to render the Table.
type TableStyle struct {
	CaseHeader           TextCase
	CaseFooter           TextCase
	CaseRows             TextCase
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
	Name                 string
}
