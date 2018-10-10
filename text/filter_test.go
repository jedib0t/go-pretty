package text

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleFilter() {
	slice := []string{"Arya Stark", "Bran Stark", "Jon Snow", "Sansa Stark"}
	filter := func(item string) bool {
		return strings.HasSuffix(item, "Stark")
	}
	fmt.Printf("%#v\n", Filter(slice, filter))

	// Output: []string{"Arya Stark", "Bran Stark", "Sansa Stark"}
}

func TestFilter(t *testing.T) {
	slice := []string{"Arya Stark", "Bran Stark", "Jon Snow", "Sansa Stark"}
	filter := func(item string) bool {
		return strings.HasSuffix(item, "Stark")
	}

	filteredSlice := Filter(slice, filter)
	assert.Equal(t, 3, len(filteredSlice))
	assert.NotContains(t, filteredSlice, "Jon Snow")
}
