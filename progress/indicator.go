package progress

import (
	"time"

	"github.com/jedib0t/go-pretty/v6/text"
)

// IndeterminateIndicator defines the structure for the indicator to indicate
// indeterminate progress. Ex.: <=>
type IndeterminateIndicator struct {
	Position int
	Text     string
}

// IndeterminateIndicatorGenerator returns an IndeterminateIndicator for cases
// where the progress percentage cannot be calculated. Ex.: [........<=>....]
type IndeterminateIndicatorGenerator func(maxLen int) IndeterminateIndicator

// IndeterminateIndicatorMovingBackAndForth returns an instance of
// IndeterminateIndicatorGenerator function that incrementally moves from the
// left to right and back for each specified duration. If duration is 0, then
// every single invocation moves the indicator.
func IndeterminateIndicatorMovingBackAndForth(indicator string, duration time.Duration) IndeterminateIndicatorGenerator {
	var indeterminateIndicator *IndeterminateIndicator
	indicatorGenerator := indeterminateIndicatorMovingBackAndForth(indicator)
	lastRenderTime := time.Now()

	return func(maxLen int) IndeterminateIndicator {
		currRenderTime := time.Now()
		if indeterminateIndicator == nil || duration == 0 || currRenderTime.Sub(lastRenderTime) > duration {
			tmpIndeterminateIndicator := indicatorGenerator(maxLen)
			indeterminateIndicator = &tmpIndeterminateIndicator
			lastRenderTime = currRenderTime
		}

		return *indeterminateIndicator
	}
}

// IndeterminateIndicatorMovingLeftToRight returns an instance of
// IndeterminateIndicatorGenerator function that incrementally moves from the
// left to right and starts from left again for each specified duration. If
// duration is 0, then every single invocation moves the indicator.
func IndeterminateIndicatorMovingLeftToRight(indicator string, duration time.Duration) IndeterminateIndicatorGenerator {
	var indeterminateIndicator *IndeterminateIndicator
	indicatorGenerator := indeterminateIndicatorMovingLeftToRight(indicator)
	lastRenderTime := time.Now()

	return func(maxLen int) IndeterminateIndicator {
		currRenderTime := time.Now()
		if indeterminateIndicator == nil || duration == 0 || currRenderTime.Sub(lastRenderTime) > duration {
			tmpIndeterminateIndicator := indicatorGenerator(maxLen)
			indeterminateIndicator = &tmpIndeterminateIndicator
			lastRenderTime = currRenderTime
		}

		return *indeterminateIndicator
	}
}

// IndeterminateIndicatorMovingRightToLeft returns an instance of
// IndeterminateIndicatorGenerator function that incrementally moves from the
// right to left and starts from right again for each specified duration. If
// duration is 0, then every single invocation moves the indicator.
func IndeterminateIndicatorMovingRightToLeft(indicator string, duration time.Duration) IndeterminateIndicatorGenerator {
	var indeterminateIndicator *IndeterminateIndicator
	indicatorGenerator := indeterminateIndicatorMovingRightToLeft(indicator)
	lastRenderTime := time.Now()

	return func(maxLen int) IndeterminateIndicator {
		currRenderTime := time.Now()
		if indeterminateIndicator == nil || duration == 0 || currRenderTime.Sub(lastRenderTime) > duration {
			tmpIndeterminateIndicator := indicatorGenerator(maxLen)
			indeterminateIndicator = &tmpIndeterminateIndicator
			lastRenderTime = currRenderTime
		}

		return *indeterminateIndicator
	}
}

func indeterminateIndicatorMovingBackAndForth(indicator string) IndeterminateIndicatorGenerator {
	increment := 1
	nextPosition := 0

	return func(maxLen int) IndeterminateIndicator {
		currentPosition := nextPosition
		if currentPosition == 0 {
			increment = 1
		} else if currentPosition+text.RuneCount(indicator) == maxLen {
			increment = -1
		}
		nextPosition += increment

		return IndeterminateIndicator{
			Position: currentPosition,
			Text:     indicator,
		}
	}
}

func indeterminateIndicatorMovingLeftToRight(indicator string) IndeterminateIndicatorGenerator {
	nextPosition := 0

	return func(maxLen int) IndeterminateIndicator {
		currentPosition := nextPosition
		nextPosition++
		if nextPosition+text.RuneCount(indicator) > maxLen {
			nextPosition = 0
		}

		return IndeterminateIndicator{
			Position: currentPosition,
			Text:     indicator,
		}
	}
}

func indeterminateIndicatorMovingRightToLeft(indicator string) IndeterminateIndicatorGenerator {
	nextPosition := -1

	return func(maxLen int) IndeterminateIndicator {
		if nextPosition == -1 {
			nextPosition = maxLen - text.RuneCount(indicator)
		}
		currentPosition := nextPosition
		nextPosition--

		return IndeterminateIndicator{
			Position: currentPosition,
			Text:     indicator,
		}
	}
}
