package main

import (
	"fmt"
)

var arr = [...]int{1, 2, 3, 4, 5, 6}

func length() {
	s := arr[1:3]
	fmt.Println(len(s)) //2
	fmt.Println(cap(s)) //5
	// fmt.Println(s[3])//宕机
	sl := s[:5]
	fmt.Println(sl)     //在原数组范围内拓展了slice
	fmt.Println(len(s)) //2
	fmt.Println(cap(s)) //5
}
func reverse(arr []int) []int { //反转一个slice
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}
func moveLeft(arr []int, a int) []int { //将一个slice左移多少位
	reverse(arr[:a])
	fmt.Println(arr)
	reverse(arr[a:])
	fmt.Println(arr)
	return reverse(arr)
}
func equal(x, y []int) bool { //判断两个slice\是否相等
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
func isNil() { //判断一个slice是否为nil,并打印长度
	var a []int
	fmt.Printf("%d\t%t\n", len(a), a == nil)
	a = nil
	fmt.Printf("%d\t%t\n", len(a), a == nil)
	a = []int(nil)
	fmt.Printf("%d\t%t\n", len(a), a == nil)
	a = []int{}
	fmt.Printf("%d\t%t\n", len(a), a == nil)
}
