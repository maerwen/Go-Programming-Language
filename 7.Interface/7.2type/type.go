package main

func main() {
	test(1, 2)
}

type Writer interface { //单方法接口
	Write(p []byte) (int, error)
}
type Reader interface {
	Read(p []byte) (int, error)
}
type Closer interface {
	Close() error
}
type ReadWriteCloser interface { //嵌入式接口
	Writer
	Reader
	Closer
}
