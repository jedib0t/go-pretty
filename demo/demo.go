package main

import (
	"fmt"

	"github.com/jedib0t/go-pretty/list"
	"github.com/jedib0t/go-pretty/table"
	"github.com/jedib0t/go-pretty/text"
)

// demoList demonstrates the capabilities of Writer with some of the
// available ready-to-use styles.
func demoList() string {
	styles := []list.Style{
		list.StyleDefault,
		list.StyleBulletCircle,
		list.StyleConnectedLight,
	}
	itemLevel1 := "Game Of Thrones"
	itemsLevel2 := []interface{}{"Winter", "Is", "Coming"}
	itemsLevel3 := []interface{}{"This", "Is", "Known"}

	var header, content []interface{}
	for _, style := range styles {
		lw := list.NewWriter()
		lw.AppendItem(itemLevel1)
		lw.Indent()
		lw.AppendItems(itemsLevel2)
		lw.Indent()
		lw.AppendItems(itemsLevel3)
		lw.SetStyle(style)

		header = append(header, style.Name)
		content = append(content, lw.Render())
	}

	tw := table.NewWriter()
	tw.AppendHeader(header)
	tw.AppendRow(content)
	tw.SetStyle(table.StyleLight)
	tw.Style().FormatHeader = text.FormatDefault
	return tw.Render()
}

// demoTable demonstrates the capabilities of Writer with some of the
// available ready-to-use styles.
func demoTable() string {
	styles := []table.Style{
		table.StyleDefault,
		table.StyleLight,
		table.StyleBold,
		table.StyleDouble,
		table.StyleRounded,
	}
	header := table.Row{"#", "First Name", "Last Name", "Salary"}
	rows1And2 := []table.Row{
		{1, "Arya", "Stark", 3000},
		{20, "Jon", "Snow", 2000, "You know nothing, Jon Snow!"},
	}
	row3 := table.Row{300, "Tyrion", "Lannister", 5000}
	footer := table.Row{"", "", "Total", 10000}
	align := []text.Align{
		text.AlignDefault,
		text.AlignLeft,
		text.AlignLeft,
		text.AlignRight,
	}

	var content []table.Row
	for _, style := range styles {
		tw := table.NewWriter()
		tw.AppendHeader(header)
		tw.AppendRows(rows1And2)
		tw.AppendRow(row3)
		tw.AppendFooter(footer)
		tw.SetAlign(align)
		tw.SetStyle(style)

		if len(content) == 0 {
			content = append(content, table.Row{"As CSV", tw.RenderCSV()})
		}
		content = append(content, table.Row{style.Name, tw.Render()})
	}

	tw := table.NewWriter()
	tw.AppendRows(content)
	tw.ShowSeparators(true)
	tw.SetAlign([]text.Align{text.AlignRight, text.AlignLeft})
	tw.SetStyle(table.StyleLight)
	tw.SetVAlign([]text.VAlign{text.VAlignMiddle, text.VAlignDefault})
	return tw.Render()
}

func main() {
	demo := table.NewWriter()
	demo.AppendHeader([]interface{}{"#", "Feature", "Samples"})
	demo.AppendRow(table.Row{"1", "List", demoList()})
	demo.AppendRow(table.Row{"2", "Table", demoTable()})
	demo.SetAlign([]text.Align{
		text.AlignDefault,
		text.AlignLeft,
		text.AlignCenter,
	})
	demo.SetCaption("Generated with go-pretty; MIT License; Copyright (c) 2018 jedib0t.")
	demo.SetStyle(table.StyleDouble)
	fmt.Println(demo.Render())
}
