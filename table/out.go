package table

import (
	"fmt"
	"io"
	"strings"
)

type outputWriter interface {
	Len() int
	Grow(n int)
	Reset()
	String() string
	WriteRune(r rune) (int, error)
	WriteString(s string) (int, error)
}

func newOutputWriter(debug io.Writer) outputWriter {
	return &ow{
		debug: debug,
		out:   strings.Builder{},
	}
}

type ow struct {
	debug io.Writer
	out   strings.Builder
}

func (o *ow) Len() int {
	return o.out.Len()
}

func (o *ow) Grow(n int) {
	if o.debug != nil {
		_, _ = o.debug.Write([]byte(fmt.Sprintf(">> grow buffer by %d bytes\n", n)))
	}
	o.out.Grow(n)
}

func (o *ow) Reset() {
	o.out.Reset()
}

func (o *ow) String() string {
	return o.out.String()
}

func (o *ow) WriteRune(r rune) (int, error) {
	if o.debug != nil {
		_, _ = o.debug.Write([]byte(fmt.Sprintf("++ [%02d] %#v\n", 1, r)))
	}
	return o.out.WriteRune(r)
}

func (o *ow) WriteString(s string) (int, error) {
	if o.debug != nil {
		_, _ = o.debug.Write([]byte(fmt.Sprintf("++ [%02d] %#v\n", len(s), s)))
	}
	return o.out.WriteString(s)
}
