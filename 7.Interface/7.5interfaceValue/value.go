package main

import (
	"bytes"
	"io"
)

func main() {
	test()
}
func test() {
	// var w io.Writer
	// w = os.Stdout
	// w = new(bytes.Buffer)
	// w = nil
	// var b *bytes.Buffer
	var b io.Writer
	if debug {
		b = new(bytes.Buffer)
	}
	f(b)
}

const debug = false

func f(out io.Writer) {
	if out != nil {
		out.Write([]byte("sadasd"))
	}
}

/*
func (b *Buffer) Write(p []byte) (n int, err error) {
	b.lastRead = opInvalid
	m, ok := b.tryGrowByReslice(len(p))//对bytes.Buffer有要求,不能为nil
	if !ok {
		m = b.grow(len(p))
	}
	return copy(b.buf[m:], p), nil
}

type Writer interface {
	Write(p []byte) (n int, err error)
}
*/
