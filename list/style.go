package list

import "github.com/jedib0t/go-pretty/text"

// Style declares how to render the List.
type Style struct {
	Format              text.Format
	CharConnectBottom   string
	CharHorizontal      string
	CharItem            string
	CharItemBottom      string
	CharItemFirst       string
	CharItemSingle      string
	CharItemTop         string
	CharPaddingLeft     string
	CharPaddingRight    string
	CharVertical        string
	CharVerticalConnect string
	LinePrefix          string
	Name                string
}

var (
	// StyleDefault renders a List like below:
	//  * Game Of Thrones
	//    * Winter
	//    * Is
	//    * Coming
	//      * This
	//      * Is
	//      * Known
	//    * The Dark Tower
	//      * The Gunslinger
	StyleDefault = Style{
		Format:              text.FormatDefault,
		CharConnectBottom:   " ",
		CharHorizontal:      " ",
		CharItem:            "*",
		CharItemBottom:      "*",
		CharItemFirst:       "*",
		CharItemSingle:      "*",
		CharItemTop:         "*",
		CharPaddingLeft:     " ",
		CharPaddingRight:    "",
		CharVertical:        " ",
		CharVerticalConnect: " ",
		LinePrefix:          "",
		Name:                "StyleDefault",
	}

	// StyleBulletCircle renders a List like below:
	//  ● Game Of Thrones
	//    ● Winter
	//    ● Is
	//    ● Coming
	//      ● This
	//      ● Is
	//      ● Known
	//    ● The Dark Tower
	//      ● The Gunslinger
	StyleBulletCircle = Style{
		Format:              text.FormatDefault,
		CharConnectBottom:   " ",
		CharHorizontal:      " ",
		CharItem:            text.BulletCircle,
		CharItemBottom:      text.BulletCircle,
		CharItemFirst:       text.BulletCircle,
		CharItemSingle:      text.BulletCircle,
		CharItemTop:         text.BulletCircle,
		CharPaddingLeft:     " ",
		CharPaddingRight:    "",
		CharVertical:        " ",
		CharVerticalConnect: " ",
		LinePrefix:          "",
		Name:                "StyleBulletCircle",
	}

	// StyleBulletFlower renders a List like below:
	//  ✽ Game Of Thrones
	//    ✽ Winter
	//    ✽ Is
	//    ✽ Coming
	//      ✽ This
	//      ✽ Is
	//      ✽ Known
	//    ✽ The Dark Tower
	//      ✽ The Gunslinger
	StyleBulletFlower = Style{
		Format:              text.FormatDefault,
		CharConnectBottom:   " ",
		CharHorizontal:      " ",
		CharItem:            text.BulletFlower,
		CharItemBottom:      text.BulletFlower,
		CharItemFirst:       text.BulletFlower,
		CharItemSingle:      text.BulletFlower,
		CharItemTop:         text.BulletFlower,
		CharPaddingLeft:     " ",
		CharPaddingRight:    "",
		CharVertical:        " ",
		CharVerticalConnect: " ",
		LinePrefix:          "",
		Name:                "StyleBulletFlower",
	}

	// StyleBulletSquare renders a List like below:
	//  ■ Game Of Thrones
	//    ■ Winter
	//    ■ Is
	//    ■ Coming
	//      ■ This
	//      ■ Is
	//      ■ Known
	//    ■ The Dark Tower
	//      ■ The Gunslinger
	StyleBulletSquare = Style{
		Format:              text.FormatDefault,
		CharConnectBottom:   " ",
		CharHorizontal:      " ",
		CharItem:            text.BulletSquare,
		CharItemBottom:      text.BulletSquare,
		CharItemFirst:       text.BulletSquare,
		CharItemSingle:      text.BulletSquare,
		CharItemTop:         text.BulletSquare,
		CharPaddingLeft:     " ",
		CharPaddingRight:    "",
		CharVertical:        " ",
		CharVerticalConnect: " ",
		LinePrefix:          "",
		Name:                "StyleBulletSquare",
	}

	// StyleBulletStar renders a List like below:
	//  ★ Game Of Thrones
	//    ★ Winter
	//    ★ Is
	//    ★ Coming
	//      ★ This
	//      ★ Is
	//      ★ Known
	//    ★ The Dark Tower
	//      ★ The Gunslinger
	StyleBulletStar = Style{
		Format:              text.FormatDefault,
		CharConnectBottom:   " ",
		CharHorizontal:      " ",
		CharItem:            text.BulletStar,
		CharItemBottom:      text.BulletStar,
		CharItemFirst:       text.BulletStar,
		CharItemSingle:      text.BulletStar,
		CharItemTop:         text.BulletStar,
		CharPaddingLeft:     " ",
		CharPaddingRight:    "",
		CharVertical:        " ",
		CharVerticalConnect: " ",
		LinePrefix:          "",
		Name:                "StyleBulletStar",
	}

	// StyleBulletTriangle renders a List like below:
	//  ▶ Game Of Thrones
	//    ▶ Winter
	//    ▶ Is
	//    ▶ Coming
	//      ▶ This
	//      ▶ Is
	//      ▶ Known
	//    ▶ The Dark Tower
	//      ▶ The Gunslinger
	StyleBulletTriangle = Style{
		Format:              text.FormatDefault,
		CharConnectBottom:   " ",
		CharHorizontal:      " ",
		CharItem:            text.BulletTrianglePointingRight,
		CharItemBottom:      text.BulletTrianglePointingRight,
		CharItemFirst:       text.BulletTrianglePointingRight,
		CharItemSingle:      text.BulletTrianglePointingRight,
		CharItemTop:         text.BulletTrianglePointingRight,
		CharPaddingLeft:     " ",
		CharPaddingRight:    "",
		CharVertical:        " ",
		CharVerticalConnect: " ",
		LinePrefix:          "",
		Name:                "StyleBulletTriangle",
	}

	// StyleConnectedBold renders a List like below:
	//  ┏━ Game Of Thrones
	//  ┗━┳━ Winter
	//    ┣━ Is
	//    ┣━ Coming
	//    ┣━┳━ This
	//    ┃ ┣━ Is
	//    ┃ ┗━ Known
	//    ┣━ The Dark Tower
	//    ┗━━━ The Gunslinger
	StyleConnectedBold = Style{
		Format:              text.FormatDefault,
		CharConnectBottom:   text.BoxBottomLeftBold,
		CharHorizontal:      text.BoxHorizontalBold,
		CharItem:            text.BoxLeftSeparatorBold,
		CharItemBottom:      text.BoxBottomLeftBold,
		CharItemFirst:       text.BoxTopSeparatorBold,
		CharItemSingle:      text.BoxHorizontalBold,
		CharItemTop:         text.BoxTopLeftBold,
		CharPaddingLeft:     text.BoxHorizontalBold,
		CharPaddingRight:    text.BoxHorizontalBold,
		CharVertical:        text.BoxVerticalBold,
		CharVerticalConnect: text.BoxLeftSeparatorBold,
		LinePrefix:          "",
		Name:                "StyleConnectedBold",
	}

	// StyleConnectedDouble renders a List like below:
	//  ╔═ Game Of Thrones
	//  ╚═╦═ Winter
	//    ╠═ Is
	//    ╠═ Coming
	//    ╠═╦═ This
	//    ║ ╠═ Is
	//    ║ ╚═ Known
	//    ╠═ The Dark Tower
	//    ╚═══ The Gunslinger
	StyleConnectedDouble = Style{
		Format:              text.FormatDefault,
		CharConnectBottom:   text.BoxBottomLeftDouble,
		CharHorizontal:      text.BoxHorizontalDouble,
		CharItem:            text.BoxLeftSeparatorDouble,
		CharItemBottom:      text.BoxBottomLeftDouble,
		CharItemFirst:       text.BoxTopSeparatorDouble,
		CharItemSingle:      text.BoxHorizontalDouble,
		CharItemTop:         text.BoxTopLeftDouble,
		CharPaddingLeft:     text.BoxHorizontalDouble,
		CharPaddingRight:    text.BoxHorizontalDouble,
		CharVertical:        text.BoxVerticalDouble,
		CharVerticalConnect: text.BoxLeftSeparatorDouble,
		LinePrefix:          "",
		Name:                "StyleConnectedDouble",
	}

	// StyleConnectedLight renders a List like below:
	//  ┌─ Game Of Thrones
	//  └─┬─ Winter
	//    ├─ Is
	//    ├─ Coming
	//    ├─┬─ This
	//    │ ├─ Is
	//    │ └─ Known
	//    ├─ The Dark Tower
	//    └─── The Gunslinger
	StyleConnectedLight = Style{
		Format:              text.FormatDefault,
		CharConnectBottom:   text.BoxBottomLeft,
		CharHorizontal:      text.BoxHorizontal,
		CharItem:            text.BoxLeftSeparator,
		CharItemBottom:      text.BoxBottomLeft,
		CharItemFirst:       text.BoxTopSeparator,
		CharItemSingle:      text.BoxHorizontal,
		CharItemTop:         text.BoxTopLeft,
		CharPaddingLeft:     text.BoxHorizontal,
		CharPaddingRight:    text.BoxHorizontal,
		CharVertical:        text.BoxVertical,
		CharVerticalConnect: text.BoxLeftSeparator,
		LinePrefix:          "",
		Name:                "StyleConnectedLight",
	}

	// StyleConnectedRounded renders a List like below:
	//  ╭─ Game Of Thrones
	//  ╰─┬─ Winter
	//    ├─ Is
	//    ├─ Coming
	//    ├─┬─ This
	//    │ ├─ Is
	//    │ ╰─ Known
	//    ├─ The Dark Tower
	//    ╰─── The Gunslinger
	StyleConnectedRounded = Style{
		Format:              text.FormatDefault,
		CharConnectBottom:   text.BoxBottomLeftRounded,
		CharHorizontal:      text.BoxHorizontal,
		CharItem:            text.BoxLeftSeparator,
		CharItemBottom:      text.BoxBottomLeftRounded,
		CharItemFirst:       text.BoxTopSeparator,
		CharItemSingle:      text.BoxHorizontal,
		CharItemTop:         text.BoxTopLeftRounded,
		CharPaddingLeft:     text.BoxHorizontal,
		CharPaddingRight:    text.BoxHorizontal,
		CharVertical:        text.BoxVertical,
		CharVerticalConnect: text.BoxLeftSeparator,
		LinePrefix:          "",
		Name:                "StyleConnectedRounded",
	}

	// StyleMarkdown renders a List like below:
	//    * Game Of Thrones
	//      * Winter
	//      * Is
	//      * Coming
	//        * This
	//        * Is
	//        * Known
	//      * The Dark Tower
	//        * The Gunslinger
	StyleMarkdown = Style{
		Format:              text.FormatDefault,
		CharConnectBottom:   " ",
		CharHorizontal:      " ",
		CharItem:            "*",
		CharItemBottom:      "*",
		CharItemFirst:       "*",
		CharItemSingle:      "*",
		CharItemTop:         "*",
		CharPaddingLeft:     " ",
		CharPaddingRight:    "",
		CharVertical:        " ",
		CharVerticalConnect: " ",
		LinePrefix:          "  ",
		Name:                "StyleMarkdown",
	}

	// styleTest renders a List like below:
	//  ^> Game Of Thrones
	//  c~f> Winter
	//    i> Is
	//    i> Coming
	//    T~f> This
	//    | i> Is
	//    | v> Known
	//    i> The Dark Tower
	//    c~I> The Gunslinger
	styleTest = Style{
		Format:              text.FormatDefault,
		CharConnectBottom:   "c",
		CharHorizontal:      "~",
		CharItem:            "i",
		CharItemBottom:      "v",
		CharItemFirst:       "f",
		CharItemSingle:      "I",
		CharItemTop:         "^",
		CharPaddingLeft:     "<",
		CharPaddingRight:    ">",
		CharVertical:        "|",
		CharVerticalConnect: "T",
		LinePrefix:          "",
		Name:                "styleTest",
	}
)
