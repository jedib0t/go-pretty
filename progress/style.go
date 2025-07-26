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
	Renderer   StyleRenderer   // custom render functions for the progress bar
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
	Pinned  text.Colors // color of the pin message
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
		Pinned:  text.Colors{text.BgHiBlack, text.FgWhite, text.Bold},
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
	DoneString              string         // "done!" string
	ErrorString             string         // "error!" string
	ETAPrecision            time.Duration  // precision for ETA
	ETAString               string         // string for ETA
	Separator               string         // text between message and tracker
	SnipIndicator           string         // text denoting message snipping
	PercentFormat           string         // formatting to use for percentage
	PercentIndeterminate    string         // when percentage cannot be computed
	SpeedPosition           Position       // where speed is displayed in stats
	SpeedPrecision          time.Duration  // precision for speed
	SpeedOverallFormatter   UnitsFormatter // formatter for the overall tracker speed
	SpeedSuffix             string         // suffix (/s)
	TimeDonePrecision       time.Duration  // precision for time when done
	TimeInProgressPrecision time.Duration  // precision for time when in progress
	TimeOverallPrecision    time.Duration  // precision for overall time
}

// StyleOptionsDefault defines sane defaults for the Options. Use this as an
// example to customize the Tracker rendering.
var StyleOptionsDefault = StyleOptions{
	DoneString:              "done!",
	ErrorString:             "fail!",
	ETAPrecision:            time.Second,
	ETAString:               "~ETA",
	PercentFormat:           "%5.2f%%",
	PercentIndeterminate:    " ??? ",
	Separator:               " ... ",
	SnipIndicator:           "~",
	SpeedPosition:           PositionRight,
	SpeedPrecision:          time.Microsecond,
	SpeedOverallFormatter:   FormatNumber,
	SpeedSuffix:             "/s",
	TimeDonePrecision:       time.Millisecond,
	TimeInProgressPrecision: time.Microsecond,
	TimeOverallPrecision:    time.Second,
}

// StyleVisibility controls what gets shown and what gets hidden.
type StyleVisibility struct {
	ETA            bool // ETA for each tracker
	ETAOverall     bool // ETA for the overall tracker
	Percentage     bool // tracker progress percentage value
	Pinned         bool // pin message
	Speed          bool // tracker speed
	SpeedOverall   bool // overall tracker speed
	Time           bool // tracker time taken
	Tracker        bool // tracker ([===========-----------])
	TrackerOverall bool // overall tracker
	Value          bool // tracker value
}

// StyleVisibilityDefault defines sane defaults for the Visibility.
var StyleVisibilityDefault = StyleVisibility{
	ETA:            false,
	ETAOverall:     true,
	Percentage:     true,
	Pinned:         true,
	Speed:          false,
	SpeedOverall:   false,
	Time:           true,
	Tracker:        true,
	TrackerOverall: false,
	Value:          true,
}

type StyleRenderer struct {
	// TrackerDeterminate will override how the progress bar is rendered.
	// value is the current value of the tracker out of total.
	// maxLen is the number of characters available for the progress bar.
	// return the complete progress bar string. E.g. [===----]
	TrackerDeterminate func(value int64, total int64, maxLen int) string

	// TrackerIndeterminate will override how the indeterminate progress bar is rendered.
	// maxLen is the number of characters available for the progress bar.
	// return the complete progress bar string. E.g. [<#>----]
	TrackerIndeterminate func(maxLen int) string
}
