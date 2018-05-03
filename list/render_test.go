package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
