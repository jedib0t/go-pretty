package list

import (
	"fmt"

	"github.com/jedib0t/go-pretty/v6/text"
)

func Example() {
	lw := NewWriter()
	// append a tree
	lw.AppendItem("George. R. R. Martin")
	lw.Indent()
	lw.AppendItem("A Song of Ice and Fire")
	lw.Indent()
	lw.AppendItems([]interface{}{
		"Arya Stark",
		"Bran Stark",
		"Rickon Stark",
		"Robb Stark",
		"Sansa Stark",
		"Jon Snow",
	})
	lw.UnIndent()
	lw.UnIndent()
	// append another tree
	lw.AppendItem("Stephen King")
	lw.Indent()
	lw.AppendItem("The Dark Tower")
	lw.Indent()
	lw.AppendItems([]interface{}{
		"Jake Chambers",
		"Randal Flagg",
		"Roland Deschain",
	})
	lw.UnIndent()
	lw.AppendItem("the shawshank redemption")
	lw.Indent()
	lw.AppendItems([]interface{}{
		"andy dufresne",
		"byron hadley",
		"ellis boyd redding",
		"samuel norton",
	})
	// customize rendering
	lw.SetStyle(StyleConnectedLight)
	lw.Style().CharItemTop = "├"
	lw.Style().Format = text.FormatTitle
	// render it
	fmt.Printf("Simple List:\n%s", lw.Render())

	// Output: Simple List:
	// ├ George. R. R. Martin
	// │  └─ A Song Of Ice And Fire
	// │     ├─ Arya Stark
	// │     ├─ Bran Stark
	// │     ├─ Rickon Stark
	// │     ├─ Robb Stark
	// │     ├─ Sansa Stark
	// │     └─ Jon Snow
	// └─ Stephen King
	//    ├─ The Dark Tower
	//    │  ├─ Jake Chambers
	//    │  ├─ Randal Flagg
	//    │  └─ Roland Deschain
	//    └─ The Shawshank Redemption
	//       ├─ Andy Dufresne
	//       ├─ Byron Hadley
	//       ├─ Ellis Boyd Redding
	//       └─ Samuel Norton
}
