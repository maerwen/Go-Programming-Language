package main

import (
	"fmt"
)

func main() {
	fmt.Println(f == nil) //true
	fmt.Printf("%T\n%T\n%T\n", f, f1, f2)
	//func(int) int
	//func(int)
	//func()
}

var f func(int) int //nil
var f1 = func(i int) {}

func func2() {}

var f2 = func2
