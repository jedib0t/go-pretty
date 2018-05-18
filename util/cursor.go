package util

import (
	"fmt"
)

// Cursor helps move the cursor on the console in multiple directions.
type Cursor rune

const (
	// CursorDown helps move the Cursor Down X lines
	CursorDown Cursor = 'B'

	// CursorLeft helps move the Cursor Left X characters
	CursorLeft Cursor = 'D'

	// CursorRight helps move the Cursor Right X characters
	CursorRight Cursor = 'C'

	// CursorUp helps move the Cursor Up X lines
	CursorUp Cursor = 'A'

	// EraseLine helps erase all characters to the Right of the Cursor in the
	// current line
	EraseLine Cursor = 'K'
)

// Sprint prints the Escape Sequence to move the Cursor "x" times.
func (c Cursor) Sprint(x ...int) string {
	if len(x) > 0 && x[0] > 0 {
		return fmt.Sprintf("%s%d%c", EscapeStart, x[0], c)
	}
	return fmt.Sprintf("%s%c", EscapeStart, c)
}
