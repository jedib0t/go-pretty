package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty"
)

var (
	tableCaptionColor = gopretty.TextColor{color.FgHiYellow}
)

func showTable() {
	tw := gopretty.NewTableWriter()
	tw.AppendHeader(gopretty.TableRow{"#", "First Name", "Last Name", "Salary"})
	tw.AppendRows([]gopretty.TableRow{
		{1, "Arya", "Stark", 3000},
		{20, "Jon", "Snow", 2000, "You know nothing, Jon Snow!"},
		{300, "Tyrion", "Lannister", 5000},
	})
	tw.AppendFooter(gopretty.TableRow{"", "", "Total", 10000})
	tw.SetCaption(tableCaptionColor.Sprintf("A Simple Table.\n"))
	tw.SetStyle(gopretty.TableStyleLight)

	fmt.Println(tw.Render())
}

func showTableColored() {
	tw := gopretty.NewTableWriter()
	tw.AppendHeader(gopretty.TableRow{"#", "First Name", "Last Name", "Salary"})
	tw.AppendRows([]gopretty.TableRow{
		{1, "Arya", "Stark", 3000},
		{20, "Jon", "Snow", 2000, "You know nothing, Jon Snow!"},
		{300, "Tyrion", "Lannister", 5000},
	})
	tw.AppendFooter(gopretty.TableRow{"", "", "Total", 10000})
	tw.SetCaption(tableCaptionColor.Sprintf("A Colorized Table.\n"))
	tw.SetStyle(gopretty.TableStyleDouble)

	colorRow := gopretty.TextColor{color.FgGreen}
	colorRowNotes := gopretty.TextColor{color.FgCyan}
	colorRowHeader := gopretty.TextColor{color.FgHiRed, color.Bold}
	colorRowFooter := gopretty.TextColor{color.FgHiBlue, color.Bold}
	tw.SetColors([]gopretty.TextColor{colorRow, colorRow, colorRow, colorRow, colorRowNotes})
	tw.SetColorsFooter([]gopretty.TextColor{{}, {}, colorRowFooter, colorRowFooter})
	tw.SetColorsHeader([]gopretty.TextColor{colorRowHeader, colorRowHeader, colorRowHeader, colorRowHeader})

	fmt.Println(tw.Render())
}

func main() {
	showTable()
	showTableColored()
}
