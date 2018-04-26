package main

import "fmt"

func main() {
	var baked = make(chan *Cake)
	var iced = make(chan *Cake)
	go baker(baked)
	go icer(baked, iced)
	for cake := range iced {
		fmt.Println(cake.state)
	}
}

// 声明一个cake类型
type Cake struct {
	state string
}

// 定义一个烘焙生产方法,把生产的成品发送给ice方法
func baker(baked chan<- *Cake) {
	for {
		cake := new(Cake)
		cake.state = "baked"
		baked <- cake
	}
}

// 定义一个加奶油加工方法,把从bake接受的成品添加奶油后,发送给下一步
func icer(baked <-chan *Cake, iced chan<- *Cake) {
	for cake := range baked {
		cake.state = "iced"
		iced <- cake
	}
}
