package list

import "github.com/jedib0t/go-pretty/text"

// Style declares how to render the List.
type Style struct {
	Format           text.Format
	CharItemSingle   string
	CharItemTop      string
	CharItemFirst    string
	CharItemMiddle   string
	CharItemVertical string
	CharItemBottom   string
	CharNewline      string
	LinePrefix       string
	Name             string
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
	//  * The Dark Tower
	//    * The Gunslinger
	StyleDefault = Style{
		Format:           text.FormatDefault,
		CharItemSingle:   "*",
		CharItemTop:      "*",
		CharItemFirst:    "*",
		CharItemMiddle:   "*",
		CharItemVertical: "  ",
		CharItemBottom:   "*",
		CharNewline:      "\n",
		LinePrefix:       "",
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
	//  ● The Dark Tower
	//    ● The Gunslinger
	StyleBulletCircle = Style{
		Format:           text.FormatDefault,
		CharItemSingle:   text.BulletCircle,
		CharItemTop:      text.BulletCircle,
		CharItemFirst:    text.BulletCircle,
		CharItemMiddle:   text.BulletCircle,
		CharItemVertical: "  ",
		CharItemBottom:   text.BulletCircle,
		CharNewline:      "\n",
		LinePrefix:       "",
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
	//  ✽ The Dark Tower
	//    ✽ The Gunslinger
	StyleBulletFlower = Style{
		Format:           text.FormatDefault,
		CharItemSingle:   text.BulletFlower,
		CharItemTop:      text.BulletFlower,
		CharItemFirst:    text.BulletFlower,
		CharItemMiddle:   text.BulletFlower,
		CharItemVertical: "  ",
		CharItemBottom:   text.BulletFlower,
		CharNewline:      "\n",
		LinePrefix:       "",
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
	//  ■ The Dark Tower
	//    ■ The Gunslinger
	StyleBulletSquare = Style{
		Format:           text.FormatDefault,
		CharItemSingle:   text.BulletSquare,
		CharItemTop:      text.BulletSquare,
		CharItemFirst:    text.BulletSquare,
		CharItemMiddle:   text.BulletSquare,
		CharItemVertical: "  ",
		CharItemBottom:   text.BulletSquare,
		CharNewline:      "\n",
		LinePrefix:       "",
		Name:             "StyleBulletSquare",
	}

	// StyleBulletStar renders a List like below:
	//  ★ Game Of Thrones
	//    ★ Winter
	//    ★ Is
	//    ★ Coming
	//      ★ This
	//      ★ Is
	//      ★ Known
	//  ★ The Dark Tower
	//    ★ The Gunslinger
	StyleBulletStar = Style{
		Format:           text.FormatDefault,
		CharItemSingle:   text.BulletStar,
		CharItemTop:      text.BulletStar,
		CharItemFirst:    text.BulletStar,
		CharItemMiddle:   text.BulletStar,
		CharItemVertical: "  ",
		CharItemBottom:   text.BulletStar,
		CharNewline:      "\n",
		LinePrefix:       "",
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
	//  ▶ The Dark Tower
	//    ▶ The Gunslinger
	StyleBulletTriangle = Style{
		Format:           text.FormatDefault,
		CharItemSingle:   text.BulletTrianglePointingRight,
		CharItemTop:      text.BulletTrianglePointingRight,
		CharItemFirst:    text.BulletTrianglePointingRight,
		CharItemMiddle:   text.BulletTrianglePointingRight,
		CharItemVertical: "  ",
		CharItemBottom:   text.BulletTrianglePointingRight,
		CharNewline:      "\n",
		LinePrefix:       "",
		Name:             "StyleBulletTriangle",
	}

	// StyleConnectedBold renders a List like below:
	//  ┏━ Game Of Thrones
	//  ┃  ┣━ Winter
	//  ┃  ┣━ Is
	//  ┃  ┗━ Coming
	//  ┃     ┣━ This
	//  ┃     ┣━ Is
	//  ┃     ┗━ Known
	//  ┗━ The Dark Tower
	//     ┗━ The Gunslinger
	StyleConnectedBold = Style{
		Format:           text.FormatDefault,
		CharItemSingle:   "━━",
		CharItemTop:      "┏━",
		CharItemFirst:    "┣━",
		CharItemMiddle:   "┣━",
		CharItemVertical: "┃  ",
		CharItemBottom:   "┗━",
		CharNewline:      "\n",
		LinePrefix:       "",
		Name:             "StyleConnectedBold",
	}

	// StyleConnectedDouble renders a List like below:
	//  ╔═ Game Of Thrones
	//  ║  ╠═ Winter
	//  ║  ╠═ Is
	//  ║  ╚═ Coming
	//  ║     ╠═ This
	//  ║     ╠═ Is
	//  ║     ╚═ Known
	//  ╚═ The Dark Tower
	//     ╚═ The Gunslinger
	StyleConnectedDouble = Style{
		Format:           text.FormatDefault,
		CharItemSingle:   "══",
		CharItemTop:      "╔═",
		CharItemFirst:    "╠═",
		CharItemMiddle:   "╠═",
		CharItemVertical: "║  ",
		CharItemBottom:   "╚═",
		CharNewline:      "\n",
		LinePrefix:       "",
		Name:             "StyleConnectedDouble",
	}

	// StyleConnectedLight renders a List like below:
	//  ┌─ Game Of Thrones
	//  │  ├─ Winter
	//  │  ├─ Is
	//  │  └─ Coming
	//  │     ├─ This
	//  │     ├─ Is
	//  │     └─ Known
	//  └─ The Dark Tower
	//     └─ The Gunslinger
	StyleConnectedLight = Style{
		Format:           text.FormatDefault,
		CharItemSingle:   "──",
		CharItemTop:      "┌─",
		CharItemFirst:    "├─",
		CharItemMiddle:   "├─",
		CharItemVertical: "│  ",
		CharItemBottom:   "└─",
		CharNewline:      "\n",
		LinePrefix:       "",
		Name:             "StyleConnectedLight",
	}

	// StyleConnectedRounded renders a List like below:
	//  ╭─ Game Of Thrones
	//  │  ├─ Winter
	//  │  ├─ Is
	//  │  ╰─ Coming
	//  │     ├─ This
	//  │     ├─ Is
	//  │     ╰─ Known
	//  ╰─ The Dark Tower
	//     ╰─ The Gunslinger
	StyleConnectedRounded = Style{
		Format:           text.FormatDefault,
		CharItemSingle:   "──",
		CharItemTop:      "╭─",
		CharItemFirst:    "├─",
		CharItemMiddle:   "├─",
		CharItemVertical: "│  ",
		CharItemBottom:   "╰─",
		CharNewline:      "\n",
		LinePrefix:       "",
		Name:             "StyleConnectedRounded",
	}

	// StyleMarkdown renders a List like below:
	//    * Game Of Thrones
	//      * Winter
	//      * Is
	//      * Coming
	//        * This
	//        * Is
	//        * Known
	//    * The Dark Tower
	//      * The Gunslinger
	StyleMarkdown = Style{
		Format:           text.FormatDefault,
		CharItemSingle:   "*",
		CharItemTop:      "*",
		CharItemFirst:    "*",
		CharItemMiddle:   "*",
		CharItemVertical: "  ",
		CharItemBottom:   "*",
		CharNewline:      "<br/>",
		LinePrefix:       "  ",
		Name:             "StyleMarkdown",
	}

	// styleTest renders a List like below:
	//  t Game Of Thrones
	//  |f Winter
	//  |m Is
	//  |b Coming
	//  | f This
	//  | m Is
	//  | b Known
	//  b The Dark Tower
	//   b The Gunslinger
	styleTest = Style{
		Format:           text.FormatDefault,
		CharItemSingle:   "s",
		CharItemTop:      "t",
		CharItemFirst:    "f",
		CharItemMiddle:   "m",
		CharItemVertical: "|",
		CharItemBottom:   "b",
		CharNewline:      "\n",
		LinePrefix:       "",
		Name:             "styleTest",
	}
)
