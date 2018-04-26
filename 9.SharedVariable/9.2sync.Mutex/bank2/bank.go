package main

import "sync"

// 利用sync.Mutex来实现
func main() {

}

var (
	mu      sync.Mutex
	balance int
)

func deposits(amount int) {
	mu.Lock()
	defer mu.Unlock()
	balance += amount
}
func balances() int {
	mu.Lock()
	defer mu.Unlock()
	return balance
}
