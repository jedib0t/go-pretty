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

func TestStripEscape_Hyperlinks(t *testing.T) {
	for _, test := range []struct {
		name  string
		input string
		want  string
	}{
		{"ST terminator", Hyperlink("https://go.dev/Docs", "Go docs"), "Go docs"},
		{"URL containing m", Hyperlink("https://example.com/Docs", "Go docs"), "Go docs"},
		{"BEL terminator", "\x1b]8;;https://go.dev/Docs\aGo docs\x1b]8;;\a", "Go docs"},
		{"surrounding text", "See " + Hyperlink("https://go.dev/Docs", "Go docs") + " here", "See Go docs here"},
		{"adjacent links", Hyperlink("https://go.dev/Docs", "Go") + Hyperlink("https://go.dev/Tour", " tour"), "Go tour"},
		{"colored label", Hyperlink("https://go.dev/Docs", "\x1b[31mGo docs\x1b[0m"), "Go docs"},
		{"Unicode label", Hyperlink("https://go.dev/Docs", "文档"), "文档"},
		{"backslash in URL", Hyperlink("C:\\Windows\\Docs", "Docs"), "Docs"},
	} {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, StripEscape(test.input))
		})
	}
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
