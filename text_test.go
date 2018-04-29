package gopretty

import (
	"testing"

	"github.com/fatih/color"
	"github.com/stretchr/testify/assert"
)

func TestAlignment_Apply(t *testing.T) {
	assert.Equal(t, "Ghost  ", AlignmentDefault.Apply("Ghost", 7))
	assert.Equal(t, "Ghost  ", AlignmentLeft.Apply("Ghost", 7))
	assert.Equal(t, " Ghost ", AlignmentCenter.Apply("Ghost", 7))
	assert.Equal(t, "  Ghost", AlignmentRight.Apply("Ghost", 7))
}

func TestAlignment_HTMLProperty(t *testing.T) {
	alignments := map[Alignment]string{
		AlignmentDefault: "",
		AlignmentLeft:    "left",
		AlignmentCenter:  "center",
		AlignmentRight:   "right",
	}
	for alignment, htmlStyle := range alignments {
		assert.Contains(t, alignment.HTMLProperty(), htmlStyle)
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

func TestVAlignment_Apply(t *testing.T) {
	assert.Equal(t, []string{"Game", "Of", "Thrones", "", ""},
		VAlignmentDefault.Apply([]string{"Game", "Of", "Thrones"}, 5))
	assert.Equal(t, []string{"Game", "Of", "Thrones", "", ""},
		VAlignmentTop.Apply([]string{"Game", "Of", "Thrones"}, 5))
	assert.Equal(t, []string{"", "Game", "Of", "Thrones", ""},
		VAlignmentMiddle.Apply([]string{"Game", "Of", "Thrones"}, 5))
	assert.Equal(t, []string{"", "", "Game", "Of", "Thrones"},
		VAlignmentBottom.Apply([]string{"Game", "Of", "Thrones"}, 5))
}

func TestVAlignment_ApplyStr(t *testing.T) {
	assert.Equal(t, []string{"Game", "Of", "Thrones", "", ""},
		VAlignmentDefault.ApplyStr("Game\nOf\nThrones", 5))
	assert.Equal(t, []string{"Game", "Of", "Thrones", "", ""},
		VAlignmentTop.ApplyStr("Game\nOf\nThrones", 5))
	assert.Equal(t, []string{"", "Game", "Of", "Thrones", ""},
		VAlignmentMiddle.ApplyStr("Game\nOf\nThrones", 5))
	assert.Equal(t, []string{"", "", "Game", "Of", "Thrones"},
		VAlignmentBottom.ApplyStr("Game\nOf\nThrones", 5))
}

func TestVAlignment_HTMLProperty(t *testing.T) {
	vAlignments := map[VAlignment]string{
		VAlignmentDefault: "",
		VAlignmentTop:     "top",
		VAlignmentMiddle:  "middle",
		VAlignmentBottom:  "bottom",
	}
	for vAlignment, htmlStyle := range vAlignments {
		assert.Contains(t, vAlignment.HTMLProperty(), htmlStyle)
	}
}
