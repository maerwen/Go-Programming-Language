package main

import (
	"fmt"
)

func main() {
	// length()
	// fmt.Println(reverse([]int{1, 2, 3, 5, 7}))
	// s := []int{1, 2, 3, 5, 7}
	// a := s[1:3]
	// b := s[2:4]
	// fmt.Println(equal(a, b))
	// fmt.Println(moveLeft(s, 3))
	// isNil()
	// testAppend()
	// testAppendInt()
	// fmt.Println(appendIntPlus([]int{1, 2, 3}, 1, 2, 3))
	fmt.Println(nonempty1([]string{"1", "2", "", "3"}))
	fmt.Println(nonempty2([]string{"1", "2", "", "3"}))
	fmt.Println(remove1([]int{1, 2, 3, 4}, 1))
	fmt.Println(remove2([]int{1, 2, 3, 4}, 1))
}
