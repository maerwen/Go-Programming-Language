package main

import (
	"fmt"
)

//吃苹果:20个,每天不能小于1个,每天吃剩下的个数的1/2+1个,问可以吃几天
func main() {
	fmt.Println(eat(20, 0))
}
func eat(count, days int) int {
	if count >= 1 {
		days++
		count -= (count/2 + 1)
		return eat(count, days)
	}
	return days
}
