package main

import (
	"fmt"
)

func main() {

}
func arr() {
	type Weekday int
	const (
		Sunday Weekday = iota
		Monday
		Tuesday
	)
	// 这居然也算个数组
	array := [...]string{Sunday: "星期日", Monday: "星期一", Tuesday: "星期二"}
	fmt.Println(Monday, array[Monday])
	r := [...]int{99: -1}
	fmt.Println(r[99])
	fmt.Println(r[1])
}
func zero(ptr *[32]byte) { //数组清零
	// ptr = [32]{}
	for i := range ptr {
		ptr[i] = 0
	}
}
