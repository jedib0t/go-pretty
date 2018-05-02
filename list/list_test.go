package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testItem1  = "Game Of Thrones"
	testItems2 = []interface{}{"Winter", "Is", "Coming"}
	testItems3 = []interface{}{"This", "Is", "Known"}
)

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

	list.AppendItem(testItem1)
	list.AppendItem(testItem1)
	assert.Equal(t, 2, list.Length())
}

func TestList_AppendItems(t *testing.T) {
	list := List{}
	assert.Equal(t, 0, list.Length())

	list.AppendItems(testItems2)
	assert.Equal(t, len(testItems2), list.Length())
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
	lw.AppendItem(testItem1)
	lw.Indent()
	lw.AppendItems(testItems2)
	lw.Indent()
	lw.AppendItems(testItems3)
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
	lw.AppendItem(testItem1)
	lw.Indent()
	lw.AppendItems(testItems2)
	lw.Indent()
	lw.AppendItems(testItems3)

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
