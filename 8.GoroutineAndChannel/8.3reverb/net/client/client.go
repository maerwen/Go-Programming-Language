package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	dial()
}
func dial() {
	c, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	defer c.Close()
	// go mustCopy(os.Stdout, c)两个线程,一个负责输入,一个负责输出.与下面无区别
	// mustCopy(c, os.Stdin)
	go mustCopy(c, os.Stdin) //当主goroutine从标准输入读取并发送到服务器的时候,第二个goroutine读取服务器的回复并且输出.
	mustCopy(os.Stdout, c)
}
func mustCopy(dest io.Writer, src io.Reader) { //消息传输
	_, err := io.Copy(dest, src)
	if err != nil {
		log.Fatal(err)
	}
}
