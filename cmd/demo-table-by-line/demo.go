package main

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
	"strings"
)

var (
	colTitleIndex     = "#"
	colTitleFirstName = "First Name"
	colTitleLastName  = "Last Name"
	colTitleSalary    = "Salary"
	colTitleQuote     = "Quote"
	rowHeader         = table.Row{colTitleIndex, colTitleFirstName, colTitleLastName, colTitleSalary, colTitleQuote}
)

func demoTableFeatures() {
	//==========================================================================
	// Initialization
	//==========================================================================
	t := table.NewWriter()
	// you can also instantiate the object directly
	tTemp := table.Table{}
	tTemp.Render() // just to avoid the compile error of not using the object

	t.AppendRow(table.Row{1, "Philip", "Fry", 856, "Shut up and take my money!"})
	t.AppendRow(table.Row{2, "Turanga", "Leela", 735, "I'm nobody's servant!"})
	t.AppendRow(table.Row{3, "Bender", "Rodriguez", 647, "Bite my shiny metal ass!"})
	t.AppendRow(table.Row{4, "Amy", "Wong", 392, "Yay, I'm useful!"})
	t.AppendRow(table.Row{5, "Hermes", "Conrad", 555, "Sweet three-toed sloth of ice planet Hoth!"})
	t.AppendRow(table.Row{6, "Professor", "Farnsworth", 1212, "Good news, everyone!"})
	t.AppendRow(table.Row{7, "Zapp", "Brannigan", 458, "If we can hit that bullseye, the rest of the dominoes will fall like a house of cards. Checkmate."})
	t.AppendRow(table.Row{8, "Kif", "Kroker", 702, "Sigh."})
	t.AppendRow(table.Row{9, "Dr. John A.", "Zoidberg", 819, "Why not Zoidberg?"})
	t.AppendRow(table.Row{10, "Nibbler", "Nibbler", 334, "You are the last hope of the universe."})
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(rowHeader)
	t.AppendFooter(rowHeader)
	t.SetBatchSize(4)
	t.Render()
	//+---+-----+------------+-----------+--------+-----------------------------+
	//|   |   # | FIRST NAME | LAST NAME | SALARY |                             |
	//+---+-----+------------+-----------+--------+-----------------------------+
	//| 1 |   1 | Arya       | Stark     |   3000 |                             |
	//| 2 |  20 | Jon        | Snow      |   2000 | You know nothing, Jon Snow! |
	//| 3 | 300 | Tyrion     | Lannister |   5000 |                             |
	//+---+-----+------------+-----------+--------+-----------------------------+
	//==========================================================================

}

func main() {
	demoWhat := "features"
	if len(os.Args) > 1 {
		demoWhat = os.Args[1]
	}

	switch strings.ToLower(demoWhat) {
	default:
		demoTableFeatures()
	}
}
