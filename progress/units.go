package progress

import (
	"fmt"
)

// UnitsNotationPosition determines notation position relative of unit value.
type UnitsNotationPosition int

// Supported unit positions relative to tracker value;
// default: UnitsNotationPositionBefore
const (
	UnitsNotationPositionBefore UnitsNotationPosition = iota
	UnitsNotationPositionAfter
)

// Units defines the "type" of the value being tracked by the Tracker.
type Units struct {
	Formatter        func(value int64) string // default: FormatNumber
	Notation         string
	NotationPosition UnitsNotationPosition // default: UnitsNotationPositionBefore
}

// Sprint prints the value as defined by the Units.
func (tu Units) Sprint(value int64) string {
	formatter := tu.Formatter
	if formatter == nil {
		formatter = FormatNumber
	}

	formattedValue := formatter(value)
	switch tu.NotationPosition {
	case UnitsNotationPositionAfter:
		return formattedValue + tu.Notation
	default: // UnitsNotationPositionBefore
		return tu.Notation + formattedValue
	}
}

var (
	// UnitsDefault doesn't define any units. The value will be treated as any
	// other number.
	UnitsDefault = Units{
		Notation:         "",
		NotationPosition: UnitsNotationPositionBefore,
		Formatter:        FormatNumber,
	}

	// UnitsBytes defines the value as a storage unit. Values will be converted
	// and printed in one of these forms: B, KB, MB, GB, TB, PB
	UnitsBytes = Units{
		Notation:         "",
		NotationPosition: UnitsNotationPositionBefore,
		Formatter:        FormatBytes,
	}

	// UnitsCurrencyDollar defines the value as a Dollar amount. Values will be
	// converted and printed in one of these forms: $x.yz, $x.yzK, $x.yzM,
	// $x.yzB, $x.yzT
	UnitsCurrencyDollar = Units{
		Notation:         "$",
		NotationPosition: UnitsNotationPositionBefore,
		Formatter:        FormatNumber,
	}

	// UnitsCurrencyEuro defines the value as a Euro amount. Values will be
	// converted and printed in one of these forms: ₠x.yz, ₠x.yzK, ₠x.yzM,
	// ₠x.yzB, ₠x.yzT
	UnitsCurrencyEuro = Units{
		Notation:         "₠",
		NotationPosition: UnitsNotationPositionBefore,
		Formatter:        FormatNumber,
	}

	// UnitsCurrencyPound defines the value as a Pound amount. Values will be
	// converted and printed in one of these forms: £x.yz, £x.yzK, £x.yzM,
	// £x.yzB, £x.yzT
	UnitsCurrencyPound = Units{
		Notation:         "£",
		NotationPosition: UnitsNotationPositionBefore,
		Formatter:        FormatNumber,
	}
)

// FormatBytes formats the given value as a "Byte".
func FormatBytes(value int64) string {
	if value < 1000 {
		return fmt.Sprintf("%dB", value)
	} else if value < 1000000 {
		return fmt.Sprintf("%.2fKB", float64(value)/1000.0)
	} else if value < 1000000000 {
		return fmt.Sprintf("%.2fMB", float64(value)/1000000.0)
	} else if value < 1000000000000 {
		return fmt.Sprintf("%.2fGB", float64(value)/1000000000.0)
	} else if value < 1000000000000000 {
		return fmt.Sprintf("%.2fTB", float64(value)/1000000000000.0)
	}
	return fmt.Sprintf("%.2fPB", float64(value)/1000000000000000.0)
}

// FormatNumber formats the given value as a "regular number".
func FormatNumber(value int64) string {
	if value < 1000 {
		return fmt.Sprintf("%d", value)
	} else if value < 1000000 {
		return fmt.Sprintf("%.2fK", float64(value)/1000.0)
	} else if value < 1000000000 {
		return fmt.Sprintf("%.2fM", float64(value)/1000000.0)
	} else if value < 1000000000000 {
		return fmt.Sprintf("%.2fB", float64(value)/1000000000.0)
	} else if value < 1000000000000000 {
		return fmt.Sprintf("%.2fT", float64(value)/1000000000000.0)
	}
	return fmt.Sprintf("%.2fQ", float64(value)/1000000000000000.0)
}
