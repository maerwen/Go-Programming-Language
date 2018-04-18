package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var global *int

func main() {
	test4()
}
func test4() { //逃逸
	x := 1
	global = &x
}
func test3() { //输出命令行参数
	n := flag.Bool("n", true, "new line")       //创建一个新的布尔标识变量
	sep := flag.String("sep", " ", "separator") //创建一个字符串变量
	//使用标识前，必须更新标识变量的默认值
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if *n {
		fmt.Println()
	}
}
func test2() { //指针测试pointer
	x := 1
	p := &x //p指向x
	fmt.Println(*p)
	*p = 2
	fmt.Println(x)
	v := 1
	fmt.Println(incr(&v))
}
func incr(p *int) int {
	*p++
	return *p
}
func test1() { //短变量声明
	var a int = 5
	var b int
	var c = 1
	d := 4
	var e, f = "aaa", 1
	var g, k int
	var x, y string = "", "\"0\""
	// var i int, j bool
	// var x string, y string = "", "\"0\""
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(k)
	fmt.Println(x)
	fmt.Println(y)
	// 短变量声明不需要声明所有左边的变量
	p, q := os.Open(os.Args[1])
	fmt.Fprintf(p, "%d", q)
	// 短变量声明最少声明一个新变量，否则，代码将无法编译通过
	s, q := os.Create(os.Args[1])
	fmt.Fprintf(s, "%d", q)

}
