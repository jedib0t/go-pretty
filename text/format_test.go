package text

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormat_Apply(t *testing.T) {
	assert.Equal(t, "jon Snow", FormatDefault.Apply("jon Snow"))
	assert.Equal(t, "jon snow", FormatLower.Apply("jon Snow"))
	assert.Equal(t, "Jon Snow", FormatTitle.Apply("jon Snow"))
	assert.Equal(t, "JON SNOW", FormatUpper.Apply("jon Snow"))
}
