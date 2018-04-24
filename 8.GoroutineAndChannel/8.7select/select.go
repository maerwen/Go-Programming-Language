package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// print()
	// countdown2()
	// countdown3()
	loop()
}
func countdown3() {
	ch := time.Tick(1 * time.Second)
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
	for i := 10; i > 0; i-- {
		fmt.Println(i)
		select {
		case <-ch:
		case <-abort:
			fmt.Println("发射取消!")
			return
		}
	}
	launch()
}
func countdown2() { //火箭发射倒计时,按键取消
	fmt.Println("10秒后将发射!")
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
	select {
	case <-time.After(10 * time.Second): //意思是多少时间段后执行什么操作,是一次性的执行
	case <-abort:
		fmt.Println("发射取消!")
		return
	}
	launch()

}
func countdown1() { //火箭发射倒计时
	fmt.Println("开始倒计时!")
	tick := time.Tick(1 * time.Second) //返回一个通道,定期发送一个事件
	for t := 10; t > 0; t-- {
		fmt.Println(t)
		<-tick
	}
	launch()
}
func abort() { //按键取消发射
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
}
func launch() {
	fmt.Println("发射!")
}
func print() {
	ch := make(chan int, 1) //容量为1的缓冲通道.如果增加容量,会使输出变得不可确定
	for i := 0; i < 10; i++ {
		select {
		case ch <- i: //偶数时发送
		case x := <-ch: //i是奇数时接收
			fmt.Println(x)
		}
	}

}
func loop() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
	select {
	case <-abort:
	default:
		fmt.Println("loop")
	}
}
