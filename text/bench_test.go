package text

import (
	"testing"
)

var (
	benchStrPlain   = "The quick brown fox jumped over the lazy dog. 生命、宇宙、そして万物についての疑問"
	benchStrEscaped = "\x1b[33mThe quick \x1b[1mbrown fox\x1b[22m jumped over the \x1b[4mlazy dog\x1b[24m.\x1b[0m 生命、宇宙"

	benchResultInt int
	benchResultStr string
)

func BenchmarkAlign_Apply(b *testing.B) {
	for _, bc := range []struct {
		name  string
		align Align
	}{
		{"Left", AlignLeft},
		{"Center", AlignCenter},
		{"Right", AlignRight},
	} {
		b.Run(bc.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				benchResultStr = bc.align.Apply("Jon Snow", 27)
			}
		})
	}
}

func BenchmarkEscSeqParser_Consume(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		esp := EscSeqParser{}
		for _, char := range benchStrEscaped {
			esp.Consume(char)
		}
	}
}

func BenchmarkStringWidthWithoutEscSequences(b *testing.B) {
	for _, bc := range []struct {
		name string
		str  string
	}{
		{"Plain", benchStrPlain},
		{"Escaped", benchStrEscaped},
	} {
		b.Run(bc.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				benchResultInt = StringWidthWithoutEscSequences(bc.str)
			}
		})
	}
}
