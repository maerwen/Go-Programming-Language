package main

import (
	"fmt"
	"reflect"
)

func main() {
	test()
}
func test() {
	x := 2
	a := reflect.ValueOf(2)
	b := reflect.ValueOf(x)
	c := reflect.ValueOf(&x)
	d := c.Elem()
	fmt.Println(a.CanAddr()) //false
	fmt.Println(b.CanAddr()) //false
	fmt.Println(c.CanAddr()) //false
	fmt.Println(d.CanAddr()) //true
}
