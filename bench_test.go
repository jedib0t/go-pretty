package gopretty

import (
	"io/ioutil"
	"testing"
	"time"

	"github.com/jedib0t/go-pretty/v6/list"
	"github.com/jedib0t/go-pretty/v6/progress"
	"github.com/jedib0t/go-pretty/v6/table"
)

var (
	listItem1      = "Game Of Thrones"
	listItems2     = []interface{}{"Winter", "Is", "Coming"}
	listItems3     = []interface{}{"This", "Is", "Known"}
	tableCaption   = "table-caption"
	tableRowFooter = table.Row{"", "", "Total", 10000}
	tableRowHeader = table.Row{"#", "First Name", "Last Name", "Salary"}
	tableRows      = []table.Row{
		{1, "Arya", "Stark", 3000},
		{20, "Jon", "Snow", 2000, "You know nothing, Jon Snow!"},
		{300, "Tyrion", "Lannister", 5000},
	}
	tracker1 = progress.Tracker{Message: "Calculating Total   # 1", Total: 1000, Units: progress.UnitsDefault}
	tracker2 = progress.Tracker{Message: "Downloading File    # 2", Total: 1000, Units: progress.UnitsBytes}
	tracker3 = progress.Tracker{Message: "Transferring Amount # 3", Total: 1000, Units: progress.UnitsCurrencyDollar}
)

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

func BenchmarkProgress_Render(b *testing.B) {
	trackSomething := func(pw progress.Writer, tracker *progress.Tracker) {
		tracker.Reset()
		pw.AppendTracker(tracker)
		time.Sleep(time.Millisecond * 100)
		tracker.Increment(tracker.Total / 2)
		time.Sleep(time.Millisecond * 100)
		tracker.Increment(tracker.Total / 2)
	}

	for i := 0; i < b.N; i++ {
		pw := progress.NewWriter()
		pw.SetAutoStop(true)
		pw.SetOutputWriter(ioutil.Discard)
		go trackSomething(pw, &tracker1)
		go trackSomething(pw, &tracker2)
		go trackSomething(pw, &tracker3)
		time.Sleep(time.Millisecond * 50)
		pw.Render()
	}
}

func generateBenchmarkTable() table.Writer {
	tw := table.NewWriter()
	tw.AppendHeader(tableRowHeader)
	tw.AppendRows(tableRows)
	tw.AppendFooter(tableRowFooter)
	tw.SetCaption(tableCaption)
	return tw
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
