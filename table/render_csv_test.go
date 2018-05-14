package table

import (
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

	expectedOut := `#,First Name,Last Name,Salary,
1,Arya,Stark,3000,
20,Jon,Snow,2000,"You know nothing\, Jon Snow!"
300,Tyrion,Lannister,5000,
0,Winter,Is,0,"Coming.
The North Remembers!
This is known."
0,Valar,Morghulis,0,Faceless    Men
,,Total,10000,`

	assert.Equal(t, expectedOut, tw.RenderCSV())
}

func TestTable_RenderCSV_Empty(t *testing.T) {
	tw := NewWriter()
	assert.Empty(t, tw.RenderCSV())
}
