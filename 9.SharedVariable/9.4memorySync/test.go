package main

import "fmt"

func main() {
	go func() {
		x = 1
		ch <- fmt.Sprint("y:", y, "")
	}()
	go func() {
		y = 1
		ch <- fmt.Sprint("x:", x, "")
	}()
loop:
	for {
		select {
		case s, ok := <-ch:
			if !ok {
				break loop
			}
			fmt.Println(s)
		default:
		}
	}
}

var ch = make(chan string, 2)
var x, y int
