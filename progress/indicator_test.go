package progress

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIndeterminateIndicatorMovingBackAndForth(t *testing.T) {
	maxLen := 10
	indicator := "<=>"
	expectedPositions := []int{
		0, 1, 2, 3, 4, 5, 6, 7, 6, 5, 4, 3, 2, 1,
		0, 1, 2, 3, 4, 5, 6, 7, 6, 5, 4, 3, 2, 1,
	}

	f := IndeterminateIndicatorMovingBackAndForth(indicator, time.Millisecond*10)
	for idx, expectedPosition := range expectedPositions {
		actual := f(maxLen)
		assert.Equal(t, expectedPosition, actual.Position, fmt.Sprintf("expectedIndeterminateIndicators[%d]", idx))
		time.Sleep(time.Millisecond * 10)
	}
}

func Test_indeterminateIndicatorMovingBackAndForth1(t *testing.T) {
	maxLen := 10
	indicator := "?"
	expectedPositions := []int{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2, 1,
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2, 1,
	}

	f := indeterminateIndicatorMovingBackAndForth(indicator)
	for idx, expectedPosition := range expectedPositions {
		actual := f(maxLen)
		assert.Equal(t, expectedPosition, actual.Position, fmt.Sprintf("expectedIndeterminateIndicators[%d]", idx))
	}
}

func Test_indeterminateIndicatorMovingBackAndForth2(t *testing.T) {
	maxLen := 10
	indicator := "<>"
	expectedPositions := []int{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 7, 6, 5, 4, 3, 2, 1,
		0, 1, 2, 3, 4, 5, 6, 7, 8, 7, 6, 5, 4, 3, 2, 1,
	}

	f := indeterminateIndicatorMovingBackAndForth(indicator)
	for idx, expectedPosition := range expectedPositions {
		actual := f(maxLen)
		assert.Equal(t, expectedPosition, actual.Position, fmt.Sprintf("expectedIndeterminateIndicators[%d]", idx))
	}
}

func Test_indeterminateIndicatorMovingBackAndForth3(t *testing.T) {
	maxLen := 10
	indicator := "<=>"
	expectedPositions := []int{
		0, 1, 2, 3, 4, 5, 6, 7, 6, 5, 4, 3, 2, 1,
		0, 1, 2, 3, 4, 5, 6, 7, 6, 5, 4, 3, 2, 1,
	}

	f := indeterminateIndicatorMovingBackAndForth(indicator)
	for idx, expectedPosition := range expectedPositions {
		actual := f(maxLen)
		assert.Equal(t, expectedPosition, actual.Position, fmt.Sprintf("expectedIndeterminateIndicators[%d]", idx))
	}
}

func TestIndeterminateIndicatorMovingLeftToRight(t *testing.T) {
	maxLen := 10
	indicator := "?"
	expectedPositions := []int{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
	}

	f := IndeterminateIndicatorMovingLeftToRight(indicator, time.Millisecond*10)
	for idx, expectedPosition := range expectedPositions {
		actual := f(maxLen)
		assert.Equal(t, expectedPosition, actual.Position, fmt.Sprintf("expectedIndeterminateIndicators[%d]", idx))
		time.Sleep(time.Millisecond * 10)
	}
}

func Test_indeterminateIndicatorMovingLeftToRight1(t *testing.T) {
	maxLen := 10
	indicator := "?"
	expectedPositions := []int{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
	}

	f := indeterminateIndicatorMovingLeftToRight(indicator)
	for idx, expectedPosition := range expectedPositions {
		actual := f(maxLen)
		assert.Equal(t, expectedPosition, actual.Position, fmt.Sprintf("expectedIndeterminateIndicators[%d]", idx))
	}
}

func Test_indeterminateIndicatorMovingLeftToRight2(t *testing.T) {
	maxLen := 10
	indicator := "<>"
	expectedPositions := []int{
		0, 1, 2, 3, 4, 5, 6, 7, 8,
		0, 1, 2, 3, 4, 5, 6, 7, 8,
	}

	f := indeterminateIndicatorMovingLeftToRight(indicator)
	for idx, expectedPosition := range expectedPositions {
		actual := f(maxLen)
		assert.Equal(t, expectedPosition, actual.Position, fmt.Sprintf("expectedIndeterminateIndicators[%d]", idx))
	}
}

func Test_indeterminateIndicatorMovingLeftToRight3(t *testing.T) {
	maxLen := 10
	indicator := "<=>"
	expectedPositions := []int{
		0, 1, 2, 3, 4, 5, 6, 7,
		0, 1, 2, 3, 4, 5, 6, 7,
	}

	f := indeterminateIndicatorMovingLeftToRight(indicator)
	for idx, expectedPosition := range expectedPositions {
		actual := f(maxLen)
		assert.Equal(t, expectedPosition, actual.Position, fmt.Sprintf("expectedIndeterminateIndicators[%d]", idx))
	}
}

func TestIndeterminateIndicatorMovingRightToLeft(t *testing.T) {
	maxLen := 10
	indicator := "?"
	expectedPositions := []int{
		9, 8, 7, 6, 5, 4, 3, 2, 1, 0,
		9, 8, 7, 6, 5, 4, 3, 2, 1, 0,
	}

	f := IndeterminateIndicatorMovingRightToLeft(indicator, time.Millisecond*10)
	for idx, expectedPosition := range expectedPositions {
		actual := f(maxLen)
		assert.Equal(t, expectedPosition, actual.Position, fmt.Sprintf("expectedIndeterminateIndicators[%d]", idx))
		time.Sleep(time.Millisecond * 10)
	}
}

func Test_indeterminateIndicatorMovingRightToLeft1(t *testing.T) {
	maxLen := 10
	indicator := "?"
	expectedPositions := []int{
		9, 8, 7, 6, 5, 4, 3, 2, 1, 0,
		9, 8, 7, 6, 5, 4, 3, 2, 1, 0,
	}

	f := indeterminateIndicatorMovingRightToLeft(indicator)
	for idx, expectedPosition := range expectedPositions {
		actual := f(maxLen)
		assert.Equal(t, expectedPosition, actual.Position, fmt.Sprintf("expectedIndeterminateIndicators[%d]", idx))
	}
}

func Test_indeterminateIndicatorMovingRightToLeft2(t *testing.T) {
	maxLen := 10
	indicator := "<>"
	expectedPositions := []int{
		8, 7, 6, 5, 4, 3, 2, 1, 0,
		8, 7, 6, 5, 4, 3, 2, 1, 0,
	}

	f := indeterminateIndicatorMovingRightToLeft(indicator)
	for idx, expectedPosition := range expectedPositions {
		actual := f(maxLen)
		assert.Equal(t, expectedPosition, actual.Position, fmt.Sprintf("expectedIndeterminateIndicators[%d]", idx))
	}
}

func Test_indeterminateIndicatorMovingRightToLeft3(t *testing.T) {
	maxLen := 10
	indicator := "<=>"
	expectedPositions := []int{
		7, 6, 5, 4, 3, 2, 1, 0,
		7, 6, 5, 4, 3, 2, 1, 0,
	}

	f := indeterminateIndicatorMovingRightToLeft(indicator)
	for idx, expectedPosition := range expectedPositions {
		actual := f(maxLen)
		assert.Equal(t, expectedPosition, actual.Position, fmt.Sprintf("expectedIndeterminateIndicators[%d]", idx))
	}
}
