package gopretty

// Unicode character constants for drawing borders & separators.
const (
	BorderBottomLeft            = "└" // \u2514 So BOX DRAWINGS LIGHT UP AND RIGHT
	BorderBottomLeftBold        = "┗" // \u2517 So BOX DRAWINGS HEAVY UP AND RIGHT
	BorderBottomLeftDouble      = "╚" // \u255A So BOX DRAWINGS DOUBLE UP AND RIGHT
	BorderBottomLeftRounded     = "╰" // \u2570 So BOX DRAWINGS LIGHT ARC UP AND RIGHT
	BorderBottomRight           = "┘" // \u2518 So BOX DRAWINGS LIGHT UP AND LEFT
	BorderBottomRightBold       = "┛" // \u251B So BOX DRAWINGS HEAVY UP AND LEFT
	BorderBottomRightDouble     = "╝" // \u255D So BOX DRAWINGS DOUBLE UP AND LEFT
	BorderBottomRightRounded    = "╯" // \u256F So BOX DRAWINGS LIGHT ARC UP AND LEFT
	BorderBottomSeparator       = "┴" // \u2534 So BOX DRAWINGS LIGHT UP AND HORIZONTAL
	BorderBottomSeparatorBold   = "┻" // \u253B So BOX DRAWINGS HEAVY UP AND HORIZONTAL
	BorderBottomSeparatorDouble = "╩" // \u2569 So BOX DRAWINGS DOUBLE UP AND HORIZONTAL
	BorderHorizontal            = "─" // \u2500 So BOX DRAWINGS LIGHT HORIZONTAL
	BorderHorizontalBold        = "━" // \u2501 So BOX DRAWINGS HEAVY HORIZONTAL
	BorderHorizontalDouble      = "═" // \u2550 So BOX DRAWINGS DOUBLE HORIZONTAL
	BorderLeft                  = BorderVertical
	BorderLeftBold              = BorderVerticalBold
	BorderLeftDouble            = BorderVerticalDouble
	BorderLeftSeparator         = "├" // \u251C So BOX DRAWINGS LIGHT VERTICAL AND RIGHT
	BorderLeftSeparatorBold     = "┣" // \u2523 So BOX DRAWINGS HEAVY VERTICAL AND RIGHT
	BorderLeftSeparatorDouble   = "╠" // \u2560 So BOX DRAWINGS DOUBLE VERTICAL AND RIGHT
	BorderRight                 = BorderVertical
	BorderRightBold             = BorderVerticalBold
	BorderRightDouble           = BorderVerticalDouble
	BorderRightSeparator        = "┤" // \u2524 So BOX DRAWINGS LIGHT VERTICAL AND LEFT
	BorderRightSeparatorBold    = "┫" // \u252B So BOX DRAWINGS HEAVY VERTICAL AND LEFT
	BorderRightSeparatorDouble  = "╣" // \u2563 So BOX DRAWINGS DOUBLE VERTICAL AND LEFT
	BorderSeparator             = "┼" // \u253C So BOX DRAWINGS LIGHT VERTICAL AND HORIZONTAL
	BorderSeparatorBold         = "╋" // \u254B So BOX DRAWINGS HEAVY VERTICAL AND HORIZONTAL
	BorderSeparatorDouble       = "╬" // \u256C So BOX DRAWINGS DOUBLE VERTICAL AND HORIZONTAL
	BorderTopLeft               = "┌" // \u250C So BOX DRAWINGS LIGHT DOWN AND RIGHT
	BorderTopLeftBold           = "┏" // \u250F So BOX DRAWINGS HEAVY DOWN AND RIGHT
	BorderTopLeftDouble         = "╔" // \u2554 So BOX DRAWINGS DOUBLE DOWN AND RIGHT
	BorderTopLeftRounded        = "╭" // \u256D So BOX DRAWINGS LIGHT ARC DOWN AND RIGHT
	BorderTopRight              = "┐" // \u2510 So BOX DRAWINGS LIGHT DOWN AND LEFT
	BorderTopRightBold          = "┓" // \u2513 So BOX DRAWINGS HEAVY DOWN AND LEFT
	BorderTopRightDouble        = "╗" // \u2557 So BOX DRAWINGS DOUBLE DOWN AND LEFT
	BorderTopRightRounded       = "╮" // \u256E So BOX DRAWINGS LIGHT ARC DOWN AND LEFT
	BorderTopSeparator          = "┬" // \u252C So BOX DRAWINGS LIGHT DOWN AND HORIZONTAL
	BorderTopSeparatorBold      = "┳" // \u2533 So BOX DRAWINGS HEAVY DOWN AND HORIZONTAL
	BorderTopSeparatorDouble    = "╦" // \u2566 So BOX DRAWINGS DOUBLE DOWN AND HORIZONTAL
	BorderVertical              = "│" // \u2502 So BOX DRAWINGS LIGHT VERTICAL
	BorderVerticalBold          = "┃" // \u2503 So BOX DRAWINGS HEAVY VERTICAL
	BorderVerticalDouble        = "║" // \u2551 So BOX DRAWINGS DOUBLE VERTICAL
	BulletCircle                = "●" // \u25CF So BLACK CIRCLE
	BulletFlower                = "✽" // \u273D So HEAVY TEARDROP-SPOKED ASTERISK
	BulletSquare                = "■" // \u25A0 So BLACK SQUARE
	BulletStar                  = "★" // \u272E So BLACK STAR
	BulletTrianglePointingDown  = "▼" // \u25BC So BLACK DOWN-POINTING TRIANGLE
	BulletTrianglePointingLeft  = "◀" // \u25C0 So BLACK LEFT-POINTING TRIANGLE
	BulletTrianglePointingRight = "▶" // \u25B6 So BLACK RIGHT-POINTING TRIANGLE
	BulletTrianglePointingTop   = "▲" // \u25B2 So BLACK UP-POINTING TRIANGLE
)
