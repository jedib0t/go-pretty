package gopretty

import (
	"testing"

	"github.com/fatih/color"
	"github.com/stretchr/testify/assert"
)

func TestAlign_Apply(t *testing.T) {
	assert.Equal(t, "Ghost  ", AlignDefault.Apply("Ghost", 7))
	assert.Equal(t, "Ghost  ", AlignLeft.Apply("Ghost", 7))
	assert.Equal(t, " Ghost ", AlignCenter.Apply("Ghost", 7))
	assert.Equal(t, "  Ghost", AlignRight.Apply("Ghost", 7))
}

func TestAlign_HTMLProperty(t *testing.T) {
	aligns := map[Align]string{
		AlignDefault: "",
		AlignLeft:    "left",
		AlignCenter:  "center",
		AlignRight:   "right",
	}
	for align, htmlStyle := range aligns {
		assert.Contains(t, align.HTMLProperty(), htmlStyle)
	}
}

func TestTextCase_Convert(t *testing.T) {
	assert.Equal(t, "jon Snow", TextCaseDefault.Convert("jon Snow"))
	assert.Equal(t, "jon snow", TextCaseLower.Convert("jon Snow"))
	assert.Equal(t, "Jon Snow", TextCaseTitle.Convert("jon Snow"))
	assert.Equal(t, "JON SNOW", TextCaseUpper.Convert("jon Snow"))
}

func TestTextColor_getColorizer(t *testing.T) {
	assert.Nil(t, TextColor{}.getColorizer())
	assert.NotNil(t, TextColor{color.BgBlack, color.FgWhite}.getColorizer())
}

func TestTextColor_Sprint(t *testing.T) {
	assert.Equal(t, "test true", TextColor{}.Sprint("test ", true))
	assert.Equal(t, "\x1b[31mtest true\x1b[0m", TextColor{color.FgRed}.Sprint("test ", true))
}

func TestTextColor_Sprintf(t *testing.T) {
	assert.Equal(t, "test true", TextColor{}.Sprintf("test %s", "true"))
	assert.Equal(t, "\x1b[31mtest true\x1b[0m", TextColor{color.FgRed}.Sprintf("test %s", "true"))
}

func TestVAlign_Apply(t *testing.T) {
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
