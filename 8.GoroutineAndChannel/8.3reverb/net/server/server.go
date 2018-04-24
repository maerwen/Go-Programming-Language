package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	testPlus()
}
func testPlus() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		c, e := listener.Accept()
		if e != nil {
			log.Fatal(e)
			continue
		}
		defer c.Close()
		handleConn(c)
	}
}
func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		s := input.Text()
		// echo(c, s, 3*time.Second) //同时只能处理最多一个回声
		go echo(c, s, 3*time.Second) //同时可以处理多个回声
	}
}
func echo(c io.Writer, s string, delay time.Duration) { //处理输出回声
	fmt.Printf("new voice:\t%s\n", s)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToUpper(s))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", s)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(s))
}
