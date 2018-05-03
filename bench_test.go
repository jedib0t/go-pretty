package gopretty

import (
	"testing"

	"github.com/jedib0t/go-pretty/list"
	"github.com/jedib0t/go-pretty/table"
	"github.com/jedib0t/go-pretty/text"
)

var (
	listItem1      = "Game Of Thrones"
	listItems2     = []interface{}{"Winter", "Is", "Coming"}
	listItems3     = []interface{}{"This", "Is", "Known"}
	tableRowAlign  = []text.Align{text.AlignDefault, text.AlignLeft, text.AlignLeft, text.AlignRight}
	tableCaption   = "table-caption"
	tableRowFooter = table.Row{"", "", "Total", 10000}
	tableRowHeader = table.Row{"#", "First Name", "Last Name", "Salary"}
	tableRows      = []table.Row{
		{1, "Arya", "Stark", 3000},
		{20, "Jon", "Snow", 2000, "You know nothing, Jon Snow!"},
		{300, "Tyrion", "Lannister", 5000},
	}
)

func generateBenchmarkTable() table.Writer {
	tw := table.NewWriter()
	tw.AppendHeader(tableRowHeader)
	tw.AppendRows(tableRows)
	tw.AppendFooter(tableRowFooter)
	tw.SetAlign(tableRowAlign)
	tw.SetCaption(tableCaption)
	return tw
}

func BenchmarkList_Render(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lw := list.NewWriter()
		lw.AppendItem(listItem1)
		lw.Indent()
		lw.AppendItems(listItems2)
		lw.Indent()
		lw.AppendItems(listItems3)
		lw.Render()
	}
}

func BenchmarkTable_Render(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generateBenchmarkTable().Render()
	}
}

func BenchmarkTable_RenderCSV(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generateBenchmarkTable().RenderCSV()
	}
}

func BenchmarkTable_RenderHTML(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generateBenchmarkTable().RenderHTML()
	}
}

func BenchmarkTable_RenderMarkdown(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generateBenchmarkTable().RenderMarkdown()
	}
}
