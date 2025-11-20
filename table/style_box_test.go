package table

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBoxStyleHorizontal(t *testing.T) {
	t.Run("empty string", func(t *testing.T) {
		horizontal := NewBoxStyleHorizontal("")
		assert.NotNil(t, horizontal)
		assert.Equal(t, "", horizontal.TitleTop)
		assert.Equal(t, "", horizontal.TitleBottom)
		assert.Equal(t, "", horizontal.HeaderTop)
		assert.Equal(t, "", horizontal.HeaderMiddle)
		assert.Equal(t, "", horizontal.HeaderBottom)
		assert.Equal(t, "", horizontal.RowTop)
		assert.Equal(t, "", horizontal.RowMiddle)
		assert.Equal(t, "", horizontal.RowBottom)
		assert.Equal(t, "", horizontal.FooterTop)
		assert.Equal(t, "", horizontal.FooterMiddle)
		assert.Equal(t, "", horizontal.FooterBottom)
	})

	t.Run("single character", func(t *testing.T) {
		horizontal := NewBoxStyleHorizontal("-")
		assert.NotNil(t, horizontal)
		assert.Equal(t, "-", horizontal.TitleTop)
		assert.Equal(t, "-", horizontal.TitleBottom)
		assert.Equal(t, "-", horizontal.HeaderTop)
		assert.Equal(t, "-", horizontal.HeaderMiddle)
		assert.Equal(t, "-", horizontal.HeaderBottom)
		assert.Equal(t, "-", horizontal.RowTop)
		assert.Equal(t, "-", horizontal.RowMiddle)
		assert.Equal(t, "-", horizontal.RowBottom)
		assert.Equal(t, "-", horizontal.FooterTop)
		assert.Equal(t, "-", horizontal.FooterMiddle)
		assert.Equal(t, "-", horizontal.FooterBottom)
	})

	t.Run("multi-character string", func(t *testing.T) {
		horizontal := NewBoxStyleHorizontal("═══")
		assert.NotNil(t, horizontal)
		assert.Equal(t, "═══", horizontal.TitleTop)
		assert.Equal(t, "═══", horizontal.TitleBottom)
		assert.Equal(t, "═══", horizontal.HeaderTop)
		assert.Equal(t, "═══", horizontal.HeaderMiddle)
		assert.Equal(t, "═══", horizontal.HeaderBottom)
		assert.Equal(t, "═══", horizontal.RowTop)
		assert.Equal(t, "═══", horizontal.RowMiddle)
		assert.Equal(t, "═══", horizontal.RowBottom)
		assert.Equal(t, "═══", horizontal.FooterTop)
		assert.Equal(t, "═══", horizontal.FooterMiddle)
		assert.Equal(t, "═══", horizontal.FooterBottom)
	})
}

func TestBoxStyle_ensureHorizontalInitialized(t *testing.T) {
	t.Run("nil horizontal initializes with MiddleHorizontal", func(t *testing.T) {
		bs := &BoxStyle{
			MiddleHorizontal: "-",
			Horizontal:       nil,
		}

		bs.ensureHorizontalInitialized()

		assert.NotNil(t, bs.Horizontal)
		assert.Equal(t, "-", bs.Horizontal.TitleTop)
		assert.Equal(t, "-", bs.Horizontal.HeaderMiddle)
		assert.Equal(t, "-", bs.Horizontal.RowMiddle)
	})

	t.Run("nil horizontal with empty MiddleHorizontal", func(t *testing.T) {
		bs := &BoxStyle{
			MiddleHorizontal: "",
			Horizontal:       nil,
		}

		bs.ensureHorizontalInitialized()

		assert.NotNil(t, bs.Horizontal)
		assert.Equal(t, "", bs.Horizontal.TitleTop)
		assert.Equal(t, "", bs.Horizontal.HeaderMiddle)
	})

	t.Run("existing horizontal not overwritten", func(t *testing.T) {
		existing := &BoxStyleHorizontal{
			TitleTop:     "0",
			TitleBottom:  "1",
			HeaderTop:    "2",
			HeaderMiddle: "3",
			HeaderBottom: "4",
			RowTop:       "5",
			RowMiddle:    "6",
			RowBottom:    "7",
			FooterTop:    "8",
			FooterMiddle: "9",
			FooterBottom: "A",
		}

		bs := &BoxStyle{
			MiddleHorizontal: "-",
			Horizontal:       existing,
		}

		bs.ensureHorizontalInitialized()

		// Should not be overwritten
		assert.Equal(t, existing, bs.Horizontal)
		assert.Equal(t, "0", bs.Horizontal.TitleTop)
		assert.Equal(t, "1", bs.Horizontal.TitleBottom)
		assert.Equal(t, "2", bs.Horizontal.HeaderTop)
		assert.Equal(t, "3", bs.Horizontal.HeaderMiddle)
		assert.Equal(t, "4", bs.Horizontal.HeaderBottom)
		assert.Equal(t, "5", bs.Horizontal.RowTop)
		assert.Equal(t, "6", bs.Horizontal.RowMiddle)
		assert.Equal(t, "7", bs.Horizontal.RowBottom)
		assert.Equal(t, "8", bs.Horizontal.FooterTop)
		assert.Equal(t, "9", bs.Horizontal.FooterMiddle)
		assert.Equal(t, "A", bs.Horizontal.FooterBottom)
	})

	t.Run("multiple calls are idempotent", func(t *testing.T) {
		bs := &BoxStyle{
			MiddleHorizontal: "-",
			Horizontal:       nil,
		}

		bs.ensureHorizontalInitialized()
		firstCall := bs.Horizontal

		bs.ensureHorizontalInitialized()
		secondCall := bs.Horizontal

		// Should be the same instance
		assert.Equal(t, firstCall, secondCall)
	})
}

