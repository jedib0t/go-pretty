package table

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTable_RenderCSV(t *testing.T) {
	tw := NewWriter()
	tw.AppendHeader(testHeader)
	tw.AppendRows(testRows)
	tw.AppendRow(testRowMultiLine)
	tw.AppendRow(testRowTabs)
	tw.AppendFooter(testFooter)
	tw.SetCaption(testCaption)
	tw.SetTitle(testTitle1)

	expectedOut := `Game of Thrones
#,First Name,Last Name,Salary,
1,Arya,Stark,3000,
20,Jon,Snow,2000,"You know nothing\, Jon Snow!"
300,Tyrion,Lannister,5000,
0,Winter,Is,0,"Coming.
The North Remembers!
This is known."
0,Valar,Morghulis,0,Faceless    Men
,,Total,10000,
A Song of Ice and Fire`
	assert.Equal(t, expectedOut, tw.RenderCSV())
}

func TestTable_RenderCSV_AutoIndex(t *testing.T) {
	tw := NewWriter()
	for rowIdx := 0; rowIdx < 10; rowIdx++ {
		row := make(Row, 10)
		for colIdx := 0; colIdx < 10; colIdx++ {
			row[colIdx] = fmt.Sprintf("%s%d", AutoIndexColumnID(colIdx), rowIdx+1)
		}
		tw.AppendRow(row)
	}
	for rowIdx := 0; rowIdx < 1; rowIdx++ {
		row := make(Row, 10)
		for colIdx := 0; colIdx < 10; colIdx++ {
			row[colIdx] = AutoIndexColumnID(colIdx) + "F"
		}
		tw.AppendFooter(row)
	}
	tw.SetAutoIndex(true)
	tw.SetStyle(StyleLight)

	expectedOut := `,A,B,C,D,E,F,G,H,I,J
1,A1,B1,C1,D1,E1,F1,G1,H1,I1,J1
2,A2,B2,C2,D2,E2,F2,G2,H2,I2,J2
3,A3,B3,C3,D3,E3,F3,G3,H3,I3,J3
4,A4,B4,C4,D4,E4,F4,G4,H4,I4,J4
5,A5,B5,C5,D5,E5,F5,G5,H5,I5,J5
6,A6,B6,C6,D6,E6,F6,G6,H6,I6,J6
7,A7,B7,C7,D7,E7,F7,G7,H7,I7,J7
8,A8,B8,C8,D8,E8,F8,G8,H8,I8,J8
9,A9,B9,C9,D9,E9,F9,G9,H9,I9,J9
10,A10,B10,C10,D10,E10,F10,G10,H10,I10,J10
,AF,BF,CF,DF,EF,FF,GF,HF,IF,JF`
	assert.Equal(t, expectedOut, tw.RenderCSV())
}

func TestTable_RenderCSV_Empty(t *testing.T) {
	tw := NewWriter()
	assert.Empty(t, tw.RenderCSV())
}

func TestTable_RenderCSV_HiddenColumns(t *testing.T) {
	tw := NewWriter()
	tw.AppendHeader(testHeader)
	tw.AppendRows(testRows)
	tw.AppendFooter(testFooter)

	// ensure sorting is done before hiding the columns
	tw.SortBy([]SortBy{
		{Name: "Salary", Mode: DscNumeric},
	})

	t.Run("every column hidden", func(t *testing.T) {
		tw.SetColumnConfigs(generateColumnConfigsWithHiddenColumns([]int{0, 1, 2, 3, 4}))

		expectedOut := ``
		assert.Equal(t, expectedOut, tw.RenderCSV())
	})

	t.Run("first column hidden", func(t *testing.T) {
		tw.SetColumnConfigs(generateColumnConfigsWithHiddenColumns([]int{0}))

		expectedOut := `First Name,Last Name,Salary,
>>Tyrion,Lannister<<,5013,
>>Arya,Stark<<,3013,
>>Jon,Snow<<,2013,"~You know nothing\, Jon Snow!~"
,Total,10000,`
		assert.Equal(t, expectedOut, tw.RenderCSV())
	})

	t.Run("column hidden in the middle", func(t *testing.T) {
		tw.SetColumnConfigs(generateColumnConfigsWithHiddenColumns([]int{1}))

		expectedOut := `#,Last Name,Salary,
307,Lannister<<,5013,
8,Stark<<,3013,
27,Snow<<,2013,"~You know nothing\, Jon Snow!~"
,Total,10000,`
		assert.Equal(t, expectedOut, tw.RenderCSV())
	})

	t.Run("last column hidden", func(t *testing.T) {
		tw.SetColumnConfigs(generateColumnConfigsWithHiddenColumns([]int{4}))

		expectedOut := `#,First Name,Last Name,Salary
307,>>Tyrion,Lannister<<,5013
8,>>Arya,Stark<<,3013
27,>>Jon,Snow<<,2013
,,Total,10000`
		assert.Equal(t, expectedOut, tw.RenderCSV())
	})
}

func TestTable_RenderCSV_Sorted(t *testing.T) {
	tw := NewWriter()
	tw.AppendHeader(testHeader)
	tw.AppendRows(testRows)
	tw.AppendRow(Row{11, "Sansa", "Stark", 6000})
	tw.AppendFooter(testFooter)
	tw.SortBy([]SortBy{{Name: "Last Name", Mode: Asc}, {Name: "First Name", Mode: Asc}})

	expectedOut := `#,First Name,Last Name,Salary,
300,Tyrion,Lannister,5000,
20,Jon,Snow,2000,"You know nothing\, Jon Snow!"
1,Arya,Stark,3000,
11,Sansa,Stark,6000,
,,Total,10000,`
	assert.Equal(t, expectedOut, tw.RenderCSV())
}
