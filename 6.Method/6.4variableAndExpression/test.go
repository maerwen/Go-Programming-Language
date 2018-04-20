package main

import "fmt"

func main() {
	// testVariable()
	testExpression(Point{2, 2}, Point{1, 1}, true)
}

type Point struct {
	X, Y int
}

func (p Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}
func (p Point) Sub(q Point) Point {
	return Point{p.X - q.X, p.Y - q.Y}
}
func testVariable() { //方法变量测试
	p := Point{2, 2}
	var add = p.Add //方法变量
	fmt.Println(add(Point{0, 0}))
}
func testExpression(p, q Point, judge bool) { //方法表达式测试
	var op func(p, q Point) Point
	if judge {
		op = Point.Add //方法表达式
	} else {
		op = Point.Sub
	}
	fmt.Println(op(p, q))
}
