package main

import (
	"fmt"

	"github.com/jedib0t/go-pretty"
)

// demoList demonstrates the capabilities of ListWriter with some of the
// available ready-to-use styles.
func demoList() string {
	styles := []gopretty.ListStyle{
		gopretty.ListStyleDefault,
		gopretty.ListStyleBulletCircle,
		gopretty.ListStyleConnectedLight,
	}
	itemLevel1 := "Game Of Thrones"
	itemsLevel2 := []interface{}{"Winter", "Is", "Coming"}
	itemsLevel3 := []interface{}{"This", "Is", "Known"}

	var header, content []interface{}
	for _, style := range styles {
		lw := gopretty.NewListWriter()
		lw.AppendItem(itemLevel1)
		lw.Indent()
		lw.AppendItems(itemsLevel2)
		lw.Indent()
		lw.AppendItems(itemsLevel3)
		lw.SetStyle(style)

		header = append(header, style.Name)
		content = append(content, lw.Render())
	}

	tw := gopretty.NewTableWriter()
	tw.AppendHeader(header)
	tw.AppendRow(content)
	tw.SetStyle(gopretty.TableStyleLight)
	tw.Style().CaseHeader = gopretty.TextCaseDefault
	return tw.Render()
}

// demoTable demonstrates the capabilities of TableWriter with some of the
// available ready-to-use styles.
func demoTable() string {
	styles := []gopretty.TableStyle{
		gopretty.TableStyleDefault,
		gopretty.TableStyleLight,
		gopretty.TableStyleBold,
		gopretty.TableStyleDouble,
		gopretty.TableStyleRounded,
	}
	header := gopretty.TableRow{"#", "First Name", "Last Name", "Salary"}
	rows1And2 := []gopretty.TableRow{
		{1, "Arya", "Stark", 3000},
		{20, "Jon", "Snow", 2000, "You know nothing, Jon Snow!"},
	}
	row3 := gopretty.TableRow{300, "Tyrion", "Lannister", 5000}
	footer := gopretty.TableRow{"", "", "Total", 10000}
	align := []gopretty.Align{
		gopretty.AlignDefault,
		gopretty.AlignLeft,
		gopretty.AlignLeft,
		gopretty.AlignRight,
	}

	var content []gopretty.TableRow
	for _, style := range styles {
		table := gopretty.NewTableWriter()
		table.AppendHeader(header)
		table.AppendRows(rows1And2)
		table.AppendRow(row3)
		table.AppendFooter(footer)
		table.SetAlign(align)
		table.SetStyle(style)

		if len(content) == 0 {
			content = append(content, gopretty.TableRow{"As CSV", table.RenderCSV()})
		}
		content = append(content, gopretty.TableRow{style.Name, table.Render()})
	}

	tw := gopretty.NewTableWriter()
	tw.AppendRows(content)
	tw.EnableSeparators()
	tw.SetAlign([]gopretty.Align{gopretty.AlignRight, gopretty.AlignLeft})
	tw.SetStyle(gopretty.TableStyleLight)
	tw.SetVAlign([]gopretty.VAlign{gopretty.VAlignMiddle, gopretty.VAlignDefault})
	return tw.Render()
}

func main() {
	demo := gopretty.NewTableWriter()
	demo.AppendHeader([]interface{}{"#", "Feature", "Samples"})
	demo.AppendRow(gopretty.TableRow{"1", "List", demoList()})
	demo.AppendRow(gopretty.TableRow{"2", "Table", demoTable()})
	demo.SetAlign([]gopretty.Align{
		gopretty.AlignDefault,
		gopretty.AlignLeft,
		gopretty.AlignCenter,
	})
	demo.SetCaption("Generated with go-pretty; MIT License; Copyright (c) 2018 jedib0t.")
	demo.SetStyle(gopretty.TableStyleDouble)
	fmt.Println(demo.Render())
}
