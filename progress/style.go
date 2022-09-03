package progress

import (
	"time"

	"github.com/jedib0t/go-pretty/v6/text"
)

// Style declares how to render the Progress/Trackers.
type Style struct {
	Name       string          // name of the Style
	Chars      StyleChars      // characters to use on the progress bar
	Colors     StyleColors     // colors to use on the progress bar
	Options    StyleOptions    // misc. options for the progress bar
	Visibility StyleVisibility // show/hide components of the progress bar(s)
}

var (
	// StyleDefault uses ASCII text to render the Trackers.
	StyleDefault = Style{
		Name:       "StyleDefault",
		Chars:      StyleCharsDefault,
		Colors:     StyleColorsDefault,
		Options:    StyleOptionsDefault,
		Visibility: StyleVisibilityDefault,
	}

	// StyleBlocks uses UNICODE Block Drawing characters to render the Trackers.
	StyleBlocks = Style{
		Name:       "StyleBlocks",
		Chars:      StyleCharsBlocks,
		Colors:     StyleColorsDefault,
		Options:    StyleOptionsDefault,
		Visibility: StyleVisibilityDefault,
	}

	// StyleCircle uses UNICODE Circle runes to render the Trackers.
	StyleCircle = Style{
		Name:       "StyleCircle",
		Chars:      StyleCharsCircle,
		Colors:     StyleColorsDefault,
		Options:    StyleOptionsDefault,
		Visibility: StyleVisibilityDefault,
	}

	// StyleRhombus uses UNICODE Rhombus runes to render the Trackers.
	StyleRhombus = Style{
		Name:       "StyleRhombus",
		Chars:      StyleCharsRhombus,
		Colors:     StyleColorsDefault,
		Options:    StyleOptionsDefault,
		Visibility: StyleVisibilityDefault,
	}
)

// StyleChars defines the characters/strings to use for rendering the Tracker.
type StyleChars struct {
	BoxLeft       string // left-border
	BoxRight      string // right-border
	Finished      string // finished block
	Finished25    string // 25% finished block
	Finished50    string // 50% finished block
	Finished75    string // 75% finished block
	Indeterminate IndeterminateIndicatorGenerator
	Unfinished    string // 0% finished block
}

var (
	// StyleCharsDefault uses simple ASCII characters.
	StyleCharsDefault = StyleChars{
		BoxLeft:       "[",
		BoxRight:      "]",
		Finished:      "#",
		Finished25:    ".",
		Finished50:    ".",
		Finished75:    ".",
		Indeterminate: IndeterminateIndicatorMovingBackAndForth("<#>", DefaultUpdateFrequency/2),
		Unfinished:    ".",
	}

	// StyleCharsBlocks uses UNICODE Block Drawing characters.
	StyleCharsBlocks = StyleChars{
		BoxLeft:       "║",
		BoxRight:      "║",
		Finished:      "█",
		Finished25:    "░",
		Finished50:    "▒",
		Finished75:    "▓",
		Indeterminate: IndeterminateIndicatorMovingBackAndForth("▒█▒", DefaultUpdateFrequency/2),
		Unfinished:    "░",
	}

	// StyleCharsCircle uses UNICODE Circle characters.
	StyleCharsCircle = StyleChars{
		BoxLeft:       "(",
		BoxRight:      ")",
		Finished:      "●",
		Finished25:    "○",
		Finished50:    "○",
		Finished75:    "○",
		Indeterminate: IndeterminateIndicatorMovingBackAndForth("○●○", DefaultUpdateFrequency/2),
		Unfinished:    "◌",
	}

	// StyleCharsRhombus uses UNICODE Rhombus characters.
	StyleCharsRhombus = StyleChars{
		BoxLeft:       "<",
		BoxRight:      ">",
		Finished:      "◆",
		Finished25:    "◈",
		Finished50:    "◈",
		Finished75:    "◈",
		Indeterminate: IndeterminateIndicatorMovingBackAndForth("◈◆◈", DefaultUpdateFrequency/2),
		Unfinished:    "◇",
	}
)

// StyleColors defines what colors to use for various parts of the Progress and
// Tracker texts.
type StyleColors struct {
	Message text.Colors // message text colors
	Error   text.Colors // error text colors
	Percent text.Colors // percentage text colors
	Stats   text.Colors // stats text (time, value) colors
	Time    text.Colors // time text colors (overrides Stats)
	Tracker text.Colors // tracker text colors
	Value   text.Colors // value text colors (overrides Stats)
	Speed   text.Colors // speed text colors
}

var (
	// StyleColorsDefault defines sane color choices - None.
	StyleColorsDefault = StyleColors{}

	// StyleColorsExample defines a few choice color options. Use this is just
	// as an example to customize the Tracker/text colors.
	StyleColorsExample = StyleColors{
		Message: text.Colors{text.FgWhite},
		Error:   text.Colors{text.FgRed},
		Percent: text.Colors{text.FgHiRed},
		Stats:   text.Colors{text.FgHiBlack},
		Time:    text.Colors{text.FgGreen},
		Tracker: text.Colors{text.FgYellow},
		Value:   text.Colors{text.FgCyan},
		Speed:   text.Colors{text.FgMagenta},
	}
)

// StyleOptions defines misc. options to control how the Tracker or its parts
// gets rendered.
type StyleOptions struct {
	DoneString              string                   // "done!" string
	ErrorString             string                   // "error!" string
	ETAPrecision            time.Duration            // precision for ETA
	ETAString               string                   // string for ETA
	Separator               string                   // text between message and tracker
	SnipIndicator           string                   // text denoting message snipping
	PercentFormat           string                   // formatting to use for percentage
	PercentIndeterminate    string                   // when percentage cannot be computed
	TimeDonePrecision       time.Duration            // precision for time when done
	TimeInProgressPrecision time.Duration            // precision for time when in progress
	TimeOverallPrecision    time.Duration            // precision for overall time
	SpeedPrecision          time.Duration            // precision for speed
	SpeedOverallFormatter   func(value int64) string // formatter for the overall tracker speed
}

var (
	// StyleOptionsDefault defines sane defaults for the Options. Use this as an
	// example to customize the Tracker rendering.
	StyleOptionsDefault = StyleOptions{
		DoneString:              "done!",
		ErrorString:             "fail!",
		ETAPrecision:            time.Second,
		ETAString:               "~ETA",
		PercentFormat:           "%5.2f%%",
		PercentIndeterminate:    " ??? ",
		Separator:               " ... ",
		SnipIndicator:           "~",
		TimeDonePrecision:       time.Millisecond,
		TimeInProgressPrecision: time.Microsecond,
		TimeOverallPrecision:    time.Second,
		SpeedPrecision:          time.Microsecond,
		SpeedOverallFormatter:   FormatNumber,
	}
)

// StyleVisibility controls what gets shown and what gets hidden.
type StyleVisibility struct {
	ETA            bool // ETA for each tracker
	ETAOverall     bool // ETA for the overall tracker
	Percentage     bool // tracker progress percentage value
	Time           bool // tracker time taken
	Tracker        bool // tracker ([===========-----------])
	TrackerOverall bool // overall tracker
	Value          bool // tracker value
	Speed          bool // tracker speed
	SpeedOverall   bool // speed for the overall tracker
}

var (
	// StyleVisibilityDefault defines sane defaults for the Visibility.
	StyleVisibilityDefault = StyleVisibility{
		ETA:            false,
		ETAOverall:     true,
		Percentage:     true,
		Time:           true,
		Tracker:        true,
		TrackerOverall: false,
		Value:          true,
		Speed:          true,
		SpeedOverall:   true,
	}
)
