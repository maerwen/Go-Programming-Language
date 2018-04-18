package main

import (
	"fmt"
	"os"
)

func main() {
	// fmt.Println(test1(3, 8))
	test2()
}
func test1(x, y int) int { //计算两个整数的最大公约数
	for y != 0 {
		x, y = y, x%y
		// 		8 3
		// 8 3  3 2
		// 3 2  2 1
		// 2 1  1 0
		// 1 0
	}
	return x
}
func test2() { //并行赋值接收表达式返回值
	_, error := os.Open("kkk")
	m := make(map[int]int)
	m[1] = 245
	_, ok := m[1]
	fmt.Println(ok)
	fmt.Println(error.Error())
}
