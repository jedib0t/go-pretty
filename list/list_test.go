package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testCSSClass = "test-css-class"
	testItem1    = "Game Of Thrones"
	testItem1ML  = testItem1 + "\n\t// George. R. R. Martin"
	testItems2   = []interface{}{"Winter", "Is", "Coming"}
	testItems2ML = []interface{}{"Winter\r\nIs\nComing", "Is", "Coming"}
	testItems3   = []interface{}{"This", "Is", "Known"}
	testItems3ML = []interface{}{"This\nIs\nKnown", "Is", "Known"}
	testItem4    = "The Dark Tower"
	testItem4ML  = testItem4 + "\n\t// Stephen King"
	testItem5    = "The Gunslinger"
)

type myMockOutputMirror struct {
	mirroredOutput string
}

func (t *myMockOutputMirror) Write(p []byte) (n int, err error) {
	t.mirroredOutput += string(p)
	return len(p), nil
}

func TestNewWriter(t *testing.T) {
	lw := NewWriter()
	assert.NotNil(t, lw.Style())
	assert.Equal(t, StyleDefault, *lw.Style())

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

	// should not indent when there is no item in the list
	list.Indent()
	assert.Equal(t, 0, list.level)

	// should indent with an item in the list
	list.AppendItem(testItem1)
	list.Indent()
	assert.Equal(t, 1, list.level)

	// should not indent if the previous item will then become "2 levels below"
	list.Indent()
	assert.Equal(t, 1, list.level)
}

func TestList_Length(t *testing.T) {
	list := List{}
	assert.Equal(t, 0, list.Length())

	list.AppendItem(testItem1)
	assert.Equal(t, 1, list.Length())
}

func TestList_Reset(t *testing.T) {
	list := List{}
	list.SetStyle(StyleBulletCircle)
	assert.Equal(t, "", list.Render())

	list.AppendItem(testItem1)
	assert.Equal(t, "● Game Of Thrones", list.Render())

	list.Reset()
	assert.Equal(t, "", list.Render())
}

func TestList_SetHTMLCSSClass(t *testing.T) {
	list := List{}
	assert.Empty(t, list.htmlCSSClass)

	list.SetHTMLCSSClass(testCSSClass)
	assert.Equal(t, testCSSClass, list.htmlCSSClass)
}

func TestList_SetOutputMirror(t *testing.T) {
	list := List{}
	list.AppendItem(testItem1)
	expectedOut := "* Game Of Thrones"
	assert.Equal(t, nil, list.outputMirror)
	assert.Equal(t, expectedOut, list.Render())

	mockOutputMirror := &myMockOutputMirror{}
	list.SetOutputMirror(mockOutputMirror)
	assert.Equal(t, mockOutputMirror, list.outputMirror)
	assert.Equal(t, expectedOut, list.Render())
	assert.Equal(t, expectedOut+"\n", mockOutputMirror.mirroredOutput)
}

func TestList_SetStyle(t *testing.T) {
	list := List{}
	assert.NotNil(t, list.Style())
	list.AppendItem(testItem1)
	list.Indent()
	list.AppendItems(testItems2)
	expectedOut := `* Game Of Thrones
  * Winter
  * Is
  * Coming`
	assert.Equal(t, expectedOut, list.Render())

	list.SetStyle(StyleConnectedLight)
	assert.NotNil(t, list.Style())
	assert.Equal(t, &StyleConnectedLight, list.Style())
	expectedOut = `── Game Of Thrones
   ├─ Winter
   ├─ Is
   └─ Coming`
	assert.Equal(t, expectedOut, list.Render())
}

func TestList_UnIndent(t *testing.T) {
	list := List{level: 2}

	list.UnIndent()
	assert.Equal(t, 1, list.level)

	list.UnIndent()
	assert.Equal(t, 0, list.level)

	list.UnIndent()
	assert.Equal(t, 0, list.level)
}
