package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList_RenderHTML(t *testing.T) {
	lw := NewWriter()
	lw.AppendItem(testItem1)
	lw.Indent()
	lw.AppendItems(testItems2)
	lw.Indent()
	lw.AppendItems(testItems3)
	lw.UnIndent()
	lw.UnIndent()
	lw.AppendItem(testItem4)
	lw.Indent()
	lw.AppendItem(testItem5)
	lw.SetHTMLCSSClass(testCSSClass)

	expectedOut := `<ul class="test-css-class">
  <li>Game Of Thrones</li>
  <ul class="test-css-class-1">
    <li>Winter</li>
    <li>Is</li>
    <li>Coming</li>
    <ul class="test-css-class-2">
      <li>This</li>
      <li>Is</li>
      <li>Known</li>
    </ul>
  </ul>
  <li>The Dark Tower</li>
  <ul class="test-css-class-1">
    <li>The Gunslinger</li>
  </ul>
</ul>`

	assert.Equal(t, expectedOut, lw.RenderHTML())
}

func TestList_RenderHTML_Complex(t *testing.T) {
	lw := NewWriter()
	lw.AppendItem("The Houses of Westeros")
	lw.Indent()
	lw.AppendItem("The Starks of Winterfell")
	lw.Indent()
	lw.AppendItem("Eddard Stark")
	lw.Indent()
	lw.AppendItems([]interface{}{"Robb Stark", "Sansa Stark", "Arya Stark", "Bran Stark", "Rickon Stark"})
	lw.UnIndent()
	lw.AppendItems([]interface{}{"Lyanna Stark", "Benjen Stark"})
	lw.UnIndent()
	lw.AppendItem("The Targaryens of Dragonstone")
	lw.Indent()
	lw.AppendItem("Aerys Targaryen")
	lw.Indent()
	lw.AppendItems([]interface{}{"Rhaegar Targaryen", "Viserys Targaryen", "Daenerys Targaryen"})
	lw.UnIndent()
	lw.UnIndent()
	lw.AppendItem("The Lannisters of Lannisport")
	lw.Indent()
	lw.AppendItem("Tywin Lannister")
	lw.Indent()
	lw.AppendItems([]interface{}{"Cersei Lannister", "Jaime Lannister", "Tyrion Lannister"})

	expectedOut := `<ul class="go-pretty-table">
  <li>The Houses of Westeros</li>
  <ul class="go-pretty-table-1">
    <li>The Starks of Winterfell</li>
    <ul class="go-pretty-table-2">
      <li>Eddard Stark</li>
      <ul class="go-pretty-table-3">
        <li>Robb Stark</li>
        <li>Sansa Stark</li>
        <li>Arya Stark</li>
        <li>Bran Stark</li>
        <li>Rickon Stark</li>
      </ul>
      <li>Lyanna Stark</li>
      <li>Benjen Stark</li>
    </ul>
    <li>The Targaryens of Dragonstone</li>
    <ul class="go-pretty-table-2">
      <li>Aerys Targaryen</li>
      <ul class="go-pretty-table-3">
        <li>Rhaegar Targaryen</li>
        <li>Viserys Targaryen</li>
        <li>Daenerys Targaryen</li>
      </ul>
    </ul>
    <li>The Lannisters of Lannisport</li>
    <ul class="go-pretty-table-2">
      <li>Tywin Lannister</li>
      <ul class="go-pretty-table-3">
        <li>Cersei Lannister</li>
        <li>Jaime Lannister</li>
        <li>Tyrion Lannister</li>
      </ul>
    </ul>
  </ul>
</ul>`

	assert.Equal(t, expectedOut, lw.RenderHTML())
}
