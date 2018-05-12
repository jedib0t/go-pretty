package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList_RenderMarkdown(t *testing.T) {
	lw := NewWriter()
	lw.AppendItem(testItem1)
	lw.Indent()
	lw.AppendItems(testItems2)
	lw.Indent()
	lw.AppendItems(testItems3)
	lw.UnIndent()
	lw.AppendItem(testItem4)
	lw.Indent()
	lw.AppendItem(testItem5)

	expectedOutMarkdown := `  * Game Of Thrones
    * Winter
    * Is
    * Coming
      * This
      * Is
      * Known
    * The Dark Tower
      * The Gunslinger`
	assert.Equal(t, expectedOutMarkdown, lw.RenderMarkdown())

	lw.SetStyle(styleTest)
	assert.NotNil(t, lw.Style())
	assert.Equal(t, styleTest.Name, lw.Style().Name)
	assert.Equal(t, expectedOutMarkdown, lw.RenderMarkdown())
	assert.NotNil(t, lw.Style())
	assert.Equal(t, styleTest.Name, lw.Style().Name)
}
