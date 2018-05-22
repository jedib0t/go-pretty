package util

// AutoIndexColumnID returns a unique Column ID/Name for the given Column Number.
// The functionality is similar to what you get in an Excel spreadsheet w.r.t.
// the Column ID/Name.
func AutoIndexColumnID(colIdx int) string {
	charIdx := colIdx % 26
	out := string(65 + charIdx)
	colIdx = colIdx / 26
	if colIdx > 0 {
		return AutoIndexColumnID(colIdx-1) + out
	}
	return out
}
