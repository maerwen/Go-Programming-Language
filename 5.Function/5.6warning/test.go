package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		a := i
		defer write(a)
		//  write(i)
	}
}
func write(i int) {
	// time.Sleep(time.Second * 2)
	fmt.Println(i)
}
