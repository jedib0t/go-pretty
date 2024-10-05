package main

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"golang.org/x/term"
)

func main() {
	w, h, err := term.GetSize(int(os.Stdin.Fd()))
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(w, h)

	tw := table.NewWriter()
	tw.SetTitle("Title")
	tw.AppendHeader(table.Row{"#", "First Name", "Last Name", "Salary"})
	tw.AppendRows([]table.Row{
		{1, "Arya", "Stark", 3000},
		{20, "Jon", "Snow", 2000, "You know nothing, Jon Snow!"},
		{300, "Tyrion", "Lannister", 5000},
	})
	tw.AppendFooter(table.Row{"", "", "Total", 10000})
	tw.SetStyle(table.StyleLight)
	tw.Style().Size = table.SizeOptions{
		WidthMin: w,
	}
	fmt.Println(tw.Render())
}
