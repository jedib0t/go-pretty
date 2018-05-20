package progress

import (
	"time"

	"github.com/jedib0t/go-pretty/text"
)

// Style declares how to render the Progress/Trackers.
type Style struct {
	Name    string
	Chars   StyleChars
	Colors  StyleColors
	Options StyleOptions
}

var (
	// StyleDefault uses ASCII text to render the Trackers.
	StyleDefault = Style{
		Name:    "StyleDefault",
		Chars:   StyleCharsDefault,
		Colors:  StyleColorsDefault,
		Options: StyleOptionsDefault,
	}

	// StyleBlocks uses UNICODE Block Drawing characters to render the Trackers.
	StyleBlocks = Style{
		Name:    "StyleBlocks",
		Chars:   StyleCharsBlocks,
		Colors:  StyleColorsDefault,
		Options: StyleOptionsDefault,
	}

	// StyleCircle uses UNICODE Circle runes to render the Trackers.
	StyleCircle = Style{
		Name:    "StyleCircle",
		Chars:   StyleCharsCircle,
		Colors:  StyleColorsDefault,
		Options: StyleOptionsDefault,
	}

	// StyleRhombus uses UNICODE Rhombus runes to render the Trackers.
	StyleRhombus = Style{
		Name:    "StyleRhombus",
		Chars:   StyleCharsRhombus,
		Colors:  StyleColorsDefault,
		Options: StyleOptionsDefault,
	}
)

// StyleChars defines the characters/strings to use for rendering the Tracker.
type StyleChars struct {
	BoxLeft    string
	BoxRight   string
	Finished   string
	Finished25 string
	Finished50 string
	Finished75 string
	Unfinished string
}

var (
	// StyleCharsDefault uses simple ASCII characters.
	StyleCharsDefault = StyleChars{
		BoxLeft:    "[",
		BoxRight:   "]",
		Finished:   "#",
		Finished25: ".",
		Finished50: ".",
		Finished75: ".",
		Unfinished: ".",
	}

	// StyleCharsBlocks uses UNICODE Block Drawing characters.
	StyleCharsBlocks = StyleChars{
		BoxLeft:    "║",
		BoxRight:   "║",
		Finished:   "█",
		Finished25: "░",
		Finished50: "▒",
		Finished75: "▓",
		Unfinished: "░",
	}

	// StyleCharsCircle uses UNICODE Circle characters.
	StyleCharsCircle = StyleChars{
		BoxLeft:    "(",
		BoxRight:   ")",
		Finished:   "●",
		Finished25: "○",
		Finished50: "○",
		Finished75: "○",
		Unfinished: "◌",
	}

	// StyleCharsRhombus uses UNICODE Rhombus characters.
	StyleCharsRhombus = StyleChars{
		BoxLeft:    "<",
		BoxRight:   ">",
		Finished:   "◆",
		Finished25: "◈",
		Finished50: "◈",
		Finished75: "◈",
		Unfinished: "◇",
	}
)

// StyleColors defines what colors to use for various parts of the Progress and
// Tracker texts.
type StyleColors struct {
	Done    text.Colors
	Message text.Colors
	Percent text.Colors
	Stats   text.Colors
	Time    text.Colors
	Tracker text.Colors
	Value   text.Colors
}

var (
	// StyleColorsDefault defines sane color choices - None.
	StyleColorsDefault = StyleColors{}

	// StyleColorsExample defines a few choice color options. Use this is just as
	// an example to customize the Tracker/text colors.
	StyleColorsExample = StyleColors{
		Done:    text.Colors{text.FgWhite, text.BgBlack},
		Message: text.Colors{text.FgWhite, text.BgBlack},
		Percent: text.Colors{text.FgHiRed, text.BgBlack},
		Stats:   text.Colors{text.FgHiBlack, text.BgBlack},
		Time:    text.Colors{text.FgGreen, text.BgBlack},
		Tracker: text.Colors{text.FgYellow, text.BgBlack},
		Value:   text.Colors{text.FgCyan, text.BgBlack},
	}
)

// StyleOptions defines misc. options to control how the Tracker or its parts
// gets rendered.
type StyleOptions struct {
	DoneString              string
	MessageTrackerSeparator string
	PercentFormat           string
	TimeDonePrecision       time.Duration
	TimeInProgressPrecision time.Duration
}

var (
	// StyleOptionsDefault defines sane defaults for the Options. Use this as an
	// example to customize the Tracker rendering.
	StyleOptionsDefault = StyleOptions{
		DoneString:              "done!",
		MessageTrackerSeparator: "...",
		PercentFormat:           "%5.2f%%",
		TimeDonePrecision:       time.Millisecond,
		TimeInProgressPrecision: time.Microsecond,
	}
)
