package table

import (
	"fmt"
	"testing"
)

var benchResultStr string

func benchGenerateTable(numRows int, numCols int) Writer {
	header := make(Row, numCols)
	for colIdx := range header {
		header[colIdx] = fmt.Sprintf("Column #%d", colIdx+1)
	}

	tw := NewWriter()
	tw.AppendHeader(header)
	for rowIdx := 0; rowIdx < numRows; rowIdx++ {
		row := make(Row, numCols)
		for colIdx := range row {
			row[colIdx] = fmt.Sprintf("value %d-%d", rowIdx+1, colIdx+1)
		}
		tw.AppendRow(row)
	}
	return tw
}

func BenchmarkTable_RenderLarge(b *testing.B) {
	tw := benchGenerateTable(1000, 10)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchResultStr = tw.Render()
	}
}

func BenchmarkTable_Filter_RegexMatch(b *testing.B) {
	tw := benchGenerateTable(1000, 3)
	tw.FilterBy([]FilterBy{
		{Number: 1, Operator: RegexMatch, Value: `^value [0-9]?[02468]0?-1$`},
	})

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchResultStr = tw.Render()
	}
}
