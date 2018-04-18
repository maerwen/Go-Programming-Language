package main

import (
	"fmt"
)

func testAppend() {
	var runes []rune
	for _, a := range "你好,小兄弟!" {
		runes = append(runes, a)
	}
	fmt.Println(runes)        //[20320 22909 44 23567 20804 24351 33]
	fmt.Printf("%c\n", runes) //[你 好 , 小 兄 弟 !]
	fmt.Printf("%q\n", runes) //['你' '好' ',' '小' '兄' '弟' '!']

	fmt.Printf("%c\n", append(runes, 's', 'a', 'd', 'f', 'o', 'x'))

}
func appendInt(a []int, b int) []int { //往一个int数组slice添加元素
	// slice仍有增长空间,拓展slice
	var c []int
	newLen := len(a) + 1
	if newLen <= cap(a) {
		c = a[:newLen]
	} else { // slice空间不足,新分配底层数组,容量翻倍
		newCap := newLen
		if newCap < len(a)*2 {
			newCap = 2 * len(a)
		}
		c = make([]int, newLen, newCap)
		copy(c, a)
	}
	c[len(a)] = b
	return c
}
func testAppendInt() {
	// 定义两个数组,一个用来接收appendint的返回值,另一个用来存储并更新数组
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d\tcap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}
func appendIntPlus(x []int, y ...int) []int { //接收可变参数列表
	var z []int
	a := len(x)
	b := len(y)
	if a+b <= cap(x) {
		z = x[:a+b]
	} else {
		newCap := a + b
		if newCap < 2*a {
			newCap = 2 * a
		}
		z = make([]int, len(x)+len(y), newCap)
		copy(z[0:len(x)], x)
	}
	copy(z[len(x):], y)
	return z
}
