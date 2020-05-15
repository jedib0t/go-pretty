package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/pkg/profile"
)

var (
	profilers = []func(*profile.Profile){
		profile.CPUProfile,
		profile.MemProfileRate(512),
	}
	tableCaption   = "Profiling a Simple Table."
	tableRowFooter = table.Row{"", "", "Total", 10000}
	tableRowHeader = table.Row{"#", "First Name", "Last Name", "Salary"}
	tableRows      = []table.Row{
		{1, "Arya", "Stark", 3000},
		{20, "Jon", "Snow", 2000, "You know nothing, Jon Snow!"},
		{300, "Tyrion", "Lannister", 5000},
	}
)

func profileRender(profiler func(profile2 *profile.Profile), n int) {
	defer profile.Start(profiler, profile.ProfilePath(".")).Stop()

	for i := 0; i < n; i++ {
		tw := table.NewWriter()
		tw.AppendHeader(tableRowHeader)
		tw.AppendRows(tableRows)
		tw.AppendFooter(tableRowFooter)
		tw.SetCaption(tableCaption)
		tw.Render()
		tw.RenderCSV()
		tw.RenderHTML()
		tw.RenderMarkdown()
	}
}

func main() {
	numRenders := 100000
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
