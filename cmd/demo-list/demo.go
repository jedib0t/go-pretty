package main

import (
	"fmt"
	"strings"

	"github.com/jedib0t/go-pretty/list"
	"github.com/jedib0t/go-pretty/text"
)

func demoPrint(title string, content string, prefix string) {
	fmt.Printf("%s:\n", title)
	fmt.Println(strings.Repeat("-", len(title)+1))
	for _, line := range strings.Split(content, "\n") {
		fmt.Printf("%s%s\n", prefix, line)
	}
	fmt.Println()
}

func main() {
	//==========================================================================
	// Initialization
	//==========================================================================
	l := list.NewWriter()
	// you can also instantiate the object directly
	lTemp := list.List{}
	lTemp.Render() // just to avoid the compile error of not using the object
	//==========================================================================

	//==========================================================================
	// A List needs Items.
	//==========================================================================
	l.AppendItem("Game Of Thrones")
	l.AppendItem("The Dark Tower")
	demoPrint("A Simple List", l.Render(), "")
	//A Simple List:
	//--------------
	//* Game Of Thrones
	//* The Dark Tower
	l.Reset()
	//==========================================================================

	//==========================================================================
	// I wanna Level Down!
	//==========================================================================
	l.AppendItem("Game Of Thrones")
	l.Indent()
	l.AppendItems([]interface{}{"Winter", "Is", "Coming"})
	l.Indent()
	l.AppendItems([]interface{}{"This", "Is", "Known"})
	l.UnIndent()
	l.AppendItem("The Dark Tower")
	l.Indent()
	l.AppendItem("The Gunslinger")
	demoPrint("A Multi-level List", l.Render(), "")
	//A Multi-level List:
	//-------------------
	//* Game Of Thrones
	//  * Winter
	//  * Is
	//  * Coming
	//    * This
	//    * Is
	//    * Known
	//  * The Dark Tower
	//    * The Gunslinger
	//==========================================================================

	//==========================================================================
	// I am Fancy!
	//==========================================================================
	l.SetStyle(list.StyleBulletCircle)
	demoPrint("A List using the Style 'StyleBulletCircle'", l.Render(), "")
	//A List using the Style 'StyleBulletCircle':
	//-------------------------------------------
	//● Game Of Thrones
	//  ● Winter
	//  ● Is
	//  ● Coming
	//    ● This
	//    ● Is
	//    ● Known
	//  ● The Dark Tower
	//    ● The Gunslinger
	l.SetStyle(list.StyleConnectedRounded)
	demoPrint("A List using the Style 'StyleConnectedRounded'", l.Render(), "")
	//A List using the Style 'StyleConnectedRounded':
	//-----------------------------------------------
	//╭─ Game Of Thrones
	//╰─┬─ Winter
	//  ├─ Is
	//  ├─ Coming
	//  ├─┬─ This
	//  │ ├─ Is
	//  │ ╰─ Known
	//  ├─ The Dark Tower
	//  ╰─── The Gunslinger
	//==========================================================================

	//==========================================================================
	// I want my own Style!
	//==========================================================================
	funkyStyle := list.Style{
		Name:              "funkyStyle",
		CharConnectBottom: "c",
		CharHorizontal:    "~",
		CharItem:          "i",
		CharItemBottom:    "v",
		CharItemFirst:     "f",
		CharItemTop:       "^",
		CharPaddingLeft:   "<",
		CharPaddingRight:  ">",
		Format:            text.FormatUpper,
	}
	l.SetStyle(funkyStyle)
	demoPrint("A List using the Style 'funkyStyle'", l.Render(), "")
	//A List using the Style 'funkyStyle':
	//------------------------------------
	//^> GAME OF THRONES
	//c~f> WINTER
	//  i> IS
	//  i> COMING
	//  ~f> THIS
	//   i> IS
	//   v> KNOWN
	//  i> THE DARK TOWER
	//  c~> THE GUNSLINGER
	//==========================================================================

	//==========================================================================
	// Can I get the list in Markdown format?
	//==========================================================================
	demoPrint("A List in Markdown format", l.RenderMarkdown(), "[Markdown] ")
	//A List in Markdown format:
	//--------------------------
	//[Markdown]   * Game Of Thrones
	//[Markdown]     * Winter
	//[Markdown]     * Is
	//[Markdown]     * Coming
	//[Markdown]       * This
	//[Markdown]       * Is
	//[Markdown]       * Known
	//[Markdown]     * The Dark Tower
	//[Markdown]       * The Gunslinger
	//==========================================================================
}
