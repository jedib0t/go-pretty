package text

import (
	"regexp"
	"strings"
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWrapHyperlinks(t *testing.T) {
	for _, terminator := range []string{"\x1b\\", "\a"} {
		for _, url := range []string{"https://a.test/x", "https://example.com/x", "https://a.test/a\\b?q=x;y", "https://a.test/a b/世界"} {
			open := "\x1b]8;id=example;" + url + terminator
			close := "\x1b]8;;" + terminator
			link := func(s string) string { return open + s + close }
			for _, wrap := range []struct {
				name string
				fn   func(string, int) string
			}{{"hard", WrapHard}, {"soft", WrapSoft}, {"text", WrapText}} {
				t.Run(wrap.name+"/"+url+"/"+terminator, func(t *testing.T) {
					assert.Equal(t, link("abcd")+"\n"+link("efgh")+"\n"+link("ij"), wrap.fn(link("abcdefghij"), 4))
					spaced := link("abc ") + "\n" + link("def ") + "\n" + link("ghi")
					if wrap.name == "soft" {
						spaced = link("abc") + " \n" + link("def") + " \n" + link("ghi")
					}
					assert.Equal(t, spaced, wrap.fn(link("abc def ghi"), 4))
					assert.Equal(t, link("abcd")+"\n"+"efgh", wrap.fn(link("abcd")+"efgh", 4))
					assert.Equal(t, "abcd\n"+link("efgh"), wrap.fn("abcd"+link("efgh"), 4))
					assert.Equal(t, "\x1b[31m"+link("abcd")+EscapeReset+"\n\x1b[31m"+link("efgh")+EscapeReset,
						wrap.fn("\x1b[31m"+link("abcdefgh")+EscapeReset, 4))
					assert.Equal(t, link("ab\x1b[1mcd")+EscapeReset+"\n\x1b[1m"+link("ef\x1b[22mgh"),
						wrap.fn(link("ab\x1b[1mcdef\x1b[22mgh"), 4))
					assert.Equal(t, link("abcd"), wrap.fn(link("abcd"), 4))
					assert.NotContains(t, wrap.fn(link("abc def ghi"), 4), "\x1b[8m")
				})
			}
		}
	}
}

func TestWrapHyperlinkPaddingAndEmptyLabel(t *testing.T) {
	open, close := "\x1b]8;;https://a.test/\x1b\\", "\x1b]8;;\x1b\\"
	assert.Equal(t, open+"ab"+close+"  \n"+open+"cdef"+close,
		WrapSoft(open+"ab cdef"+close, 4))
	for _, wrap := range []func(string, int) string{WrapHard, WrapSoft, WrapText} {
		assert.Equal(t, open+close+"abcd\nefgh", wrap(open+close+"abcdefgh", 4))
		assert.Equal(t, "abcd\nefgh"+open+close, wrap("abcdefgh"+open+close, 4))
		assert.Equal(t, open+"abcd"+close+"\n"+open+"efgh"+close, wrap(open+"abcdefgh   "+close, 4))
		assert.Equal(t, "abcd\n"+open+close+"efgh", wrap("abcd "+open+close+" efgh", 4))
	}
}

func TestWrapHyperlinkTargets(t *testing.T) {
	// Read only real OSC 8 and SGR sequences, independently of EscSeqParser.
	sequences := regexp.MustCompile("\x1b(?:\\]8;[^;]*;([^\x1b\a]*)(?:\a|\x1b\\\\)|\\[[0-9;]*m)")
	for _, end := range []string{"\a", "\x1b\\"} {
		for _, target := range []string{"https://a.test/a b/世界?q=x;y", "https://a.test/" + strings.Repeat("m", 2100)} {
			open, close := "\x1b]8;id=target;"+target+end, "\x1b]8;;"+end
			for _, label := range []string{"abcdefghij", "ab cdef ghij", "abcd\nefgh\n\nijkl"} {
				for _, width := range []int{1, 2, 3, 4, 5} {
					for _, wrap := range []func(string, int) string{WrapHard, WrapSoft, WrapText} {
						result := wrap(open+label+close+"!", width)
						active, cursor, visible := "", 0, ""
						consume := func(text string) {
							for _, r := range text {
								if r == '\n' {
									require.Empty(t, active, "link crosses generated line boundary")
								}
								if unicode.IsSpace(r) {
									continue
								}
								if r == '!' {
									require.Empty(t, active, "suffix became linked")
								} else {
									require.Equal(t, target, active)
								}
								visible += string(r)
							}
						}
						for _, indices := range sequences.FindAllStringSubmatchIndex(result, -1) {
							consume(result[cursor:indices[0]])
							if indices[2] >= 0 {
								active = result[indices[2]:indices[3]]
							}
							cursor = indices[1]
						}
						consume(result[cursor:])
						require.Empty(t, active)
						expected := strings.Map(func(r rune) rune {
							if unicode.IsSpace(r) {
								return -1
							}
							return r
						}, label) + "!"
						require.Equal(t, expected, visible)
						require.NotContains(t, result, "\x1b[8m")
					}
				}
			}
		}
	}
}

func TestWrapHyperlinkTabs(t *testing.T) {
	open, close := "\x1b]8;id=a\tb;https://a.test/path\x1b\\", "\x1b]8;;\x1b\\"
	for _, wrap := range []func(string, int) string{WrapHard, WrapSoft, WrapText} {
		assert.Equal(t, open+"ab    cd"+close, wrap(open+"ab\tcd"+close, 10))
		assert.Contains(t, wrap(open+"abcdefgh"+close, 4), open)
		assert.NotContains(t, wrap(open+"abcdefgh"+close, 4), "id=a    b")
	}
}

func TestWrapAdjacentHyperlinks(t *testing.T) {
	left, right := "\x1b]8;;https://left.test/\x1b\\", "\x1b]8;;https://right.test/\x1b\\"
	close := "\x1b]8;;\x1b\\"
	for _, wrap := range []func(string, int) string{WrapHard, WrapSoft, WrapText} {
		input := left + "ab" + close + right + "cdef" + close + "gh"
		expected := left + "ab" + close + right + "cd" + close + "\n" + right + "ef" + close + "gh"
		assert.Equal(t, expected, wrap(input, 4))
	}
}

func TestWrapHyperlinkParagraphs(t *testing.T) {
	open, close := "\x1b]8;;https://a.test/\x1b\\", "\x1b]8;;\x1b\\"
	input := open + "abcdef\nghijkl" + close
	assert.Equal(t, open+"abcd"+close+"\n"+open+"ef"+close+"\n"+open+"ghij"+close+"\n"+open+"kl"+close, WrapText(input, 4))
	for _, wrap := range []func(string, int) string{WrapHard, WrapSoft} {
		result := wrap(open+"abcdefgh\n\nijklmnop"+close, 4)
		assert.Equal(t, "abcd\nefgh\n\nijkl\nmnop", StripEscape(result))
		for _, line := range strings.Split(result, "\n") {
			if line != "" {
				assert.True(t, strings.HasSuffix(line, close), "%q", line)
			}
		}
	}
}
