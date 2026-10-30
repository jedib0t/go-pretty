package progress

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatBytes(t *testing.T) {
	assert.Equal(t, "1B", FormatBytes(1))
	assert.Equal(t, "1.50KB", FormatBytes(1500))
	assert.Equal(t, "1.50MB", FormatBytes(1500000))
	assert.Equal(t, "1.50GB", FormatBytes(1500000000))
	assert.Equal(t, "1.50TB", FormatBytes(1500000000000))
	assert.Equal(t, "1.50PB", FormatBytes(1500000000000000))
	assert.Equal(t, "1500.00PB", FormatBytes(1500000000000000000))
	assert.Equal(t, "999B", FormatBytes(999))
	assert.Equal(t, "999.99KB", FormatBytes(999994))
	assert.Equal(t, "1.00MB", FormatBytes(999995))
	assert.Equal(t, "1.00MB", FormatBytes(999999))
	assert.Equal(t, "1.00GB", FormatBytes(999999999))
}

func TestFormatNumber(t *testing.T) {
	assert.Equal(t, "1", FormatNumber(1))
	assert.Equal(t, "1.50K", FormatNumber(1500))
	assert.Equal(t, "1.50M", FormatNumber(1500000))
	assert.Equal(t, "1.50B", FormatNumber(1500000000))
	assert.Equal(t, "1.50T", FormatNumber(1500000000000))
	assert.Equal(t, "1.50Q", FormatNumber(1500000000000000))
	assert.Equal(t, "1500.00Q", FormatNumber(1500000000000000000))
	assert.Equal(t, "1.00M", FormatNumber(999999))
	assert.Equal(t, "1.00B", FormatNumber(999999999))
	assert.Equal(t, "-1", FormatNumber(-1))
	assert.Equal(t, "-1.50K", FormatNumber(-1500))
	assert.Equal(t, "-1.50M", FormatNumber(-1500000))
	assert.Equal(t, "-1.00M", FormatNumber(-999999))
	assert.Equal(t, "-9223.37Q", FormatNumber(math.MinInt64))
}

func TestUnits_Sprint(t *testing.T) {
	assert.Equal(t, "1.50K", UnitsDefault.Sprint(1500))
	assert.Equal(t, "1.50KB", UnitsBytes.Sprint(1500))
	assert.Equal(t, "$1.50K", UnitsCurrencyDollar.Sprint(1500))
	assert.Equal(t, "₠1.50K", UnitsCurrencyEuro.Sprint(1500))
	assert.Equal(t, "£1.50K", UnitsCurrencyPound.Sprint(1500))

	customUnits := Units{Notation: "#"}
	assert.Equal(t, "#1.50K", customUnits.Sprint(1500))
}

func TestUnits_NotationPosition(t *testing.T) {
	afterUnits := Units{Notation: " ₽", NotationPosition: UnitsNotationPositionAfter}
	assert.Equal(t, "1.50K ₽", afterUnits.Sprint(1500))

	unknownNotationPosition := Units{Notation: "* ", NotationPosition: UnitsNotationPosition(999)}
	assert.Equal(t, "* 1.50K", unknownNotationPosition.Sprint(1500))
}
