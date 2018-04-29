package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty"
)

var (
	listCaptionColor = gopretty.TextColor{color.FgHiYellow}
)

func showList() {
	for _, style := range []gopretty.ListStyle{gopretty.ListStyleBulletCircle, gopretty.ListStyleConnectedBold} {
		lw := gopretty.NewListWriter()
		lw.AppendItem("Game Of Thrones")
		lw.Indent()
		lw.AppendItems([]interface{}{"Winter", "Is", "Coming"})
		lw.Indent()
		lw.AppendItems([]interface{}{"This", "Is", "Known"})
		lw.SetStyle(style)

		fmt.Println(lw.Render())
		fmt.Println(listCaptionColor.Sprintf("A List using the style '%s'.\n", lw.Style().Name))
	}
}

func main() {
	showList()
}
