package main

import (
	"fmt"
	"strings"

	"github.com/jedib0t/go-pretty/v6/list"
	"github.com/jedib0t/go-pretty/v6/text"
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
	//* The Dark Tower
	//  * The Gunslinger
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
	//● The Dark Tower
	//  ● The Gunslinger
	l.SetStyle(list.StyleConnectedRounded)
	demoPrint("A List using the Style 'StyleConnectedRounded'", l.Render(), "")
	//A List using the Style 'StyleConnectedRounded':
	//-----------------------------------------------
	//╭─ Game Of Thrones
	//├─┬─ Winter
	//│ ├─ Is
	//│ ├─ Coming
	//│ ╰─┬─ This
	//│   ├─ Is
	//│   ╰─ Known
	//├─ The Dark Tower
	//╰─── The Gunslinger
	//==========================================================================

	//==========================================================================
	// I want my own Style!
	//==========================================================================
	funkyStyle := list.Style{
		CharItemSingle:   "s",
		CharItemTop:      "t",
		CharItemFirst:    "f",
		CharItemMiddle:   "m",
		CharItemVertical: "|",
		CharItemBottom:   "b",
		CharNewline:      "\n",
		Format:           text.FormatUpper,
		LinePrefix:       "",
		Name:             "styleTest",
	}
	l.SetStyle(funkyStyle)
	demoPrint("A List using the Style 'funkyStyle'", l.Render(), "")
	//A List using the Style 'funkyStyle':
	//------------------------------------
	//t GAME OF THRONES
	//|f WINTER
	//|m IS
	//|b COMING
	//| f THIS
	//| m IS
	//| b KNOWN
	//b THE DARK TOWER
	// b THE GUNSLINGER
	//==========================================================================

	//==========================================================================
	// I want to use it in a HTML file!
	//==========================================================================
	demoPrint("A List in HTML format", l.RenderHTML(), "[HTML] ")
	//A List in HTML format:
	//----------------------
	//[HTML] <ul class="go-pretty-table">
	//[HTML]   <li>Game Of Thrones</li>
	//[HTML]   <ul class="go-pretty-table-1">
	//[HTML]     <li>Winter</li>
	//[HTML]     <li>Is</li>
	//[HTML]     <li>Coming</li>
	//[HTML]     <ul class="go-pretty-table-2">
	//[HTML]       <li>This</li>
	//[HTML]       <li>Is</li>
	//[HTML]       <li>Known</li>
	//[HTML]     </ul>
	//[HTML]   </ul>
	//[HTML]   <li>The Dark Tower</li>
	//[HTML]   <ul class="go-pretty-table-1">
	//[HTML]     <li>The Gunslinger</li>
	//[HTML]   </ul>
	//[HTML] </ul>
	//==========================================================================

	//==========================================================================
	// Can I get the list in Markdown format?
	//==========================================================================
	demoPrint("A List in Markdown format", l.RenderMarkdown(), "[Markdown] ")
	fmt.Println()
	//A List in Markdown format:
	//--------------------------
	//[Markdown]   * Game Of Thrones
	//[Markdown]     * Winter
	//[Markdown]     * Is
	//[Markdown]     * Coming
	//[Markdown]       * This
	//[Markdown]       * Is
	//[Markdown]       * Known
	//[Markdown]   * The Dark Tower
	//[Markdown]     * The Gunslinger
	//==========================================================================
}
