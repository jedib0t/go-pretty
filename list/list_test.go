package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	listItem1  = "Game Of Thrones"
	listItems2 = []interface{}{"Winter", "Is", "Coming"}
	listItems3 = []interface{}{"This", "Is", "Known"}
)

func BenchmarkList_Render(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lw := NewWriter()
		lw.AppendItem(listItem1)
		lw.Indent()
		lw.AppendItems(listItems2)
		lw.Indent()
		lw.AppendItems(listItems3)
		lw.Render()
	}
}

func TestNewWriter(t *testing.T) {
	lw := NewWriter()
	assert.Nil(t, lw.Style())

	lw.SetStyle(StyleConnectedBold)
	assert.NotNil(t, lw.Style())
	assert.Equal(t, StyleConnectedBold, *lw.Style())
}

func TestList_AppendItem(t *testing.T) {
	list := List{}
	assert.Equal(t, 0, list.Length())

	list.AppendItem(listItem1)
	list.AppendItem(listItem1)
	assert.Equal(t, 2, list.Length())
}

func TestList_AppendItems(t *testing.T) {
	list := List{}
	assert.Equal(t, 0, list.Length())

	list.AppendItems(listItems2)
	assert.Equal(t, len(listItems2), list.Length())
}

func TestList_Indent(t *testing.T) {
	list := List{}
	assert.Equal(t, 0, list.level)

	list.Indent()
	assert.Equal(t, 1, list.level)

	list.Indent()
	assert.Equal(t, 2, list.level)
}

func TestList_Render(t *testing.T) {
	lw := NewWriter()
	lw.AppendItem(listItem1)
	lw.Indent()
	lw.AppendItems(listItems2)
	lw.Indent()
	lw.AppendItems(listItems3)
	lw.SetStyle(styleTest)

	expectedOut := `^> Game Of Thrones
c<f> Winter
  i> Is
  i> Coming
  c<f> This
    i> Is
    v> Known`

	assert.Equal(t, expectedOut, lw.Render())
}

func TestList_RenderWithoutStyle(t *testing.T) {
	lw := NewWriter()
	lw.AppendItem(listItem1)
	lw.Indent()
	lw.AppendItems(listItems2)
	lw.Indent()
	lw.AppendItems(listItems3)

	expectedOut := `- Game Of Thrones
--- Winter
  - Is
  - Coming
  --- This
    - Is
    - Known`

	assert.Equal(t, expectedOut, lw.Render())
}

func TestList_SetStyle(t *testing.T) {
	list := List{}
	assert.Nil(t, list.Style())

	list.SetStyle(StyleDefault)
	assert.NotNil(t, list.Style())
	assert.Equal(t, &StyleDefault, list.Style())
}
