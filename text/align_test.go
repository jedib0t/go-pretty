package text

import (
	"testing"

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
