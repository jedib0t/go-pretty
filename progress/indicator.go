package progress

import (
	"strings"
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

// IndeterminateIndicatorDominoes returns an instance of
// IndeterminateIndicatorGenerator function that simulates a bunch of dominoes
// falling.
func IndeterminateIndicatorDominoes(duration time.Duration) IndeterminateIndicatorGenerator {
	var indeterminateIndicator *IndeterminateIndicator
	indicatorGenerator := indeterminateIndicatorDominoes()
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

// IndeterminateIndicatorPacMan returns an instance of
// IndeterminateIndicatorGenerator function that simulates a Pac-Man character
// chomping through the progress bar.
func IndeterminateIndicatorPacMan(duration time.Duration) IndeterminateIndicatorGenerator {
	var indeterminateIndicator *IndeterminateIndicator
	indicatorGenerator := indeterminateIndicatorPacMan()
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

func indeterminateIndicatorDominoes() IndeterminateIndicatorGenerator {
	direction := 1 // positive == left to right; negative == right to left
	nextPosition := 0

	out := strings.Builder{}
	generateIndicator := func(currentPosition int, maxLen int) string {
		out.Reset()
		out.WriteString(strings.Repeat("/", currentPosition))
		out.WriteString(strings.Repeat("\\", maxLen-currentPosition))
		return out.String()
	}

	return func(maxLen int) IndeterminateIndicator {
		currentPosition := nextPosition

		if currentPosition == 0 {
			direction = 1
		} else if currentPosition == maxLen {
			direction = -1
		}
		nextPosition += direction

		return IndeterminateIndicator{
			Position: 0,
			Text:     generateIndicator(currentPosition, maxLen),
		}
	}
}

func indeterminateIndicatorMovingBackAndForth(indicator string) IndeterminateIndicatorGenerator {
	direction := 1 // positive == left to right; negative == right to left
	nextPosition := 0

	return func(maxLen int) IndeterminateIndicator {
		currentPosition := nextPosition

		if currentPosition == 0 {
			direction = 1
		} else if currentPosition+text.RuneCount(indicator) == maxLen {
			direction = -1
		}
		nextPosition += direction

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

func indeterminateIndicatorPacMan() IndeterminateIndicatorGenerator {
	pacManMovingRight, pacManMovingLeft := "ᗧ", "ᗤ"
	direction := 1 // positive == left to right; negative == right to left
	indicator := pacManMovingRight
	nextPosition := 0

	out := strings.Builder{}
	generateIndicator := func(currentPosition int, maxLen int) string {
		out.Reset()
		if currentPosition > 0 {
			out.WriteString(strings.Repeat(" ", currentPosition))
		}
		out.WriteString(indicator)
		out.WriteString(strings.Repeat(" ", maxLen-currentPosition-1))
		return out.String()
	}

	return func(maxLen int) IndeterminateIndicator {
		currentPosition := nextPosition
		currentText := generateIndicator(currentPosition, maxLen)

		if currentPosition == 0 {
			direction = 1
			indicator = pacManMovingRight
		} else if currentPosition+text.RuneCount(indicator) == maxLen {
			direction = -1
			indicator = pacManMovingLeft
		}
		nextPosition += direction

		return IndeterminateIndicator{
			Position: 0,
			Text:     currentText,
		}
	}
}
