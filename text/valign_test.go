package text

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleVAlign_Apply() {
	lines := []string{"Game", "Of", "Thrones"}
	maxLines := 5
	fmt.Printf("VAlignDefault: %#v\n", VAlignDefault.Apply(lines, maxLines))
	fmt.Printf("VAlignTop    : %#v\n", VAlignTop.Apply(lines, maxLines))
	fmt.Printf("VAlignMiddle : %#v\n", VAlignMiddle.Apply(lines, maxLines))
	fmt.Printf("VAlignBottom : %#v\n", VAlignBottom.Apply(lines, maxLines))

	// Output: VAlignDefault: []string{"Game", "Of", "Thrones", "", ""}
	// VAlignTop    : []string{"Game", "Of", "Thrones", "", ""}
	// VAlignMiddle : []string{"", "Game", "Of", "Thrones", ""}
	// VAlignBottom : []string{"", "", "Game", "Of", "Thrones"}
}

func TestVAlign_Apply(t *testing.T) {
	assert.Equal(t, []string{"Game", "Of", "Thrones"},
		VAlignDefault.Apply([]string{"Game", "Of", "Thrones"}, 1))
	assert.Equal(t, []string{"Game", "Of", "Thrones"},
		VAlignDefault.Apply([]string{"Game", "Of", "Thrones"}, 3))
	assert.Equal(t, []string{"Game", "Of", "Thrones", "", ""},
		VAlignDefault.Apply([]string{"Game", "Of", "Thrones"}, 5))

	assert.Equal(t, []string{"Game", "Of", "Thrones"},
		VAlignTop.Apply([]string{"Game", "Of", "Thrones"}, 1))
	assert.Equal(t, []string{"Game", "Of", "Thrones"},
		VAlignTop.Apply([]string{"Game", "Of", "Thrones"}, 3))
	assert.Equal(t, []string{"Game", "Of", "Thrones", "", ""},
		VAlignTop.Apply([]string{"Game", "Of", "Thrones"}, 5))

	assert.Equal(t, []string{"Game", "Of", "Thrones"},
		VAlignMiddle.Apply([]string{"Game", "Of", "Thrones"}, 1))
	assert.Equal(t, []string{"Game", "Of", "Thrones"},
		VAlignMiddle.Apply([]string{"Game", "Of", "Thrones"}, 3))
	assert.Equal(t, []string{"", "Game", "Of", "Thrones", ""},
		VAlignMiddle.Apply([]string{"Game", "Of", "Thrones"}, 5))

	assert.Equal(t, []string{"Game", "Of", "Thrones"},
		VAlignBottom.Apply([]string{"Game", "Of", "Thrones"}, 1))
	assert.Equal(t, []string{"Game", "Of", "Thrones"},
		VAlignBottom.Apply([]string{"Game", "Of", "Thrones"}, 3))
	assert.Equal(t, []string{"", "", "Game", "Of", "Thrones"},
		VAlignBottom.Apply([]string{"Game", "Of", "Thrones"}, 5))
}

func ExampleVAlign_ApplyStr() {
	str := "Game\nOf\nThrones"
	maxLines := 5
	fmt.Printf("VAlignDefault: %#v\n", VAlignDefault.ApplyStr(str, maxLines))
	fmt.Printf("VAlignTop    : %#v\n", VAlignTop.ApplyStr(str, maxLines))
	fmt.Printf("VAlignMiddle : %#v\n", VAlignMiddle.ApplyStr(str, maxLines))
	fmt.Printf("VAlignBottom : %#v\n", VAlignBottom.ApplyStr(str, maxLines))

	// Output: VAlignDefault: []string{"Game", "Of", "Thrones", "", ""}
	// VAlignTop    : []string{"Game", "Of", "Thrones", "", ""}
	// VAlignMiddle : []string{"", "Game", "Of", "Thrones", ""}
	// VAlignBottom : []string{"", "", "Game", "Of", "Thrones"}
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

func ExampleVAlign_HTMLProperty() {
	fmt.Printf("VAlignDefault: '%s'\n", VAlignDefault.HTMLProperty())
	fmt.Printf("VAlignTop    : '%s'\n", VAlignTop.HTMLProperty())
	fmt.Printf("VAlignMiddle : '%s'\n", VAlignMiddle.HTMLProperty())
	fmt.Printf("VAlignBottom : '%s'\n", VAlignBottom.HTMLProperty())

	// Output: VAlignDefault: ''
	// VAlignTop    : 'valign="top"'
	// VAlignMiddle : 'valign="middle"'
	// VAlignBottom : 'valign="bottom"'
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
