package main

import (
	"fmt"
	"github.com/jedib0t/go-pretty/text"
	"github.com/jedib0t/go-pretty/util"
)

func main() {
	fmt.Printf("%#v\n", text.FgRed.Sprint("Red"))
	fmt.Printf("%s\n", text.FgRed.Sprint("Red"))
	fmt.Printf("%c[31mRed%c[0m\n", util.EscapeStartRune, util.EscapeStartRune)
	fmt.Printf("[31mRed[0m\n")
}
