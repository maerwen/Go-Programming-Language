package main

import (
	"fmt"
)

// 结构清晰,但访问很麻烦
type Point struct {
	X, Y int
}
type Circle1 struct {
	Point  Point
	Radius int
}
type Wheel1 struct {
	Circle Circle1
	Spokes int
}

// 匿名成员
type Circle2 struct {
	Point
	Radius int
}
type Wheel2 struct {
	Circle2
	Spokes int
}

func test() {
	// var w1 Wheel1
	var w2, w3 Wheel2
	// w1 = Wheel1{X8, 8, 5, 20}                      //不能以快捷方式初始化结构体
	// w2 = Wheel2{X: 8, Y: 8, Radius: 5, Spokes: 20} //不能以快捷方式初始化结构体
	w2 = Wheel2{Circle2{Point{8, 8}, 5}, 20}
	w3 = Wheel2{Circle2: Circle2{
		Point:  Point{X: 8, Y: 8},
		Radius: 5,
	}, Spokes: 20}
	fmt.Printf("%v\n%v\n", w2, w3)
}
