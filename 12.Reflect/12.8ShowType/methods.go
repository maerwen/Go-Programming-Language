package main

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

func main() {
	t := time.Hour
	Print(t)
	// r := new(http.Request)
	// Print(r)
}
func Print(x interface{}) {
	v := reflect.ValueOf(x)
	t := v.Type()
	fmt.Printf("type\t%s\n", t)
	for i := 0; i < v.NumMethod(); i++ {
		methType := v.Method(i).Type()
		fmt.Printf("func (%s) %s %s\n", t, t.Method(i).Name, strings.TrimPrefix(methType.String(), "func"))
		// fmt.Println(methType.String()) //参数列表与返回值
	}
}
