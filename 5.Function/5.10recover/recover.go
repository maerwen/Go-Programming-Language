package main

import "fmt"

func main() {
	i(111)
}
func i(i int) {
	defer func() {
		i := recover()
		fmt.Println(i)
	}()
	panic(i)
}