func TestBoxStyle_middleHorizontal(t *testing.T) {
	t.Run("with nil horizontal initializes automatically", func(t *testing.T) {
		bs := &BoxStyle{
			MiddleHorizontal: "-",
			Horizontal:       nil,
		}

		result := bs.middleHorizontal(separatorTypeRowMiddle)

		assert.NotNil(t, bs.Horizontal)
		assert.Equal(t, "-", result)
	})

	t.Run("all separator types return correct values", func(t *testing.T) {
		bs := &BoxStyle{
			Horizontal: &BoxStyleHorizontal{
				TitleTop:     "0",
				TitleBottom:  "1",
				HeaderTop:    "2",
				HeaderMiddle: "3",
				HeaderBottom: "4",
				RowTop:       "5",
				RowMiddle:    "6",
				RowBottom:    "7",
				FooterTop:    "8",
				FooterMiddle: "9",
				FooterBottom: "A",
			},
		}

		testCases := []struct {
			name           string
			separatorType  separatorType
			expectedResult string
		}{
			{"TitleTop", separatorTypeTitleTop, "0"},
			{"TitleBottom", separatorTypeTitleBottom, "1"},
			{"HeaderTop", separatorTypeHeaderTop, "2"},
			{"HeaderMiddle", separatorTypeHeaderMiddle, "3"},
			{"HeaderBottom", separatorTypeHeaderBottom, "4"},
			{"RowTop", separatorTypeRowTop, "5"},
			{"RowMiddle", separatorTypeRowMiddle, "6"},
			{"RowBottom", separatorTypeRowBottom, "7"},
			{"FooterTop", separatorTypeFooterTop, "8"},
			{"FooterMiddle", separatorTypeFooterMiddle, "9"},
			{"FooterBottom", separatorTypeFooterBottom, "A"},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				result := bs.middleHorizontal(tc.separatorType)
				assert.Equal(t, tc.expectedResult, result)
			})
		}
	})

	t.Run("default case returns RowMiddle", func(t *testing.T) {
		bs := &BoxStyle{
			Horizontal: &BoxStyleHorizontal{
				RowMiddle: "default-value",
			},
		}

		// Use an invalid separatorType that doesn't match any case
		// separatorTypeCount is the last valid value, so anything >= separatorTypeCount
		// should fall through to default
		result := bs.middleHorizontal(separatorTypeCount)
		assert.Equal(t, "default-value", result)

		// Also test with a value beyond separatorTypeCount
		result2 := bs.middleHorizontal(separatorTypeCount + 1)
		assert.Equal(t, "default-value", result2)
	})

	t.Run("with pre-initialized horizontal", func(t *testing.T) {
		bs := &BoxStyle{
			MiddleHorizontal: "should-not-be-used",
			Horizontal: &BoxStyleHorizontal{
				TitleTop:  "custom-title",
				RowMiddle: "custom-row",
			},
		}

		// Should use the custom values, not MiddleHorizontal
		assert.Equal(t, "custom-title", bs.middleHorizontal(separatorTypeTitleTop))
		assert.Equal(t, "custom-row", bs.middleHorizontal(separatorTypeRowMiddle))
	})
}
