package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	test()
}
func Fprintf(w io.Writer, format string, args ...interface{}) (int, error) { //模拟
	return fmt.Fprintf(w, format, args)
}
func Printf(format string, args ...interface{}) (int, error) {
	return Fprintf(os.Stdout, format, args)
}
func Sprintf(format string, args ...interface{}) string {
	var b bytes.Buffer
	Fprintf(&b, format, args)
	return b.String()
}

type Writer interface { //简单的接口
	Write(p []byte) (int, error)
}
type ByteCounter int

func (bc *ByteCounter) Write(p []byte) (int, error) {
	*bc += ByteCounter(len(p))
	return len(p), nil
}
func test() {
	var bc ByteCounter
	bc.Write([]byte("hello"))
	fmt.Println(bc)
	Fprintf(&bc, "hello,%s", "Dolly")
	fmt.Println(bc)
}
