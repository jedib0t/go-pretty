package main

import (
	"os"
	"fmt"
	"strconv"

	"github.com/jedib0t/go-pretty/list"
	"github.com/pkg/profile"
)

var (
	listItem1  = "Game Of Thrones"
	listItems2 = []interface{}{"Winter", "Is", "Coming"}
	listItems3 = []interface{}{"This", "Is", "Known"}
	profilers  = []func(*profile.Profile){
		profile.CPUProfile,
		profile.MemProfile,
	}
)

func profileRender(profiler func(profile2 *profile.Profile), n int) {
	defer profile.Start(profiler, profile.ProfilePath("./")).Stop()

	for i := 0; i < n; i++ {
		lw := list.NewWriter()
		lw.AppendItem(listItem1)
		lw.Indent()
		lw.AppendItems(listItems2)
		lw.Indent()
		lw.AppendItems(listItems3)
		lw.Render()
	}
}

func main() {
	numRenders := 10000
	if len(os.Args) > 1 {
		var err error
		numRenders, err = strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Printf("Invalid Argument: '%s'\n", os.Args[2])
			os.Exit(1)
		}
	}

	for _, profiler := range profilers {
		profileRender(profiler, numRenders)
	}
}
