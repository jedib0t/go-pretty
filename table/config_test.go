package table

import (
	"testing"

	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/stretchr/testify/assert"
)

func TestColumnConfig_getWidthMaxEnforcer(t *testing.T) {
	t.Run("no width enforcer", func(t *testing.T) {
		cc := ColumnConfig{}

		widthEnforcer := cc.getWidthMaxEnforcer()
		assert.Equal(t, "1234567890", widthEnforcer("1234567890", 0))
		assert.Equal(t, "1234567890", widthEnforcer("1234567890", 1))
		assert.Equal(t, "1234567890", widthEnforcer("1234567890", 5))
		assert.Equal(t, "1234567890", widthEnforcer("1234567890", 10))
		assert.Equal(t, "1234567890", widthEnforcer("1234567890", 100))
		assert.Equal(t, "1234567890", widthEnforcer("1234567890", 1000))
	})

	t.Run("default width enforcer", func(t *testing.T) {
		cc := ColumnConfig{
			WidthMax: 10,
		}

		widthEnforcer := cc.getWidthMaxEnforcer()
		assert.Equal(t, "", widthEnforcer("1234567890", 0))
		assert.Equal(t, "1\n2\n3\n4\n5\n6\n7\n8\n9\n0", widthEnforcer("1234567890", 1))
		assert.Equal(t, "12345\n67890", widthEnforcer("1234567890", 5))
		assert.Equal(t, "1234567890", widthEnforcer("1234567890", 10))
		assert.Equal(t, "1234567890", widthEnforcer("1234567890", 100))
		assert.Equal(t, "1234567890", widthEnforcer("1234567890", 1000))
	})

	t.Run("custom width enforcer (1)", func(t *testing.T) {
		cc := ColumnConfig{
			WidthMax:         10,
			WidthMaxEnforcer: text.Trim,
		}

		widthEnforcer := cc.getWidthMaxEnforcer()
		assert.Equal(t, text.Trim("1234567890", 0), widthEnforcer("1234567890", 0))
		assert.Equal(t, text.Trim("1234567890", 1), widthEnforcer("1234567890", 1))
		assert.Equal(t, text.Trim("1234567890", 5), widthEnforcer("1234567890", 5))
		assert.Equal(t, text.Trim("1234567890", 10), widthEnforcer("1234567890", 10))
		assert.Equal(t, text.Trim("1234567890", 100), widthEnforcer("1234567890", 100))
		assert.Equal(t, text.Trim("1234567890", 1000), widthEnforcer("1234567890", 1000))
	})

	t.Run("custom width enforcer (2)", func(t *testing.T) {
		cc := ColumnConfig{
			WidthMax: 10,
			WidthMaxEnforcer: func(col string, maxLen int) string {
				return "foo"
			},
		}

		widthEnforcer := cc.getWidthMaxEnforcer()
		assert.Equal(t, "foo", widthEnforcer("1234567890", 0))
		assert.Equal(t, "foo", widthEnforcer("1234567890", 1))
		assert.Equal(t, "foo", widthEnforcer("1234567890", 5))
		assert.Equal(t, "foo", widthEnforcer("1234567890", 10))
		assert.Equal(t, "foo", widthEnforcer("1234567890", 100))
		assert.Equal(t, "foo", widthEnforcer("1234567890", 1000))
	})
}
