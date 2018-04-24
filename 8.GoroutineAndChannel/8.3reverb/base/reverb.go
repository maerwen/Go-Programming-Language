package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	handle()
}
func handle() { //测试最基础的回声程序
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		// echo(input.Text(), 3*time.Second)
		go echo(input.Text(), 3*time.Second)
	}
}
func echo(s string, delay time.Duration) { //回声发送
	time.Sleep(delay)
	fmt.Printf("\t%s\n", strings.ToUpper(s))
	time.Sleep(delay)
	fmt.Printf("\t%s\n", s)
	time.Sleep(delay)
	fmt.Printf("\t%s\n", strings.ToLower(s))
}
