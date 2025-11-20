package main

import (
	"fmt"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

func main() {
	tw := table.NewWriter()
	tw.AppendRows([]table.Row{
		{renderGrid(false)},
		{renderGrid(true)},
	})
	tw.SetTitle("256-Color Palette")
	tw.SetStyle(table.StyleLight)
	tw.Style().Options.SeparateRows = true
	tw.Style().Title.Align = text.AlignCenter
	fmt.Println(tw.Render())
}

func alignCode(code int) string {
	return " " + text.AlignRight.Apply(fmt.Sprint(code), 3) + " "
}

func blankTable() table.Writer {
	tw := table.NewWriter()
	tw.SetStyle(table.StyleLight)
	style := tw.Style()
	style.Box.PaddingLeft = ""
	style.Box.PaddingRight = ""
	style.Options.DrawBorder = false
	style.Options.SeparateRows = false
	style.Options.SeparateColumns = false
	return tw
}

func buildRow(start, end int, isBackground bool) table.Row {
	row := make(table.Row, 0, end-start)
	for i := start; i < end; i++ {
		row = append(row, cellValue(i, isBackground))
	}
	return row
}

func cellValue(code int, isBackground bool) string {
	if isBackground {
		return text.Colors{text.Bg256Color(code), text.FgBlack}.Sprint(alignCode(code))
	}
	return text.Colors{text.BgBlack, text.Fg256Color(code)}.Sprint(alignCode(code))
}

func renderGrid(isBackground bool) string {
	tw := table.NewWriter()
	tw.SetIndexColumn(1)
	title := "Foreground Colors"
	if isBackground {
		title = "Background Colors"
	}
	tw.SetTitle(text.Underline.Sprint(title) + "\n")
	tw.SetStyle(table.StyleLight)
	style := tw.Style()
	style.Box.PaddingLeft = ""
	style.Box.PaddingRight = ""
	style.Options.DrawBorder = false
	style.Options.SeparateRows = false
	style.Options.SeparateColumns = false
	style.Title.Align = text.AlignCenter
	tw.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, Align: text.AlignCenter},
	})

	// Standard 16 colors (0-15)
	row16 := blankTable()
	row16.AppendRow(buildRow(0, 16, isBackground))
	tw.AppendRows([]table.Row{{row16.Render()}, {""}})

	// RGB cube colors (16-231) - 216 colors in 6 blocks of 36
	row216 := blankTable()
	blockRow := make(table.Row, 0)
	for block := 0; block < 6; block++ {
		blockStart := 16 + 36*block
		blockTable := blankTable()
		colors := buildRow(blockStart, blockStart+36, isBackground)
		for i := 0; i < len(colors); i += 6 {
			end := i + 6
			if end > len(colors) {
				end = len(colors)
			}
			blockTable.AppendRow(colors[i:end])
		}
		blockRow = append(blockRow, blockTable.Render())
		if len(blockRow) == 3 {
			row216.AppendRow(blockRow)
			blockRow = make(table.Row, 0)
		}
	}
	if len(blockRow) > 0 {
		row216.AppendRow(blockRow)
	}
	tw.AppendRows([]table.Row{{row216.Render()}, {""}})

	// Grayscale colors (232-255) - 24 colors
	rowGrayscale := blankTable()
	colors := buildRow(232, 256, isBackground)
	for i := 0; i < len(colors); i += 12 {
		end := i + 12
		if end > len(colors) {
			end = len(colors)
		}
		rowGrayscale.AppendRow(colors[i:end])
	}
	tw.AppendRows([]table.Row{{rowGrayscale.Render()}})

	return tw.Render()
}
