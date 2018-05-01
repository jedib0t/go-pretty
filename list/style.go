package list

import "github.com/jedib0t/go-pretty/text"

// Style declares how to render the List.
type Style struct {
	Format           text.Format
	CharConnect      string
	CharItem         string
	CharItemBottom   string
	CharItemFirst    string
	CharItemTop      string
	CharPaddingLeft  string
	CharPaddingRight string
	Name             string
}

var (
	// StyleDefault renders a List like below:
	//  - Game Of Thrones
	//  --- Winter
	//    - Is
	//    - Coming
	//    --- This
	//      - Is
	//      - Known
	StyleDefault = Style{
		Format:           text.FormatDefault,
		CharConnect:      "-",
		CharItem:         "-",
		CharItemBottom:   "-",
		CharItemFirst:    "-",
		CharItemTop:      "-",
		CharPaddingLeft:  "-",
		CharPaddingRight: "",
		Name:             "StyleDefault",
	}

	// StyleBulletCircle renders a List like below:
	//  ● Game Of Thrones
	//    ● Winter
	//    ● Is
	//    ● Coming
	//      ● This
	//      ● Is
	//      ● Known
	StyleBulletCircle = Style{
		Format:           text.FormatDefault,
		CharConnect:      " ",
		CharItem:         text.BulletCircle,
		CharItemBottom:   text.BulletCircle,
		CharItemFirst:    text.BulletCircle,
		CharItemTop:      text.BulletCircle,
		CharPaddingLeft:  " ",
		CharPaddingRight: "",
		Name:             "StyleBulletCircle",
	}

	// StyleBulletFlower renders a List like below:
	//  ✽ Game Of Thrones
	//    ✽ Winter
	//    ✽ Is
	//    ✽ Coming
	//      ✽ This
	//      ✽ Is
	//      ✽ Known
	StyleBulletFlower = Style{
		Format:           text.FormatDefault,
		CharConnect:      " ",
		CharItem:         text.BulletFlower,
		CharItemBottom:   text.BulletFlower,
		CharItemFirst:    text.BulletFlower,
		CharItemTop:      text.BulletFlower,
		CharPaddingLeft:  " ",
		CharPaddingRight: "",
		Name:             "StyleBulletFlower",
	}

	// StyleBulletSquare renders a List like below:
	//  ■ Game Of Thrones
	//    ■ Winter
	//    ■ Is
	//    ■ Coming
	//      ■ This
	//      ■ Is
	//      ■ Known
	StyleBulletSquare = Style{
		Format:           text.FormatDefault,
		CharConnect:      " ",
		CharItem:         text.BulletSquare,
		CharItemBottom:   text.BulletSquare,
		CharItemFirst:    text.BulletSquare,
		CharItemTop:      text.BulletSquare,
		CharPaddingLeft:  " ",
		CharPaddingRight: "",
		Name:             "StyleBulletSquare",
	}

	// StyleBulletStar renders a List like below:
	//  ✭ Game Of Thrones
	//    ✭ Winter
	//    ✭ Is
	//    ✭ Coming
	//      ✭ This
	//      ✭ Is
	//      ✭ Known
	StyleBulletStar = Style{
		Format:           text.FormatDefault,
		CharConnect:      " ",
		CharItem:         text.BulletStar,
		CharItemBottom:   text.BulletStar,
		CharItemFirst:    text.BulletStar,
		CharItemTop:      text.BulletStar,
		CharPaddingLeft:  " ",
		CharPaddingRight: "",
		Name:             "StyleBulletStar",
	}

	// StyleBulletTriangle renders a List like below:
	//  ▶ Game Of Thrones
	//    ▶ Winter
	//    ▶ Is
	//    ▶ Coming
	//      ▶ This
	//      ▶ Is
	//      ▶ Known
	StyleBulletTriangle = Style{
		Format:           text.FormatDefault,
		CharConnect:      " ",
		CharItem:         text.BulletTrianglePointingRight,
		CharItemBottom:   text.BulletTrianglePointingRight,
		CharItemFirst:    text.BulletTrianglePointingRight,
		CharItemTop:      text.BulletTrianglePointingRight,
		CharPaddingLeft:  " ",
		CharPaddingRight: "",
		Name:             "StyleBulletTriangle",
	}

	// StyleConnectedBold renders a List like below:
	//  ┏━ Game Of Thrones
	//  ┗━┳━ Winter
	//    ┣━ Is
	//    ┣━ Coming
	//    ┗━┳━ This
	//      ┣━ Is
	//      ┗━ Known
	StyleConnectedBold = Style{
		Format:           text.FormatDefault,
		CharConnect:      text.BoxBottomLeftBold,
		CharItem:         text.BoxLeftSeparatorBold,
		CharItemBottom:   text.BoxBottomLeftBold,
		CharItemFirst:    text.BoxTopSeparatorBold,
		CharItemTop:      text.BoxTopLeftBold,
		CharPaddingLeft:  text.BoxHorizontalBold,
		CharPaddingRight: text.BoxHorizontalBold,
		Name:             "StyleConnectedBold",
	}

	// StyleConnectedDouble renders a List like below:
	//  ╔═ Game Of Thrones
	//  ╚═╦═ Winter
	//    ╠═ Is
	//    ╠═ Coming
	//    ╚═╦═ This
	//      ╠═ Is
	//      ╚═ Known
	StyleConnectedDouble = Style{
		Format:           text.FormatDefault,
		CharConnect:      text.BoxBottomLeftDouble,
		CharItem:         text.BoxLeftSeparatorDouble,
		CharItemBottom:   text.BoxBottomLeftDouble,
		CharItemFirst:    text.BoxTopSeparatorDouble,
		CharItemTop:      text.BoxTopLeftDouble,
		CharPaddingLeft:  text.BoxHorizontalDouble,
		CharPaddingRight: text.BoxHorizontalDouble,
		Name:             "StyleConnectedDouble",
	}

	// StyleConnectedLight renders a List like below:
	//  ┌─ Game Of Thrones
	//  └─┬─ Winter
	//    ├─ Is
	//    ├─ Coming
	//    └─┬─ This
	//      ├─ Is
	//      └─ Known
	StyleConnectedLight = Style{
		Format:           text.FormatDefault,
		CharConnect:      text.BoxBottomLeft,
		CharItem:         text.BoxLeftSeparator,
		CharItemBottom:   text.BoxBottomLeft,
		CharItemFirst:    text.BoxTopSeparator,
		CharItemTop:      text.BoxTopLeft,
		CharPaddingLeft:  text.BoxHorizontal,
		CharPaddingRight: text.BoxHorizontal,
		Name:             "StyleConnectedLight",
	}

	// StyleConnectedRounded renders a List like below:
	//  ╭─ Game Of Thrones
	//  ╰─┬─ Winter
	//    ├─ Is
	//    ├─ Coming
	//    ╰─┬─ This
	//      ├─ Is
	//      ╰─ Known
	StyleConnectedRounded = Style{
		Format:           text.FormatDefault,
		CharConnect:      text.BoxBottomLeftRounded,
		CharItem:         text.BoxLeftSeparator,
		CharItemBottom:   text.BoxBottomLeftRounded,
		CharItemFirst:    text.BoxTopSeparator,
		CharItemTop:      text.BoxTopLeftRounded,
		CharPaddingLeft:  text.BoxHorizontal,
		CharPaddingRight: text.BoxHorizontal,
		Name:             "StyleConnectedRounded",
	}

	// styleTest renders a List like below:
	//  ^> Game Of Thrones
	//  c<f> Winter
	//    i> Is
	//    i> Coming
	//    c<f> This
	//      i> Is
	//      v> Known
	styleTest = Style{
		Format:           text.FormatDefault,
		CharConnect:      "c",
		CharItem:         "i",
		CharItemBottom:   "v",
		CharItemFirst:    "f",
		CharItemTop:      "^",
		CharPaddingLeft:  "<",
		CharPaddingRight: ">",
		Name:             "styleTest",
	}
)
