package gopretty

var (
	// ListStyleDefault renders a List like below:
	//  - Game Of Thrones
	//  --- Winter
	//    - Is
	//    - Coming
	//    --- This
	//      - Is
	//      - Known
	ListStyleDefault = ListStyle{
		Case:             TextCaseDefault,
		CharConnect:      "-",
		CharItem:         "-",
		CharItemBottom:   "-",
		CharItemFirst:    "-",
		CharItemTop:      "-",
		CharPaddingLeft:  "-",
		CharPaddingRight: "",
		Name:             "ListStyleDefault",
	}

	// ListStyleBulletCircle renders a List like below:
	//  ● Game Of Thrones
	//    ● Winter
	//    ● Is
	//    ● Coming
	//      ● This
	//      ● Is
	//      ● Known
	ListStyleBulletCircle = ListStyle{
		Case:             TextCaseDefault,
		CharConnect:      " ",
		CharItem:         BulletCircle,
		CharItemBottom:   BulletCircle,
		CharItemFirst:    BulletCircle,
		CharItemTop:      BulletCircle,
		CharPaddingLeft:  " ",
		CharPaddingRight: "",
		Name:             "ListStyleBulletCircle",
	}

	// ListStyleBulletFlower renders a List like below:
	//  ✽ Game Of Thrones
	//    ✽ Winter
	//    ✽ Is
	//    ✽ Coming
	//      ✽ This
	//      ✽ Is
	//      ✽ Known
	ListStyleBulletFlower = ListStyle{
		Case:             TextCaseDefault,
		CharConnect:      " ",
		CharItem:         BulletFlower,
		CharItemBottom:   BulletFlower,
		CharItemFirst:    BulletFlower,
		CharItemTop:      BulletFlower,
		CharPaddingLeft:  " ",
		CharPaddingRight: "",
		Name:             "ListStyleBulletFlower",
	}

	// ListStyleBulletSquare renders a List like below:
	//  ■ Game Of Thrones
	//    ■ Winter
	//    ■ Is
	//    ■ Coming
	//      ■ This
	//      ■ Is
	//      ■ Known
	ListStyleBulletSquare = ListStyle{
		Case:             TextCaseDefault,
		CharConnect:      " ",
		CharItem:         BulletSquare,
		CharItemBottom:   BulletSquare,
		CharItemFirst:    BulletSquare,
		CharItemTop:      BulletSquare,
		CharPaddingLeft:  " ",
		CharPaddingRight: "",
		Name:             "ListStyleBulletSquare",
	}

	// ListStyleBulletStar renders a List like below:
	//  ✭ Game Of Thrones
	//    ✭ Winter
	//    ✭ Is
	//    ✭ Coming
	//      ✭ This
	//      ✭ Is
	//      ✭ Known
	ListStyleBulletStar = ListStyle{
		Case:             TextCaseDefault,
		CharConnect:      " ",
		CharItem:         BulletStar,
		CharItemBottom:   BulletStar,
		CharItemFirst:    BulletStar,
		CharItemTop:      BulletStar,
		CharPaddingLeft:  " ",
		CharPaddingRight: "",
		Name:             "ListStyleBulletStar",
	}

	// ListStyleBulletTriangle renders a List like below:
	//  ▶ Game Of Thrones
	//    ▶ Winter
	//    ▶ Is
	//    ▶ Coming
	//      ▶ This
	//      ▶ Is
	//      ▶ Known
	ListStyleBulletTriangle = ListStyle{
		Case:             TextCaseDefault,
		CharConnect:      " ",
		CharItem:         BulletTrianglePointingRight,
		CharItemBottom:   BulletTrianglePointingRight,
		CharItemFirst:    BulletTrianglePointingRight,
		CharItemTop:      BulletTrianglePointingRight,
		CharPaddingLeft:  " ",
		CharPaddingRight: "",
		Name:             "ListStyleBulletTriangle",
	}

	// ListStyleConnectedBold renders a List like below:
	//  ┏━ Game Of Thrones
	//  ┗━┳━ Winter
	//    ┣━ Is
	//    ┣━ Coming
	//    ┗━┳━ This
	//      ┣━ Is
	//      ┗━ Known
	ListStyleConnectedBold = ListStyle{
		Case:             TextCaseDefault,
		CharConnect:      BorderBottomLeftBold,
		CharItem:         BorderLeftSeparatorBold,
		CharItemBottom:   BorderBottomLeftBold,
		CharItemFirst:    BorderTopSeparatorBold,
		CharItemTop:      BorderTopLeftBold,
		CharPaddingLeft:  BorderHorizontalBold,
		CharPaddingRight: BorderHorizontalBold,
		Name:             "ListStyleConnectedBold",
	}

	// ListStyleConnectedDouble renders a List like below:
	//  ╔═ Game Of Thrones
	//  ╚═╦═ Winter
	//    ╠═ Is
	//    ╠═ Coming
	//    ╚═╦═ This
	//      ╠═ Is
	//      ╚═ Known
	ListStyleConnectedDouble = ListStyle{
		Case:             TextCaseDefault,
		CharConnect:      BorderBottomLeftDouble,
		CharItem:         BorderLeftSeparatorDouble,
		CharItemBottom:   BorderBottomLeftDouble,
		CharItemFirst:    BorderTopSeparatorDouble,
		CharItemTop:      BorderTopLeftDouble,
		CharPaddingLeft:  BorderHorizontalDouble,
		CharPaddingRight: BorderHorizontalDouble,
		Name:             "ListStyleConnectedDouble",
	}

	// ListStyleConnectedLight renders a List like below:
	//  ┌─ Game Of Thrones
	//  └─┬─ Winter
	//    ├─ Is
	//    ├─ Coming
	//    └─┬─ This
	//      ├─ Is
	//      └─ Known
	ListStyleConnectedLight = ListStyle{
		Case:             TextCaseDefault,
		CharConnect:      BorderBottomLeft,
		CharItem:         BorderLeftSeparator,
		CharItemBottom:   BorderBottomLeft,
		CharItemFirst:    BorderTopSeparator,
		CharItemTop:      BorderTopLeft,
		CharPaddingLeft:  BorderHorizontal,
		CharPaddingRight: BorderHorizontal,
		Name:             "ListStyleConnectedLight",
	}

	// ListStyleConnectedRounded renders a List like below:
	//  ╭─ Game Of Thrones
	//  ╰─┬─ Winter
	//    ├─ Is
	//    ├─ Coming
	//    ╰─┬─ This
	//      ├─ Is
	//      ╰─ Known
	ListStyleConnectedRounded = ListStyle{
		Case:             TextCaseDefault,
		CharConnect:      BorderBottomLeftRounded,
		CharItem:         BorderLeftSeparator,
		CharItemBottom:   BorderBottomLeftRounded,
		CharItemFirst:    BorderTopSeparator,
		CharItemTop:      BorderTopLeftRounded,
		CharPaddingLeft:  BorderHorizontal,
		CharPaddingRight: BorderHorizontal,
		Name:             "ListStyleConnectedRounded",
	}

	// listStyleTest renders a List like below:
	//  ^> Game Of Thrones
	//  c<f> Winter
	//    i> Is
	//    i> Coming
	//    c<f> This
	//      i> Is
	//      v> Known
	listStyleTest = ListStyle{
		Case:             TextCaseDefault,
		CharConnect:      "c",
		CharItem:         "i",
		CharItemBottom:   "v",
		CharItemFirst:    "f",
		CharItemTop:      "^",
		CharPaddingLeft:  "<",
		CharPaddingRight: ">",
		Name:             "listStyleTest",
	}
)

// ListStyle declares how to render the List.
type ListStyle struct {
	Case             TextCase
	CharConnect      string
	CharItem         string
	CharItemBottom   string
	CharItemFirst    string
	CharItemTop      string
	CharPaddingLeft  string
	CharPaddingRight string
	Name             string
}
