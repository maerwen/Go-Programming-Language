package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond) //新创建一个goroutine
	const n = 46
	fibN := fib(n)
	fmt.Printf("\rFib(%d) = %d\n", n, fibN)
}
func fib(x int) int { //斐波那契数列
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)

}
func spinner(delay time.Duration) { //提示
	for { //死循环
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r) //\r清空本行,光标回到行首.%c字符
			time.Sleep(delay)     //推迟线程执行的时间
		}
	}
}
