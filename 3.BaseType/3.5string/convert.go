package main

import (
	"fmt"
	"strconv"
)

func IntToStr1(x int) string {
	return fmt.Sprintf("%d", x)
}
func IntToStr2(x int) string {
	return strconv.Itoa(x)
}
func StrToInt1(s string) int {
	x, _ := strconv.Atoi(s)
	return x
}
func StrToInt2(s string) int {
	x, _ := strconv.ParseInt(s, 10, 64)
	return int(x)
}
