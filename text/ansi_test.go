package text

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEscape(t *testing.T) {
	assert.Equal(t, "\x1b[91mGhost\x1b[0m", Escape("Ghost", FgHiRed.EscapeSeq()))
	assert.Equal(t, "\x1b[94mGhost\x1b[0m\x1b[91mLady\x1b[0m", Escape(FgHiBlue.Sprint("Ghost")+"Lady", FgHiRed.EscapeSeq()))
	assert.Equal(t, "\x1b[91mNymeria\x1b[94mGhost\x1b[0m\x1b[91mLady\x1b[0m", Escape("Nymeria"+FgHiBlue.Sprint("Ghost")+"Lady", FgHiRed.EscapeSeq()))
	assert.Equal(t, "\x1b[91mNymeria \x1b[94mGhost\x1b[0m\x1b[91m Lady\x1b[0m", Escape("Nymeria "+FgHiBlue.Sprint("Ghost")+" Lady", FgHiRed.EscapeSeq()))
}

func ExampleEscape() {
	fmt.Printf("Escape(%#v, %#v) == %#v\n", "Ghost", "", Escape("Ghost", ""))
	fmt.Printf("Escape(%#v, %#v) == %#v\n", "Ghost", FgHiRed.EscapeSeq(), Escape("Ghost", FgHiRed.EscapeSeq()))
	fmt.Printf("Escape(%#v, %#v) == %#v\n", FgHiBlue.Sprint("Ghost")+"Lady", FgHiRed.EscapeSeq(), Escape(FgHiBlue.Sprint("Ghost")+"Lady", FgHiRed.EscapeSeq()))
	fmt.Printf("Escape(%#v, %#v) == %#v\n", "Nymeria"+FgHiBlue.Sprint("Ghost")+"Lady", FgHiRed.EscapeSeq(), Escape("Nymeria"+FgHiBlue.Sprint("Ghost")+"Lady", FgHiRed.EscapeSeq()))
	fmt.Printf("Escape(%#v, %#v) == %#v\n", "Nymeria "+FgHiBlue.Sprint("Ghost")+" Lady", FgHiRed.EscapeSeq(), Escape("Nymeria "+FgHiBlue.Sprint("Ghost")+" Lady", FgHiRed.EscapeSeq()))

	// Output: Escape("Ghost", "") == "Ghost"
	// Escape("Ghost", "\x1b[91m") == "\x1b[91mGhost\x1b[0m"
	// Escape("\x1b[94mGhost\x1b[0mLady", "\x1b[91m") == "\x1b[94mGhost\x1b[0m\x1b[91mLady\x1b[0m"
	// Escape("Nymeria\x1b[94mGhost\x1b[0mLady", "\x1b[91m") == "\x1b[91mNymeria\x1b[94mGhost\x1b[0m\x1b[91mLady\x1b[0m"
	// Escape("Nymeria \x1b[94mGhost\x1b[0m Lady", "\x1b[91m") == "\x1b[91mNymeria \x1b[94mGhost\x1b[0m\x1b[91m Lady\x1b[0m"
}

func TestStripEscape(t *testing.T) {
	assert.Equal(t, "Ghost", StripEscape(FgHiRed.Sprint("Ghost")))
	assert.Equal(t, "GhostLady", StripEscape(FgHiBlue.Sprint("Ghost")+"Lady"))
	assert.Equal(t, "NymeriaGhostLady", StripEscape("Nymeria"+FgHiBlue.Sprint("Ghost")+"Lady"))
	assert.Equal(t, "Nymeria Ghost Lady", StripEscape("Nymeria "+FgHiBlue.Sprint("Ghost")+" Lady"))
}

func ExampleStripEscape() {
	fmt.Printf("StripEscape(%#v) == %#v\n", "Ghost", StripEscape("Ghost"))
	fmt.Printf("StripEscape(%#v) == %#v\n", "\x1b[91mGhost\x1b[0m", StripEscape("\x1b[91mGhost\x1b[0m"))
	fmt.Printf("StripEscape(%#v) == %#v\n", "\x1b[94mGhost\x1b[0m\x1b[91mLady\x1b[0m", StripEscape("\x1b[94mGhost\x1b[0m\x1b[91mLady\x1b[0m"))
	fmt.Printf("StripEscape(%#v) == %#v\n", "\x1b[91mNymeria\x1b[94mGhost\x1b[0m\x1b[91mLady\x1b[0m", StripEscape("\x1b[91mNymeria\x1b[94mGhost\x1b[0m\x1b[91mLady\x1b[0m"))
	fmt.Printf("StripEscape(%#v) == %#v\n", "\x1b[91mNymeria \x1b[94mGhost\x1b[0m\x1b[91m Lady\x1b[0m", StripEscape("\x1b[91mNymeria \x1b[94mGhost\x1b[0m\x1b[91m Lady\x1b[0m"))

	// Output: StripEscape("Ghost") == "Ghost"
	// StripEscape("\x1b[91mGhost\x1b[0m") == "Ghost"
	// StripEscape("\x1b[94mGhost\x1b[0m\x1b[91mLady\x1b[0m") == "GhostLady"
	// StripEscape("\x1b[91mNymeria\x1b[94mGhost\x1b[0m\x1b[91mLady\x1b[0m") == "NymeriaGhostLady"
	// StripEscape("\x1b[91mNymeria \x1b[94mGhost\x1b[0m\x1b[91m Lady\x1b[0m") == "Nymeria Ghost Lady"
}
