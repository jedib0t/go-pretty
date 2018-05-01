package text

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVAlign_Apply(t *testing.T) {
	assert.Equal(t, []string{"Game", "Of", "Thrones"},
		VAlignDefault.Apply([]string{"Game", "Of", "Thrones"}, 3))

	assert.Equal(t, []string{"Game", "Of", "Thrones", "", ""},
		VAlignDefault.Apply([]string{"Game", "Of", "Thrones"}, 5))
	assert.Equal(t, []string{"Game", "Of", "Thrones", "", ""},
		VAlignTop.Apply([]string{"Game", "Of", "Thrones"}, 5))
	assert.Equal(t, []string{"", "Game", "Of", "Thrones", ""},
		VAlignMiddle.Apply([]string{"Game", "Of", "Thrones"}, 5))
	assert.Equal(t, []string{"", "", "Game", "Of", "Thrones"},
		VAlignBottom.Apply([]string{"Game", "Of", "Thrones"}, 5))
}

func TestVAlign_ApplyStr(t *testing.T) {
	assert.Equal(t, []string{"Game", "Of", "Thrones", "", ""},
		VAlignDefault.ApplyStr("Game\nOf\nThrones", 5))
	assert.Equal(t, []string{"Game", "Of", "Thrones", "", ""},
		VAlignTop.ApplyStr("Game\nOf\nThrones", 5))
	assert.Equal(t, []string{"", "Game", "Of", "Thrones", ""},
		VAlignMiddle.ApplyStr("Game\nOf\nThrones", 5))
	assert.Equal(t, []string{"", "", "Game", "Of", "Thrones"},
		VAlignBottom.ApplyStr("Game\nOf\nThrones", 5))
}

func TestVAlign_HTMLProperty(t *testing.T) {
	vAligns := map[VAlign]string{
		VAlignDefault: "",
		VAlignTop:     "top",
		VAlignMiddle:  "middle",
		VAlignBottom:  "bottom",
	}
	for vAlign, htmlStyle := range vAligns {
		assert.Contains(t, vAlign.HTMLProperty(), htmlStyle)
	}
}
