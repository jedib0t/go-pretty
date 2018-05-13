package util

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterStrings(t *testing.T) {
	slice := []string{"Arya Stark", "Bran Stark", "Jon Snow", "Sansa Stark"}
	filter := func(item string) bool {
		return strings.HasSuffix(item, "Stark")
	}

	filteredSlice := FilterStrings(slice, filter)
	assert.Equal(t, 3, len(filteredSlice))
	assert.NotContains(t, filteredSlice, "Jon Snow")
}
