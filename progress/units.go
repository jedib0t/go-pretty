package progress

import (
	"fmt"
)

// Units defines the "type" of the value being tracked by the Tracker.
type Units struct {
	Notation  string
	Formatter func(value int64) string
}

var (
	// UnitsDefault doesn't define any units. The value will be treated as any
	// other number.
	UnitsDefault = Units{
		Notation:  "",
		Formatter: FormatNumber,
	}

	// UnitsBytes defines the value as a storage unit. Values will be converted
	// and printed in one of these forms: B, KB, MB, GB, TB, PB
	UnitsBytes = Units{
		Notation:  "",
		Formatter: FormatBytes,
	}

	// UnitsCurrencyDollar defines the value as a Dollar amount. Values will be
	// converted and printed in one of these forms: $x.yz, $x.yzK, $x.yzM,
	// $x.yzB, $x.yzT
	UnitsCurrencyDollar = Units{
		Notation:  "$",
		Formatter: FormatNumber,
	}

	// UnitsCurrencyEuro defines the value as a Euro amount. Values will be
	// converted and printed in one of these forms: ₠x.yz, ₠x.yzK, ₠x.yzM,
	// ₠x.yzB, ₠x.yzT
	UnitsCurrencyEuro = Units{
		Notation:  "₠",
		Formatter: FormatNumber,
	}

	// UnitsCurrencyPound defines the value as a Pound amount. Values will be
	// converted and printed in one of these forms: £x.yz, £x.yzK, £x.yzM,
	// £x.yzB, £x.yzT
	UnitsCurrencyPound = Units{
		Notation:  "£",
		Formatter: FormatNumber,
	}
)

// Sprint prints the value as defined by the Units.
func (tu Units) Sprint(value int64) string {
	if tu.Formatter == nil {
		return tu.Notation + FormatNumber(value)
	}
	return tu.Notation + tu.Formatter(value)
}

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
