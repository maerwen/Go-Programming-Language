package main

import "fmt"

func main() {
	test()
}
func test() { //左移与右移
	x := 4
	fmt.Println(x << 2)
	fmt.Println(x >> 2)
}
