// bank一个只有一个账户的并发安全银行
package main

import (
	"fmt"
	"time"
)

func main() {
	start()
	for i := 0; i < 100; i++ {
		go func() {
			deposit(100)
		}()
	}
	time.Sleep(3 * time.Second)
	fmt.Println(balance())

}

var deposits = make(chan int) //发送存款额
var balances = make(chan int) //接收余额
func deposit(amount int) {
	deposits <- amount
}
func balance() int {
	return <-balances
}
func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}
func start() {
	go teller() //监控goroutine
}
