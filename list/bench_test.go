package list

import (
	"testing"
)

func BenchmarkList_Render(b *testing.B) {
	items2 := []interface{}{"Winter", "Is", "Coming"}
	items3 := []interface{}{"This", "Is", "Known"}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		lw := NewWriter()
		lw.AppendItem("Game Of Thrones")
		lw.Indent()
		lw.AppendItems(items2)
		lw.Indent()
		lw.AppendItems(items3)
		lw.Render()
	}
}
