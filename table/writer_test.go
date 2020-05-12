package table

import (
	"fmt"

	"github.com/jedib0t/go-pretty/v6/text"
)

func Example_simple() {
	// simple table with zero customizations
	tw := NewWriter()
	// append a header row
	tw.AppendHeader(Row{"#", "First Name", "Last Name", "Salary"})
	// append some data rows
	tw.AppendRows([]Row{
		{1, "Arya", "Stark", 3000},
		{20, "Jon", "Snow", 2000, "You know nothing, Jon Snow!"},
		{300, "Tyrion", "Lannister", 5000},
	})
	// append a footer row
	tw.AppendFooter(Row{"", "", "Total", 10000})
	// render it
	fmt.Printf("Table without any customizations:\n%s", tw.Render())

	// Output: Table without any customizations:
	// +-----+------------+-----------+--------+-----------------------------+
	// |   # | FIRST NAME | LAST NAME | SALARY |                             |
	// +-----+------------+-----------+--------+-----------------------------+
	// |   1 | Arya       | Stark     |   3000 |                             |
	// |  20 | Jon        | Snow      |   2000 | You know nothing, Jon Snow! |
	// | 300 | Tyrion     | Lannister |   5000 |                             |
	// +-----+------------+-----------+--------+-----------------------------+
	// |     |            | TOTAL     |  10000 |                             |
	// +-----+------------+-----------+--------+-----------------------------+
}

func Example_styled() {
	// table with some amount of customization
	tw := NewWriter()
	// append a header row
	tw.AppendHeader(Row{"First Name", "Last Name", "Salary"})
	// append some data rows
	tw.AppendRows([]Row{
		{"Jaime", "Lannister", 5000},
		{"Arya", "Stark", 3000, "A girl has no name."},
		{"Sansa", "Stark", 4000},
		{"Jon", "Snow", 2000, "You know nothing, Jon Snow!"},
		{"Tyrion", "Lannister", 5000, "A Lannister always pays his debts."},
	})
	// append a footer row
	tw.AppendFooter(Row{"", "Total", 10000})
	// auto-index rows
	tw.SetAutoIndex(true)
	// sort by last name and then by salary
	tw.SortBy([]SortBy{{Name: "Last Name", Mode: Dsc}, {Name: "Salary", Mode: AscNumeric}})
	// use a ready-to-use style
	tw.SetStyle(StyleLight)
	// customize the style and change some stuff
	tw.Style().Format.Header = text.FormatLower
	tw.Style().Format.Row = text.FormatLower
	tw.Style().Format.Footer = text.FormatLower
	tw.Style().Options.SeparateColumns = false
	// render it
	fmt.Printf("Table with customizations:\n%s", tw.Render())

	// Output: Table with customizations:
	// ┌──────────────────────────────────────────────────────────────────────┐
	// │    first name  last name  salary                                     │
	// ├──────────────────────────────────────────────────────────────────────┤
	// │ 1  arya        stark        3000  a girl has no name.                │
	// │ 2  sansa       stark        4000                                     │
	// │ 3  jon         snow         2000  you know nothing, jon snow!        │
	// │ 4  jaime       lannister    5000                                     │
	// │ 5  tyrion      lannister    5000  a lannister always pays his debts. │
	// ├──────────────────────────────────────────────────────────────────────┤
	// │                total       10000                                     │
	// └──────────────────────────────────────────────────────────────────────┘
}
