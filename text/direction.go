package text

// Direction defines the overall flow of text. Similar to bidi.Direction, but
// simplified and specific to this package.
type Direction int

// Available Directions.
const (
	Default Direction = iota
	LeftToRight
	RightToLeft
)

func (d Direction) Modifier() string {
	switch d {
	case Default:
		return ""
	case LeftToRight:
		return "‪"
	case RightToLeft:
		return "‫"
	}
	return ""
}
