package text

// Unicode character constants for drawing borders & separators.
const (
	BoxBottomLeft               = "└" // \u2514 So BOX DRAWINGS LIGHT UP AND RIGHT
	BoxBottomLeftBold           = "┗" // \u2517 So BOX DRAWINGS HEAVY UP AND RIGHT
	BoxBottomLeftDouble         = "╚" // \u255A So BOX DRAWINGS DOUBLE UP AND RIGHT
	BoxBottomLeftRounded        = "╰" // \u2570 So BOX DRAWINGS LIGHT ARC UP AND RIGHT
	BoxBottomRight              = "┘" // \u2518 So BOX DRAWINGS LIGHT UP AND LEFT
	BoxBottomRightBold          = "┛" // \u251B So BOX DRAWINGS HEAVY UP AND LEFT
	BoxBottomRightDouble        = "╝" // \u255D So BOX DRAWINGS DOUBLE UP AND LEFT
	BoxBottomRightRounded       = "╯" // \u256F So BOX DRAWINGS LIGHT ARC UP AND LEFT
	BoxBottomSeparator          = "┴" // \u2534 So BOX DRAWINGS LIGHT UP AND HORIZONTAL
	BoxBottomSeparatorBold      = "┻" // \u253B So BOX DRAWINGS HEAVY UP AND HORIZONTAL
	BoxBottomSeparatorDouble    = "╩" // \u2569 So BOX DRAWINGS DOUBLE UP AND HORIZONTAL
	BoxHorizontal               = "─" // \u2500 So BOX DRAWINGS LIGHT HORIZONTAL
	BoxHorizontalBold           = "━" // \u2501 So BOX DRAWINGS HEAVY HORIZONTAL
	BoxHorizontalDouble         = "═" // \u2550 So BOX DRAWINGS DOUBLE HORIZONTAL
	BoxLeft                     = BoxVertical
	BoxLeftBold                 = BoxVerticalBold
	BoxLeftDouble               = BoxVerticalDouble
	BoxLeftSeparator            = "├" // \u251C So BOX DRAWINGS LIGHT VERTICAL AND RIGHT
	BoxLeftSeparatorBold        = "┣" // \u2523 So BOX DRAWINGS HEAVY VERTICAL AND RIGHT
	BoxLeftSeparatorDouble      = "╠" // \u2560 So BOX DRAWINGS DOUBLE VERTICAL AND RIGHT
	BoxRight                    = BoxVertical
	BoxRightBold                = BoxVerticalBold
	BoxRightDouble              = BoxVerticalDouble
	BoxRightSeparator           = "┤" // \u2524 So BOX DRAWINGS LIGHT VERTICAL AND LEFT
	BoxRightSeparatorBold       = "┫" // \u252B So BOX DRAWINGS HEAVY VERTICAL AND LEFT
	BoxRightSeparatorDouble     = "╣" // \u2563 So BOX DRAWINGS DOUBLE VERTICAL AND LEFT
	BoxSeparator                = "┼" // \u253C So BOX DRAWINGS LIGHT VERTICAL AND HORIZONTAL
	BoxSeparatorBold            = "╋" // \u254B So BOX DRAWINGS HEAVY VERTICAL AND HORIZONTAL
	BoxSeparatorDouble          = "╬" // \u256C So BOX DRAWINGS DOUBLE VERTICAL AND HORIZONTAL
	BoxTopLeft                  = "┌" // \u250C So BOX DRAWINGS LIGHT DOWN AND RIGHT
	BoxTopLeftBold              = "┏" // \u250F So BOX DRAWINGS HEAVY DOWN AND RIGHT
	BoxTopLeftDouble            = "╔" // \u2554 So BOX DRAWINGS DOUBLE DOWN AND RIGHT
	BoxTopLeftRounded           = "╭" // \u256D So BOX DRAWINGS LIGHT ARC DOWN AND RIGHT
	BoxTopRight                 = "┐" // \u2510 So BOX DRAWINGS LIGHT DOWN AND LEFT
	BoxTopRightBold             = "┓" // \u2513 So BOX DRAWINGS HEAVY DOWN AND LEFT
	BoxTopRightDouble           = "╗" // \u2557 So BOX DRAWINGS DOUBLE DOWN AND LEFT
	BoxTopRightRounded          = "╮" // \u256E So BOX DRAWINGS LIGHT ARC DOWN AND LEFT
	BoxTopSeparator             = "┬" // \u252C So BOX DRAWINGS LIGHT DOWN AND HORIZONTAL
	BoxTopSeparatorBold         = "┳" // \u2533 So BOX DRAWINGS HEAVY DOWN AND HORIZONTAL
	BoxTopSeparatorDouble       = "╦" // \u2566 So BOX DRAWINGS DOUBLE DOWN AND HORIZONTAL
	BoxVertical                 = "│" // \u2502 So BOX DRAWINGS LIGHT VERTICAL
	BoxVerticalBold             = "┃" // \u2503 So BOX DRAWINGS HEAVY VERTICAL
	BoxVerticalDouble           = "║" // \u2551 So BOX DRAWINGS DOUBLE VERTICAL
	BulletCircle                = "●" // \u25CF So BLACK CIRCLE
	BulletFlower                = "✽" // \u273D So HEAVY TEARDROP-SPOKED ASTERISK
	BulletSquare                = "■" // \u25A0 So BLACK SQUARE
	BulletStar                  = "★" // \u272E So BLACK STAR
	BulletTrianglePointingDown  = "▼" // \u25BC So BLACK DOWN-POINTING TRIANGLE
	BulletTrianglePointingLeft  = "◀" // \u25C0 So BLACK LEFT-POINTING TRIANGLE
	BulletTrianglePointingRight = "▶" // \u25B6 So BLACK RIGHT-POINTING TRIANGLE
	BulletTrianglePointingTop   = "▲" // \u25B2 So BLACK UP-POINTING TRIANGLE
)
