package main

import "fmt"

func main() {
	test()
}
func test() { //对一系列自增的数字求平方并打印
	naturals := make(chan int)
	squares := make(chan int)
	go func() {
		for i := 0; i < 999; i++ {
			naturals <- i
		}
		close(naturals)
	}()
	go func() {
		for {
			i, ok := <-naturals
			if ok {
				squares <- i * i
				continue
			}
			break
		}
		close(squares)
	}()
	for i := range squares {
		fmt.Printf("%d\t", i)
	}
}
