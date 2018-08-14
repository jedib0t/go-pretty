package progress

import (
	"time"

	"github.com/jedib0t/go-pretty/text"
)

// Style declares how to render the Progress/Trackers.
type Style struct {
	Name    string       // name of the Style
	Chars   StyleChars   // characters to use on the progress bar
	Colors  StyleColors  // colors to use on the progress bar
	Options StyleOptions // misc. options for the progress bar
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
	BoxLeft    string // left-border
	BoxRight   string // right-border
	Finished   string // finished block
	Finished25 string // 25% finished block
	Finished50 string // 50% finished block
	Finished75 string // 75% finished block
	Unfinished string // 0% finished block
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
	Message text.Colors // message text colors
	Percent text.Colors // percentage text colors
	Stats   text.Colors // stats text (time, value) colors
	Time    text.Colors // time text colors (overrides Stats)
	Tracker text.Colors // tracker text colors
	Value   text.Colors // value text colors (overrides Stats)
}

var (
	// StyleColorsDefault defines sane color choices - None.
	StyleColorsDefault = StyleColors{}

	// StyleColorsExample defines a few choice color options. Use this is just as
	// an example to customize the Tracker/text colors.
	StyleColorsExample = StyleColors{
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
	DoneString              string        // "done!" string
	Separator               string        // text between message and tracker
	SnipIndicator           string        // text denoting message snipping
	PercentFormat           string        // formatting to use for percentage
	TimeDonePrecision       time.Duration // precision for time when done
	TimeInProgressPrecision time.Duration // precision for time when in progress
}

var (
	// StyleOptionsDefault defines sane defaults for the Options. Use this as an
	// example to customize the Tracker rendering.
	StyleOptionsDefault = StyleOptions{
		DoneString:              "done!",
		PercentFormat:           "%5.2f%%",
		Separator:               " ... ",
		SnipIndicator:           "~",
		TimeDonePrecision:       time.Millisecond,
		TimeInProgressPrecision: time.Microsecond,
	}
)
