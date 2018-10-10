package text

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleFormat_Apply() {
	fmt.Printf("FormatDefault: '%s'\n", FormatDefault.Apply("jon Snow"))
	fmt.Printf("FormatLower  : '%s'\n", FormatLower.Apply("jon Snow"))
	fmt.Printf("FormatTitle  : '%s'\n", FormatTitle.Apply("jon Snow"))
	fmt.Printf("FormatUpper  : '%s'\n", FormatUpper.Apply("jon Snow"))

	// Output: FormatDefault: 'jon Snow'
	// FormatLower  : 'jon snow'
	// FormatTitle  : 'Jon Snow'
	// FormatUpper  : 'JON SNOW'
}

func TestFormat_Apply(t *testing.T) {
	assert.Equal(t, "jon Snow", FormatDefault.Apply("jon Snow"))
	assert.Equal(t, "jon snow", FormatLower.Apply("jon Snow"))
	assert.Equal(t, "Jon Snow", FormatTitle.Apply("jon Snow"))
	assert.Equal(t, "JON SNOW", FormatUpper.Apply("jon Snow"))
}
