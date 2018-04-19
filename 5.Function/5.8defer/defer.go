package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2))
}
func sum(x, y int) int {
	defer func() {
		fmt.Println("---------")
	}()
	return x + y
}
