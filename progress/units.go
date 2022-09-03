package progress

import (
	"fmt"
)

// UnitsNotationPosition determines notation position relative to unit value.
type UnitsNotationPosition int

// Supported unit positions relative to tracker value;
// default: UnitsNotationPositionBefore
const (
	UnitsNotationPositionBefore UnitsNotationPosition = iota
	UnitsNotationPositionAfter
)

// UnitsFormatter defines a function that prints a value in a specific style.
type UnitsFormatter func(value int64) string

// Units defines the "type" of the value being tracked by the Tracker.
type Units struct {
	Formatter        UnitsFormatter // default: FormatNumber
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
	return formatNumber(value, map[int64]string{
		1000000000000000: "PB",
		1000000000000:    "TB",
		1000000000:       "GB",
		1000000:          "MB",
		1000:             "KB",
		0:                "B",
	})
}

// FormatNumber formats the given value as a "regular number".
func FormatNumber(value int64) string {
	return formatNumber(value, map[int64]string{
		1000000000000000: "Q",
		1000000000000:    "T",
		1000000000:       "B",
		1000000:          "M",
		1000:             "K",
		0:                "",
	})
}

var (
	unitScales = []int64{
		1000000000000000,
		1000000000000,
		1000000000,
		1000000,
		1000,
	}
)

func formatNumber(value int64, notations map[int64]string) string {
	for _, unitScale := range unitScales {
		if value >= unitScale {
			return fmt.Sprintf("%.2f%s", float64(value)/float64(unitScale), notations[unitScale])
		}
	}
	return fmt.Sprintf("%d%s", value, notations[0])
}
