package main

import (
	"fmt"

	"github.com/jedib0t/go-pretty/list"
	"github.com/jedib0t/go-pretty/text"
)

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
	fmt.Printf("A Simple List:\n%s\n\n", l.Render())
	//A Simple List:
	//- Game Of Thrones
	//==========================================================================

	//==========================================================================
	// I wanna Level Down!
	//==========================================================================
	l.Indent()
	l.AppendItems([]interface{}{"Winter", "Is", "Coming"})
	l.Indent()
	l.AppendItems([]interface{}{"This", "Is", "Known"})
	fmt.Printf("A Multi-level List:\n%s\n\n", l.Render())
	//A Multi-level List:
	//- Game Of Thrones
	//--- Winter
	//  - Is
	//  - Coming
	//  --- This
	//    - Is
	//    - Known
	//==========================================================================

	//==========================================================================
	// I am Fancy!
	//==========================================================================
	l.SetStyle(list.StyleBulletCircle)
	fmt.Printf("A List using the Style 'StyleBulletCircle':\n%s\n\n", l.Render())
	//A List using the Style 'StyleBulletCircle':
	//● Game Of Thrones
	//  ● Winter
	//  ● Is
	//  ● Coming
	//    ● This
	//    ● Is
	//    ● Known
	//
	l.SetStyle(list.StyleConnectedRounded)
	fmt.Printf("A List using the Style 'StyleConnectedRounded':\n%s\n\n", l.Render())
	//A List using the Style 'StyleConnectedRounded':
	//╭─ Game Of Thrones
	//╰─┬─ Winter
	//  ├─ Is
	//  ├─ Coming
	//  ╰─┬─ This
	//    ├─ Is
	//    ╰─ Known
	//==========================================================================

	//==========================================================================
	// I want my own Style!
	//==========================================================================
	funkyStyle := list.Style{
		CharConnect:      "c",
		CharItem:         "i",
		CharItemBottom:   "v",
		CharItemFirst:    "f",
		CharItemTop:      "^",
		CharPaddingLeft:  "<",
		CharPaddingRight: ">",
		Format:           text.FormatUpper,
		Name:             "funkyStyle",
	}
	l.SetStyle(funkyStyle)
	fmt.Printf("A List using the Style 'funkyStyle':\n%s\n\n", l.Render())
	//A List using the Style 'funkyStyle':
	//^> 1
	//i> TWO
	//i> 3.14159265359
	//c<f> GAME
	//  i> OF
	//  i> THRONES
	//  c<f> SEASON
	//    v> 7
	//==========================================================================
}
