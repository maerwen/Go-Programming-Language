package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	test5()
}
func test1() {
	x := 2
	a := reflect.ValueOf(2)  //	整数2的一个副本
	b := reflect.ValueOf(x)  //	x的一个副本
	c := reflect.ValueOf(&x) //	指针&x的一个副本
	d := c.Elem()            //	对c的指针提领得来,可寻址
	fmt.Println(a.CanAddr()) //false
	fmt.Println(b.CanAddr()) //false
	fmt.Println(c.CanAddr()) //false
	fmt.Println(d.CanAddr()) //true
}
func test2() {
	x := 2
	d := reflect.ValueOf(&x).Elem()
	px := d.Addr().Interface().(*int)
	*px = 3
	fmt.Println(x)
}
func test3() {
	x := 2
	d := reflect.ValueOf(&x).Elem()
	d.Set(reflect.ValueOf(4))
	fmt.Println(x)
}
func test4() {
	x := 2
	// a := reflect.ValueOf(x)
	// x.Set(reflect.ValueOf(4))//错误
	b := reflect.ValueOf(&x).Elem()
	// b.set(reflect.ValueOf(int64(5)))//错误
	b.SetInt(2)

	var y interface{}
	ry := reflect.ValueOf(&y).Elem()
	// ry.SetInt(2) //错误
	ry.Set(reflect.ValueOf(2))
	fmt.Println(y)
	ry.Set(reflect.ValueOf("fsdf"))
	fmt.Println(y)
}
func test5() {
	stdout := reflect.ValueOf(os.Stdout).Elem()
	fmt.Println(stdout.Type())
	fd := stdout.FieldByName("fd")
	// fd.SetInt(4)
	fmt.Println(fd.CanAddr(), fd.CanSet()) //是否可寻址	是否可改变
}
